import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-number.js';

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/NumberInput',
  component: 'form-number',
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
    },
    onBlur: {
      name: "onBlur",
      description: "onBlur click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onBlur" 
    }
  }
};

export function Template({ id, name, theme, value, integer, min, max,
  disabled, readonly, autofocus, full, style, 
  onChange, onEnter, onBlur }) {
  const component = html`<form-number
    id="${id}"
    name="${name}"
    .value="${value}"
    .integer="${integer}"
    ?disabled="${disabled}"
    ?readonly="${readonly}"
    ?autofocus="${autofocus}"
    ?full="${full}"
    .min="${min}"
    .max="${max}"
    label="test" 
    .style="${style}"
    .onChange=${onChange}
    .onEnter=${onEnter}
    .onBlur=${onBlur}
  ></form-number>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_number",
  name: undefined,
  theme: APP_THEME.LIGHT,
  value: 0,
  integer: false,
  label: undefined,
  max: undefined,
  min: undefined,
  disabled: false,
  readonly: false,
  autofocus: false,
  full: false,
  style: {}
}

export const Integer = Template.bind({});
Integer.args = {
  ...Default.args,
  integer: true,
  min: 0,
  max: 100,
  full: true
}