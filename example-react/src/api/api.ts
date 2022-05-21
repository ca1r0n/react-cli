import axios, {AxiosInstance, AxiosRequestConfig} from "axios";
import {config} from "../config/config";


const apiConfig: AxiosRequestConfig = {
    baseURL: config.api
}
const api: AxiosInstance = axios.create(apiConfig)

api.interceptors.response.use()

api.interceptors.request.use()

export {api}