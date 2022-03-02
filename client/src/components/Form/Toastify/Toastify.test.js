import Toastify from './Toastify'
import { toast } from 'react-toastify';

jest.mock("react-toastify");

describe('Toastify', () => {

  beforeEach(() => {
    toast.mockReturnValue({
      error: jest.fn(),
      warning: jest.fn(),
      success: jest.fn(),
      info: jest.fn(),
    })
  });
  
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('showToast', () => {
    Toastify({ toastType: "error", message: "message", toastTime: 0 })
    Toastify({ toastType: "warning", message: "message", toastTime: 5000 })
    Toastify({ toastType: "success", message: "message", toastTime: 0 })
    Toastify({ toastType: "info", message: "message", toastTime: 3000 })
    Toastify({ toastType: "default", message: "message", toastTime: 0 })
  });
})