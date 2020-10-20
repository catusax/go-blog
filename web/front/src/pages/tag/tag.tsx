import React from 'react';
import { Link } from 'umi';
import request from '@/utils/request';
import Pagination from '@/layouts/components/pagination';
declare var title :string

class Archive extends React.Component<any> {
  constructor(props: any) {
    super(props)
  }
  state = {
    tag: this.props.match.params.name,
    pagination: {
      current: this.props.match.params.page||1,
      pagesize: 10,
      total: 0,
    },
    data: [{
        ID: 0,
        Title: "",
        Update: "",
      }]
  }

  
  get = (page?:number) => {
    return request("/api/public/tag", {
      method: "get",
      params: {
        page: page || this.props.match.params.page,
        pagesize: this.state.pagination.pagesize,
        tag: this.state.tag
      },
    })
  }

  getdata = async (page?:number) => {
    let data = await this.get(page||undefined)
    this.setState({
      data: data.posts,
      pagination:{
        current: this.state.pagination.current,
        pagesize: this.state.pagination.pagesize,
        total:data.total,
      }
    })
  }

  componentDidUpdate(){
    document.title = this.state.tag+ ' · '+ title
  }

  componentDidMount() {
    this.getdata()
  }

  paginationhandle = async(page:number)=>{
    this.getdata(page)
  }


  render() {
    const elements: any = []
      this.state.data.forEach((post) => {
        let link = "/post/" + post.ID
        elements.push(
          <div className="post-item">
            <div className="post-info">{post.Update}</div>
            <Link className="post-title-link" to={link} >{post.Title}</Link>
          </div>
        )
      })

    return (
      <><section className="container">
        <div className="archive">
        <h2 className="archive-year">{this.state.tag}</h2>
          {elements}
        </div>
      </section>
        <Pagination current={this.state.pagination.current} pagesize={this.state.pagination.pagesize} total={this.state.pagination.total} onChange={this.paginationhandle.bind(this)} /></>
    );
  }
}

export default Archive