import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-shipping.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Shipping',
  component: 'modal-shipping',
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
  id, partnumber, description, unit, batch_no, qty, theme, onModalEvent 
}) {
  const component = html`<modal-shipping
    id="${id}"
    unit="${unit}"
    batch_no="${batch_no}"
    qty="${qty}"
    partnumber="${partnumber}"
    description="${description}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-shipping>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "shipping",
  theme: APP_THEME.LIGHT,
  partnumber: "DMPROD/00008",
  description: "DMPROD/00008 | Pallet",
  unit: "piece",
  batch_no: "",
  qty: 6,
}

export const DarkShipping = Template.bind({});
DarkShipping.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
}