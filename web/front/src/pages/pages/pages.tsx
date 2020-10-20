import React from 'react';
import request from '@/utils/request';
import { Link } from 'umi';
declare var hljs: { highlightBlock: (arg0: Element) => void; }
declare var title:string

class Archive extends React.Component<any> {
  constructor(props: any) {
    super(props)
  }
  state = {
    post: {
      ID: "",
      Title: "",
      Update: "",
      HTML: "",
    }
  }

  get = () => {
    return request("/api/pages/page", {
      method: "get",
      params: {
        page: this.props.match.params.pageid,
      },
    })
  }

  getdata = async () => {
    let data = await this.get()
    this.setState({
      post: data.page
    })
  }

  componentDidUpdate(prevProps: any) {
    if (this.props.match.params.pageid !== prevProps.match.params.pageid){
    this.getdata()
    this.highlightCallBack();
    document.title = this.state.post.Title+ ' · '+ title}
  }

  highlightCallBack = () => {
    document.querySelectorAll("pre code").forEach(block => {
      try { hljs.highlightBlock(block); }
      catch (e) { console.log(e); }
    });
  };

  componentDidMount() {
    this.getdata()
    this.highlightCallBack();
  }


  render() {
    return (
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
    );
  }
}

export default Archive