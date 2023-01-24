import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './sidebar-search.js';

import { SIDE_VISIBILITY, APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Sidebar/Search',
  component: 'sidebar-search',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    side: {
      control: 'select',
      options: Object.values(SIDE_VISIBILITY),
    },
    onSideEvent: {
      name: "onSideEvent",
      description: "onSideEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onSideEvent" 
    },
  }
};

export function Template({ id, side, groupKey, auditFilter, theme, onSideEvent, msg }) {
  const component = html`<sidebar-search
    id="${id}"
    side="${side}"
    groupKey="${groupKey}"
    .auditFilter="${auditFilter}"
    .onEvent=${{ onSideEvent }}
    .msg=${msg}
  ></sidebar-search>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "side_bar",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  groupKey: "transitem", 
  auditFilter: {
    trans: {
      offer: ["all", 1],
      order: ["all", 1],
      worksheet: ["all", 1],
      rent: ["all", 1],
      invoice: ["all", 1],
      receipt: ["all", 1],
      bank: ["all", 1],
      cash: ["all", 1],
      delivery: ["all", 1],
      inventory: ["all", 1],
      waybill: ["all", 1],
      production: ["all", 1],
      formula: ["all", 1]
    },
    menu: {},
    report: {},
    customer: ["all", 1],
    product: ["all", 1],
    employee: ["all", 1],
    tool: ["all", 1],
    project: ["all", 1],
    setting: ["all", 1],
    audit: ["all", 1]
  },
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const Office = Template.bind({});
Office.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  groupKey: "office",
  auditFilter: {
    trans: {
      offer: ["all", 1],
      order: ["all", 1],
      worksheet: ["all", 1],
      rent: ["all", 1],
      invoice: ["all", 1],
      receipt: ["all", 1],
      bank: ["disabled", 1],
      cash: ["disabled", 1],
      delivery: ["disabled", 1],
      inventory: ["disabled", 1],
      waybill: ["disabled", 1],
      production: ["disabled", 1],
      formula: ["disabled", 1]
    },
    menu: {},
    report: {},
    customer: ["disabled", 1],
    product: ["disabled", 1],
    employee: ["disabled", 1],
    tool: ["disabled", 1],
    project: ["disabled", 1],
    setting: ["all", 1],
    audit: ["all", 1]
  }
}