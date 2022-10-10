import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-button.js';
import '../NtLabel/nt-label.js';

import { BUTTON_TYPE } from './NtButton.js'

export default {
  title: 'Form/NtButton',
  component: 'nt-button',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    },
    type: {
      control: 'select',
      options: Object.values(BUTTON_TYPE),
    },
    onClick: {
      name: "onClick",
      description: "onClick click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onClick" 
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

export function Template({ id, name, theme, type, style, disabled, small, autofocus, full, value, 
  onClick, onEnter }) {
  const component = html`<nt-button
    id="${id}"
    name="${name}"
    .value="${value}"
    type="${type}"
    ?disabled="${disabled}"
    ?autofocus="${autofocus}"
    ?small="${small}"
    ?full="${full}"
    label="test" 
    .style="${style}"
    .onClick=${onClick}
    .onEnter=${onEnter}
  >${value}</nt-button>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_button",
  theme: "light",
  name: undefined,
  type: BUTTON_TYPE.PRIMARY,
  small: false,
  full: false,
  disabled: false,
  value: "Primary",
  style: {}
};

export const PrimaryDark = Template.bind({});
PrimaryDark.args = {
  ...Default.args,
  theme: "dark",
  type: BUTTON_TYPE.PRIMARY,
  value: "Primary dark",
  full: true
};

export const Secondary = Template.bind({});
Secondary.args = {
  ...Default.args,
  theme: "light",
  type: BUTTON_TYPE.SECONDARY,
  value: "Secondary",
};

export const SecondaryDark = Template.bind({});
SecondaryDark.args = {
  ...Default.args,
  theme: "dark",
  type: BUTTON_TYPE.SECONDARY,
  value: "Secondary dark",
};

export const Border = Template.bind({});
Border.args = {
  ...Default.args,
  theme: "light",
  type: BUTTON_TYPE.BORDER,
  value: "Border",
};

export const BorderDark = Template.bind({});
BorderDark.args = {
  ...Default.args,
  theme: "dark",
  type: BUTTON_TYPE.BORDER,
  value: "Border dark",
};

export const SmallDisabled = Template.bind({});
SmallDisabled.args = {
  ...Default.args,
  disabled: true,
  small: true
};

export const ButtonStyle = Template.bind({});
ButtonStyle.args = {
  ...Default.args,
  type: BUTTON_TYPE.BORDER,
  value: "Button style",
  style: {"border-color": "green", color: "red", "border-radius": "3px"}
};

export const LabelButton = Template.bind({});
LabelButton.args = {
  ...Default.args,
  value: html`<nt-label value="Label" leftIcon="InfoCircle" 
    .style="${{color: "#FFFFFF"}}" .iconStyle="${{color: "#FFFFFF"}}" ></nt-label>`,
};

export const Invalid = Template.bind({});
Invalid.args = {
  ...Default.args,
  type: "type",
};

