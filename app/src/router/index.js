import Vue from 'vue'
import Router from 'vue-router'
import DownLoad from '@/view/download'   // DownLoad组件中包含了所有页面公共的侧边栏

Vue.use(Router)

// 不需要用户权限的静态路由, path: 路由的路径,redirect: 重定向路由的路径
// component: 跳转的组件, meta:元数据{title: 页面的标题, icon: 侧边栏的按钮,
// hidden: 该路由是否在侧边栏被隐藏, 默认为false}
const constantRoutes = [
  {
    path: '/',
    redirect: '/login',
    meta: { hidden: true }
  },
  {
    path: '/login',  // 登陆界面
    name: 'Login',
    component: () => import('@/view/login'),
    meta: { title: '登陆', hidden: true }
  },
  {
    path: '/homePage',
    name: 'HomePage',
    component: DownLoad,
    meta: { title: 'GDB', icon: 'el-icon-coin', hidden: false },
    redirect:'/index',
    children:[
      {
        path:'/index',
        component: ()=>import('@/view/firstPage'),
        meta: { title: 'GDB', icon: 'el-icon-coin', hidden: false }
      }
    ]
  },
  {
    path: '/search',
    component: DownLoad,
    name: 'Search',
    meta: { title: '资源管理', icon: 'el-icon-menu', hidden: false },
    children: [
      {
        path: '/groups',
        name: 'Group',
        component: () => import('@/view/group'),
        meta: { title: '分组管理', icon: 'el-icon-film', hidden: false },
      }, 
      {
        path: '/calc',
        name: 'Calc',
        component: () => import('@/view/calc'),
        meta: { title: '二次计算', icon: 'el-icon-cpu', hidden: false }
      }
    ]
  }
]

const asyncRoutes = [
  {
    path: '/user',
    component: DownLoad,
    name: 'User',
    meta: { title: '用户中心', role: ['developer', 'super_user', 'common_user'], icon: 'el-icon-user-solid', hidden: false },
    children: [
      {
        path: '/document',
        name: 'Document',
        component: () => import('@/view/user/document'),
        meta: { title: '系统文档', role: ['developer', 'super_user', 'common_user'], icon: 'el-icon-document', hidden: false }  // 路由元数据
      },
      {
        path :'/log',
        name: 'Log',
        component: ()=>import('@/view/log'),
        meta: { title: '运行日志', role: ['developer', 'super_user', 'common_user'], icon: 'el-icon-document', hidden: false }  // 路由元数据
      },
      {
        path: '/userManagement',
        name: 'UserManagement',
        component: () => import('@/view/user/userManagement'),
        meta: { title: '用户管理', role: ['developer', 'super_user'], icon: 'el-icon-user-solid', hidden: false }  // 路由元数据
      },
    ]
  },
  {
    path: '*',
    name: 'Page404',
    component: () => import('@/view/404'),
    meta: { role: ['developer', 'super_user', 'common_user'], hidden: true }
  }
]

// 解决VUE路由跳转出现Redirected when going from "/xxx" to "/yyy" via a navigation guard.报错
const originalPush = Router.prototype.push
Router.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

const router = new Router({
  routes: constantRoutes,
  mode: 'history'
})

const createRouter = () => {
  return new Router({
    routes: constantRoutes,
    mode: 'history'
  })
}

// addRoutes 方法仅仅是帮你注入新的路由，并没有帮你剔除其它路由,所以需要手动清空路由
// 为了确保其完成，需要使用async异步操作
const resetRouter = async () => {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}


export default router
export { constantRoutes, asyncRoutes, resetRouter }