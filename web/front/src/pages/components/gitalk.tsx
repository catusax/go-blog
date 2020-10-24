import React from 'react'
import './gitalk.css'
import Gitalk from 'gitalk'
import siteinfo from '@/utils/siteinfo'
interface props {
    title: string
}
export default class Comment extends React.Component<props>{
    constructor(props: props) {
        super(props)
    }
    componentDidUpdate(){
        console.log(this.props.title)
        var gitalk = new Gitalk({
            clientID: siteinfo.Gitalk.ClientID,
            clientSecret: siteinfo.Gitalk.ClientSecret,
            repo: siteinfo.Gitalk.Repo,
            owner: siteinfo.Gitalk.Owner,
            admin: [siteinfo.Gitalk.Owner],
            id: this.props.title,      // Ensure uniqueness and length less than 50
            distractionFreeMode: false  // Facebook-like distraction free mode
          })
          
          gitalk.render('gitalk-container')
    }
    render() {
        return (<div id="gitalk-container"></div>)
    }
}