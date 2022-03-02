import { Label } from "./Label";

import Icon from 'components/Form/Icon'

export default {
  title: "Form/Label",
  component: Label
}

const Template = (args) => <Label {...args} />

export const Default = Template.bind({});
Default.args = {
  value: "Label",
}

export const LeftIcon = Template.bind({});
LeftIcon.args = {
  ...Default.args,
  leftIcon: <Icon iconKey="InfoCircle" />,
}

export const RightIcon = Template.bind({});
RightIcon.args = {
  ...Default.args,
  rightIcon: <Icon iconKey="InfoCircle" />,
  iconWidth: "20px",
  style: { width: "100px" },
}

export const Centered = Template.bind({});
Centered.args = {
  ...Default.args,
  leftIcon: <Icon iconKey="InfoCircle" />,
  center: true
}