import { Item } from "./Item";

import { getText as appGetText, store } from 'config/app';
import { Forms } from 'containers/Controller/Forms'
import { Item as EditorItem } from 'components/Editor/Editor/Editor.stories'

export default {
  title: "Editor/Item",
  component: Item,
}

const Template = (args) => <Item {...args} />
const getText = (key)=>appGetText({ locales: store.session.locales, lang: "en", key: key })

export const Default = Template.bind({});
Default.args = {
  current: EditorItem.args.current,  
  dataset: EditorItem.args.dataset, 
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: getText
}

export const NewItem = Template.bind({});
NewItem.args = {
  ...Default.args,
  current: {
    ...Default.args.current,
    form: {
      ...Default.args.current.form,
      id: null,
    },
    form_template: Forms({ getText: getText })["item"]({ id: null }, { current: { transtype: "invoice" } })
  },
  className: "dark",
}