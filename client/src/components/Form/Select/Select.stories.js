import { Select } from "./Select";

export default {
  title: "Form/Select",
  component: Select
}

const Template = (args) => <Select {...args} />

export const Default = Template.bind({});
Default.args = {
  value: "value1",
  options: [
    { value: "value1", text: "Text 1" },
    { value: "value2", text: "Text 2" },
    { value: "value3", text: "Text 3" }
  ],
  className: "light",
  //onChange: (value) => {}
}

export const Placeholder = Template.bind({});
Placeholder.args = {
  ...Default.args,
  value: "",
  placeholder: "Placeholder text",
}