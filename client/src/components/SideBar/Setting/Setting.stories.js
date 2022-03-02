import { Setting, SIDE_VISIBILITY } from "./Setting";

import { getText, store } from 'config/app';

export default {
  title: "SideBar/Setting",
  component: Setting,
}

const Template = (args) => <Setting {...args} />

export const Default = Template.bind({});
Default.args = {
  side: SIDE_VISIBILITY.AUTO,
  module: { 
    dirty: false, 
    group_key: "group_admin"
  }, 
  auditFilter: {
    trans: {
      offer: ["all", 1],
      order: ["all", 1],
      worksheet: ["all", 1],
      rent: ["all", 1],
      invoice: ["all", 1],
      receipt: ["all", 1],
      bank: ["all", 1],
      cash: ["all", 1],
      delivery: ["all", 1],
      inventory: ["all", 1],
      waybill: ["all", 1],
      production: ["all", 1],
      formula: ["all", 1]
    },
    menu: {},
    report: {},
    customer: ["all", 1],
    product: ["all", 1],
    employee: ["all", 1],
    tool: ["all", 1],
    project: ["all", 1],
    setting: ["all", 1],
    audit: ["all", 1]
  },
  username: "admin",
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
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