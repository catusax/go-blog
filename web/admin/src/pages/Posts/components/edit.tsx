import { Modal, Button, message } from 'antd';
import React from 'react';
import request from '@/utils/request';
import Editor from "react-markdown-editor-lite";
import 'react-markdown-editor-lite/lib/index.css';
import ReactMarkdown from 'react-markdown';

class EditModal extends React.Component<any> {

    constructor(props: any) {
        super(props);
    }

    state = {
        contentmd: "---\n"+this.props.record.Yaml+"\n---\n"+this.props.record.Content,
        visible: false,
        confirmLoading: false,
    };

    handleEditorChange = ({ text }: any) => {
        this.setState({
            contentmd: text,
        });
    };

    save = (): any => {
        return request('/api/posts/new', {
            method: 'POST',
            data: {
                "Content": this.state.contentmd,
                "ID": this.props.record.ID
            },
        });
    }

    showModal = () => {
        this.setState({
            visible: true,
        });
    };

    handleOk = async () => {
        this.setState({
            ModalText: '正在删除',
            confirmLoading: true,
        });
        const result = await this.save()
        if (result['status'] == 'ok')
            this.setState({
                visible: false,
                confirmLoading: false,
            });
        else message.error(result['msg']);
        this.props['action']() //刷新数据
    };

    handleCancel = () => {
        this.setState({
            visible: false,
        });
    };

    render() {
        const { visible, confirmLoading, contentmd } = this.state;
        return (
            <>
                <Button type="default" danger onClick={this.showModal}>
                    编辑
        </Button>
                <Modal
                    title="Title"
                    visible={visible}
                    onOk={this.handleOk}
                    confirmLoading={confirmLoading}
                    onCancel={this.handleCancel}
                >
                    <Editor
                    style={{ height: "500px" }}
                    value={contentmd}
                    onChange={this.handleEditorChange}
                    renderHTML={(text: string | undefined) => <ReactMarkdown source={text} />}
                    />
                </Modal>
            </>
        );
    }
}
export default EditModal