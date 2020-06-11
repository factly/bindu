import React from "react";
import { useDispatch, useSelector } from 'react-redux';

import { Card } from "antd";
import Chart from "../charts/index.js";

function Templates() {
  const dispatch = useDispatch();
  
  const templates = useSelector(state => state.templates);
  const {options, selectedOption} = templates;


  if (selectedOption === -1) {
    return (
      <React.Fragment>
        <h2>Choose a Template</h2>
        <div className="example-container">
          {options.map((d, i) => {
            return (
              <Card
                onClick={() => dispatch({ type: "set-chart", index: i })}
                key={i}
                hoverable
                cover={
                  <img
                    alt="example"
                    src={d.icon}
                  />
                }
              >
                {d.name}
              </Card>
            );
          })}
        </div>
      </React.Fragment>
    );
  } else {
    return <Chart />;
  }
}

export default Templates;