import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './setting-view.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Setting/View',
  component: 'setting-view',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onSettingEvent: {
      name: "onSettingEvent",
      description: "onSettingEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onSettingEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, data, paginationPage, theme, onSettingEvent
}) {
  const component = html`<setting-view
    id="${id}"
    .data="${data}"
    paginationPage="${paginationPage}"
    .onEvent=${{ onSettingEvent }}
    .msg=${msg}
  ></setting-view>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "setting_view",
  theme: APP_THEME.LIGHT,
  data: {
    caption: "DEFAULT SETTINGS",
    icon: "Cog",
    view: {
      type: "list",
      result: [
        { description: "business year", fieldname: "transyear", fieldtype: "integer",
          id: 13, lslabel: "business year", lsvalue: "2022", notes: null, value: "2022", valuelist: null },
        { description: "default bank place no.", fieldname: "default_bank", fieldtype: "string",
          id: 1, lslabel: "default bank place no.", lsvalue: "bank", notes: null, value: "bank", valuelist: null },
        { description: "default country", fieldname: "default_country", fieldtype: "string",
          id: 4, lslabel: "default country", lsvalue: "EU", notes: null, value: "EU", valuelist: null },
        { description: "default currency", fieldname: "default_currency", fieldtype: "string",
          id: 6, lslabel: "default currency", lsvalue: "EUR", notes: null, value: "EUR", valuelist: null },
      ],
      fields: null
    },
    actions: {
      new: { action: "newItem" },
      edit: { action: "editItem" },
      delete: { action: "deleteItem" },
    },
    page: 0
  },
  paginationPage: 3
}

export const Table = Template.bind({});
Table.args = {
  ...Default.args,
  data: {
    caption: "CURRENCY",
    icon: "Dollar",
    view: {
      type: "table",
      result: [
        { cround: 0, curr: "EUR", defrate: 0, description: "euro", digit: 2, id: 1 },
        { cround: 0, curr: "USD", defrate: 0, description: "dollar", digit: 2, id: 2 }
      ],
      fields: {
        curr: { fieldtype: "string", label: "Currency" },
        description: { fieldtype: "string", label: "Description" },
        digit: { fieldtype: "number", label: "Digit" },
        cround: { fieldtype: "number", label: "Def.Rate" },
        defrate: { fieldtype: "number", label: "Round" }
      }
    },
    actions: {
      new: { action: "newItem" },
      edit: { action: "editItem" },
      delete: { action: "deleteItem" }
    },
    page: 0
  }
}

export const ReadOnlyList = Template.bind({});
ReadOnlyList.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    actions: {}
  }
}

export const ReadOnlyTable = Template.bind({});
ReadOnlyTable.args = {
  ...Table.args,
  data: {
    ...Table.args.data,
    actions: {}
  },
  paginationPage: 1
}