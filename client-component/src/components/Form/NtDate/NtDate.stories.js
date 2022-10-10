import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-date.js';

import { DATE_TYPE } from './NtDate.js'

export default {
  title: 'Form/NtDate',
  component: 'nt-date',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    },
    type: {
      control: 'select',
      options: Object.values(DATE_TYPE),
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
  const component = html`<nt-date
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
  ></nt-date>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_date",
  name: undefined,
  theme: "light",
  value: "",
  type: DATE_TYPE.DATE,
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
  theme: "dark",
  isnull: false,
  picker: true,
  type: DATE_TYPE.DATETIME,
  full: true
}

export const Time = Template.bind({});
Time.args = {
  ...Default.args,
  type: DATE_TYPE.TIME,
}