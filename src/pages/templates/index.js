import React, { useState } from "react";
import { Card } from "antd";
import Chart from "../charts/index.js";

const charts = [
  {
    name: "Grouped Bar Chart",
    path: "GroupedBarChart",
  }
];

function Templates() {
  const [chartSelectedIndex, setchartSelectedIndex] = useState(-1);

  if (chartSelectedIndex === -1) {
    return (
      <>
        <h2>Choose a Template</h2>
        <div className="example-container">
          {charts.map((d, i) => {
            return (
              <Card
                onClick={() => setchartSelectedIndex(i)}
                key={i}
                hoverable
                cover={
                  <img
                    alt="example"
                    src="https://public.flourish.studio/uploads/ebed7412-cc56-4ef0-a3ba-97865643f545.png"
                  />
                }
              >
                {d.name}
              </Card>
            );
          })}
        </div>
      </>
    );
  } else {
    return <Chart path={charts[chartSelectedIndex].path} />;
  }
}

export default Templates;
