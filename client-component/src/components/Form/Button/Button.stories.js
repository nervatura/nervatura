import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-button.js';
import '../Label/form-label.js';

import { BUTTON_TYPE, APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Button',
  component: 'form-button',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
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

export function Template({ 
  id, name, theme, 
  type, style, disabled, small, autofocus, full, selected, hidelabel, value, icon,
  onClick, onEnter }) {
  const component = html`<form-button
    id="${id}"
    name="${name}"
    .value="${value}"
    .icon="${icon}"
    type="${type}"
    ?disabled="${disabled}"
    ?autofocus="${autofocus}"
    ?small="${small}"
    ?full="${full}"
    ?selected="${selected}"
    ?hidelabel="${hidelabel}"
    label="test" 
    .style="${style}"
    .onClick=${onClick}
    .onEnter=${onEnter}
  >${value}</form-button>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_button",
  theme: APP_THEME.LIGHT,
  name: undefined,
  type: BUTTON_TYPE.DEFAULT,
  small: false,
  full: false,
  selected: false,
  hidelabel: false,
  disabled: false,
  value: "Default",
  style: {}
};

export const Dark = Template.bind({});
Dark.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  type: BUTTON_TYPE.DEFAULT,
  value: "Default dark",
  icon: "InfoCircle",
  hidelabel: true
};

export const Primary = Template.bind({});
Primary.args = {
  ...Default.args,
  theme: APP_THEME.LIGHT,
  type: BUTTON_TYPE.PRIMARY,
  value: "Primary",
  icon: "Check",
  selected: true
};

export const PrimaryDark = Template.bind({});
PrimaryDark.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  type: BUTTON_TYPE.PRIMARY,
  value: "Primary dark"
};

export const Border = Template.bind({});
Border.args = {
  ...Default.args,
  theme: APP_THEME.LIGHT,
  type: BUTTON_TYPE.BORDER,
  value: "Border",
  full: true
};

export const BorderDark = Template.bind({});
BorderDark.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
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
  value: html`<form-label value="Label" leftIcon="InfoCircle"></form-label>`,
};

