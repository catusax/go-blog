import { PageContainer } from '@ant-design/pro-layout';
import React, { useState, useEffect } from 'react';
import { Spin } from 'antd';
import styles from './index.less';
import TableAjax from './TableAjax';
export default () => {
  const [loading, setLoading] = useState<boolean>(true);
  useEffect(() => {
    setTimeout(() => {
      setLoading(false);
    }, 0); //TODO 在加载列表后修改spin状态
  }, []);
  return (
    <PageContainer content="这是一个新页面，从这里进行开发！" className={styles.main}>
      <TableAjax />
      <div
        style={{
          paddingTop: 100,
          textAlign: 'center',
        }}
      >
        <Spin spinning={loading} size="large" />
      </div>
    </PageContainer>
  );
};
