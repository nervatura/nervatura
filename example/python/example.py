import os

from dotenv import load_dotenv
load_dotenv(dotenv_path=".env.example")

from lib import cgo, cli, rest, rpc

api_map = {
  "cgo": cgo, "cli": cli, "rest": rest, "rpc": rpc
}
fn_map = {
   "cgo": {
     "Database": [{"alias": "demo", "demo": True}],
     "Function": [{"name": "test", "values": {}}],
     "ResetPassword": [{"alias": "demo", "code": "USR0000000000N1"}],
     "Create": ["customer", {"alias": "demo"}, {"code": "CGO0000000000N1", "customer_name": "Python Test"}],
     "Update": ["customer", {"alias": "demo", "code": "CGO0000000000N1"}, {"customer_meta": {"account": "1234567890"}}],
     "Get": ["customer", {"alias": "demo", "code": "CGO0000000000N1"}],
     "Query": ["customer", {"alias": "demo", "customer_type": "CUSTOMER_COMPANY"}],
     "View": [{"alias": "demo", "name": "VIEW_CUSTOMER_EVENTS", 
       "filters": [{"field":"like_subject","value":"visit"}, {"field":"place","value":"City1"}], 
       "limit":10}],
     "Delete": ["customer", {"alias": "demo", "code": "CGO0000000000N1"}],
   },
  "cli": {
#    "Database": [{"alias": "demo", "demo": True}],
    "Function": [{"name": "test", "values": {}}],
    "ResetPassword": [{"alias": "demo", "code": "USR0000000000N1"}],
    "Create": ["customer", {"alias": "demo"}, {"code": "CLI0000000000N1", "customer_name": "Python Test"}],
    "Update": ["customer", {"alias": "demo", "code": "CLI0000000000N1"}, {"customer_meta": {"account": "1234567890"}}],
    "Get": ["customer", {"alias": "demo", "code": "CLI0000000000N1"}],
    "Query": ["customer", {"alias": "demo", "customer_type": "CUSTOMER_COMPANY"}],
    "View": [{"alias": "demo", "name": "VIEW_CUSTOMER_EVENTS", "filters": [{"field":"like_subject","value":"visit"}, {"field":"place","value":"City1"}], "limit":10}],
    "Delete": ["customer", {"alias": "demo", "code": "CLI0000000000N1"}],
  },
   "rest": {
    "Post": ["", "customer", {"code": "REST0000000000N1", "customer_name": "Python Test"}],
    "Put": ["", "customer/REST0000000000N1", {"customer_name": "Test Python"}],
    "Get": ["", "customer", {"customer_type": "CUSTOMER_COMPANY"}],
    "Delete": ["", "customer/REST0000000000N1"],
   },
   "rpc": {
    #"Database": [{"alias": "demo", "demo": True}],
    "CustomerUpdate": ["", {"code": "RPC0000000000N1", "customer_name": "Python Test"}],
    "CustomerGet": ["", {"code": "RPC0000000000N1"}],
    "CustomerQuery": ["", {"customer_type": "CUSTOMER_COMPANY"}],
    "Delete": ["", {"code": "RPC0000000000N1", "model": "customer"}],
    "Function": ["", {"name": "product_price", "values": {"product_code":"PRD0000000000N1", "currency_code":"EUR", "price_type":"PRICE_CUSTOMER"}}],
    "View": ["", {"name": "VIEW_CUSTOMER_EVENTS", "filters": [{"field":"like_subject","value":"visit"}, {"field":"place","value":"City1"}], "limit":10}],
   }
}
for api_key in api_map:
  for fn_key in fn_map[api_key]:
    result, err = api_map[api_key].__dict__[fn_key](*fn_map[api_key][fn_key])
    if err:
      print(api_key+" "+fn_key+" error: "+str(err))
    else:
      print(api_key+" "+fn_key+" OK")
