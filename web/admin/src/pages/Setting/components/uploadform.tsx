import React from 'react';
import ImgCrop from 'antd-img-crop';
import { Button, Avatar, Card, Upload } from 'antd'
import 'antd/es/modal/style';
import 'antd/es/slider/style';
import style from './uploadform.less'

export default class UploadForm extends React.Component {
    render() {
        return (
            <div className={style.uploadform}>
                <div className={style.uploadcard}>
                    <p>头像</p>
                    <div className={style.imgcenter}>
                            <ImgCrop>
                                <Upload
                                    withCredentials
                                    name="avatar"
                                    action={process.env.url + "/api/settings/avatar"}
                                >
                                     <Avatar className={style.imgicon} size={96} src={process.env.url + "/avatar.png"} />
                                </Upload>
                            </ImgCrop>
                    </div>
                    </div>
                    <div className={style.uploadcard}>
                    <p>网站图标</p>
                    <div className={style.imgcenter}>
                            <ImgCrop>
                                <Upload
                                    withCredentials
                                    name="favicon"
                                    action={process.env.url + "/api/settings/favicon"}
                                >
                                    <Avatar size={96} className={style.imgicon} src={process.env.url + "/favicon.ico"} />
                                </Upload>
                            </ImgCrop>
                    </div>
                </div>

            </div>
        );
    }
}
