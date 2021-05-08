import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { List, Card, Popconfirm, Button, Typography } from 'antd';
import { Link } from 'react-router-dom';
import { EditOutlined, DeleteOutlined, PlusOutlined } from '@ant-design/icons';

import { deleteTemplate, getTemplates } from '../../actions/templates';

const { Meta } = Card;
const { Title } = Typography;

function TemplatesGroup({ ids = [], category = {} }) {
  const dispatch = useDispatch();
  const templates = useSelector(({ templates }) => ids.map((id) => templates.details[id]));

  const handleDelete = (id) => {
    dispatch(deleteTemplate(id));
  };

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
            actions={[
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
            ]}
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
      header={
        <Link to={'/templates/create'}>
          <Button type="primary" icon={<PlusOutlined />} size={'large'}>
            Add new
          </Button>
        </Link>
      }
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
