## OSP-heart-disease-risk-predictor
Project for the Open Source Programming course, to predict whether a person is susceptible to heart disease or not

## Prediction Engine:
1. Decision tree logic implemented
2. Using Celery with Redis to distribute tasks and maintain queue

## Rest Endpoints:
1. To validate user input
2. Make REST interface indepenent from ML code, for future purposes
3. Made using golang

### Sample Request

__Sample Request Body__
```json
{
	"name":"Reejh Ghosh",
	"age":43.0,
	"sex":1,
	"cp":0,	
	"trestbps":120,	
	"chol":177,	
	"fbs":0,	
	"restecg":0,	
	"thalach":120,	
	"exang":1,	
	"oldpeak":2.5,	
	"slope":1,	
	"ca":0	,
	"thal":3
}
```

__Sample Response__
```json
{
    "Name": "Reejh Ghosh",
    "Label": 0,
    "Date": "2018-09-17T10:08:31.679127494+05:30"
}
```

## UI sample

![alt text](https://raw.githubusercontent.com/r4reejh/OSP-heart-disease-risk-predictor/master/sample1.png)
![alt text](https://raw.githubusercontent.com/r4reejh/OSP-heart-disease-risk-predictor/master/sample2.png)
