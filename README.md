## Heart Disease Risk Predictor
Project for the Open Source Programming course, to predict whether a person is susceptible to heart disease or not

## Technology Stack
- Golang: The primary programming language for building the application
- Python3: Utilized for implementing the decision tree logic
- Redis: Employed with celery to distribute tasks and manage queues
- MySQL: Used for database storage and retrieval
- Celery: Task queue management to enhance performance and scalability

## Installation and Setup
To get started with the Heart Disease Risk Predictor, follow these steps:

1. Clone the project repository and redirect to the cloned repo
```
git clone https://github.com/yourusername/HeartRiskPredictor.git
cd heart-disease-predictor
```
2. Install project dependencies:
```
go get -d ./...
```
3. Run the executable located in the executable folder depending on your operating system (Windows/Linux_x64)
- For Windows users, execute heart-predictor.exe
- For Linux (x64) users, run heart-predictor

## Prediction Engine:
1. Decision tree logic implemented
2. Using Celery with Redis to distribute tasks and maintain queue

## Rest Endpoints:
1. To validate user input
2. Make REST interface indepenent from ML code, for future purposes

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

## Contribution
If you'd like to contribute to the project, please follow these steps:
1. Fork the repository: Click the "Fork button on the repository page to create your own fork
2. Clone your fork: Clone your forked repository to your local development environment
3. Make Changes: Make your changes, commit them, then push to your fork
4. Create a Pull Request: Visit the repository on GitHub and create a new pull request with a detailed description of all changes.
