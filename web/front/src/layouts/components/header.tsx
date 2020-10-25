import request from '@/utils/request';
import React from 'react';
import { NavLink } from 'umi';
import './header.css'

class Header extends React.Component<{ siteinfo: siteinfo }> {
    constructor(props: { siteinfo: siteinfo }) {
        super(props)
    }
    state = {
        Github: '',
        Sitename: '',
        pages: [{
            Comment: true,
            Enable: true,
            Title: "",
            MenuName: "",
            ID: 0
        }]
    }

    componentDidMount() {
        this.getpages()
    }

    componentDidUpdate(prevprops: any) {
        if (!prevprops.siteinfo)
            this.setState({
                Github: this.props.siteinfo.Github
            })
    }

    getpages = async () => {
        const resp = await request(
            "/api/pages/getlist", {
            method: "get"
        }
        )
        this.setState({
            pages: resp.pages
        })
    }

    render() {
        let elements: any = []
        let { pages } = this.state
        pages.forEach((value) => {
            let link = "/pages/" + value.ID
            if (value.Enable)
                elements.push(
                    <li key={value.ID + value.Title} className="nav-list-item"><NavLink key={value.ID} activeClassName="active" isActive={
                        (_, location) => {
                            if (location.pathname.search("^/pages/" + value.ID))
                                return false
                            else return true
                        }
                    } className="nav-list-link" to={link}>{value.MenuName}</NavLink></li>
                )
        })
        if (this.state.Github != "")
            elements.push(<li className="nav-list-item"><a className="nav-list-link" href={this.state.Github} target="_blank">GITHUB</a></li>)

        return (
            <header>
                <a className="logo-link" href="/"><img src="/static/avatar.png"></img></a>
                <ul className="nav nav-list">
                    <li className="nav-list-item"><NavLink activeClassName="active" className="nav-list-link" to="/" isActive={(_, location) => {
                        if (location.pathname != '/' && location.pathname.search("^/page/")) return false
                        return true
                    }} >BLOG</NavLink></li>
                    <li className="nav-list-item"><NavLink activeClassName="active" className="nav-list-link" to="/archives">ARCHIVE</NavLink></li>
                    {elements}
                    <li className="nav-list-item"><a className="nav-list-link" href="/atom.xml" target="_self">RSS</a></li>
                </ul>
            </header>
        )
    }
}
export default Header