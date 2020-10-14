import { Modal, Button , message} from 'antd';
import React from 'react';
import request from '@/utils/request';

class DeleteModal extends React.Component<any> {

    constructor(props: any) {
        super(props);
    }

    state = {
        ModalText: '确认删除' + this.props.record.Title + '?',
        visible: false,
        confirmLoading: false,
    };

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
        const result = await this.delete()
        if (result['status'] == 'ok')
            this.setState({
                visible: false,
                confirmLoading: false,
            });
        else message.error(result['msg']);
        this.props['action']() //刷新数据
    };

    delete = () => {
        return request('/api/posts/delete', {
            method: 'DELETE',
            data: {
                "ID": this.props['record'].ID,
            },
        });
    }


    handleCancel = () => {
        console.log('Clicked cancel button');
        this.setState({
            visible: false,
        });
    };

    render() {
        const { visible, confirmLoading, ModalText } = this.state;
        return (
            <>
                <Button type="primary" danger onClick={this.showModal}>
                    删除
        </Button>
                <Modal
                    title="Title"
                    visible={visible}
                    onOk={this.handleOk}
                    confirmLoading={confirmLoading}
                    onCancel={this.handleCancel}
                >
                    <p>{ModalText}</p>
                </Modal>
            </>
        );
    }
}
export default DeleteModal