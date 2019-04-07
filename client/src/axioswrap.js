import axios from 'axios'
import qs from 'qs'

const get = (path) => {
    return axios.get(path, getOptions())
        .then(res => {
            return dataMiddleware(res)
        })
        .catch(handlePromissError)
}

const put = (path, data) => {
    return axios.put(path, qs.stringify(data), getOptions())
        .then(res => {
            return dataMiddleware(res)
        })
        .catch(handlePromissError)
}

const post = (path, data) => {
    return axios.post(path, qs.stringify(data), getOptions())
        .then(res => {
            return dataMiddleware(res)
        })
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

const dataMiddleware = (response) => {
    const serverData = response.data
    switch (serverData.code) {
        case 1:
            window.vm.$router.push("/login")
            break
    }
    return serverData
}

const handlePromissError = (err) => {
    console.error(err)
}

export default {
    get,
    put,
    post
}