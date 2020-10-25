import request from "@/utils/request";
import { Card, List } from "antd";
import React from "react";
import style from '../Welcome.less'

export default class Mostread extends React.Component {
    constructor(props: any) {
        super(props)
        this.getdata()
    }
    state = {
        data: [{
            Title:'',
            Update:'',
            Count:0,
            PostID:0,
        }]
    }
    getdata = async () => {
        let resp = await request("/api/statistic/mostread")
        this.setState({
            data: resp
        })
    }

    render() {
        return (
            <Card className={style.card}
            title="最多阅读"
            >
            <List
                itemLayout="horizontal"
                dataSource={this.state.data}
                renderItem={item => (
                    <List.Item>
                        <List.Item.Meta
                        title={<a href={"/posts/"+item.PostID}>{item.Title}</a>}
                        description={"更新日期："+item.Update+" 阅读量："+item.Count}
                        />
                        
                    </List.Item>
                )}
            />
            </Card>
        )
    }
}