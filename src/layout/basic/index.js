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

  const selected = useSelector(({ spaces }) => spaces.selected);

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
          <Card key={selected.toString()} bordered={false} className="wrap-children-content">
            {children}
          </Card>
        </Content>
      </Layout>
    </Layout>
  );
}

export default BasicLayout;
