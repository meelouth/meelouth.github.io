import axios from "axios";

export const HTTP = axios.create({
    baseURL: 'http://127.0.0.1:8081',
    headers: {
        'Content-Type' : 'application/json'
    }
});

export default HTTP;