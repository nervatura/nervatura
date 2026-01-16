import os
import requests
import lib.utils as utils

service_url = "http://localhost:"+str(os.getenv("NT_HTTP_PORT"))+"/api/v6"
headers = {
  "Content-Type": "application/json; charset=utf-8"
}

def decodeResult(resultStr, code):
  result = utils.checkJson(resultStr)
  if code not in [200, 201, 204]:
    return [None, result]
  return [result, None]

def set_headers(token):
  if token != "":
    headers["Authorization"] = "Bearer "+token
  else:
    headers["X-API-KEY"] = os.getenv("NT_API_KEY")

def Get(token, path, query):
  set_headers(token)
  try:
    path = service_url+"/"+path
    query_str = "&".join([f"{k}={v}" for k, v in query.items()])
    if query_str != "":
      path += "?"+query_str
    response = requests.get(path, headers=headers)
    return decodeResult(response.text, response.status_code)
  except requests.ConnectionError:
    return [None, "Connection failure"]

def Post(token, path, data):
  set_headers(token)
  try:
    path = service_url+"/"+path
    response = requests.post(path, headers=headers, json=data)
    return decodeResult(response.text, response.status_code)
  except requests.ConnectionError:
    return [None, "Connection failure"]

def Put(token, path, data):
  set_headers(token)
  try:
    path = service_url+"/"+path
    response = requests.put(path, headers=headers, json=data)
    return decodeResult(response.text, response.status_code)
  except requests.ConnectionError:
    return [None, "Connection failure"]

def Delete(token, path):
  set_headers(token)
  try:
    path = service_url+"/"+path
    response = requests.delete(path, headers=headers)
    return decodeResult(response.text, response.status_code)
  except requests.ConnectionError:
    return [None, "Connection failure"]