import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory, useParams } from 'react-router-dom';

import { getTemplate, updateTemplate } from '../../actions/templates';
import TemplateForm from './components/TemplateForm';

function EditTemplate() {
  const history = useHistory();
  const { templateId: id } = useParams();

  const dispatch = useDispatch();

  const { template, loading } = useSelector(({ templates }) => {
    return {
      template: templates.details[id] ? templates.details[id] : null,
      loading: templates.loading,
    };
  });

  React.useEffect(() => {
    dispatch(getTemplate(id));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id]);

  if (loading) return null;

  const onUpdate = (values) => {
    dispatch(
      updateTemplate({
        ...values,
        id: template.id,
      }),
    ).then(() => {
      history.push('/templates');
    });
  };

  return <TemplateForm data={template} onSubmit={onUpdate} />;
}

export default EditTemplate;
