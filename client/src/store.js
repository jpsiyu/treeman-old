import Vuex from 'vuex'
import { MacroGender } from './macro'

const person = {
    state: {
        pageAddOn: false,
        members: [
            { id: 1, name: "张三", age: 18, gender: MacroGender.Male, living: "广州", birthPlace: "佛山" },
            { id: 2, name: "李四", age: 29, gender: MacroGender.Female, living: "广州", birthPlace: "武汉" },
            { id: 3, name: "王五", age: 39, gender: MacroGender.Male, living: "广州", birthPlace: "武汉" },
        ]
    },
    mutations: {
        showAdd: function (state) {
            state.pageAddOn = true
        },
        hideAdd: function (state) {
            state.pageAddOn = false
        },
        addMember: function (state, payload) {
            const member = Object.assign({ id: state.members.length + 1 }, payload)
            state.members.push(member)
        }

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