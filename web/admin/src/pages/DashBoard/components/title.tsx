import React from "react";
import { Card } from 'antd'
import request from '@/utils/request'

export default class Title extends React.Component {
    state = {
        Year: new Date().getFullYear(),
        Mon: new Date().getMonth(),
        Day: new Date().getDate(),
        Hour: new Date().getHours(),
        min: new Date().getMinutes(),
        sec: new Date().getSeconds(),
        Visitors: 1,
        Posts: 1
    }
    componentDidMount() {
        setInterval(this.tick, 1000);
        this.getdata()
    }
    getdata = async () => {
        let resp = await request("/api/statistic/total")
        this.setState({
            Visitors: resp.Visitors,
            Posts: resp.Posts
        })
    }

    tick = () => {
        let date = new Date()
        this.setState({
            Year: date.getFullYear(),
            Mon: date.getMonth(),
            Day: date.getDate(),
            Hour: date.getHours(),
            min: date.getMinutes(),
            sec: date.getSeconds()
        })
    }

    render() {
        return (
            <Card style={{ margin: "10px" }} >
                <Card.Meta
                    title={"欢迎来到Blog后台，现在是" + this.state.Year + '年' + this.state.Mon + '月' + this.state.Day + '日   ' + this.state.Hour + ':' + this.state.min + ':' + this.state.sec}
                    description={<p
                        style={{ float: "right", paddingRight: "30px" }}
                    >总文章数：{this.state.Posts}<span style={{ width: "20px", display: 'inline-block' }}> </span>总访客数：{this.state.Visitors}</p>}
                />
            </Card>
        )
    }
}