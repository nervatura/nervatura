import { expect } from '@open-wc/testing';

import { Sql } from './Sql.js'

it('Sql', () => {
  const sql = Sql({ msg: (key)=> key })

  expect(sql.all.deffield_prop("nervatype")).to.exist
  
  expect(sql.currency.delete_state()).to.exist
  expect(sql.currency.currency_view()).to.exist

  expect(sql.customer.delete_state()).to.exist
  expect(sql.customer.address()).to.exist
  expect(sql.customer.contact()).to.exist
  expect(sql.customer.event()).to.exist

  expect(sql.deffield.delete_state()).to.exist
  expect(sql.deffield.deffield_view()).to.exist

  expect(sql.employee.delete_state()).to.exist
  expect(sql.employee.employee()).to.exist
  expect(sql.employee.address()).to.exist
  expect(sql.employee.contact()).to.exist
  expect(sql.employee.event()).to.exist

  expect(sql.event.event()).to.exist

  expect(sql.groups.delete_state()).to.exist
  expect(sql.groups.groups_view()).to.exist

  expect(sql.log.result()).to.exist

  expect(sql.numberdef.numberdef_view()).to.exist

  expect(sql.place.delete_state()).to.exist
  expect(sql.place.contact()).to.exist
  expect(sql.place.place_view()).to.exist

  expect(sql.printqueue.server_printers()).to.exist
  expect(sql.printqueue.items("")).to.exist
  expect(sql.printqueue.items({})).to.exist
  expect(sql.printqueue.items({ nervatype: "", startdate: "", enddate: "", transnumber: "", username: "" })).to.exist
  expect(sql.printqueue.items({ nervatype: "customer" })).to.exist
  expect(sql.printqueue.items({ nervatype: "product" })).to.exist
  expect(sql.printqueue.items({ nervatype: "employee" })).to.exist
  expect(sql.printqueue.items({ nervatype: "tool" })).to.exist
  expect(sql.printqueue.items({ nervatype: "project" })).to.exist

  expect(sql.product.delete_state()).to.exist
  expect(sql.product.barcode()).to.exist
  expect(sql.product.barcode_check()).to.exist
  expect(sql.product.price()).to.exist
  expect(sql.product.discount()).to.exist
  expect(sql.product.event()).to.exist

  expect(sql.project.delete_state()).to.exist
  expect(sql.project.project()).to.exist
  expect(sql.project.address()).to.exist
  expect(sql.project.contact()).to.exist
  expect(sql.project.event()).to.exist

  expect(sql.rate.rate()).to.exist
  expect(sql.rate.delete_state()).to.exist

  expect(sql.report.report("report")).to.exist
  expect(sql.report.report("printqueue")).to.exist
  expect(sql.report.report("default")).to.exist

  expect(sql.setting.setting_view()).to.exist

  expect(sql.tax.tax_view()).to.exist
  expect(sql.tax.delete_state()).to.exist

  expect(sql.template.template_view()).to.exist
  expect(sql.template.template()).to.exist

  expect(sql.tool.delete_state()).to.exist
  expect(sql.tool.tool()).to.exist
  expect(sql.tool.event()).to.exist

  expect(sql.trans.delete_state()).to.exist
  expect(sql.trans.item()).to.exist
  expect(sql.trans.element_count()).to.exist
  expect(sql.trans.payment()).to.exist
  expect(sql.trans.movement_delivery()).to.exist
  expect(sql.trans.movement_transfer()).to.exist
  expect(sql.trans.movement_inventory()).to.exist
  expect(sql.trans.movement_waybill()).to.exist
  expect(sql.trans.movement_formula_head()).to.exist
  expect(sql.trans.movement_formula()).to.exist
  expect(sql.trans.movement_production_head()).to.exist
  expect(sql.trans.movement_production()).to.exist
  expect(sql.trans.formula_head()).to.exist
  expect(sql.trans.formula_items()).to.exist
  expect(sql.trans.trans()).to.exist
  expect(sql.trans.translink()).to.exist
  expect(sql.trans.cancel_link()).to.exist
  expect(sql.trans.invoice_customer()).to.exist
  expect(sql.trans.invoice_link()).to.exist
  expect(sql.trans.payment_link()).to.exist
  expect(sql.trans.tool_movement()).to.exist
  expect(sql.trans.transitem_invoice()).to.exist
  expect(sql.trans.transitem_shipping()).to.exist
  expect(sql.trans.shipping_items()).to.exist
  expect(sql.trans.shipping_delivery()).to.exist
  expect(sql.trans.shipping_stock()).to.exist

  expect(sql.ui_menu.ui_menufields()).to.exist
  expect(sql.ui_menu.ui_menu_view()).to.exist

  expect(sql.usergroup.delete_state()).to.exist
  expect(sql.usergroup.reportkey()).to.exist
  expect(sql.usergroup.menukey()).to.exist
  expect(sql.usergroup.datafilter()).to.exist
  expect(sql.usergroup.audit()).to.exist
  expect(sql.usergroup.usergroup_view()).to.exist

})