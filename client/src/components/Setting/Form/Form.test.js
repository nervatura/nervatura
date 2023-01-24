import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './setting-form.js';
import { Template, Default, Items, Log } from  './Form.stories.js';

describe('Setting-Form', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onSettingEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onSettingEvent
    }));
    const form = element.querySelector('#setting_form');
    expect(form).to.exist;

    const inputRow = form.shadowRoot.querySelector('#row_2')
    const fieldvalue = inputRow.shadowRoot.querySelector('#field_fieldvalue_value').shadowRoot.querySelector('#field_fieldvalue_value')
    fieldvalue._onInput({ target: { value: "abc" } })
    sinon.assert.callCount(onSettingEvent, 1);

  })

  it('renders in the Items state', async () => {
    const onSettingEvent = sinon.spy()
    let element = await fixture(Template({
      ...Items.args, onSettingEvent
    }));
    let form = element.querySelector('#setting_form');
    expect(form).to.exist;

    const formView = form.shadowRoot.querySelector('#form_view')
    const btnAdd = formView.shadowRoot.querySelector('#btn_add')
    btnAdd.click()
    sinon.assert.callCount(onSettingEvent, 1);

    const rowEdit = formView.shadowRoot.querySelector('#edit_1')
    rowEdit.click()
    sinon.assert.callCount(onSettingEvent, 2);

    const rowDelete = formView.shadowRoot.querySelector('#delete_1')
    rowDelete.click()
    sinon.assert.callCount(onSettingEvent, 3);

    const delete_data = {
      ...Items.args.data,
      current: {
        ...Items.args.data.current,
        template: {
          ...Items.args.data.current.template,
          view: {
            ...Items.args.data.current.template.view,
            items: {
              ...Items.args.data.current.template.view.items,
              actions: {
                ...Items.args.data.current.template.view.items.actions,
                delete: null
              }
            }
          }
        }
      }
    }
    element = await fixture(Template({
      ...Items.args, data: delete_data
    }));
    form = element.querySelector('#setting_form');
    expect(form).to.exist;

  })

  it('renders in the Log state', async () => {
    const element = await fixture(Template({
      ...Log.args
    }));
    const form = element.querySelector('#setting_form');
    expect(form).to.exist;

  })

})