import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-datetime.js';

import { DATETIME_TYPE, APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/DateTime',
  component: 'form-datetime',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    type: {
      control: 'select',
      options: Object.values(DATETIME_TYPE),
    },
    onChange: {
      name: "onChange",
      description: "onChange click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onChange" 
    },
    onEnter: {
      name: "onEnter",
      description: "onEnter click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEnter" 
    }
  }
};

export function Template({ id, name, theme, value, type, isnull, picker,
  disabled, readonly, autofocus, full, style, 
  onChange, onEnter }) {
  const component = html`<form-datetime
    id="${id}"
    name="${name}"
    .value="${value}"
    type="${type}"
    ?disabled="${disabled}"
    ?readonly="${readonly}"
    ?autofocus="${autofocus}"
    ?full="${full}"
    .isnull="${isnull}"
    .picker="${picker}"
    label="test" 
    .style="${style}"
    .onChange=${onChange}
    .onEnter=${onEnter}
  ></form-datetime>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_date",
  name: undefined,
  theme: APP_THEME.LIGHT,
  value: "",
  type: DATETIME_TYPE.DATE,
  label: undefined,
  isnull: true,
  picker: false,
  disabled: false,
  readonly: false,
  autofocus: false,
  full: false,
  style: {}
}

export const DateTime = Template.bind({});
DateTime.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  isnull: false,
  picker: true,
  type: DATETIME_TYPE.DATETIME,
  full: true
}

export const Time = Template.bind({});
Time.args = {
  ...Default.args,
  type: DATETIME_TYPE.TIME,
}