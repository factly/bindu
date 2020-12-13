import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory, useParams } from 'react-router-dom';

import { getChart, updateChart } from '../../actions/charts';
import EditChart from './index';

function EditCharts() {
  const history = useHistory();
  const { id } = useParams();

  const dispatch = useDispatch();

  const { chart, loading } = useSelector((state) => {
    return {
      chart: state.charts.details[id] ? state.charts.details[id] : null,
      loading: state.charts.loading,
    };
  });

  React.useEffect(() => {
    dispatch(getChart(id));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [id]);

  if (loading) return null;

  const onUpdate = (values) => {
    dispatch(
      updateChart({
        ...values,
        id: chart.id,
      }),
    ).then(() => {
      history.push('/charts/saved');
    });
  };

  return <EditChart data={chart} onSubmit={onUpdate} />;
}

export default EditCharts;
