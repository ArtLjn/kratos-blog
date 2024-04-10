import { createWebHistory,createRouter } from "vue-router";

const routes = [
    {
        path:"/login",
        component:() => import(/* webpackChunkName:"login" */'./components/authentication/login.vue'),
    },
    {
        path: "/test",
        component:() => import(/* webpackChunkName:"test" */'./components/edit/test.vue')
    },
    {
        path:"/",
        redirect:"/login",
    },
    {
        path:"/main",
        component:() => import(/* webpackChunkName:"main" */'./components/main.vue'),
        redirect:'/main/blog',
        children:[
            {
                path:"blog",
                component:() => import(/* webpackChunkName:"blog" */'./components/BlogManager/Blog.vue'),
            },
            {
                path:"edit",
                component:() => import(/* webpackChunkName:"edit" */'./components/edit/edit.vue'),
            },
            {
                path:`watch/:id`,
                name:"watch",
                component:() => import(/* webpackChunkName:"watch" */'./components/BlogManager/watch.vue'),
            },
            {
                path:"photo",
                component:() => import(/* webpackChunkName:"photo" */'./components/BlogManager/blogPhoto.vue'),
            },
            {
                path:"tagManager",
                component:() => import(/* webpackChunkName:"tagManager" */'./components/BlogManager/TagManager.vue'),
            },
            {
                path:'friendManager',
                component:() => import(/* webpackChunkName:"friendManager" */'./components/BlogManager/FriendManager.vue'),
            },
            {
                path:"setting",
                component:() => import(/* webpackChunkName:"setting" */'./components/setting/set.vue'),
            }
        ]
    }
]

const router = createRouter({
    history:createWebHistory(),
    routes
})


export default router;  