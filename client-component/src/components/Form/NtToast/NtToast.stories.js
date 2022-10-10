import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import '../NtButton/nt-button.js'
import './nt-toast.js';

import { TOAST_TYPE } from './NtToast.js'

export default {
  title: 'Form/NtToast',
  component: 'nt-toast',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    },
    type: {
      control: 'select',
      options: Object.values(TOAST_TYPE),
    }
  }
};

export function Template({ id, name, theme, type, style, value, store }) {
  const component = html`<nt-toast
    id="${id}"
    name="${name}"
    type="${type}"
    .style="${style}"
    .store="${store}"
  >${value}</nt-toast>`
  return html`<story-container theme="${theme}">
  ${component}
  <nt-button id="toast_show" .onClick=${()=>{
    document.querySelector(`#${id}`).show()
  }}>Show</nt-button>
  </story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_toast",
  name: undefined,
  theme: "light",
  type: TOAST_TYPE.INFO,
  value: "This is an info message.",
  style: {}
}

export const Error = Template.bind({});
Error.args = {
  ...Default.args,
  type: TOAST_TYPE.ERROR,
  value: html`<i>This is an error message.</i>`,
  store: {
    data: {},
    setData: ()=>{}
  }
}

export const Success = Template.bind({});
Success.args = {
  ...Default.args,
  type: TOAST_TYPE.SUCCESS,
  value: "This is an success message. This is an success message. This is an success message. This is an success message"
}

export const Invalid = Template.bind({});
Invalid.args = {
  ...Default.args,
  type: "invalid",
}