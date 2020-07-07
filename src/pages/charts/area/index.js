import React, { useEffect } from "react";

import { Collapse } from "antd";
import ChartProperties from "../../../components/shared/chart_properties.js";
import Colors from "../../../components/shared/colors.js";
import XAxis from "../../../components/shared/x_axis.js";
import YAxis from "../../../components/shared/y_axis.js";
import DataLabels from "../../../components/shared/data_labels.js";

import Dots from "../../../components/line/dots.js";
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
      name: "Data Labels",
      component: <DataLabels />
    }
  ];

  return (
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
  );
}

export default GroupedBarChart;