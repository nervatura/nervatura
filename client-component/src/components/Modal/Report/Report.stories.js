import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-report.js';

import { APP_THEME } from '../../../config/enums.js'
import { store } from '../../../config/app.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Print',
  component: 'modal-report',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    size: {
      control: 'select',
      options: store.ui.report_size.map(item => item[0]),
    },
    orient: {
      control: 'select',
      options: store.ui.report_orientation.map(item => item[0]),
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
  id, title, template, templates, orient, size, copy, theme, onModalEvent 
}) {
  const component = html`<modal-report
    id="${id}"
    title="${title}"
    template="${template}"
    copy="${copy}"
    orient="${orient}"
    size="${size}"
    .templates=${templates}
    .report_size=${store.ui.report_size}
    .report_orientation=${store.ui.report_orientation}
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-report>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "report",
  theme: APP_THEME.LIGHT,
  title: "DMINV/00001",
  template: "ntr_invoice_en",
  templates: [
    { text: "Invoice EN", value: "ntr_invoice_en" },
    { text: "Lasku FI", value: "ntr_invoice_fi" }
  ],
  orient: "portrait",
  size: "a4", 
  copy: 1,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  template: "",
  templates: []
}