import { expect } from '@open-wc/testing';
import { Forms } from './Forms.js'
import * as locales from '../config/locales.js';
import { store as storeConfig } from '../config/app.js'

const msg = (defaultValue, props) => locales.en[props.id] || defaultValue
const form = Forms({ msg: (key)=>msg(key,{ id: key }), getSetting: (key)=>storeConfig.ui[key] })

describe('Forms', () => {
  it('address', () => {
    let address = form.address()
    expect(address.options.icon).to.equal("Home")
    address = form.address({ id: null })
    expect(address.options.panel.new).to.exist
    address = form.address({ id: 1 })
    expect(address.options.panel.new).to.not.exist
  })

  it('bank', () => {
    let bank = form.bank()
    expect(bank.options.icon).to.equal("Money")
    bank = form.bank({ id: null })
    expect(bank.options.panel.new).to.exist
    bank = form.bank({ id: 1 }, { dataset: { translink: [] } })
    expect(bank.options.panel.new).to.not.exist
    bank = form.bank({ id: 1 }, { dataset: { translink: [{ transtype: "transtype" }] } })
    expect(bank.rows[0].columns[0].name).to.equal("id")
  })

  it('barcode', () => {
    let barcode = form.barcode()
    expect(barcode.options.icon).to.equal("Barcode")
    barcode = form.barcode({ id: null })
    expect(barcode.options.panel.new).to.exist
    barcode = form.barcode({ id: 1 })
    expect(barcode.options.panel.new).to.not.exist
  })

  it('cash', () => {
    let cash = form.cash()
    expect(cash.options.icon).to.equal("Money")
    cash = form.cash(
      { id: null, direction: 1 }, 
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(cash.options.panel.new).to.exist
    cash = form.cash(
      { id: 1, direction: 1 }, 
      { dataset: { translink: [], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(cash.options.panel.new).to.not.exist
    cash = form.cash(
      { id: 1, direction: 1 }, 
      { dataset: { translink: [{ transtype: "transtype" }], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(cash.rows[3].columns[0].map.source).to.equal("translink")
    cash = form.cash(
      { id: 1, direction: 1 }, 
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(cash.rows[3].columns[0].map.source).to.equal("cancel_link")
  })

  it('contact', () => {
    let contact = form.contact()
    expect(contact.options.icon).to.equal("Phone")
    contact = form.contact({ id: null })
    expect(contact.options.panel.new).to.exist
    contact = form.contact({ id: 1 })
    expect(contact.options.panel.new).to.not.exist
  })

  it('currency', () => {
    let currency = form.currency()
    expect(currency.options.icon).to.equal("Dollar")
    currency = form.currency({ id: null })
    expect(currency.options.panel.new).to.exist
    currency = form.currency({ id: 1 })
    expect(currency.options.panel.new).to.not.exist
  })

  it('customer', () => {
    let customer = form.customer()
    expect(customer.options.icon).to.equal("User")
    customer = form.customer(
      { id: null, custtype: 1 },
      { dataset: { custtype: {}, groups: [ { id: 1, groupname: "custtype", groupvalue: "own" } ] } }
    )
    expect(customer.options.panel.new).to.exist
    customer = form.customer(
      { id: 1, custtype: 2 },
      { dataset: { custtype: {}, groups: [ { id: 1, groupname: "custtype", groupvalue: "own" } ] } }
    )
    expect(customer.options.panel.new).to.not.exist
    customer = form.customer(
      { id: 1, custtype: 2 },
      { dataset: { groups: [ { id: 1, groupname: "custtype", groupvalue: "own" } ] } }
    )
    expect(customer.options.panel.new).to.not.exist
  })

  it('deffield', () => {
    let deffield = form.deffield()
    expect(deffield.options.icon).to.equal("Tag")
    deffield = form.deffield({ id: null })
    expect(deffield.options.panel.new).to.exist
    deffield = form.deffield(
      { id: 1, fieldtype: 1 },
      { dataset: { fieldtype: [ { id: 1, groupvalue: "valuelist" } ] } }
    )
    expect(deffield.options.panel.new).to.not.exist
    deffield = form.deffield(
      { id: 1, fieldtype: 1 },
      { dataset: { fieldtype: [ { id: 2, groupvalue: "valuelist" } ] } }
    )
    expect(deffield.options.panel.new).to.not.exist
  })

  it('delivery', () => {
    let delivery = form.delivery()
    expect(delivery.options.icon).to.equal("Truck")
    delivery = form.delivery(
      { id: null, direction: 1 },
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.options.panel.new).to.exist
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.rows[0].columns[1].name).to.equal("ref_transnumber")
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { translink: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.rows[0].columns[1].name).to.equal("id")
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    expect(delivery.rows[0].columns[1].name).to.equal("id")
    delivery = form.delivery(
      { id: 1, direction: 1 },
      { dataset: { cancel_link: [{ transtype: "in" }], groups: [ { id: 1, groupvalue: "transfer" } ] } }
    )
    delivery = form.delivery(
      { id: null, direction: 1 },
      { dataset: { translink: [], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(delivery.options.panel.new).to.exist
  })

  it('discount', () => {
    let discount = form.discount()
    expect(discount.options.icon).to.equal("Dollar")
    discount = form.discount({ id: null })
    expect(discount.options.panel.new).to.exist
    discount = form.discount({ id: 1 })
    expect(discount.options.panel.new).to.not.exist
  })

  it('employee', () => {
    let employee = form.employee()
    expect(employee.options.icon).to.equal("Male")
    employee = form.employee({ id: null })
    expect(employee.options.panel.new).to.exist
    employee = form.employee({ id: 1 })
    expect(employee.options.panel.new).to.not.exist
  })

  it('event', () => {
    let event = form.event()
    expect(event.options.icon).to.equal("Calendar")
    event = form.event({ id: null })
    expect(event.options.panel.new).to.exist
    event = form.event({ id: 1 })
    expect(event.options.panel.new).to.not.exist
  })

  it('formula', () => {
    let formula = form.formula()
    expect(formula.options.icon).to.equal("Magic")
    formula = form.formula({ id: null })
    expect(formula.options.panel.new).to.exist
    formula = form.formula(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(formula.options.panel.new).to.not.exist
    formula = form.formula(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(formula.rows[2].columns[0].name).to.equal("id")
  })

  it('groups', () => {
    let groups = form.groups()
    expect(groups.options.icon).to.equal("Th")
    groups = form.groups({ id: null })
    expect(groups.options.panel.new).to.exist
    groups = form.groups({ id: 1 })
    expect(groups.options.panel.new).to.not.exist
  })

  it('inventory', () => {
    let inventory = form.inventory()
    expect(inventory.options.icon).to.equal("Truck")
    inventory = form.inventory({ id: null })
    expect(inventory.options.panel.new).to.exist
    inventory = form.inventory({ id: 1 })
    expect(inventory.options.panel.new).to.not.exist
  })

  it('invoice_link', () => {
    let invoice_link = form.invoice_link()
    expect(invoice_link.options.icon).to.equal("Money")
    invoice_link = form.invoice_link({ id: null })
    expect(invoice_link.options.panel.new).to.exist
    invoice_link = form.invoice_link({ id: 1 })
    expect(invoice_link.options.panel.new).to.not.exist
  })

  it('invoice', () => {
    let invoice = form.invoice()
    expect(invoice.options.icon).to.equal("FileText")
    invoice = form.invoice({ id: null })
    expect(invoice.options.panel.new).to.exist
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal" },
      { dataset: { translink: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(invoice.options.panel.new).to.not.exist
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal", deleted: 0 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(invoice.options.panel.corrective).to.exist
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(invoice.options.panel.cancellation).to.exist
    invoice = form.invoice(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(invoice.options.panel.cancellation).to.exist
  })

  it('item', () => {
    let item = form.item()
    expect(item.options.icon).to.equal("ListOl")
    item = form.item(
      { id: null },
      { current: { transtype: "invoice" } }
    )
    expect(item.options.panel.new).to.exist
    item = form.item(
      { id: 1 },
      { current: { transtype: "offer" } }
    )
    expect(item.options.panel.new).to.not.exist
    item = form.item(
      { id: 1 },
      { current: { transtype: "order" } }
    )
    expect(item.options.panel.new).to.not.exist
  })

  it('log', () => {
    const log = form.log()
    expect(log.options.icon).to.equal("InfoCircle")
  })

  it('movement', () => {
    let movement = form.movement(
      undefined,
      { current: { transtype: "default" } }
    )
    expect(movement.options.icon).to.equal("Truck")
    movement = form.movement(
      { id: null },
      { current: { transtype: "delivery" } }
    )
    expect(movement.options.panel.new).to.exist
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "inventory" } }
    )
    expect(movement.options.panel.new).to.not.exist
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "production" } }
    )
    expect(movement.options.panel.new).to.not.exist
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "formula" } }
    )
    expect(movement.options.panel.new).to.not.exist
    movement = form.movement(
      { id: 1 },
      { current: { transtype: "waybill" } }
    )
    expect(movement.options.panel.new).to.not.exist
  })

  it('numberdef', () => {
    const numberdef = form.numberdef()
    expect(numberdef.options.icon).to.equal("ListOl")
  })

  it('offer', () => {
    let offer = form.offer()
    expect(offer.options.icon).to.equal("FileText")
    offer = form.offer({ id: null })
    expect(offer.options.panel.new).to.exist
    offer = form.offer(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(offer.options.panel.new).to.not.exist
    offer = form.offer(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(offer.options.panel.new).to.not.exist
  })

  it('order', () => {
    let order = form.order()
    expect(order.options.icon).to.equal("FileText")
    order = form.order({ id: null })
    expect(order.options.panel.new).to.exist
    order = form.order(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(order.options.panel.new).to.not.exist
    order = form.order(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(order.options.panel.new).to.not.exist
  })

  it('password', () => {
    const password = form.password()
    expect(password.options.icon).to.equal("Lock")
  })

  it('payment_link', () => {
    let payment_link = form.payment_link()
    expect(payment_link.options.icon).to.equal("FileText")
    payment_link = form.payment_link({ id: null })
    expect(payment_link.options.panel.new).to.exist
    payment_link = form.payment_link({ id: 1 })
    expect(payment_link.options.panel.new).to.equal(false)
  })

  it('payment', () => {
    let payment = form.payment()
    expect(payment.options.icon).to.equal("Money")
    payment = form.payment({ id: null })
    expect(payment.options.panel.new).to.exist
    payment = form.payment({ id: 1 })
    expect(payment.options.panel.new).to.not.exist
  })

  it('place', () => {
    let place = form.place()
    expect(place.options.icon).to.equal("Map")
    place = form.place({ id: null })
    expect(place.options.panel.new).to.exist
    place = form.place(
      { id: 1, placetype: 2 },
      { dataset: { placetype: [ { id: 1, groupvalue: "warehouse" } ] } }
    )
    expect(place.options.panel.new).to.not.exist
    place = form.place(
      { id: 1, placetype: 1 },
      { dataset: { groups: [ { id: 1, groupname: "placetype", groupvalue: "warehouse" } ] } }
    )
    expect(place.options.panel.new).to.not.exist
  })

  it('price', () => {
    let price = form.price()
    expect(price.options.icon).to.equal("Dollar")
    price = form.price({ id: null })
    expect(price.options.panel.new).to.exist
    price = form.price({ id: 1 })
    expect(price.options.panel.new).to.not.exist
  })

  it('printqueue', () => {
    const printqueue = form.printqueue(
      undefined, undefined,
      { printqueue_type: "printqueue_type", printqueue_mode: "printqueue_mode",
        report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size"
      }
    )
    expect(printqueue.options.icon).to.equal("Filter")
  })

  it('product', () => {
    let product = form.product()
    expect(product.options.icon).to.equal("ShoppingCart")
    product = form.product({ id: null })
    expect(product.options.panel.new).to.exist
    product = form.product({ id: 1 })
    expect(product.options.panel.new).to.not.exist
  })

  it('production', () => {
    let production = form.production()
    expect(production.options.icon).to.equal("Flask")
    production = form.production({ id: null })
    expect(production.options.panel.new).to.exist
    production = form.production(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(production.options.panel.new).to.not.exist
    production = form.production(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(production.options.panel.new).to.not.exist
  })

  it('project', () => {
    let project = form.project()
    expect(project.options.icon).to.equal("Clock")
    project = form.project({ id: null })
    expect(project.options.panel.new).to.exist
    project = form.project({ id: 1 })
    expect(project.options.panel.new).to.not.exist
  })

  it('rate', () => {
    let rate = form.rate()
    expect(rate.options.icon).to.equal("Strikethrough")
    rate = form.rate(
      { id: null },
      { dataset: { ratetype: [ { id: 1, groupvalue: "rate" } ], settings: [ { id: 1, fieldname: "default_currency", value: "EUR" } ] } }
    )
    expect(rate.options.panel.new).to.exist
    rate = form.rate(
      { id: null },
      { dataset: { ratetype: [ { id: 1, groupvalue: "rate" } ], settings: [] } }
    )
    expect(rate.options.panel.new).to.exist
    rate = form.rate({ id: 1 })
    expect(rate.options.panel.new).to.not.exist
  })

  it('program', () => {
    const program = form.program()
    expect(program.options.icon).to.equal("Keyboard")
  })

  it('receipt', () => {
    let receipt = form.receipt()
    expect(receipt.options.icon).to.equal("FileText")
    receipt = form.receipt({ id: null })
    expect(receipt.options.panel.new).to.exist
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal" },
      { dataset: { translink: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "in" } ] } }
    )
    expect(receipt.options.panel.new).to.not.exist
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal", deleted: 0 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(receipt.options.panel.corrective).to.exist
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [{ transtype: "transtype" }], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(receipt.options.panel.cancellation).to.exist
    receipt = form.receipt(
      { id: 1, direction: 1, transcast: "normal", deleted: 1 },
      { dataset: { translink: [], cancel_link: [], groups: [ { id: 1, groupvalue: "out" } ] } }
    )
    expect(receipt.options.panel.cancellation).to.exist
  })

  it('rent', () => {
    let rent = form.rent()
    expect(rent.options.icon).to.equal("FileText")
    rent = form.rent({ id: null })
    expect(rent.options.panel.new).to.exist
    rent = form.rent(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(rent.options.panel.new).to.not.exist
    rent = form.rent(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(rent.options.panel.new).to.not.exist
  })

  it('report', () => {
    let report = form.report(
      undefined, undefined,
      { report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size" }
    )
    expect(report.options.panel.print).to.exist
    report = form.report(
      { id: 1, ftype: "csv" }, undefined,
      { report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size" }
    )
    expect(report.options.panel.print).to.exist
    report = form.report(
      { id: 1, ftype: "pdf" }, undefined,
      { report_orientation: "report_orientation", page_orient: "page_orient",
        report_size: "report_size", page_size: "page_size" }
    )
    expect(report.options.panel.print).to.exist
  })

  it('setting', () => {
    const setting = form.setting()
    expect(setting.options.icon).to.equal("Cog")
  })

  it('shipping', () => {
    const shipping = form.shipping()
    expect(shipping.options.icon).to.equal("Truck")
  })

  it('tax', () => {
    let tax = form.tax()
    expect(tax.options.icon).to.equal("Ticket")
    tax = form.tax({ id: null })
    expect(tax.options.panel.new).to.exist
    tax = form.tax({ id: 1 })
    expect(tax.options.panel.new).to.not.exist
  })

  it('template', () => {
    const template = form.template()
    expect(template.options.icon).to.equal("TextHeight")
  })

  it('tool', () => {
    let tool = form.tool()
    expect(tool.options.icon).to.equal("Wrench")
    tool = form.tool({ id: null })
    expect(tool.options.panel.new).to.exist
    tool = form.tool({ id: 1 })
    expect(tool.options.panel.new).to.not.exist
  })

  it('ui_menu', () => {
    let ui_menu = form.ui_menu()
    expect(ui_menu.options.icon).to.equal("Share")
    ui_menu = form.ui_menu({ id: null })
    expect(ui_menu.rows[0].columns[0].disabled).to.not.exist
    ui_menu = form.ui_menu({ id: 1 })
    expect(ui_menu.rows[0].columns[0].disabled).to.exist
  })

  it('usergroup', () => {
    let usergroup = form.usergroup()
    expect(usergroup.options.icon).to.equal("Key")
    usergroup = form.usergroup({ id: null })
    expect(usergroup.options.panel.new).to.exist
    usergroup = form.usergroup({ id: 1 })
    expect(usergroup.options.panel.new).to.not.exist
  })

  it('waybill', () => {
    let waybill = form.waybill()
    expect(waybill.options.icon).to.equal("Briefcase")
    waybill = form.waybill({ id: null })
    expect(waybill.options.panel.new).to.exist
    waybill = form.waybill(
      { id: 1, customer_id: null, employee_id: null },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(waybill.options.panel.new).to.not.exist
    waybill = form.waybill(
      { id: 1, customer_id: null, employee_id: null },
      { dataset: { translink: [] } }
    )
    expect(waybill.options.panel.new).to.not.exist
    waybill = form.waybill(
      { id: 1, customer_id: 1 },
    )
    expect(waybill.rows[1].columns[1].map.seltype).to.equal("customer")
    waybill = form.waybill(
      { id: 1, customer_id: null, employee_id: 1 },
    )
    expect(waybill.rows[1].columns[1].map.seltype).to.equal("employee")
  })

  it('worksheet', () => {
    let worksheet = form.worksheet()
    expect(worksheet.options.icon).to.equal("FileText")
    worksheet = form.worksheet({ id: null })
    expect(worksheet.options.panel.new).to.exist
    worksheet = form.worksheet(
      { id: 1 },
      { dataset: { translink: [{ transtype: "transtype" }] } }
    )
    expect(worksheet.options.panel.new).to.not.exist
    worksheet = form.worksheet(
      { id: 1 },
      { dataset: { translink: [] } }
    )
    expect(worksheet.options.panel.new).to.not.exist
  })

})