import { defineConfig } from 'umi';
const title = 'Just4fun'
export default defineConfig({
  styles:[
    //'https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@10.2.1/build/styles/default.min.css'
  ],
  scripts:[
    'https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@10.2.1/build/highlight.min.js',
  ],
  nodeModulesTransform: {
    type: 'none',
  },
  define:{
  },
  routes: [
    {
      path: '/',
      component: '@/layouts/index',
      routes: [
        {
          path: '/',
          component: '@/pages/index',
        },
        {
          path: '/page/:page',
          component: '@/pages/index', 
          title: title
        },
        {
          path: '/archives',
          component: '@/pages/archive/archive',
        },
        {
          path: '/archives/page/:page',
          component: '@/pages/archive/archive',
        },
        {
          path: '/post/:id',
          component: '@/pages/post/post'
        },
        {
          path: '/tag/:name',
          component: '@/pages/tag/tag'
        },
        {
          path: '/tag/:name/page/:page',
          component: '@/pages/tag/tag'
        },
        {
          path: '/pages/:pageid',
          component:'@/pages/pages/pages'
        },
        {
          component:'@/pages/404'
        }
      ],
    },
  ],
});
