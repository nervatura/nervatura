import { DateTime } from "./DateTime";
import formatISO from 'date-fns/formatISO'

export default {
  title: "Form/DateTime",
  component: DateTime
}

const Template = (args) => <DateTime {...args} />

export const Default = Template.bind({});
Default.args = {
  className: "light",
  value: formatISO(new Date()),
  dateTime: true,
  isEmpty: true,
  showTimeSelectOnly: false,
}

export const DateInput = Template.bind({});
DateInput.args = {
  ...Default.args,
  className: "dark",
  value: formatISO(new Date(), { representation: 'date' }),
  dateTime: false,
  isEmpty: false,
  locale: "de"
}

export const TimeInput = Template.bind({});
TimeInput.args = {
  ...Default.args,
  value: formatISO(new Date(), { representation: 'time' }),
  showTimeSelectOnly: true,
}
