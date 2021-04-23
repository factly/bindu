import React from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { Form } from 'antd';

import TemplateForm from './components/TemplateForm';
import Display from '../charts/display';
import { addTemplate } from '../../actions/templates';
import { collapseSider } from '../../actions/settings';

function CreateTemplate() {
  const [spec, setSpec] = React.useState({});
  const [form] = Form.useForm();

  const history = useHistory();
  const dispatch = useDispatch();

  React.useEffect(() => {
    setSpec(form.getFieldValue('schema'));
  }, [form]);

  React.useEffect(() => {
    dispatch(collapseSider());
  }, []);

  const onCreate = (values) => {
    dispatch(addTemplate(values)).then(() => history.push('/templates'));
  };

  return (
    <div style={{ display: 'flex', height: '80vh' }}>
      <div style={{ flex: 1, height: '100%', overflow: 'auto' }}>
        <Display spec={spec} />
      </div>
      <div style={{ flex: 1, height: '100%', overflow: 'auto' }}>
        <TemplateForm form={form} onSubmit={onCreate} />
      </div>
    </div>
  );
}

export default CreateTemplate;
