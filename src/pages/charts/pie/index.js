import React, { useEffect } from "react";

import { Collapse } from "antd";
import ChartProperties from "../../../components/shared/chart_properties.js";
import Bars from "../../../components/bar/bars.js";
import Legend from "../../../components/bar/legend.js";
import LegendLabel from "../../../components/bar/legend_label.js";
import XAxis from "../../../components/bar/x_axis.js";
import YAxis from "../../../components/bar/y_axis.js";
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
            <ChartProperties />
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
