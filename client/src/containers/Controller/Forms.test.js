import { Forms } from './Forms'
import { getText, store } from 'config/app';

const form = Forms({ getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }) })

describe('Forms', () => {
  it('address', () => {
    let address = form.address()
    expect(address.options.icon).toBe("Home")
    address = form.address({ id: null })
    expect(address.options.panel["new"]).toBeDefined()
    address = form.address({ id: 1 })
    expect(address.options.panel["new"]).toBeUndefined()
  })

  it('bank', () => {
    let bank = form.bank()
    expect(bank.options.icon).toBe("Money")
    bank = form.bank({ id: null })
    expect(bank.options.panel["new"]).toBeDefined()
    bank = form.bank({ id: 1 }, { dataset: { translink: [] } })
    expect(bank.options.panel["new"]).toBeUndefined()
    bank = form.bank({ id: 1 }, { dataset: { translink: [{ transtype: "transtype" }] } })
    expect(bank.rows[0].columns[0].name).toBe("id")
  })

  it('barcode', () => {
    let barcode = form.barcode()
    expect(barcode.options.icon).toBe("Barcode")
    barcode = form.barcode({ id: null })
    expect(barcode.options.panel["new"]).toBeDefined()
    barcode = form.barcode({ id: 1 })
    expect(barcode.options.panel["new"]).toBeUndefined()
  })

  it('cash', () => {
    let cash = form.cash()
    expect(cash.options.icon).toBe("Money")
    cash = form.cash(
      { id: null, direction: 1 }, 
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(cash.options.panel["new"]).toBeDefined()
    cash = form.cash(
      { id: 1, direction: 1 }, 
      { dataset: { translink: [], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(cash.options.panel["new"]).toBeUndefined()
    cash = form.cash(
      { id: 1, direction: 1 }, 
      { dataset: { translink: [{ transtype: "transtype" }], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(cash.rows[3].columns[0].map["source"]).toBe("translink")
    cash = form.cash(
      { id: 1, direction: 1 }, 
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(cash.rows[3].columns[0].map["source"]).toBe("cancel_link")
  })

  it('contact', () => {
    let contact = form.contact()
    expect(contact.options.icon).toBe("Phone")
    contact = form.contact({ id: null })
    expect(contact.options.panel["new"]).toBeDefined()
    contact = form.contact({ id: 1 })
    expect(contact.options.panel["new"]).toBeUndefined()
  })

  it('currency', () => {
    let currency = form.currency()
    expect(currency.options.icon).toBe("Dollar")
    currency = form.currency({ id: null })
    expect(currency.options.panel["new"]).toBeDefined()
    currency = form.currency({ id: 1 })
    expect(currency.options.panel["new"]).toBeUndefined()
  })

  it('customer', () => {
    let customer = form.customer()
    expect(customer.options.icon).toBe("User")
    customer = form.customer(
      { id: null, custtype: 1 },
      { dataset: { custtype: {}, groups: [ { id: 1, groupname: "custtype", groupvalue: "own" } ] } }
    )
    expect(customer.options.panel["new"]).toBeDefined()
    customer = form.customer(
      { id: 1, custtype: 2 },
      { dataset: { custtype: {}, groups: [ { id: 1, groupname: "custtype", groupvalue: "own" } ] } }
    )
    expect(customer.options.panel["new"]).toBeUndefined()
    customer = form.customer(
      { id: 1, custtype: 2 },
      { dataset: { groups: [ { id: 1, groupname: "custtype", groupvalue: "own" } ] } }
    )
    expect(customer.options.panel["new"]).toBeUndefined()
  })

  it('deffield', () => {
    let deffield = form.deffield()
    expect(deffield.options.icon).toBe("Tag")
    deffield = form.deffield({ id: null })
    expect(deffield.options.panel["new"]).toBeDefined()
    deffield = form.deffield(
      { id: 1, fieldtype: 1 },
      { dataset: { fieldtype: [ { id: 1, groupvalue: "valuelist" } ] } }
    )
    expect(deffield.options.panel["new"]).toBeUndefined()
    deffield = form.deffield(
      { id: 1, fieldtype: 1 },
      { dataset: { fieldtype: [ { id: 2, groupvalue: "valuelist" } ] } }
    )
    expect(deffield.options.panel["new"]).toBeUndefined()
  })

  it('delivery', () => {
    let delivery = form.delivery()
    expect(delivery.options.icon).toBe("Truck")
    delivery = form.delivery(
      { id: null, direction: 1 },
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.options.panel["new"]).toBeDefined()
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.rows[0].columns[1].name).toBe("ref_transnumber")
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { translink: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.rows[0].columns[1].name).toBe("id")
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.rows[0].columns[1].name).toBe("id")
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { cancel_link: [{ transtype: "in" }], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    delivery = form.delivery(
      { id: null, direction: 1 },
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(delivery.options.panel["new"]).toBeDefined()
  })

  it('discount', () => {
    let discount = form.discount()
    expect(discount.options.icon).toBe("Dollar")
    discount = form.discount({ id: null })
    expect(discount.options.panel["new"]).toBeDefined()
    discount = form.discount({ id: 1 })
    expect(discount.options.panel["new"]).toBeUndefined()
  })

  it('employee', () => {
    let employee = form.employee()
    expect(employee.options.icon).toBe("Male")
    employee = form.employee({ id: null })
    expect(employee.options.panel["new"]).toBeDefined()
    employee = form.employee({ id: 1 })
    expect(employee.options.panel["new"]).toBeUndefined()
  })

  it('event', () => {
    let event = form.event()
    expect(event.options.icon).toBe("Calendar")
    event = form.event({ id: null })
    expect(event.options.panel["new"]).toBeDefined()
    event = form.event({ id: 1 })
    expect(event.options.panel["new"]).toBeUndefined()
  })

  it('formula', () => {
    let formula = form.formula()
    expect(formula.options.icon).toBe("Magic")
    formula = form.formula({ id: null })
    expect(formula.options.panel["new"]).toBeDefined()
    formula = form.formula(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(formula.options.panel["new"]).toBeUndefined()
    formula = form.formula(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(formula.rows[2].columns[0].name).toBe("id")
  })

  it('groups', () => {
    let groups = form.groups()
    expect(groups.options.icon).toBe("Th")
    groups = form.groups({ id: null })
    expect(groups.options.panel["new"]).toBeDefined()
    groups = form.groups({ id: 1 })
    expect(groups.options.panel["new"]).toBeUndefined()
  })

  it('inventory', () => {
    let inventory = form.inventory()
    expect(inventory.options.icon).toBe("Truck")
    inventory = form.inventory({ id: null })
    expect(inventory.options.panel["new"]).toBeDefined()
    inventory = form.inventory({ id: 1 })
    expect(inventory.options.panel["new"]).toBeUndefined()
  })

  it('invoice_link', () => {
    let invoice_link = form.invoice_link()
    expect(invoice_link.options.icon).toBe("Money")
    invoice_link = form.invoice_link({ id: null })
    expect(invoice_link.options.panel["new"]).toBeDefined()
    invoice_link = form.invoice_link({ id: 1 })
    expect(invoice_link.options.panel["new"]).toBeUndefined()
  })

  it('invoice', () => {
    let invoice = form.invoice()
    expect(invoice.options.icon).toBe("FileText")
    invoice = form.invoice({ id: null })
    expect(invoice.options.panel["new"]).toBeDefined()
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal" },
      { dataset: { translink: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(invoice.options.panel["new"]).toBeUndefined()
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal", deleted: 0 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(invoice.options.panel["corrective"]).toBeDefined()
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(invoice.options.panel["cancellation"]).toBeDefined()
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(invoice.options.panel["cancellation"]).toBeDefined()
  })

  it('item', () => {
    let item = form.item()
    expect(item.options.icon).toBe("ListOl")
    item = form.item(
      { id: null },
      { current: { transtype: "invoice" } }
    )
    expect(item.options.panel["new"]).toBeDefined()
    item = form.item(
      { id: 1 },
      { current: { transtype: "offer" } }
    )
    expect(item.options.panel["new"]).toBeUndefined()
    item = form.item(
      { id: 1 },
      { current: { transtype: "order" } }
    )
    expect(item.options.panel["new"]).toBeUndefined()
  })

  it('log', () => {
    let log = form.log()
    expect(log.options.icon).toBe("InfoCircle")
  })

  it('movement', () => {
    let movement = form.movement(
      undefined,
      { current: { transtype: "default" } }
    )
    expect(movement.options.icon).toBe("Truck")
    movement = form.movement(
      { id: null },
      { current: { transtype: "delivery" } }
    )
    expect(movement.options.panel["new"]).toBeDefined()
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "inventory" } }
    )
    expect(movement.options.panel["new"]).toBeUndefined()
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "production" } }
    )
    expect(movement.options.panel["new"]).toBeUndefined()
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "formula" } }
    )
    expect(movement.options.panel["new"]).toBeUndefined()
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "waybill" } }
    )
    expect(movement.options.panel["new"]).toBeUndefined()
  })

  it('numberdef', () => {
    let numberdef = form.numberdef()
    expect(numberdef.options.icon).toBe("ListOl")
  })

  it('offer', () => {
    let offer = form.offer()
    expect(offer.options.icon).toBe("FileText")
    offer = form.offer({ id: null })
    expect(offer.options.panel["new"]).toBeDefined()
    offer = form.offer(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(offer.options.panel["new"]).toBeUndefined()
    offer = form.offer(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(offer.options.panel["new"]).toBeUndefined()
  })

  it('order', () => {
    let order = form.order()
    expect(order.options.icon).toBe("FileText")
    order = form.order({ id: null })
    expect(order.options.panel["new"]).toBeDefined()
    order = form.order(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(order.options.panel["new"]).toBeUndefined()
    order = form.order(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(order.options.panel["new"]).toBeUndefined()
  })

  it('password', () => {
    let password = form.password()
    expect(password.options.icon).toBe("Lock")
  })

  it('payment_link', () => {
    let payment_link = form.payment_link()
    expect(payment_link.options.icon).toBe("FileText")
    payment_link = form.payment_link({ id: null })
    expect(payment_link.options.panel["new"]).toBeDefined()
    payment_link = form.payment_link({ id: 1 })
    expect(payment_link.options.panel["new"]).toBe(false)
  })

  it('payment', () => {
    let payment = form.payment()
    expect(payment.options.icon).toBe("Money")
    payment = form.payment({ id: null })
    expect(payment.options.panel["new"]).toBeDefined()
    payment = form.payment({ id: 1 })
    expect(payment.options.panel["new"]).toBeUndefined()
  })

  it('place', () => {
    let place = form.place()
    expect(place.options.icon).toBe("Map")
    place = form.place({ id: null })
    expect(place.options.panel["new"]).toBeDefined()
    place = form.place(
      { id: 1, placetype: 2 },
      { dataset: { placetype: [ { id: 1, groupvalue: "warehouse" } ] } }
    )
    expect(place.options.panel["new"]).toBeUndefined()
    place = form.place(
      { id: 1, placetype: 1 },
      { dataset: { groups: [ { id: 1, groupname: "placetype", groupvalue: "warehouse" } ] } }
    )
    expect(place.options.panel["new"]).toBeUndefined()
  })

  it('price', () => {
    let price = form.price()
    expect(price.options.icon).toBe("Dollar")
    price = form.price({ id: null })
    expect(price.options.panel["new"]).toBeDefined()
    price = form.price({ id: 1 })
    expect(price.options.panel["new"]).toBeUndefined()
  })

  it('printqueue', () => {
    let printqueue = form.printqueue(
      undefined, undefined,
      { printqueue_type: "printqueue_type", printqueue_mode: "printqueue_mode",
        report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size"
      }
    )
    expect(printqueue.options.icon).toBe("Filter")
  })

  it('product', () => {
    let product = form.product()
    expect(product.options.icon).toBe("ShoppingCart")
    product = form.product({ id: null })
    expect(product.options.panel["new"]).toBeDefined()
    product = form.product({ id: 1 })
    expect(product.options.panel["new"]).toBeUndefined()
  })

  it('production', () => {
    let production = form.production()
    expect(production.options.icon).toBe("Flask")
    production = form.production({ id: null })
    expect(production.options.panel["new"]).toBeDefined()
    production = form.production(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(production.options.panel["new"]).toBeUndefined()
    production = form.production(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(production.options.panel["new"]).toBeUndefined()
  })

  it('project', () => {
    let project = form.project()
    expect(project.options.icon).toBe("Clock")
    project = form.project({ id: null })
    expect(project.options.panel["new"]).toBeDefined()
    project = form.project({ id: 1 })
    expect(project.options.panel["new"]).toBeUndefined()
  })

  it('rate', () => {
    let rate = form.rate()
    expect(rate.options.icon).toBe("Strikethrough")
    rate = form.rate(
      { id: null },
      { dataset: { ratetype: [ { id: 1, groupvalue: "rate" } ], settings: [ { id: 1, fieldname: "default_currency", value: "EUR" } ] } }
    )
    expect(rate.options.panel["new"]).toBeDefined()
    rate = form.rate(
      { id: null },
      { dataset: { ratetype: [ { id: 1, groupvalue: "rate" } ], settings: [] } }
    )
    expect(rate.options.panel["new"]).toBeDefined()
    rate = form.rate({ id: 1 })
    expect(rate.options.panel["new"]).toBeUndefined()
  })

  it('program', () => {
    let program = form.program()
    expect(program.options.icon).toBe("Keyboard")
  })

  it('receipt', () => {
    let receipt = form.receipt()
    expect(receipt.options.icon).toBe("FileText")
    receipt = form.receipt({ id: null })
    expect(receipt.options.panel["new"]).toBeDefined()
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal" },
      { dataset: { translink: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(receipt.options.panel["new"]).toBeUndefined()
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal", deleted: 0 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(receipt.options.panel["corrective"]).toBeDefined()
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(receipt.options.panel["cancellation"]).toBeDefined()
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(receipt.options.panel["cancellation"]).toBeDefined()
  })

  it('rent', () => {
    let rent = form.rent()
    expect(rent.options.icon).toBe("FileText")
    rent = form.rent({ id: null })
    expect(rent.options.panel["new"]).toBeDefined()
    rent = form.rent(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(rent.options.panel["new"]).toBeUndefined()
    rent = form.rent(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(rent.options.panel["new"]).toBeUndefined()
  })

  it('report', () => {
    let report = form.report(
      undefined, undefined,
      { report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size" }
    )
    expect(report.options.panel["print"]).toBeDefined()
    report = form.report(
      { id: 1, ftype: "csv" }, undefined,
      { report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size" }
    )
    expect(report.options.panel["print"]).toBeDefined()
    report = form.report(
      { id: 1, ftype: "pdf" }, undefined,
      { report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size" }
    )
    expect(report.options.panel["print"]).toBeDefined()
  })

  it('setting', () => {
    let setting = form.setting()
    expect(setting.options.icon).toBe("Cog")
  })

  it('shipping', () => {
    let shipping = form.shipping()
    expect(shipping.options.icon).toBe("Truck")
  })

  it('tax', () => {
    let tax = form.tax()
    expect(tax.options.icon).toBe("Ticket")
    tax = form.tax({ id: null })
    expect(tax.options.panel["new"]).toBeDefined()
    tax = form.tax({ id: 1 })
    expect(tax.options.panel["new"]).toBeUndefined()
  })

  it('template', () => {
    let template = form.template()
    expect(template.options.icon).toBe("TextHeight")
  })

  it('tool', () => {
    let tool = form.tool()
    expect(tool.options.icon).toBe("Wrench")
    tool = form.tool({ id: null })
    expect(tool.options.panel["new"]).toBeDefined()
    tool = form.tool({ id: 1 })
    expect(tool.options.panel["new"]).toBeUndefined()
  })

  it('ui_menu', () => {
    let ui_menu = form.ui_menu()
    expect(ui_menu.options.icon).toBe("Share")
    ui_menu = form.ui_menu({ id: null })
    expect(ui_menu.rows[0].columns[0]["disabled"]).toBeUndefined()
    ui_menu = form.ui_menu({ id: 1 })
    expect(ui_menu.rows[0].columns[0]["disabled"]).toBeDefined()
  })

  it('usergroup', () => {
    let usergroup = form.usergroup()
    expect(usergroup.options.icon).toBe("Key")
    usergroup = form.usergroup({ id: null })
    expect(usergroup.options.panel["new"]).toBeDefined()
    usergroup = form.usergroup({ id: 1 })
    expect(usergroup.options.panel["new"]).toBeUndefined()
  })

  it('waybill', () => {
    let waybill = form.waybill()
    expect(waybill.options.icon).toBe("Briefcase")
    waybill = form.waybill({ id: null })
    expect(waybill.options.panel["new"]).toBeDefined()
    waybill = form.waybill(
      { id: 1, customer_id: null, employee_id: null },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(waybill.options.panel["new"]).toBeUndefined()
    waybill = form.waybill(
      { id: 1, customer_id: null, employee_id: null },
      { dataset: { translink: [] } }
    )
    expect(waybill.options.panel["new"]).toBeUndefined()
    waybill = form.waybill(
      { id: 1, customer_id: 1 },
    )
    expect(waybill.rows[1].columns[1]["map"]["seltype"]).toBe("customer")
    waybill = form.waybill(
      { id: 1, customer_id: null, employee_id: 1 },
    )
    expect(waybill.rows[1].columns[1]["map"]["seltype"]).toBe("employee")
  })

  it('worksheet', () => {
    let worksheet = form.worksheet()
    expect(worksheet.options.icon).toBe("FileText")
    worksheet = form.worksheet({ id: null })
    expect(worksheet.options.panel["new"]).toBeDefined()
    worksheet = form.worksheet(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(worksheet.options.panel["new"]).toBeUndefined()
    worksheet = form.worksheet(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(worksheet.options.panel["new"]).toBeUndefined()
  })

})