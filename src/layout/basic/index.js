import React from 'react';
import { Layout, Card } from 'antd';

import './basic.css';
import Sidebar from '../../components/GlobalNav/Sidebar';
import Header from '../../components/GlobalNav/Header';

const { Content } = Layout;

function BasicLayout(props) {
  const { children } = props;
  return (
    <Layout hasSider={true}>
      <Sidebar />
      <Layout>
        <Header />
        <Content className="layout-content">
          <Card bordered={false} className="wrap-children-content">
            {children}
          </Card>
        </Content>
      </Layout>
    </Layout>
  );
}

export default BasicLayout;
