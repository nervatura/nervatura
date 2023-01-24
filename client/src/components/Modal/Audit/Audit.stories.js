import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-audit.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Audit',
  component: 'modal-audit',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onModalEvent: {
      name: "onModalEvent",
      description: "onModalEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onModalEvent" 
    },
  }
};

export function Template({ 
  id, idKey, usergroup, nervatype, subtype, inputfilter, supervisor,
  typeOptions, subtypeOptions, inputfilterOptions, theme, onModalEvent 
}) {
  const component = html`<modal-audit
    id="${id}"
    idKey="${idKey}"
    usergroup="${usergroup}"
    nervatype="${nervatype}"
    subtype="${subtype}"
    inputfilter="${inputfilter}"
    supervisor="${supervisor}"
    .typeOptions="${typeOptions}"
    .subtypeOptions="${subtypeOptions}"
    .inputfilterOptions="${inputfilterOptions}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-audit>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "audit",
  theme: APP_THEME.LIGHT,
  idKey: null,
  usergroup: 1,
  nervatype: 6,
  typeOptions: [
    { value: "6", text: "audit" },
    { value: "10", text: "customer" },
    { value: "12", text: "employee" },
    { value: "13", text: "event" },
    { value: "18", text: "menu" },
    { value: "24", text: "price" },
    { value: "25", text: "product" },
    { value: "26", text: "project" },
    { value: "28", text: "report" },
    { value: "32", text: "setting" },
    { value: "30", text: "tool" },
    { value: "31", text: "trans" },
  ],
  subtype: null,
  subtypeOptions: [
    { value: "66", text: "bank", type: "trans" },
    { value: "67", text: "cash", type: "trans" },
    { value: "61", text: "delivery", type: "trans" },
    { value: "65", text: "formula", type: "trans" },
    { value: "62", text: "inventory", type: "trans" },
    { value: "55", text: "invoice", type: "trans" },
    { value: "58", text: "offer", type: "trans" },
    { value: "57", text: "order", type: "trans" },
    { value: "64", text: "production", type: "trans" },
    { value: "60", text: "rent", type: "trans" },
    { value: "63", text: "waybill", type: "trans" },
    { value: "59", text: "worksheet", type: "trans" },
    { value: "3", text: "ntr_bank_en", type: "report" },
    { value: "4", text: "ntr_cash_in_en", type: "report" },
    { value: "5", text: "ntr_cash_out_en", type: "report" },
  ],
  inputfilter: 109,
  inputfilterOptions: [
    { value: "109", text: "all" },
    { value: "106", text: "disabled" },
    { value: "107", text: "readonly" },
    { value: "108", text: "update" }
  ],
  supervisor: 1,
}

export const Existing = Template.bind({});
Existing.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  idKey: 1,
  nervatype: 31,
  subtype: 61,
  inputfilter: 107,
  supervisor: 0,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  nervatype: 18,
}