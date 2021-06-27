import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory, useParams } from 'react-router-dom';

import { getTemplate, updateTemplate } from '../../actions/templates';
import TemplateForm from './components/TemplateForm';
import Display from '../charts/display';

function EditTemplate() {
  const history = useHistory();
  const { templateId: id } = useParams();
  const [spec, setSpec] = React.useState({});
  const [mode, setMode] = React.useState('');

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

  React.useEffect(() => {
    if (template?.spec) setSpec(template.spec);
    if (template?.mode) setMode(template.mode);
  }, [template]);

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
  const onChange = (values) => {
    if (values.spec) {
      try {
        const spec = JSON.parse(values.spec);
        setSpec(spec);
      } catch {
        console.log('Spec is not JSON');
      }
    }
  };
  return (
    <div style={{ display: 'flex', height: '80vh' }}>
      <div style={{ flex: 1, height: '100%', overflow: 'auto' }}>
        <Display spec={spec} mode={mode} />
      </div>
      <div style={{ flex: 1, height: '100%', overflow: 'auto' }}>
        <TemplateForm
          onSubmit={onUpdate}
          onChange={onChange}
          data={template}
          onModeChange={setMode}
        />
      </div>
    </div>
  );
}

export default EditTemplate;
