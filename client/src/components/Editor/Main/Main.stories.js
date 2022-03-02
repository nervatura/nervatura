import { Main } from "./Main";

import { getText, store } from 'config/app';
import { Default as EditorDefault, Report as EditorReport } from 'components/Editor/Editor/Editor.stories'

export default {
  title: "Editor/Main",
  component: Main,
}

const Template = (args) => <Main {...args} />

export const Default = Template.bind({});
Default.args = {
  current: EditorDefault.args.current, 
  template: EditorDefault.args.template, 
  dataset: EditorDefault.args.dataset, 
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const Report = Template.bind({});
Report.args = {
  current: EditorReport.args.current, 
  template: EditorReport.args.template, 
  dataset: EditorReport.args.dataset, 
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}