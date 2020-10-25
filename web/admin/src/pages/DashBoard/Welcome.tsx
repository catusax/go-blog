import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import Visitors from './components/visitors'
import Mostread from './components/mostread'
import Recentupdated from './components/recentupdated'
import Browsers from './components/browsers'
import Os from './components/os'
import Title from './components/title'
import style from './Welcome.less'
import { Col, Row } from 'antd';

export default () => {
  return (
    <PageContainer className={style.main}>
      <Row gutter={1}>
        <Col flex={1}>
          <Title />
        </Col>
      </Row>
      <Row gutter={12}>
        <Col flex={3} style={{ maxWidth: "60%" }}>
          <Visitors />
        </Col>
        <Col flex={2}>
          <Mostread />
        </Col>
      </Row>

      <Row gutter={12}>
        <Col flex={2}>
          <Recentupdated />
        </Col>
        <Col flex={2} style={{ maxWidth: "30%" }} >
          <Browsers />
        </Col>
        <Col flex={2} style={{ maxWidth: "30%" }} >
          <Os />
        </Col>
      </Row>
      <div
        style={{
          paddingTop: 100,
          textAlign: 'center',
        }}
      >
      </div>
    </PageContainer>

  )
}
