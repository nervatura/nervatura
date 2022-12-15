import merge from 'deepmerge';
// use createSpaConfig for bundling a Single Page App
import { createSpaConfig } from '@open-wc/building-rollup';

// use createBasicConfig to do regular JS to JS bundling
// import { createBasicConfig } from '@open-wc/building-rollup';

import copy from 'rollup-plugin-copy';
import replace from '@rollup/plugin-replace';
import packageData from './package.json' assert { type: 'json' };

const baseConfig = createSpaConfig({
  // use the outputdir option to modify where files are output
  outputDir: 'dist',

  // if you need to support older browsers, such as IE11, set the legacyBuild
  // option to generate an additional build just for this browser
  // legacyBuild: true,

  // development mode creates a non-minified build for debugging or development
  developmentMode: process.env.ROLLUP_WATCH === 'true',

  // set to true to inject the service worker registration into your index.html
  injectServiceWorker: false,
  workbox: false,
  html: {
    extractAssets: false,
    minify: true
  }
});

export default merge(baseConfig, {
  // if you use createSpaConfig, you can use your index.html as entrypoint,
  // any <script type="module"> inside will be bundled by rollup
  input: './index.html',
  output: {
    entryFileNames: 'client-[hash].js',
    chunkFileNames: '[name]-[hash].js',
    // assetFileNames: '[hash][extname]',
    sourcemap: false,
    manualChunks: {
      lit: ['lit'],
      icon: ['./src/components/Form/Icon/form-icon.js'],
      locales: ['./src/config/locales.js'],
      // app: ['./src/controllers/AppController.js'],
    },

  },
  plugins: [
    replace({
      preventAssignment: true,
      __SERVER__: process.env.APP_CONFIG,
      __VERSION__: packageData.version
    }),
    copy({
      targets: [
        { src: 'assets/robots.txt', dest: 'dist' },
        { src: 'assets/meta.json', dest: 'dist' },
        { src: 'manifest.json', dest: 'dist' },
        { src: 'favicon.svg', dest: 'dist' }
      ],
    }),
  ]
  // alternatively, you can use your JS as entrypoint for rollup and
  // optionally set a HTML template manually
  // input: './app.js',
});
