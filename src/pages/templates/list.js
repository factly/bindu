import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { List, Card, Popconfirm, Typography } from 'antd';
import { Link, useLocation } from 'react-router-dom';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';

import { deleteTemplate, getTemplates } from '../../actions/templates';
import routes from '../../config/routesConfig';

const { Meta } = Card;
const { Title } = Typography;

function TemplatesGroup({ ids = [], category = {} }) {
  const dispatch = useDispatch();
  const location = useLocation();
  const templates = useSelector(({ templates }) => ids.map((id) => templates.details[id]));

  const handleDelete = (id) => {
    dispatch(deleteTemplate(id));
  };

  const actions = (template) => [
    <Link to={'/templates/' + template.id + '/edit'}>
      <EditOutlined key="edit" />
    </Link>,
    <Popconfirm
      title="Are you sure to delete this template?"
      onConfirm={() => handleDelete(template.id)}
      okText="Yes"
      cancelText="No"
    >
      <DeleteOutlined key="ellipsis" />
    </Popconfirm>,
  ];

  return (
    <List
      style={{ width: '100%' }}
      header={<Title level={2}>{category.name}</Title>}
      grid={{
        gutter: 16,
        xs: 1,
        sm: 2,
        md: 4,
        lg: 6,
        xl: 6,
        xxl: 6,
      }}
      dataSource={templates}
      renderItem={(template) => (
        <List.Item>
          <Card
            hoverable
            cover={
              <img
                alt="example"
                src={template.medium ? template.medium.url.proxy : null}
                height={200}
              />
            }
            actions={location.pathname === routes.templatesList.path ? [] : actions(template)}
          >
            <Link to={'/chart/' + template.id}>
              <Meta title={template.title} />
            </Link>
          </Card>
        </List.Item>
      )}
    />
  );
}

function Templates() {
  const dispatch = useDispatch();
  const categories = useSelector((state) => Object.values(state.categories.details));

  useEffect(() => {
    dispatch(getTemplates());
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <List
      dataSource={categories}
      renderItem={(category) => (
        <List.Item style={{ width: '100%' }}>
          <TemplatesGroup ids={category.template_ids} category={category} />
        </List.Item>
      )}
    />
  );
}

export default Templates;
