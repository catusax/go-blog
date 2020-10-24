import { PageContainer } from '@ant-design/pro-layout';
import React from 'react';
import styles from './index.less';
import PageForm from './components/form'

export default () => {

  return (
    <PageContainer content="在此处编辑友链页面" className={styles.main}>
      <PageForm name="LINKS" id={1} />
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
