import { Button } from "./Button";

import Label from "components/Form/Label";
import Icon from 'components/Form/Icon'

export default {
  title: "Form/Button",
  component: Button
}

const Template = (args) => <Button {...args} />

export const Default = Template.bind({});
Default.args = {
  label: "Label",
  className: "light",
  small: false
}

export const PrimaryIconLabel = Template.bind({});
PrimaryIconLabel.args = {
  ...Default.args,
  value: <Label value="Label" leftIcon={<Icon iconKey="QuestionCircle" height={14} width={14} />} />,
  className: "primary"
}

export const IconBorderButton = Template.bind({});
IconBorderButton.args = {
  ...Default.args,
  value: <Icon iconKey="QuestionCircle" height={16} width={16} />,
  className: "light border-button"
}

export const Disabled = Template.bind({});
Disabled.args = {
  ...Default.args,
  disabled: "disabled"
}

export const SmallButton = Template.bind({});
SmallButton.args = {
  ...Default.args,
  small: true
}