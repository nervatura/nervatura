import { Quick } from './Quick'
import { getText, store } from 'config/app';

const quick = Quick({ getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }) })

it('Quick', () => {
  expect(quick.customer().sql).toBeDefined()
  expect(quick.employee().sql).toBeDefined()
  expect(quick.payment().sql).toBeDefined()
  expect(quick.place().sql).toBeDefined()
  expect(quick.place_bank().sql).toBeDefined()
  expect(quick.place_cash().sql).toBeDefined()
  expect(quick.place_warehouse().sql).toBeDefined()
  expect(quick.product().sql).toBeDefined()
  expect(quick.product_item().sql).toBeDefined()
  expect(quick.project().sql).toBeDefined()
  expect(quick.report(1).sql).toBeDefined()
  expect(quick.servercmd(1).sql).toBeDefined()
  expect(quick.tool().sql).toBeDefined()
  expect(quick.transitem().sql).toBeDefined()
  expect(quick.transitem_invoice().sql).toBeDefined()
  expect(quick.transitem_delivery().sql).toBeDefined()
  expect(quick.transmovement().sql).toBeDefined()
  expect(quick.transpayment().sql).toBeDefined()
  
})