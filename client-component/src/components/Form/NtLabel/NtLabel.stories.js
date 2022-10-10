import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-label.js';

export default {
  title: 'Form/NtLabel',
  component: 'nt-label',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    }
  }
};

export function Template({ id, theme, value, centered, leftIcon, rightIcon, style, iconStyle }) {
  const component = html`<nt-label
    id="${id}"
    value="${value}"
    .style="${style}"
    ?centered="${centered}"
    .leftIcon="${leftIcon}"
    .rightIcon="${rightIcon}"
    .iconStyle="${iconStyle}"
  ></nt-label>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_label",
  theme: "light",
  value: "Label",
  style: {},
  iconStyle: {}
};

export const LeftIcon = Template.bind({});
LeftIcon.args = {
  ...Default.args,
  theme: "dark",
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