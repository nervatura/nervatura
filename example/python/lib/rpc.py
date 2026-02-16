import os
import grpc

from nervatura import store_pb2 as pb
from nervatura import store_pb2_grpc as api

import lib.utils as utils

def client():
  channel = grpc.insecure_channel("localhost:"+os.getenv("NT_GRPC_PORT"))
  return api.APIStub(channel)

def metadata(token):
  if token != "":
    return [("authorization", "Bearer "+token)]
  else:
    return [("x-api-key", os.getenv("NT_API_KEY"))]

def Database(options):
  try:
    response = client().Database(pb.RequestDatabase(
      alias=options["alias"],
      demo=options["demo"]
    ), metadata=metadata(""))
    result = utils.checkJson(response.data)
    return [result, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]

def CustomerUpdate(token, data):
  try:
    response = client().CustomerUpdate(pb.Customer(
      code=data["code"],
      customer_name=data["customer_name"]
    ), metadata=metadata(token))
    return [response, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]

def CustomerGet(token, options):
  try:
    response = client().CustomerGet(pb.RequestGet(
      code=options["code"]
    ), metadata=metadata(token))
    return [response, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]

def CustomerQuery(token, filters):
  try:
    response = client().CustomerQuery(pb.RequestQuery(
      filters=[pb.RequestQueryFilter(field=field, value=value) for field, value in filters.items()]
    ), metadata=metadata(token))
    result = response.data
    return [result, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]

def Delete(token, options):
  try:
    response = client().Delete(pb.RequestDelete(
      code=options["code"],
      model=pb.Model.Value(options["model"].upper())
    ), metadata=metadata(token))
    return [response, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]

def Function(token, options):
  try:
    response = client().Function(pb.RequestFunction(
      function=options["name"],
      args=options["values"]
    ), metadata=metadata(token))
    result = utils.checkJson(response.data)
    return [result, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]

def View(token, options):
  try:
    filters = [pb.RequestQueryFilter(field=filter["field"], value=filter["value"]) for filter in options["filters"]]
    response = client().View(pb.RequestView(
      name=pb.ViewName.Value(options["name"].upper()),
      filters=filters,
      limit=options["limit"],
    ), metadata=metadata(token))
    result = utils.checkJson(response.data)
    return [result, None]
  except grpc.RpcError as err:
    return [None, err.details()]
  except Exception as e:
    return [None, str(e)]