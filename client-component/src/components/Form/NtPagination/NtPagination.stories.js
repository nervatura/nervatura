import { html } from 'lit';

import '../../StoryContainer/story-container.js';
import './nt-pagination.js';

export default {
  title: 'Form/NtPagination',
  component: 'nt-pagination',
  excludeStories: ['Template'],
  argTypes: {
    theme: {
      control: 'select',
      options: ["light", "dark"],
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
  const component = html`<nt-pagination
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
  ></nt-pagination>`
  return html`<story-container theme="${theme}">${component}</story-container>`;
}

export const Default = Template.bind({});
Default.args = {
  id: "test_pagination",
  theme: "light",
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
  theme: "dark",
  pageIndex: 1, 
  pageSize: 10, 
  pageCount: 3,  
  canPreviousPage: true,
  canNextPage: true,
  hidePageSize: false
}