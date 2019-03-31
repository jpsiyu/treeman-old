import Vuex from 'vuex'
import { MacroGender } from './macro'

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

const behavior = {
    state: {},
    mutations: {},
}

const createStore = () => {
    return new Vuex.Store({
        modules: {
            person,
            behavior,
        }
    })
}

export { createStore }