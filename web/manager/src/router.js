import { createWebHistory,createRouter } from "vue-router";

const routes = [
    {
        path:"/login",
        component:() => import(/* webpackChunkName:"login" */'./view/authentication/login.vue'),
    },
    {
        path:"/",
        redirect:"/login",
    },
    {
        path:"/main",
        component:() => import(/* webpackChunkName:"main" */'./view/main.vue'),
        redirect:'/main/blog',
        children:[
            {
                path:"blog",
                component:() => import(/* webpackChunkName:"blog" */'./view/BlogManager/Blog.vue'),
            },
            {
                path:"edit",
                component:() => import(/* webpackChunkName:"edit" */'./view/edit/edit.vue'),
            },
            {
                path:`watch/:id`,
                name:"watch",
                component:() => import(/* webpackChunkName:"watch" */'./view/BlogManager/watch.vue'),
            },
            {
                path:"photo",
                component:() => import(/* webpackChunkName:"photo" */'./view/BlogManager/blogPhoto.vue'),
            },
            {
                path:"tagManager",
                component:() => import(/* webpackChunkName:"tagManager" */'./view/BlogManager/TagManager.vue'),
            },
            {
                path:'friendManager',
                component:() => import(/* webpackChunkName:"friendManager" */'./view/BlogManager/FriendManager.vue'),
            },
            {
                path:"setting",
                component:() => import(/* webpackChunkName:"setting" */'./view/configCenter/set.vue'),
            },
            {
                path:"user",
                component:() => import(/* webpackChunkName:"user" */'./view/usermanager/userManager.vue'),
            }
        ]
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})


export default router;  