import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './client-setting.js';

import { APP_THEME, SIDE_VISIBILITY } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

import { Default as ViewDefault } from '../View/View.stories.js'
import { Default as FormDefault } from '../Form/Form.stories.js'

export default {
  title: 'Setting/Setting',
  component: 'client-setting',
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
    onSettingEvent: {
      name: "onSettingEvent",
      description: "onSettingEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onSettingEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, data, side, auditFilter, username, paginationPage, theme, 
  onSideEvent, onSettingEvent
}) {
  const component = html`<client-setting
    id="${id}"
    .data=${data}
    side="${side}"
    .auditFilter="${auditFilter}"
    username="${username}"
    paginationPage=${paginationPage}
    .msg=${msg}
    .onEvent=${{ 
      onSideEvent, onSettingEvent, setModule: ()=>{}
    }}
  ></client-setting>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "setting",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  data: {
    ...ViewDefault.args.data,
    group_key: "group_admin",
    dirty: false,
  },
  auditFilter: {
    setting: ["all", 1],
    audit: ["all", 1]
  },
  username: "admin",
  paginationPage: 10
}

export const Form = Template.bind({});
Form.args = {
  ...Default.args,
  data: {
    ...FormDefault.args.data,
    group_key: "group_admin",
    dirty: false,
  },
}

export const Empty = Template.bind({});
Empty.args = {
  ...Default.args,
  data: {
    ...Default.args.data,
    view: null,
  },
}