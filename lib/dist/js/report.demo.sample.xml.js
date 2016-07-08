/*
This file is part of the Nervatura Framework
http://www.nervatura.com
Copyright © 2011-2016, Csaba Kappel
License: LGPLv3
https://raw.githubusercontent.com/nervatura/nervatura/master/LICENSE
*/

define([], function () {
  function getXml(f) {
    return f.toString().replace(/^[^\/]+\/\*!?/, '').replace(/\*\/[^\/]+$/, '');}

  return {
    Template: getXml(function() {
      /*!<template>
      <report title="Nervatura Report" left-margin="15" top-margin="15" right-margin="15" />
      <header>
        <row height="10">
          <columns>
            <image src="logo" />
            <cell name="label" value="labels.title" font-style="bolditalic" font-size="26" color="#D8DBDA" />
            <cell name="label" value="XML Sample" align="right" font-style="bold" />
          </columns>
        </row>
        <vgap height="2" />
        <hline border-color="100" />
        <vgap height="2" />
      </header>
      <details>
        <vgap height="2" />
        <row>
          <columns>
            <cell name="label" width="50%" font-style="bold" value="labels.left_text" border="LT" border-color="100" background-color="230" />
            <cell name="label" font-style="bold" value="labels.left_text" border="LTR" border-color="100" background-color="230" />
          </columns>
        </row>
        <row>
          <columns>
            <cell name="short_text" width="50%" value="head.short_text" border="L" border-color="100" />
            <cell name="short_text" value="head.short_text" border="LR" border-color="100" />
          </columns>
        </row>
        <row>
          <columns>
            <cell name="short_text" width="50%" value="head.short_text" border="LB" border-color="100" />
            <cell name="short_text" value="head.short_text" border="LBR" border-color="100" />
          </columns>
        </row>
        <row>
          <columns>
            <cell name="label" width="40" font-style="bold" value="labels.left_text" border="LB" border-color="100" />
            <cell name="label" align="center" width="30" font-style="bold" value="labels.center_text" border="LB" border-color="100" />
            <cell name="label" align="right" width="40" font-style="bold" value="labels.right_text" border="LB" border-color="100" />
            <cell name="label" font-style="bold" value="labels.left_text" border="LBR" border-color="100" />
          </columns>
        </row>
        <row>
          <columns>
            <cell name="short_text" width="40" value="head.short_text" border="LB" border-color="100" />
            <cell name="date" align="center" width="30" value="head.date" border="LB" border-color="100" />
            <cell name="amount" align="right" width="40" value="head.number" border="LB" border-color="100" />
            <cell name="short_text" value="head.short_text" border="LBR" border-color="100" />
          </columns>
        </row>
        <row>
          <columns>
            <cell name="label" font-style="bold" value="labels.left_text" border="LB" border-color="100" />
            <cell name="short_text" width="50" value="head.short_text" border="LB" border-color="100" />
            <cell name="label" font-style="bold" value="labels.left_text" border="LB" border-color="100" />
            <cell name="short_text" value="head.short_text" border="LBR" border-color="100" />
          </columns>
        </row>
        <row>
          <columns>
            <cell name="long_text" multiline="true" value="head.long_text" border="LBR" border-color="100" />
          </columns>
        </row>
        <vgap height="2" />
        <row hgap="2">
          <columns>
            <cell name="label" value="labels.left_text" font-style="bold" border="1" border-color="100" background-color="230" />
            <cell name="short_text" value="head.short_text" border="1" border-color="100" />
            <cell name="label" value="labels.left_text" font-style="bold" border="1" border-color="100" background-color="230" />
            <cell name="short_text" value="head.short_text" border="1" border-color="100" />
          </columns>
        </row>
        <vgap height="2" />
        <row hgap="2">
          <columns>
            <cell name="label" value="labels.long_text" font-style="bold" border="1" border-color="100" background-color="230" />
            <cell name="long_text" multiline="true" value="head.long_text" border="1" border-color="100" />
          </columns>
        </row>
        <vgap height="2"/>
        <row hgap="2">
          <columns>
            <cell name="label" value="labels.long_text" font-style="bold" border="1" border-color="100" background-color="230"/>
            <cell name="long_text" multiline="true" value="head.long_text" border="1" border-color="100"/>
          </columns>
        </row>
        <vgap height="2"/>
        <hline border-color="100"/>
        <vgap height="2"/>
        <row hgap="3">
          <columns>
            <cell name="label" value="Barcode (Interleaved 2of5)" font-style="bold" font-size="10" 
              border="1" border-color="100" background-color="230"/>
            <barcode code-type="ITF" value="1234567890" visible-value="1"/>
            <cell name="label" value="Barcode (Code 39)" font-style="bold" font-size="10" 
              border="1" border-color="100" background-color="230"/>
            <barcode code-type="CODE_39" value="1234567890ABCDEF" visible-value="1"/>
          </columns>
        </row>
        <vgap height="3"/>
        <row>
          <columns>
            <cell name="label" value="Datagrid Sample" align="center" font-style="bold" border="1" border-color="100" 
                  background-color="230"/>
          </columns>
        </row>
        <vgap height="2" />
        <datagrid name="items" databind="items" border="1" border-color="100" header-background="230" footer-background="230">
          <columns>
            <column width="8%" fieldname="counter" align="right" label="labels.counter" footer="labels.total" />
            <column width="20%" fieldname="date" align="center" label="labels.center_text" />
            <column width="15%" fieldname="number" align="right" label="labels.right_text" 
                    footer="items_footer.items_total" footer-align="right" />
            <column fieldname="text" label="labels.left_text" />
          </columns>
        </datagrid>
        <vgap height="5"/>
        <html fieldname="html_text"><![CDATA[<i>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</i>
          ={{html_text}} Nulla a pretium nunc, in cursus quam.]]></html>
      </details>
      <footer>
        <vgap height="2" />
        <hline border-color="100" />
        <row height="10">
          <columns>
            <cell value="Nervatura Report Template" font-style="bolditalic" />
            <cell value="{{page}}" align="right" font-style="bold" />
          </columns>
        </row>
      </footer>
      <data>
        <labels title="REPORT TEMPLATE" left_text="Short text" center_text="Centered text" right_text="Right text" long_text="Long text"
                counter="No." total="Total"/>
        <head short_text="Lorem ipsum dolor" number="123 456" date="2015.01.01" 
                long_text="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim. Nulla a pretium nunc, in cursus quam." />
        <html_text><![CDATA[<p><b>Pellentesque eu mattis diam, sed dapibus justo. In eget augue nisi. Cras eget odio vel mi vulputate interdum. Curabitur consequat sapien at lacus tincidunt, at sagittis felis lobortis. Aenean porta maximus quam eu porta. Fusce sed leo ut justo commodo facilisis. Vivamus vitae tempor erat, at ultrices enim.</b></p>]]></html_text>
        <items>
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
          <items text="Lorem ipsum dolor" number="123 456" date="2015.01.01" />
        </items>
        <items_footer items_total="3 703 680" />        
        <logo>data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABoAAAAaCAYAAACpSkzOAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAFxGAABcRgEUlENBAAAAB3RJTUUH4AcCFiIAfaA8WwAAAeBJREFUSMftlj9IVlEUwH/nfR/+G3QwaPlAg3D5pKkPDIoIghYVm4yCFKlPcGoQdGppa4mm4IJbQdFYRKDgoFR0o5ZwcL5LLoKCoPW+43JfvO/6ns9P0ckDj8c5l3t+55x73rkPzqUFqdRrTe/Tho0CRGeQ2FWA8ilmcgH4CUyeGFSp13DGpvUOoB+YSwDALwBp0VF6jwAKXAbGgIfAYOBi1Rl7oxDknXf56IaBKtDhnzagvcBHzRn7IxOUZFGp12aAaeDKMSv7D+gBdpyxzSBfmnZgGbh2wn5Yc8ZWEyXKOI+lY0C2RGQ9sD1OK/9BvlyjwPUWIZsi8lFVB1K2DWfsl/REKAdle94i5LuIbKnq/cC+mAR/ICMvvUcliMgroFNVb2csT4WGcoGeJbsisqCq4zmBLThj90JjmNFeQbt+i0ReqOpMDiR2xj7KmtYh6E/WZhH5WoqiCWCooTp/SDC3wrPJ7Dr/9Sdn8L5UKt0T5JOqVuNG401BSaedsSt5d08ULFwENoBNVb0Zx/FbRUeA7gLIU2esCWdjU/P41h4HZv3d8Q74DTw7QmP8Baacsa8Pg6S7bBvoSw9BP0yfAJ05gM/AXSDOO5cDGeVdC972ALgDXAJ2gQ/O2Jfnf0RnJvumbKT0gnMTFgAAAABJRU5ErkJggg==</logo>
      </data>
    </template>
    */
    })
  }

});
