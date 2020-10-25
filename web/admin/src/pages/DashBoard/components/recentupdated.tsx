import request from "@/utils/request";
import { Card, List } from "antd";
import React from "react";
import style from '../Welcome.less'

export default class Recentupdated extends React.Component {
    constructor(props: any) {
        super(props)
        this.getdata()
    }
    state = {
        data: [{
            Title:'',
            Update:'',
            PostID:0,
        }]
    }
    getdata = async () => {
        let resp = await request("/api/statistic/recentpost")
        this.setState({
            data: resp
        })
    }

    render() {
        return (
            <Card className={style.card} style={{overflowY:"scroll"}}
            title="最近更新"
            >
            <List
                itemLayout="horizontal"
                dataSource={this.state.data}
                renderItem={item => (
                    <List.Item>
                        <List.Item.Meta
                        title={<a href={"/posts/"+item.PostID}>{item.Title}</a>}
                        description={"更新日期："+item.Update}
                        />
                        
                    </List.Item>
                )}
            />
            </Card>
        )
    }
}