import React from 'react';
import { useDispatch } from 'react-redux';
import { useHistory, useParams } from 'react-router-dom';

import { addChart } from '../../actions/charts';
import CreateChart from './index';

function CreateCharts() {
  const history = useHistory();
  const { templateId } = useParams();

  const dispatch = useDispatch();

  const onCreate = (values) => {
    dispatch(
      addChart({
        ...values,
        template_id: templateId,
      }),
    ).then(() => {
      history.push('/charts/saved');
    });
  };

  return <CreateChart onSubmit={onCreate} />;
}

export default CreateCharts;
