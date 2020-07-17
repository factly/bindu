import React from 'react';
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_AREA_LINES,
  SET_AREA_LINE_WIDTH,
  SET_AREA_LINE_OPACITY,
  SET_AREA_LINE_CURVE,
  SET_AREA_LINE_DASHED,
} from '../../constants/area_lines.js';
const { Option } = Select;

function Lines(props) {
  const spec = useSelector((state) => state.chart.spec);
  const markObj = props.properties.find((d) => d.prop === 'mark');
  const mark = getValueFromNestedPath(spec, markObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Enable</label>
        </Col>
        <Col span={12}>
          <Input
            type="checkbox"
            onChange={(e) =>
              dispatch({
                type: SET_AREA_LINES,
                payload: { value: e.target.checked, path: markObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      {mark.hasOwnProperty('line') ? (
        <React.Fragment>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Width</label>
            </Col>
            <Col span={12}>
              <Input
                value={mark.line.strokeWidth}
                type="number"
                onChange={(e) =>
                  dispatch({
                    type: SET_AREA_LINE_WIDTH,
                    payload: { value: e.target.value, path: markObj.path },
                    chart: 'shared',
                  })
                }
              />
            </Col>
          </Row>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Opacity</label>
            </Col>
            <Col span={12}>
              <Input
                value={mark.line.opacity}
                min={0}
                max={1}
                step={0.05}
                type="number"
                onChange={(e) =>
                  dispatch({
                    type: SET_AREA_LINE_OPACITY,
                    payload: { value: e.target.value, path: markObj.path },
                    chart: 'shared',
                  })
                }
              />
            </Col>
          </Row>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Line Curve</label>
            </Col>
            <Col span={12}>
              <Select
                value={mark.line.interpolate}
                onChange={(value) =>
                  dispatch({
                    type: SET_AREA_LINE_CURVE,
                    payload: { value: value, path: markObj.path },
                    chart: 'shared',
                  })
                }
              >
                <Option value="linear">Linear</Option>
                <Option value="linear-closed">Linear Closed</Option>
                <Option value="step">Step</Option>
                <Option value="basis">Basis</Option>
                <Option value="monotone">Monotone</Option>
              </Select>
            </Col>
          </Row>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Dash Width</label>
            </Col>
            <Col span={12}>
              <Input
                value={mark.line.strokeDash}
                type="number"
                onChange={(e) =>
                  dispatch({
                    type: SET_AREA_LINE_DASHED,
                    payload: { value: e.target.value, path: markObj.path },
                    chart: 'shared',
                  })
                }
              />
            </Col>
          </Row>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default Lines;
