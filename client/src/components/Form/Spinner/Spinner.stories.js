import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-spinner.js';

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Spinner',
  component: 'form-spinner',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    }
  }
};

export function Template({ theme }) {
  return html`<story-container theme="${theme}">
      <form-spinner></form-spinner>
    </story-container>
  `;
}

export const Default = Template.bind({});
Default.args = {
  theme: APP_THEME.LIGHT
};

export const Dark = Template.bind({});
Dark.args = {
  theme: APP_THEME.DARK
};