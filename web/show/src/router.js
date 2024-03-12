import { createRouter, createWebHistory } from 'vue-router';
import main from './components/blog/main.vue'
const routes = [
    {
        path:"/main",
        component:main,
        meta:{
            title:'红豆南墙 | 个人博客'
        },
        redirect:"/",
        children:[
            {
                name:"blogmian",
                path:"/",
                component:()=>import(/* webpackChunkName:"blogmain" */'./components/blog/blogmain.vue'),
                meta:{
                    title:'红豆南墙 | 个人博客'
                }
            },
            {
                path:`/watch/:id`,
                name:"watch",
                component:()=>import(/* webpackChunkName:"watch" */'./components/blog/watch.vue'),
                meta:{
                    title:'文章 | 红豆南墙 | 个人博客'
                }
            },
            {
                path:"/about",
                component:()=>import(/* webpackChunkName:"about" */'./components/blog/followme.vue'),
                meta:{
                    title:'关于 | 红豆南墙 | 个人博客'
                }
            },
            {
                path:"/archives",
                component:()=>import(/* webpackChunkName:"archives" */'./components/blog/archives.vue'),
                meta:{
                    title:'归档 | 红豆南墙 | 个人博客'
                }
            },
            {
                path:"/message",
                component:()=>import(/* webpackChunkName:"message" */'./components/blog/message.vue'),
                meta:{
                    title:'留言 | 红豆南墙 | 个人博客'
                }
            },
            {
                path:"/photo",
                component:()=>import(/* webpackChunkName:"photo" */'./components/blog/photo.vue'),
                meta:{
                    title:'相册 | 红豆南墙 | 个人博客'
                }
            },
            {
                path:"/tag",
                component:()=>import(/* webpackChunkName:"tag" */'./components/blog/tag.vue'),
                meta:{
                    title:'标签 | 红豆南墙 | 个人博客'
                }
            },
            {
                path:"/friendlink",
                component:()=>import(/* webpackChunkName:"frinedLink" */'./components/blog/friendLink.vue'),
                meta:{
                    title:'友链 | 红豆南墙 | 个人博客'
                }
            },
        ]
    },
    {
        path:'/404',
        component: ()=>import(/* webpackChunkName:"404" */'./components/404.vue'),
        meta:{
            title:"404"
        },
    },
    {
        path:'/:pathMatch(.*)*',
        redirect:'/404'
    },
    {
        path:"/test",
        component:()=>import(/* webpackChunkName:"test" */'./components/blog/test.vue'),
    },
]
const router = createRouter({
    history:createWebHistory(),
    routes
})
export default router;