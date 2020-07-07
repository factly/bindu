import React, { useEffect } from "react";

import { Collapse } from "antd";
import ChartProperties from "../../../components/shared/chart_properties.js";
import Colors from "../../../components/shared/colors.js";

import Bars from "../../../components/bar/bars.js";
import Lines from "../../../components/line_bar/lines.js";
import Dots from "../../../components/line_bar/dots.js";

import Legend from "../../../components/shared/legend.js";
import LegendLabel from "../../../components/shared/legend_label.js";
import XAxis from "../../../components/shared/x_axis.js";
import YAxis from "../../../components/shared/y_axis.js";
import DataLabels from "../../../components/shared/data_labels.js";
import { useDispatch } from 'react-redux';

import Spec from "./default.json";
const { Panel } = Collapse;

function GroupedBarChart() {
	const dispatch = useDispatch();
	useEffect(() => {
		dispatch({type: "set-config", value: Spec});
	}, [dispatch]);
  
  const properties = [
    {
      name: "Chart Properties",
      component: <ChartProperties />
    },
    {
      name: "Colors",
      component: <Colors />
    },
    {
      name: "Bars",
      component: <Bars />
    },
    {
      name: "Lines",
      component: <Lines />
    },
    {
      name: "Dots",
      component: <Dots />
    },
    {
      name: "X Axis",
      component: <XAxis />
    },
    {
      name: "Y Axis",
      component: <YAxis />
    },
    {
      name: "Legend",
      component: <Legend />
    },
    {
      name: "Legend Label",
      component: <LegendLabel />
    },
    {
      name: "Data Labels",
      component: <DataLabels />
    }
  ];

  return (
    <div className="options-container">
    		<Collapse
          className="option-item-collapse"
        >
          {
            properties.map((d, i) => {
              return (
                <Panel className="option-item-panel" header={d.name} key={i}>
                  {d.component}
                </Panel>
              )
            })
          }
        </Collapse>
    </div>
  );
}

export default GroupedBarChart;
