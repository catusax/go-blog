import React from 'react';
import './archive.less';
import { Link } from 'umi';
import request from '@/utils/request';
import Pagination from '@/layouts/components/pagination'

class Archive extends React.Component<any> {
  constructor(props: any) {
    super(props)
  }
  state = {
    pagination: {
      current: this.props.match.params.page||1,
      pagesize: 10,
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

  
  get = (page?:number) => {
    return request("/api/public/archives", {
      method: "get",
      params: {
        page: page||this.props.match.params.page,
        pagesize: this.state.pagination.pagesize,
      },
    })
  }

  getdata = async (page?:number) => {
    let data = await this.get(page||undefined)
    this.setState({
      data: data.archives,
      pagination:{
        current: page||this.state.pagination.current,
        pagesize: this.state.pagination.pagesize,
        total:data.total,
      }
    })
  }

  componentDidMount() {
    this.getdata()
  }

  paginationhandle = async(page:number)=>{
    this.getdata(page)
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
        <Pagination current={this.state.pagination.current} pagesize={this.state.pagination.pagesize} total={this.state.pagination.total} onChange={this.paginationhandle.bind(this)} /></>
    );
  }
}

export default Archive