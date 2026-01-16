import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './client-template.js';

import { APP_THEME, SIDE_VISIBILITY } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

import { Default as EditorDefault } from '../Editor/Editor.stories.js'

export default {
  title: 'Template/Template',
  component: 'client-template',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onSideEvent: {
      name: "onSideEvent",
      description: "onSideEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onSideEvent" 
    },
    onTemplateEvent: {
      name: "onTemplateEvent",
      description: "onTemplateEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onTemplateEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, data, side, paginationPage, theme, onSideEvent, onTemplateEvent
}) {
  const component = html`<client-template
    id="${id}"
    .data=${data}
    side="${side}"
    paginationPage=${paginationPage}
    .msg=${msg}
    .onEvent=${{ 
      onSideEvent, onTemplateEvent, setModule: ()=>{}
    }}
  ></client-template>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "template",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  data: {
    ...EditorDefault.args.data,
    key: "template",
    dirty: false,
  },
  paginationPage: 10
}