import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-trans.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Trans',
  component: 'modal-trans',
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
  id, baseTranstype, transtype, direction, doctypes, directions, 
  refno, nettoDiv, netto, fromDiv, from, elementCount,
  theme, onModalEvent 
}) {
  const component = html`<modal-trans
    id="${id}"
    baseTranstype="${baseTranstype}"
    transtype="${transtype}"
    direction="${direction}"
    .doctypes="${doctypes}"
    .directions="${directions}"
    ?refno="${refno}"
    ?nettoDiv="${nettoDiv}"
    ?netto="${netto}"
    ?fromDiv="${fromDiv}"
    ?from="${from}"
    elementCount="${elementCount}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-trans>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "trans",
  theme: APP_THEME.LIGHT,
  baseTranstype: "offer",
  transtype: "order",
  direction: "out",
  doctypes: ["offer","order","worksheet","rent"],
  directions: ["in","out"],
  refno: true, 
  nettoDiv: false, 
  netto: true, 
  fromDiv: false, 
  from: false,
  elementCount: 10,
}

export const Order = Template.bind({});
Order.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  baseTranstype: "order",
  transtype: "invoice",
  direction: "in",
  doctypes: ["offer","order","worksheet","rent","invoice","receipt"],
  refno: true, 
  nettoDiv: true, 
  netto: true, 
  fromDiv: true, 
  from: true,
  elementCount: 0,
}

export const Worksheet = Template.bind({});
Worksheet.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  baseTranstype: "worksheet",
  transtype: "invoice",
  direction: "in",
  doctypes: ["offer","order","worksheet","rent","invoice","receipt"],
  refno: true, 
  nettoDiv: false, 
  netto: true, 
  fromDiv: true, 
  from: false,
  elementCount: 10,
}