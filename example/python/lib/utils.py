import json

def encodeOptions(data):
  return json.dumps(data).encode("utf-8")

def checkJson(data):
  try:
    jdata = json.loads(data)
  except ValueError as err:
    return data
  return jdata