import { Button,message } from 'antd';
import request from '@/utils/request';
import React from 'react';

class ChangePublish extends React.Component<any> {
    constructor(props: any) {
        super(props);
    }

    state = {
        id: this.props.record.ID,
        status: this.props.record.Publish,
      };
    publish=async () =>{
        let data = await this.change(true)
        if (data['status'] == "ok")
        this.setState({
            status: true,
        })
        else message.error(data['msg'])
    }
    cancel=async () =>{
        let data = await this.change(false)
        if (data['status'] == "ok")
        this.setState({
            status: false,
        })
        else message.error(data['msg'])
    }

    change = (publish:boolean):Promise<any>=>{
        return request("/api/posts/changestatus",{
            method:"PUT",
                data: {
                ID:this.state.id,
                Publish:publish,
            }
        })
    }

    render(){
        if (this.state.status)
        return (
            <Button type="dashed" onClick={this.cancel} size="small" >已发布</Button>
        )
        else return (
            <Button type="dashed" onClick={this.publish} size="small" danger>编辑中</Button>
        )
    }
}

export default ChangePublish