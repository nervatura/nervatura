import { MenuBar, SIDE_VISIBILITY, APP_MODULE } from "./MenuBar";

import { getText, store } from 'config/app';

export default {
  title: "MenuBar",
  component: MenuBar,
  argTypes: {
    loadModule: {
      name: "loadModule",
      description: "loadModule click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "loadModule" 
    },
  }
}

const Template = (args) => <MenuBar {...args} />

export const Default = Template.bind({});
Default.args = {
  side: SIDE_VISIBILITY.AUTO,
  scrollTop: false,
  module: APP_MODULE.SEARCH,
  loadModule: undefined, 
  sideBar: undefined, 
  setScroll: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const ScrollTop = Template.bind({});
ScrollTop.args = {
  ...Default.args,
  side: SIDE_VISIBILITY.SHOW,
  scrollTop: true,
  module: APP_MODULE.SETTING,
}