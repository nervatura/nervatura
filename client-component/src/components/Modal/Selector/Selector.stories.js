import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-selector.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Modal/Selector',
  component: 'modal-selector',
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

export function Template({ id, isModal, view, columns, result, filter, theme, onModalEvent, msg }) {
  const component = html`<modal-selector
    id="${id}"
    ?isModal="${isModal}"
    view="${view}"
    .columns=${columns}
    .result=${result}
    filter="${filter}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-selector>`
  return html`<story-container theme="${theme}" .style=${{ "padding": (!isModal) ? "20px 0" : "" }} >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "selector",
  theme: APP_THEME.LIGHT,
  isModal: true,
  view: "customer",
  columns: [
    ["custname", "c.custname"],
    ["custnumber", "c.custnumber"],
    ["city", "addr.city"],
    ["street", "addr.street"]
  ],
  result: [
    {
      city: 'City1',
      custname: 'First Customer Co.',
      custnumber: 'DMCUST/00001',
      id: 'customer-2',
      label: 'First Customer Co.',
      street: 'street 1.'
    },
    {
      city: 'City3',
      custname: 'Second Customer Name',
      custnumber: 'DMCUST/00002',
      id: 'customer-3',
      label: 'Second Customer Name',
      street: 'street 3.',
      deleted: 1
    },
    {
      city: 'City4',
      custname: 'Third Customer Foundation',
      custnumber: 'DMCUST/00003',
      id: 'customer-4',
      label: 'Third Customer Foundation',
      street: 'street 4.'
    }
  ],
  filter: "",
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const QuickView = Template.bind({});
QuickView.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  isModal: false
}