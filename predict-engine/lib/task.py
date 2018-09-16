import sys
from celery import Celery
from modelLoaderV2 import ValuePredictor as VP

app = Celery('tasks',backend='redis://localhost:6379', broker='redis://localhost:6379')

@app.task
def add(ll):
    x = VP(ll)
    return x