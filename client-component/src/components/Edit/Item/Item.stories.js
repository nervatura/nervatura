import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './edit-item.js';

import { Item as EditorItem  } from '../Editor/Editor.stories.js'
import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

import { Forms } from '../../../controllers/Forms.js'

export default {
  title: 'Edit/Item',
  component: 'edit-item',
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
  id, current, dataset, audit, theme, onEditEvent
}) {
  const component = html`<edit-item
    id="${id}"
    audit="${audit}"
    .current="${current}"
    .dataset="${dataset}"
    .onEvent=${{ onEditEvent }}
    .msg=${msg}
  ></edit-item>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "item",
  theme: APP_THEME.LIGHT,
  current: EditorItem.args.current,  
  dataset: EditorItem.args.dataset, 
  audit: "all",
}

export const NewItem = Template.bind({});
NewItem.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  current: {
    ...Default.args.current,
    form: {
      ...Default.args.current.form,
      id: null,
    },
    form_template: Forms({ msg: (key)=> key } ).item({ id: null }, { current: { transtype: "invoice" } })
  },
}