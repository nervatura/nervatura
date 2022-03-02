import update from 'immutability-helper';
import { getSetting } from 'config/app'

export const Forms = ({ getText }) => {
  return {
    address: (item) => {
      let address = {
        options: {
          title: getText("address_view"),
          title_field: "",
          icon: "Home",
          panel: {}
        },
        rows: [
          {rowtype: "col3", columns: [
            {name: "country", label: getText("address_country"), datatype: "string"},
            {name: "state", label: getText("address_state"), datatype: "string"},
            {name: "zipcode", label: getText("address_zipcode"), datatype: "string"}]},
          {rowtype: "field", name: "city", label: getText("address_city"), datatype: "string"},
          {rowtype: "field", name: "street", label: getText("ddress_street"), datatype: "string"},
          {rowtype: "field", name:"notes", label: getText("address_notes"), datatype: "text"}]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          address = update(address, { options: { 
            panel: {$merge: {
              new: false, delete: false
            }} 
          }})
        }
      }
      return address;
    },

    bank: (item, edit) => {
      let bank = {
        options: {
          title: getText("title_bank"),
          title_field: "transnumber",
          icon: "Money",
          fieldvalue: true,
          pattern: true,
          panel: {arrow:true, more:true, trans:true, create:false,
            bookmark:["editor","trans","transnumber"], help:"document/payment"}},
        view: {
          payment: {
            type: "table",
            icon: "ListOl",
            title: getText("item_view"),
            data: "payment",
            total:{
              expense: getText("payment_expense"),
              income: getText("payment_income"),
              balance: getText("payment_balance")
            },
            fields: {
              rid: {fieldtype:'number', label: getText("payment_item")},
              paiddate: {fieldtype:'date', label: getText("payment_paiddate2")},
              amount: {fieldtype:'number', label: getText("payment_amount")},
              notes: {fieldtype:'string', label: getText("payment_description")}}
          },
          payment_link: {
            type: "list",
            data:"payment_link",
            icon: "FileText",
            title:getText("invoice_view"),
            actions: {
              new: null, 
              edit: {action: "editEditorItem", fkey: "payment_link"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"ref_transnumber", label:getText("document_ref_transnumber"), datatype:"string"},
            {name:"crdate", label:getText("bank_crdate"), datatype:"date", disabled: false},
            {name:"transtate", label:getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"transdate", label:getText("bank_transdate"), datatype:"date"},
            {name:"place_id", label:getText("payment_place_bank"), datatype:"selector",
              empty: false, map:{seltype:"place_bank", table:"trans", fieldname:"place_id", 
              lnktype:"place", transtype:"", label_field:"planumber"}},
            {name:"closed", label:getText("document_closed"), datatype:"flip"}]},
          {rowtype:"field", name:"notes", label:getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label:getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          bank = update(bank, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false, 
                report: false, bookmark: false, trans: false
              }} 
            }
          })  
        } else {
          if (edit.dataset.translink.length > 0) {
            bank = update(bank, { rows: { 0: { columns: { 0: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: { source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return bank;
    },

    barcode: (item) => {
      let barcode = {
        options: {
          title: getText("barcode_view"),
          title_field: "",
          icon: "Barcode",
          panel: {}},
        rows: [
          {rowtype:"field", name:"code", label: getText("barcode_code"), datatype:"string"},
          {rowtype:"col3", columns: [
            {name:"barcodetype", label: getText("barcode_barcodetype"), datatype:"select", 
              map: {source:"barcodetype", value:"id", text:"description" }},
            {name:"qty", label: getText("barcode_qty"), datatype:"float"},
            {name:"defcode", label: getText("barcode_defcode"), datatype:"flip"}]},
          {rowtype:"field", name:"description", label:getText("barcode_description"), datatype:"text"}]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          barcode = update(barcode, { options: { 
            panel: {$merge: {
              new: false, delete: false
            }} 
          }})
        }
      }
      return barcode;
    },

    cash: (item, edit) => {
      let cash = {
        options: {
          title: getText("title_cash"),
          title_field: "transnumber",
          icon: "Money",
          fieldvalue:true,
          pattern:true,
          extend: "payment",
          panel: {arrow:true, more:true, trans:true, create:false,
            cancellation:true, bookmark:["editor","trans","transnumber"], help:"document/payment",
            link: true, link_type:"payment_link", link_field:"ref_id_1",
            link_label: getText("label_link_invoice")}},
        view: {
          payment_link: {
            type: "list",
            data:"payment_link",
            icon: "FileText",
            title: getText("invoice_view"),
            actions: {
              new: null, 
              edit: {action: "editEditorItem", fkey: "payment_link"}, 
              delete: null}
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("invoice_crdate"), datatype:"date", disabled: true},
            {name:"closed", label: getText("document_closed"), datatype:"flip"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("document_direction"), datatype:"select", empty: false,
              map: {source:"direction", value:"id", text:"groupvalue", label:"cash" }},
            {name:"id", label: getText("payment_paiddate"), datatype:"date",
              map: {source:"payment", value:"trans_id", text:"paiddate", extend:true}},
            {name:"id", label: getText("payment_amount"), datatype:"float", opposite:true,
              map: {source:"payment", value:"trans_id", text:"amount", extend:true}}]},
          {rowtype:"col2", columns: [
            {name:"place_id", label: getText("payment_place_cash"), datatype:"selector",
              empty: false, map:{seltype:"place_cash", table:"trans", fieldname:"place_id", 
              lnktype:"place", transtype:"", label_field:"planumber"}},
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}}]},
          {rowtype:"col2", columns: [
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"intnotes", label: getText("document_intnotes"), datatype:"text"}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        const direction = edit.dataset.groups.filter((group)=> {
          return (group.id === item.direction)
        })[0].groupvalue
        if(direction === "out"){
          cash = update(cash, {
            options: {$merge: {
              opposite: true
            }}
          }) 
        }
        if (item.id === null) {
          cash = update(cash, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false, 
                report: false, bookmark: false, trans: false,
                link: false
              }} 
            }
          })
        } else {
          cash = update(cash, { rows: { 1: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          if (edit.dataset.translink.length > 0) {
            cash = update(cash, { rows: { 3: { columns: { 0: {$set: {
              name:"id", 
              label:getText("document_ref_transnumber"), 
              datatype:"link",
              map: {source: "translink", value: "ref_id_1", text: "ref_id_2",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          } else if (edit.dataset.cancel_link.length > 0) {
            cash = update(cash, { rows: { 3: { columns: { 0: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype: "link",
              map: { source:"cancel_link", value: "ref_id_2", text: "ref_id_1",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.cancel_link[0].transtype
              }
            }}}}}})
          }
        }
      }
      return cash;
    },

    contact: (item) => {
      let contact = {
        options: {
          title: getText("contact_view"),
          title_field: "",
          icon: "Phone",
          panel: {}},
        rows: [
          {rowtype:"col2", columns: [
            {name:"firstname", label:getText("contact_firstname"), datatype:"string"},
            {name:"surname", label:getText("contact_surname"), datatype:"string"}]},
          {rowtype:"col2", columns: [
            {name:"status", label:getText("contact_status"), datatype:"string"},
            {name:"phone", label:getText("contact_phone"), datatype:"string"}]},
          {rowtype:"col2", columns: [
            {name:"mobil", label:getText("contact_mobil"), datatype:"string"},
            {name:"fax", label:getText("contact_fax"), datatype:"string"}]},
          {rowtype:"field", name:"email", label:getText("contact_email"), datatype:"string"},
          {rowtype:"field", name:"notes", label:getText("contact_notes"), datatype:"text"}]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          contact = update(contact, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return contact;
    },
    
    currency: (item) => {
      let currency = {
        options: {
          icon: "Dollar",
          data: "currency",
          title: getText("title_currency"),
          panel: {page:"setting", more:false, help:"settings/currency"}},
        view: {
          setting: {
            type:"table",
            actions: {
              new: {action: "newItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem"}},
            fields: {
                curr: {fieldtype:'string', label:getText("currency_curr")},
                description: {fieldtype:'string', label:getText("currency_description")},
                digit: {fieldtype:'number', label:getText("currency_digit")},
                cround: {fieldtype:'number', label:getText("currency_cround")},
                defrate: {fieldtype:'number', label:getText("currency_defrate")}}}
        },
        rows: [
          {rowtype:"col2", columns: [
            {name:"curr", label:getText("currency_curr"), datatype:"string"},
            {name:"description", label:getText("currency_description"), datatype:"string"}]},
          {rowtype:"col3", columns: [
            {name:"digit", label:getText("currency_digit"), datatype:"integer"},
            {name:"cround", label:getText("currency_cround"), datatype:"integer"},
            {name:"defrate", label:getText("currency_defrate"), datatype:"float"}]}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id !== null) {
          currency = update(currency, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
        } else {
          currency = update(currency, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return currency;
    },
    
    customer: (item, edit) => { 
      let customer = {
        options: {
          title: getText("title_customer"),
          title_field: "custnumber",
          icon: "User",
          fieldvalue:true,
          panel: {more:true, bookmark:["editor","customer","custname","custnumber"], help:"resources/customer"}},
        view: {
          address: {
            type: "list",
            data: "address",
            icon: "Home",
            title: getText("address_view")},
          contact: {
            type: "list",
            data: "contact",
            icon: "Phone",
            title: getText("contact_view")},
          event: {
            type: "list",
            data: "event",
            icon: "Calendar",
            title: getText("event_view"),
            actions: {
              new: {action: "loadEditor", ntype: "event", ttype: null}, 
              edit: {action: "loadEditor", ntype: "event", ttype: null}, 
              delete: {action: "deleteEditorItem", fkey: "event", table: "event"}
            }
          }
      },
      rows: [
        {rowtype:"field", name:"custname", 
          label:getText("customer_custname"), datatype:"string"},
        {rowtype:"col3", columns: [
          {name:"custnumber", label:getText("customer_custnumber"), datatype:"string"},
          {name:"taxnumber", label:getText("customer_taxnumber"), datatype:"string"},
          {name:"account", label:getText("customer_account"), datatype:"string"}]}
        ]
      };
      if (typeof item !== "undefined") {
        if (typeof edit.dataset.custtype !== "undefined") {
          if (item.custtype === edit.dataset.groups.filter((group)=> {
              return ((group.groupname === "custtype") && (group.groupvalue ==="own"))
            })[0].id) {
            customer = update(customer, {
              options: { $merge: {
                title: getText("title_company"),
                icon: "Home"
              }}
            })
            customer = update(customer, {
              options: { 
                panel: {$merge: {
                  new: false, delete: false
                }} 
              }
            })
          } else {
            customer = update(customer, { rows: { $push: [
              { rowtype:"col3", columns: [
                {name:"creditlimit", label:getText("customer_creditlimit"), datatype:"float"},
                {name:"terms", label:getText("customer_terms"), datatype:"integer"},
                {name:"discount", label:getText("customer_discount"), datatype:"float", min:0, max:100}
              ]},
              {rowtype:"col3", columns: [
                {name:"custtype", label:getText("customer_custtype"), datatype:"select", 
                  map: {source:"custtype", value:"id", text:"groupvalue" }},
                {name:"inactive", label:getText("customer_inactive"), datatype:"flip"},
                {name:"notax", label:getText("customer_notax"), datatype:"flip"}
              ]}
            ]}})
          }
          customer = update(customer, { rows: { $push: [
            { rowtype:"field", name:"notes", 
              label:getText("customer_notes"), datatype:"text" }
          ]}})
        }
        if (item.id === null) {
          customer = update(customer, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, 
                report: false, bookmark: false
              }} 
            }
          })
        } else {
          customer = update(customer, { rows: { 1: {$merge: {
            columns: customer.rows[1].columns.slice(1,customer.rows[1].columns.length),
            rowtype: "col2"
          }}}})
        }
      }
      return customer;
    },

    deffield: (item, setting) => {
      let deffield = {
        options: {
          icon: "Tag",
          data: "deffield",
          title: getText("title_deffield"),
          panel: {page:"setting", more:false, help:"settings/metadata"}},
        view: {
          setting: {
            type: "list",
            actions: {
              new: {action: "newItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem"}}}
        },
        rows: [
          {rowtype:"field", name:"fieldname", label: getText("deffield_fieldname"), 
            datatype:"string", disabled: true},
          {rowtype:"col2", columns: [
            {name:"nervatype", label: getText("deffield_nervatype"), datatype:"select", empty: true,
              map: {source:"nervatype", value:"id", text:"groupvalue" }},
            {name:"fieldtype", label: getText("deffield_fieldtype"), datatype:"select", empty: true,
              map: {source:"fieldtype", value:"id", text:"groupvalue" }}]},
          {rowtype:"field", name:"description", label: getText("deffield_description"), datatype:"string"},
          {rowtype:"col3", columns: [
            {name:"addnew", label: getText("deffield_addnew"), datatype:"flip"},
            {name:"visible", label: getText("deffield_visible"), datatype:"flip"},
            {name:"readonly", label: getText("deffield_readonly"), datatype:"flip"}]},
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id !== null) {
          deffield = update(deffield, { rows: { 1: { columns: { 
            0: {$merge: {
              disabled: true
            }},
            1: {$merge: {
              disabled: true
            }}
          }}}})
          if (item.fieldtype === setting.dataset.fieldtype.filter((group)=> {
              return ((group.groupvalue === "valuelist"))
            })[0].id) {
              deffield = update(deffield, { rows: { $push: [
                { rowtype:"field", name:"valuelist", 
                  label: getText("deffield_valuelist"), 
                  datatype:"text" }
               ]}})
            }
        } else {
          deffield = update(deffield, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return deffield;
    },

    delivery: (item, edit) => {
      let delivery = {
        options: {
          title: getText("title_delivery"),
          title_field: "transnumber",
          icon: "Truck",
          fieldvalue: true,
          pattern: true,
          panel: {
            arrow:true, more:true, trans:true, create:false, copy:false, 
            cancellation:true, delete:false, new:false,
            bookmark: ["editor","trans","transnumber"], help:"stock/delivery"
          }
        },
        view: {
          movement: {
            type: "table",
            icon: "ListOl",
            title: getText("item_view"),
            data: "movement",
            edited: false,
            fields: {
              product: {fieldtype:'string', label: getText("product_description")},
              unit: {fieldtype:'string', label: getText("product_unit")},
              notes: {fieldtype:'string', label: getText("movement_batchnumber")},
              qty: {fieldtype:'number', label: getText("movement_qty")}}
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("delivery_direction"), datatype:"string",
              map: {source:"groups", value:"id", text:"groupvalue", label:"delivery" }},
            {name:"id", label: getText("document_ref_transnumber"), datatype:"link",
              map: {source:"movement", value:"trans_id", text:"item_ref_id",
                label_field:"item_refnumber", lnktype:"trans", transtype:"order"}},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("delivery_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("delivery_transdate"), datatype:"date", disabled: true},
            {name:"id", label: getText("delivery_place"), datatype:"string",
              map: {source:"movement", value:"trans_id", text:"planumber" }}]}
        ]};
      if (typeof item !== "undefined") {
        const direction = edit.dataset.groups.filter((group)=> {
          return (group.id === item.direction)
        })[0].groupvalue
        if (direction === "transfer") {
          if (edit.dataset.translink && (edit.dataset.translink.length > 0)) {
            delivery = update(delivery, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype: "link",
              map: {
                source: "translink", value: "ref_id_1", text: "ref_id_2",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          } else if (edit.dataset.cancel_link && (edit.dataset.cancel_link.length > 0)) {
            delivery = update(delivery, { rows: { 0: { columns: { 1: {$set: {
              name: "id", 
              label: getText("document_ref_transnumber"), datatype: "link",
              map: {
                source: "cancel_link", value: "ref_id_2", text: "ref_id_1",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.cancel_link[0].transtype
              }
            }}}}}})
          } else{
            delivery = update(delivery, { rows: { 0: { columns: { 1: {$set: {
              name:"ref_transnumber", 
              label: getText("document_ref_transnumber"), 
              datatype:"string"
            }}}}}})
          }
          delivery = update(delivery, { rows: { 1: { columns: { 
            1: {$merge: {
              disabled:false
            }},
            2: {$set: {
              name:"closed", 
              label: getText("document_closed"), 
              datatype:"flip"
            }}
          }}}})
          delivery = update(delivery, { rows: { $push: [
            {rowtype:"col2", columns: [
              {name:"place_id", label: getText("delivery_place"), 
                datatype:"selector", empty: false, 
                map:{seltype:"place_warehouse", table:"trans", fieldname:"place_id", 
                lnktype:"place", transtype:"", label_field:"planumber"}},
              {name:"target_place", label: getText("movement_target"), 
                datatype:"selector", empty: false, disabled: true,
                map:{seltype:"place_warehouse", table:"trans", fieldname:"target_place", 
                lnktype:"place", transtype:"", label_field:"target_planumber"}}
            ]}
          ]}})
          if (item.id === null) {
            delivery = update(delivery, {
              view: {$set: {}},
              options: { 
                panel: {$merge: {
                  arrow: false, new: false, delete: false,
                  report: false, bookmark: false, trans: false
                }} 
              }
            })
          } else {
            delivery = update(delivery, {
              options: { 
                panel: {$merge: {
                  copy: true, new: true
                }} 
              },
              view: {
                movement: {$merge: {
                  edit: true,
                  data: "movement_transfer"
                }}
              },
              rows: {
                2: {
                  columns: {
                    0: {$merge: {
                      disabled: true
                    }}
                  }
                }
              }
            })
          }
        }
        delivery = update(delivery, { rows: { $push: [
          { rowtype:"field", name:"notes", 
            label: getText("document_notes"), datatype:"text" },
          { rowtype:"field", name:"intnotes", 
            label: getText("document_intnotes"), datatype:"text" }
        ]}})
      }
      return delivery;
    },

    discount: (item) => {
      let discount = {
        options: {
          title: getText("discount_view"),
          title_field: "",
          icon: "Dollar",
          panel: {}
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"validfrom", label: getText("price_validfrom"), datatype: "date"},
            {name:"validto", label: getText("price_validto"), datatype:"date", empty: true},
            {name:"vendor", label: getText("price_vendor"), datatype: "flip"}]},
          {rowtype:"col3", columns: [
            {name:"curr", label: getText("price_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"qty", label: getText("price_qty"), datatype:"float"},
            {name:"pricevalue", label: getText("price_limit"), datatype:"float"}]},
          {rowtype:"col2", columns: [
            {name:"calcmode", label: getText("price_calcmode"), datatype:"select",
              map: {source:"calcmode", value:"id", text:"description"}},
            {name:"discount", label: getText("price_discount"), datatype:"float"}]},
          {rowtype:"field", name:"id", label: getText("customer_custname"), datatype:"selector",
            empty: true, map:{seltype:"customer", table:"discount", fieldname:"customer_id", 
            lnktype:"customer", transtype:"", label_field:"custname"}}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          discount = update(discount, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return discount;
    },

    employee: (item) => { 
      let employee = {
        options: {
          title: getText("title_employee"),
          title_field: "empnumber",
          icon: "Male",
          extend: "contact",
          fieldvalue: true,
          panel: { 
            more:true, password:true,
            bookmark:["editor","employee","empnumber","empnumber"], 
            help:"resources/employee"}
          },
        view: {
          address: {
            type: "list",
            data: "address",
            icon: "Home",
            title: getText("address_view")
          },
          event: {
            type: "list",
            data: "event",
            icon: "Calendar",
            title: getText("event_view"),
            actions: {
              new: {action: "loadEditor", ntype: "event", ttype: null}, 
              edit: {action: "loadEditor", ntype: "event", ttype: null}, 
              delete: {action: "deleteEditorItem", fkey: "event", table: "event"}
            }
          }
      },
      rows: [
        {rowtype:"col3", columns: [
          {name:"empnumber", label: getText("employee_empnumber"), datatype:"string"},
          {name:"id", label: getText("contact_firstname"), datatype:"string",
            map: {source:"contact", value:"ref_id", text:"firstname", extend:true}},
          {name:"id", label: getText("contact_surname"), datatype:"string",
            map: {source:"contact", value:"ref_id", text:"surname", extend:true}}]},
        {rowtype:"col3", columns: [
          {name:"id", label: getText("contact_status"), datatype:"string",
            map: {source:"contact", value:"ref_id", text:"status", extend:true}},
          {name:"id", label: getText("contact_phone"), datatype:"string",
            map: {source:"contact", value:"ref_id", text:"phone", extend:true}},
          {name:"id", label: getText("contact_mobil"), datatype:"string",
            map: {source:"contact", value:"ref_id", text:"mobil", extend:true}}]},
        {rowtype:"field", name:"id", label: getText("contact_email"), datatype:"string",
          map: {source:"contact", value:"ref_id", text:"email", extend:true}},
        {rowtype:"col3", columns: [
          {name:"startdate", label: getText("employee_startdate"), datatype:"date", empty: true},
          {name:"enddate", label: getText("employee_enddate"), datatype:"date", empty: true},
          {name:"department", label: getText("employee_department"), datatype:"select", empty: true,
            map: {source:"department", value:"id", text:"groupvalue"}}]},
        {rowtype:"col3", columns: [
          {name:"usergroup", label: getText("employee_usergroup"), datatype:"select", empty: false,
            map: {source:"usergroup", value:"id", text:"groupvalue"}},
          {name:"username", label: getText("employee_username"), datatype:"string"},
          {name:"inactive", label: getText("employee_inactive"), datatype:"flip"}]},
        {rowtype:"field", name:"id", label: getText("employee_notes"), datatype:"text",
          map: {source:"contact", value:"ref_id", text:"notes", extend:true}}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          employee = update(employee, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, 
                report: false, bookmark: false, password: false
              }} 
            }
          })
        } else {
          employee = update(employee, { rows: { 0: {$merge: {
            columns: employee.rows[0].columns.slice(1,employee.rows[0].columns.length),
            rowtype: "col2"
          }}}})
        }
      }
      return employee;
    },

    event: (item) => { 
      let event = {
        options: {
          title: getText("title_event"),
          title_field: "calnumber",
          icon: "Calendar",
          fieldvalue: true,
          panel: {
            back: true, more: true, 
            bookmark:["editor","event","calnumber","calnumber"], 
            help: "resources/event", 
            export_event: true, report: false
          }
        },
        view: {},
        rows: [
          {rowtype:"field", name:"subject", label: getText("event_subject"), datatype:"string"},
          {rowtype:"col2", columns: [
            {name:"place", label: getText("event_place"), datatype:"string"},
            {name:"eventgroup", label: getText("event_group"), datatype:"select", empty: true,
              map: {source:"eventgroup", value:"id", text:"groupvalue" }}]},
          {rowtype:"col2", columns: [
            {name:"fromdate", label: getText("event_fromdate"), datatype:"datetime", empty: true},
            {name:"todate", label: getText("event_todate"), datatype:"datetime", empty: true}]},
          {rowtype:"field", name:"description", label: getText("event_description"), datatype:"text" }
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          event = update(event, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, 
                report: false, bookmark: false, export_event: false
              }} 
            }
          })
        }
      }
      return event;
    },

    formula: (item, edit) => {
      let formula = {
        options: {
          title: getText("title_formula"),
          title_field: "transnumber",
          icon: "Magic",
          fieldvalue: true,
          pattern: true,
          extend: "movement_head",
          panel: {
            arrow: true, more: true, trans: true, create: false,
            bookmark:["editor","trans","transnumber"], 
            help:"stock/formula"
          }
        },
        view: {
          movement: {
            type: "table",
            icon: "ListOl",
            title: getText("item_view"),
            data: "movement",
            fields: {
              product: {fieldtype:'string', label: getText("product_description")},
              unit: {fieldtype:'string', label: getText("product_unit")},
              cb_shared: {fieldtype:'bool', label: getText("formula_shared")},
              qty: {fieldtype:'number', label: getText("movement_qty")}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("invoice_crdate"), datatype:"date", disabled: true},
            {name:"closed", label: getText("document_closed"), datatype:"flip"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"field", name:"product_id", label: getText("product_partnumber"), datatype:"selector",
              empty: false, barcode: true, map:{seltype:"product_item", table:"movement_head", fieldname:"product_id", 
              lnktype:"product", transtype:"", label_field:"product", extend:true}},
          {rowtype:"col2", columns: [
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"qty", label: getText("movement_qty"), datatype:"float", map: {text:"qty", extend:true}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"intnotes", name:"notes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          formula = update(formula, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false, 
                report: false, bookmark: false, password: false, trans: false
              }} 
            }
          })
        } else {
          if (edit.dataset.translink.length > 0) {
            formula = update(formula, { rows: { 2: { columns: { 0: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source: "translink", value: "ref_id_1", text: "ref_id_2",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return formula;
    },

    groups: (item) => {
      let groups = {
        options: {
          icon: "Th",
          data: "groups",
          title: getText("title_groups"),
          panel: {
            page: "setting", more: false, help: "settings/groups"
          }
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: {action: "newItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem"}
            }
          }
        },
        rows: [
          {rowtype:"field", name:"groupvalue", label: getText("groups_groupvalue"), 
            datatype:"string"},
          {rowtype:"col2", columns: [
            {name:"groupname", label: getText("groups_groupname"), 
              datatype:"select", default: "", 
              options: [["",""],["department","department"],["eventgroup","eventgroup"],["paidtype","paidtype"],
                ["toolgroup","toolgroup"],["rategroup","rategroup"]]},
            {name:"inactive", label: getText("groups_inactive"), datatype:"flip"}]},
          {rowtype:"field", name:"description", label: getText("groups_description"), 
            datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id !== null) {
          groups = update(groups, { rows: { 1: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
        } else {
          groups = update(groups, {
            options: { 
              panel: {$merge: {
                new: false, 
                delete: false
              }} 
            }
          })
        }
      }
      return groups;
    },

    inventory: (item) => {
      let inventory = {
        options: {
          title: getText("title_inventory"),
          title_field: "transnumber",
          icon: "Truck",
          fieldvalue: true,
          pattern: true,
          panel: {
            arrow: true, more: true, trans: true, 
            create: false, cancellation: true, delete: false, 
            bookmark: ["editor","trans","transnumber"], 
            help: "stock/inventory"
          }
        },
        view: {
          movement: {
            type: "table",
            icon: "ListOl",
            title: getText("item_view"),
            data: "movement",
            fields: {
              product: {fieldtype:'string', label: getText("product_description")},
              unit: {fieldtype:'string', label: getText("product_unit")},
              notes: {fieldtype:'string', label: getText("movement_batchnumber")},
              qty: {fieldtype:'number', label: getText("movement_qty")}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"place_id", label: getText("delivery_place"), datatype:"selector",
              empty: false, map:{seltype:"place_warehouse", table:"trans", fieldname:"place_id", 
              lnktype:"place", transtype:"", label_field:"planumber"}},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("delivery_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("inventory_posdate"), datatype:"date"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          inventory = update(inventory, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        }
      }
      return inventory;
    },

    invoice_link: (item) => {
      let link = {
        options: {
          data: "link",
          title: getText("payment_view"),
          title_field: "",
          icon: "Money",
          panel: {}
        },
        rows: [
          {rowtype:"field", name:"ref_id_1", label: getText("payment_paidnumber"), datatype:"selector",
            empty: false, map:{seltype:"payment", table:"invoice_link", fieldname:"ref_id_1", 
            lnktype:"trans", transtype:"", lnkid:"trans_id", label_field:"transnumber"}},
          {rowtype:"col3", columns: [
            {name:"id", label: getText("payment_curr"), datatype:"string",
              map: {source:"invoice_link", value:"id", text:"curr"}},
            {name:"link_qty", label: getText("payment_amount"), datatype:"float",
              map: {source:"invoice_link_fieldvalue", value:"fieldname", text:"value"}},
            {name:"link_rate", label: getText("payment_rate"), datatype:"float",
              map: {source:"invoice_link_fieldvalue", value:"fieldname", text:"value"}}]}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          link = update(link, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return link;
    },

    invoice: (item, edit) => {
      let invoice = {
        options: {
          title: getText("title_invoice"),
          title_field: "transnumber",
          icon: "FileText",
          fieldvalue: true,
          pattern: true,
          panel: {
            arrow: true, more: true, trans: true,
            bookmark: ["editor","trans","transnumber"], 
            help: "document/document"
          }
        },
        view: {
          item: {
            type: "table",
            data: "item",
            icon: "ListOl",
            title: getText("item_view"),
            total:{
              netamount: getText("item_netamount"),
              vatamount: getText("item_vatamount"),
              amount: getText("item_amount")
            },
            fields: {
              description: {fieldtype:'string', label: getText("item_description")},
              unit: {fieldtype:'string', label: getText("item_unit")},
              qty: {fieldtype:'number', label: getText("item_qty")},
              amount: {fieldtype:'number', label: getText("item_amount")}
            }
          },
          invoice_link: {
            type: "list",
            data: "invoice_link",
            icon: "Money",
            title: getText("payment_view"),
            actions: {
              new: {action: "newEditorItem", fkey: "invoice_link"}, 
              edit: {action: "editEditorItem", fkey: "invoice_link"}, 
              delete: {action: "deleteEditorItem", fkey: "invoice_link", table: "link"}
            }
          },
          tool_movement: {
            type: "list",
            data: "tool_movement",
            icon: "Briefcase",
            title: getText("toolmovement_view"),
            audit_type: "trans",
            audit_transtype: "waybill",
            actions: {
              new: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("invoce_direction"), datatype:"select", empty: false,
              map: {source:"direction", value:"id", text:"groupvalue", label:"invoice" }},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("invoice_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("invoice_transdate"), datatype:"date"},
            {name:"duedate", label: getText("invoice_duedate"), datatype:"date"}]},
          {rowtype:"field", name:"customer_id", label: getText("customer_custname"), datatype:"selector",
              empty: false, map:{seltype:"customer", table:"trans", fieldname:"customer_id", 
              lnktype:"customer", transtype:"", label_field:"custname"}},
          {rowtype:"col4", columns: [
            {name:"curr", label: getText("document_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"acrate", label: getText("document_acrate"), datatype:"float", default:0},
            {name:"paid", label: getText("invoice_paid"), datatype:"flip"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"paidtype", label: getText("document_paidtype"), datatype:"select", empty: false,
              map: {source:"paidtype", value:"id", text:"groupvalue", label:"paidtype"}},
            {name:"department", label: getText("document_department"), datatype:"select", empty: true,
              map: {source:"department", value:"id", text:"groupvalue"}}]},
          {rowtype:"col2", columns: [
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}},
            {name:"project_id", label: getText("project_pronumber"), datatype:"selector",
              empty: true, map:{seltype:"project", table:"trans", fieldname:"project_id", 
              lnktype:"project", transtype:"", label_field:"pronumber"}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          invoice = update(invoice, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          invoice = update(invoice, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          if (edit.dataset.translink.length > 0) {
            invoice = update(invoice, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
          else if (edit.dataset.cancel_link.length > 0) {
            invoice = update(invoice, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype: "link",
              map: {
                source: "cancel_link", value: "ref_id_2", text: "ref_id_1",
                label_field: "transnumber", lnktype: "trans", 
                transtype: edit.dataset.cancel_link[0].transtype
              }
            }}}}}})
          }
          const direction = edit.dataset.groups.filter((group)=> {
            return (group.id === item.direction)
          })[0].groupvalue
          if (direction==="out" && item.transcast === "normal") {
            if (item.deleted === 0) {
              invoice = update(invoice, {
                options: { 
                  panel: {$merge: {
                    corrective: true
                  }} 
                }
              })
            } else {
              invoice = update(invoice, {
                options: { 
                  panel: {$merge: {
                    cancellation: true
                  }} 
                }
              })
            }
          }
        }
      }
      return invoice;
    },

    item: (item, edit) => { 
      let itemrow = {
        options: {
          title: getText("item_view"),
          title_field: "",
          icon: "ListOl",
          panel: {
            help:"document/item"
          }
        },
        rows: [
          {rowtype:"field", name:"product_id", label: getText("product_partnumber"), datatype:"selector",
            empty:  false, barcode:  true, map:{seltype:"product", table:"item", fieldname:"product_id", 
            lnktype:"product", transtype:"", label_field:"partnumber"}},
          {rowtype:"field", name:"description", label: getText("item_description"), datatype:"text"},
          {rowtype:"col2", columns: [
            {name:"unit", label: getText("item_unit"), datatype:"string"},
            {name:"ownstock", label: getText("item_ownstock"), datatype:"float"}]},
          {rowtype:"col3", columns: [
            {name:"qty", label: getText("item_qty"), datatype:"float"},
            {name:"discount", label: getText("item_discount"), datatype:"float", min:0, max:100},
            {name:"fxprice", label: getText("item_fxprice"), datatype:"float"}]},
          {rowtype:"col3", columns: [
            {name:"netamount", label: getText("item_netamount"), datatype:"float"},
            {name:"tax_id", label: getText("item_taxcode"), datatype:"select", empty:  true,
              map: {source:"tax", value:"id", text:"taxcode"}},
            {name:"amount", label: getText("item_amount"), datatype:"float"}]}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          itemrow = update(itemrow, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
        switch (edit.current.transtype) {
          case "invoice":
            itemrow = update(itemrow, { rows: { 
              2: {$set: {
                rowtype:"col3", 
                columns: [
                  {name:"unit", label: getText("item_unit"), datatype:"string"},
                  {name:"ownstock", label: getText("item_ownstock"), datatype:"float"},
                  {name:"deposit", label: getText("item_deposit_1"), datatype:"flip"}
                ]
              }} 
            }})
            break;
          case "offer":
            itemrow = update(itemrow, { rows: { 
              2: {$set: {
                rowtype:"col3", 
                columns: [
                  {name:"unit", label: getText("item_unit"), datatype:"string"},
                  {name:"ownstock", label: getText("item_ownstock"), datatype:"float"},
                  {name:"deposit", label: getText("item_deposit_2"), datatype:"flip"}
                ]
              }} 
            }})
            break;
          default:
            break;}
      }
      return itemrow;
    },

    log: () => { return {
      options: {
        title: getText("title_log"),
        title_field: "",
        edited: false,
        icon: "InfoCircle",
        panel: {}
      },
      view: {
        setting: {
          type:"table",
          actions: {
            new: null, 
            edit: null, 
            delete: null},
          fields: {
            crdate: {fieldtype:'date', label: getText("log_crdate")},
            empnumber: {fieldtype:'string', label: getText("log_empnumber")},
            logstate: {fieldtype:'string', label: getText("log_logstate")},
            nervatype: {fieldtype:'string', label: getText("log_nervatype")},
            refnumber: {fieldtype:'string', label: getText("log_refnumber")}
          }
        }
      },
      rows: [
        {rowtype:"col3", columns: [
          {name:"fromdate", label: getText("log_fromdate"), datatype:"date"},
          {name:"todate", label: getText("log_todate"), datatype:"date", empty: true},
          {name:"empnumber", label: getText("log_empnumber"), datatype:"string"}]},
        {rowtype:"col3", columns: [
          {name:"logstate", label: getText("log_logstate"), datatype:"select", empty: false,
            options: [["update","update"],["closed","closed"],["deleted","deleted"],
                ["print","print"],["login","login"],["logout","logout"]]},
          {name:"nervatype", label: getText("log_nervatype"), datatype:"select", default: "",
            options: [["",""], ["customer","customer"], ["employee","employee"], 
            ["event","event"], ["place","place"], ["product","product"], 
            ["project","project"], ["tool","tool"], ["trans","trans"]]},
          {name:"log_search", title: getText("browser_search"), label:"", focus: true,
            class:"full", icon: "Search", datatype:"button"}]}
      ]
    }},

    movement: (item, edit) => {
      let movement = {
        options: {
          title: getText("movement_view"),
          title_field: "",
          icon: "Truck",
          opposite: true,
          panel: {}
        }
      };
      switch (edit.current.transtype){
        case "delivery":
          movement.rows = [
            {rowtype:"col2", columns: [
              {name:"place_id", label: getText("movement_target"), datatype:"selector",
                empty: false, map:{seltype:"place_warehouse", table:"movement", fieldname:"place_id", 
                lnktype:"place", transtype:"", label_field:"planumber"}},
              {name:"trans_id", label: getText("movement_place"), datatype:"link",
                map: {source:"trans", value:"id", text:"place_id",
                  label_field:"planumber", lnktype:"place", transtype:""}}]},
            {rowtype:"field", name:"product_id", 
              label: getText("product_partnumber"), datatype:"selector",
              empty: false, barcode: true, map:{seltype:"product_item", table:"movement", fieldname:"product_id", 
              lnktype:"product", transtype:"", label_field:"product"}},
            {rowtype:"col3", columns: [
              {name:"trans_id", label: getText("movement_shippingdate"), datatype:"date",
                map: {source:"trans", value:"id", text:"transdate"}},
              {name:"notes", label: getText("movement_batchnumber"), datatype:"string"},
              {name:"qty", label: getText("movement_qty"), datatype:"float"}]}];
          break;
        case "inventory":
          movement.rows = [
            {rowtype:"field", name:"product_id", 
              label: getText("product_partnumber"), datatype:"selector",
              empty: false, barcode: true, map:{seltype:"product_item", table:"movement", fieldname:"product_id", 
              lnktype:"product", transtype:"", label_field:"product"}},
            {rowtype:"col3", columns: [
              {name:"trans_id", label: getText("movement_shippingdate"), datatype:"date",
                map: {source:"trans", value:"id", text:"transdate"}},
              {name:"notes", label: getText("movement_batchnumber"), datatype:"string"},
              {name:"qty", label: getText("movement_qty"), datatype:"float"}]}];
          break;
        case "production":
          movement.rows = [
            {rowtype:"col2", columns: [
              {name:"shippingdate", label: getText("movement_shippingdate"), 
                datatype:"datetime", empty: false},
              {name:"place_id", label: getText("movement_place"), datatype:"selector",
                empty: false, map:{seltype:"place_warehouse", table:"movement", fieldname:"place_id", 
                lnktype:"place", transtype:"", label_field:"planumber"}}]},
            {rowtype:"field", name:"product_id", label: getText("product_partnumber"), datatype:"selector",
                empty: false, barcode: true, map:{seltype:"product_item", table:"movement", fieldname:"product_id", 
                lnktype:"product", transtype:"", label_field:"product"}},
            {rowtype:"col2", columns: [
              {name:"notes", label: getText("movement_batchnumber"), datatype:"string"},
              {name:"qty", label: getText("movement_qty"), datatype:"float", opposite:true}]}];
          break;
        case "formula":
          movement.rows = [
            {rowtype:"field", name:"product_id", label: getText("product_partnumber"), datatype:"selector",
                empty: false, barcode: true, map:{seltype:"product_item", table:"movement", fieldname:"product_id", 
                lnktype:"product", transtype:"", label_field:"product"}},
            {rowtype:"col3", columns: [
              {name:"qty", label: getText("movement_qty"), datatype:"float"},
              {name:"shared", label: getText("formula_shared"), datatype:"flip"},
              {name:"place_id", label: getText("movement_place"), datatype:"selector",
                empty: false, map:{seltype:"place_warehouse", table:"movement", 
                fieldname:"place_id", 
                lnktype:"place", transtype:"", label_field:"planumber"}}]},
            {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"}];
          break;
        case "waybill":
          movement.rows = [
            {rowtype:"col2", columns: [
              {name:"shippingdate", label: getText("movement_shippingdate"), 
                datatype:"datetime", empty: false},
              {name:"tool_id", label: getText("tool_serial"), datatype:"selector",
                empty: false, map:{seltype:"tool", table:"movement", fieldname:"tool_id", 
                lnktype:"tool", transtype:"", label_field:"serial"}}]},
            {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"}];
          break;
        default:
          break;
      }
      if (typeof item !== "undefined") {
        if (item.id === null) {
          movement = update(movement, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return movement;
    },

    numberdef: () => {
      let numberdef = {
        options: {
          icon: "ListOl",
          data: "numberdef",
          title: getText("title_numberdef"),
          panel: {
            page:"setting", delete:false, new:false, 
            more:false, help:"settings/numberdef"
          }
        },
        view: {
          setting: {
            type:"table",
            actions: {
              new: null, 
              edit: {action: "editItem"}, 
              delete: null
            },
            fields: {
              numberkey: {fieldtype:'string', label: getText("numberdef_numberkey")},
              prefix: {fieldtype:'string', label: getText("numberdef_prefix")},
              is_year: {fieldtype:'string', label: getText("numberdef_isyear"), align:"center"},
              sep: {fieldtype:'string', label: getText("numberdef_sep_short"), align:"center"},
              len: {fieldtype:'number', label: getText("numberdef_len")},
              curvalue: {fieldtype:'number', align:"right", label: getText("numberdef_curvalue")}
            }
          }
        },
        rows: [
          {rowtype:"field", name:"numberkey", label: getText("numberdef_numberkey"), 
            datatype:"string", disabled: true},
          {rowtype:"col2", columns: [
            {name:"prefix", label: getText("numberdef_prefix"), datatype:"string"},
            {name:"curvalue", label: getText("numberdef_curvalue"), datatype:"integer"}]},
          {rowtype:"col3", columns: [
            {name:"isyear", label: getText("numberdef_isyear"), datatype:"flip"},
            {name:"sep", label: getText("numberdef_sep"), datatype:"string", length:1},
            {name:"len", label: getText("numberdef_len"), datatype:"integer"}]},
          {rowtype:"field", name:"description", label: getText("numberdef_description"), 
            datatype:"text"}
        ]
      };
      return numberdef;
    },

    offer: (item, edit) => {
      let offer = {
        options: {
          title: getText("title_offer"),
          title_field: "transnumber",
          icon: "FileText",
          fieldvalue: true,
          pattern: true,
          panel: {
            arrow:true, more:true, trans:true,
            bookmark:["editor","trans","transnumber"], 
            help:"document/document"
          }
        },
        view: {
          item: {
            type: "table",
            data: "item",
            icon: "ListOl",
            title: getText("item_view"),
            total:{
              netamount: getText("item_netamount"),
              vatamount: getText("item_vatamount"),
              amount: getText("item_amount")
            },
            fields: {
              description: {fieldtype:'string', label: getText("item_description")},
              unit: {fieldtype:'string', label: getText("item_unit")},
              qty: {fieldtype:'number', label: getText("item_qty")},
              amount: {fieldtype:'number', label: getText("item_amount")}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("offer_direction"), datatype:"select", empty: false,
              map: {source:"direction", value:"id", text:"groupvalue", label:"offer" }},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("offer_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("offer_transdate"), datatype:"date"},
            {name:"duedate", label: getText("offer_duedate"), datatype:"date"}]},
          {rowtype:"field", name:"customer_id", label: getText("customer_custname"), datatype:"selector",
              empty: false, map:{seltype:"customer", table:"trans", fieldname:"customer_id", 
              lnktype:"customer", transtype:"", label_field:"custname"}},
          {rowtype:"col4", columns: [
            {name:"curr", label: getText("document_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"acrate", label: getText("offer_acrate"), datatype:"float", default:0},
            {name:"paid", label: getText("offer_paid"), datatype:"flip"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"paidtype", label: getText("document_paidtype"), datatype:"select", empty: false,
              map: {source:"paidtype", value:"id", text:"groupvalue", label:"paidtype"}},
            {name:"department", label: getText("document_department"), datatype:"select", empty: true,
              map: {source:"department", value:"id", text:"groupvalue"}}]},
          {rowtype:"col2", columns: [
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}},
            {name:"project_id", label: getText("project_pronumber"), datatype:"selector",
              empty: true, map:{seltype:"project", table:"trans", fieldname:"project_id", 
              lnktype:"project", transtype:"", label_field:"pronumber"}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          offer = update(offer, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          offer = update(offer, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          if (edit.dataset.translink.length > 0) {
            offer = update(offer, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return offer;
    },

    order: (item, edit) => {
      let order = {
        options: {
          title: getText("title_order"),
          title_field: "transnumber",
          icon: "FileText",
          fieldvalue: true,
          pattern: true,
          panel: {
            arrow:true, more:true, trans:true,
            bookmark:["editor","trans","transnumber"], 
            help:"document/document"
          }
        },
        view: {
          item: {
            type: "table",
            data: "item",
            icon: "ListOl",
            title: getText("item_view"),
            total:{
              netamount: getText("item_netamount"),
              vatamount: getText("item_vatamount"),
              amount: getText("item_amount")
            },
            fields: {
              description: {fieldtype:'string', label: getText("item_description")},
              unit: {fieldtype:'string', label: getText("item_unit")},
              qty: {fieldtype:'number', label: getText("item_qty")},
              amount: {fieldtype:'number', label: getText("item_amount")}
            }
          },
          transitem_invoice: {
            type: "list",
            data: "transitem_invoice",
            icon: "FileText",
            title: getText("invoice_view"),
            audit_type: "trans",
            audit_transtype: "invoice",
            actions: {
              new: null, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "invoice"}, 
              delete: null
            }
          },
          transitem_shipping: {
            type: "table",
            data: "transitem_shipping",
            icon: "Truck",
            title: getText("shipping_view"),
            new_icon: "Truck",
            new_label: getText("title_shipping"),
            actions: {
              new: {action: "loadShipping"}, 
              edit: null, 
              delete: null
            },
            fields: {
              item_product: {fieldtype:'string', label: getText("shipping_item_product")},
              movement_product: {fieldtype:'string', label: getText("shipping_movement_product")},
              sqty: {fieldtype:'number', label: getText("shipping_sqty")}
            }
          },
          tool_movement: {
            type: "list",
            data: "tool_movement",
            icon: "Briefcase",
            title: getText("toolmovement_view"),
            audit_type: "trans",
            audit_transtype: "waybill",
            actions: {
              new: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("order_direction"), datatype:"select", empty: false,
              map: {source:"direction", value:"id", text:"groupvalue", label:"order" }},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("order_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("order_transdate"), datatype:"date"},
            {name:"duedate", label: getText("order_duedate"), datatype:"date"}]},
          {rowtype:"field", name:"customer_id", label: getText("customer_custname"), datatype:"selector",
              empty: false, map:{seltype:"customer", table:"trans", fieldname:"customer_id", 
              lnktype:"customer", transtype:"", label_field:"custname"}},
          {rowtype:"col4", columns: [
            {name:"curr", label: getText("document_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"acrate", label: getText("order_acrate"), datatype:"float", default:0},
            {name:"paid", label: getText("order_paid"), datatype:"flip"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"paidtype", label: getText("document_paidtype"), datatype:"select", empty: false,
              map: {source:"paidtype", value:"id", text:"groupvalue", label:"paidtype"}},
            {name:"department", label: getText("document_department"), datatype:"select", empty: true,
              map: {source:"department", value:"id", text:"groupvalue"}}]},
          {rowtype:"col2", columns: [
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}},
            {name:"project_id", label: getText("project_pronumber"), datatype:"selector",
              empty: true, map:{seltype:"project", table:"trans", fieldname:"project_id", 
              lnktype:"project", transtype:"", label_field:"pronumber"}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          order = update(order, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          order = update(order, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          if (edit.dataset.translink.length > 0) {
            order = update(order, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return order;
    },

    password: () => { return {
      options: {
        title: getText("title_password"),
        title_field: "",
        edited: false,
        icon: "Lock",
        panel: {
          delete: false, new: false
        }
      },
      view: {},
      rows: [
        {rowtype:"col3", columns: [
          {name:"username", label: getText("password_username"), datatype:"string", disabled: true},
          {name:"password_1", label: getText("password_new"), datatype:"password"},
          {name:"password_2", label: getText("password_verify"), datatype:"password"}
        ]}
      ]
    };},

    payment_link: (item) => {
      let link = {
        options: {
          data: "link",
          title: getText("invoice_view"),
          title_field: "",
          icon: "FileText",
          panel: {
            new: false
          }
        },
        rows: [
          {rowtype:"field", name:"ref_id_2", label: getText("invoice_transnumber"), datatype:"selector",
            empty: false, map:{seltype:"transitem_invoice", table:"payment_link", fieldname:"ref_id_2", 
            lnktype:"trans", transtype:"invoice", lnkid:"trans_id", label_field:"transnumber"}},
          {rowtype:"col3", columns: [
            {name:"id", label: getText("payment_curr"), datatype:"string",
              map: {source:"payment_link", value:"id", text:"curr"}},
            {name:"link_qty", label: getText("payment_amount"), datatype:"float",
              map: {source:"payment_link_fieldvalue", value:"fieldname", text:"value"}},
            {name:"link_rate", label: getText("payment_rate"), datatype:"float",
              map: {source:"payment_link_fieldvalue", value:"fieldname", text:"value"}}]}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          link = update(link, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false,
              }} 
            }
          })
        }
      }
      return link;
    },

    payment: (item) => {
      let payment = {
        options: {
          title: getText("payment_view"),
          title_field: "",
          icon: "Money",
          panel: {
            link: true, link_type:"payment_link", link_field:"ref_id_1",
            link_label: getText("label_link_invoice")
          }
        },
        rows: [
          {rowtype:"col2", columns: [
            {name:"paiddate", label: getText("payment_paiddate"), datatype:"date"},
            {name:"amount", label: getText("payment_amount"), datatype:"float"}]},
          {rowtype:"field", name:"notes", label: getText("payment_description"), datatype:"text"}]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          payment = update(payment, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, link: false
              }} 
            }
          })
        }
      }
      return payment;
    },

    place: (item, page) => { 
      let place = {
        options: {
          title: getText("title_place"),
          title_field: "planumber",
          icon: "Map",
          extend: "address",
          fieldvalue: true,
          panel: {
            more: true, report: false, 
            bookmark: ["editor","place","description","planumber"], 
            help: "settings/place"
          }
        },
        view: {
          setting: {
            type: "table",
            actions: {
              new: {action: "editItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem", tablename: "place"}
            },
            fields: {
              planumber: {fieldtype:'string', label: getText("place_planumber")},
              place_type: {fieldtype:'string', label: getText("place_placetype")},
              description: {fieldtype:'string', label: getText("place_description")}
            }
          },
          contact: {
            type: "list",
            data: "contact",
            icon: "Phone",
            title: getText("contact_view")
          },
        },
        rows: [
          {rowtype:"field", name:"description", label: getText("place_description"), datatype:"string"},
          {rowtype:"col2", columns: [
            {name:"placetype", label: getText("place_placetype"), datatype:"select", empty: true,
              map: {source:"placetype", value:"id", text:"groupvalue" }},
            {name:"inactive", label: getText("place_inactive"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"id", label: getText("address_zipcode"), datatype:"string",
              map: {source:"address", value:"ref_id", text:"zipcode", extend:true}},
            {name:"id", label: getText("address_city"), datatype:"string",
              map: {source:"address", value:"ref_id", text:"city", extend:true}}]},
          {rowtype:"field", name:"id", label: getText("address_street"), datatype:"string", 
            map: {source:"address", value:"ref_id", text:"street", extend:true}},
          {rowtype:"field", name:"notes", label: getText("place_notes"), datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          place = update(place, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, report: false, bookmark: false
              }} 
            }
          })
        } else {
          place = update(place, { rows: { 1: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          let placetype;
          if (page.dataset.placetype) {
            placetype = page.dataset.placetype.filter(
              (group) => (group.groupvalue === "warehouse"))[0]
          } else {
            placetype = page.dataset.groups.filter(
              (group) => ((group.groupname === "placetype") && (group.groupvalue === "warehouse")))[0]
          }
          if (item.placetype !== placetype.id) {
            place = update(place, { rows: { 1: {$merge: {
              rowtype: "col3"
            }}}})
            place = update(place, { rows: { 1: { columns: {$push: [{
              name:"curr", 
              label: getText("place_curr"), 
              datatype:"select", empty: false,
              map: {source:"currency", value:"curr", text:"curr"}
            }]}}}})
          }
        }
      }
      return place;
    },

    price: (item) => {
      let price = {
        options: {
          title: getText("price_view"),
          title_field: "",
          icon: "Dollar",
          panel: {}
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"validfrom", label: getText("price_validfrom"), datatype:"date"},
            {name:"validto", label: getText("price_validto"), datatype:"date", empty: true},
            {name:"vendor", label: getText("price_vendor"), datatype:"flip"}]},
          {rowtype:"col3", columns: [
            {name:"curr", label: getText("price_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"qty", label: getText("price_qty"), datatype:"float"},
            {name:"pricevalue", label: getText("price_pricevalue"), datatype:"float"}]},
          {rowtype:"field", name:"id", label: getText("customer_custname"), datatype:"selector",
            empty: true, map:{seltype:"customer", table:"price", fieldname:"customer_id", 
            lnktype:"customer", transtype:"", label_field:"custname"}}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          price = update(price, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return price;
    },

    printqueue: (item, edit, config) => {
      let printqueue = {
        options: {
          search_form: true,
          title: getText("title_printqueue"),
          title_field: getText("printqueue_head_title"),
          icon: "Filter",
          panel: {
            save:false, new:false, delete:false, more:true, report:false,
            search:true, export_all:true, print:false, bookmark:false, help:"program/printqueue"
          }
        },
        view: {
          items: {
            type: "list",
            data: "items",
            icon: "Print",
            edit_icon: "Check",
            title: getText("printqueue_selected_items"),
            actions: {
              new: null, 
              edit: {action: "exportQueueItem"}, 
              delete: {action: "deleteEditorItem", fkey: "items", table: "ui_printqueue"}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"nervatype", label: getText("printqueue_type"), datatype:"select", 
              empty: true, options: config.printqueue_type},
            {name:"startdate", label: getText("printqueue_startdate"), datatype:"date", empty: true},
            {name:"enddate", label: getText("printqueue_enddate"), datatype:"date", empty: true}]},
          {rowtype:"col2", columns: [
            {name:"transnumber", label: getText("printqueue_transnumber"), datatype:"string"},
            {name:"username", label: getText("printqueue_username"), datatype:"string"}
          ]},
          {rowtype:"col3", columns: [
            {name:"mode", label: getText("printqueue_mode"), datatype:"select", 
              empty: false, options: config.printqueue_mode},
            {name:"orientation", label: getText("report_orientation"), datatype:"select", 
              empty: false, options: config.report_orientation, default: config.page_orient},
            {name:"size", label: getText("report_size"), datatype:"select", 
              empty: false, options: config.report_size, default: config.page_size}]}
        ]};
      return printqueue;
    },

    product: (item) => { 
      let product = {
        options: {
          title: getText("title_product"),
          title_field: "partnumber",
          icon: "ShoppingCart",
          fieldvalue: true,
          panel: {
            more: true, 
            bookmark: ["editor","product","description","partnumber"], 
            help: "resources/product"
          }
        },
        view: {
          barcode: {
            type: "list",
            data: "barcode",
            icon: "Barcode",
            title: getText("barcode_view")},
          price: {
            type: "table",
            icon: "Dollar",
            title: getText("price_view"),
            data: "price",
            fields: {
              validfrom: {fieldtype:'date', label: getText("price_validfrom")},
              curr: {fieldtype:'string', label: getText("price_curr")},
              qty: {fieldtype:'number', label: getText("price_qty")},
              pricevalue: {fieldtype:'number', label: getText("price_pricevalue")}
            }
          },
          event: {
            type: "list",
            data: "event",
            icon: "Calendar",
            title: getText("event_view"),
            actions: {
              new: {action: "loadEditor", ntype: "event", ttype: null}, 
              edit: {action: "loadEditor", ntype: "event", ttype: null}, 
              delete: {action: "deleteEditorItem", fkey: "event", table: "event"}
            }
          }
      },
      rows: [
        {rowtype:"field", name:"description", 
          label: getText("product_description"), datatype:"string"},
        {rowtype:"col3", columns: [
          {name:"partnumber", label: getText("product_partnumber"), datatype:"string"},
          {name:"protype", label: getText("product_protype"), datatype:"select", 
            map: {source:"protype", value:"id", text:"groupvalue" }},
          {name:"unit", label: getText("product_unit"), datatype:"string"}]},
        {rowtype:"col3", columns: [
          {name:"tax_id", label: getText("product_tax"), datatype:"select", empty: false,
            map: {source:"tax", value:"id", text:"taxcode"}},
          {name:"webitem", label: getText("product_webitem"), datatype:"flip"},
          {name:"inactive", label: getText("product_inactive"), datatype:"flip"}]},
        {rowtype:"field", name:"notes", label: getText("product_notes"), datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {            
        if (item.id === null) {
          product = update(product, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, 
                report: false, bookmark: false
              }} 
            }
          })
        } else {
          product = update(product, { rows: { 1: { columns: { 
            0: {$merge: {
              disabled: true
            }},
            1: {$merge: {
              disabled: true
            }}
          }}}})
        }
      }
      return product;
    },

    production: (item, edit) => {
      let production = {
        options: {
          title: getText("title_production"),
          title_field: "transnumber",
          icon: "Flask",
          fieldvalue: true,
          pattern: true,
          extend: "movement_head",
          panel: {
            arrow:true, more:true, trans:true, create:false, formula:true,
            bookmark:["editor","trans","transnumber"], help:"stock/production"
          }
        },
        view: {
          movement: {
            type: "table",
            icon: "ListOl",
            title: getText("item_view"),
            data: "movement",
            fields: {
              product: {fieldtype:'string', label: getText("product_description")},
              unit: {fieldtype:'string', label: getText("product_unit")},
              notes: {fieldtype:'string', label: getText("movement_batchnumber")},
              opposite_qty: {fieldtype:'number', label: getText("movement_qty")}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("invoice_crdate"), datatype:"date", disabled: true},
            {name:"closed", label: getText("document_closed"), datatype:"flip"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col2", columns: [
            {name:"transdate", label: getText("production_transdate"), datatype:"date"},
            {name:"duedate", label: getText("production_duedate"), datatype:"datetime", empty: false}]},
          {rowtype:"field", name:"product_id", label: getText("product_partnumber"), datatype:"selector",
            empty: false, barcode: true, map:{seltype:"product_item", table:"movement_head", fieldname:"product_id", 
            lnktype:"product", transtype:"", label_field:"product", extend:true}},
          {rowtype:"col2", columns: [
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"place_id", label: getText("delivery_place"), datatype:"selector",
              empty: false, map:{seltype:"place_warehouse", table:"trans", fieldname:"place_id", 
              lnktype:"place", transtype:"", label_field:"planumber"}}]},
          {rowtype:"col2", columns: [
            {name:"batchnumber", label: getText("movement_batchnumber"), datatype:"string", 
              map: {text:"notes", extend:true}},
            {name:"qty", label: getText("movement_qty"), datatype:"float", map: {text:"qty", extend:true}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          production = update(production, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          if (edit.dataset.translink.length > 0) {
            production = update(production, { rows: { 3: { columns: { 0: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return production;
    },

    project: (item) => { 
      let project = {
        options: {
          title: getText("title_project"),
          title_field: "pronumber",
          icon: "Clock",
          fieldvalue: true,
          panel: {
            more:true, bookmark:["editor","project","description","pronumber"], help:"resources/project"
          }
        },
        view: {
          address: {
            type: "list",
            data: "address",
            icon: "Home",
            title: getText("address_view")
          },
          contact: {
            type: "list",
            data: "contact",
            icon: "Phone",
            title: getText("contact_view")
          },
          event: {
            type: "list",
            data: "event",
            icon: "Calendar",
            title: getText("event_view"),
            actions: {
              new: {action: "loadEditor", ntype: "event", ttype: null}, 
              edit: {action: "loadEditor", ntype: "event", ttype: null}, 
              delete: {action: "deleteEditorItem", fkey: "event", table: "event"}
            }
          }
      },
      rows: [
        {rowtype:"col3", columns: [
          {name:"startdate", label: getText("project_startdate"), datatype:"date", empty: true},
          {name:"enddate", label: getText("project_enddate"), datatype:"date", empty: true},
          {name:"inactive", label: getText("project_inactive"), datatype:"flip"}]},
        {rowtype:"col2", columns: [
          {name:"pronumber", label: getText("project_pronumber"), datatype:"string"},
          {name:"description", label: getText("project_description"), datatype:"string"}]},
        {rowtype:"field", name:"customer_id", label: getText("project_customer"), datatype:"selector",
          empty: true, map:{seltype:"customer", table:"project", fieldname:"customer_id", 
          lnktype:"customer", transtype:"", label_field:"custname"}},
        {rowtype:"field", name:"notes", label: getText("project_notes"), datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          project = update(project, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false,
                report: false, bookmark: false
              }} 
            }
          })
        } else {
          project = update(project, { rows: { 1: {$set: { 
            rowtype:"field", name:"description", 
            label: getText("project_description"), 
            datatype:"string"
          }}}})
        }
      }
      return project;
    },

    rate: (item, edit) => { 
      let rate = {
        options: {
          title: getText("title_rate"),
          icon: "Strikethrough",
          fieldvalue: true,
          panel: {
            more:true, report:false, 
            bookmark:false, help:"settings/rate"
          }
        },
        view: {},
      rows: [
        {rowtype:"col2", columns: [
          {name:"ratetype", label: getText("rate_ratetype"), datatype:"select", empty: false,
            map: {source:"ratetype", value:"id", text:"groupvalue"}},
          {name:"ratedate", label: getText("rate_ratedate"), datatype:"date"}]},
        {rowtype:"col2", columns: [
          {name:"curr", label: getText("rate_curr"), datatype:"select", empty: false,
            map: {source:"currency", value:"curr", text:"curr"}},
          {name:"ratevalue", label: getText("rate_ratevalue"), datatype:"float"}]},
        {rowtype:"col2", columns: [
          {name:"rategroup", label: getText("rate_rategroup"), datatype:"select", empty: true,
            map: {source:"rategroup", value:"id", text:"groupvalue"}},
          {name:"place_id", label: getText("rate_planumber"), datatype:"selector",
            empty: true, map:{seltype:"place_bank", table:"rate", fieldname:"place_id", 
            lnktype:"place", transtype:"", label_field:"planumber"}}]}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          rate = update(rate, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false,
              }} 
            }
          })
          item.ratetype = edit.dataset.ratetype.filter(item => (item.groupvalue === "rate"))[0].id
          let def_rate_currency = edit.dataset.settings.filter(item => (item.fieldname === "default_currency"))[0]
          if (typeof def_rate_currency !== "undefined") {
            item.curr = def_rate_currency.value;
          }
        }
      }
      return rate;
    },

    program: () => { return {
      options: {
        title: getText("title_program"),
        title_field: "",
        edited: false,
        icon: "Keyboard",
        panel: {}
      },
      view: {},
      rows: [
        {rowtype:"col4", columns: [
          {name:"paginationPage", label: getText("program_page"), datatype:"integer"},
          {name:"history", label: getText("program_history"), datatype:"integer"},
          {name:"export_sep", label: getText("program_export_sep"), datatype:"string", length:1},
          {name:"decimal_sep", label: getText("program_decimal_sep"), datatype:"select", 
            empty: false, options: getSetting("separators")},
        ]},
        {rowtype:"col3", columns: [
          {name:"page_size", label: getText("program_page_size"), datatype:"select", 
            empty: false, options: getSetting("report_size")},
          {name:"dateFormat", label: getText("program_date_format"), datatype:"select", 
            empty: false, options: getSetting("dateStyle")},
          {name:"calendar", label: getText("program_calendar"), datatype:"select", 
            empty: false, options: getSetting("calendarLocales")},
        ]}
      ]
    };},

    receipt: (item, edit) => {
      let receipt = {
        options: {
          title: getText("title_receipt"),
          title_field: "transnumber",
          icon: "FileText",
          fieldvalue: true,
          pattern: true,
          panel: {
            arrow:true, more:true, trans:true, create:false,
            bookmark:["editor","trans","transnumber"], 
            help:"document/document"
          }
        },
        view: {
          item: {
            type: "table",
            data: "item",
            icon: "ListOl",
            title: getText("item_view"),
            total:{
              netamount: getText("item_netamount"),
              vatamount: getText("item_vatamount"),
              amount: getText("item_amount")
            },
            fields: {
              description: {fieldtype:'string', label: getText("item_description")},
              unit: {fieldtype:'string', label: getText("item_unit")},
              qty: {fieldtype:'number', label: getText("item_qty")},
              amount: {fieldtype:'number', label: getText("item_amount")}
            }
          },
          tool_movement: {
            type: "list",
            data: "tool_movement",
            icon: "Briefcase",
            title: getText("toolmovement_view"),
            audit_type: "trans",
            audit_transtype: "waybill",
            actions: {
              new: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("invoce_direction"), datatype:"select", empty: false, disabled: true,
              map: {source:"direction", value:"id", text:"groupvalue", label:"receipt" }},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("receipt_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("receipt_transdate"), datatype:"date"},
            {name:"duedate", label: getText("receipt_duedate"), datatype:"date"}]},
          {rowtype:"col4", columns: [
            {name:"curr", label: getText("document_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"acrate", label: getText("document_acrate"), datatype:"float", default:0},
            {name:"paid", label: getText("receipt_paid"), datatype:"flip"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"paidtype", label: getText("document_paidtype"), datatype:"select", empty: false,
              map: {source:"paidtype", value:"id", text:"groupvalue", label:"paidtype"}},
            {name:"department", label: getText("document_department"), datatype:"select", empty: true,
              map: {source:"department", value:"id", text:"groupvalue"}}]},
          {rowtype:"col2", columns: [
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}},
            {name:"project_id", label: getText("project_pronumber"), datatype:"selector",
              empty: true, map:{seltype:"project", table:"trans", fieldname:"project_id", 
              lnktype:"project", transtype:"", label_field:"pronumber"}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          receipt = update(receipt, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          if (edit.dataset.translink.length > 0) {
            receipt = update(receipt, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          } else if (edit.dataset.cancel_link.length > 0) {
            receipt = update(receipt, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source:"cancel_link", value:"ref_id_2", text:"ref_id_1",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.cancel_link[0].transtype
              }
            }}}}}})
          }
          const direction = edit.dataset.groups.filter((group)=> {
            return (group.id === item.direction)
          })[0].groupvalue
          if (direction==="out" && item.transcast === "normal") {
            if (item.deleted === 0) {
              receipt = update(receipt, {
                options: { 
                  panel: {$merge: {
                    corrective: true
                  }} 
                }
              })
            } else {
              receipt = update(receipt, {
                options: { 
                  panel: {$merge: {
                    cancellation: true
                  }} 
                }
              })
            }
          }
        }
      }
      return receipt;
    },

    rent: (item, edit) => {
      let rent = {
        options: {
          title: getText("title_rent"),
          title_field: "transnumber",
          icon: "FileText",
          fieldvalue: true,
          pattern: true,
          edited: false,
          panel: {
            arrow:true, more:true, trans:true,
            bookmark:["editor","trans","transnumber"], 
            help:"document/document"
          }
        },
        view: {
          item: {
            type: "table",
            data: "item",
            icon: "ListOl",
            title: getText("item_view"),
            total:{
              netamount: getText("item_netamount"),
              vatamount: getText("item_vatamount"),
              amount: getText("item_amount")
            },
            fields: {
              description: {fieldtype:'string', label: getText("item_description")},
              unit: {fieldtype:'string', label: getText("item_unit")},
              qty: {fieldtype:'number', label: getText("item_qty")},
              amount: {fieldtype:'number', label: getText("item_amount")}
            }
          },
          transitem_invoice: {
            type: "list",
            data: "transitem_invoice",
            icon: "FileText",
            title: getText("invoice_view"),
            audit_type: "trans",
            audit_transtype: "invoice",
            actions: {
              new: null, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "invoice"}, 
              delete: null
            }
          },
          transitem_shipping: {
            type: "table",
            data: "transitem_shipping",
            icon: "Truck",
            title: getText("shipping_view"),
            actions: {
              new: {action: "loadShipping"}, 
              edit: null, 
              delete: null
            },
            fields: {
              item_product: {fieldtype:'string', label: getText("shipping_item_product")},
              movement_product: {fieldtype:'string', label: getText("shipping_movement_product")},
              sqty: {fieldtype:'number', label: getText("shipping_sqty")}
            }
          },
          tool_movement: {
            type: "list",
            data: "tool_movement",
            icon: "Briefcase",
            title: getText("toolmovement_view"),
            audit_type: "trans",
            audit_transtype: "waybill",
            actions: {
              new: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("rental_direction"), datatype:"select", empty: false,
              map: {source:"direction", value:"id", text:"groupvalue", label:"rent" }},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("rental_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("rental_transdate"), datatype:"date"},
            {name:"duedate", label: getText("rental_duedate"), datatype:"date"}]},
          {rowtype:"field", name:"customer_id", label: getText("customer_custname"), datatype:"selector",
              empty: false, map:{seltype:"customer", table:"trans", fieldname:"customer_id", 
              lnktype:"customer", transtype:"", label_field:"custname"}},
          {rowtype:"col3", columns: [
            {name:"trans_reholiday", label: getText("trans_reholiday"), datatype:"float",
              map: {source:"fieldvalue", value:"fieldname", text:"value"}},
            {name:"trans_rebadtool", label: getText("trans_rebadtool"), datatype:"float",
              map: {source:"fieldvalue", value:"fieldname", text:"value"}},
            {name:"trans_reother", label: getText("trans_reother"), datatype:"float",
              map: {source:"fieldvalue", value:"fieldname", text:"value"}}]},
          {rowtype:"field", name:"trans_rentnote", label: getText("trans_rentnote"), datatype:"string",
            map: {source:"fieldvalue", value:"fieldname", text:"value"}},
          {rowtype:"col4", columns: [
            {name:"curr", label: getText("document_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"acrate", label: getText("rental_acrate"), datatype:"float"},
            {name:"paid", label: getText("rental_paid"), datatype:"flip"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"paidtype", label: getText("document_paidtype"), datatype:"select", empty: false,
              map: {source:"paidtype", value:"id", text:"groupvalue", label:"paidtype"}},
            {name:"department", label: getText("document_department"), datatype:"select", empty: true,
              map: {source:"department", value:"id", text:"groupvalue"}}]},
          {rowtype:"col2", columns: [
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}},
            {name:"project_id", label: getText("project_pronumber"), datatype:"selector",
              empty: true, map:{seltype:"project", table:"trans", fieldname:"project_id", 
              lnktype:"project", transtype:"", label_field:"pronumber"}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id === null) {
          rent = update(rent, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false,
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          rent = update(rent, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          if (edit.dataset.translink.length > 0) {
            rent = update(rent, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), datatype:"link",
              map: {
                source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return rent;
    },

    report: (item, edit, config) => {
      let report = {
        options: {
          title: getText("title_report"),
          title_field: "repname",
          icon: "ChartBar",
          panel: {
            save:false, new:false, delete:false, more:true, report:false,
            print:true, export_pdf:true, export_xml:true, bookmark:false, help:"program/report"
          }
        },
        view: {},
        rows: [
          {rowtype:"label", name:"description"},
          {rowtype:"col3", columns: [
            {name:"oslabel", label: getText("report_orientation")+" / "+getText("report_size"), datatype:"label"},
            {name:"orientation", label:"", datatype:"select", 
              empty: false, options: config.report_orientation, default: config.page_orient},
            {name:"size", label:"", datatype:"select", 
              empty: false, options: config.report_size, default: config.page_size}]}
        ]};
      if (typeof item !== "undefined") {
        if (item.ftype === "csv") {
          report = update(report, {
            rows: {$set: [{rowtype:"label", name:"description"}]},
            options: { 
              panel: {$merge: {
                print: false, export_pdf: false, export_xml: false, export_csv: true
              }} 
            }
          })
        }
      }
      return report;
    },

    setting: () => {
      let setting = {
        options: {
          icon: "Cog",
          data: "fieldvalue",
          title: getText("title_dbsettings"),
          panel: {
            page:"setting", delete:false, new:false, more:false, help:"settings/setting"
          }
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: null, 
              edit: {action: "editItem"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"field", name:"fieldname", label: getText("fields_fieldname"), 
            datatype:"string", disabled: true},
          {rowtype:"field", name:"label", label: getText("fields_fielddef"), 
            datatype:"string", disabled: true},
          {rowtype:"field", name:"fieldvalue_value", label: getText("fields_value"), 
            datatype:"fieldvalue"},
          {rowtype:"field", name:"fieldvalue_notes", label: getText("fields_notes"), 
            datatype:"text"}
        ]
      };
      return setting;
    },

    shipping: (item) => {
      let shipping = {
        options: {
          title: getText("title_shipping"),
          title_field: "transnumber",
          icon: "Truck",
          panel: {
            back:true, save:false, delete:false, new:false, shipping:true, help:"stock/shipping"
          }
        },
        view: {
          shipping_items: {
            type: "table",
            data: "shipping_items_",
            icon: "ListOl",
            edit_icon: "Plus",
            delete_icon: "Book",
            title: getText("shipping_items"),
            actions: {
              new: null, 
              edit: {action: "addShippingRow"},
              delete: {action: "showShippingStock"}
            },
            fields: {
              product: {fieldtype:'string', label: getText("shipping_movement_product")},
              qty: {fieldtype:'number', label: getText("shipping_qty")},
              tqty: {fieldtype:'number', label: getText("shipping_turnover")},
              diff: {fieldtype:'number', label: getText("shipping_diff"), format:true}
            }
          },
          shiptemp_items: {
            type: "table",
            data: "shiptemp",
            title: getText("shipping_create"),
            icon: "Plus",
            actions: {
              new: null, 
              edit: {action: "editShippingRow"}, 
              delete: {action: "deleteShippingRow"}
            },
            fields: {
              product: {fieldtype:'string', label: getText("shipping_product")},
              batch_no: {fieldtype:'string', label: getText("movement_batchnumber")},
              qty: {fieldtype:'number', label: getText("movement_qty")},
              diff: {fieldtype:'number', label: getText("shipping_diff"), format:true}
            }
          },
          shipping_delivery: {
            type: "list",
            data: "shipping_delivery",
            title: getText("shipping_delivery"),
            icon: "Truck",
            actions: {
              new: null, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "delivery"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col2", columns: [
            {name:"delivery_type", label: getText("delivery_direction"), 
              datatype:"string", disabled: true},
            {name:"id", label: getText("customer_custname"), datatype:"link",
              map: {source:"trans", value:"id", text:"customer_id",
                label_field:"custname", lnktype:"customer", transtype:""}}]},
          {rowtype:"col2", columns: [
            {name:"shippingdate", label: getText("movement_shippingdate"), 
              datatype:"datetime", empty: false},
            {name:"shipping_place_id", label: getText("movement_place"), datatype:"selector", 
              empty: false, map:{seltype:"place_warehouse", table:"trans", fieldname:"shipping_place_id", 
              lnktype:"place", transtype:"", label_field:"planumber"}}]}
        ]};
      return shipping;
    },

    tax: (item) => {
      let tax = {
        options: {
          icon: "Ticket",
          data: "tax",
          title: getText("title_tax"),
          panel: {
            page:"setting", more:false
          }
        },
        view: {
          setting: {
            type: "table",
            actions: {
              new: {action: "newItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem"}
            },
            fields: {
              taxcode: {fieldtype:'string', label: getText("tax_taxcode")},
              description: {fieldtype:'string', label: getText("tax_description")},
              rate: {fieldtype:'number', label: getText("tax_rate")},
              inact: {fieldtype:'string', label: getText("tax_inactive"), align:"center"}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"taxcode", label: getText("tax_taxcode"), datatype:"string"},
            {name:"rate", label: getText("tax_rate"), datatype:"float"},
            {name:"inactive", label: getText("tax_inactive"), datatype:"flip"}]},
          {rowtype:"field", name:"description", label: getText("tax_description"), datatype:"string"}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id !== null) {
          tax = update(tax, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
        } else {
          tax = update(tax, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return tax;
    },
    
    template: (item) => {
      let template = {
        options: {
          icon: "TextHeight",
          data: "template",
          title: getText("title_report_editor"),
          panel: {}
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: null, 
              edit: {action: "editItem"}, 
              delete: null
            }
          }
        },
        rows: []
      };
      return template;
    },

    tool: (item) => { 
      let tool = {
        options: {
          title: getText("title_tool"),
          title_field: "serial",
          icon: "Wrench",
          fieldvalue: true,
          panel: {
            more:true, 
            bookmark:["editor","tool","description","serial"], 
            help:"resources/tool"
          }
        },
        view: {
          event: {
            type: "list",
            data: "event",
            icon: "Calendar",
            title: getText("event_view"),
            actions: {
              new: {action: "loadEditor", ntype: "event", ttype: null}, 
              edit: {action: "loadEditor", ntype: "event", ttype: null}, 
              delete: {action: "deleteEditorItem", fkey: "event", table: "event"}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"serial", label: getText("tool_serial"), datatype:"string"},
            {name:"toolgroup", label: getText("tool_toolgroup"), datatype:"select", empty: true,
              map: {source:"toolgroup", value:"id", text:"groupvalue"}},
            {name:"inactive", label: getText("tool_inactive"), datatype:"flip"}]},
          {rowtype:"field", name:"description", label: getText("tool_description"), datatype:"string"},
          {rowtype:"field", name:"product_id", 
            label: getText("product_partnumber"), datatype:"selector",
            empty: false, barcode: true, map:{seltype:"product_item", table:"tool", fieldname:"product_id", 
            lnktype:"product", transtype:"", label_field:"product"}},
          {rowtype:"field", name:"notes", label: getText("tool_notes"), datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {            
        if (item.id === null) {
          tool = update(tool, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                new: false, delete: false, report: false, bookmark: false
              }} 
            }
          })
        } else {
          tool = update(tool, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
        }
      }
      return tool;
    },

    ui_menu: (item) => {
      let ui_menu = {
        options: {
          icon: "Share",
          data: "ui_menu",
          title: getText("title_menucmd"),
          panel: {
            page:"setting", more:false, help:"settings/uimenu"
          }
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: {action: "newItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem"}
            }
          },
          items: {
            type:"table",
            data:"ui_menufields",
            actions: {
              new: {action: "editMenuField"}, 
              edit: {action: "editMenuField"}, 
              delete: {action: "deleteItemRow", table:"ui_menufields"}
            },
            fields: {
              fieldname: {fieldtype:'string', label: getText("menufields_fieldname")},
              description: {fieldtype:'string', label: getText("menufields_description")},
              fieldtype_name: {fieldtype:'string', label: getText("menufields_fieldtype")},
              orderby: {fieldtype:'number', label: getText("menufields_orderby")}
            }
          }
        },
        rows: [
          {rowtype:"col2", columns: [
            {name:"menukey", label: getText("menucmd_menukey"), datatype:"string"},
            {name:"description", label: getText("menucmd_description"), datatype:"string"}]},
          {rowtype:"col3", columns: [
            {name:"method", label:getText("menucmd_method"), datatype:"select", 
              map: {source:"method", value:"id", text:"groupvalue" }},
            {name:"modul", label: getText("menucmd_modul"), datatype:"string"},
            {name:"icon", label: getText("menucmd_icon"), datatype:"string"}
          ]},
          {rowtype:"col2", columns: [
            {name:"funcname", label: getText("menucmd_funcname"), datatype:"string"},
            {name:"address", label: getText("menucmd_address"), datatype:"string"}]}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id !== null) {
          ui_menu = update(ui_menu, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
        }
      }
      return ui_menu;
    },
    
    usergroup: (item) => {
      let usergroup = {
        options: {
          icon: "Key",
          data: "groups",
          title: getText("title_usergroup"),
          panel: {
            page:"setting", more:false, help:"settings/usergroup"
          }
        },
        view: {
          setting: {
            type: "list",
            actions: {
              new: {action: "newItem"}, 
              edit: {action: "editItem"}, 
              delete: {action: "deleteItem"}
            }
          },
          items: {
            type:"table",
            data:"audit",
            actions: {
              new: {action: "editAudit"}, 
              edit: {action: "editAudit"}, 
              delete: {action: "deleteItemRow", table:"ui_audit"}
            },
            fields: {
              nervatype_name: {fieldtype:'string', label: getText("audit_nervatype")},
              subtype_name: {fieldtype:'string', label: getText("audit_subtype")},
              inputfilter_name: {fieldtype:'string', label: getText("audit_inputfilter")},
              supervisor_name: {fieldtype:'string', label: getText("audit_supervisor")}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"groupvalue", label: getText("groups_groupvalue"), datatype:"string"},
            {name:"transfilter", label: getText("groups_transfilter"), 
              datatype:"select",  empty: true,
              map: {source:"transfilter", value:"id", text:"groupvalue" }},
            {name:"inactive", label: getText("groups_inactive"), datatype:"flip"}]},
          {rowtype:"field", name:"description", label: getText("groups_description"), 
            datatype:"text"}
        ]
      };
      if (typeof item !== "undefined") {
        if (item.id !== null) {
          usergroup = update(usergroup, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
        } else {
          usergroup = update(usergroup, {
            options: { 
              panel: {$merge: {
                new: false, delete: false
              }} 
            }
          })
        }
      }
      return usergroup;
    },

    waybill: (item, edit) => {
      let waybill = {
        options: {
          title: getText("title_waybill"),
          title_field: "transnumber",
          icon: "Briefcase",
          fieldvalue: true,
          pattern: true,
          extend: "refvalue",
          panel: {
            arrow:true, more:true, trans:true, create:false,
            bookmark:["editor","trans","transnumber"], 
            help:"stock/waybill"
          }
        },
        view: {
          movement: {
            type: "table",
            icon: "ListOl",
            title: getText("item_view"),
            data: "movement",
            fields: {
              shippingdate: {fieldtype:'date', label: getText("movement_shippingdate2")},
              serial: {fieldtype:'string', label: getText("tool_serial")},
              tooldesc: {fieldtype:'string', label: getText("tool_description")}
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("document_direction"), datatype:"select", empty: false,
              map: {source:"direction", value:"id", text:"groupvalue", label:"waybill" }},
            {name:"crdate", label: getText("waybill_crdate"), datatype:"date", disabled: true},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col2", columns: [
            {name:"seltype", label: getText("waybill_seltype"), datatype:"select", 
              empty: false, olabel:"waybill", extend:true,
              options: [["transitem","transitem"], ["customer","customer"], 
                ["employee","employee"]]},
            {name:"ref_id", label: getText("waybill_reference"), datatype:"selector",
              empty: false, map:{seltype:"transitem", table:"extend", fieldname:"ref_id", 
              lnktype:"trans", transtype:"", label_field:"refnumber", extend:true}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          waybill = update(waybill, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false, 
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          waybill = update(waybill, { rows: { 0: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          waybill = update(waybill, { rows: { 1: { columns: { 0: {$merge: {
            disabled: true
          }}}}}})
          if (item.customer_id !== null) {
            waybill = update(waybill, { rows: { 1: { columns: { 1: { map: {$merge: {
              seltype: "customer",
              lnktype: "customer"
            }}}}}}})
          } else if (item.employee_id !== null) {
            waybill = update(waybill, { rows: { 1: { columns: { 1: { map: {$merge: {
              seltype: "employee",
              lnktype: "employee"
            }}}}}}})
          } else {
            waybill = update(waybill, { rows: { 1: { columns: { 1: { map: {$merge: {
              seltype: "transitem",
              lnktype: "trans"
            }}}}}}})
            if (edit.dataset.translink.length > 0) {
              waybill = update(waybill, { rows: { 1: { columns: { 1: { map: {$merge: {
                transtype: edit.dataset.translink[0].transtype
              }}}}}}})
            }
          }
        }
      }
      return waybill;
    },
    
    worksheet: (item, edit) => {
      let worksheet = {
        options: {
          title: getText("title_worksheet"),
          title_field: "transnumber",
          icon: "FileText",
          fieldvalue: true,
          pattern: true,
          edited: false,
          panel: {
            arrow:true, more:true, trans:true,
            bookmark:["editor","trans","transnumber"], 
            help:"document/document"
          }
        },
        view: {
          item: {
            type: "table",
            data: "item",
            icon: "ListOl",
            title: getText("item_view"),
            total:{
              netamount: getText("item_netamount"),
              vatamount: getText("item_vatamount"),
              amount: getText("item_amount")
            },
            fields: {
              description: {fieldtype:'string', label: getText("item_description")},
              unit: {fieldtype:'string', label: getText("item_unit")},
              qty: {fieldtype:'number', label: getText("item_qty")},
              amount: {fieldtype:'number', label: getText("item_amount")}}
          },
          transitem_invoice: {
            type: "list",
            data: "transitem_invoice",
            icon: "FileText",
            title: getText("invoice_view"),
            audit_type: "trans",
            audit_transtype: "invoice",
            actions: {
              new: null, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "invoice"}, 
              delete: null
            }
          },
          transitem_shipping: {
            type: "table",
            data: "transitem_shipping",
            icon: "Truck",
            title: getText("shipping_view"),
            actions: {
              new: {action: "loadShipping"}, 
              edit: null, 
              delete: null
            },
            fields: {
              item_product: {fieldtype:'string', label: getText("shipping_item_product")},
              movement_product: {fieldtype:'string', label: getText("shipping_movement_product")},
              sqty: {fieldtype:'number', label: getText("shipping_sqty")}
            }
          },
          tool_movement: {
            type: "list",
            data: "tool_movement",
            icon: "Briefcase",
            title: getText("toolmovement_view"),
            audit_type: "trans",
            audit_transtype: "waybill",
            actions: {
              new: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              edit: {action: "loadEditor", ntype: "trans", ttype: "waybill"}, 
              delete: null
            }
          }
        },
        rows: [
          {rowtype:"col3", columns: [
            {name:"direction", label: getText("worksheet_direction"), datatype:"select", empty: false, disabled: true,
              map: {source:"direction", value:"id", text:"groupvalue", label:"worksheet" }},
            {name:"ref_transnumber", label: getText("document_ref_transnumber"), datatype:"string"},
            {name:"transtate", label: getText("document_transtate"), datatype:"select", empty: false,
              map: {source:"transtate", value:"id", text:"groupvalue", label:"state"}}]},
          {rowtype:"col3", columns: [
            {name:"crdate", label: getText("worksheet_crdate"), datatype:"date", disabled: true},
            {name:"transdate", label: getText("worksheet_transdate"), datatype:"date"},
            {name:"duedate", label: getText("worksheet_duedate"), datatype:"date"}]},
          {rowtype:"field", name:"customer_id", label: getText("customer_custname"), datatype:"selector",
              empty: false, map:{seltype:"customer", table:"trans", fieldname:"customer_id", 
              lnktype:"customer", transtype:"", label_field:"custname"}},
          {rowtype:"col3", columns: [
            {name:"trans_wsdistance", label: getText("trans_wsdistance"), datatype:"float",
              map: {source:"fieldvalue", value:"fieldname", text:"value"}},
            {name:"trans_wsrepair", label: getText("trans_wsrepair"), datatype:"float",
              map: {source:"fieldvalue", value:"fieldname", text:"value"}},
            {name:"trans_wstotal", label: getText("trans_wstotal"), datatype:"float",
              map: {source:"fieldvalue", value:"fieldname", text:"value"}}]},
          {rowtype:"field", name:"trans_wsnote", label: getText("trans_wsnote"), datatype:"string",
            map: {source:"fieldvalue", value:"fieldname", text:"value"}},
          {rowtype:"col4", columns: [
            {name:"curr", label: getText("document_curr"), datatype:"select", empty: true,
              map: {source:"currency", value:"curr", text:"curr"}},
            {name:"acrate", label: getText("worksheet_acrate"), datatype:"float", default:0},
            {name:"paid", label: getText("worksheet_paid"), datatype:"flip"},
            {name:"closed", label: getText("document_closed"), datatype:"flip"}]},
          {rowtype:"col2", columns: [
            {name:"paidtype", label: getText("document_paidtype"), datatype:"select", empty: false,
              map: {source:"paidtype", value:"id", text:"groupvalue", label:"paidtype"}},
            {name:"department", label: getText("document_department"), datatype:"select", empty: true,
              map: {source:"department", value:"id", text:"groupvalue"}}]},
          {rowtype:"col2", columns: [
            {name:"employee_id", label: getText("employee_empnumber"), datatype:"selector",
              empty: true, map:{seltype:"employee", table:"trans", fieldname:"employee_id", 
              lnktype:"employee", transtype:"", label_field:"empnumber"}},
            {name:"project_id", label: getText("project_pronumber"), datatype:"selector",
              empty: true, map:{seltype:"project", table:"trans", fieldname:"project_id", 
              lnktype:"project", transtype:"", label_field:"pronumber"}}]},
          {rowtype:"field", name:"notes", label: getText("document_notes"), datatype:"text"},
          {rowtype:"field", name:"intnotes", label: getText("document_intnotes"), datatype:"text"}
        ]};
      if (typeof item !== "undefined") {
        if (item.id === null) {
          worksheet = update(worksheet, {
            view: {$set: {}},
            options: { 
              panel: {$merge: {
                arrow: false, new: false, delete: false, 
                report: false, bookmark: false, trans: false
              }} 
            }
          })
        } else {
          if (edit.dataset.translink.length > 0) {
            worksheet = update(worksheet, { rows: { 0: { columns: { 1: {$set: {
              name:"id", 
              label: getText("document_ref_transnumber"), 
              datatype:"link",
              map: {
                source:"translink", value:"ref_id_1", text:"ref_id_2",
                label_field:"transnumber", lnktype:"trans", 
                transtype: edit.dataset.translink[0].transtype
              }
            }}}}}})
          }
        }
      }
      return worksheet;
    }

  }
}