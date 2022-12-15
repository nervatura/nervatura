import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './sidebar-edit.js';

import { SIDE_VISIBILITY, SIDE_VIEW, APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Sidebar/Edit',
  component: 'sidebar-edit',
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
    view: {
      control: 'select',
      options: Object.values(SIDE_VIEW),
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

export function Template({ id, side, view, newFilter, auditFilter, module, forms, theme, onSideEvent, msg }) {
  const component = html`<sidebar-edit
    id="${id}"
    side="${side}"
    view="${view}"
    .newFilter="${newFilter}"
    .auditFilter="${auditFilter}"
    .module="${module}"
    .forms="${forms}"
    .onEvent=${{ onSideEvent }}
    .msg=${msg}
  ></sidebar-edit>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "side_bar",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  view: SIDE_VIEW.NEW,
  module: {
    current: {
      type: "invoice",
      item: {
        id: 1, nervatype: 1, ref_id: 1
      }
    }, 
    form_dirty: false,
    dirty: false,
    panel: {
      arrow: true,
      more: true,
      trans: true,
      bookmark: ["editor", "trans", "transnumber"],
      help: "document",
      corrective: true,
      form: true,
      state: "normal",
      formula: true,
      link: true,
      password: true,
      shipping: true,
      report: true,
      search: true,
      export_all: true,
      print: true,
      export_pdf: true,
      export_xml: true,
      export_csv: true,
      export_event: true,
      link_label: "LINK"
    }, 
    dataset: {
      shiptemp: [{}]
    }, 
    group_key: ""
  }, 
  newFilter: [
    ["offer","order","worksheet","rent","invoice","receipt"],
    ["bank","cash"],
    ["delivery","inventory","waybill","production","formula"],
    ["customer","product","employee","tool","project"]
  ],
  auditFilter: {
    trans: {
      offer: ["all", 1],
      order: ["all", 1],
      worksheet: ["all", 1],
      rent: ["all", 1],
      invoice: ["all", 1],
      receipt: ["disabled", 1],
      bank: ["all", 1],
      cash: ["disabled", 1],
      delivery: ["all", 1],
      inventory: ["all", 1],
      waybill: ["all", 1],
      production: ["disabled", 1],
      formula: ["all", 1]
    },
    menu: {},
    report: {},
    customer: ["all", 1],
    product: ["all", 1],
    employee: ["all", 1],
    tool: ["all", 1],
    project: ["disabled", 1],
    setting: ["all", 1],
    audit: ["all", 1]
  }, 
  forms: {
    delivery: ()=>({ options: { icon: "Truck" }}),
    inventory: ()=>({ options: { icon: "Truck" }}),
    waybill: ()=>({ options: { icon: "Briefcase" }}),
    // production: ()=>({ options: { icon: "Flask" }}),
    formula: ()=>({ options: { icon: "Magic" }}),
    customer: ()=>({ options: { icon: "User" }}),
    product: ()=>({ options: { icon: "ShoppingCart" }}),
    employee: ()=>({ options: { icon: "Male" }}),
    tool: ()=>({ options: { icon: "Wrench" }}),
    // project: ()=>({ options: { icon: "Clock" }})
  },
  msg: (defaultValue, props) => locales.en[props.id] || defaultValue
}

export const NewItem = Template.bind({});
NewItem.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  view: SIDE_VIEW.EDIT,
  module: {
    ...Default.args.module,
    current: {},
    group_key: "new_transitem"
  },
  newFilter: [
    ["offer","order","worksheet","rent","invoice","receipt"],
    [],
    [],
    []
  ],
}

export const NewPayment = Template.bind({});
NewPayment.args = {
  ...Default.args,
  module: {
    ...Default.args.module,
    current: {
      form: {}
    },
    group_key: "new_transpayment"
  },
  newFilter: [
    [],
    ["bank","cash"],
    ["delivery","inventory","waybill","production","formula"],
    ["customer","product","employee","tool","project"]
  ],
}

export const NewMovement = Template.bind({});
NewMovement.args = {
  ...Default.args,
  module: {
    ...Default.args.module,
    group_key: "new_transmovement"
  }
}

export const NewResource = Template.bind({});
NewResource.args = {
  ...Default.args,
  module: {
    ...Default.args.module,
    group_key: "new_resources"
  }
}

export const Document = Template.bind({});
Document.args = {
  ...Default.args,
  view: SIDE_VIEW.EDIT,
  module: {
    ...Default.args.module,
    form_dirty: true,
    dirty: true,
    dataset: {
      shiptemp: []
    }
  }
}

export const DocumentDeleted = Template.bind({});
DocumentDeleted.args = {
  ...Default.args,
  view: SIDE_VIEW.EDIT,
  module: {
    ...Default.args.module, 
    panel: {
      ...Default.args.module.panel,
      state: "deleted",
      cancellation: true,
      more: false,
    },
    current: {
      form: {}
    },
    form_dirty: true,
  }
}

export const DocumentCancellation = Template.bind({});
DocumentCancellation.args = {
  ...Default.args,
  view: SIDE_VIEW.EDIT,
  module: {
    ... Default.args.module, 
    panel: {
      ...Default.args.module.panel,
      state: "cancellation",
      back: true,
      arrow: false,
      help: false,
      print: false,
      report: false,
      search: false,
      formula: false,
      copy: false,
      create: false
    }
  }
}

export const DocumentClosed = Template.bind({});
DocumentClosed.args = {
  ...Default.args,
  view: SIDE_VIEW.EDIT,
  module: {
    ...Default.args.module, 
    panel: {
      ...Default.args.module.panel,
      state: "closed",
      save: false,
      delete: false,
      new: false
    }
  }
}

export const DocumentReadonly = Template.bind({});
DocumentReadonly.args = {
  ...Default.args,
  view: SIDE_VIEW.EDIT,
  module: {
    ...Default.args.module, 
    panel: {
      ...Default.args.module.panel,
      state: "readonly",
    }
  }
}

export const DocumentNoOptions = Template.bind({});
DocumentNoOptions.args = {
  ...Default.args,
  view: SIDE_VIEW.EDIT,
  module: {
    ...Default.args.module, 
    panel: undefined,
  }
}