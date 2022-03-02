import { getText, getSetting, formatNumber } from './app'

it('getText', () => {
  const locales = {
    en: {
      lkey: "en_value"
    }
  }
  let value = getText({locales, lang: "en", key: "lkey", defaultValue: "value" })
  expect(value).toBe("en_value")
  value = getText({locales, lang: "de", key: "lkey" })
  expect(value).toBe("en_value")
  value = getText({locales, lang: "de", key: "mkey", defaultValue: "value" })
  expect(value).toBe("value")
})

it('getSetting', () => {
  localStorage.setItem("toastTime", 7000)
  let value = getSetting("ui")
  expect(value).toBeDefined()
  value = getSetting("toastTime")
  expect(value).toBe("7000")
  value = getSetting("missing")
  expect(value).toBe("")
})

it('formatNumber', () => {
  formatNumber(123.123, 0);
  formatNumber(123.123);
  formatNumber("number");
})