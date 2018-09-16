from setuptools import setup
from setuptools.command.install import install
import os    


setup(name='heart-risk',
    version='0.1.0',
    description='a service to predict risk of heart disease',
    author='Reejh Ghosh, Ritvik Khanna',
    license='MIT',
    packages=['lib'],
    install_requires=[
        'sklearn',
        'pandas',
        'numpy',
        'celery'
    ],
    zip_safe=False)