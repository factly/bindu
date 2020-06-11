import React, { useEffect } from "react";

import { Collapse } from "antd";
import Dimensions from "../../../components/shared/Dimensions.js";
import Bars from "../../../components/bar/Bars.js";
import Legend from "../../../components/bar/Legend.js";
import LegendLabel from "../../../components/bar/LegendLabel.js";
import Properties from "../../../components/bar/Properties.js";
import XAxis from "../../../components/bar/XAxis.js";
import YAxis from "../../../components/bar/YAxis.js";
import { useDispatch } from 'react-redux';

import Spec from "./default.json";
const { Panel } = Collapse;

function GroupedBarChart() {
	const dispatch = useDispatch();
	useEffect(() => {
		dispatch({type: "set-config", value: Spec});
	}, []);

  return (
    <div className="options-container">
    		<Collapse
          className="option-item-collapse"
        >
          <Panel className="option-item-panel" header={"Chart Properties"} key="1">
            <Dimensions />
          </Panel>
          <Panel className="option-item-panel" header={"Bars"} key="2">
            <Bars />
          </Panel>
          <Panel className="option-item-panel" header={"X Axis"} key="3">
            <XAxis />
          </Panel>
          <Panel className="option-item-panel" header={"Y Axis"} key="4">
            <YAxis />
          </Panel>
        </Collapse>
    </div>
  );
}

export default GroupedBarChart;
