import { Sql } from './Sql'
import { getText, store } from 'config/app';

const sql = Sql({ getText: (key)=>getText({ locales: store.session.locales, lang: "en", key: key }) })

it('Sql', () => {
  expect(sql.all.deffield_prop("nervatype")).toBeDefined()
  
  expect(sql.currency.delete_state()).toBeDefined()
  expect(sql.currency.currency_view()).toBeDefined()

  expect(sql.customer.delete_state()).toBeDefined()
  expect(sql.customer.address()).toBeDefined()
  expect(sql.customer.contact()).toBeDefined()
  expect(sql.customer.event()).toBeDefined()

  expect(sql.deffield.delete_state()).toBeDefined()
  expect(sql.deffield.deffield_view()).toBeDefined()

  expect(sql.employee.delete_state()).toBeDefined()
  expect(sql.employee.employee()).toBeDefined()
  expect(sql.employee.address()).toBeDefined()
  expect(sql.employee.contact()).toBeDefined()
  expect(sql.employee.event()).toBeDefined()

  expect(sql.event.event()).toBeDefined()

  expect(sql.groups.delete_state()).toBeDefined()
  expect(sql.groups.groups_view()).toBeDefined()

  expect(sql.log.result()).toBeDefined()

  expect(sql.numberdef.numberdef_view()).toBeDefined()

  expect(sql.place.delete_state()).toBeDefined()
  expect(sql.place.contact()).toBeDefined()
  expect(sql.place.place_view()).toBeDefined()

  expect(sql.printqueue.server_printers()).toBeDefined()
  expect(sql.printqueue.items("")).toBeDefined()
  expect(sql.printqueue.items({})).toBeDefined()
  expect(sql.printqueue.items({ nervatype: "", startdate: "", enddate: "", transnumber: "", username: "" })).toBeDefined()
  expect(sql.printqueue.items({ nervatype: "customer" })).toBeDefined()
  expect(sql.printqueue.items({ nervatype: "product" })).toBeDefined()
  expect(sql.printqueue.items({ nervatype: "employee" })).toBeDefined()
  expect(sql.printqueue.items({ nervatype: "tool" })).toBeDefined()
  expect(sql.printqueue.items({ nervatype: "project" })).toBeDefined()

  expect(sql.product.delete_state()).toBeDefined()
  expect(sql.product.barcode()).toBeDefined()
  expect(sql.product.barcode_check()).toBeDefined()
  expect(sql.product.price()).toBeDefined()
  expect(sql.product.discount()).toBeDefined()
  expect(sql.product.event()).toBeDefined()

  expect(sql.project.delete_state()).toBeDefined()
  expect(sql.project.project()).toBeDefined()
  expect(sql.project.address()).toBeDefined()
  expect(sql.project.contact()).toBeDefined()
  expect(sql.project.event()).toBeDefined()

  expect(sql.rate.rate()).toBeDefined()
  expect(sql.rate.delete_state()).toBeDefined()

  expect(sql.report.report("report")).toBeDefined()
  expect(sql.report.report("printqueue")).toBeDefined()
  expect(sql.report.report("default")).toBeDefined()

  expect(sql.setting.setting_view()).toBeDefined()

  expect(sql.tax.tax_view()).toBeDefined()
  expect(sql.tax.delete_state()).toBeDefined()

  expect(sql.template.template_view()).toBeDefined()
  expect(sql.template.template()).toBeDefined()

  expect(sql.tool.delete_state()).toBeDefined()
  expect(sql.tool.tool()).toBeDefined()
  expect(sql.tool.event()).toBeDefined()

  expect(sql.trans.delete_state()).toBeDefined()
  expect(sql.trans.item()).toBeDefined()
  expect(sql.trans.element_count()).toBeDefined()
  expect(sql.trans.payment()).toBeDefined()
  expect(sql.trans.movement_delivery()).toBeDefined()
  expect(sql.trans.movement_transfer()).toBeDefined()
  expect(sql.trans.movement_inventory()).toBeDefined()
  expect(sql.trans.movement_waybill()).toBeDefined()
  expect(sql.trans.movement_formula_head()).toBeDefined()
  expect(sql.trans.movement_formula()).toBeDefined()
  expect(sql.trans.movement_production_head()).toBeDefined()
  expect(sql.trans.movement_production()).toBeDefined()
  expect(sql.trans.formula_head()).toBeDefined()
  expect(sql.trans.formula_items()).toBeDefined()
  expect(sql.trans.trans()).toBeDefined()
  expect(sql.trans.translink()).toBeDefined()
  expect(sql.trans.cancel_link()).toBeDefined()
  expect(sql.trans.invoice_customer()).toBeDefined()
  expect(sql.trans.invoice_link()).toBeDefined()
  expect(sql.trans.payment_link()).toBeDefined()
  expect(sql.trans.tool_movement()).toBeDefined()
  expect(sql.trans.transitem_invoice()).toBeDefined()
  expect(sql.trans.transitem_shipping()).toBeDefined()
  expect(sql.trans.shipping_items()).toBeDefined()
  expect(sql.trans.shipping_delivery()).toBeDefined()
  expect(sql.trans.shipping_stock()).toBeDefined()

  expect(sql.ui_menu.ui_menufields()).toBeDefined()
  expect(sql.ui_menu.ui_menu_view()).toBeDefined()

  expect(sql.usergroup.delete_state()).toBeDefined()
  expect(sql.usergroup.reportkey()).toBeDefined()
  expect(sql.usergroup.menukey()).toBeDefined()
  expect(sql.usergroup.datafilter()).toBeDefined()
  expect(sql.usergroup.audit()).toBeDefined()
  expect(sql.usergroup.usergroup_view()).toBeDefined()

})