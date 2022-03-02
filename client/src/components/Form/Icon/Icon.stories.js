import Icon, { ICON_KEY } from "./Icon";

export default {
  title: "Form/Icon",
  component: Icon,
  argTypes: {
    iconKey: {
      options: ICON_KEY,
    }
  }
}

const Template = (args) => <Icon {...args} />

export const Default = Template.bind({});
Default.args = {
  iconKey: "ExclamationTriangle",
}

export const ColorPointer = Template.bind({});
ColorPointer.args = {
  iconKey: "Copy",
  width: 42,
  height: 48,
  color: "red",
  onClick: undefined,
  cursor: "pointer"
}