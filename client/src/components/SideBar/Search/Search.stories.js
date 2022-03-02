import { Search, SIDE_VISIBILITY } from "./Search";

import { getText, store } from 'config/app';

export default {
  title: "SideBar/Search",
  component: Search,
}

const Template = (args) => <Search {...args} />

export const Default = Template.bind({});
Default.args = {
  side: SIDE_VISIBILITY.AUTO,
  groupKey: "transitem", 
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
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const Office = Template.bind({});
Office.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  groupKey: "office",
  auditFilter: {
    trans: {
      offer: ["all", 1],
      order: ["all", 1],
      worksheet: ["all", 1],
      rent: ["all", 1],
      invoice: ["all", 1],
      receipt: ["all", 1],
      bank: ["disabled", 1],
      cash: ["disabled", 1],
      delivery: ["disabled", 1],
      inventory: ["disabled", 1],
      waybill: ["disabled", 1],
      production: ["disabled", 1],
      formula: ["disabled", 1]
    },
    menu: {},
    report: {},
    customer: ["disabled", 1],
    product: ["disabled", 1],
    employee: ["disabled", 1],
    tool: ["disabled", 1],
    project: ["disabled", 1],
    setting: ["all", 1],
    audit: ["all", 1]
  }
}