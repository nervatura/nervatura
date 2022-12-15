import { html } from 'lit';
import { ifDefined } from 'lit/directives/if-defined.js';

import '../../StoryContainer/story-container.js';
import './edit-note.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

export default {
  title: 'Edit/Note',
  component: 'edit-note',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onEditEvent: {
      name: "onEditEvent",
      description: "onEditEvent handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEditEvent" 
    }
  }
};

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export function Template({ 
  id, value, patternId, patterns, readOnly, theme, onEditEvent
}) {
  const component = html`<edit-note
    id="${id}"
    value="${value}"
    patternId="${ifDefined(patternId)}"
    .patterns="${patterns}"
    ?readOnly="${readOnly}"
    .onEvent=${{ onEditEvent }}
    .msg=${msg}
  ></edit-note>`
  return html`<story-container theme="${theme}" >${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "note",
  theme: APP_THEME.LIGHT,
  value: "<p>A long and <strong><em>rich text</em></strong> at the bottom of the invoice...</p><p>Can be multiple lines ...</p>",
  patternId: 1,
  patterns: [
    { id: 1, description: "first pattern", transtype: 55, 
      notes: "pattern text", defpattern: 0, deleted: 0 },
    { id: 2, description: "default pattern", transtype: 55, 
      notes: null, defpattern: 1, deleted: 0 },
  ],
  readOnly: false,
}

export const Empty = Template.bind({});
Empty.args = {
  ...Default.args,
  value: null,
  patternId: undefined,
  patterns: [],
}

export const ReadOnly = Template.bind({});
ReadOnly.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  readOnly: true,
}
