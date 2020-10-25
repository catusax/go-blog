import { Button, Form, Input, InputNumber, message } from "antd";
import React from "react";
import style from './settingform.less'
import request from '@/utils/request'

export default class SettingForm extends React.Component {
    state = {
        port: 8080,
        sitename: "wersefsfe",
        db: {
            host: "",
            port: 5432,
            user: "",
            password: "",
            name: "",
        },
        disqus: {
            shortname: "",
            apikey: "",
            siteName: "",
            api: "https://disqus.skk.moe/disqus/",
            admin: "",
            adminLabel: "",
        }
    }
    async getdata() {
        let data = await request("/api/settings/getconfig")
        this.setState(data)
    }

    componentDidMount() {
        this.getdata()
    }

    layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };
    validateMessages = {
        required: '${label} is required!',
        types: {
            number: '${label} is not a validate number!',
        },
        number: {
            range: '${label} must be between ${min} and ${max}',
        },
    };

    onFinish = (values: any) => {
        request("/api/settings/changeconfig", {
            method: "post",
            data: values
        }).then(
            message.success("修改成功")
        )
    };

    render() {
        return (
            <div className={style.container}>
                <p>基本设置</p>
                <Form key={this.state.sitename} {...this.layout} name="nest-messages" onFinish={this.onFinish} validateMessages={this.validateMessages} initialValues={this.state}>
                    <Form.Item
                        name="sitename"
                        label="博客名称"
                        rules={[
                            {
                                required: true,
                            },
                        ]}
                    >
                        <Input/>
                    </Form.Item>
                    <Form.Item
                        name='github'
                        label="Github地址"
                        extra='选填，菜单栏的Github链接'
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name="port"
                        label="http端口"
                        rules={[
                            {
                                required: true,
                                type: 'number',
                                min: 1,
                                max: 65535,
                            },
                        ]}
                    >
                        <InputNumber />
                    </Form.Item>

                    <Form.Item
                        name={['db', 'host']}
                        label="数据库地址"
                        rules={[{ required: true, }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        name={['db', 'port']}
                        label="数据库端口"
                        rules={[{ required: true, type: 'number' }]}
                    >
                        <InputNumber />
                    </Form.Item>

                    <Form.Item
                        name={['db', 'user']}
                        label="数据库用户名"
                        rules={[{ required: true, }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        name={['db', 'password']}
                        label="数据库密码"
                        rules={[{ required: true, }]}
                    >
                        <Input.Password />
                    </Form.Item>

                    <Form.Item
                        name={['disqus', 'shortname']}
                        label="Disqus shortname"
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        name={['disqus', 'apikey']}
                        label="Disqus apikey"
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name={['disqus', 'sitename']}
                        label="Disqus siteName"
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name={['disqus', 'api']}
                        label="Disqus api"
                        extra="选填，默认为：https://disqus.skk.moe/disqus/"
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name={['disqus', 'admin']}
                        label="Disqus admin"
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        name={['disqus', 'adminlabel']}
                        label="Disqus adminLabel"
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item wrapperCol={{ ...this.layout.wrapperCol, offset: 8 }}>
                        <Button type="primary" htmlType="submit">
                            Submit
              </Button>
                    </Form.Item>
                </Form>


            </div>
        )
    }
}