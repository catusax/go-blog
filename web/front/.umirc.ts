import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
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
        },
        {
          path: '/archives',
          component: '@/pages/archive/archive'
        },
        {
          path: '/archives/page/:page',
          component: '@/pages/archive/archive'
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
        }
      ],
    },
  ],
});
