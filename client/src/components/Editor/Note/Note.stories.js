import { Note } from "./Note";

import { getText, store } from 'config/app';

export default {
  title: "Editor/Note",
  component: Note,
}

const Template = (args) => <Note {...args} />

export const Default = Template.bind({});
Default.args = {
  value: "<p>A long and <strong><em>rich text</em></strong> at the bottom of the invoice...</p><p>Can be multiple lines ...</p>",
  patternId: 1,
  patterns: [
    { id: 1, description: "first pattern", transtype: 55, 
      notes: "pattern text", defpattern: 0, deleted: 0 },
    { id: 2, description: "default pattern", transtype: 55, 
      notes: null, defpattern: 1, deleted: 0 },
  ],
  className: "light",
  readOnly: false,
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const Empty = Template.bind({});
Empty.args = {
  value: null,
  patternId: undefined,
  patterns: [],
  className: "light",
  readOnly: false,
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const ReadOnly = Template.bind({});
ReadOnly.args = {
  ...Default.args,
  readOnly: true,
  className: "dark",
}