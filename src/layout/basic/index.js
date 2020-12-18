import React from 'react';
import { Layout, Card, notification } from 'antd';

import './basic.css';
import Sidebar from '../../components/GlobalNav/Sidebar';
import Header from '../../components/GlobalNav/Header';
import { useDispatch, useSelector } from 'react-redux';

import { getSpaces } from '../../actions/spaces';

const { Content } = Layout;

function BasicLayout(props) {
  const { children } = props;

  const dispatch = useDispatch();

  const { permission, orgs, loading, selected } = useSelector((state) => {
    const { selected, orgs, loading } = state.spaces;

    if (selected > 0) {
      return {
        permission: state.spaces.details[selected].permissions || [],
        orgs: orgs,
        loading: loading,
        selected: selected,
      };
    }
    return { orgs: orgs, loading: loading, permission: [], selected: selected };
  });

  const { type, message, description } = useSelector((state) => state.notifications);

  React.useEffect(() => {
    dispatch(getSpaces());
  }, [dispatch, selected]);

  React.useEffect(() => {
    if (type && message && description && selected !== 0) {
      notification[type]({
        message: message,
        description: description,
      });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [description]);

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
