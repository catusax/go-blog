//siteinfo是全局配置信息，通过路由传到每个路由组件
interface siteinfo {
    SiteName: string,
    Github: string,
    Disqus: {
        shortname: string,
        apikey: string,
        siteName: string,
        api: string,
        admin: string,
        adminLabel: string,
    }
}

declare var hljs: { highlightBlock: (arg0: Element) => void; }
