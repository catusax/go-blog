import React from "react";
import { Button, message } from 'antd'
import { history } from 'umi';
import Editor from "react-markdown-editor-lite";
import request from '@/utils/request';
import "react-markdown-editor-lite/lib/index.css";
import "./mdeditor.less"
import ReactMarkdown from "react-markdown";
import description from "./mdeditorplugin"

export default () => {
    Editor.use(description);
    const mdEditor = React.useRef<Editor | null>(null);
    const [value, setValue] = React.useState("---\n\
title: \n\
tags: \n\n\
---\n");

    const handleEditorChange = ({ text }: any) => {
        setValue(text);
    };

    const save = (): any => {
        if (mdEditor.current) {
            return request('/api/posts/new', {
                method: 'POST',
                data: {
                    "Content": mdEditor.current.getMdValue(),
                    "Publish": false,
                },
            });
        }
    }
    const publish = (): any => {
        if (mdEditor.current) {
            return request('/api/posts/new', {
                method: 'POST',
                data: {
                    "Content": mdEditor.current.getMdValue(),
                    "Publish": true,
                },
            });
        }
    }



    const savepost = async () => {
        const resp = await save()
        if (resp['status'] == 'ok') {
            message.success('保存成功')
            history.push('/posts')
        }
        else
            message.error(resp['msg'])
    }

    const publishpost = async () => {
        const resp = await publish()
        if (resp['status'] == 'ok') {
            message.success('发布成功')
            history.push('/posts')
        }
        else
            message.error(resp['msg'])
    }

    return (
        <div className="MDeditor">
            <Editor
                ref={mdEditor}
                value={value}
                style={{
                    height: "500px"
                }}
                onChange={handleEditorChange}
                renderHTML={text => <ReactMarkdown source={text} />}
            />
            <div className="submitbtn">
                <Button onClick={savepost} type="dashed">保存</Button>
                <Button onClick={publishpost} type="primary">发布</Button>
            </div>
        </div>
    );
}
