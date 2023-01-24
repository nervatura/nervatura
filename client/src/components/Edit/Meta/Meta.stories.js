import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './edit-meta.js';

import { Default as EditorDefault } from '../Editor/Editor.stories.js'
import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Edit/Meta',
  component: 'edit-meta',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onEditEvent: {
      name: "onEditEvent",
      description: "onEditEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEditEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, current, dataset, audit, pageSize, theme, onEditEvent
}) {
  const component = html`<edit-meta
    id="${id}"
    audit="${audit}"
    .current="${current}"
    .dataset="${dataset}"
    pageSize="${pageSize}"
    .onEvent=${{ onEditEvent }}
    .msg=${msg}
  ></edit-meta>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "meta",
  theme: APP_THEME.LIGHT,
  current: {
    ...EditorDefault.args.current,
    view: "fieldvalue",
    page: 0
  }, 
  dataset: EditorDefault.args.dataset, 
  audit: "all",
  pageSize: 5
}

export const NewField = Template.bind({});
NewField.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    deffield: "trans_transitem_link",
    page: 5,
  },
  dataset: {
    ...Default.args.dataset,
    deffield_prop: []
  },
  pageSize: 2
}

export const ReadOnly = Template.bind({});
ReadOnly.args = {
  ...Default.args,
  audit: "readonly"
}

export const Customer = Template.bind({});
Customer.args = {
  ...Default.args,
  current: {
    type: "customer",
    transtype: "",
    item: { id: 2, custtype: 116, custnumber: "DMCUST/00001", custname: "First Customer Co.",
      taxnumber: "12345678-1-12", account: null, notax: 0, terms: 8, creditlimit: 1000000,
      discount: 2, notes: null, inactive: 0, deleted: 0 },
    state: "normal",
    page: -1,
    fieldvalue: [
      { deleted: 0, fieldname: "sample_customer_date", id: 27, notes: "", 
        ref_id: 2, value: "2022-08-12" },
      {
        deleted: 0, fieldname: "sample_customer_float", id: 28, notes: "",
        ref_id: 2, value: "123.4" }
    ],
    view: "form"
  },
}