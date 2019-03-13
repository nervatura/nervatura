#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2019, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
"""

import sys

class Pylib(object):
 
  def __init__(self, argv):
    self.params = argv[2:]
  
  def check_python(self):
    print(sys.version_info.major)
  
  def get_source(self):
    import inspect
    func_ = getattr(self, self.params[0], None)
    if callable(func_):
      print(inspect.getsource(func_))
  
  def read_file(self, filename, mode='r'):
    f = open(filename, mode)
    try:
      return f.read()
    finally:
      f.close()
            
  def create_excel(self):
    from xlwt import easyxf
    from xlwt import Workbook
    from StringIO import StringIO
    import json
    import binascii
    
    styles = {
      'header': easyxf(
        'font: bold true, height 160;'
        'alignment: horizontal left, vertical center;'
        'pattern: back_colour gray25;'
        'borders: left thin, right thin, top thin, bottom thin;'),
      'float': easyxf(
        'alignment: horizontal right, vertical center;'
        'borders: left thin, right thin, top thin, bottom thin;',
        num_format_str='# ### ##0.00'),
      'integer': easyxf(
        'alignment: horizontal right, vertical center;'
        'borders: left thin, right thin, top thin, bottom thin;',
        num_format_str='# ### ##0'),
      'date': easyxf(
        'alignment: horizontal center, vertical center;'
        'borders: left thin, right thin, top thin, bottom thin;',
        num_format_str='yyyy.mm.dd'),
      'bool': easyxf(
        'alignment: horizontal center, vertical center;'
        'borders: left thin, right thin, top thin, bottom thin;'),
      'string': easyxf(
        'alignment: horizontal left, vertical center;'
        'borders: left thin, right thin, top thin, bottom thin;')}
    
    if len(self.params)>0:
      params = json.loads(self.params[0])
    
    book = Workbook(encoding='utf-8')
    
    for skey in params["data"].keys():
      if type(params["data"][skey]).__name__=="list":
        if len(params["data"][skey])>0:
          sheetName = skey
          columns = []
          if params["template"].has_key(skey):
            if params["template"][skey].has_key("sheetName"):
              sheetName = params["template"][skey]["sheetName"]
            if params["template"][skey].has_key("columns"):
              columns = params["template"][skey]["columns"]
          else:
            for colname in params["data"][skey][0].keys():
              columns.append({"name":colname,"label":colname,"type":"string"})
          if params["data"]["labels"]:
            for col in columns:
              if params["data"]["labels"].has_key(col["name"]):
                col["label"] = params["data"]["labels"][col["name"]]
          
          sheet = book.add_sheet(sheetName)     
          colnum = 0;
          for col in columns:
            sheet.write(0, colnum, col["label"], styles["header"])
            colnum = colnum + 1
          rownum = 1  
          for row in params["data"][skey]:
            colnum = 0
            for col in columns:
              if col["type"]=="float":
                sheet.write(rownum, colnum, row[col["name"]], styles["float"])
              if col["type"]=="integer":
                sheet.write(rownum, colnum, row[col["name"]], styles["integer"])
              if col["type"]=="date":
                sheet.write(rownum, colnum, row[col["name"]], styles["date"])
              if col["type"]=="bool":
                sheet.write(rownum, colnum, row[col["name"]], styles["bool"])
              if col["type"]=="string":
                sheet.write(rownum, colnum, row[col["name"]], styles["string"])
              colnum = colnum + 1
            rownum = rownum + 1
    
    output = StringIO()
    book.save(output)
    contents = output.getvalue()
    output.close
    sys.stdout.write(binascii.b2a_base64(contents))

if len(sys.argv)>1:
  pl = Pylib(sys.argv)
  func_ = getattr(pl, sys.argv[1], None)
  if callable(func_):
    try:
      func_()
    except Exception as err:
      raise RuntimeError(err)
  else:
    print("Invalid function:"+sys.argv[1])
else:
  print("Missing function name!")
  