import { Audit } from "./Audit";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Audit",
  component: Audit
}

const Template = (args) => <Audit {...args} />

export const Default = Template.bind({});
Default.args = {
  idKey: null,
  usergroup: 1,
  nervatype: 6,
  typeOptions: [
    { value: "6", text: "audit" },
    { value: "10", text: "customer" },
    { value: "12", text: "employee" },
    { value: "13", text: "event" },
    { value: "18", text: "menu" },
    { value: "24", text: "price" },
    { value: "25", text: "product" },
    { value: "26", text: "project" },
    { value: "28", text: "report" },
    { value: "32", text: "setting" },
    { value: "30", text: "tool" },
    { value: "31", text: "trans" },
  ],
  subtype: null,
  subtypeOptions: [
    { value: "66", text: "bank", type: "trans" },
    { value: "67", text: "cash", type: "trans" },
    { value: "61", text: "delivery", type: "trans" },
    { value: "65", text: "formula", type: "trans" },
    { value: "62", text: "inventory", type: "trans" },
    { value: "55", text: "invoice", type: "trans" },
    { value: "58", text: "offer", type: "trans" },
    { value: "57", text: "order", type: "trans" },
    { value: "64", text: "production", type: "trans" },
    { value: "60", text: "rent", type: "trans" },
    { value: "63", text: "waybill", type: "trans" },
    { value: "59", text: "worksheet", type: "trans" },
    { value: "3", text: "ntr_bank_en", type: "report" },
    { value: "4", text: "ntr_cash_in_en", type: "report" },
    { value: "5", text: "ntr_cash_out_en", type: "report" },
  ],
  inputfilter: 109,
  inputfilterOptions: [
    { value: "109", text: "all" },
    { value: "106", text: "disabled" },
    { value: "107", text: "readonly" },
    { value: "108", text: "update" }
  ],
  supervisor: 1,
  className: "light",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
  onClose: undefined,
  onAudit: undefined,
}

export const Existing = Template.bind({});
Existing.args = {
  ...Default.args,
  className: "dark",
  idKey: 1,
  nervatype: 31,
  subtype: 61,
  inputfilter: 107,
  supervisor: 0,
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  nervatype: 18,
}