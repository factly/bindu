import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { Layout } from 'antd';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons';
import { toggleSider } from '../../actions/settings';
import AccountMenu from './AccountMenu';
import OrganisationSelector from '../OrganisationSelector';

function Header() {
  const collapsed = useSelector((state) => state.settings.sider.collapsed);
  const dispatch = useDispatch();
  const MenuFoldComponent = collapsed ? MenuUnfoldOutlined : MenuFoldOutlined;

  return (
    <Layout.Header className="layout-header">
      <div style={{ display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
        <MenuFoldComponent
          style={{ fontSize: '16px', marginRight: 'auto' }}
          onClick={() => dispatch(toggleSider())}
        />

        <OrganisationSelector />

        <AccountMenu />
      </div>
    </Layout.Header>
  );
}

export default Header;
