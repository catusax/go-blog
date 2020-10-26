import React from 'react';
import './archive.css';
import { Link, history } from 'umi';
import request from '@/utils/request';
import Pagination from '@/layouts/components/pagination'

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
      Year: 0,
      Posts: [{
        ID: 0,
        Title: "",
        Update: "",
      }]
    }],
  }


  get = (page?: number) => {
    return request("/api/public/archives", {
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
      data: data.archives,
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
      document.title = 'Archive · ' + this.props.siteinfo.SiteName
  }

  paginationhandle = (page: number) => {
    history.push('/archives/page/' + page)
  }

  render() {
    const elements: any = []
    this.state.data.forEach((archive) => { //最外层，增加年份
      elements.push(
        <h2 key={archive.Year} className="archive-year">{archive.Year}</h2>
      )
      archive.Posts.forEach((post) => {
        let link = "/post/" + post.ID
        elements.push(
          <div className="post-item">
            <div className="post-info">{post.Update}</div>
            <Link className="post-title-link" to={link} >{post.Title}</Link>
          </div>
        )
      })
    }

    )
    return (
      <><section className="container">
        <div className="archive">
          {elements}
        </div>
      </section>
        <Pagination current={this.state.pagination.current} pageSize={this.state.pagination.pageSize} total={this.state.pagination.total} onChange={this.paginationhandle.bind(this)} /></>
    );
  }
}

export default Archive