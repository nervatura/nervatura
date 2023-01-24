import { expect } from '@open-wc/testing';

import { Quick } from './Quick.js'

it('Quick', () => {
  expect(Quick.customer().sql).to.exist
  expect(Quick.employee().sql).to.exist
  expect(Quick.payment().sql).to.exist
  expect(Quick.place().sql).to.exist
  expect(Quick.place_bank().sql).to.exist
  expect(Quick.place_cash().sql).to.exist
  expect(Quick.place_warehouse().sql).to.exist
  expect(Quick.product().sql).to.exist
  expect(Quick.product_item().sql).to.exist
  expect(Quick.project().sql).to.exist
  expect(Quick.report(1).sql).to.exist
  expect(Quick.servercmd(1).sql).to.exist
  expect(Quick.tool().sql).to.exist
  expect(Quick.transitem().sql).to.exist
  expect(Quick.transitem_invoice().sql).to.exist
  expect(Quick.transitem_delivery().sql).to.exist
  expect(Quick.transmovement().sql).to.exist
  expect(Quick.transpayment().sql).to.exist
})