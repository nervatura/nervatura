import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './client-edit.js';

import { APP_THEME, SIDE_VISIBILITY } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

import { Document as SideDocument } from '../../SideBar/Edit/Edit.stories.js'
import { Default as EditorDefault } from '../Editor/Editor.stories.js'

export default {
  title: 'Edit/Edit',
  component: 'client-edit',
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
  id, data, side, auditFilter, newFilter, forms, paginationPage, selectorPage, theme, 
  onSideEvent, onEditEvent
}) {
  const component = html`<client-edit
    id="${id}"
    .data=${data}
    side="${side}"
    .auditFilter="${auditFilter}"
    .newFilter="${newFilter}"
    .forms="${forms}"
    paginationPage=${paginationPage}
    selectorPage=${selectorPage}
    .msg=${msg}
    .onEvent=${{ 
      onSideEvent, onEditEvent, setModule: ()=>{}
    }}
  ></client-edit>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "edit",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  data: {
    side_view: SideDocument.args.view,
    caption: EditorDefault.args.caption,
    current: EditorDefault.args.current,
    template: EditorDefault.args.template,
    dataset: EditorDefault.args.dataset,
    audit: EditorDefault.args.audit,
    panel: SideDocument.args.module.panel
  },
  auditFilter: SideDocument.args.auditFilter,
  newFilter: SideDocument.args.newFilter,
  forms: SideDocument.args.forms,
  paginationPage: 10,
  selectorPage: 5
}

export const New = Template.bind({});
New.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    current: {
      ...Default.args.data.current,
      item: null
    }
  }
}