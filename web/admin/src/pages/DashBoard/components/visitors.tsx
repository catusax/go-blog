import React from "react";
import request from "@/utils/request";
import { Area } from "@ant-design/charts";
import { Card } from "antd";
import style from '../Welcome.less'

export default class Visitors extends React.Component {
    constructor(props: any) {
        super(props)
        this.getdata()
    }
    state = {
        data: [{
            Date: '2020',
            Count: 9,
        }]
    }
    getdata = async () => {
        let resp = await request("/api/statistic/recentvisit")
        if (resp = null)
            this.setState({
                data: resp
            })
    }

    render() {
        return (
            <Card title="最近访客" className={style.card}
            >
                <Area style={{ width: "100%" }}
                    height={300}

                    data={this.state.data}
                    xField='Date'
                    yField='Count'
                />
            </Card>
        )
    }
}