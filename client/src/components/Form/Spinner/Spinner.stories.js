import { Spinner } from "./Spinner";

export default {
  title: "Form/Spinner",
  component: Spinner
}

const Template = (args) => <Spinner {...args} />

export const Default = Template.bind({});
Default.args = {
  id: "app_loading"
}