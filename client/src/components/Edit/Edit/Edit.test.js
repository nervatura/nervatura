import { fixture, expect } from '@open-wc/testing';

import './client-edit.js';
import { Template, Default, New } from  './Edit.stories.js';

describe('Edit-Edit', () => {
  it('renders in the Default state', async () => {
    const element = await fixture(Template({
      ...Default.args
    }));
    const edit = element.querySelector('#edit');
    expect(edit).to.exist;

    const totalFrm = edit.modalFormula({})
    expect(totalFrm).to.exist;

    const modalSelector = edit.modalSelector({})
    expect(modalSelector).to.exist;

    const modalTrans = edit.modalTrans({})
    expect(modalTrans).to.exist;

    const modalStock = edit.modalStock({})
    expect(modalStock).to.exist;

    const modalShipping = edit.modalShipping({})
    expect(modalShipping).to.exist;

    const modalReport = edit.modalReport({})
    expect(modalReport).to.exist;
  })

  it('renders in the New state', async () => {
    const element = await fixture(Template({
      ...New.args
    }));
    const edit = element.querySelector('#edit');
    expect(edit).to.exist;
  })

})
