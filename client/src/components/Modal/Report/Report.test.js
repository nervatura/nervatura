import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './modal-report.js';
import { Template, Default, Disabled } from  './Report.stories.js';

describe('Report', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onModalEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onModalEvent
    }));
    const report = element.querySelector('#report');
    expect(report).to.exist;

    const closeIcon = report.shadowRoot.querySelector('#closeIcon')
    closeIcon.click()
    sinon.assert.callCount(onModalEvent, 1);

    const btnPrint = report.shadowRoot.querySelector('#btn_print')
    btnPrint.click()
    sinon.assert.callCount(onModalEvent, 2);

    const btnPdf = report.shadowRoot.querySelector('#btn_pdf')
    btnPdf.click()
    sinon.assert.callCount(onModalEvent, 3);

    const btnXml = report.shadowRoot.querySelector('#btn_xml')
    btnXml.click()
    sinon.assert.callCount(onModalEvent, 4);

    const btnPrintqueue = report.shadowRoot.querySelector('#btn_printqueue')
    btnPrintqueue.click()
    sinon.assert.callCount(onModalEvent, 5);

    const template = report.shadowRoot.querySelector('#template');
    template._onInput({ target: { value: "ntr_invoice_fi" } })
    expect(template.value).to.equal("ntr_invoice_fi");

    const orient = report.shadowRoot.querySelector('#orient');
    orient._onInput({ target: { value: "landscape" } })
    expect(orient.value).to.equal("landscape");

    const size = report.shadowRoot.querySelector('#size');
    size._onInput({ target: { value: "a5" } })
    expect(size.value).to.equal("a5");

    const copy = report.shadowRoot.querySelector('#copy');
    copy._onInput({ target: { valueAsNumber: 12 } })
    expect(copy.value).to.equal(12);
  })

  it('renders in the Disabled state', async () => {
    const element = await fixture(Template({
      ...Disabled.args
    }));
    const report = element.querySelector('#report');
    expect(report).to.exist;
  })

})