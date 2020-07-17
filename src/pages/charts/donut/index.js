import React, { useEffect } from "react";

import { Collapse } from "antd";
import ChartProperties from "../../../components/shared/chart_properties.js";
import Colors from "../../../components/shared/colors.js";
import Legend from "../../../components/shared/legend.js";
import LegendLabel from "../../../components/shared/legend_label.js";

import Segment from "../../../components/shared/segment.js";
import DataLabels from "../../../components/shared/pie_data_labels.js";

import { useDispatch } from 'react-redux';

import Spec from "./default.json";
const { Panel } = Collapse;

function PieChar() {
	const dispatch = useDispatch();
	useEffect(() => {
		dispatch({type: "set-config", value: Spec});
	}, [dispatch]);

  const properties = [{
   name: "Chart Properties",
    properties: [
        {
          prop: "title",
          path: ["title"]
        },
        {
          prop: "width",
          path: ["width"]
        },
        {
          prop: "height",
          path: ["height"]
        },
        {
          prop: "background",
          path: ["background"]
        }
    ],
    Component: ChartProperties
  },{
    name: "Colors",
    properties: [
        {
          prop: "color",
          type: "array",
          path: ["layer", 0, "encoding","color","scale", "range"]
        }
      ],
    Component: Colors
  },{
    name: "Segment",
    properties: [
        {
          prop: "inner_radius",
          path: ["layer", 0, "mark","innerRadius"]
        },
        {
          prop: "corner_radius",
          path: ["layer", 0, "mark","cornerRadius"]
        },
        {
          prop: "pad_angle",
          path: ["layer", 0, "mark","padAngle"]
        },
        {
          prop: "outer_radius",
          path: ["layer", 0, "mark","outerRadius"]
        }
      ],
    Component: Segment
  },{
    name: "Legend",
    properties: [
        {
          prop: "title",
          path: ["layer", 0, "encoding", "color", "legend", "title"]
        },
        {
          prop: "fill_color",
          path: ["layer", 0, "encoding", "color", "legend", "fillColor"]
        },
        {
          prop: "symbol_type",
          path: ["layer", 0, "encoding", "color", "legend", "symbolType"]
        },
        {
          prop: "symbol_size",
          path: ["layer", 0, "encoding", "color", "legend", "symbolSize"]
        },
        {
          prop: "orient",
          path: ["layer", 0, "encoding", "color", "legend", "orient"]
        }
      ],
    Component: Legend
  },{
    name: "Legend Label",
    properties: [
        {
          prop: "label_align",
          path: ["layer", 0, "encoding", "color", "legend", "labelAlign"]
        },
        {
          prop: "label_baseline",
          path: ["layer", 0, "encoding", "color", "legend", "labelBaseline"]
        },
        {
          prop: "label_color",
          path: ["layer", 0, "encoding", "color", "legend", "labelColor"]
        }
      ],
    Component: LegendLabel
  },
    {
      name: "Data Labels",
      properties: [
        {
          prop: "color",
          path: ["layer", 1, "encoding", "color", "value"]
        },
        {
          prop: "font_size",
          path: ["layer", 1, "mark", "fontSize"]
        },
        {
          prop: "format",
          path: ["layer", 1, "encoding", "text", "format"]
        }
      ],
      Component: DataLabels
    }];

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

export default PieChar;
