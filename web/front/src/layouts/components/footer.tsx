import React from 'react'
import './footer.css'

class Footer extends React.Component<{ siteinfo: siteinfo }> {
    constructor(props: { siteinfo: siteinfo }) {
        super(props)
    }
    state = {
        SiteName: ''
    }
    componentDidUpdate(prevprops: any) {
        if (!prevprops.siteinfo)
            this.setState(this.props.siteinfo)
    }
    render() {
        return (
            <footer>
                <div className="copyright">
                    <p>© 2015 - 2020 <a href="/">{this.state.SiteName}</a>, powered by <a href="https://github.com/coolrc136/go-blog" target="_blank">Go-blog</a>.
                    </p>
                </div>
            </footer>
        )
    }
}

export default Footer