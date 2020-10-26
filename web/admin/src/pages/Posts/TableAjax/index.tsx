import React from "react";
import "./index.less";
import { Table, Input, Popover, Divider, Tag } from "antd";
import request from '@/utils/request';
import DeleteModal from "../components/deleteconfirm"
import EditModal from "../components/edit"
import ChangePublish from "../components/changepublish"

const { Search } = Input;

class App extends React.Component {
  state = {
    data: [],
    pagination: {
      current: 1,
      pageSize: 10,
    },
    loading: false,
    keyword: "",
  };

  componentDidMount() {
    this.fetch(this.state.pagination);
  }

  handleTableChange = (pagination: any) => {
    this.fetch(pagination);
  };

  fetch = (pagination?: any) => {
    this.setState({ loading: true });
    request("/api/posts/getlist", {
      method: "GET",
      params: {
        page: pagination.current,
        pageSize: pagination.pageSize,
        word: this.state.keyword
      },
    }).then(data => {
      this.setState({
        loading: false,
        data: data.posts,
        pagination: {
          ...pagination,
          total: data.total,
        }
      });
    });
  };

  searchhandle = (value: string) => {
    this.setState({
      keyword: value,
    })
    setTimeout(() => { //等待state更新完成再执行
      this.componentDidMount()
    }, 0);
  }

  render() {
    const { data, pagination, loading } = this.state;
    return (
      <div>
        <div className="search">
          <Search
            className="searchbar"
            enterButton="Search"
            placeholder="input search text"
            onSearch={this.searchhandle}
          />
        </div>
        <Table
          columns={
            [
              {
                title: "标题",
                dataIndex: "Title",
                width: "20%",
                render:(Title,record:any)=>{
                  return(
                  <a style={{color:"#333"}} href={"/post/"+record.ID}>{Title}</a>
                  )
                }
              },
              {
                title: "创建日期",
                dataIndex: "Update"
              },
              {
                title: "摘要",
                dataIndex: "Description",
                width: "20%",
                render: (Description) =>{
                  const content =(
                    <div style={{width:"500px",maxHeight:"400px"}} dangerouslySetInnerHTML={{__html: Description}} ></div>
                  )
                  return(
                    <Popover content={ content} title="Title">
                    <div className="description" >{ Description.replace(/<(style|script|iframe)[^>]*?>[\s\S]+?<\/\1\s*>/gi,'').replace(/<[^>]+?>/g,'').replace(/\s+/g,' ').replace(/ /g,' ').replace(/>/g,' ') }</div>
                  </Popover>
                  )
                }
              },
              {
                title: "Tags",
                dataIndex: "Tags",
                width: "20%",
                render:tags => (
                  <>
                    {tags.map((tag:any) => {
                      let color = tag.Name.length > 3 ? 'gold' : 'green';
                      color = tag.Name.length > 5 ? 'volcano' : color;
                      color = tag.Name.length > 7 ? 'red' : color;
                      return (
                        <Tag color={color} key={tag}>
                          {tag.Name}
                        </Tag>
                      );
                    })}
                  </>
                ),
              },
              {
                title: "状态",
                dataIndex: 'Publish',
                render: (_, record) => {
                  return (
                    <ChangePublish record={record} />
                  )
                }
              },
              {
                title: "操作",
                dataIndex: "option",
                render: (_: any, record: any) => (
                  <>
                  <div className="optionbtn">
                    <DeleteModal record={record} action={this.componentDidMount.bind(this)} /> {/* 传入刷新数据的方法，绑定上下文 */}
                    <Divider type="vertical" />
                    <EditModal record={record} action={this.componentDidMount.bind(this)} />
                    </div>
                  </>
                ),
              }
            ]}
          className="listtable"
          rowKey={record => record['ID']}
          dataSource={data}
          pagination={pagination}
          loading={loading}
          onChange={this.handleTableChange}
        />
      </div>
    );
  }
}

export default () => (
  <div id="components-table-demo-ajax">
    <App />
  </div>
);
