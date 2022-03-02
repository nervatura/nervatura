import { InputBox } from "./InputBox";

export default {
  title: "Modal/InputBox",
  component: InputBox
}

const Template = (args) => <InputBox {...args} />

export const Default = Template.bind({});
Default.args = {
  className: "light",
  title: "Input value title",
  message: "Input value message",
  infoText: undefined,
  value: "",
  labelCancel: "Cancel",
  labelOK: "OK",
  defaultOK: true,
  showValue: false,
}

export const InputValue = Template.bind({});
InputValue.args = {
  ...Default.args,
  className: "dark",
  infoText: "Input value info text",
  value: "default value",
  showValue: true,
}

export const DefaultCancel = Template.bind({});
DefaultCancel.args = {
  ...Default.args,
  defaultOK: false,
}