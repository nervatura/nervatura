module.exports = {
  stories: [
    "../src/**/*.stories.js", 
    "../src/**/*.stories.@(js|jsx|ts|tsx)", 
    "../src/components/**/**/*.stories.js", 
    "../src/components/**/**/*.stories.@(js|jsx|ts|tsx)"
  ],
  addons: [
    "@storybook/addon-links", 
    "@storybook/addon-essentials", 
    "@storybook/preset-create-react-app"
  ],
  core: {
    builder: "webpack5"
  },
  staticDirs: [
    '../public',
  ]
};