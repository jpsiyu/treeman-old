import axioswrap from './axioswrap'

const getToken = (user, password) => {
    return axioswrap.get(`/pubapi/login?user=${user}&password=${password}`)
}

const genPerson = (name, age, gender) => {
    const ajaxData = { name, age, gender }
    return axioswrap.post("/api/genperson", ajaxData)
}

const getAllPerson = () => {
    return axioswrap.get("/api/allperson")
}

const updatePerson = (id, name, age, gender) => {
    const ajaxData = {id, name, age, gender}
    return axioswrap.put("/api/updateperson", ajaxData)
}

const delPerson = (id) => {
    return axioswrap.put("/api/deleteperson", { id })
}

const getPersonByName = (name) => {
    return axioswrap.get(`/api/findperson?name=${name}`)
}

const getRecord = (id) => {
    return axioswrap.get(`/api/record?id=${id}`)
}

const delRecord = (id) => {
    const ajaxData = { id };
    return axioswrap.put("/api/deleterecord", ajaxData)
}

const addRecord = (id, detail, comment) => {
    const ajaxData = { detail, comment, id }
    return axioswrap.post("/api/addrecord", ajaxData)
}

const updateRecord = (id, detail, comment) => {
    const ajaxData = { detail, comment, id }
    return axioswrap.put("/api/updaterecord", ajaxData)
}

export default {
    getToken,
    genPerson,
    getAllPerson,
    delPerson,
    updatePerson,
    getPersonByName,
    getRecord,
    delRecord,
    addRecord,
    updateRecord,
}