import Vuex from 'vuex'

const app = {
    state: {
        tokenStr: ""
    },
    mutations: {
        setTokenStr: function(state, payload){
            state.tokenStr = payload
        }
    }
}

const person = {
    state: {
        pageAddOn: false,
        members: []
    },
    mutations: {
        initMembers: function (state, payload) {
            state.members = payload
        },
    },
}

const createStore = () => {
    return new Vuex.Store({
        modules: {
            app,
            person,
        }
    })
}

export { createStore }