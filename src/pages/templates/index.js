import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';

import { List, Card, Popconfirm, Button } from 'antd';
import { Link } from 'react-router-dom';
import { EditOutlined, DeleteOutlined, PlusOutlined } from '@ant-design/icons';

import { deleteTemplate, getTemplates } from '../../actions/templates';

const { Meta } = Card;

function Templates() {
  const dispatch = useDispatch();
  const templates = useSelector((state) => Object.values(state.templates.details));

  useEffect(() => {
    dispatch(getTemplates());
  }, []);

  const handleDelete = (id) => {
    dispatch(deleteTemplate(id));
  };

  return (
    <List
      header={
        <Link to={'/templates/create'}>
          <Button type="primary" icon={<PlusOutlined />} size={'large'}>
            Add new
          </Button>
        </Link>
      }
      grid={{ gutter: 16, column: 5 }}
      dataSource={templates}
      renderItem={(template) => (
        <List.Item>
          <Card
            hoverable
            cover={<img alt="example" src={template.icon} />}
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

export default Templates;
