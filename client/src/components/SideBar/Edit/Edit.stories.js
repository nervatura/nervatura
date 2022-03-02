import update from 'immutability-helper';
import { Edit, SIDE_VISIBILITY } from "./Edit";

import { getText, store } from 'config/app';

export default {
  title: "SideBar/Edit",
  component: Edit,
}

const Template = (args) => <Edit {...args} />

export const Default = Template.bind({});
Default.args = {
  side: SIDE_VISIBILITY.AUTO,
  edit: false, 
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
    delivery: ()=>{
      return { options: { icon: "Truck" }}
    },
    inventory: ()=>{
      return { options: { icon: "Truck" }}
    },
    waybill: ()=>{
      return { options: { icon: "Briefcase" }}
    },
    //production: ()=>{
    //  return { options: { icon: "Flask" }}
    //},
    formula: ()=>{
      return { options: { icon: "Magic" }}
    },
    customer: ()=>{
      return { options: { icon: "User" }}
    },
    product: ()=>{
      return { options: { icon: "ShoppingCart" }}
    },
    employee: ()=>{
      return { options: { icon: "Male" }}
    },
    tool: ()=>{
      return { options: { icon: "Wrench" }}
    },
    //project: ()=>{
    //  return { options: { icon: "Clock" }}
    //}
  },
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const NewItem = Template.bind({});
NewItem.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  edit: true,
  module: update(Default.args.module, {$merge: {
    current: {},
    group_key: "new_transitem"
  }}),
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
  module: update(Default.args.module, {$merge: {
    current: {
      form: {}
    },
    group_key: "new_transpayment"
  }}),
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
  module: update(Default.args.module, {$merge: {
    group_key: "new_transmovement"
  }})
}

export const NewResource = Template.bind({});
NewResource.args = {
  ...Default.args,
  module: update(Default.args.module, {$merge: {
    group_key: "new_resources"
  }})
}

export const Document = Template.bind({});
Document.args = {
  ...Default.args,
  edit: true,
  module: update(Default.args.module, {$merge: {
    form_dirty: true,
    dirty: true,
    dataset: {
      shiptemp: []
    }
  }})
}

export const DocumentDeleted = Template.bind({});
DocumentDeleted.args = {
  ...Default.args,
  edit: true,
  module: update(update(Default.args.module, {panel: {$merge: {
    state: "deleted",
    cancellation: true,
    more: false,
  }}}), {$merge: {
    current: {
      form: {}
    },
    form_dirty: true,
  }})
}

export const DocumentCancellation = Template.bind({});
DocumentCancellation.args = {
  ...Default.args,
  edit: true,
  module: update(Default.args.module, {panel: {$merge: {
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
  }}})
}

export const DocumentClosed = Template.bind({});
DocumentClosed.args = {
  ...Default.args,
  edit: true,
  module: update(Default.args.module, {panel: {$merge: {
    state: "closed",
    save: false,
    delete: false,
    new: false
  }}})
}

export const DocumentReadonly = Template.bind({});
DocumentReadonly.args = {
  ...Default.args,
  edit: true,
  module: update(Default.args.module, {panel: {$merge: {
    state: "readonly",
  }}})
}

export const DocumentNoOptions = Template.bind({});
DocumentNoOptions.args = {
  ...Default.args,
  edit: true,
  module: update(Default.args.module, {$merge: {
    panel: undefined,
    
  }})
}