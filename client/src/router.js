import VueRouter from 'vue-router'
import App from "./entrance/App"
import PageRecord from "./record/PageRecord"
import Miss from './entrance/Miss'
import Login from './login/Login'

const routes = [
    { path: "/", component: App },
    { path: "/login", component: Login },
    { path: "/record", component: PageRecord },
    { path: "*", component: Miss },
]

const createRouter = () => {
    return new VueRouter({
        mode: "history",
        routes,
    })
}

export { createRouter }