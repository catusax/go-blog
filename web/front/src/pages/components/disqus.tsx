import React from 'react';
import 'disqusjs/dist/disqusjs.css'
import DisqusJS from 'disqusjs'
import getinfo from '@/utils/siteinfo'
import './disqus.css'
interface props {
    title: string
}
export default class Comment extends React.Component<props>{
    constructor(props: props) {
        super(props)
    }

    componentDidMount() {
        getinfo().then((siteinfo)=>{
            const {Disqus} = siteinfo
            DisqusJS({
                ...Disqus,
                title: this.props.title
            });
        })

    }
    render() {
        return (<div id="disqus_thread"></div>
        )
    }
}