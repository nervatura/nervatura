import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './nt-row.js';

import { Template, Default, FlipStringOn, FlipStringOff, FlipTextOn, FlipTextOff, FlipImageOn, FlipImageOff,
  FlipChecklistOn, FlipChecklistOff, Field, Reportfield, ReportfieldEmpty, Fieldvalue,
  Col2, Col3, Col4, Missing, Label } from './NtRow.stories.js';

  describe('NtField', () => {
    afterEach(() => {
      // Restore the default sandbox here
      sinon.restore();
    });
  
    it('renders in the Default state', async () => {
      const element = await fixture(Template({
        ...Default.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the FlipStringOn state', async () => {
      const onEdit = sinon.spy()
      const element = await fixture(Template({
        ...FlipStringOn.args, onEdit
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

      const checkbox = testRow.shadowRoot.querySelector('#checkbox_title');
      checkbox.click()
      sinon.assert.calledOnce(onEdit);
    })

    it('renders in the FlipStringOff state', async () => {
      const element = await fixture(Template({
        ...FlipStringOff.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the FlipTextOn state', async () => {
      const element = await fixture(Template({
        ...FlipTextOn.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the FlipTextOff state', async () => {
      const element = await fixture(Template({
        ...FlipTextOff.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the FlipImageOn state', async () => {
      const onEdit = sinon.spy()
      let element = await fixture(Template({
        ...FlipImageOn.args, onEdit
      }));
      let testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

      testRow._onTextInput({ target: { value: "data:image" } })
      sinon.assert.calledOnce(onEdit);

      const fileSrc = testRow.shadowRoot.querySelector('#file_src')
      fileSrc.onChange({ value: "" })
      sinon.assert.calledTwice(onEdit);
      expect(fileSrc.value).to.equal('');

      element = await fixture(Template({
        ...FlipImageOn.args,
        values: { src: "" }
      }));
      testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

    })

    it('renders in the FlipImageOff state', async () => {
      const element = await fixture(Template({
        ...FlipImageOff.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the FlipChecklistOn state', async () => {
      const onEdit = sinon.spy()
      const element = await fixture(Template({
        ...FlipChecklistOn.args, onEdit
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

      const checklist = testRow.shadowRoot.querySelector('#checklist_border_0');
      checklist.click()
      sinon.assert.calledOnce(onEdit);
    })

    it('renders in the FlipChecklistOff state', async () => {
      const element = await fixture(Template({
        ...FlipChecklistOff.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the Field state', async () => {
      const element = await fixture(Template({
        ...Field.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the Reportfield state', async () => {
      const onEdit = sinon.spy()
      const element = await fixture(Template({
        ...Reportfield.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

      const posdate = testRow.shadowRoot.querySelector('#cb_posdate');
      posdate.click()
      sinon.assert.callCount(onEdit, 0);
    })

    it('renders in the ReportfieldEmpty state', async () => {
      const onEdit = sinon.spy()
      const element = await fixture(Template({
        ...ReportfieldEmpty.args, onEdit
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

      const curr = testRow.shadowRoot.querySelector('#cb_curr');
      curr.click()
      sinon.assert.callCount(onEdit, 1);
    })

    it('renders in the Fieldvalue state', async () => {
      const onEdit = sinon.spy()
      const element = await fixture(Template({
        ...Fieldvalue.args, onEdit
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;

      const deleteRow = testRow.shadowRoot.querySelector('#delete_sample_customer_float');
      deleteRow.click()
      sinon.assert.callCount(onEdit, 1);

      const customerFloat = testRow.shadowRoot.querySelector('#notes_sample_customer_float')
      customerFloat._onInput({ target: { value: "12", valueAsNumber: 12 } })
      sinon.assert.calledTwice(onEdit);
      expect(customerFloat.value).to.equal("12");
    })

    it('renders in the Col2 state', async () => {
      const element = await fixture(Template({
        ...Col2.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the Col3 state', async () => {
      const element = await fixture(Template({
        ...Col3.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the Col4 state', async () => {
      const element = await fixture(Template({
        ...Col4.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the Missing state', async () => {
      const element = await fixture(Template({
        ...Missing.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

    it('renders in the Label state', async () => {
      const element = await fixture(Template({
        ...Label.args
      }));
      const testRow = element.querySelector('#test_row');
      expect(testRow).to.exist;
    })

})