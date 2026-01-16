import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-label.js';

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Label',
  component: 'form-label',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    }
  }
};

export function Template({ id, theme, value, centered, leftIcon, rightIcon, style, iconStyle }) {
  const component = html`<form-label
    id="${id}"
    value="${value}"
    .style="${style}"
    ?centered="${centered}"
    .leftIcon="${leftIcon}"
    .rightIcon="${rightIcon}"
    .iconStyle="${iconStyle}"
  ></form-label>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_label",
  theme: APP_THEME.LIGHT,
  value: "Label",
  style: {},
  iconStyle: {}
};

export const LeftIcon = Template.bind({});
LeftIcon.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  leftIcon: "InfoCircle",
}

export const RightIcon = Template.bind({});
RightIcon.args = {
  ...Default.args,
  rightIcon: "InfoCircle",
}

export const Centered = Template.bind({});
Centered.args = {
  ...Default.args,
  leftIcon: "InfoCircle",
  centered: true
}


export const LabelStyle = Template.bind({});
LabelStyle.args = {
  ...Default.args,
  leftIcon: "InfoCircle",
  style: {color: "red"},
  iconStyle: {fill: "orange"}
}