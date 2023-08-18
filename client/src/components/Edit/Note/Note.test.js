import { fixture, expect } from '@open-wc/testing';
import sinon from 'sinon'

import './edit-note.js';
import { Template, Default, Empty, ReadOnly } from  './Note.stories.js';

describe('Edit-Note', () => {
  afterEach(() => {
    // Restore the default sandbox here
    sinon.restore();
  });

  it('renders in the Default state', async () => {
    const onEditEvent = sinon.spy()
    const element = await fixture(Template({
      ...Default.args, onEditEvent
    }));
    const note = element.querySelector('#note');
    expect(note).to.exist;

    const patternDefault = note.shadowRoot.querySelector('#btn_pattern_default')
    patternDefault.click()
    sinon.assert.callCount(onEditEvent, 1);

    const patternLoad = note.shadowRoot.querySelector('#btn_pattern_load')
    patternLoad.click()
    sinon.assert.callCount(onEditEvent, 2);

    const patternSave = note.shadowRoot.querySelector('#btn_pattern_save')
    patternSave.click()
    sinon.assert.callCount(onEditEvent, 3);

    const patternNew = note.shadowRoot.querySelector('#btn_pattern_new')
    patternNew.click()
    sinon.assert.callCount(onEditEvent, 4);

    const patternDelete = note.shadowRoot.querySelector('#btn_pattern_delete')
    patternDelete.click()
    sinon.assert.callCount(onEditEvent, 5);

    const selPattern = note.shadowRoot.querySelector('#sel_pattern')
    selPattern._onInput({ target: { value: "2" } })
    sinon.assert.callCount(onEditEvent, 6);

    note._onInput()
    sinon.assert.callCount(onEditEvent, 6);

    const bold = note.shadowRoot.querySelector('#btn_bold')
    bold.click()

    const italic = note.shadowRoot.querySelector('#btn_italic')
    italic.click()

    note._onLostFocus()

  })

  it('renders in the ReadOnly state', async () => {
    const element = await fixture(Template({
      ...ReadOnly.args
    }));
    const note = element.querySelector('#note');
    expect(note).to.exist;
  })

  it('renders in the Empty state', async () => {
    const element = await fixture(Template({
      ...Empty.args
    }));
    const note = element.querySelector('#note');
    expect(note).to.exist;
  })

})