import axios from 'axios'
import qs from 'qs'

const genPerson = (name, age, gender) => {
    const ajaxData = { name, age, gender }
    return axios.post("/api/genperson", qs.stringify(ajaxData))
}

const getAllPerson = () => {
    return axios.get("/api/allperson")
}

const delPerson = (id) => {
    return axios.put("/api/deleteperson", qs.stringify({ id}))
}

const getRecord = (id) => {
    return axios.get(`/api/record?id=${id}`)
}

const delRecord = (id) => {
    const ajaxData = { id };
    return axios.put("/api/deleterecord", qs.stringify(ajaxData))
}

const addRecord = (id, detail, comment) => {
    const ajaxData = { detail, comment, id }
    return axios.post("/api/addrecord", qs.stringify(ajaxData))
}

const updateRecord = (id, detail, comment) => {
    const ajaxData = { detail, comment, id }
    return axios.put("/api/updaterecord", qs.stringify(ajaxData))
}

export default {
    genPerson,
    getAllPerson,
    delPerson,
    getRecord,
    delRecord,
    addRecord,
    updateRecord,
}