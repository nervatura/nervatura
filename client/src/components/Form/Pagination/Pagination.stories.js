import { Pagination } from "./Pagination";

export default {
  title: "Form/Pagination",
  component: Pagination
}

const Template = (args) => <Pagination {...args} />

export const Default = Template.bind({});
Default.args = {
  pageIndex: 0, 
  pageSize: 5, 
  pageCount: 0, 
  canPreviousPage: false,
  canNextPage: false,
  hidePageSize: true,
  className: "light",
  onEvent: undefined,
}

export const Items = Template.bind({});
Items.args = {
  ...Default.args,
  pageIndex: 1, 
  pageSize: 10, 
  pageCount: 3,  
  canPreviousPage: true,
  canNextPage: true,
  hidePageSize: false,
  className: "dark",
}