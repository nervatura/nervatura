import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './modal-formula.js';

import { APP_THEME } from '../../../config/enums.js'
import * as locales from '../../../config/locales.js';

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue

export default {
  title: 'Modal/Formula',
  component: 'modal-formula',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
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

export function Template({ 
  id, formula, formulaValues, partnumber, description, theme, onModalEvent 
}) {
  const component = html`<modal-formula
    id="${id}"
    formula="${formula}"
    .formulaValues="${formulaValues}"
    partnumber="${partnumber}"
    description="${description}"
    .onEvent=${{ onModalEvent }}
    .msg=${msg}
  ></modal-formula>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "formula",
  theme: APP_THEME.LIGHT,
  formula: "18",
  formulaValues: [
    {value: "18", text: 'DMFRM/00001'},
    {value: "19", text: 'DMFRM/00002'}
  ],
  partnumber: "DMPROD/00004",
  description: "Car",
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  formula: "",
  formulaValues: []
}