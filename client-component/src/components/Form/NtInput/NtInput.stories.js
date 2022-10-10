import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-input.js';

import { INPUT_TYPE } from './NtInput.js'

export default {
  title: 'Form/NtInput',
  component: 'nt-input',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    },
    type: {
      control: 'select',
      options: Object.values(INPUT_TYPE),
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

export function Template({ id, name, theme, value, type, placeholder, accept,
  disabled, readonly, autofocus, full, style, 
  onChange, onEnter }) {
  const component = html`<nt-input
    id="${id}"
    name="${name}"
    .value="${value}"
    type="${type}"
    ?disabled="${disabled}"
    ?readonly="${readonly}"
    ?autofocus="${autofocus}"
    placeholder="${placeholder}"
    accept="${accept}"
    ?full="${full}"
    label="test" 
    .style="${style}"
    .onChange=${onChange}
    .onEnter=${onEnter}
  ></nt-input>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_input",
  name: undefined,
  theme: "light",
  value: "",
  type: INPUT_TYPE.TEXT,
  label: undefined,
  placeholder: "placeholder text",
  disabled: false,
  readonly: false,
  autofocus: false,
  full: false,
  style: {}
}

export const Password = Template.bind({});
Password.args = {
  ...Default.args,
  type: INPUT_TYPE.PASSWORD,
  theme: "dark",
  value: "secret",
  full: true
}

export const File = Template.bind({});
File.args = {
  ...Default.args,
  type: INPUT_TYPE.FILE,
  accept: ".jpg,.png"
}

export const Color = Template.bind({});
Color.args = {
  ...Default.args,
  type: INPUT_TYPE.COLOR,
  value: "#845185"
}