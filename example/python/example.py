
from lib import cgo

api_map = {
  "cgo": cgo
  # , "cli": cli, "rpc": rpc, "http": rest
}
fn_map = {
  "cgo": {
    "Database": [{"alias": "demo", "demo": True}],
    "ResetPassword": [{"alias": "demo", "code": "USR0000000000N1"}],
    "Create": ["customer", {"alias": "demo"}, {"code": "PYT0000000000N1", "customer_name": "Python Test"}],
    "Update": ["customer", {"alias": "demo", "code": "PYT0000000000N1"}, {"customer_meta": {"account": "1234567890"}}],
    "Get": ["customer", {"alias": "demo", "code": "PYT0000000000N1"}],
    "Query": ["customer", {"alias": "demo", "customer_type": "CUSTOMER_COMPANY"}],
    "View": [{"alias": "demo", "name": "VIEW_CUSTOMER_EVENTS", "filter":"subject like '%visit%' and place='City1'", "limit":10}],
    "Delete": ["customer", {"alias": "demo", "code": "PYT0000000000N1"}],
  }
}
for api_key in api_map:
  for fn_key in fn_map[api_key]:
    result, err = api_map[api_key].__dict__[fn_key](*fn_map[api_key][fn_key])
    if err:
      print(api_key+" "+fn_key+" error: "+err)
    else:
      print(api_key+" "+fn_key+" OK")
