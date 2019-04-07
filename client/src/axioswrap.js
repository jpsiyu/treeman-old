import axios from 'axios'
import qs from 'qs'

const get = (path) => {
    return axios.get(path, getOptions())
        .catch(handlePromissError)
}

const put = (path, data) => {
    return axios.put(path, qs.stringify(data), getOptions())
        .catch(handlePromissError)
}

const post = (path, data) => {
    return axios.post(path, qs.stringify(data), getOptions())
        .catch(handlePromissError)
}

const getOptions = () => {
    if (!window.vm) return {}
    return {
        headers: {
            Authorization: window.vm.$store.state.app.tokenStr
        }
    }
}

const handlePromissError = (err) => {
    const res = err.response
    switch(res.status){
        case 401:
            window.vm.$router.push("/login")
            break
    }
}

export default {
    get,
    put,
    post
}