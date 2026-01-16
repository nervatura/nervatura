import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './form-pagination.js';

import { APP_THEME } from '../../../config/enums.js'

export default {
  title: 'Form/Pagination',
  component: 'form-pagination',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: Object.values(APP_THEME),
    },
    onEvent: {
      name: "onEvent",
      description: "onEvent click handler",
      table: {
        type: { 
          summary: "func", 
        },
      },
      action: "onEvent" 
    }
  }
}

export function Template({ 
  id, theme, name, 
  pageIndex, pageSize, pageCount, canPreviousPage, canNextPage, hidePageSize, style,
  onEvent
}) {
  const component = html`<form-pagination
    id="${id}"
    name="${name}"
    pageIndex="${pageIndex}"
    pageSize="${pageSize}"
    pageCount="${pageCount}"
    ?canPreviousPage="${canPreviousPage}"
    ?canNextPage="${canNextPage}"
    ?hidePageSize="${hidePageSize}"
    .style="${style}"
    .onEvent=${onEvent}
  ></form-pagination>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_pagination",
  theme: APP_THEME.LIGHT,
  name: undefined,
  pageIndex: 0, 
  pageSize: 5, 
  pageCount: 0, 
  canPreviousPage: false,
  canNextPage: false,
  hidePageSize: true,
  style: {}
}

export const Items = Template.bind({});
Items.args = {
  ...Default.args,
  theme: APP_THEME.DARK,
  pageIndex: 1, 
  pageSize: 10, 
  pageCount: 3,  
  canPreviousPage: true,
  canNextPage: true,
  hidePageSize: false
}