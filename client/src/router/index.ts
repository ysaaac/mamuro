import { createRouter, createWebHistory } from 'vue-router'
import InboxView from '@/views/InboxView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'inbox',
      component: InboxView
    },
    {
      path: '/sent',
      name: 'sent_messages',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/SentMessagesView.vue')
    }
  ]
})

export default router
