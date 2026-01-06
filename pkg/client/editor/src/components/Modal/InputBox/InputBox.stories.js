import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-inputbox.js';

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Modal/InputBox',
  component: 'modal-inputbox',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
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
  id, title, message, infoText, value,
  labelCancel, labelOK, defaultOK, showValue, theme, onModalEvent 
}) {
  const component = html`<modal-inputbox
    id="${id}"
    title="${title}"
    message="${message}"
    infoText="${infoText}"
    value="${value}"
    labelCancel="${labelCancel}"
    labelOK="${labelOK}"
    ?defaultOK="${defaultOK}"
    ?showValue="${showValue}"
    .onEvent=${{ onModalEvent }}
  ></modal-inputbox>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "inputbox",
  theme: APP_THEME.LIGHT,
  title: "Input value title",
  message: "Input value message",
  infoText: undefined,
  value: "",
  labelCancel: "Cancel",
  labelOK: "OK",
  defaultOK: true,
  showValue: false,
}

export const InputValue = Template.bind({});
InputValue.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  infoText: "Input value info text",
  value: "default value",
  showValue: true,
}

export const DefaultCancel = Template.bind({});
DefaultCancel.args = {
  ...Default.args,
  defaultOK: false,
}