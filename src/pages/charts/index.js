import React from "react";
import Display from "./display.js";
import Option from "./options.js";
import "./index.css";

function Chart() {
  return (
    <div className="chart-wrapper">
      <Display />
      <Option />
    </div>
  );
}

export default Chart;