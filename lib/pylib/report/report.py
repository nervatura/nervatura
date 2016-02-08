# -*- coding: utf-8 -*-

"""
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
"""

from fpdf.fpdf import FPDF
import xml.etree.ElementTree as et
from fpdf_html import HTMLMixin

from io import BytesIO
try:
  try:
    import Image
  except:
    from PIL import Image
except ImportError:
  Image = None
                
class Report(FPDF, HTMLMixin):
  
  CLASS_VERSION='2.141007'
  template,databind = {},{}
  images_folder=None
  
  def __init__(self,orientation='P',unit='mm',format='A4'):  # @ReservedAssignment
    self.template = {
      "document":{
        "title":"Nervatura Report", "author":"", "creator":"", "subject":"", "keywords":""},
      "style":{
        "font-family":"times", "font-style":"", "font-size":12, "color":0, "border-color":0, "background-color":255},
      "margins":{
        "left-margin": 12, "top-margin": 12, "right-margin": 12, "bottom-margin": 12},
      "elements": {
        "header":et.Element("header") , "details":et.Element("details"), "footer":et.Element("footer")},
      "xml": {
        "header":"", "details":""},
      "html": {
        "header":"", "details":"", "footer":""},
      "page": {
        "favicon":"/nerva2py/static/favicon.ico",
        "footer_start":None, "row_height":7, "decode":"utf-8", "encode":"latin_1"}
      }
    self.databind = {}
    self.images_folder=None
    
    self.orientation = self.parse_value("orientation", orientation)
    self.unit = self.parse_value("unit", unit)
    FPDF.__init__(self,orientation, self.unit, self.parse_value("format", format))
    self.set_style(self.template["style"])
    self.alias_nb_pages('{{pages}}')
    
  def header(self):
    if self.template["elements"]["header"]!=None:
      ehtml=""
      for element in self.template["elements"]["header"]:
        ehtml+=self.show_element("header",element)
      if self.page_no()==1:
        self.template["html"]["header"]=ehtml.replace("{{pages}}", "1")
  
  def footer(self):
    if self.template["elements"]["footer"]!=None:
      ehtml=""
      if self.page_no()==1:
        if self.y < self.page_break_trigger-5:
          self.set_y(self.page_break_trigger-5)
        self.template["page"]["footer_start"] = self.get_y()
      elif self.template["page"]["footer_start"]!=self.get_y():
        self.set_y(self.template["page"]["footer_start"])
      for element in self.template["elements"]["footer"]:
        ehtml+=self.show_element("footer",element)
      if self.page_no()==1:
        self.template["html"]["footer"]=ehtml.replace("{{pages}}", "1")
    
  def show_element(self, section, element):
    if element.tag =="row":
      return self.create_row(section,element)
    elif element.tag =="vgap":
      height = element.get("height","")
      if height!="":
        height=float(height)
      self.ln(h=height)
      return '<div style="height:'+str(height)+self.unit+';">&nbsp;</div>'
    elif element.tag =="hline":
      width = element.get("width",0)
      if width!="":
        width=float(width)
      if width==0:
        style="width:100%;"
      else:
        style='width:'+str(width)+self.unit+';'
      gap = element.get("gap",0)
      if gap!="":
        gap=float(gap)
      if element.get("border-color",None)!=None:
        style+="border-color:"+hex(int(element.get("border-color",0))).replace("0x","#").upper()+" !important;"
        self.set_draw_color(*self.rgb(int(element.get("border-color",0))))
      else:
        self.set_draw_color(*self.rgb(int(self.template["style"]["border-color"])))
      self.cell(w=width, h=gap, txt="", border="TB", ln=1)
      return '<hr style="'+style+'margin:3px;"/>'
    elif element.tag =="html":
      html = element.text
      fieldname=element.get("fieldname","head")
      html = self.setHtmlValue(html,fieldname)
      x = self.get_x()
      self.set_style(element.attrib)
      font = element.get("font-family",self.template["style"]["font-family"])
      fontsize=int(element.get("font-size",self.template["style"]["font-size"]))
      self.write_html(html.decode(self.template["page"]["decode"]).encode(self.template["page"]["encode"],'replace'),font=font,fontsize=fontsize)
      self.set_x(x)
      self.ln()
      return html
    elif element.tag =="datagrid":
      datagrid_ = self.create_datagrid(element)
      datagrid = datagrid_[0]
      if datagrid!="":
        self.lasth=0
        x = self.get_x()
        if datagrid_[1]!=1:
          self.ln(-self.font_size_pt/72.0*25.4*1.5)
        font = element.get("font-family",self.template["style"]["font-family"])
        fontsize=int(element.get("font-size",self.template["style"]["font-size"]))
        self.write_html(datagrid.decode(self.template["page"]["decode"]).encode(self.template["page"]["encode"],'replace'),font=font,fontsize=fontsize)
        self.set_x(x)
        self.ln(-self.font_size_pt/72.0*25.4*1.5)
      return datagrid
    return ""
  
  def rgb(self,col):
    return (col // 65536), (col // 256 % 256), (col% 256)
  
  def set_style(self,attr):
    style=""
    if attr.has_key("font-style"):
      if str(attr["font-style"]).find("B")>-1:
        style+="font-weight:bold;"
      if str(attr["font-style"]).find("I")>-1:
        style+="font-style:italic;"
      if str(attr["font-style"]).find("U")>-1:
        style+="text-decoration:underline;"
    if attr.has_key("font-family"):
      family=attr["font-family"]
      style+="font-family:"+attr["font-family"]+";"
    else:
      style+="font-family:"+self.template["style"]["font-family"]+";"
      family=self.template["style"]["font-family"]
    if attr.has_key("font-style"):
      fstyle=attr["font-style"]
    else:
      fstyle=self.template["style"]["font-style"]
    self.set_font(family=family, style=fstyle)  
    if attr.has_key("font-size"):
      style+="font-size:"+str(attr["font-size"])+"px;"
      self.set_font_size(int(attr["font-size"]))
    else:
      style+="font-size:"+str(self.template["style"]["font-size"])+"px;"
      self.set_font_size(int(self.template["style"]["font-size"]))
    if attr.has_key("color"):
      style+="color:"+hex(int(attr["color"])).replace("0x","#").upper()+";"
      self.set_text_color(*self.rgb(int(attr["color"])))
    else:
      style+="color:"+hex(int(self.template["style"]["color"])).replace("0x","#").upper()+";"
      self.set_text_color(*self.rgb(int(self.template["style"]["color"])))
    if attr.has_key("border-color"):
      style+="border-color:"+hex(int(attr["border-color"])).replace("0x","#").upper()+" !important;"
      self.set_draw_color(*self.rgb(int(attr["border-color"])))
    else:
      style+="border-color:"+hex(int(self.template["style"]["border-color"])).replace("0x","#").upper()+" !important;"
      self.set_draw_color(*self.rgb(int(self.template["style"]["border-color"])))
    if attr.has_key("background-color"):
      style+="background-color:"+hex(int(attr["background-color"])).replace("0x","#").upper()+";"
      self.set_fill_color(*self.rgb(int(attr["background-color"])))
    else:
      style+="background-color:"+hex(int(self.template["style"]["background-color"])).replace("0x","#").upper()+";"
      self.set_fill_color(*self.rgb(int(self.template["style"]["background-color"])))
    return style
  
  def setValue(self,value):
    
    def getValue(value):
      if str(value).find("{{page}}")>-1:
        value = str(value).replace("{{page}}", str(self.page_no()))
      dbv = str(value).split(".")
      if self.databind.has_key(dbv[0]):
        if type(self.databind[dbv[0]]).__name__=="list":
          try:
            if self.databind[dbv[0]][int(dbv[1])][dbv[2]]!=None:
              return str(self.databind[dbv[0]][int(dbv[1])][dbv[2]])
            else: return ""
          except Exception:
            return value
        elif type(self.databind[dbv[0]]).__name__=="dict":
          try:
            if self.databind[dbv[0]][dbv[1]]!=None:
              return str(self.databind[dbv[0]][dbv[1]])
            else: return ""
          except Exception:
            return value
        elif self.databind[dbv[0]]!=None:
          if self.databind[dbv[0]]!=None:
            return str(self.databind[dbv[0]])
          else: return ""
        else:
          return value
      else:
        return value
          
    if str(value).find("={{")>-1 and str(value).find("}}")>-1:
      _value = value[str(value).find("={{")+3:str(value).find("}}")]
      value = value.replace("={{"+_value+"}}",getValue(_value));
      if str(value).find("={{")>-1 and str(value).find("}}")>-1:
        return self.setValue(value)
      else:
        return value
    else:
      return getValue(value)
  
  def setHtmlValue(self,html,fieldname):
    start_index=str(html).find("={{")
    if start_index>-1:
      value_old = html[start_index:str(html).find("}}",start_index)+2]
      value_new = self.setValue(value_old)
      html = str(html).replace(value_old, value_new)
      self.template["xml"]["details"] += "\n      <"+str(fieldname)+"><![CDATA["+str(value_new)+"]]></"+str(fieldname)+">"
      if str(html).find("={{")>-1:
        return self.setHtmlValue(html,fieldname)
      else:
        return html
    else:
      return html
  
  def get_datagrid_col_align(self,align):
    if align in("L","left"):
      align="left"
    elif align in("R","right"):
      align="right"
    elif align in("C","center"):
      align="center"
    elif align in("J","justify"):
      align="justify"
    return align
  
  def set_datagrid_col_number_format(self,col,value):
    digit = col.get("digit","")
    if digit!="":
      try:
        value = str("{0:."+str(digit)+"f}").format(float(value))
      except Exception:
        pass
    thousands = col.get("thousands","")
    if thousands!="":
      value = self.splitThousands(value,thousands)
    return value
              
  def create_datagrid(self,grid_element):
    databind = grid_element.get("databind","")
    xname = grid_element.get("name","items")
    header,columns=None,None
    for element in grid_element:
      if element.tag =="header":
        header = element
      elif element.tag =="columns":
        columns = element
    border = grid_element.get("border","1")
    if border=="1":
      border_sty="border:2px solid;"
    else:
      border_sty="border:none;"
    dgwidth = grid_element.get("width","100%")
    fontsize = grid_element.get("font-size",self.template["style"]["font-size"])
    style = self.set_style(grid_element.attrib)
    style+=border_sty
    style='style="vertical-align:top !important;padding:1px;font-size:'+str(fontsize)+';border-collapse: separate !important;'+style+'"'
    dgrid = '<table '+style+' border="'+border+'" width="'+dgwidth+'" cellpadding="0" cellspacing="0">'
    header_background = grid_element.get("header-background","")
    if header_background!="":
      header_background=' bgcolor="'+hex(int(header_background)).replace("0x","#").upper()+'"'
    footer_background = grid_element.get("footer-background","")
    if footer_background!="":
      footer_background=' bgcolor="'+hex(int(footer_background)).replace("0x","#").upper()+'"'
    if header!=None:
      bgcolor = header.get("background-color","")
      merge = header.get("merge",0)
      if bgcolor!="":
        bgcolor=' bgcolor="'+hex(int(bgcolor)).replace("0x","#").upper()+'"'
    else:
      bgcolor=""
      merge = grid_element.get("merge",0)
    dfoot='<tfoot><tr '+bgcolor+footer_background+'>'
    dgrid+='<thead><tr '+bgcolor+header_background+'>'
    self.template["xml"]["details"] += "\n    <"+str(xname)+"_footer"+">"
    style = style.replace("border:2px", "border:1px").replace("padding:1px", "padding:4px")
    total_ws = self.w - self.r_margin - self.l_margin
    col_check = {"pc":0, "nc":0, "nc_width":0}
    footer_cols = [0]
    for col in columns:
      width = col.get("width","")
      if width=="":
        col_check["nc"]+=1
      elif str(width).endswith("%"):
        col_check["pc"]+=float(str(width).replace("%", ""))
      else:
        col_check["pc"]+=int(float(width)/self.k/total_ws*100)
      if col.get("footer","")!="":
        footer_cols.append(1)
        if len(footer_cols)==2:
          footer_cols[1]+=footer_cols[0]
      else:
        footer_cols[len(footer_cols)-1]+=1
    del footer_cols[0]
    if total_ws > total_ws*col_check["pc"]/100:
      col_check["nc_width"] = 100-(total_ws*col_check["pc"]/100)/total_ws*100+1
    for cn in range(len(columns)):
      width = columns[cn].get("width","")
      label = self.setValue(columns[cn].get("label",""))
      if width=="":
        width = str(int(col_check["nc_width"]/col_check["nc"]))+"%"
      elif not str(width).endswith("%"):
        width = str(int(float(width)/self.k/total_ws*100))+"%"
      if col_check["nc"]==0 and col_check["nc_width"]>0 and cn==len(columns)-1:
        width = str(int(float(str(width).replace("%", ""))+col_check["nc_width"]))+"%"
      if col_check["pc"]>100 and cn==len(columns)-1:
        if float(str(width).replace("%", ""))-(col_check["pc"]-100)>0:
          width = str(int(float(str(width).replace("%", ""))-(col_check["pc"]-100))+1)+"%"
      columns[cn].set("width",width)
      dgrid+='<th '+style+' width="'+width+'" align="'+self.get_datagrid_col_align(columns[cn].get("header-align","left"))+'">'+label+'</th>'
      if columns[cn].get("footer","")!="":
        if len(footer_cols)>0:
          fospan = ' colspan="'+str(footer_cols[0])+'" '
          del footer_cols[0]
        else:
          fospan = ' colspan="1" '
        footer = self.setValue(columns[cn].get("footer",""))
        footer = self.set_datagrid_col_number_format(columns[cn],footer)
        dfoot+='<td '+fospan+style+' align="'+self.get_datagrid_col_align(columns[cn].get("footer-align","left"))+'">'+footer+'</td>'
        fieldname = columns[cn].get("fieldname","")
        self.template["xml"]["details"] += "\n      <"+str(fieldname)+"><![CDATA["+str(footer)+"]]></"+str(fieldname)+">"
    self.template["xml"]["details"] += "\n    </"+str(xname)+"_footer"+">"
    dfoot+="</tr></tfoot>"
    dgrid+="</tr></thead>"
    dgrid+="<tbody>"    
    if self.databind.has_key(databind) and columns!=None:
      rows = self.databind[databind]
      if len(rows)>0:
        for rn in range(len(rows)):
          rheight = self.get_row_height(rows[rn], columns)
          row="<tr>"
          self.template["xml"]["details"] += "\n    <"+str(xname)+">"
          svalue=""
          for col in columns:
            fieldname = col.get("fieldname","")
            if fieldname=="counter":
              value=str(rn+1)
            elif rows[rn].has_key(fieldname):
              if rows[rn][fieldname]!=None and rows[rn][fieldname]!="":
                value=str(rows[rn][fieldname])
              else:
                value="&nbsp; "
            else:
              value="&nbsp; "
            value = self.set_datagrid_col_number_format(col,value)
            if svalue!="": 
              svalue+=" "
            svalue += value
            if merge!=1:
              row+='<td height="'+str(int(rheight["mrh"]))+'" multiline="'+rheight["rh"][fieldname]+'" '+style+' align="'+self.get_datagrid_col_align(col.get("align","left"))+'">'+value+'</td>'
            self.template["xml"]["details"] += "\n      <"+str(fieldname)+"><![CDATA["+str(value)+"]]></"+str(fieldname)+">"
          if merge==1:
            row+='<td width="100%" height="'+str(int(rheight["mrh"]))+'" multiline="'+rheight["rh"][fieldname]+'" '+style+'>'+svalue+'</td>'
          dgrid+=row+"</tr>"
          self.template["xml"]["details"] += "\n    </"+str(xname)+">"
      else:
        dgrid=""
    else:
      dgrid=""
    if dgrid!="": 
      dgrid+="</tbody>"+dfoot+'</table>'
    return [dgrid, merge]
  
  def get_row_height(self, row,columns):
    retval = {"mrh":1, "rh":{}}
    for col in columns:
      fieldname = col.get("fieldname","")
      col_width = col.get("width","0")
      if col_width=="0":
        col_width=str(int(100/len(columns)))+"%"
      if col_width[-1]=='%':
        total = self.w - self.r_margin - self.l_margin
        col_width = (float(col_width[:-1])-1) * total / 100
      else:
        col_width = col_width.replace(col_width.rstrip('0123456789'),"")
        col_width = float(col_width)
      if fieldname=="counter":
        rh=1
      else:
        rh = len(self.multi_cell(col_width,None,str(row[fieldname]),0,'L',0,True))
      if rh>1:
        retval["rh"][fieldname]="multi"
        if retval["mrh"]<rh:
          retval["mrh"]=rh
      else:
        retval["rh"][fieldname]="single"
    return retval
  
  def handle_data(self, txt):
    pass
  
  def create_row(self,section,row_element):
    existing = row_element.get("visible",None)
    if existing:
      if self.databind.has_key(existing):
        if len(self.databind[existing])==0:
          return ""
      else:
        return ""
    hrow = '<table width="100%" cellpadding="0" cellspacing="0" style="border-collapse: separate !important;"><tr>'
    hgap = float(row_element.get("hgap",0))
    max_height = row_height = float(row_element.get("height",self.template["page"]["row_height"]))
    if row_height!=self.template["page"]["row_height"]:
      row_height=float(row_height)
    row_element = row_element.getiterator("columns")[0]
    for ei in range(len(row_element)):
      style = self.set_style(row_element[ei].attrib)
      if row_height>0:
        style+="height:"+str(row_height)+self.unit+";"
      if row_element[ei].tag =="cell":
        value = self.setValue(row_element[ei].get("value",""))
        
        width = row_element[ei].get("width","0")
        if width[-1]=='%':
          style+="width:"+width+";"
          total = self.w - self.r_margin - self.l_margin
          width = float(width[:-1]) * total / 100
        else:
          width = float(width)
          if len(row_element)-1>ei and width==0:
            width = self.get_string_width(str(value))+4
          if width>0:
            style+="width:"+str(width)+self.unit+";"
        
        border = row_element[ei].get("border","0")
        if border=="0":
          border=0
          style+="border:none;"
        if border=="1":
          border=1
          style+="border:2px solid;"
        if str(border).find("L")>-1:
          style+="border-left:solid;"
        if str(border).find("R")>-1:
          style+="border-right:solid;"
        if str(border).find("B")>-1:
          style+="border-bottom:solid;"
        if str(border).find("T")>-1:
          style+="border-top:solid;"
        
        align = row_element[ei].get("align","L")
        if str(align).find("L")>-1:
          style+="text-align:left;"
        elif str(align).find("R")>-1:
          style+="text-align:right;"
        elif str(align).find("C")>-1:
          style+="text-align:center;"
        elif str(align).find("J")>-1:
          style+="text-align:justify;"
        
        if row_element[ei].get("background-color",None)!=None: fill=1
        else: fill=0 
        link = self.setValue(row_element[ei].get("link",""))
        multiline = row_element[ei].get("multiline","false")
        page_no = self.page_no(); y = self.get_y(); x = self.get_x()
        if multiline=="true":
          self.multi_cell(w=width, h=row_height, txt=str(value).decode(self.template["page"]["decode"]).encode(self.template["page"]["encode"],'replace'), 
                          border=border, align=align, fill=fill)
          if self.page_no()>page_no:
            self.set_y(self.get_y()-self.lasth); max_height = row_height
          else:
            if self.get_y()-y>max_height: max_height = self.get_y()-y
            self.set_y(y);
          if ei<len(row_element)-1:
            self.set_x(x+width)
        else:
          self.cell(w=width, h=row_height, txt=str(value).decode(self.template["page"]["decode"]).encode(self.template["page"]["encode"],'replace'), 
                    border=border, ln=0, align=align, fill=fill, link=link)
          if self.get_y()-y>max_height: max_height = self.get_y()-y
        style+="padding:2px 4px 2px 4px;"
        if value=="":
          value="&nbsp;"
        if link!="":
          hrow+='<td style="'+style+'"><a href="'+link+'">'+str(value)+'</a></td>'
        else:
          hrow+='<td style="'+style+'">'+str(value)+'</td>'
        xname = row_element[ei].get("name","head")
        if section=="header" and xname!="label":
          self.template["xml"]["header"] += "\n    <"+str(xname)+"><![CDATA["+str(value)+"]]></"+str(xname)+">"
        if section=="details" and xname!="label":
          self.template["xml"]["details"] += "\n    <"+str(xname)+"><![CDATA["+str(value)+"]]></"+str(xname)+">"
      if row_element[ei].tag =="image":
        width = float(row_element[ei].get("width",0))
        if width>0:
          style+="width:"+str(width)+self.unit+";"
        
        img_src = row_element[ei].get("src","");
        link = self.setValue(row_element[ei].get("link",""))
        if img_src == "":
          name = self.setValue(row_element[ei].get("file",""))
          if name.find("/")==-1 and self.images_folder and name!="":
            name = self.images_folder+'/'+name
          if name!="":
            import os
            if not os.path.isfile(name):
              import urllib2
              try:
                urllib2.urlopen(urllib2.Request(name))
              except:
                name = ""
            if name!="":  
              self.image(name=name, x=self.get_x()+0.5, y=self.get_y()+0.5, w=width, link=link)
              self.set_x(self.get_x()+(self.images[name]['w']/self.k)+0.5)
            if link!="":
              hrow+='<td width="1'+self.unit+'"><a href="'+link+'"><img style="'+style+'" src="'+name+'"/></a></td>'
            else:
              hrow+='<td width="1 '+self.unit+'"><img style="'+style+'" src="'+name+'"/></td>'
        else:
          value_data = self.setValue(img_src);
          name = "img_data" if value_data==img_src else img_src
          info = None
          if not self.images.has_key(name):
            img_data = value_data.split(",")
            if len(img_data)>1 and img_data[0].startswith("data:image"):
              img_type = img_data[0].split(";")[0].split(":")[1]
              if img_data[0].find("base64")>-1:
                bimg = BytesIO(img_data[1].decode("base64"))
              else:
                bimg = BytesIO(img_data[1])
              if img_type == "image/jpeg" or img_type == "image/jpg":
                im = Image.open(bimg)
                bimg.seek(0)
                info = {'w':im.size[0],'h':im.size[1],'bpc':8,'f':'DCTDecode','data':bimg.read()}
                if im.mode == 'RGB':
                  info['cs']='DeviceRGB'
                elif im.mode == 'CMYK':
                  info['cs']='DeviceCMYK'
                else:
                  info['cs']='DeviceGray'
              elif img_type == "image/png":
                info = self.parsepng_(bimg)
              else:
                pass
              bimg.close()
              
              if info:
                info['i']=len(self.images)+1
                info['odata']=value_data
                self.images[name]=info
          if self.images.has_key(name):
            self.image(name=name, x=self.get_x()+0.5, y=self.get_y()+0.5, w=width, link=link)
            self.set_x(self.get_x()+(self.images[name]['w']/self.k)+0.5)
            if link!="":
              hrow+='<td width="1'+self.unit+'"><a href="'+link+'"><img style="'+style+'" src="'+self.images[name]['odata']+'"/></a></td>'
            else:
              hrow+='<td width="1 '+self.unit+'"><img style="'+style+'" src="'+self.images[name]['odata']+'"/></td>'
        if self.images[name]['h']/self.k>max_height:
          max_height = self.images[name]['h']/self.k
      if row_element[ei].tag =="separator":
        gap = row_element[ei].get("gap",0)
        if gap!="":
          gap=float(gap)
        if gap>0:
          width=gap
          border="LR"
        elif ei==0:
          width=1
          border="L"
        else:
          width=1
          border="R"
        self.cell(w=width, h=max_height, txt="", border=border, ln=0)
      if row_element[ei].tag =="hgap":
        width = row_element[ei].get("width",0)
        if width!=0:
          width=float(width)
        self.cell(w=width, txt="", border=0, ln=0)
        hrow+='<td style="width:'+str(width)+self.unit+';"> </td>'
      if row_element[ei].tag =="barcode":
        code_type = row_element[ei].get("code-type","code39")
        wide = float(row_element[ei].get("wide",self.get_def_barcode_size(code_type, "wide")))
        narrow = float(row_element[ei].get("narrow",self.get_def_barcode_size(code_type, "narrow")))
        code = row_element[ei].get("value","")
        visible = row_element[ei].get("visible-value",0)
        if visible==1:
          bc_height = narrow+self.font_size_pt/self.k+2
        else: bc_height = narrow 
        if self.get_y() + bc_height > self.page_break_trigger and not self.in_footer:
          x = self.get_x()
          self.add_page(self.cur_orientation)
          self.set_x(x)
        if code_type=="code39":
          x = self.code39_(code, self.get_x(), self.get_y(), wide, narrow, (visible==1))
          hrow+='<td><img src="'+code39_prev+'"/></td>'
        elif code_type=="i2of5":
          x = self.interleaved2of5_(code, self.get_x(), self.get_y(), wide, narrow, (visible==1))
          hrow+='<td><img src="'+i2of5_prev+'"/></td>'
        self.set_x(x+0.5)
        if bc_height>max_height: max_height = bc_height
      if len(row_element)-1==ei: 
        self.ln(max_height)
      elif len(row_element)-1>ei and hgap>0:
        self.cell(w=hgap, txt="", border=0, ln=0)
        hrow+='<td style="width:'+str(hgap)+self.unit+';"> </td>'
    hrow+='</tr></table>'
    return hrow
  
  def get_def_barcode_size(self, code_type, size_type):
    if code_type=="code39":
      if size_type=="wide": return 1.5
      else: return 5
    if code_type=="i2of5":
      if size_type=="wide": return 1
      else: return 10
    return 0
  
  def check_barchar(self,bar_char,code):
    for i in xrange (0, len(code)-1):
      char_bar = code[i]
      if not char_bar in bar_char.keys():
        return False
    return True
  
  def interleaved2of5_(self, txt, x, y, w=1.0, h=10.0, visible=False):
    "Barcode I2of5 (numeric), adds a 0 if odd lenght"
    narrow = w / 3.0
    wide = w
    
    self.set_font_size(10)
    start_x = x
    
    # wide/narrow codes for the digits
    bar_char={'0': 'nnwwn', '1': 'wnnnw', '2': 'nwnnw', '3': 'wwnnn',
              '4': 'nnwnw', '5': 'wnwnn', '6': 'nwwnn', '7': 'nnnww',
              '8': 'wnnwn', '9': 'nwnwn', 'A': 'nn', 'Z': 'wn'}
    
    if not self.check_barchar(bar_char, txt):
      self.text(x, y+h+self.font_size_pt/self.k, 'Invalid barcode: '+txt)
      return (self.get_x(),h+self.font_size_pt/self.k+1)
    
    self.set_fill_color(0)
    code = txt
    # add leading zero if code-length is odd
    if len(code) % 2 != 0:
      code = '0' + code

    # add start and stop codes
    code = 'AA' + code.lower() + 'ZA'

    for i in xrange(0, len(code), 2):
      # choose next pair of digits
      char_bar = code[i]
      char_space = code[i+1]
      
      # create a wide/narrow-seq (first digit=bars, second digit=spaces)
      seq = ''
      for s in xrange(0, len(bar_char[char_bar])):
        seq += bar_char[char_bar][s] + bar_char[char_space][s]

      for bar in xrange(0, len(seq)):
        # set line_width depending on value
        if seq[bar] == 'n':
          line_width = narrow
        else:
          line_width = wide

        # draw every second value, the other is represented by space
        if bar % 2 == 0:
          self.rect(x, y, line_width, h, 'F')
        
        x += line_width
    if visible:
      start_x = start_x+(x-start_x-self.get_string_width(txt))/2
      self.text(start_x, y+h+self.font_size_pt/self.k, txt)
    return x
  
  def code39_(self, txt, x, y, w=1.5, h=5.0, visible=False):
    "Barcode 3of9"
    wide = w
    narrow = w / 3.0
    gap = narrow
    
    self.set_font_size(10)
    start_x = x
    
    bar_char={'0': 'nnnwwnwnn', '1': 'wnnwnnnnw', '2': 'nnwwnnnnw',
              '3': 'wnwwnnnnn', '4': 'nnnwwnnnw', '5': 'wnnwwnnnn',
              '6': 'nnwwwnnnn', '7': 'nnnwnnwnw', '8': 'wnnwnnwnn',
              '9': 'nnwwnnwnn', 'A': 'wnnnnwnnw', 'B': 'nnwnnwnnw',
              'C': 'wnwnnwnnn', 'D': 'nnnnwwnnw', 'E': 'wnnnwwnnn',
              'F': 'nnwnwwnnn', 'G': 'nnnnnwwnw', 'H': 'wnnnnwwnn',
              'I': 'nnwnnwwnn', 'J': 'nnnnwwwnn', 'K': 'wnnnnnnww',
              'L': 'nnwnnnnww', 'M': 'wnwnnnnwn', 'N': 'nnnnwnnww',
              'O': 'wnnnwnnwn', 'P': 'nnwnwnnwn', 'Q': 'nnnnnnwww',
              'R': 'wnnnnnwwn', 'S': 'nnwnnnwwn', 'T': 'nnnnwnwwn',
              'U': 'wwnnnnnnw', 'V': 'nwwnnnnnw', 'W': 'wwwnnnnnn',
              'X': 'nwnnwnnnw', 'Y': 'wwnnwnnnn', 'Z': 'nwwnwnnnn',
              '-': 'nwnnnnwnw', '.': 'wwnnnnwnn', ' ': 'nwwnnnwnn',
              '*': 'nwnnwnwnn', '$': 'nwnwnwnnn', '/': 'nwnwnnnwn',
              '+': 'nwnnnwnwn', '%': 'nnnwnwnwn'}
    
    code = txt.upper()
    if not self.check_barchar(bar_char, txt):
      self.text(x, y+h+self.font_size_pt/self.k, 'Invalid barcode: '+txt)
      return (self.get_x(),h+self.font_size_pt/self.k+1)
    
    self.set_fill_color(0)
    for i in xrange (0, len(code), 2):
      char_bar = code[i]

      seq= ''
      for s in xrange(0, len(bar_char[char_bar])):
        seq += bar_char[char_bar][s]

      for bar in xrange(0, len(seq)):
        if seq[bar] == 'n':
          line_width = narrow
        else:
          line_width = wide

        if bar % 2 == 0:
          self.rect(x, y, line_width, h, 'F')
        x += line_width
    x += gap
    if visible:
      start_x = start_x+(x-start_x-self.get_string_width(txt))/2
      self.text(start_x, y+h+self.font_size_pt/self.k, txt)
    return x
                      
  def parsepng_(self, f):
    from fpdf.php import substr
    import zlib, re
    
    #Check signature
    if(f.read(8)!='\x89'+'PNG'+'\r'+'\n'+'\x1a'+'\n'):
      return None
    #Read header chunk
    f.read(4)
    if(f.read(4)!='IHDR'):
      return None
    w=self._freadint(f)
    h=self._freadint(f)
    bpc=ord(f.read(1))
    if(bpc>8):
      return None
    ct=ord(f.read(1))
    if(ct==0 or ct==4):
      colspace='DeviceGray'
    elif(ct==2 or ct==6):
      colspace='DeviceRGB'
    elif(ct==3):
      colspace='Indexed'
    else:
      return None
    if(ord(f.read(1))!=0):
      return None
    if(ord(f.read(1))!=0):
      return None
    if(ord(f.read(1))!=0):
      return None
    f.read(4)
    dp='/Predictor 15 /Colors '
    if colspace == 'DeviceRGB':
      dp+='3'
    else:
      dp+='1'
    dp+=' /BitsPerComponent '+str(bpc)+' /Columns '+str(w)+''
    #Scan chunks looking for palette, transparency and image data
    pal=''
    trns=''
    data=''
    n=1
    while n != None:
      n=self._freadint(f)
      ftype=f.read(4)
      if(ftype=='PLTE'):
        #Read palette
        pal=f.read(n)
        f.read(4)
      elif(ftype=='tRNS'):
        #Read transparency info
        t=f.read(n)
        if(ct==0):
          trns=[ord(substr(t,1,1)),]
        elif(ct==2):
          trns=[ord(substr(t,1,1)),ord(substr(t,3,1)),ord(substr(t,5,1))]
        else:
          pos=t.find('\x00')
          if(pos!=-1):
            trns=[pos,]
        f.read(4)
      elif(ftype=='IDAT'):
        #Read image data block
        data+=f.read(n)
        f.read(4)
      elif(ftype=='IEND'):
        break
      else:
        f.read(n+4)
    if(colspace=='Indexed' and not pal):
      return None
    f.close()
    info = {'w':w,'h':h,'cs':colspace,'bpc':bpc,'f':'FlateDecode','dp':dp,'pal':pal,'trns':trns,}
    if(ct>=4):
      # Extract alpha channel
      data = zlib.decompress(data)
      color = '';
      alpha = '';
      if(ct==4):
        # Gray image
        length = 2*w
        for i in range(h):
          pos = (1+length)*i
          color += data[pos]
          alpha += data[pos]
          line = substr(data, pos+1, length)
          try:
            color += re.sub('(.).',lambda m: m.group(1),line, flags=re.DOTALL)
            alpha += re.sub('.(.)',lambda m: m.group(1),line, flags=re.DOTALL)
          except:
            color += re.sub('(?s)(.).',lambda m: m.group(1),line) 
            alpha += re.sub('(?s).(.)',lambda m: m.group(1),line)
      else:
        # RGB image
        length = 4*w
        for i in range(h):
          pos = (1+length)*i
          color += data[pos]
          alpha += data[pos]
          line = substr(data, pos+1, length)
          try:
            color += re.sub('(.{3}).',lambda m: m.group(1),line, flags=re.DOTALL)
            alpha += re.sub('.{3}(.)',lambda m: m.group(1),line, flags=re.DOTALL)
          except:
            color += re.sub('(?s)(.{3}).',lambda m: m.group(1),line) 
            alpha += re.sub('(?s).{3}(.)',lambda m: m.group(1),line)
      del data
      data = zlib.compress(color)
      info['smask'] = zlib.compress(alpha)
      if (self.pdf_version < '1.4'):
        self.pdf_version = '1.4'
    info['data'] = data
    return info

  def splitThousands(self, s, tSep=',', dSep='.'):
    if s == None:
      return 0
    if not isinstance( s, str ):
      s = str( s )
    cnt=0
    numChars=dSep+'0123456789'
    ls=len(s)
    while cnt < ls and s[cnt] not in numChars: cnt += 1
    lhs = s[ 0:cnt ]
    s = s[ cnt: ]
    if dSep == '':
      cnt = -1
    else:
      cnt = s.rfind( dSep )
    if cnt > 0:
      rhs = dSep + s[ cnt+1: ]
      s = s[ :cnt ]
    else:
      rhs = ''
    splt=''
    while s != '':
      splt= s[ -3: ] + tSep + splt
      s = s[ :-3 ]
    return lhs + splt[ :-1 ] + rhs
  
  def hex2greyscale(self, hexvalue):
    if str(hexvalue).startswith("#"): hexvalue = str(hexvalue)[1:7]
    return (int(str(hexvalue)[0:2],16)+int(str(hexvalue)[2:4],16)+int(str(hexvalue)[4:6],16))/3
  
  def get_value(self, value, defvalue):
    return value if value!=None else defvalue
    
  def parse_value(self, vname, value):
    #value or default
    if vname in("author","creator","subject","title"):
      value = self.get_value(value,self.template["document"][vname])
      return value
    
    elif vname=="format":
      value = self.get_value(str(value).lower(),"a4")
      if value not in ("a3","a4","a5","letter","legal"):
        value = "a4"
      return value
    
    elif vname=="orientation":
      value = self.get_value(str(value).lower(),"p")
      if value in ("p","portrait"):
        value = "P"
      elif value in ("l","landscape"):
        value = "L"
      else:
        value = "P"
      return value
    
    elif vname=="unit":
      value = self.get_value(str(value).lower(),"mm")
      if value not in ("pt","mm","cm","in"):
        value = "mm"
      return value
      
    #integer
    elif vname in("font-size","digit"):
      try:
        return int(value)
      except Exception:
        return None
      
    #float
    elif vname in("left-margin","right-margin","top-margin","bottom-margin","height","gap","hgap","wide","narrow"):
      try:
        return float(value)
      except Exception:
        return None
    
    elif vname=="width":
      value = self.get_value(str(value).replace(str(value).lstrip('0123456789%'),""),"")
      if value=="": return None
      else: return value
    
    #string
    elif vname in("src","value","name","databind","label","fieldname","thousands","footer",
                  "decpoint","visible","border","decode","encode","link","file","html","code-type"):
      return value
        
    elif vname in("merge","visible-value"):
      try:
        return int(self.get_value(value,0))
      except Exception:
        value = 0
      if value>0: return 1 
      else: return value
      
    elif vname=="multiline":
      value = self.get_value(value,"false")
      if value!="false": return "true" 
      else: return value
      
    elif vname=="compression":
      value = self.get_value(value,"medium")
      if value in("fast","medium","slow"):
        pass
      else:
        value = "medium"
      return value
        
    elif vname in("align","header-align","footer-align"):
      value = self.get_value(value,"L")
      if value=="left": return "L"
      elif value=="right": return "R"
      elif value=="center": return "C"
      elif value=="justify" and vname=="align": return "J"
      elif value in("L","C","R"): return value
      else: return "L"
        
    #check fonts
    elif vname=="font-family":
      value = self.get_value(value,None)
      if value not in ("times","helvetica","courier"):
        value = "times"
      return value
        
    #check font-style
    elif vname=="font-style":
      value = self.get_value(value,None);
      if value!=None:
        if value=="normal":
          value = ""
        elif value=="bold":
          value = "B"
        elif value=="italic":
          value = "I"
        elif value=="bolditalic":
          value = "BI"
        elif value in("B","I","BI",""):
          pass
        else:
          value = ""
        return value
      return None
              
    #color
    elif vname in("color","border-color","background-color","footer-background","header-background"):
      value =self. get_value(value,None)
      if value!=None:
        try:
          if str(value).startswith("#"):
            value = int(str(value).replace("#",""),16)
          else:
            value = int(value)
          if value<=255:
            return value* 0x00010101
          else: return value
        except Exception:
          return None
      return None
    return None
  
  def get_parent(self, element):
    for parent in self.template["elements"]["details"].getiterator():
      for child in parent:
        if element==child: return parent
    for parent in self.template["elements"]["header"].getiterator():
      for child in parent:
        if element==child: return parent
    for parent in self.template["elements"]["footer"].getiterator():
      for child in parent:
        if element==child: return parent
        
  def parse_xml(self, xml):
    for elements in xml:
      if elements.tag in("row","datagrid","html","vgap","hline"):
        for aname in elements.attrib.keys():
          elements.attrib[aname] = self.parse_value(aname, elements.attrib[aname])
          if elements.attrib[aname]==None:
            del elements.attrib[aname]
        if elements.tag == "row":
          if len(elements.getiterator("columns"))==0:
            items = []
            for element in elements:
              items.append(element)
            columns = et.SubElement(elements, "columns")
            for element in items:
              if (element!=columns):
                elements.remove(element)
              if element.tag in("cell","image","separator","barcode","hgap"):
                columns.append(element)
          else:
            columns = elements.getiterator("columns")[0]
          for element in columns:
            if element.tag in("cell","image","separator","barcode","hgap"):
              for aname in element.attrib.keys():
                element.attrib[aname] = self.parse_value(aname, element.attrib[aname])
                if element.attrib[aname]==None:
                  del element.attrib[aname]
            else:
              columns.remove(element)
        elif elements.tag=="datagrid":
          for element in elements:
            if element.tag in("header","footer"):
              if element.attrib.has_key("background-color"):
                element.attrib["background-color"] = self.parse_value("background-color", element.attrib["background-color"])
                elements.set(element.tag+"_background",element.attrib["background-color"])
            elif element.tag=="columns":
              for column in element:
                for aname in column.attrib.keys():
                  column.attrib[aname] = self.parse_value(aname, column.attrib[aname])
                  if column.attrib[aname]==None:
                    del column.attrib[aname]
      else:
        xml.remove(elements)
          
#public functions
  def loadDefinition(self,data):
    if data!="" or data!=None:
      repdef = et.XML(data)
      if len(repdef.getiterator("report"))>0:
        report_attr = repdef.getiterator("report")[0]
        self.template["document"]["author"] = self.parse_value("author", 
          report_attr.get("author",self.template["document"]["author"]))
        self.template["document"]["creator"] = self.parse_value("creator", 
          report_attr.get("creator",self.template["document"]["creator"]))
        self.template["document"]["subject"] = self.parse_value("subject", 
          report_attr.get("subject",self.template["document"]["subject"]))
        self.template["document"]["title"] = self.parse_value("title", 
          report_attr.get("title",self.template["document"]["title"]))
        
        self.template["margins"]["left-margin"] = self.parse_value("left-margin", 
          report_attr.get("left-margin",self.template["margins"]["left-margin"]))
        self.template["margins"]["right-margin"] = self.parse_value("right-margin", 
          report_attr.get("right-margin",self.template["margins"]["right-margin"]))
        self.template["margins"]["top-margin"] = self.parse_value("top-margin", 
          report_attr.get("top-margin",self.template["margins"]["top-margin"]))
        
        self.template["page"]["decode"] = self.parse_value("decode",
          report_attr.get("decode",self.template["page"]["decode"]))
        self.template["page"]["encode"] = self.parse_value("encode",
          report_attr.get("encode",self.template["page"]["encode"]))
        
        self.template["style"]["font-family"]=self.parse_value("font-family",
          report_attr.get("font-family",self.template["style"]["font-family"]))
        self.template["style"]["font-style"]=self.parse_value("font-style",
          report_attr.get("font-style",self.template["style"]["font-style"]))
        self.template["style"]["font-size"]=self.parse_value("font-size",
          report_attr.get("font-size",self.template["style"]["font-size"]))
        self.template["style"]["color"]=self.parse_value("color",
          report_attr.get("color",self.template["style"]["color"]))
        self.template["style"]["border-color"]=self.parse_value("border-color",
          report_attr.get("border-color",self.template["style"]["border-color"]))
        self.template["style"]["background-color"]=self.parse_value("background-color",
          report_attr.get("background-color",self.template["style"]["background-color"]))
          
      if len(repdef.getiterator("header"))>0:
        self.parse_xml(repdef.getiterator("header")[0])
        self.template["elements"]["header"] = repdef.getiterator("header")[0]
        self.header_show = self.template["elements"]["header"].get("show_in","all")
      if len(repdef.getiterator("footer"))>0:
        self.parse_xml(repdef.getiterator("footer")[0])
        self.template["elements"]["footer"] = repdef.getiterator("footer")[0]
        self.footer_show = self.template["elements"]["footer"].get("show_in","all")
      if len(repdef.getiterator("details"))>0:
        self.parse_xml(repdef.getiterator("details")[0])
        self.template["elements"]["details"] = repdef.getiterator("details")[0]
      
      if len(repdef.getiterator("data"))>0:
        for data in repdef.getiterator("data")[0]:
          if len(data.attrib)>0:
            self.databind[data.tag] = data.attrib
          elif len(data._children)>0:
            self.databind[data.tag] = []
            for item in data._children:
              self.databind[data.tag].append(item.attrib)
          elif data.text!="":
            self.databind[data.tag] = data.text
      return True
    return False
  
  def setData(self, key, value):
    if self.databind.has_key(key) and type(value).__name__=="dict":
      if type(self.databind[key]).__name__=="dict":
        for value_key in value.keys():
          self.databind[key][value_key] = value[value_key]
      else:
        self.databind[key] = value;
    else:
      self.databind[key] = value
    return True
  
  def getData(self, key):
    return self.databind[key];
  
  def insertElement(self, parent, ename="row", index=-1, values={}):
    if parent==None:
      parent = self.template["elements"]["details"]
    if not ename in("row","datagrid","vgap","hline","html","column","cell","image","separator","barcode","hgap"):
      return None
    if parent.tag in("row","datagrid"):
      parent = parent.getiterator("columns")[0]
    for aname in values.keys():
      values[aname] = self.parse_value(aname, values[aname])
      if values[aname]==None:
        del values[aname]
    element = parent.makeelement(ename, values.copy())
    if ename in("row","datagrid"):
      et.SubElement(element, "columns", {})
    elif ename == "html":
      if element.attrib.has_key("html"):
        element.text = element.attrib["html"]
        del element.attrib["html"]
    if index<0 or index>=len(parent):
      parent.append(element)
    else:
      parent.insert(index, element)
    return element
  
  def editElement(self, element, values={}):
    if et.iselement(element):
      for vkey in values.keys():
        if element.attrib.has_key(vkey) and values[vkey]==None:
          del element.attrib[vkey]
        else:
          element.attrib[vkey] = values[vkey]
      return True
    else: return False
  
  def deleteElement(self, element, parent=None):
    if et.iselement(element):
      if parent==None: parent = self.get_parent(element)
      if et.iselement(parent):
        if parent.tag in("row","datagrid"): parent = parent.getiterator("columns")[0]
        parent.remove(element)
        return True
  
  def createReport(self):
    self.set_author(self.template["document"]["author"])
    self.set_creator(self.template["document"]["creator"])
    self.set_subject(self.template["document"]["subject"])
    self.set_title(self.template["document"]["title"])
    
    self.set_left_margin(self.template["margins"]["left-margin"])
    self.set_right_margin(self.template["margins"]["right-margin"])
    self.set_top_margin(self.template["margins"]["top-margin"])
    
    self.set_style(self.template["style"])
    self.template["html"]["details"]=""
    self.add_page()
    if self.template["elements"]["details"]!=None:
      for element in self.template["elements"]["details"]:
        self.template["html"]["details"]+=self.show_element("details",element)
    return True
                    
  def save2PdfFile(self,path):
    return self.output(name=path,dest='F')
  
  def save2Pdf(self):
    return self.output(dest='S')
  
  def save2Html(self,details=False):
    style=""
    if self.orientation=="P":
      style+='width:'+str(self.fw)+self.unit+';'
    else:
      style+='width:'+str(self.fh)+self.unit+';'
    style+='word-wrap: break-word; overflow: auto;background-color:#FFFFFF;'
    style+='padding-left:'+str(self.l_margin)+self.unit+';padding-top:'+str(self.t_margin)+self.unit+';padding-right:'+str(self.r_margin)+self.unit+';'
    ohtml = '<html style="background-color:#CCCCCC;"><head><title>'+self.template["document"]["title"]+'</title><link rel="shortcut icon" href="'+self.template["page"]["favicon"]+'" type="image/x-icon"></head><body style="'+style+'">'
    if details==False:
      ohtml +=self.template["html"]["header"]
    ohtml +=self.template["html"]["details"]
    if details==False:
      ohtml +=self.template["html"]["footer"]
    ohtml +='<div style="height:'+str(self.t_margin)+self.unit+'"/></body></html>'
    return ohtml
  
  def save2Xml(self):
    return "<data>"+self.template["xml"]["header"]+self.template["xml"]["details"]+"\n</data>"
  

code39_prev = "data:image/png;base64,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"
i2of5_prev = "data:image/png;base64,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"