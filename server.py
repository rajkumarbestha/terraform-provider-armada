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
    if request.method == 'GET':
        print('In POST method.')
        #ome = request.json
        print(request.args.get('service')+' Service called from Terraform.')
        print('AccessKey from terraform is '+request.args.get('access_key'))
        
        longstr = request.args.get('service')+"_req201908131127"
        print(longstr+' request ID sent from server.')
        sys.stdout.flush()
        return jsonify({'requestID':longstr})

if __name__ == "__main__":
    app.run(debug=True)