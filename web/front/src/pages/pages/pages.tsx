import React from 'react';
import request from '@/utils/request';
import Comment from '../components/disqus'
import '../post/post.css'

export default class extends React.Component<any> {
  constructor(props: any) {
    super(props)
    this.getdata()
  }
  state = {
    post: {
      ID: "",
      Title: "",
      Update: "",
      HTML: "",
      Comment: true,
    }
  }

  get = (page?: number) => {
    return request("/api/pages/page", {
      method: "get",
      params: {
        page: page || this.props.match.params.pageid,
      },
    })
  }

  getdata = async (page?: number) => {
    let data = await this.get(page || undefined)
    this.setState({
      post: data.page
    })
  }

  componentDidUpdate(prevProps: any) {
    if (this.props.match.params.pageid != prevProps.match.params.pageid)
      this.getdata(this.props.match.params.pageid)
    this.highlightCallBack();
    if (this.props.siteinfo)
      document.title = this.state.post.Title + ' · ' + this.props.siteinfo.SiteName
  }

  highlightCallBack = () => {
    document.querySelectorAll("pre code").forEach(block => {
      try { hljs.highlightBlock(block); }
      catch (e) { console.log(e); }
    });
  };

  render() {
    const elements = []
    if (this.state.post.Comment)
      elements.push(<Comment title={this.state.post.Title} siteinfo={this.props.siteinfo} />)
    return (
      <>
        <section className="container">
          <div className="post">
            <article className="post-block">
              <h1 className="post-title">{this.state.post.Title}</h1>
              <div className="post-info">{this.state.post.Update}</div>
              <div dangerouslySetInnerHTML={{ __html: this.state.post.HTML }} className="post-content">
              </div>
              <div className="post-info">last updated: {this.state.post.Update}</div></article>
          </div>
        </section>
        {elements}
      </>
    );
  }
}