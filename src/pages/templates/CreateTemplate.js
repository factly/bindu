import React from 'react';
import { useDispatch } from 'react-redux';
import { addTemplate } from '../../actions/templates';
import { useHistory } from 'react-router-dom';
import TemplateForm from './components/TemplateForm';

function CreateTemplate() {
  const history = useHistory();

  const dispatch = useDispatch();
  const onCreate = (values) => {
    dispatch(addTemplate(values)).then(() => history.push('/templates'));
  };
  return <TemplateForm onSubmit={onCreate} />;
}

export default CreateTemplate;
