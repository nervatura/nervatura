/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

/*jshint -W117 */
/*eslint-env amd, browser*/

define(["jspdf"], function (JsPdf) {
  return function Report(orientation, unit, format) {

    var self = this;
    //var CLASS_VERSION = '1.150708';
    self.orientation = parse_value("orientation",orientation);
    self.unit = parse_value("unit",unit);
    self.format = parse_value("format",format);
    self.template = {
        document:{
          title:"Nervatura Report",
          author:"",
          creator:"",
          subject:"",
          keywords:""
        },
        margins:{
          "left-margin": 12,
          "top-margin": 12,
          "right-margin": 12,
          "bottom-margin": 12
        },
        style: {
          "font-family": "times",
          "font-style": "normal",
          "font-size": 12,
          color: "#000000",
          "border-color": 0,
          "background-color": 255
        },
        elements: {
          report: {}, 
          header: [],
          details: [],
          footer: [],
          data: {}
        },
        xml: {
          header: "",
          details: ""
        },
        page: {
          cursor: {x:0, y:0},
          padding: 2,
          header_height: 0,
          footer_height: 0,
          current_page: 1 }};
    var orig_data = {};
    //orientation, unit, format, compress, textColor, drawColor, fontSize, lineHeight, lineWidth
    var doc = new JsPdf({orientation: self.orientation, unit: self.unit, format: self.format,
                        textColor: self.template.style.color, drawColor: self.template.style["border-color"],
                        fontSize: self.template.style["font-size"]});

    function hex2greyscale(hex) {
      if (hex.charAt(0)==="#") {hex = hex.substring(1,7);}
      return (parseInt(hex.substring(0,2),16)+parseInt(hex.substring(2,4),16)+parseInt(hex.substring(4,6),16))/3;}
    
    function get_value(value,defvalue) {
      if ((typeof value!=="undefined")&&(value!==null)) {
        return value;} else {return defvalue;}}
        
    function check_page_break(current_y, step) {
      var dline = doc.internal.pageSize.height-self.template.margins["bottom-margin"];
      if (current_y<doc.internal.pageSize.height-self.template.page.footer_height-self.template.margins["bottom-margin"]) {
        dline -= self.template.page.footer_height;}
      if (current_y+step>dline) {
        return true;} else {return false;}}

    function setValue(value) {
      function getValue(value) {
        if (value.indexOf("{{page}}")>-1) {value = value.replace("{{page}}", self.template.page.current_page.toString());}
        var dbv = value.split(".");
        if (typeof self.template.elements.data[dbv[0]]!=="undefined") {
          if (typeof self.template.elements.data[dbv[0]]==="object") {
            if (Array.isArray(self.template.elements.data[dbv[0]])) {
              try {
                if (self.template.elements.data[dbv[0]][parseInt(dbv[1],10)][dbv[2]]!==null) {
                  return self.template.elements.data[dbv[0]][parseInt(dbv[1],10)][dbv[2]].toString();}
                else {return "";}}
              catch(err) {return value;}}
            else {
              try {
                if (self.template.elements.data[dbv[0]][dbv[1]]!==null) {
                  return self.template.elements.data[dbv[0]][dbv[1]].toString();}
                else {return "";}}
              catch(err) {return value;}}}
          else {
            if (self.template.elements.data[dbv[0]]!==null) {
              return self.template.elements.data[dbv[0]].toString();}
            else {return "";}}}
        else {return value;}}

      if (value.indexOf("={{")>-1 && value.indexOf("}}")>-1) {
        var _value = value.substring(value.indexOf("={{")+3, value.indexOf("}}"));
        value = value.replace("={{"+_value+"}}",getValue(_value));
        if (value.indexOf("={{")>-1 && value.indexOf("}}")>-1) {
          return setValue(value);} else {return value;}
      } else {
        return getValue(value);}}

    function setHtmlValue(html,fieldname) {
      if (html.indexOf("={{")>-1 && html.indexOf("}}")>-1) {
        var _value = html.substring(html.indexOf("={{"), html.indexOf("}}")+2);
        var value_ = setValue(_value);
        html = html.replace(_value,value_);
        self.template.xml.details += "\n    <"+fieldname.toString()+"><![CDATA["+value_.toString()+"]]></"+fieldname.toString()+">";
        if (html.indexOf("={{")>-1 && html.indexOf("}}")>-1) {
          html = setHtmlValue(html,fieldname);
        } else {return html;}
      } else {return html;}}
    
    function get_element_type(element) {
      if (Object.getOwnPropertyNames(element).length>0) {
        return Object.getOwnPropertyNames(element)[0];
      } else {return null;}}
    
    function get_element_style(element, options){
      if ((typeof options==="undefined")||(options===null)) {
        options = {};}
      options["font-family"] = get_value(element["font-family"],null);
      options["font-style"] = get_value(element["font-style"],null);
      options["font-size"] = get_value(element["font-size"],null);
      
      options.color = get_value(element.color,null);
      options["border-color"] = get_value(element["border-color"],null);
      options["background-color"] = get_value(element["background-color"],null);
      return options;}

    function set_page_style(options){
      //font-family, font-size, font-style, color,background-color,border-color

      doc.setFont(self.template.style["font-family"]);
      if ((typeof options["font-family"]!=="undefined")&&(options["font-family"]!==null)) {
        doc.setFont(options["font-family"]);}
      doc.setFontType(self.template.style["font-style"]);
      if ((typeof options["font-style"]!=="undefined")&&(options["font-style"]!==null)) {
        doc.setFontType(options["font-style"]);}
      doc.setFontSize(self.template.style["font-size"]);
      if ((typeof options["font-size"]!=="undefined")&&(options["font-size"]!==null)) {
        doc.setFontSize(options["font-size"]);}
      
      if ((typeof options.color==="undefined")||(options.color===null)) {
        options.color = self.template.style.color;}
      doc.setTextColor(options.color);
      if ((typeof options["background-color"]==="undefined")||(options["background-color"]===null)) {
        options["background-color"] = self.template.style["background-color"];}
      doc.setFillColor(options["background-color"]);
      if ((typeof options["border-color"]==="undefined")||(options["border-color"]===null)) {
        options["border-color"] = self.template.style["border-color"];}
      doc.setDrawColor(options["border-color"]);}

    function set_number_format(number, decimals, dec_point, thousands_sep) {
      var digit = decimals;
      if(decimals === "auto"){
        if(number.toString().indexOf(dec_point)>-1){
          digit = number.toString().substr(number.indexOf(dec_point)+1).length;}
        else {
          digit = 0;}}
      var n = !isFinite(+number) ? 0 : +number,
        prec = !isFinite(+digit) ? 0 : Math.abs(digit),
        sep = (typeof thousands_sep === 'undefined') ? ',' : thousands_sep,
        dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
        toFixedFix = function (n, prec) {
          var k = Math.pow(10, prec);
          return Math.round(n * k) / k;},
          s = (prec ? toFixedFix(n, prec) : Math.round(n)).toString().split('.');
      if (s[0].length > 3) {
        s[0] = s[0].replace(/\B(?=(?:\d{3})+(?!\d))/g, sep);}
      if ((s[1] || '').length < prec) {
        s[1] = s[1] || '';
        s[1] += new Array(prec - s[1].length + 1).join('0');}
      return s.join(dec);}

    function get_height(text,width,padding,fonttype) {
      fonttype = (fonttype === null) ? "normal" : fonttype;
      doc.setFontType(fonttype);
      var stext = doc.splitTextToSize(text,width-2*padding);
      return {text:stext, height:((doc.internal.getLineHeight()/doc.internal.scaleFactor))*(stext.length)+padding};}
    
    function get_img_size(img_src) {
      var isize = {width: 0, height: 0};
      var img = document.createElement('img'); img.id = "pimg"; img.src = img_src;
      document.body.appendChild(img);
      isize.width = img.width/doc.internal.scaleFactor; isize.height = img.height/doc.internal.scaleFactor;
      document.body.removeChild(img);
      return isize;}
      
    function header() {
      var eheader = self.template.elements.header;
      for(var i = 0; i < eheader.length; i++) {
        var etype = get_element_type(eheader[i]);
        switch(etype) {
          case "row":
          case "vgap":
          case "hline":
            create_elements("header", etype, eheader[i][etype]);
            break;}}}

    function footer() {
      var efooter = self.template.elements.footer;
      for(var i = 0; i < efooter.length; i++) {
        var etype = get_element_type(efooter[i]);
        switch(etype) {
          case "row":
          case "vgap":
          case "hline":
            create_elements("footer", etype, efooter[i][etype]);
            break;}}}

    function get_footer_height() {
      var footer_height = 0;
      var efooter = self.template.elements.footer;
      for(var i = 0; i < efooter.length; i++) {
        var itype = get_element_type(efooter[i]);
        var foo_item = efooter[i][itype];
        switch(itype) {
          case "row":
            var row_height = get_value(foo_item.height,0);
            for(var index = 0; index < Object.keys(foo_item.columns).length; index++) {
              var etype = get_element_type(foo_item.columns[index]);
              var element = foo_item.columns[index][etype];
              switch(etype) {
                case "cell":
                  var cell_height = get_value(element.height,null);
                  if (cell_height===null) {
                    cell_height = parseFloat(doc.internal.getLineHeight()/doc.internal.scaleFactor)+(self.template.page.padding);}
                  else {cell_height = parseFloat(cell_height);}
                  if (cell_height>row_height){row_height=cell_height;}
                  break;
                case "barcode":
                  var code_type = get_value(element["code-type"],"code39");
                  var visible = get_value(element["visible-value"],0);
                  var narrow = get_value(element.wide,get_def_barcode_size(code_type, "narrow"));
                  if (visible===1) {
                    narrow += (doc.internal.getFontSize()/doc.internal.scaleFactor);}
                  if (narrow>row_height){row_height=narrow;}
                  break;
                case "image":
                  var img_src = get_value(element.src,null);
                  if (img_src !== null) {
                    if (img_src.toString().substr(0,10)!=="data:image") {
                      img_src = setValue(img_src);}
                    if (img_src.toString().substr(0,10)==="data:image") {
                      var img_height = get_value(element.height,0);
                      if (img_height===0) {img_height = get_img_size(img_src).height;}
                      if (img_height>row_height) {row_height = img_height;}}}
                  break;}}
            footer_height += row_height;
            break;
          case "vgap":
            footer_height += get_value(foo_item.height,0);
            break;
          case "hline":
            footer_height += 1 + get_value(foo_item.gap,0);
            break;}}
      return footer_height;}

    doc.internal.events.subscribe('addPage', function (page) {
      self.template.page.current_page = page.pageNumber;
      var footer_height = self.template.page.footer_height; self.template.page.footer_height = 0;
      self.template.page.cursor.y = doc.internal.pageSize.height-self.template.margins["bottom-margin"]-footer_height;
      footer(); self.template.page.cursor.y = self.template.margins["top-margin"]; self.template.page.footer_height = footer_height;
      header(); self.template.page.cursor.x = self.template.margins["left-margin"];
      self.template.page.cursor.y = self.template.margins["top-margin"] + self.template.page.header_height;});

    function create_elements(section, etype, element) {
      var options = {};
      var existing = get_value(element.visible,"");
      if (existing!=="") {
        if (typeof self.template.elements.data[existing]==="undefined") {return "";}
        else {if (self.template.elements.data[existing].length===0) {return "";}}}
      switch(etype) {
        case "row":
          create_row(section,element);
          break;
        case "vgap":
          var height = get_value(element.height,0);
          if (height>0) {
            if (check_page_break(self.template.page.cursor.y, height)) {
              doc.addPage();}
            else {self.template.page.cursor.y += height;}}
          break;
        case "hline":
          var width = parseFloat(get_value(element.width,0));
          var gap = get_value(element.gap,0);
          if (width===0) {
            width=doc.internal.pageSize.width-self.template.margins["left-margin"]-self.template.margins["right-margin"];}
            if (typeof element["border-color"]!=="undefined"){
              options = {"border-color": element["border-color"]};
              set_page_style(options);}
            doc.line(self.template.page.cursor.x, self.template.page.cursor.y,
                self.template.page.cursor.x+width, self.template.page.cursor.y);
            if (gap>0) {
              doc.line(self.template.page.cursor.x, self.template.page.cursor.y+gap,
                  self.template.page.cursor.x+width, self.template.page.cursor.y+gap);}
            self.template.page.cursor.x = self.template.margins["left-margin"];
            self.template.page.cursor.y += 1+gap;
          break;
        case "html":
          var html = element.html; var fieldname = get_value(element.fieldname,"head");
          options = get_element_style(element); set_page_style(options); doc.text("", 0, 0);
          html = setHtmlValue(html,fieldname);
          doc.fromHTML(html, self.template.page.cursor.x, self.template.page.cursor.y,
              {'width': doc.internal.pageSize.width-self.template.margins["left-margin"]-self.template.margins["right-margin"],
               'elementHandlers':function(){}},
              function (dispose) {
                 self.template.page.cursor.x = self.template.margins["left-margin"];
                 self.template.page.cursor.y = dispose.y;
                 return true;}, {top: self.template.margins["top-margin"]+self.template.page.header_height,
                 bottom: self.template.margins["bottom-margin"]+self.template.page.footer_height});
          break;
        case "datagrid":
          create_datagrid(element);
          break;}
      return true;}

    function create_datagrid(grid_element) {
      var rows = get_value(grid_element.databind,null); if (rows===null) {return true;} 
      rows = self.template.elements.data[rows];if (typeof rows==="undefined") {return true;}
      if (rows.length===0) {return true;}
      
      var grid_options = get_element_style(grid_element);//font-size,color,border-color
      grid_options.xname = get_value(grid_element.name,"items");
      grid_options.border = get_value(grid_element.border,"1");
      
      var header_options={"font-size":grid_options["font-size"], color:grid_options.color, "border-color":grid_options["border-color"],
          border:grid_options.border, merge:0, "font-style":"bold", text:"", height:0,
          columns: [], columns_width: 0, grid_width: get_value(grid_element.width,"100%")};
      header_options.merge = get_value(grid_element.merge,0);
      header_options["background-color"] = get_value(grid_element["header-background"],null);
      
      var nwidth = doc.internal.pageSize.width-self.template.margins["right-margin"]-self.template.margins["left-margin"];
      if (header_options.grid_width.indexOf("%")>-1) {
        header_options.grid_width=nwidth*(parseFloat(header_options.grid_width.replace("%",""))/100);}
      else {header_options.grid_width = parseFloat(header_options.grid_width);
        if (header_options.grid_width>nwidth) {header_options.grid_width=nwidth;}}

      var footer_options={"font-size":grid_options["font-size"], color:grid_options.color, "border-color":grid_options["border-color"],
          border:grid_options.border, "font-style":"bold", text:"", height:0};
      footer_options["background-color"] = get_value(grid_element["footer-background"],null);
      header_options.extend = grid_options.extend = footer_options.extend = (header_options.grid_width===nwidth);

      var zcol = 0; var footers=[]; var footer_width = 0;
      if (grid_element.columns.length>0) {
        var columns_elements=grid_element.columns;
        for(var index = 0; index < columns_elements.length; index++) {
          var column = columns_elements[index].column;
          if (header_options.columns_width>=header_options.grid_width) {
            return false;}
          var column_options={"font-size":grid_options["font-size"], color:grid_options.color, "border-color":grid_options["border-color"],
              border:grid_options.border};
          column_options.fieldname = get_value(column.fieldname,"");
          column_options.label = setValue(get_value(column.label,""));
          if (header_options.merge===0) {
            column_options.width = get_value(column.width,null);
            if (column_options.width!==null) {
              if (column_options.width.toString().indexOf("%")>-1) {
                column_options.width = header_options.grid_width*(parseFloat(column_options.width.replace("%",""))/100);}
              else {column_options.width = parseFloat(column_options.width);}}
            else {
              column_options.width = parseFloat(
                doc.getStringUnitWidth(column_options.label)*doc.internal.getFontSize()/doc.internal.scaleFactor+self.template.page.padding*2)+1; 
              zcol += 1;}
            if (header_options.columns_width+column_options.width>=header_options.grid_width) {
              column_options.width = header_options.grid_width-header_options.columns_width;}
            header_options.columns_width += column_options.width;
            var cheight = get_height(column_options.label,column_options.width,self.template.page.padding,header_options["font-style"]).height;
            if (cheight>header_options.height){header_options.height = cheight;}}
          column_options.header_align = get_value(column["header-align"],"left");
          column_options.align = get_value(column.align,"left");
          if (column_options.align==="left") {
            column_options.multiline = "true";}
          else {column_options.multiline = "false";}

          if ((typeof column.thousands!=="undefined")||(column.digit!=="undefined")||(typeof column.decpoint!=="undefined")) {
            column_options.numbercol = true;
            column_options.thousands = get_value(column.thousands,"");
            column_options.digit = get_value(column.digit,"auto");
            column_options.decpoint = get_value(column.decpoint,".");}
          var footer_value = setValue(get_value(column.footer,""));
          var footer_align = get_value(column["footer-align"],"left");
          if (column_options.numbercol) {
            if (!isNaN(+footer_value) && isFinite(footer_value)) {
              footer_value = set_number_format(footer_value, column_options.digit, column_options.decpoint, column_options.thousands);}}
          if (get_value(column.footer,"")!=="") {
            if (footers.length===0) {
              footers.push({text:footer_value, align:footer_align, width:footer_width+column_options.width});}
            else {
              footers[footers.length-1].width += footer_width;
              footers.push({text:footer_value, align:footer_align, width:column_options.width});}
            footer_width = 0;}
          else {footer_width += column_options.width;}
          if (columns_elements.length-1===index) {column_options.ln=1;} else {column_options.ln=0;}
          header_options.columns.push(column_options);}}
      else {return true;}
      if (header_options.merge===0) {create_grid_header(header_options);}

      rows.forEach(function(row, row_index){
        self.template.xml.details += "\n    <"+grid_options.xname.toString()+">";
        grid_options.height = 0; grid_options.text ="";
        header_options.columns.forEach(function(column){
          if (column.fieldname==="counter") {
            column.text = (row_index+1).toString();}
          else {
            if ((typeof row[column.fieldname]!=="undefined") && (row[column.fieldname]!==null)) {column.text = row[column.fieldname];}
            else {column.text = "";}}
          if (column.numbercol) {
            if (!isNaN(+column.text) && isFinite(column.text) && (column.text!=="")) {
              column.text = set_number_format(column.text, column.digit, column.decpoint, column.thousands);}}
          if (header_options.merge===0) {
            cheight = get_height(column.text,column.width,self.template.page.padding,grid_options["font-style"]).height;
            if (cheight>grid_options.height){grid_options.height = cheight;}}
          else {
            grid_options.text += " "+column.text;
            self.template.xml.details += "\n      <"+column.fieldname.toString()+"><![CDATA["+column.text.toString()+"]]></"+column.fieldname.toString()+">";}});
        if (header_options.merge===0) {
          header_options.columns.forEach(function(column){
            grid_options.text = column.text;
            grid_options.width = column.width;
            grid_options.align = column.align;
            grid_options.ln = column.ln;
            grid_options.multiline = column.multiline;
            create_cell(grid_options, header_options);
            self.template.xml.details += "\n      <"+column.fieldname.toString()+"><![CDATA["+column.text.toString()+"]]></"+column.fieldname.toString()+">";});}
        else {
          grid_options.text = grid_options.text.trim();
          grid_options.width = null;
          grid_options.ln = 1;
          grid_options.multiline = "true";
          grid_options.height = null;
          create_cell(grid_options);}
        self.template.xml.details += "\n    </"+grid_options.xname.toString()+">";});

      if (header_options.merge===0) {
        footers.forEach(function(column){
          cheight = get_height(column.text,column.width,self.template.page.padding,footer_options["font-style"]).height;
          if (cheight>footer_options.height){footer_options.height = cheight;}});
        footers.forEach(function(column, index){
          footer_options.text = column.text;
          footer_options.width = column.width;
          footer_options.align = column.align;
          if (footer_options.align === "left") {
            footer_options.multiline = "true";}
          else {footer_options.multiline = "false";}
          if (footers.length-1===index) {footer_options.ln=1;} else {footer_options.ln=0;}
          create_cell(footer_options, header_options);
          self.template.xml.details += "\n    <"+grid_options.xname.toString()+"_footer"+"><![CDATA["+column.text.toString()+
          "]]></"+grid_options.xname.toString()+"_footer"+">";});}}

    function create_grid_header(header_options) {
      header_options.height=0;header_options.x=null;header_options.y=null;
      header_options.columns.forEach(function(column){
        if (header_options.merge===0) {
          header_options.text = column.label;
          header_options.ln = column.ln;
          if (column.width===0) {
            column.width = (header_options.grid_width-header_options.columns_width)/header_options.columns.length;}
          header_options.width = column.width;
        } else {header_options.text += " "+column.label;}
        header_options.align = column.header_align;
        if (header_options.align === "left") {
          header_options.multiline = "true";}
        else {header_options.multiline = "false";}
        create_cell(header_options, header_options);
        if (header_options.width>column.width) {column.width = header_options.width;}});}

    function create_cell(options, header_options){
      //x, y, width, height, text, border, ln, align, padding, multiline, extend
      set_page_style(options);
      if ((typeof options.text==="undefined")||(options.text===null)) {
        options.text = "";}

      if ((typeof options.x==="undefined")||(options.x===null)) {
        options.x = parseFloat(self.template.page.cursor.x);}
      else {options.x = parseFloat(options.x);}
      if ((typeof options.y==="undefined")||(options.y===null)) {
        options.y = parseFloat(self.template.page.cursor.y);}
      else {options.y = parseFloat(options.y);}

      if ((typeof options.padding==="undefined")||(options.padding===null)) {
        options.padding = parseInt(self.template.page.padding,10);}
      else {options.padding = parseInt(options.padding,10);}
      if ((typeof options.multiline==="undefined")||(options.multiline===null)) {
        options.multiline = "false";}

      if ((typeof options.ln==="undefined")||(options.ln===null)) {
        options.ln = 0;}
      else {options.ln = parseInt(options.ln,10);}
      if (options.ln>0 && options.extend) {
        options.width = doc.internal.pageSize.width - self.template.margins["right-margin"] -options.x;}
      else {
        if ((typeof options.width==="undefined")||(options.width===null)) {
          options.width = 
            parseFloat(doc.getStringUnitWidth(options.text)*doc.internal.getFontSize()/doc.internal.scaleFactor+options.padding*2)+1;}
        else {
          if (options.width.toString().indexOf("%")>-1) {
            var nwidth = doc.internal.pageSize.width-self.template.margins["right-margin"]-self.template.margins["left-margin"];
            options.width=nwidth*(parseFloat(options.width.toString().replace("%",""))/100);}
          else {
            options.width = parseFloat(options.width);}}
        if (options.x+options.padding+options.width > doc.internal.pageSize.width - self.template.margins["right-margin"]) {
          options.width = doc.internal.pageSize.width - self.template.margins["right-margin"] -options.x;}}
      if (options.width<0) {options.width=0; options.text="";}

      var theight = parseFloat(doc.internal.getLineHeight()/doc.internal.scaleFactor)+(options.padding);
      if ((typeof options.height==="undefined")||(options.height===null)) {
        options.height = theight; }
      else {
        if (theight > parseFloat(options.height)) {options.height = theight;} else {options.height = parseFloat(options.height);}}

      if (options.multiline==="true" && options.text!=="") {
        var mheight = get_height(options.text,options.width,options.padding,options["font-style"]); options.text = mheight.text;
        if (mheight.height>options.height){options.height = mheight.height;}}

      if (check_page_break(options.y, options.height)) {
        doc.addPage();
        if (typeof header_options!=="undefined"){
          create_grid_header(header_options); options.x = parseFloat(self.template.page.cursor.x);}
        set_page_style(options); options.y = parseFloat(self.template.page.cursor.y);}

      if ((typeof options["background-color"]!=="undefined")&&(options["background-color"]!==null)) {
        doc.rect(options.x, options.y, options.width, options.height, "F");}

      if ((typeof options.border!=="undefined")&&(options.border!==null)) {
        if (options.border==="1") {
          doc.rect(options.x, options.y, options.width, options.height);}
        else {
          if (options.border.indexOf("L")>-1) {
            doc.line(options.x, options.y, options.x, options.y+options.height);}
          if (options.border.indexOf("R")>-1) {
            doc.line(options.x+options.width, options.y, options.x+options.width, options.y+options.height);}
          if (options.border.indexOf("T")>-1) {
            doc.line(options.x, options.y, options.x+options.width, options.y);}
          if (options.border.indexOf("B")>-1) {
            doc.line(options.x, options.y+options.height, options.x+options.width, options.y+options.height);}}}

      var text_x = options.x + options.padding;
      if ((typeof options.align==="undefined")||(options.align===null)) {
        options.align = "left";}
      if (options.multiline==="false") {
        switch(options.align) {
          case "right":
            text_x = options.x + (options.width -parseFloat(doc.getStringUnitWidth(options.text)*doc.internal.getFontSize()/doc.internal.scaleFactor)-options.padding);
            break;
          case "center":
            text_x = options.x + (options.width -parseFloat(doc.getStringUnitWidth(options.text)*doc.internal.getFontSize()/doc.internal.scaleFactor))/2;
            break;
          default:
            //left
            text_x = options.x + options.padding;}
        doc.text(options.text, text_x,options.y + (0.5*options.height+ 0.3*doc.internal.getFontSize()/doc.internal.scaleFactor));}
      else {
        for (var i=0; i<options.text.length; i++) {
          doc.text(options.text[i], text_x, options.y+((doc.internal.getLineHeight()/doc.internal.scaleFactor))*(i+1));}}

      if (options.ln>0) {
        options.y = options.y + options.height;
        if (options.ln===1) {options.x = self.template.margins["left-margin"];}}
      else {
        options.x = options.x + options.width;}
      self.template.page.cursor.x = options.x; self.template.page.cursor.y = options.y;

      return true;}
    
    function create_row(section,row_element) {
      var hgap = get_value(row_element.hgap,0);
      var row_height = get_value(row_element.height,null); var max_height = row_height;
      
      for(var index = 0; index < row_element.columns.length; index++) {
        if (self.template.page.cursor.x!==self.template.margins["left-margin"]) {
          self.template.page.cursor.x+=parseFloat(hgap);}
        var rtype = get_element_type(row_element.columns[index]);
        var element = row_element.columns[index][rtype];
        switch(rtype) {
          case "cell":
            var options = {height:row_height};
            options = get_element_style(element,options);
            options.text = setValue(get_value(element.value,""));
            options.width = get_value(element.width,null);
            options.border = get_value(element.border,null);
            options.align = get_value(element.align,"left");
            if (row_element.columns.length-1===index) {options.ln=1;} else {options.ln=0;}
            options.multiline = "false"; options.extend = true;
            if (section==="details") {
              options.multiline = get_value(element.multiline,"false");}
            create_cell(options);
            if (options.height>max_height || max_height===null) {max_height = options.height;}

            var xname = get_value(element.name,"head");
            if (section==="header" && xname!=="label") {
              self.template.xml.header += "\n    <"+xname.toString()+"><![CDATA["+options.text.toString()+"]]></"+xname.toString()+">";}
            if (section==="details" && xname!=="label") {
              self.template.xml.details += "\n    <"+xname.toString()+"><![CDATA["+options.text.toString()+"]]></"+xname.toString()+">";}
            break;
          case "image":
            var img_src = get_value(element.src,null);
            if (img_src !== null) {
              if (img_src.toString().substr(0,10)!=="data:image") {
                img_src = setValue(img_src);}
              if (img_src.toString().substr(0,10)==="data:image") {
                var width = parseFloat(get_value(element.width,0));
                var img_height = get_value(element.height,0);
                if (img_height===0 || width===0) {
                  var isize = get_img_size(img_src);
                  width = isize.width; img_height = isize.height;}
                var compression = get_value(element.compression,"medium");
                if (check_page_break(self.template.page.cursor.y, img_height)) {
                  doc.addPage();}
                doc.addImage(img_src, "", self.template.page.cursor.x, self.template.page.cursor.y, width, img_height, null, compression);
                self.template.page.cursor.x += width;
                if (max_height<img_height || max_height===null) {max_height = img_height;}}}
            break;
          case "barcode":
            var code_type = get_value(element["code-type"],"code39");
            var code = get_value(element.value,"");
            var visible = get_value(element["visible-value"],0);
            var wide = get_value(element.wide,get_def_barcode_size(code_type, "wide"));
            var narrow = get_value(element.wide,get_def_barcode_size(code_type, "narrow"));
            var bc_height = narrow;
            if (visible===1) {
              bc_height = narrow +(doc.internal.getFontSize()/doc.internal.scaleFactor);}
            if (check_page_break(self.template.page.cursor.y, bc_height)) {
              doc.addPage();}
            switch (code_type) {
              case "code39":
                code39(code, wide, narrow, visible);
                break;
              case "i2of5":
                interleaved2of5(code, wide, narrow, visible);
                break;}
            if (max_height<bc_height || max_height===null) {max_height = bc_height;}
            break;
          case "separator":
            if (max_height!==null) {
              var gap = get_value(element.gap,0);
              doc.line(self.template.page.cursor.x+gap, self.template.page.cursor.y,
                  self.template.page.cursor.x+gap, self.template.page.cursor.y+parseFloat(max_height));
              if (row_element.columns.length-1<index) {self.template.page.cursor.x += gap;}}
            break;}
        if (rtype!=="cell") {
          if (row_element.columns.length-1===index) {
            self.template.page.cursor.y+=max_height;
            self.template.page.cursor.x = self.template.margins["left-margin"];}}}}
    
    function get_def_barcode_size(code_type, size_type) {
      switch (code_type) {
        case "code39":
          if (size_type==="wide")
            return 1.5;
          else return 5;
        case "i2of5":
          if (size_type==="wide")
            return 1;
          else return 10;}
      return 0;}
    
    function check_barchar(bar_char, code) {
      for (var i=0; i<code.toString().length; i++) {
        var char_bar = code[i];
        if (!bar_char.hasOwnProperty(char_bar)) {return false;}}
      return true;}
    
    function code39(txt, w, h, visible) {
      w = w || 1.5; h = h || 5; visible = visible || false;
      var narrow = w / 3.0; var wide = w; var gap = narrow;
      var code = txt.toUpperCase();
      doc.setFontSize(10);
      var start_x = self.template.page.cursor.x;
      var bar_char={'0': 'nnnwwnwnn', '1': 'wnnwnnnnw', '2': 'nnwwnnnnw', '3': 'wnwwnnnnn', '4': 'nnnwwnnnw', '5': 'wnnwwnnnn',
                '6': 'nnwwwnnnn', '7': 'nnnwnnwnw', '8': 'wnnwnnwnn', '9': 'nnwwnnwnn', 'A': 'wnnnnwnnw', 'B': 'nnwnnwnnw',
                'C': 'wnwnnwnnn', 'D': 'nnnnwwnnw', 'E': 'wnnnwwnnn', 'F': 'nnwnwwnnn', 'G': 'nnnnnwwnw', 'H': 'wnnnnwwnn',
                'I': 'nnwnnwwnn', 'J': 'nnnnwwwnn', 'K': 'wnnnnnnww', 'L': 'nnwnnnnww', 'M': 'wnwnnnnwn', 'N': 'nnnnwnnww',
                'O': 'wnnnwnnwn', 'P': 'nnwnwnnwn', 'Q': 'nnnnnnwww', 'R': 'wnnnnnwwn', 'S': 'nnwnnnwwn', 'T': 'nnnnwnwwn',
                'U': 'wwnnnnnnw', 'V': 'nwwnnnnnw', 'W': 'wwwnnnnnn', 'X': 'nwnnwnnnw', 'Y': 'wwnnwnnnn', 'Z': 'nwwnwnnnn',
                '-': 'nwnnnnwnw', '.': 'wwnnnnwnn', ' ': 'nwwnnnwnn', '*': 'nwnnwnwnn', '$': 'nwnwnwnnn', '/': 'nwnwnnnwn',
                '+': 'nwnnnwnwn', '%': 'nnnwnwnwn'};
      if (!check_barchar(bar_char, txt)) {
        var err_str = 'Invalid barcode: '+txt;
        self.template.page.cursor.x+=self.template.page.padding;
        doc.text(err_str, self.template.page.cursor.x, self.template.page.cursor.y+(doc.internal.getLineHeight()/doc.internal.scaleFactor));
        self.template.page.cursor.x+=parseFloat(doc.getStringUnitWidth(err_str)*doc.internal.getFontSize()/doc.internal.scaleFactor);
        self.template.page.cursor.x+=self.template.page.padding;
        return false;}
      
      for(var i = 0; i < code.length; i+=2) {
        var char_bar = code[i]; var seq = '';
        for(var s = 0; s < bar_char[char_bar].length; s++) {
          seq += bar_char[char_bar][s];}
        for(var bar = 0; bar < seq.length; bar++) {
          var line_width;
          if (seq[bar] === 'n') {
            line_width = narrow;}
          else {line_width = wide;}
          if (bar % 2 === 0)
            doc.rect(self.template.page.cursor.x, self.template.page.cursor.y, line_width, h, 'F');
          self.template.page.cursor.x += line_width;}}
      self.template.page.cursor.x += gap;
      if (visible) {
        var txt_w = parseFloat(doc.getStringUnitWidth(txt)*doc.internal.getFontSize()/doc.internal.scaleFactor);
        start_x = start_x+(self.template.page.cursor.x-start_x-txt_w)/2;
        doc.text(txt, start_x, self.template.page.cursor.y+h+(doc.internal.getFontSize()/doc.internal.scaleFactor));}
      return true;}
    
    function interleaved2of5(txt, w, h, visible) {
      w = w || 1; h = h || 10; visible = visible || false;
      var narrow = w / 3.0; var wide = w;
      doc.setFontSize(10);
      var start_x = self.template.page.cursor.x;
      var bar_char={
        '0': 'nnwwn', '1': 'wnnnw', '2': 'nwnnw', '3': 'wwnnn', '4': 'nnwnw', '5': 'wnwnn', '6': 'nwwnn', '7': 'nnnww',
        '8': 'wnnwn', '9': 'nwnwn', 'A': 'nn', 'Z': 'wn'};
      if (!check_barchar(bar_char, txt)) {
        var err_str = 'Invalid barcode: '+txt;
        self.template.page.cursor.x+=self.template.page.padding;
        doc.text(err_str, self.template.page.cursor.x, self.template.page.cursor.y+(doc.internal.getLineHeight()/doc.internal.scaleFactor));
        self.template.page.cursor.x+=parseFloat(doc.getStringUnitWidth(err_str)*doc.internal.getFontSize()/doc.internal.scaleFactor);
        self.template.page.cursor.x+=self.template.page.padding;
        return false;}
      
      var code = txt;
      if (code.length % 2 !== 0)
        code = '0' + code;
      code = 'AA' + code.toLocaleLowerCase() + 'ZA';
      for(var i = 0; i < code.length; i+=2) {
        var char_bar = code[i]; var char_space = code[i+1]; var seq = '';
        for(var s = 0; s < bar_char[char_bar].length; s++) {
          seq += bar_char[char_bar][s] + bar_char[char_space][s];}
        for(var bar = 0; bar < seq.length; bar++) {
          var line_width;
          if (seq[bar] === 'n') {
            line_width = narrow;}
          else {line_width = wide;}
          if (bar % 2 === 0)
            doc.rect(self.template.page.cursor.x, self.template.page.cursor.y, line_width, h, 'F');
          self.template.page.cursor.x += line_width;}}
      if (visible) {
        var txt_w = parseFloat(doc.getStringUnitWidth(txt)*doc.internal.getFontSize()/doc.internal.scaleFactor);
        start_x = start_x+(self.template.page.cursor.x-start_x-txt_w)/2;
        doc.text(txt, start_x, self.template.page.cursor.y+h+(doc.internal.getFontSize()/doc.internal.scaleFactor));}
      return true;}
    
    function parse_value(vname, value) {
      switch(vname) {
        
        case "format":
          value = get_value(value,"a4");
          switch (value) {
            case "a3":
            case "a4":
            case "a5":
            case "letter":
            case "legal":
              break;
            default:
              value = "a4";}
          return value;
        
        case "orientation":
          value = get_value(value,"p");
          switch (value) {
            case "p":
            case "portrait":
            case "l":
            case "landscape":
              break;
            default:
              value = "p";}
          return value;
        
        case "unit":
          value = get_value(value,"mm");
          switch (value) {
            case "pt":
            case "mm":
            case "cm":
            case "in":
              break;
            default:
              value = "mm";}
          return value;
          
        case "author":
        case "creator":
        case "subject":
        case "title":
          value = get_value(value,self.template.document[vname]);
          return value;
        
        //integer
        case "font-size":
        case "digit":
          value = value.toString().replace(/[^0-9\-]|\-(?=.)/g,'');
          value = parseInt(get_value(value,null),10);
          if (!isNaN(value)) {
            return parseInt(get_value(value,null),10);}
          return null;
        
        //float
        case "left-margin":
        case "right-margin":
        case "top-margin":
        case "bottom-margin":
        case "height":
        case "gap":
        case "hgap":
        case "wide":
        case "narrow":
          value = value.toString().replace(/[^0-9\-\.,]|[\-](?=.)|[\.,](?=[0-9]*[\.,])/g,'');
          value = parseFloat(get_value(value,null));
          if (!isNaN(value)) {
            return parseFloat(get_value(value,null));}
          return null;
        
        //percent or float
        case "width":
          value = value.toString().replace(/[^0-9\-\.,%]|[\-](?=.)|[\.,%](?=[0-9]*[\.,%])/g,'');
          if (value==="") {
            value = null;}
          else if (value.indexOf("%")>-1) {
            var pv = parseInt(value.replace("%",""),10);
            if (pv>100) {pv=100;}
            value = pv.toString()+"%";}
          else {
            value = parseFloat(value);}
          return value;
        
        //string
        case "src":
        case "value":
        case "name":
        case "databind":
        case "label":
        case "fieldname":
        case "thousands":
        case "footer":
        case "decpoint":
        case "visible":
        case "border":
        case "html":
        case "code-type":
          return value;
          
        case "merge":
        case "visible-value":
          value = parseInt(get_value(value,0),10);
          if (value===0 || value==="") {
            return 0;} 
          return 1;
        
        case "multiline":
          value = get_value(value,"false");
          if (value==="false" || value==="") {
            return "false";} 
          return "true";
        
        case "compression":
          value = get_value(value,"medium");
          switch (value) {
            case "fast":
            case "medium":
            case "slow":
              break;
            default:
              value = "medium";}
          return value;
          
        case "align":
        case "header-align":
        case "footer-align":
          value = get_value(value,"left");
          switch(value) {
            case "R":
              return "right";
            case "C":
              return "center";
            case "left":
            case "center":
            case "right":
              return value;
            default:
              return "left";}
          
        //check fonts
        case "font-family":
          value = get_value(value,null);
          if (value!=="times" || value!=="helvetica" || value!=="courier") {
            value = "times";}
          return value;
          
        //check font-style
        case "font-style":
          value = get_value(value,null);
          if (value!==null) {
            switch(value) {
              case "B":
                value = "bold";
                break;
              case "I":
                value = "italic";
                break;
              case "BI":
              case "IB":
                value = "bolditalic";
                break;
              case "bold":
              case "italic":
              case "bolditalic":
              case "normal":
                break;
              default:
                value = "normal";}
            return value;}
          return null;
        
        //hex.color
        case "color":
          value = get_value(value,null);
          if (value!==null) {
            if (value.toString().charAt(0)!=="#") {
              if(!isNaN(parseInt(value,10))) {
                value = "#"+parseInt(value,10).toString(16).toUpperCase();
                return value;
              } else {return null;}
            } else {return value;}}
          return null;
          
        //rgb color
        case "border-color":
        case "background-color":
        case "footer-background":
        case "header-background":
          value = get_value(value,null);
          if (value!==null) {
            if (value.toString().charAt(0)==="#") {
              value = hex2greyscale(value); return value;}
            else {
              if(!isNaN(parseInt(value,10))) {
                if (parseInt(value,10)>255) {
                  value = hex2greyscale(parseInt(value,10).toString(16).toUpperCase());} 
                return parseInt(value,10);} 
              else {return null;}}}
          return null;}
      return null;
    }
    
    function copy_attributes(from, to) {
      for(var i = 0; i < from.attributes.length; i++) {
        var value = parse_value(from.attributes[i].name, from.attributes[i].value);
        if (value!==null) {
          to[from.attributes[i].name] = value;}}}
    
    function parse_xml(section, xml) {
      for(var hi = 0; hi < xml.childNodes.length; hi++) {
        switch(xml.childNodes[hi].nodeName) {
          case "row":
            var hrow = xml.childNodes[hi]; var rprop ={}; rprop[xml.childNodes[hi].nodeName] = {columns:[]};
            copy_attributes(hrow, rprop[xml.childNodes[hi].nodeName]);
            var cells = hrow.childNodes;
            for(var rci = 0; rci < cells.length; rci++) {
              if (cells[rci].nodeName!=="#text") {cells = cells[rci];}}
            for(var ri = 0; ri < cells.childNodes.length; ri++) {
              switch(cells.childNodes[ri].nodeName) {
                case "cell":
                case "image":
                case "barcode":
                case "separator":
                  var cprop = {}; cprop[cells.childNodes[ri].nodeName] = {};
                  copy_attributes(cells.childNodes[ri], cprop[cells.childNodes[ri].nodeName]);
                  rprop[xml.childNodes[hi].nodeName].columns.push(cprop);
                  break;}}
            self.template.elements[section].push(rprop);
            break;
          case "datagrid":
            var grow = xml.childNodes[hi]; var gprop ={}; gprop[xml.childNodes[hi].nodeName] = {columns:[]};
            copy_attributes(grow, gprop[xml.childNodes[hi].nodeName]);
            for(var gi = 0; gi < grow.childNodes.length; gi++) {
              switch(grow.childNodes[gi].nodeName) {
                case "header":
                case "footer":
                  var value = parse_value("background-color", grow.childNodes[gi].attributes["background-color"].value);
                  if (value!==null) {
                    gprop[xml.childNodes[hi].nodeName][grow.childNodes[gi].nodeName+"_background"] = value;}
                  break;
                case "columns":
                  var columns = grow.childNodes[gi];
                  for(var ci = 0; ci < columns.childNodes.length; ci++) {
                    if(columns.childNodes[ci].nodeName==="column") {
                      var cgprop = {}; cgprop[columns.childNodes[ci].nodeName] = {};
                    copy_attributes(columns.childNodes[ci], cgprop[columns.childNodes[ci].nodeName]);
                    gprop[xml.childNodes[hi].nodeName].columns.push(cgprop);}}
                  break;}}
            self.template.elements[section].push(gprop);
            break;
          case "html":
            var html = {}; html[xml.childNodes[hi].nodeName] = {};
            copy_attributes(xml.childNodes[hi], html[xml.childNodes[hi].nodeName]);
            html[xml.childNodes[hi].nodeName].html = xml.childNodes[hi].textContent || "";
            self.template.elements[section].push(html);
            break;
          case "vgap":
          case "hline":
            var hprop = {}; hprop[xml.childNodes[hi].nodeName] = {};
            copy_attributes(xml.childNodes[hi], hprop[xml.childNodes[hi].nodeName]);
            self.template.elements[section].push(hprop);
            break;}}}
    
    function copy_properties(from, to) {
      for(var i = 0; i < Object.keys(from).length; i++) {
        var pname = Object.keys(from)[i];
        var value = parse_value(pname, from[pname]);
        if (value!==null) {
          to[pname] = value;}}}
    
    function parse_json(section, obj) {
      for(var hi = 0; hi < obj.length; hi++) {
        var otype = get_element_type(obj[hi]);
        switch(otype) {
          case "row":
            var hrow = obj[hi][otype]; var rprop ={}; rprop[otype] = {columns:[]};
            copy_properties(hrow, rprop[otype]);
            for(var ri = 0; ri < hrow.columns.length; ri++) {
              var ctype = get_element_type(hrow.columns[ri]);
              var cprop = {}; cprop[ctype] = {};
              copy_properties(hrow.columns[ri][ctype], cprop[ctype]);
              rprop[otype].columns.push(cprop);}
            self.template.elements[section].push(rprop);
            break;
          case "datagrid":
            var grow = obj[hi][otype]; var gprop ={}; gprop[otype] = {columns:[]};
            copy_properties(grow, gprop[otype]);
            for(var ci = 0; ci < grow.columns.length; ci++) {
              var column = grow.columns[ci].column;
              var cgprop = {}; cgprop.column = {};
              copy_properties(column, cgprop.column);
              gprop.datagrid.columns.push(cgprop);}
            self.template.elements[section].push(gprop);
            break;
          case "html":
            var html = {}; html[otype] = {};
            copy_properties(obj[hi][otype], html[otype]);
            html[otype].html = obj[hi][otype].html;
            self.template.elements[section].push(html);
            break;
          case "vgap":
          case "hline":
            var hprop = {}; hprop[otype] = {};
            copy_properties(obj[hi][otype], hprop[otype]);
            self.template.elements[section].push(hprop);
            break;}}}
    
    function get_parent_section(element, section) {
      for(var i = 0; i < section.length; i++) {
        if (element === section[i]) {return section;}
        else {
          var etype = get_element_type(section[i]);
          if (etype==="row" || etype==="datagrid") {
            for(var ri = 0; ri < section[i][etype].columns.length; ri++) {
              if (element === section[i][etype].columns[ri]) {return section[i][etype].columns;}
            }}}}
      return null;}
    
    function get_parent(element) {
      var parent = get_parent_section(element, self.template.elements.details);
      if (parent === null) {
        parent = get_parent_section(element, self.template.elements.header);}
      if (parent === null) {
        parent = get_parent_section(element, self.template.elements.footer);}
      return parent;}
    
    function create_xml_element(obj, xdoc, parent) {
      for(var i = 0; i < obj.length; i++) {
        var etype = get_element_type(obj[i]);
        var element = parent.appendChild(xdoc.createElement(etype));
        for(var ai = 0; ai < Object.keys(obj[i][etype]).length; ai++) {
          var pname = Object.keys(obj[i][etype])[ai];
          var pval = obj[i][etype][pname];
          if (typeof pval === "object") {
            if (Array.isArray(pval)) {
              var columns = element.appendChild(xdoc.createElement(pname));
              create_xml_element(pval, xdoc, columns);
            } else {element.setAttribute(pname, pval);}} 
          else {
            if (pname==="html") {element.appendChild(xdoc.createCDATASection(pval));}
            else {element.setAttribute(pname, pval);}}}}}
    
    function report_ini() {
      self.template.document.author = self.template.elements.report.author||self.template.document.author;
      self.template.document.creator = self.template.elements.report.creator||self.template.document.creator;
      self.template.document.subject = self.template.elements.report.subject||self.template.document.subject;
      self.template.document.title = self.template.elements.report.title||self.template.document.title;

      if (typeof self.template.elements.report["font-family"]!=="undefined"){
        self.template.style["font-family"] = self.template.elements.report["font-family"];}
      if (typeof self.template.elements.report["font-style"]!=="undefined"){
        self.template.style["font-style"] = self.template.elements.report["font-style"];}
      if (typeof self.template.elements.report["font-size"]!=="undefined"){
        self.template.style["font-size"] = self.template.elements.report["font-size"];}

      if (typeof self.template.elements.report.color!=="undefined"){
        self.template.style.color = self.template.elements.report.color;}
      if (typeof self.template.elements.report["border-color"]!=="undefined"){
        self.template.style["border-color"] = self.template.elements.report["border-color"];}
      if (typeof self.template.elements.report["background-color"]!=="undefined"){
        self.template.style["background-color"] = self.template.elements.report["background-color"];}

      self.template.margins["left-margin"] = self.template.elements.report["left-margin"]||self.template.margins["left-margin"];
      self.template.margins["right-margin"] = self.template.elements.report["right-margin"]||self.template.margins["right-margin"];
      self.template.margins["top-margin"] = self.template.elements.report["top-margin"]||self.template.margins["top-margin"];
      self.template.margins["bottom-margin"] = self.template.elements.report["bottom-margin"]||self.template.margins["bottom-margin"];}
    
  //public functions
    self.loadJsonDefinition = function(data) {
      if ((data===null) || (data==="") || (typeof data==="undefined")) {return true;}
      if (typeof data==="string") {
        data = JSON.parse(data);}
      self.template.elements = {report: {}, header: [], details: [], footer: [], data: {}};
      self.template.element_matrix = [];
      if (typeof data.report !== "undefined") {
        copy_properties(data.report, self.template.elements.report);}
      if (typeof data.header !== "undefined") {
          parse_json("header", data.header);}
      if (typeof data.details !== "undefined") {
          parse_json("details", data.details);}
      if (typeof data.footer !== "undefined") {
          parse_json("footer", data.footer);}
      if (typeof data.data !== "undefined") {
        var db = self.template.elements.data;
        for(var i = 0; i < Object.keys(data.data).length; i++) {
          var pname = Object.keys(data.data)[i];
          if (typeof data.data[pname] === "object") {
            if (!Array.isArray(data.data[pname])) {
              db[pname] = {};
              for(var pi = 0; pi < Object.keys(data.data[pname]).length; pi++) {
                db[pname][Object.keys(data.data[pname])[pi]] = data.data[pname][Object.keys(data.data[pname])[pi]];}
            } else {
              db[pname] = [];
              for(var ai = 0; ai < data.data[pname].length; ai++) {
                var element = {};
                for(var pai = 0; pai < Object.keys(data.data[pname][ai]).length; pai++) {
                  element[Object.keys(data.data[pname][ai])[pai]] = data.data[pname][ai][Object.keys(data.data[pname][ai])[pai]];}
                db[pname].push(element);}}}
          else {
            db[pname] = data.data[pname];}}}
      report_ini(); return true;};
    
    self.loadDefinition = function(data) {
      if ((data!==null) && (data!=="") && (typeof data==="string")) {
        self.template.elements = {report: {}, header: [], details: [], footer: [], data: {}};
        self.template.element_matrix = [];
        var xml_elements = new DOMParser().parseFromString(data, "application/xml");
        
        if (xml_elements.getElementsByTagName("report").length>0) {
          copy_attributes(xml_elements.getElementsByTagName("report")[0], self.template.elements.report);}
            
        if (xml_elements.getElementsByTagName("header").length>0) {
          parse_xml("header", xml_elements.getElementsByTagName("header")[0]);}
        if (xml_elements.getElementsByTagName("details").length>0) {
          parse_xml("details", xml_elements.getElementsByTagName("details")[0]);}
        if (xml_elements.getElementsByTagName("footer").length>0) {
          parse_xml("footer", xml_elements.getElementsByTagName("footer")[0]);}
        if (xml_elements.getElementsByTagName("data").length>0) {
          var db = self.template.elements.data;
          var xdata = xml_elements.getElementsByTagName("data")[0];
          for(var i = 0; i < xdata.childNodes.length; i++) {
            if (xdata.childNodes[i].nodeName!=="#text") {
              if (xdata.childNodes[i].childNodes.length>0) {
                if (xdata.childNodes[i].childNodes.length===1 && xdata.childNodes[i].childNodes[0].nodeName==="#text") {
                  db[xdata.childNodes[i].nodeName] = xdata.childNodes[i].childNodes[0].textContent;
                } else {
                  if (xdata.childNodes[i].childNodes.length===1 && 
                    xdata.childNodes[i].childNodes[0].nodeName==="#cdata-section") {
                    db[xdata.childNodes[i].nodeName] = xdata.childNodes[i].childNodes[0].data;
                  } else {
                    db[xdata.childNodes[i].nodeName] = [];
                    for(var ci = 0; ci < xdata.childNodes[i].childNodes.length; ci++) {
                      if (xdata.childNodes[i].childNodes[ci].nodeName!=="#text") {
                        var xelement = xdata.childNodes[i].childNodes[ci]; var xobj = {};
                      for(var cai = 0; cai < xelement.attributes.length; cai++) {
                        xobj[xelement.attributes[cai].nodeName] = xelement.attributes[cai].value;}
                      db[xdata.childNodes[i].nodeName].push(xobj);}}}}
              } else {
                if (xdata.childNodes[i].attributes.length>0) {
                  db[xdata.childNodes[i].nodeName] = {};
                  for(var ai = 0; ai < xdata.childNodes[i].attributes.length; ai++) {
                    db[xdata.childNodes[i].nodeName][xdata.childNodes[i].attributes[ai].name] = xdata.childNodes[i].attributes[ai].value;}} 
                else {
                  self.template.elements.data[xdata.childNodes[i].nodeName] = xdata.childNodes[i].textContent;}}}}
          orig_data = self.template.elements.data;}
        report_ini(); return true;}
      return false;};

    self.setData = function(key, value) {
      if ((typeof value === "object") && (typeof self.template.elements.data[key] === "object")){
        if (!Array.isArray(value) && !Array.isArray(self.template.elements.data[key])) {
          for(var pi = 0; pi < Object.keys(value).length; pi++) {
            var pname = Object.keys(value)[pi];
            self.template.elements.data[key][pname] = value[pname];}}
        else {
          self.template.elements.data[key] = value;}}
      else {
        self.template.elements.data[key] = value;}
      return true;};
    
    self.getData = function(key) {
      return self.template.elements.data[key];};
    
    self.createReport = function() {
      if (self.template.elements.details.length>0) {
        self.template.page.cursor = {x:self.template.margins["left-margin"], y:self.template.margins["top-margin"]};
        doc.setProperties({
          author: self.template.document.author, creator: self.template.document.creator,
          subject: self.template.document.subject, title: self.template.document.title});
        set_page_style({}); var footer_height = get_footer_height();
        self.template.page.cursor.y = doc.internal.pageSize.height-self.template.margins["bottom-margin"]-footer_height;
        footer();
        self.template.page.cursor.y = self.template.margins["top-margin"]; self.template.page.footer_height = footer_height;
        self.template.page.cursor.x = self.template.margins["left-margin"];
        header(); self.template.page.header_height = self.template.page.cursor.y - self.template.margins["top-margin"];
        var edetails = self.template.elements.details;
          for(var i = 0; i < edetails.length; i++) {
            var etype = get_element_type(edetails[i]);
            create_elements("details", etype, edetails[i][etype]);}}
      return true;};

    self.save2PdfFile = function(filename) {
      doc.save(filename);};

    self.save2Pdf = function() {
      return doc.output('arraybuffer');};

    self.save2DataUrl = function() {
      return doc.output('dataurlnewwindow');};
    
    self.save2DataUrlString = function() {
      return doc.output('dataurlstring');};
    
    self.save2Blob = function() {
      return doc.output('blob');};
    
    self.save2BlobUrl = function() {
      return doc.output('bloburl');};
    
    self.save2Xml = function() {
      return "<data>"+self.template.xml.header+self.template.xml.details+"\n</data>";};
    
    self.getXmlTemplate = function() {
      var xml_str = "<template><report/><header></header><details></details><footer></footer><data></data></template>";
      var xml_elements = new DOMParser().parseFromString(xml_str, "application/xml");
      
      for(var i = 0; i < Object.keys(self.template.elements.report).length; i++) {
        var pname = Object.keys(self.template.elements.report)[i];
        xml_elements.getElementsByTagName("report")[0].setAttribute(pname, self.template.elements.report[pname]);}
      create_xml_element(self.template.elements.header, xml_elements, xml_elements.getElementsByTagName("header")[0]);
      create_xml_element(self.template.elements.details, xml_elements, xml_elements.getElementsByTagName("details")[0]);
      create_xml_element(self.template.elements.footer, xml_elements, xml_elements.getElementsByTagName("footer")[0]);
      
      for(var i = 0; i < Object.keys(self.template.elements.data).length; i++) {
        var ename = Object.keys(self.template.elements.data)[i];
        var element = xml_elements.getElementsByTagName("data")[0].appendChild(
          xml_elements.createElement(ename));
        if (typeof self.template.elements.data[ename] === "object") {
          if (!Array.isArray(self.template.elements.data[ename])) {
            for(var pi = 0; pi < Object.keys(self.template.elements.data[ename]).length; pi++) {
              var pname = Object.keys(self.template.elements.data[ename])[pi];
              element.setAttribute(pname, self.template.elements.data[ename][pname]);}
          } else {
            for(var ai = 0; ai < self.template.elements.data[ename].length; ai++) {
              var erow = element.appendChild(xml_elements.createElement(ename));
              for(var pai = 0; pai < Object.keys(self.template.elements.data[ename][ai]).length; pai++) {
                var epname = Object.keys(self.template.elements.data[ename][ai])[pai];
                erow.setAttribute(epname, self.template.elements.data[ename][ai][epname]);}}}}
        else {
          element.appendChild(xml_elements.createCDATASection(self.template.elements.data[ename]));}}
      
      return new XMLSerializer().serializeToString(xml_elements);};
    
    self.insertElement = function(parent, ename, index, values) {
      if ((typeof parent==="undefined")||(parent===null)) {parent=self.template.elements.details;}
      if ((typeof ename==="undefined")||(ename===null)) {ename="row";}
      if ((typeof index==="undefined")||(index===null)) {index=-1;}
      if ((typeof values==="undefined")||(values===null)) {values={};}
      if (typeof parent === "object") {
        if (!Array.isArray(parent)) {
          var ptype = get_element_type(parent);
          if (ptype==="row" || ptype==="datagrid") {parent = parent[ptype].columns;}
        }} else {return null;}
      var element = {}; element[ename] = {};
      switch (ename) {
        case "row":
        case "datagrid":
          element[ename].columns=[];
          break;
        case "vgap":
        case "hline":
        case "html":
        case "column":
        case "cell":
        case "image":
        case "barcode":
        case "separator":
          break;
        default:
          return null;}
      if (index>parent.length || index<0) {index=parent.length;}
      parent.splice(index, 0, element);
      for(var i = 0; i < Object.keys(values).length; i++) {
        var pname = Object.keys(values)[i];
        element[ename][pname] = parse_value(pname,values[pname]);}
      return element;};
    
    self.editElement = function(element, values) {
      var ptype = get_element_type(element);
      if (ptype!==null) {
        for(var i = 0; i < Object.keys(values).length; i++) {
          var pname = Object.keys(values)[i];
          if (element[ptype].hasOwnProperty(pname) && values[pname]===null) {
            delete element[ptype][pname];
          } else {
            element[ptype][pname] = values[pname];}}
        return true;
      } else {return false;}};
    
    self.deleteElement = function(element, parent) {
      if ((typeof parent==="undefined")||(parent===null)) {
        parent = get_parent(element);}
      if (parent!==null) {
        if (!Array.isArray(parent)) {
          var ptype = get_element_type(parent);
          if (ptype==="row" || ptype==="datagrid") {parent = parent[ptype].columns;}
          else {return false;}}}
      if (parent!==null) {
        var index = parent.indexOf(element);
        if (index > -1) {parent.splice(index, 1);}
        return true;}
      return false;};
  
  };
});
