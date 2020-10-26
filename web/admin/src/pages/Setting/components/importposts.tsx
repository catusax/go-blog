import { UploadOutlined } from '@ant-design/icons';
import { Button, message, Upload } from 'antd';
import React from 'react'

export default class extends React.Component {

    onChange(info: any) {
        if (info.file.status !== 'uploading') {
            console.log(info.file, info.fileList);
        }
        if (info.file.status === 'done') {
            message.success(`${info.file.name} file uploaded successfully`);
            message.success(info.file.status.response)
        } else if (info.file.status === 'error') {
            message.error(`${info.file.name} file upload failed.`);
        }
    }

    render() {
        return (
            <><div style={{padding:"30px"}}>
            <p>
                导入Hexo的MD文件
            </p>
                <Upload
                    multiple={true}
                    name='files'
                    action={process.env.url + '/api/posts/import'}
                    withCredentials={true}
                    onChange={this.onChange}
                >
                    <Button icon={<UploadOutlined />}>点击上传文件</Button>
                </Upload>
                </div>
            </>
        )
    }
}