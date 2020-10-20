import React from "react";
import { Button, Checkbox, Input, message, Space } from 'antd';
import request from '@/utils/request'
import Editor from "react-markdown-editor-lite";
import "react-markdown-editor-lite/lib/index.css";
import ReactMarkdown from "react-markdown";
import './form.less'
interface props {
    name: string
    id: number
}

export default class PageForm extends React.Component<props>{
    constructor(props: props) {
        super(props)
    }
    state = {
        ID: this.props.id,
        Title: "",
        MenuName: this.props.name,
        Content: "",
        Comment: true,
        Enable: true
    }

    componentDidMount(){
        this.getpage()
    }
    getpage = async()=>{
        let data = await request(
            "/api/pages/page",{
                method: "get",
                params:{
                    page: this.state.ID
                }
            }
        )
        if (data.page)
        this.setState(
            {
                ID: data.page.ID,
                Title: data.page.Title,
                MenuName: data.page.MenuName,
                Content: data.page.Content,
                Comment: data.page.Comment,
                Enable: data.page.Enable
            }
        )
    }

    submit = async () => {
        let resp = await request(
            "/api/pages/new", {
            method: "post",
            data: {
                ID: this.state.ID,
                Title: this.state.Title,
                MenuName: this.state.MenuName,
                Content: this.state.Content,
                Comment: this.state.Comment,
                Enable: this.state.Enable
            }
        }
        )
        if (resp.status == "ok") message.success("提交成功")
        else message.error(resp.msg)
    }

    render() {
        return (
            <>
                <div className="formtitle">
                    <Input addonBefore="标题" value={this.state.Title} onChange={(e) => { this.setState({ Title: e.target.value }) }} style={{ width: "70%" }} />
                    <Space className="formtitleitem" >
                        <Checkbox checked={this.state.Comment} onChange={(e) => { this.setState({ Comment: e.target.checked }) }}>开启评论</Checkbox>
                        <Checkbox checked={this.state.Enable} onChange={(e) => { this.setState({ Enable: e.target.checked }) }}>启用此页面</Checkbox>
                        <Button onClick={this.submit} type="primary">保存</Button></Space>
                </div>
                <Editor
                    value={this.state.Content}
                    onChange={(data) => { this.setState({ Content: data.text }) }}
                    style={{
                        height: "500px"
                    }}
                    renderHTML={text => <ReactMarkdown source={text} />}
                />
            </>
        )
    }
}