import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus';
import 'element-plus/theme-chalk/index.css'

import {  library } from '@fortawesome/fontawesome-svg-core';
import { faUserSecret } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import Prism from 'prismjs';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import VMdPreviewHtml from '@kangc/v-md-editor/lib/preview-html';
import VMdEditor from '@kangc/v-md-editor';
import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/preview-html.css';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
import createEmojiPlugin from '@kangc/v-md-editor/lib/plugins/emoji/index';
import '@kangc/v-md-editor/lib/plugins/emoji/emoji.css';
import createLineNumbertPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index';
import createHighlightLinesPlugin from '@kangc/v-md-editor/lib/plugins/highlight-lines/index';
import '@kangc/v-md-editor/lib/plugins/highlight-lines/highlight-lines.css';
import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index';
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css';
import createTodoListPlugin from '@kangc/v-md-editor/lib/plugins/todo-list/index';
import '@kangc/v-md-editor/lib/plugins/todo-list/todo-list.css';
import createTipPlugin from '@kangc/v-md-editor/lib/plugins/tip/index';
import '@kangc/v-md-editor/lib/plugins/tip/tip.css';

library.add(faUserSecret)
VMdEditor.use(vuepressTheme, {
    Prism,
})
VueMarkdownEditor.use(createEmojiPlugin,createLineNumbertPlugin,createHighlightLinesPlugin,
    createCopyCodePlugin,createTodoListPlugin,createTipPlugin)
router.beforeEach((to,from,next) => {
  const authData = JSON.parse(localStorage.getItem('authData'));
  const isLoggedIn = authData && authData.isLoggedIn;
  if(to.path !== '/login' && !isLoggedIn) {
    next("/login")
  } else {
    next()
  }
})
const app = createApp(App)
app.component("font-awesome-icon", FontAwesomeIcon)
app.use(router).use(ElementPlus).use(VMdEditor).use(VMdPreviewHtml).mount('#app')
