import React from "react";
import { useSelector } from 'react-redux';

import { Card } from "antd";
import { Link } from 'react-router-dom';

const { Meta } = Card;
function Templates() {
  
  const templates = useSelector(state => state.templates);
  const {options} = templates;
  return (
    <React.Fragment>
      <h2>Choose a Template</h2>
      <div className="example-container">
        {options.map((d, i) => {
          return (
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
          );
        })}
      </div>
    </React.Fragment>
  );
}

export default Templates;