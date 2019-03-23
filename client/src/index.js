import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
import VueForm from 'vue-form'

import { createStore } from './store'
import { createRouter } from './router'

// set router
Vue.use(VueRouter)
const router = createRouter()

// set store
Vue.use(Vuex)
const store = createStore()

// set form validator
Vue.use(VueForm)

// start vue instance
new Vue({
    router,
    store,
    template: '<router-view></router-view>'
}).$mount("#root")

