import React from 'react'
import './comment.css'
import Gitalk from 'gitalk'
declare const gitalkconf: any
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
            clientID: gitalkconf.clientID,
            clientSecret: gitalkconf.clientSecret,
            repo: gitalkconf.repo,
            owner: gitalkconf.owner,
            admin: [gitalkconf.owner],
            id: this.props.title,      // Ensure uniqueness and length less than 50
            distractionFreeMode: false  // Facebook-like distraction free mode
          })
          
          gitalk.render('gitalk-container')
    }
    render() {
        return (<div id="gitalk-container"></div>)
    }
}