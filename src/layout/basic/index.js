import React from 'react';
import { Layout, Card, Menu, Space } from 'antd';

import { Link } from 'react-router-dom';
import OrganisationSelector from '../../components/OrganisationSelector';

const { Header, Content } = Layout;

function BasicLayout(props) {
  const { children } = props;
  return (
    <Layout>
      <Header className="header">
        <Space direction="horizontal" size={48}>
          <OrganisationSelector />
          <Menu theme="light" mode="horizontal">
            <Menu.Item key="1">
              <Link to={'/templates'}>Templates</Link>
            </Menu.Item>
            <Menu.Item key="2">Saved Charts</Menu.Item>
          </Menu>
        </Space>
      </Header>
      <Content className="layout-content">
        <Card bordered={false} className="wrap-children-content">
          {children}
        </Card>
      </Content>
    </Layout>
  );
}

export default BasicLayout;
