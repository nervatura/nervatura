import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-stock.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Stock',
  component: 'modal-stock',
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
  id, partnumber, partname, rows, selectorPage,
  theme, onModalEvent 
}) {
  const component = html`<modal-stock
    id="${id}"
    partnumber="${partnumber}"
    partname="${partname}"
    selectorPage="${selectorPage}"
    .rows="${rows}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-stock>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "stock",
  theme: APP_THEME.LIGHT,
  partnumber: "DMPROD/00001",
  partname: "Big product",
  rows: [
    {
      batch_no: "demo",
      description: "Big product",
      id: 4,
      partnumber: "DMPROD/00001",
      shipping: "2020-11-08T00:00:00",
      sqty: 5,
      unit: "piece",
      warehouse: "material | Raw material",
    },
    {
      batch_no: "demo",
      description: "Big product",
      id: 3,
      partnumber: "DMPROD/00001",
      shipping: "2020-12-10T00:00:00",
      sqty: 2,
      unit: "piece",
      warehouse: "warehouse | Warehouse"
    }
  ],
  selectorPage: 5
}

export const DarkShipping = Template.bind({});
DarkShipping.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
}