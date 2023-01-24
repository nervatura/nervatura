import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './sidebar-setting.js';

import { SIDE_VISIBILITY, APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Sidebar/Setting',
  component: 'sidebar-setting',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    side: {
      control: 'select',
      options: Object.values(SIDE_VISIBILITY),
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
  }
};

export function Template({ id, side, username, auditFilter, module, theme, onSideEvent, msg }) {
  const component = html`<sidebar-setting
    id="${id}"
    side="${side}"
    username="${username}"
    .auditFilter="${auditFilter}"
    .module="${module}"
    .onEvent=${{ onSideEvent }}
    .msg=${msg}
  ></sidebar-setting>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "side_bar",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  module: { 
    dirty: false, 
    group_key: "group_admin"
  }, 
  auditFilter: {
    setting: ["all", 1],
    audit: ["all", 1]
  },
  username: "admin",
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const AdminGroup = Template.bind({});
AdminGroup.args = {
  ...Default.args,
  auditFilter: {
    setting: ["all", 1],
    audit: ["disabled", 1]
  }
}

export const DatabaseGroup = Template.bind({});
DatabaseGroup.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  module: {  
    group_key: "group_database"
  },
}

export const UserGroup = Template.bind({});
UserGroup.args = {
  ...Default.args,
  module: {  
    group_key: "group_user"
  },
  auditFilter: {
    setting: ["disabled", 1],
  }
}

export const FormItemAll = Template.bind({});
FormItemAll.args = {
  ...Default.args,
  module: {
    dirty: false,  
    current: {
      form: {
        id: 1
      }
    },
    panel: {
      save: true,
      delete: true,
      new: true,
      help: "help"
    }
  },
}

export const FormItemNew = Template.bind({});
FormItemNew.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  module: {
    dirty: true,  
    current: {
      form: {
        id: null
      }
    },
    panel: {
      save: true,
      delete: true,
      new: true,
    }
  },
}

export const FormItemRead = Template.bind({});
FormItemRead.args = {
  ...Default.args,
  module: {
    dirty: false,  
    current: {
      form: {
        id: 1
      }
    },
    panel: {
      save: false
    }
  },
}
