# -*- coding: utf-8 -*-

"""
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
"""

from fpdf.html import HTML2FPDF

def hex2dec(color = "#000000"):
  if color:
    r = int(color[1:3], 16)
    g = int(color[3:5], 16)
    b = int(color[5:7], 16)
    return r, g, b
      
class html2pdf(HTML2FPDF):
  def __init__(self, pdf):
    HTML2FPDF.__init__(self, pdf)
    
  def handle_data(self, txt):
    if self.td is not None: # drawing a table?
      if 'width' not in self.td and 'colspan' not in self.td:
        l = [self.table_col_width[self.table_col_index]]
      elif 'colspan' in self.td:
        i = self.table_col_index
        colspan = int(self.td['colspan'])
        l = self.table_col_width[i:i+colspan]
      else:
        l = [self.td.get('width','240')]
      w = sum([self.width2mm(lenght) for lenght in l])
      h = int(self.td.get('height', 1))*(self.h*1.50)
      self.table_h = h
      border = int(self.table.get('border', 0))
      if not self.th:
        align = self.td.get('align', 'L')[0].upper()
        border = border and 'LRB'
      else:
        self.set_style('B',True)
        border = border or 'B'
        align = self.td.get('align', 'C')[0].upper()
      bgcolor = hex2dec(self.td.get('bgcolor', self.tr.get('bgcolor', '')))
      # parsing table header/footer (drawn later):
      if self.thead is not None:
        self.theader.append(((w,h,txt,border,0,align), bgcolor))
      if self.tfoot is not None:
        self.tfooter.append(((w,h,txt,border,0,align), bgcolor))
      # check if reached end of page, add table footer and header:
      height = h + (self.tfooter and self.tfooter[0][0][1] or 0)
      if self.pdf.y+height>self.pdf.page_break_trigger and not self.th:
        self.output_table_footer()
        self.pdf.add_page()
        self.theader_out = self.tfooter_out = False
      if self.tfoot is None and self.thead is None:
        if not self.theader_out:
            self.output_table_header()
        self.box_shadow(w, h, bgcolor)
        multiline = self.td.get('multiline', "single")
        if multiline!="multi":
          self.pdf.cell(w,h,txt,0,0,align)
        else:
          cy = self.pdf.get_y()
          cx = self.pdf.get_x()
          self.pdf.multi_cell(w,(self.h*1.50),txt,0,align)
          self.pdf.set_y(cy)
          self.pdf.set_x(cx+w)
        if border!=0:
          self.pdf.rect(self.pdf.get_x()-w, self.pdf.get_y(), w, h, '')
    elif self.table is not None:
      # ignore anything else than td inside a table
      pass
    elif self.align:
      self.pdf.cell(0,self.h,txt,0,1,self.align[0].upper(), self.href)
    else:
      txt = txt.replace("\n"," ")
      if self.href:
        self.put_link(self.href,txt)
      else:
        self.pdf.write(self.h,txt)
            
class HTMLMixin():
  def write_html(self, text, **kwargs):
    "Parse HTML and convert it to PDF"
    h2p = html2pdf(self)
    if kwargs.has_key("font") and kwargs.has_key("fontsize"):
      h2p.set_font(kwargs["font"], kwargs["fontsize"])
    h2p.feed(text)
        