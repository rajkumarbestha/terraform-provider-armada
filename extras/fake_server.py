# -*- coding: utf-8 -*-
"""
Created on Tue Aug 13 23:09:36 2019

@author: Rajkumar
"""
import sys
import json
from flask import Flask, request, jsonify
app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def hello_world():
    if request.method == 'POST':
        print('In POST method.')
        values = request.json
        print(values)
        sys.stdout.flush()
        longstr = "23eduor8943tgruobg"
    return jsonify({'RequestIDOut':longstr})

if __name__ == "__main__":
    app.run(debug=True)
