import React from "react";
import request from "@/utils/request";
import { Pie } from "@ant-design/charts";
import { Card } from "antd";
import style from '../Welcome.less'

export default class Os extends React.Component {
    constructor(props: any) {
        super(props)
        this.getdata()
    }
    state = {
        data: [{
            Os: '',
            Count: 0
        }]
    }
    getdata = async () => {
        let resp = await request("/api/statistic/os")
        if (resp != null)
        this.setState({
            data: resp
        })
    }

    render() {
        return (
            <Card title="操作系统" className={style.card}
            >
                <Pie
                    data={this.state.data}
                    color= {['#feb64d','#ff7c7c','#9287e7','#60acfc' ,'#32d3eb','#5bc49f']}
                    angleField='Count'
                    colorField='Os'
                    radius={0.8}
                    height={300}
                    width={300}
                />
            </Card>
        )
    }
}