import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-spinner.js';

export default {
  title: 'Form/NtSpinner',
  component: 'nt-spinner',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
    }
  }
};

export function Template({ theme }) {
  return html`<story-container theme="${theme}">
      <nt-spinner></nt-spinner>
    </story-container>
  `;
}

export const Default = Template.bind({});
Default.args = {
  theme: "light"
};

export const Dark = Template.bind({});
Dark.args = {
  theme: "dark"
};