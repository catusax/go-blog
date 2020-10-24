import request from './request'

var siteinfo = {
    SiteName: "BLOG",
    Gitalk: {
        Owner: 'qqwwee',
        Repo: '11223',
        ClientID: 'fghfhf3660',
        ClientSecret: 'asdfsdf',
    }
}

const getinfo = async () => {
    let data = await request("/api/public/info")
    siteinfo.SiteName = data.SiteName
    siteinfo.Gitalk = data.Gitalk
}
getinfo()
export default siteinfo