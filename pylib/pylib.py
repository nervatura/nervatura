#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2017, Csaba Kappel
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
  
  def create_report(self):
    from report.report import Report
    import json

    if len(self.params)>0:
      params = json.loads(self.params[0])

    if not "orientation" in params:
      params["orientation"] = "p"
    if not "size" in params:
      params["size"] = "a4"
    if not "output" in params:
      params["output"] = "pdf"
    rpt = Report(orientation=params["orientation"],format=params["size"])

    if "data" in params:
      rpt.databind = params["data"]
    if "template" in params:
      rpt.loadDefinition(params["template"])
    rpt.createReport()
    if params["output"] == "pdf" or params["output"] == "base64":
      import binascii
      sys.stdout.write(binascii.b2a_base64(rpt.save2Pdf()))
    elif output == "xml":
      sys.stdout.write(rpt.save2Xml())
    else:
      sys.stdout.write(rpt.save2Html())
  
  def create_report_sample(self):
    from report.report import Report

    orient = "p"
    if len(self.params)>0:
      orient = self.params[0]
    output = "pdf"  
    if len(self.params)>1:
      output = self.params[1]

    rpt = Report(orient)
    #default values
    rpt.template["document"]["title"] = "Nervatura Report"
    rpt.template["margins"]["left-margin"] = 15
    rpt.template["margins"]["top-margin"] = 15
    rpt.template["margins"]["right-margin"] = 15
    rpt.template["margins"]["bottom-margin"] = 15
    rpt.template["style"]["font-family"] = "times"

    #header
    header = rpt.template["elements"]["header"]
    row_data = rpt.insertElement(header, "row", -1, {"height": 10})
    rpt.insertElement(row_data, "image",-1,{"src":"logo"})
    rpt.insertElement(row_data, "cell",-1,{
      "name":"label", "value":"labels.title", "font-style": "bolditalic", "font-size": 26, "color": "#D8DBDA"})
    rpt.insertElement(row_data, "cell",-1,{
      "name":"label", "value":"Python Sample", "font-style": "bold", "align": "right"})
    rpt.insertElement(header, "vgap", -1, {"height": 2})
    rpt.insertElement(header, "hline", -1, {"border-color": 218})
    rpt.insertElement(header, "vgap", -1, {"height": 2})

    #details
    details = rpt.template["elements"]["details"]
    rpt.insertElement(details, "vgap", -1, {"height": 2})
    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "width": "50%", "font-style": "bold", "value": "labels.left_text", "border": "LT", 
      "border-color": 218, "background-color": 245})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LTR", 
      "border-color": 218, "background-color": 245})

    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "width": "50%", "value": "head.short_text", "border": "L", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "value": "head.short_text", "border": "LR", "border-color": 218})
    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "width": "50%", "value": "head.short_text", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "value": "head.short_text", "border": "LBR", "border-color": 218})

    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "width": "40", "font-style": "bold", "value": "labels.left_text", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "align": "center", "width": "30", "font-style": "bold", "value": "labels.center_text", 
      "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "align": "right", "width": "40", "font-style": "bold", "value": "labels.right_text", 
      "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LBR", "border-color": 218})

    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "width": "40", "value": "head.short_text", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "date", "align": "center", "width": "30", "value": "head.date", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "amount", "align": "right", "width": "40", "value": "head.number", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "value": "head.short_text", "border": "LBR", "border-color": 218})

    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "width": "50", "value": "head.short_text", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "font-style": "bold", "value": "labels.left_text", "border": "LB", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "value": "head.short_text", "border": "LBR", "border-color": 218})

    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "long_text", "multiline": "true", "value": "head.long_text", "border": "LBR", "border-color": 218})

    rpt.insertElement(details, "vgap", -1, {"height": 2})
    row_data = rpt.insertElement(details, "row", -1, {"hgap": 2})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "labels.left_text", "font-style": "bold", "border": "1", "border-color": 245, 
      "background-color": 245})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "value": "head.short_text", "border": "1", "border-color": 218})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "labels.left_text", "font-style": "bold", "border": "1", "border-color": 245, "background-color": 245})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "short_text", "value": "head.short_text", "border": "1", "border-color": 218})

    rpt.insertElement(details, "vgap", -1, {"height": 2})
    row_data = rpt.insertElement(details, "row", -1, {"hgap": 2})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "labels.long_text", "font-style": "bold", "border": "1", "border-color": 245, "background-color": 245})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "long_text", "multiline": "true", "value": "head.long_text", "border": "1", "border-color": 218})

    rpt.insertElement(details, "vgap", -1, {"height": 2})
    rpt.insertElement(details, "hline", -1, {"border-color": 218})
    rpt.insertElement(details, "vgap", -1, {"height": 2})

    row_data = rpt.insertElement(details, "row", -1, {"hgap": 3})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "Barcode (Interleaved 2of5)", "font-style": "bold", "font-size": 10,
      "border": "1", "border-color": 245, "background-color": 245})
    rpt.insertElement(row_data, "barcode",-1,{"code-type": "ITF", "value": "1234567890", "visible-value":1})
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "Barcode (Code 39)", "font-style": "bold", "font-size": 10, 
      "border": "1", "border-color": 245, "background-color": 245})
    rpt.insertElement(row_data, "barcode",-1,{"code-type": "CODE_39", "value": "1234567890ABCDEF", "visible-value":1})

    rpt.insertElement(details, "vgap", -1, {"height": 3})
      
    row_data = rpt.insertElement(details, "row")
    rpt.insertElement(row_data, "cell",-1,{
      "name": "label", "value": "Datagrid Sample", "align": "center", "font-style": "bold", 
      "border": "1", "border-color": 245, "background-color": 245})
    rpt.insertElement(details, "vgap", -1, {"height": 2})

    grid_data = rpt.insertElement(details, "datagrid", -1, {
      "name": "items", "databind": "items", "border": "1", "border-color": 218, "header-background": 245, "footer-background": 245})
    rpt.insertElement(grid_data, "column",-1,{
      "width": "8%", "fieldname": "counter", "align": "right", "label": "labels.counter", "footer": "labels.total"})
    rpt.insertElement(grid_data, "column",-1,{
      "width": "20%", "fieldname": "date", "align": "center", "label": "labels.center_text"})
    rpt.insertElement(grid_data, "column",-1,{
      "width": "15%", "fieldname": "number", "align": "right", "label": "labels.right_text", 
      "footer": "items_footer.items_total", "footer-align": "right"})
    rpt.insertElement(grid_data, "column",-1,{
      "fieldname": "text", "label": "labels.left_text"})

    rpt.insertElement(details, "vgap", -1, {"height": 5})
    rpt.insertElement(details, "html", -1, {"fieldname": "html_text", 
      "html": "<i>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</i> ={{html_text}} Nulla a pretium nunc, in cursus quam."})

    #footer
    footer = rpt.template["elements"]["footer"]
    rpt.insertElement(footer, "vgap", -1, {"height": 2})
    rpt.insertElement(footer, "hline", -1, {"border-color": 218})
    row_data = rpt.insertElement(footer, "row", -1, {"height": 10})
    rpt.insertElement(row_data, "cell",-1,{"value": "Nervatura Report Template", "font-style": "bolditalic"})
    rpt.insertElement(row_data, "cell",-1,{"value": "{{page}}", "align": "right", "font-style": "bold"})
              
    #data
    rpt.setData("labels", {"title": "REPORT TEMPLATE", "left_text": "Short text", "center_text": "Centered text", 
                                          "right_text": "Right text", "long_text": "Long text", "counter": "No.", "total": "Total"})
    rpt.setData("head", {"short_text": "Lorem ipsum dolor", "number": "123 456", "date": "2015.01.01", 
                          "long_text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim. Nulla a pretium nunc, in cursus quam."})
    rpt.setData("html_text", "<p><b>Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim.</b></p>")
    rpt.setData("items_footer", {"items_total": "3 703 680"})
    items = []
    for i in range(1, 30):
      items.append({"text": str(i)+". Lorem ipsum dolor", "number": "123 456", "date": "2015.01.01"})
    rpt.setData("items", items)
    rpt.setData("logo", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABoAAAAaCAYAAACpSkzOAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAFxGAABcRgEUlENBAAAAB3RJTUUH4AcCFiIAfaA8WwAAAeBJREFUSMftlj9IVlEUwH/nfR/+G3QwaPlAg3D5pKkPDIoIghYVm4yCFKlPcGoQdGppa4mm4IJbQdFYRKDgoFR0o5ZwcL5LLoKCoPW+43JfvO/6ns9P0ckDj8c5l3t+55x73rkPzqUFqdRrTe/Tho0CRGeQ2FWA8ilmcgH4CUyeGFSp13DGpvUOoB+YSwDALwBp0VF6jwAKXAbGgIfAYOBi1Rl7oxDknXf56IaBKtDhnzagvcBHzRn7IxOUZFGp12aAaeDKMSv7D+gBdpyxzSBfmnZgGbh2wn5Yc8ZWEyXKOI+lY0C2RGQ9sD1OK/9BvlyjwPUWIZsi8lFVB1K2DWfsl/REKAdle94i5LuIbKnq/cC+mAR/ICMvvUcliMgroFNVb2csT4WGcoGeJbsisqCq4zmBLThj90JjmNFeQbt+i0ReqOpMDiR2xj7KmtYh6E/WZhH5WoqiCWCooTp/SDC3wrPJ7Dr/9Sdn8L5UKt0T5JOqVuNG401BSaedsSt5d08ULFwENoBNVb0Zx/FbRUeA7gLIU2esCWdjU/P41h4HZv3d8Q74DTw7QmP8Baacsa8Pg6S7bBvoSw9BP0yfAJ05gM/AXSDOO5cDGeVdC972ALgDXAJ2gQ/O2Jfnf0RnJvumbKT0gnMTFgAAAABJRU5ErkJggg==")

    rpt.createReport()
    if output == "pdf":
      import binascii
      sys.stdout.write(binascii.b2a_base64(rpt.save2Pdf()))
    elif output == "xml":
      sys.stdout.write(rpt.save2Xml())
    else:
      sys.stdout.write(rpt.save2Html())
          
  def load_report_xml(self):
    from report.report import Report
    
    orient = "p"
    if len(self.params)>0:
      orient = self.params[0]
    output = "pdf"  
    if len(self.params)>1:
      output = self.params[1]
    deffile = self.params[2]

    rpt = Report(orient)
    xdata = self.read_file(deffile)
    rpt.loadDefinition(xdata)
    rpt.createReport()
    if output == "pdf":
      import binascii
      sys.stdout.write(binascii.b2a_base64(rpt.save2Pdf()))
    elif output == "xml":
      sys.stdout.write(rpt.save2Xml())
    else:
      sys.stdout.write(rpt.save2Html())

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
  