import { PageContainer } from '@ant-design/pro-layout';
import React from 'react';
import styles from './index.less';
import TableAjax from './TableAjax';

export default () => {
  return (
    <PageContainer className={styles.main}>
      <TableAjax />
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
