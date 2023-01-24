import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './edit-view.js';

import { Default as EditorDefault } from '../Editor/Editor.stories.js'
import { APP_THEME, ACTION_EVENT } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Edit/View',
  component: 'edit-view',
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
  id, viewName, current, dataset, audit, template, pageSize, theme, onEditEvent
}) {
  const component = html`<edit-view
    id="${id}"
    viewName="${viewName}"
    audit="${audit}"
    .current="${current}"
    .dataset="${dataset}"
    .template="${template}"
    pageSize="${pageSize}"
    .onEvent=${{ onEditEvent }}
    .msg=${msg}
  ></edit-view>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "view",
  theme: APP_THEME.LIGHT,
  viewName: "item",
  current: EditorDefault.args.current, 
  template: EditorDefault.args.template, 
  dataset: EditorDefault.args.dataset, 
  audit: "all",
  pageSize: 10
}

export const List = Template.bind({});
List.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  viewName: "invoice_link"
}

export const ReadOnly = Template.bind({});
ReadOnly.args = {
  ...Default.args,
  audit: "readonly",
}

export const ReadOnlyList = Template.bind({});
ReadOnlyList.args = {
  ...Default.args,
  viewName: "invoice_link",
  template: {
    ...EditorDefault.args.template,
    view: {
      invoice_link: {
        ...EditorDefault.args.template.view.invoice_link,
        actions: {
          new: null,
          edit: null,
          delete: null
        }
      }
    }
  },
}

export const Empty = Template.bind({});
Empty.args = {
  ...Default.args,
  template: {
    ...EditorDefault.args.template,
    view: {
      item: {
        ...EditorDefault.args.template.view.item,
        data: "missing",
        edit_icon: "Edit",
        delete_icon: "Times",
        new_icon: "Plus",
        edited: false,
        new_label: "NEW"
      }
    }
  }
}

export const DeleteOnly = Template.bind({});
DeleteOnly.args = {
  ...Default.args,
  template: {
    ...EditorDefault.args.template,
    view: {
      item: {
        ...EditorDefault.args.template.view.item,
        actions: {
          new: null,
          edit: null,
          delete: {action: ACTION_EVENT.DELETE_EDITOR_ITEM, fkey: "item"}
        }
      }
    }
  }
}
