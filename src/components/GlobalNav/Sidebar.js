import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { Link } from 'react-router-dom';
import { Layout, Menu } from 'antd';
import { toggleSider } from '../../actions/settings';

import logo from '../../assets/logo.svg';
import {
  TagOutlined,
  BarChartOutlined,
  LineChartOutlined,
  AppstoreOutlined,
} from '@ant-design/icons';

const { Sider } = Layout;

function Sidebar() {
  const {
    sider: { collapsed },
    navTheme,
  } = useSelector((state) => state.settings);
  const dispatch = useDispatch();

  return (
    <Sider
      breakpoint="lg"
      width="256"
      theme={navTheme}
      collapsible
      collapsed={collapsed}
      trigger={null}
      onBreakpoint={(broken) => {
        dispatch(toggleSider());
      }}
    >
      <div className="menu-header">
        <img alt="logo" className="menu-logo" src={logo} />
        <span hidden={collapsed} className="menu-company">
          BINDU
        </span>
      </div>
      <Menu theme="dark" mode="inline" className="slider-menu">
        <Menu.Item key="1" icon={<BarChartOutlined />}>
          <Link to={'/templates'}>Templates</Link>
        </Menu.Item>
        <Menu.Item key="2" icon={<LineChartOutlined />}>
          <Link to={'/charts/saved'}>Saved Charts</Link>
        </Menu.Item>
        <Menu.Item key="3" icon={<TagOutlined />}>
          <Link to={'/tags'}>Tags</Link>
        </Menu.Item>
        <Menu.Item key="4" icon={<AppstoreOutlined />}>
          <Link to={'/categories'}>Categories</Link>
        </Menu.Item>
      </Menu>
    </Sider>
  );
}

export default Sidebar;
