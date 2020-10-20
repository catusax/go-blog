import React from 'react';
import { Link } from 'umi';
import request from '@/utils/request';
import Pagination from '@/layouts/components/pagination'

class Archive extends React.Component<any> {
  constructor(props: any) {
    super(props)
  }
  state = {
    pagination: {
      current: 1,
      pagesize: 10,
      total: 0,
    },
    data: [{
      ID: 0,
      Title: "",
      Update: "",
      Description: "",
    }]
  }


  get = (page?:number) => {
    return request("/api/public/index", {
      method: "get",
      params: {
        page: page || this.props.match.params.page,
        pagesize: this.state.pagination.pagesize,
      },
    })
  }

  getdata = async (page?:number) => {
    let data = await this.get(page||undefined)
    this.setState({
      data: data.posts,
      pagination: {
        current: page||this.state.pagination.current,
        pagesize: this.state.pagination.pagesize,
        total: data.total,
      }
    })
  }

  componentDidMount() {
    this.getdata()
  }

  paginationhandle = async (page: number) => {
    this.getdata(page)
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
        <Pagination current={this.state.pagination.current} pagesize={this.state.pagination.pagesize} total={this.state.pagination.total} onChange={this.paginationhandle.bind(this)} />
      </section>
    );
  }
}

export default Archive
