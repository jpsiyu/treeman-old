import VueRouter from 'vue-router'
import App from "./entrance/App"
import PageRecord from "./record/PageRecord"
import Miss from './entrance/Miss'

const routes = [
    { path: "/", component: App },
    { path: "/record", component: PageRecord},
    { path: "*", component: Miss },
]

const createRouter = () => {
    return new VueRouter({
        mode: "history",
        routes,
    })
}

export { createRouter }