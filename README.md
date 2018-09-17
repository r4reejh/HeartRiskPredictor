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
	"age":"57",
	"sex":"0",
	"cp":"1",	
	"trestbps":"136",	
	"chol":"236",	
	"fbs":"0",	
	"restecg":"0",	
	"thalach":"174",	
	"exang":"0",	
	"oldpeak":"0",	
	"slope":"1",	
	"ca":"1"	,
	"thal":"2"
}
```

__Sample Response__
```json
{
  "Label":1
}
```

## UI sample

![alt text](https://raw.githubusercontent.com/r4reejh/OSP-heart-disease-risk-predictor/master/sample1.png)
![alt text](https://raw.githubusercontent.com/r4reejh/OSP-heart-disease-risk-predictor/master/sample2.png)
