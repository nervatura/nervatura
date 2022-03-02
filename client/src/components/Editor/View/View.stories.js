import { View } from "./View";

import { getText, store } from 'config/app';
import { Default as EditorDefault } from 'components/Editor/Editor/Editor.stories'

export default {
  title: "Editor/View",
  component: View,
}

const Template = (args) => <View {...args} />

export const Default = Template.bind({});
Default.args = {
  viewName: "item",
  current: EditorDefault.args.current, 
  template: EditorDefault.args.template, 
  dataset: EditorDefault.args.dataset, 
  audit: "all",
  className: "light",
  onEvent: undefined,
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key })
}

export const List = Template.bind({});
List.args = {
  ...Default.args,
  viewName: "invoice_link",
  className: "dark",
}

export const ReadOnly = Template.bind({});
ReadOnly.args = {
  ...Default.args,
  audit: "readonly",
}

export const ReadOnlyList = Template.bind({});
ReadOnlyList.args = {
  ...Default.args,
  viewName: "invoice_link",
  template: {
    ...EditorDefault.args.template,
    view: {
      invoice_link: {
        ...EditorDefault.args.template.view.invoice_link,
        actions: {
          new: null,
          edit: null,
          delete: null
        }
      }
    }
  },
}

export const Empty = Template.bind({});
Empty.args = {
  ...Default.args,
  template: {
    ...EditorDefault.args.template,
    view: {
      item: {
        ...EditorDefault.args.template.view.item,
        data: "missing",
        edit_icon: "Edit",
        delete_icon: "Times",
        new_icon: "Plus",
        edited: false,
        new_label: "NEW"
      }
    }
  }
}

export const DeleteOnly = Template.bind({});
DeleteOnly.args = {
  ...Default.args,
  template: {
    ...EditorDefault.args.template,
    view: {
      item: {
        ...EditorDefault.args.template.view.item,
        actions: {
          new: null,
          edit: null,
          delete: {action: "deleteEditorItem", fkey: "item"}
        }
      }
    }
  }
}