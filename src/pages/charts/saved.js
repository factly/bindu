import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { List, Card } from 'antd';

import { getCharts } from '../../actions/charts';

const { Meta } = Card;
function SavedCharts() {
  const dispatch = useDispatch();
  const { loading, savedCharts } = useSelector(({ charts }) => ({
    loading: charts.loading,
    savedCharts: Object.values(charts.details),
  }));

  React.useEffect(() => {
    dispatch(getCharts());
  }, []);

  return loading ? null : (
    <List
      grid={{ gutter: 16, column: 5 }}
      dataSource={savedCharts}
      renderItem={(chart, index) => {
        return (
          <List.Item>
            <Link key={index}>
              <Card hoverable cover={<img alt="example" src={chart.medium?.url.proxy} />}>
                <Meta title={chart.title} />
              </Card>
            </Link>
          </List.Item>
        );
      }}
    />
  );
}

export default SavedCharts;
