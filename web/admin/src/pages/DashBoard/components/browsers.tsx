import React from "react";
import request from "@/utils/request";
import { Pie } from "@ant-design/charts";
import { Card } from "antd";
import style from '../Welcome.less'

export default class Browsers extends React.Component {
    constructor(props: any) {
        super(props)
        this.getdata()
    }
    state = {
        data: [{
            Browser: '',
            Count: 0
        }]
    }
    getdata = async () => {
        let resp = await request("/api/statistic/browser")
        if (resp != null)
        this.setState({
            data: resp
        })
    }

    render() {
        return (
            <Card title="浏览器" className={style.card}
            >
                <Pie
                    data={this.state.data}
                    angleField='Count'
                    colorField='Browser'
                    radius={0.8}
                    height={300}
                    width={300}
                />
            </Card>
        )
    }
}