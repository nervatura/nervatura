import { memo, Fragment } from 'react';
import PropTypes from 'prop-types';

import SideBar from 'components/SideBar/Template'
import TemplateEditor from 'components/Report/PDF/TemplateEditor'

export const TemplateView = ({
  data, current, onEvent, getText
}) => {
  return (
    <Fragment>
      <SideBar side={current.side} 
        templateKey={data.key} dirty={data.dirty} 
        onEvent={onEvent} getText={getText} />
      <div className={`${"page padding-normal"} ${current.theme}`} >
        <TemplateEditor data={data} onEvent={onEvent} getText={getText} />
      </div>
    </Fragment>
  )
}

TemplateView.propTypes = {
  data: PropTypes.object.isRequired,
  current: PropTypes.object.isRequired,
  onEvent: PropTypes.func,
  getText: PropTypes.func
}

TemplateView.defaultProps = {
  data: {}, 
  current: {},
  onPage: undefined,
  getText: undefined
}

export default memo(TemplateView, (prevProps, nextProps) => {
  return (
    (prevProps.data === nextProps.data) &&
    (prevProps.current.side === nextProps.current.side)
  )
})

export const templateElements = ({ getText }) => {
  return {
    report:{
      options: {
        title: "REPORT"},
      rows: [
        {rowtype:"flip", name:"title", datatype:"string", default:"Nervatura Report"},
        {rowtype:"flip", name:"author", datatype:"string"},
        {rowtype:"flip", name:"creator", datatype:"string"},
        {rowtype:"flip", name:"subject", datatype:"string"},
        {rowtype:"flip", name:"keywords", datatype:"string"},
        {rowtype:"groupline"},
        {rowtype:"flip", name:"font-style", datatype:"select", default: "",
          options: [["",""],["bold","bold"],["italic","italic"],["bolditalic","bolditalic"]], 
          info: getText("info_font-style")},
        {rowtype:"flip", name:"font-size", datatype:"integer", default: 12},
        {rowtype:"flip", name:"color", datatype:"color", info:getText("info_color")},
        {rowtype:"flip", name:"border-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_border-color")},
        {rowtype:"flip", name:"background-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_background-color")},
        {rowtype:"groupline"},
        {rowtype:"flip", name:"left-margin", datatype:"integer", default: 12},
        {rowtype:"flip", name:"right-margin", datatype:"integer", default: 12},
        {rowtype:"flip", name:"top-margin", datatype:"integer", default: 12}
      ]
    },
    row:{
      options: {
        title: "ROW"},
      rows: [
        {rowtype:"flip", name:"height", datatype:"float", default: 0},
        {rowtype:"flip", name:"hgap", datatype:"integer", default: 0, info: getText("info_hgap")},
        {rowtype:"flip", name:"visible", datatype:"string", info: getText("info_visible")}
      ]
    },
    cell:{
      options: {
        title: "CELL"},
      rows: [
        {rowtype:"flip", name:"name", datatype:"string", default: "head", info: getText("info_name")},
        {rowtype:"flip", name:"value", datatype:"string", info: getText("info_value")},
        {rowtype:"flip", name:"width", datatype:"percent", info: getText("info_width")},
        {rowtype:"flip", name:"align", datatype:"select", default: "left",
          options: [["left","left"],["right","right"],["center","center"]], info: getText("info_align")},
        {rowtype:"flip", name:"multiline", datatype:"select", default: "false",
          options: [["false","false"],["true","true"]], info: getText("info_multiline")},
        {rowtype:"groupline"},
        {rowtype:"flip", name:"font-style", datatype:"select", default: "",
          options: [["",""],["bold","bold"],["italic","italic"],["bolditalic","bolditalic"]], 
          info: getText("info_font-style")},
        {rowtype:"flip", name:"font-size", datatype:"integer", default: 12},
        {rowtype:"flip", name:"color", datatype:"color", info:getText("info_color")},
        {rowtype:"flip", name:"border-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_border-color")},
        {rowtype:"flip", name:"background-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_background-color")},
        {rowtype:"flip", name:"border", datatype:"checklist", 
          values: ["1|All", "L|Left", "T|Top", "R|Right", "B|Bottom"]}
      ]
    },
    image:{
      options: {
        title: "IMAGE"},
      rows: [
        {rowtype:"flip", name:"src", datatype:"image", info: getText("info_src")},
        {rowtype:"flip", name:"height", datatype:"float", default: 0, info: getText("info_height_image")}
      ]
    },
    barcode:{
      options: {
        title: "BARCODE"},
      rows: [
        {rowtype:"flip", name:"code-type", datatype:"select", default: "ITF",
          options: [["ITF","ITF"],["CODE_39","CODE_39"],["CODE_128","CODE_128"],["EAN","EAN"],["QR","QR"]], info: getText("info_code-type")},
        {rowtype:"flip", name:"value", datatype:"string", default: "", info: getText("info_barcode_value")},
        {rowtype:"flip", name:"visible-value", datatype:"select", default: "0",
          options: [["0","0"],["1","1"]], info: getText("info_visible-value")},
        {rowtype:"flip", name:"wide", datatype:"float", default: 0, info: getText("info_optional")},
        {rowtype:"flip", name:"narrow", datatype:"float", default: 0, info: getText("info_optional")}
          ]
    },
    separator:{
      options: {
        title: "SEPARATOR"},
      rows: [
        {rowtype:"flip", name:"hgap", datatype:"integer", default: 0, info: getText("info_hgap")}]
    },
    vgap:{
      options: {
        title: "VGAP"},
      rows: [
        {rowtype:"flip", name:"height", datatype:"float", default: 0, info: getText("info_height")}]
    },
    hline:{
      options: {
        title: "HLINE"},
      rows: [
        {rowtype:"flip", name:"width", datatype:"percent", info: getText("info_width")},
        {rowtype:"flip", name:"gap", datatype:"integer", default:0, info: getText("info_gap")},
        {rowtype:"flip", name:"border-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_border-color")}
        ]
    },
    html:{
      options: {
        title: "HTML"},
      rows: [
        {rowtype:"flip", name:"html", datatype:"text", default: "", info: getText("info_html")}]
    },
    datagrid:{
      options: {
        title: "DATAGRID"},
      rows: [
        {rowtype:"flip", name:"name", datatype:"string", default: "items", info: getText("info_datagrid_name")},
        {rowtype:"flip", name:"databind", datatype:"string", default: "", info: getText("info_databind")},
        {rowtype:"flip", name:"width", datatype:"percent", info: getText("info_width")},
        {rowtype:"flip", name:"merge", datatype:"select", default: "0",
          options: [["0","0"],["1","1"]], info: getText("info_merge")},
        {rowtype:"flip", name:"font-size", datatype:"integer", default: 12},
        {rowtype:"flip", name:"border", datatype:"checklist", 
          values: ["1|All", "L|Left", "T|Top", "R|Right", "B|Bottom"]},
        {rowtype:"flip", name:"color", datatype:"color", info:getText("info_color")},
        {rowtype:"flip", name:"border-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_border-color")},
        {rowtype:"flip", name:"background-color", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_background-color")},
        {rowtype:"flip", name:"header-background", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_background-color")},
        {rowtype:"flip", name:"footer-background", datatype:"integer", default: 0, max: 255, min: 0,
          info: getText("info_background-color")}
        ]
    },
    column:{
      options: {
        title: "COLUMN"},
      rows: [
        {rowtype:"flip", name:"fieldname", datatype:"string", default: "", info: getText("info_fieldname")},
        {rowtype:"flip", name:"label", datatype:"string", default: "", info: getText("info_label")},
        {rowtype:"flip", name:"width", datatype:"percent", info: getText("info_width")},
        {rowtype:"flip", name:"align", datatype:"select", default: "left",
          options: [["left","left"],["right","right"],["center","center"]], info: getText("info_align")},
        {rowtype:"flip", name:"header-align", datatype:"select", default: "left",
          options: [["left","left"],["right","right"],["center","center"]], info: getText("info_align")},
        {rowtype:"flip", name:"footer-align", datatype:"select", default: "left",
          options: [["left","left"],["right","right"],["center","center"]], info: getText("info_align")},
        {rowtype:"flip", name:"thousands", datatype:"string", default: "", info: getText("info_thousands")},
        {rowtype:"flip", name:"digit", datatype:"integer", default: 0, info: getText("info_digit")},
        {rowtype:"flip", name:"footer", datatype:"string", default: "", info: getText("info_footer")}
      ]
    },
    header:{
      options: {
        title: "HEADER"},
      rows: []
    },
    details:{
      options: {
        title: "DETAILS"},
      rows: []
    },
    footer:{
      options: {
        title: "FOOTER"},
      rows: []
    }
  }
}