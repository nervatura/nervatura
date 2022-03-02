import { Queries } from './Queries'
import { getText, store } from 'config/app';

const query = Queries({ getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }) })

it('Queries', () => {
  expect(query.customer().options).toBeDefined()
  expect(query.employee().options).toBeDefined()
  expect(query.product().options).toBeDefined()
  expect(query.project().options).toBeDefined()
  expect(query.rate().options).toBeDefined()
  expect(query.tool().options).toBeDefined()
  expect(query.transitem().options).toBeDefined()
  expect(query.transmovement().options).toBeDefined()
  expect(query.transpayment().options).toBeDefined()
})