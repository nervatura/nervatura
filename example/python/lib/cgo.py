import ctypes

import lib.utils as utils

ntura = ctypes.cdll.LoadLibrary("dist/nervatura.so")

def decodeResult(resultStr):
    result = utils.checkJson(resultStr)
    if type(result) == dict:
        if "code" in result:
            if result["code"] not in [200, 201, 204]:
                return [None, result["data"]]
    return [result, None]

def Database(options):
    Database = ntura.Database
    Database.argtypes = [ctypes.c_char_p]
    Database.restype = ctypes.c_void_p
    result = ctypes.string_at(Database(utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)

def Function(options):
    Function = ntura.Function
    Function.argtypes = [ctypes.c_char_p]
    Function.restype = ctypes.c_void_p
    result = ctypes.string_at(Function(utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)

def ResetPassword(options):
    ResetPassword = ntura.ResetPassword
    ResetPassword.argtypes = [ctypes.c_char_p]
    ResetPassword.restype = ctypes.c_void_p
    result = ctypes.string_at(ResetPassword(utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)

def Create(model, options, data):
    Create = ntura.Create
    Create.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
    Create.restype = ctypes.c_void_p
    result = ctypes.string_at(Create(model.encode("utf-8"), utils.encodeOptions(options), utils.encodeOptions(data))).decode("utf-8")
    return decodeResult(result)

def Update(model, options, data):
    Update = ntura.Update
    Update.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
    Update.restype = ctypes.c_void_p
    result = ctypes.string_at(Update(model.encode("utf-8"), utils.encodeOptions(options), utils.encodeOptions(data))).decode("utf-8")
    return decodeResult(result)

def Delete(model, options):
    Delete = ntura.Delete
    Delete.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
    Delete.restype = ctypes.c_void_p
    result = ctypes.string_at(Delete(model.encode("utf-8"), utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)

def Get(model, options):
    Get = ntura.Get
    Get.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
    Get.restype = ctypes.c_void_p
    result = ctypes.string_at(Get(model.encode("utf-8"), utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)

def Query(model, options):
    Query = ntura.Query
    Query.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
    Query.restype = ctypes.c_void_p
    result = ctypes.string_at(Query(model.encode("utf-8"), utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)

def View(options):
    View = ntura.View
    View.argtypes = [ctypes.c_char_p]
    View.restype = ctypes.c_void_p
    result = ctypes.string_at(View(utils.encodeOptions(options))).decode("utf-8")
    return decodeResult(result)
