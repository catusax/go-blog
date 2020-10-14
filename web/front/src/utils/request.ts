import {extend }from 'umi-request'

let devserver = process.env.url;

const request = extend ({
    prefix: devserver,
    credentials: 'include',
})


export default request;