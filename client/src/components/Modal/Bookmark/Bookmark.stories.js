import { Bookmark } from "./Bookmark";

import { getText, store } from 'config/app';

export default {
  title: "Modal/Bookmark",
  component: Bookmark
}

const Template = (args) => <Bookmark {...args} />

export const Default = Template.bind({});
Default.args = {
  className: "light",
  bookmark: { history: null, bookmark: [] },
  tabView: "bookmark",
  getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }),
}

export const BookmarkData = Template.bind({});
BookmarkData.args = {
  ...Default.args,
  bookmark: {
    history: {
      cfgroup: '2021-11-23T23:05:40+02:00',
      cfname: 3,
      cfvalue: '[{"datetime":"2021-11-23T23:05:40+02:00","type":"delete","ntype":"product","transtype":"","id":14,"title":"PRODUCT | PRO/00001"},{"datetime":"2021-11-23T23:05:33+02:00","type":"save","ntype":"product","transtype":"","id":null,"title":"PRODUCT | null"},{"datetime":"2021-11-23T23:04:17+02:00","type":"save","ntype":"trans","transtype":"invoice","id":5,"title":"INVOICE | DMINV/00001"}]',
      employee_id: 1,
      id: 3,
      orderby: 0,
      section: 'history'
    },
    bookmark: [
      {
        cfgroup: 'editor',
        cfname: 'First Customer Co.',
        cfvalue: '{"date":"2021-11-23","ntype":"customer","transtype":"","id":2,"info":"DMCUST/00001"}',
        employee_id: 1,
        id: 1,
        orderby: 0,
        section: 'bookmark'
      },
      {
        cfgroup: 'browser',
        cfname: 'Invoices',
        cfvalue: '{"date":"2021-11-23","vkey":"transitem","view":"TransItemHeadView","filters":[{"id":"1637701420094","fieldtype":"string","fieldname":"transtype","sqlstr":"case when mst.msg is null then tg.groupvalue else mst.msg end ","wheretype":"where","filtertype":"===","value":"invoice"}],"columns":{"transtype":true,"transnumber":true,"transdate":true,"custname":true}}',
        employee_id: 1,
        id: 2,
        orderby: 0,
        section: 'bookmark'
      },
      {
        cfgroup: 'editor',
        cfname: 'Big product',
        cfvalue: '{"date":"2021-11-23","ntype":"product","transtype":"","id":1,"info":"DMPROD/00001"}',
        employee_id: 1,
        id: 4,
        orderby: 0,
        section: 'bookmark'
      },
      {
        cfgroup: 'editor',
        cfname: 'DMINV/00001',
        cfvalue: '{"date":"2021-11-23","ntype":"trans","transtype":"invoice","id":5,"info":"First Customer Co."}',
        employee_id: 1,
        id: 5,
        orderby: 0,
        section: 'bookmark'
      }
    ]
  },
}

export const DarkTheme = Template.bind({});
DarkTheme.args = {
  ...BookmarkData.args,
  className: "dark",
}