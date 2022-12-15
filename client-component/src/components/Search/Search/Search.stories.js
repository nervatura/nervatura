import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './client-search.js';

import { APP_THEME, SIDE_VISIBILITY } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

import { Default as BrowserDefault } from '../Browser/Browser.stories.js'
import { Default as SideBarDefault } from '../../SideBar/Search/Search.stories.js'
import { Default as QuickDefault } from '../../Modal/Selector/Selector.stories.js'
import { Queries } from '../../../controllers/Queries.js'
import { Quick } from '../../../controllers/Quick.js'

export default {
  title: 'Search/Search',
  component: 'client-search',
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
    onBrowserEvent: {
      name: "onBrowserEvent",
      description: "onBrowserEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onBrowserEvent" 
    },
    onModalEvent: {
      name: "onModalEvent",
      description: "onModalEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onModalEvent" 
    },
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, data, side, auditFilter, paginationPage, theme, 
  onSideEvent, onBrowserEvent, onModalEvent
}) {
  const component = html`<client-search
    id="${id}"
    .data=${data}
    side="${side}"
    .auditFilter="${auditFilter}"
    paginationPage=${paginationPage}
    .queries=${Queries({ msg: (key)=> key })}
    .quick=${{...Quick}}
    .onEvent=${{ 
      onSideEvent, onBrowserEvent, onModalEvent, setModule: ()=>{}
    }}
    .msg="${msg}"
  ></client-search>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "search",
  theme: APP_THEME.LIGHT,
  side: SIDE_VISIBILITY.AUTO,
  data: {
    seltype: "selector",
    group_key: "customer",
    qview: "customer",
    qfilter: "",
    result: QuickDefault.args.result,
  },
  auditFilter: SideBarDefault.args.auditFilter,
  paginationPage: 10
}

export const Browser = Template.bind({});
Browser.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  data: {
    ...BrowserDefault.args.data,
    seltype: "browser",
    group_key: "customer",
  },
}