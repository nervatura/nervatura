import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './template-editor.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

import { templateElements } from './Elements.js'
import { sample } from '../../../config/sample.js';

export default {
  title: 'Template/Editor',
  component: 'template-editor',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onTemplateEvent: {
      name: "onTemplateEvent",
      description: "onTemplateEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onTemplateEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, data, paginationPage, theme, onTemplateEvent
}) {
  const component = html`<template-editor
    id="${id}"
    paginationPage="${paginationPage}"
    .data="${data}"
    .onEvent=${{ onTemplateEvent }}
    .msg=${msg}
  ></template-editor>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "editor",
  theme: APP_THEME.LIGHT,
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
      form: templateElements({ msg }).report
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
  paginationPage: 10
}

export const Details = Template.bind({});
Details.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
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
      form: templateElements({ msg }).details
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
      form: templateElements({ msg }).row,
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
      form: templateElements({ msg }).cell,
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
      form: templateElements({ msg }).vgap,
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
      items: Object.keys(sample.data.labels).map(key => ({ lslabel: key, lsvalue: sample.data.labels[key] }))
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
      items: sample.data.items.map((item, index) => ({ ...item, _index: index })),
      fields: (()=>{ 
        const fields = {}
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
