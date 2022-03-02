import { Input, INPUT_TYPE } from "./Input";
import { DECIMAL_SEPARATOR } from 'config/app'

export default {
  title: "Form/Input",
  component: Input
}

const Template = (args) => <Input {...args} />

export const Default = Template.bind({});
Default.args = {
  value: "value",
  className: "light",
  //onChange: (value)=>{},
  //onEnter: (value)=>{},
}

export const NumberInput = Template.bind({});
NumberInput.args = {
  ...Default.args,
  type: INPUT_TYPE.NUMBER,
  separator: DECIMAL_SEPARATOR.POINT,
  value: 123.55,
}

export const IntegerInput = Template.bind({});
IntegerInput.args = {
  ...Default.args,
  type: INPUT_TYPE.INTEGER,
  value: 111,
  minValue: 0,
  maxValue: 100,
}

export const CommaSeparator = Template.bind({});
CommaSeparator.args = {
  ...NumberInput.args,
  type: INPUT_TYPE.NUMBER,
  separator: DECIMAL_SEPARATOR.COMMA,
  value: 123.55,
  className: "dark"
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  disabled: "disabled",
}