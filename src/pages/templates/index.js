import React from "react";
import { useSelector } from 'react-redux';

import { List, Card } from "antd";
import { Link } from 'react-router-dom';

const { Meta } = Card;
function Templates() {
  
  const templates = useSelector(state => state.templates);
  const {options} = templates;
  return (
    <List
    grid={{ gutter: 16, column: 5 }}
    dataSource={options}
    renderItem={(d, i) => (
      <List.Item>
        <Link key={i} to={'/chart/'+i}>
              <Card
                hoverable
                cover={
                  <img
                    alt="example"
                    src={d.icon}
                  />
                }
              >
                <Meta title={d.name} />
              </Card>
            </Link>
      </List.Item>
    )}
  />
  );
}

export default Templates;