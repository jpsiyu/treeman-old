import VueRouter from 'vue-router'
import App from "./entrance/App"
import Miss from './entrance/Miss'

const routes = [
    { path: "/", component: App },
    { path: "*", component: Miss },
]

const createRouter = () => {
    return new VueRouter({
        mode: "history",
        routes,
    })
}

export { createRouter }