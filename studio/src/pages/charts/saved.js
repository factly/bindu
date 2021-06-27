import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { List, Card, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';

import { getCharts } from '../../actions/charts';
import routes from '../../config/routesConfig';

const { Meta } = Card;
function SavedCharts() {
  const dispatch = useDispatch();
  const { loading, savedCharts } = useSelector(({ charts }) => ({
    loading: charts.loading,
    savedCharts: Object.values(charts.details),
  }));

  React.useEffect(() => {
    dispatch(getCharts());
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <List
      header={
        <Link to={routes.templatesList.path}>
          <Button type="primary" icon={<PlusOutlined />} size={'large'}>
            Create New Chart
          </Button>
        </Link>
      }
      loading={loading}
      grid={{ gutter: 16, column: 5 }}
      dataSource={savedCharts}
      renderItem={(chart, index) => {
        return (
          <List.Item>
            <div style={{ position: 'relative' }}>
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
                <Link key={index} to={`/charts/${chart.id}/edit`}>
                  <Meta title={chart.title} />
                </Link>
              </Card>
              {chart.is_public && (
                <div
                  style={{
                    position: 'absolute',
                    backgroundColor: '#108ee9',
                    color: '#fff',
                    top: 0,
                    left: 0,
                    paddingLeft: 6,
                    paddingRight: 6,
                    borderRadius: 20,
                    boxShadow: '0px 0px 10px 2px #aaa',
                  }}
                >
                  Public
                </div>
              )}
            </div>
          </List.Item>
        );
      }}
    />
  );
}

export default SavedCharts;
