import request from './request'
export default async function getinfo() {
    var siteinfo = {
        SiteName: "BLOG",
        Gitalk: {
            Owner: 'qqwwee',
            Repo: '11223',
            ClientID: 'fghfhf3660',
            ClientSecret: 'asdfsdf',
        },
        Disqus: {
            shortname: "name",
            apikey: "apikey",
            siteName: '',
            api: "https://disqus.skk.moe/disqus/",
            admin: '',
            adminLabel: '',
        }
    }
    let data = await request("/api/public/info")
    siteinfo.SiteName = data.SiteName
    siteinfo.Gitalk = data.Gitalk
    siteinfo.Disqus = data.Disqus
    return siteinfo
}