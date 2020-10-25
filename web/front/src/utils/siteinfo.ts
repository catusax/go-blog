import request from './request'
export default async function getinfo() {
    let resp:siteinfo = await request("/api/public/info")
    return resp
}