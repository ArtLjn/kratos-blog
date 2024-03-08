import { createWebHistory,createRouter } from "vue-router";
import Blog from './components/BlogManager/Blog.vue'
import edit from './components/edit/edit.vue'
import main from './components/main.vue'
import watch from './components/BlogManager/watch.vue'
import login from './components/authentication/login.vue'
import bloPhoto from './components/BlogManager/blogPhoto.vue'
import TagManager from './components/BlogManager/TagManager.vue'
import FriendManager from './components/BlogManager/FriendManager.vue'
const routes = [
    {
        path:"/login",
        component:login
    },
    {
        path:"/",
        component:main,
        redirect:"/login",
    },
    {
        path:"/main",
        component:main,
        redirect:'/main/blog',
        children:[
            {
                path:"blog",
                component:Blog
            },
            {
                path:"edit",
                component:edit
            },
            {
                path:`watch/:id`,
                name:"watch",
                component:watch
            },
            {
                path:"photo",
                component:bloPhoto
            },
            {
                path:"tagManager",
                component:TagManager
            },
            {
                path:'friendManager',
                component:FriendManager
            },
            {
                path:"talkManager",
                component:()=>import(/* webpackChunkName:"frinedLink" */'./components/edit/TalkManager.vue'),
            },
        ]
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})
export default router;  