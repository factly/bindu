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
            <Link key={index} to={`/charts/${chart.id}/edit`}>
              <Card
                hoverable
                cover={
                  <img
                    width={180}
                    height={150}
                    alt="example"
                    src={chart.medium?.url.url.replace('minio', '127.0.0.1')}
                  />
                }
              >
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
