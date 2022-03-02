import { Template as SideTemplate, SIDE_VISIBILITY } from "./Template";

import { getText, store } from 'config/app';

export default {
  title: "SideBar/Template",
  component: SideTemplate,
}

const Template = (args) => <SideTemplate {...args} />

export const Default = Template.bind({});
Default.args = {
  side: SIDE_VISIBILITY.AUTO,
  templateKey: "template",
  dirty: false, 
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const Sample = Template.bind({});
Sample.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  templateKey: "_sample",
}

export const Dirty = Template.bind({});
Dirty.args = {
  ...Default.args,
  dirty: true,
}