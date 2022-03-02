import 'sanitize.css/sanitize.css';
import '../src/styles/theme.css';
import '../src/styles/global.css';

// Import CSS reset and Global Styles
export const parameters = {
  actions: { argTypesRegex: "^on[A-Z].*" },
  controls: {
    matchers: {
      color: /(background|color)$/i,
      date: /Date$/,
    },
  },
}