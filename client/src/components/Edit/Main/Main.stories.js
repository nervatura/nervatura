import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './edit-main.js';

import { Default as EditorDefault, Report as EditorReport } from '../Editor/Editor.stories.js'
import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Edit/Main',
  component: 'edit-main',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onEditEvent: {
      name: "onEditEvent",
      description: "onEditEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEditEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, current, template, dataset, audit, theme, onEditEvent
}) {
  const component = html`<edit-main
    id="${id}"
    audit="${audit}"
    .current="${current}"
    .template="${template}"
    .dataset="${dataset}"
    .onEvent=${{ onEditEvent }}
    .msg=${msg}
  ></edit-main>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "main",
  theme: APP_THEME.LIGHT,
  current: EditorDefault.args.current, 
  template: EditorDefault.args.template, 
  dataset: EditorDefault.args.dataset, 
  audit: "all",
}

export const Report = Template.bind({});
Report.args = {
  id: "main",
  theme: APP_THEME.LIGHT,
  current: EditorReport.args.current, 
  template: EditorReport.args.template, 
  dataset: EditorReport.args.dataset, 
  audit: "all",
}