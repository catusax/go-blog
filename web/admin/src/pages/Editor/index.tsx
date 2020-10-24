import { PageContainer } from '@ant-design/pro-layout';
import React from 'react';
import styles from './index.less';
import MDeditor from './components/mdeditor'

export default () => {
  return (
    <PageContainer content="新建一篇文章" className={styles.main}>
      <MDeditor />
      <div
        style={{
          paddingTop: 50,
          textAlign: 'center',
        }}
      >
      </div>
    </PageContainer>
  );
};
