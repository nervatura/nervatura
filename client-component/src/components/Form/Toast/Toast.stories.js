import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import '../Button/form-button.js'
import './form-toast.js';

import { TOAST_TYPE, APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Toast',
  component: 'form-toast',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    type: {
      control: 'select',
      options: Object.values(TOAST_TYPE),
    }
  }
};

export function Template({ id, name, theme, type, style, value, timeout, setData }) {
  const component = html`<form-toast
    id="${id}"
    name="${name}"
    .style="${style}"
    .setData="${setData}"
  ></form-toast>`
  return html`<story-container theme="${theme}">
  ${component}
  <form-button id="toast_show" .onClick=${()=>{
    document.querySelector(`#${id}`).show({type, value, timeout})
  }}>Show</form-button>
  </story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_toast",
  name: undefined,
  theme: APP_THEME.LIGHT,
  type: TOAST_TYPE.INFO,
  value: "This is an info message.",
  style: {}
}

export const Error = Template.bind({});
Error.args = {
  ...Default.args,
  type: TOAST_TYPE.ERROR,
  value: html`<i>This is an error message.</i>`,
  timeout: 0,
  setData: ()=>{}
}

export const Success = Template.bind({});
Success.args = {
  ...Default.args,
  type: TOAST_TYPE.SUCCESS,
  value: "This is an success message. This is an success message. This is an success message. This is an success message"
}