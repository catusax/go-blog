import React from 'react';
import {NavLink} from 'umi'

class Header extends React.Component<any> {
    constructor(props:any){
        super(props)
    }

    render(){
        return (
            <header>
                <a className="logo-link" href="/"><img src="/favicon.png"></img></a>
                <ul className="nav nav-list">
                    <li className="nav-list-item"><NavLink activeClassName="active" className="nav-list-link" to="/" isActive={(_, location) => {
          if (location.pathname != '/' && location.pathname.search("^/page")) return false
          return true
        }} >BLOG</NavLink></li>
                    <li className="nav-list-item"><NavLink activeClassName="active" className="nav-list-link" to="/archives">ARCHIVE</NavLink></li>
                    <li className="nav-list-item"><NavLink activeClassName="active" className="nav-list-link" to="/links">LINKS</NavLink></li>
                    <li className="nav-list-item"><a className="nav-list-link" href="https://github.com/coolrc136" target="_blank">GITHUB</a></li>
                    <li className="nav-list-item"><a className="nav-list-link" href="/atom.xml" target="_self">RSS</a></li>
                </ul>
            </header>
        )
        }
}
export default Header