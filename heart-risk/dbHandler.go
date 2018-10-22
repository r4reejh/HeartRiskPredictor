package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func dbInit() {
	var err error
	db, err = sql.Open("mysql", "r4reejh:RxWwDI2UQCI8yWkf@/HeartRiskPredictor?charset=utf8&parseTime=true")
	fmt.Println("database initialized")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// DBcheckUserValid check if user is valid
func DBcheckUserValid(u string, p string) (status bool) {
	var password string
	row := db.QueryRow("select password from user where username=?", u)
	switch err := row.Scan(&password); err {
	case sql.ErrNoRows:
		return false
	case nil:
		// match password
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(p))
		// error !nil if password doesnt match
		if err != nil {
			return false
		}
		return true
	default:
		fmt.Println(err)
		panic(err)
	}
}

// DBfindUser check if user exists
func DBfindUser(uname string, email string) (status bool) {
	var u string
	row := db.QueryRow("select username from user where username=? OR email=?", uname, email)
	switch err := row.Scan(&u); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}
}

// DBfetchName return user's name
func DBfetchName(uname string) (name string) {
	var u string
	row := db.QueryRow("select name from user where username=?", uname)
	switch err := row.Scan(&u); err {
	case sql.ErrNoRows:
		return ""
	case nil:
		return u
	default:
		panic(err)
	}
}

// DBCreateUser to create a new validated user
func DBCreateUser(US User) {
	stmt, err := db.Prepare("insert user set username=?, email=?, password=?, name=?")
	checkErr(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(US.Password), bcrypt.DefaultCost)
	checkErr(err)

	_, err = stmt.Exec(US.Username, US.Email, string(hashedPassword), US.Name)
	checkErr(err)
	return
}

// DBfetchUserHistory to fetch array of historical scans for a user
func DBfetchUserHistory(username string) (History []Scan) {
	query := "select * from scan where username=? order by createdat desc"
	rows, err := db.Query(query, username)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		T := Scan{}
		err = rows.Scan(&T.ID,
			&T.UserID,
			&T.Username,
			&T.Age,
			&T.Sex,
			&T.Cp,
			&T.Trestbps,
			&T.Chol,
			&T.Fbs,
			&T.Restecg,
			&T.Thalach,
			&T.Exang,
			&T.Oldpeak,
			&T.Slope,
			&T.Ca,
			&T.Thal,
			&T.CreatedAt,
			&T.Result)

		History = append(History, T)
	}
	return
}

// DBrecordScan insert scans to DB
func DBrecordScan(T InputStruct, x float64) {

	username, err := redis.String(cache.Do("GET", T.Token))
	checkErr(err)

	statement := `
			insert scan set userid=?,
			username=?,
			age=?, 
			sex=?, 
			cp=?, 
			trestbps=?, 
			Chol=?, 
			fbs=?, 
			Restecg=?,
			Thalach=?,
			Exang=?,
			Oldpeak=?,
			Slope=?,
			Ca=?,
			Thal=?,
			Result=?`

	stmt, err := db.Prepare(statement)
	checkErr(err)

	_, err = stmt.Exec(T.UserID,
		username,
		T.Age,
		T.Sex,
		T.Cp,
		T.Trestbps,
		T.Chol,
		T.Fbs,
		T.Restecg,
		T.Thalach,
		T.Exang,
		T.Oldpeak,
		T.Slope,
		T.Ca,
		T.Thal,
		x)

	checkErr(err)
}

/*func rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if row exists '%s' %v", args, err)
	}
	return exists
}*/
