import restwrap from './restwrap'

const getToken = () => {
    return restwrap.get("/pubapi/login")
}

const genPerson = (name, age, gender) => {
    const ajaxData = { name, age, gender }
    return restwrap.post("/api/genperson", ajaxData)
}

const getAllPerson = () => {
    return restwrap.get("/api/allperson")
}

const delPerson = (id) => {
    return restwrap.put("/api/deleteperson", { id })
}

const getPersonByName = (name) => {
    return restwrap.get(`/api/findperson?name=${name}`)
}

const getRecord = (id) => {
    return restwrap.get(`/api/record?id=${id}`)
}

const delRecord = (id) => {
    const ajaxData = { id };
    return restwrap.put("/api/deleterecord", ajaxData)
}

const addRecord = (id, detail, comment) => {
    const ajaxData = { detail, comment, id }
    return restwrap.post("/api/addrecord", ajaxData)
}

const updateRecord = (id, detail, comment) => {
    const ajaxData = { detail, comment, id }
    return restwrap.put("/api/updaterecord", ajaxData)
}

export default {
    getToken,
    genPerson,
    getAllPerson,
    delPerson,
    getPersonByName,
    getRecord,
    delRecord,
    addRecord,
    updateRecord,
}