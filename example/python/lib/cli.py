import subprocess
import os

import lib.utils as utils

service_path = os.getenv("NT_SERVICE_PATH")

def decodeResult(output):
    err = output.stderr.decode()
    if err != "":
        return [None, err]
    result = output.stdout.decode()
    result = utils.checkJson(result.split("\n")[len(result.split("\n"))-2])
    if type(result) == dict:
        if "code" in result:
            if result["code"] not in [200, 201, 204]:
                return [None, result["data"]]
    return [result, None]

def Database(options):
    output = subprocess.run([service_path, "-c", "database", "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)

def ResetPassword(options):
    output = subprocess.run([service_path, "-c", "reset", "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)

def Create(model, options, data):
    output = subprocess.run([service_path, "-c", "create", 
    "-m", model, "-o", utils.encodeOptions(options), "-d", utils.encodeOptions(data)], capture_output=True)
    return decodeResult(output)

def Update(model, options, data):
    output = subprocess.run([service_path, "-c", "update", 
    "-m", model, "-o", utils.encodeOptions(options), "-d", utils.encodeOptions(data)], capture_output=True)
    return decodeResult(output)

def Delete(model, options):
    output = subprocess.run([service_path, "-c", "delete", 
    "-m", model, "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)

def Get(model, options):
    output = subprocess.run([service_path, "-c", "get", 
    "-m", model, "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)

def Query(model, options):
    output = subprocess.run([service_path, "-c", "query", 
    "-m", model, "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)

def View(options):
    output = subprocess.run([service_path, "-c", "view", "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)

def Function(options):
    output = subprocess.run([service_path, "-c", "function", "-o", utils.encodeOptions(options)], capture_output=True)
    return decodeResult(output)
