import axios from 'axios'
import qs from 'qs'

const get = (path) => {
    return axios.get(path, getOptions())
        .then(res => {
            return handleResponse(res)
        })
        .catch(handlePromissError)
}

const put = (path, data) => {
    return axios.put(path, qs.stringify(data), getOptions())
        .then(res => {
            return handleResponse(res)
        })
        .catch(handlePromissError)
}

const post = (path, data) => {
    return axios.post(path, qs.stringify(data), getOptions())
        .then(res => {
            return handleResponse(res)
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

const handleResponse = (response) => {
    const serverData = response.data
    handleServerError(serverData.code)
    return serverData
}

const handleServerError = (code) => {
    switch (code) {
        case 0:
            break
        case 1:
            window.vm.$router.push("/login")
            break
    }
    if (code != 0)
        console.error(`server error code: ${code}`)
}

const handlePromissError = (err) => {
    console.error(err)
}

export default {
    get,
    put,
    post
}