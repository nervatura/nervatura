import { TemplateEditor } from "./TemplateEditor";

import { getText as appGetText, store } from 'config/app';
import { templateElements } from 'containers/Template/Template'

export default {
  title: "Report/TemplateEditor",
  component: TemplateEditor,
}

const Template = (args) => <TemplateEditor {...args} />
const sample = require('../../../config/sample.json')
const getText = (key)=>appGetText({ locales: store.session.locales, lang: "en", key: key })

export const Default = Template.bind({});
Default.args = {
  data: {
    title: "Sample Report", 
    tabView: "template", 
    template: {
      ...sample,
      sources: {
        head: {
          default: "select * from table"
        }
      },
      footer: []
    }, 
    current: {
      id: "tmp_report",
      section: "report",
      type: "report",
      item: sample.report,
      index: null,
      parent: null,
      parent_type: null,
      parent_index: null,
      form: templateElements({ getText: getText })["report"]
    }, 
    current_data: null, 
    dataset: [
      { lslabel: "labels", lsvalue: "list" },
      { lslabel: "head", lsvalue: "list" },
      { lslabel: "html_text", lsvalue: "string" },
      { lslabel: "items", lsvalue: "table" },
      { lslabel: "items_footer", lsvalue: "list" },
      { lslabel: "logo", lsvalue: "string" }
    ]
  },
  className: "light",
  onEvent: undefined,
  getText: getText
}

export const Details = Template.bind({});
Details.args = {
  ...Default.args,
  className: "dark",
  data: {
    ...Default.args.data,
    current: {
      id: "tmp_details",
      section: "details",
      type: "details",
      item: sample.details,
      index: null,
      parent: null,
      parent_type: null,
      parent_index: null,
      form: templateElements({ getText: getText })["details"]
    },
  } 
}

export const Row = Template.bind({});
Row.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    current: {
      id: "tmp_details_1_row",
      section: "details",
      type: "row",
      item: sample.details[1].row.columns,
      item_base: sample.details[1].row,
      index: 1,
      parent: sample.details,
      parent_type: "details",
      parent_index: null,
      form: templateElements({ getText: getText })["row"],
      add_item: "cell"
    },
  } 
}

export const Cell = Template.bind({});
Cell.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    current: {
      id: "tmp_details_1_row_0_cell",
      section: "details",
      type: "cell",
      item: sample.details[1].row.columns[0].cell,
      index: 0,
      parent: sample.details[1].row.columns,
      parent_type: "row",
      parent_index: 1,
      form: templateElements({ getText: getText })["cell"],
    },
  } 
}

export const VGap = Template.bind({});
VGap.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    current: {
      id: "tmp_details_0_vgap",
      section: "details",
      type: "vgap",
      item: sample.details[0].vgap,
      index: 0,
      parent: sample.details,
      parent_type: "details",
      parent_index: null,
      form: templateElements({ getText: getText })["vgap"],
    },
  } 
}

export const Data = Template.bind({});
Data.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "data",
  } 
}

export const StringData = Template.bind({});
StringData.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "data",
    current_data: {
      name: 'html_text',
      type: 'string'
    }
  } 
}

export const ListData = Template.bind({});
ListData.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "data",
    current_data: {
      name: 'labels',
      type: 'list',
      items: Object.keys(sample.data.labels).map(key => { return { lslabel: key, lsvalue: sample.data.labels[key] } })
    }
  } 
}

export const ListDataItem = Template.bind({});
ListDataItem.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "data",
    current_data: {
      ...ListData.args.data.current_data,
      item: "title"
    }
  } 
}

export const TableData = Template.bind({});
TableData.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "data",
    current_data: {
      name: 'items',
      type: 'table',
      items: sample.data.items.map((item, index) => { return { ...item, _index: index } }),
      fields: (()=>{ 
        let fields = {}
        Object.keys(sample.data.items[0]).forEach((key) => {
          fields[key] = { fieldtype: 'string', label: key }
        }) 
        return fields 
      })()
    }
  } 
}

export const TableDataItem = Template.bind({});
TableDataItem.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "data",
    current_data: {
      ...TableData.args.data.current_data,
      item: { ...sample.data.items[0], _index: 0 }
    }
  } 
}

export const Meta = Template.bind({});
Meta.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    tabView: "meta",
  } 
}