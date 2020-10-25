import React from 'react';
import 'disqusjs/dist/disqusjs.css'
import DisqusJS from 'disqusjs'
import getinfo from '@/utils/siteinfo'
import './disqus.css'
interface props {
    title: string
    siteinfo: siteinfo
}
export default class Comment extends React.Component<props>{
    constructor(props: props) {
        super(props)
    }

    // componentDidMount() {
    //     getinfo().then((siteinfo) => {
    //         const { Disqus } = siteinfo
    //         DisqusJS({
    //             ...Disqus,
    //             title: this.props.title
    //         });
    //     })

    // }
    componentDidUpdate() {
        if (this.props.siteinfo) {
            console.log(this.props.siteinfo)
            DisqusJS({
                ...this.props.siteinfo.Disqus,
                title: this.props.title
            });
        }
    }
    render() {
        return (<div id="disqus_thread"></div>
        )
    }
}