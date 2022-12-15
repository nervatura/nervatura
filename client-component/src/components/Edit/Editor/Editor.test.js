import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './edit-editor.js';
import { Template, Default, New, Report, PrintQueue, Notes, Item, Meta, View } from  './Editor.stories.js';

describe('Edit-Editor', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEditEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    const btnForm = editor.shadowRoot.querySelector('#btn_form')
    btnForm.click()
    sinon.assert.callCount(onEditEvent, 1);

    const btnFieldvalue = editor.shadowRoot.querySelector('#btn_fieldvalue')
    btnFieldvalue.click()
    sinon.assert.callCount(onEditEvent, 2);

  })

  it('renders in the New state', async () => {
    const element = await fixture(Template({
      ...New.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;
  })

  it('renders in the Report state', async () => {
    const element = await fixture(Template({
      ...Report.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;
  })

  it('renders in the PrintQueue state', async () => {
    const element = await fixture(Template({
      ...PrintQueue.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;
  })

  it('renders in the Notes state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Notes.args, onEditEvent
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;

    editor._onNoteBlur({ target: { _value: "", value: "abc" } })
    sinon.assert.callCount(onEditEvent, 1);
    editor._onNoteBlur({ target: { _value: "abc", value: "abc" } })
    sinon.assert.callCount(onEditEvent, 1);
  })

  it('renders in the Item state', async () => {
    const element = await fixture(Template({
      ...Item.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;
  })

  it('renders in the Meta state', async () => {
    const element = await fixture(Template({
      ...Meta.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;
  })

  it('renders in the View state', async () => {
    const element = await fixture(Template({
      ...View.args
    }));
    const editor = element.querySelector('#editor');
    expect(editor).to.exist;
  })

})