import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-select.js';

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Select',
  component: 'form-select',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
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

export function Template({ id, name, theme, value, options, autofocus, disabled, isnull, full, style, 
  onChange, onEnter }) {
  const component = html`<form-select
    id="${id}"
    name="${name}"
    .value="${value}"
    .options="${options}"
    ?disabled="${disabled}"
    ?autofocus="${autofocus}"
    ?full="${full}"
    .isnull="${isnull}"
    label="test" 
    .style="${style}"
    .onChange=${onChange}
    .onEnter=${onEnter}
  ></form-select>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_select",
  theme: APP_THEME.LIGHT,
  value: "value1",
  options: [
    { value: "value1", text: "Text 1" },
    { value: "value2", text: "Text 2" },
    { value: "value3", text: "Text 3" }
  ],
  name: undefined,
  isnull: true,
  disabled: false,
  autofocus: false,
  full: false,
  style: {}
};

export const NotNull = Template.bind({});
NotNull.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  isnull: false,
  full: true
}
