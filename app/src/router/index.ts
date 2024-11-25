import {
  createRouter,
  createMemoryHistory,
} from 'vue-router'

/** 根路由: / */
export const ROOT_ROUTE: AuthRoute.Route = {
  name: 'root',
  path: '/',
  redirect: '/home',
  meta: {
    title: 'Root'
  }
}

/** 固定的路由 */
export const routes: AuthRoute.Route[] = [
  ROOT_ROUTE,
  {
    name: 'login',
    path: '/login',
    component: 'self',
    props: (route) => {
      const moduleType =
        (route.params.module as UnionKey.LoginModule) || 'pwd-login'
      return {
        module: moduleType
      }
    },
    meta: {
      title: '登录',
      // dynamicPath: `/login/:module(${getLoginModuleRegExp()})?`,
      dynamicPath: `/login`,
      singleLayout: 'blank'
    }
  },
  {
    name: 'home',
    path: '/home',
    component: ()=>import('@/views/home/home.vue'),
    meta: {
      title: '主页',
      singleLayout: 'blank'
    }
  },
  {
    name: 'constant-page',
    path: '/constant-page',
    component: 'self',
    meta: {
      title: '固定页面',
      singleLayout: 'blank'
    }
  },
  {
    name: '403',
    path: '/403',
    component: 'self',
    meta: {
      title: '无权限',
      singleLayout: 'blank'
    }
  },
  {
    name: '404',
    path: '/404',
    component: 'self',
    meta: {
      title: '未找到',
      singleLayout: 'blank'
    }
  },
  {
    name: '500',
    path: '/500',
    component: 'self',
    meta: {
      title: '服务器错误',
      singleLayout: 'blank'
    }
  },
  // 匹配无效路径的路由
  {
    name: 'not-found',
    path: '/:pathMatch(.*)*',
    component: 'blank',
    meta: {
      title: '未找到',
      singleLayout: 'blank'
    }
  }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export async function setupRouter(app: App) {
  app.use(router)
  await router.isReady()
}
