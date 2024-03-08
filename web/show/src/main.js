import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css'
import router from './router';
import * as ElIcons from "@element-plus/icons-vue";
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import {  library } from '@fortawesome/fontawesome-svg-core';
import { faUserSecret } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import emitt from './plugins/bus';
import animated from 'animate.css';
import VueWechatTitle from 'vue-wechat-title';
// markdown
import Prism from 'prismjs'; //高亮包
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js'; //主题包
import VMdEditor from '@kangc/v-md-editor'; //编辑器包
import VueMarkdownEditor from '@kangc/v-md-editor'; //编辑器包
import '@kangc/v-md-editor/lib/theme/style/vuepress.css'  ; //主题css
import '@kangc/v-md-editor/lib/plugins/emoji/emoji.css'; //表情css
import createLineNumbertPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index'; //行高
import createHighlightLinesPlugin from '@kangc/v-md-editor/lib/plugins/highlight-lines/index'; //行数高亮包
import '@kangc/v-md-editor/lib/plugins/highlight-lines/highlight-lines.css'; //高亮样式包
import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index'; //复制包
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css';//复制样式包
import createTipPlugin from '@kangc/v-md-editor/lib/plugins/tip/index';//提示插件
import '@kangc/v-md-editor/lib/plugins/tip/tip.css';//提示主题包

import VueLazyLoad from 'vue-lazyload'


library.add(faUserSecret)
VMdEditor.use(vuepressTheme,{   //编辑器注册样式
  Prism,
  extend(md) {},
  config:{
    toc:{
      includeLevel: [3, 4], //显示三，四级导航
    },
  }
})
VueMarkdownEditor.use(createLineNumbertPlugin()).use(createHighlightLinesPlugin()).use(createCopyCodePlugin()) //注册插件
.use(createTipPlugin());
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
for (const name in ElIcons) {
  app.component(name,ElIcons[name]);
}
app.component("font-awesome-icon", FontAwesomeIcon)
app.use(ElementPlus).use(router).use(emitt).use(VueWechatTitle)
.use(animated).use(VMdEditor).use(VueLazyLoad,{
  preLoad:1,
  attmept:2
})
.mount('#app')