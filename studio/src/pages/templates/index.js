import React from 'react';
import { Button } from 'antd';
import { Link } from 'react-router-dom';
import { PlusOutlined } from '@ant-design/icons';

import TemplatesList from './list';

function Templates() {
  return (
    <>
      <Link to={'/templates/create'}>
        <Button type="primary" icon={<PlusOutlined />} size={'large'}>
          Create New Template
        </Button>
      </Link>
      <TemplatesList />
    </>
  );
}

export default Templates;
