import React from "react";
import { Layout, Card, Menu } from 'antd';

import {
  Link
} from "react-router-dom";

const { Header, Content } = Layout;

function BasicLayout(props){
	const { children } = props;
	return (
		<Layout>
		    <Header className="header">
		      <Menu theme="light" mode="horizontal">
		        <Menu.Item key="1"><Link to={'/templates'}>Templates</Link></Menu.Item>
		        <Menu.Item key="2">Saved Charts</Menu.Item>
		      </Menu>
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