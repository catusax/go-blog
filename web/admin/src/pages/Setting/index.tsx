import { PageContainer } from '@ant-design/pro-layout';
import React from 'react';
import style from './index.less';
import UploadForm from './components/uploadform'
import SettingForm from './components/settingform'
import { Col, Card } from 'antd';
import { FormattedDisplayName } from 'umi';

export default () => {
  return (
    <PageContainer className={style.main}>
      <Card>
        <h3>修改设置</h3>
        <div className={style.settingform}>
          <Col flex="50%">
            <SettingForm />
          </Col>
          <Col flex="10%">
            <UploadForm />
          </Col>
        </div>
        <div style={{ paddingTop: 100, textAlign: 'center' }}>
        </div>
      </Card>
    </PageContainer>
  );
};