const components = require('prismjs/components');
const allLanguages = Object.keys(components.languages).filter((item) => item !== 'meta');
module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],
  "plugins": [
    ["@babel/plugin-transform-runtime", {
      "corejs": false, // 默认值，可以省略
      "helpers": true,
      "regenerator": true,
      "useESModules": false
    }]
  ],
  plugins: [
    [
      'prismjs',
      {
        languages: allLanguages,
      },
    ],
  ],
}
