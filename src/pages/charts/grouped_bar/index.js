import React, { useEffect } from "react";

import { Collapse } from "antd";
import ChartProperties from "../../../components/shared/chart_properties.js";
import Colors from "../../../components/shared/colors.js";
import Bars from "../../../components/shared/bars.js";
import XAxis from "../../../components/shared/x_axis.js";
import YAxis from "../../../components/shared/y_axis.js";

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
      Component: ChartProperties
    },
    {
      name: "Colors",
      properties: [
        {
          prop: "color",
          type: "array",
          path: ["encoding","color","scale", "range"]
        }
      ],
      Component: Colors
    },
    {
      name: "Bars",
      properties: [
        {
          prop: "opacity",
          path: ["encoding","opacity","value"]
        },
        {
          prop: "corner_radius",
          path: ["mark","cornerRadius"]
        }
      ],
      Component: Bars
    },
    {
      name: "X Axis",
      properties: [
        {
          prop: "title",
          path: ["encoding","x","axis", "title"]
        },
        {
          prop: "orient",
          path: ["encoding","x","axis", "orient"]
        },
        {
          prop: "format",
          path: ["encoding","x","axis", "format"]
        },
        {
          prop: "label_color",
          path: ["encoding","x","axis", "labelColor"]
        }
      ],
      Component: XAxis
    },
    {
      name: "Y Axis",
      properties: [
        {
          prop: "title",
          path: ["encoding","y","axis", "title"]
        },
        {
          prop: "orient",
          path: ["encoding","y","axis", "orient"]
        },
        {
          prop: "format",
          path: ["encoding","y","axis", "format"]
        },
        {
          prop: "label_color",
          path: ["encoding","y","axis", "labelColor"]
        }
      ],
      Component: YAxis
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
                  <d.Component properties = {d.properties}/>
                </Panel>
              )
            })
          }
        </Collapse>
    </div>
  );
}

export default GroupedBarChart;
