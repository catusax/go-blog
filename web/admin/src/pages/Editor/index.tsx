import { PageContainer } from '@ant-design/pro-layout';
import React, { useState, useEffect } from 'react';
import { Spin } from 'antd';
import styles from './index.less';
import MDeditor from './components/mdeditor'

export default () => {
  const [loading, setLoading] = useState<boolean>(true);
  useEffect(() => {
    setTimeout(() => {
      setLoading(false);
    }, 0);
  }, []);
  return (
    <PageContainer content="新建一篇文章" className={styles.main}>
      <MDeditor />
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
