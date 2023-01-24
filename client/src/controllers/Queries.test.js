import { expect } from '@open-wc/testing';

import { Queries } from './Queries.js'

const query = Queries({ msg: (key)=> key })

it('Queries', () => {
  expect(query.customer().options).to.exist
  expect(query.employee().options).to.exist
  expect(query.product().options).to.exist
  expect(query.project().options).to.exist
  expect(query.rate().options).to.exist
  expect(query.tool().options).to.exist
  expect(query.transitem().options).to.exist
  expect(query.transmovement().options).to.exist
  expect(query.transpayment().options).to.exist
})