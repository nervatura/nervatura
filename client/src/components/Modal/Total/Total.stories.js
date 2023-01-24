import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-total.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Modal/Total',
  component: 'modal-total',
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

export function Template({ id, total, theme, onModalEvent, msg }) {
  const component = html`<modal-total
    id="${id}"
    .total=${total}
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-total>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "total",
  theme: APP_THEME.LIGHT,
  total: {
    totalFields: {
      netamount: 11221,
      vatamount: 2097.2,
      amount: 13318.2,
      acrate: 0,
    },
    totalLabels: {
      netamount: "Net Amount",
      vatamount: "VAT",
      amount: "Amount",
      acrate: "Acc.Rate",
    },
    count: 4,
  },
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const DarkTotal = Template.bind({});
DarkTotal.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
}
