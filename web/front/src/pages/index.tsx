import React from 'react';
import { Link, history } from 'umi';
import request from '@/utils/request';
import Pagination from '@/layouts/components/pagination'
import './index.css'

class Archive extends React.Component<any> {
  constructor(props: any) {
    super(props)
    this.getdata()
  }
  state = {
    pagination: {
      current: parseInt(this.props.match.params.page) || 1,
      pageSize: 10,
      total: 0,
    },
    data: [{
      ID: 0,
      Title: "",
      Update: "",
      Description: "",
    }]
  }


  get = (page?: number) => {
    return request("/api/public/index", {
      method: "get",
      params: {
        page: page || this.props.match.params.page,
        pageSize: this.state.pagination.pageSize,
      },
    })
  }

  getdata = async (page?: number) => {
    let data = await this.get(page || undefined)
    this.setState({
      data: data.posts,
      pagination: {
        current: page || this.state.pagination.current,
        pageSize: this.state.pagination.pageSize,
        total: data.total,
      }
    })
  }

  componentDidUpdate(prevprops: any) {
    if (this.props.match.params.page != prevprops.match.params.page)
      this.getdata(parseInt(this.props.match.params.page))
    if (this.props.siteinfo)
      document.title = this.props.siteinfo.SiteName
      this.highlightCallBack()
  }

  highlightCallBack = () => {
    document.querySelectorAll("pre code").forEach(block => {
      try { hljs.highlightBlock(block); }
      catch (e) { console.log(e); }
    });
  };

  paginationhandle = (page: number) => {
    history.push("/page/" + page)
  }


  render() {
    const elements: any = []
    this.state.data.forEach((post) => {
      let link = "/post/" + post.ID
      elements.push(

        <li className="post-list-item">
          <article className="post-block">
            <h2 className="post-title">
              <Link className="post-title-link" to={link} >{post.Title}</Link>
            </h2>
            <div className="post-info">{post.Update}</div>
            <div className="post-content" dangerouslySetInnerHTML={{ __html: post.Description }}></div>
            <Link className="read-more" to={link} >...more</Link>
          </article>
        </li>
      )
    }

    )
    return (
      <section className="container">
        <ul className="home post-list">
          {elements}
        </ul>
        <Pagination current={this.state.pagination.current} pageSize={this.state.pagination.pageSize} total={this.state.pagination.total} onChange={this.paginationhandle.bind(this)} />
      </section>
    );
  }
}

export default Archive
