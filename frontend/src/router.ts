import { createRouter, createWebHashHistory } from 'vue-router'
import {
    Headset,
    Management,
    Notebook,
    VideoPause,
  } from '@element-plus/icons-vue'
  
export const routes = [
    { path: '/', icon: Management, component: () => import('@views/BookShelf.vue') },
    { path: '/BookShelf', name: "书架管理", type: 'menu', icon: Management, component: () => import('@views/BookShelf.vue') },
    { path: '/BookSites', name: "书源管理", type: 'menu',  icon: Notebook, component: () => import('@views/BookSites.vue') },
    { path: '/book2', name: "书籍制作", type: 'menu', icon: Notebook, component: () => import('@components/HelloWorld.vue') },
    { path: '/book3', name: "音频制作", type: 'menu', icon: Headset, component: () => import('@views/SideMain.vue') },
    { path: '/book4', name: "音频管理", type: 'menu', icon: VideoPause, component: () => import('@components/HelloWorld.vue') },
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes, // `routes: routes` 的缩写
})