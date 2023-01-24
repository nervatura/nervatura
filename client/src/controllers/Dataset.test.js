import { expect } from '@open-wc/testing';

import { Dataset } from './Dataset.js'

it('Dataset', () => {

  expect(Dataset.currency().length).greaterThanOrEqual(0)
  expect(Dataset.customer().length).greaterThanOrEqual(0)
  expect(Dataset.deffield().length).greaterThanOrEqual(0)
  expect(Dataset.employee().length).greaterThanOrEqual(0)
  expect(Dataset.event().length).greaterThanOrEqual(0)
  expect(Dataset.groups().length).greaterThanOrEqual(0)
  expect(Dataset.numberdef().length).greaterThanOrEqual(0)
  expect(Dataset.place().length).greaterThanOrEqual(0)
  expect(Dataset.printqueue().length).greaterThanOrEqual(0)
  expect(Dataset.product().length).greaterThanOrEqual(0)
  expect(Dataset.project().length).greaterThanOrEqual(0)
  expect(Dataset.rate().length).greaterThanOrEqual(0)
  expect(Dataset.report().length).greaterThanOrEqual(0)
  expect(Dataset.setting().length).greaterThanOrEqual(0)
  expect(Dataset.tax().length).greaterThanOrEqual(0)
  expect(Dataset.template().length).greaterThanOrEqual(0)
  expect(Dataset.tool().length).greaterThanOrEqual(0)
  expect(Dataset.trans("bank").length).greaterThanOrEqual(0)
  expect(Dataset.trans("invoice").length).greaterThanOrEqual(0)
  expect(Dataset.trans("formula").length).greaterThanOrEqual(0)
  expect(Dataset.trans("production").length).greaterThanOrEqual(0)
  expect(Dataset.trans("delivery").length).greaterThanOrEqual(0)
  expect(Dataset.trans("inventory").length).greaterThanOrEqual(0)
  expect(Dataset.trans("waybill").length).greaterThanOrEqual(0)
  expect(Dataset.trans("order").length).greaterThanOrEqual(0)
  expect(Dataset.trans("").length).greaterThanOrEqual(0)
  expect(Dataset.ui_menu().length).greaterThanOrEqual(0)
  expect(Dataset.usergroup().length).greaterThanOrEqual(0)

})