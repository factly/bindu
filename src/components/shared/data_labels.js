import React from 'react';
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_DATA_LABELS,
  SET_DATA_LABELS_COLOR,
  SET_DATA_LABELS_SIZE,
  SET_DATA_LABELS_FORMAT,
} from '../../constants/data_labels.js';

function DataLabels(props) {
  const spec = useSelector((state) => state.chart.spec);
  const layersLength = spec.layer.length;

  const dispatch = useDispatch();

  const colorObj = props.properties.find((d) => d.prop === 'color');
  const fontSizeObj = props.properties.find((d) => d.prop === 'font_size');
  const formatObj = props.properties.find((d) => d.prop === 'format');

  let color;
  let fontSize;
  let format;
  if (layersLength > 1 && colorObj.path) {
    color = getValueFromNestedPath(spec, colorObj.path);
  }
  if (layersLength > 1 && fontSizeObj.path) {
    fontSize = getValueFromNestedPath(spec, fontSizeObj.path);
  }
  if (layersLength > 1 && formatObj.path) {
    format = getValueFromNestedPath(spec, formatObj.path);
  }

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
              dispatch({ type: SET_DATA_LABELS, value: e.target.checked, chart: 'shared' })
            }
          />
        </Col>
      </Row>
      {layersLength > 1 ? (
        <React.Fragment>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Color</label>
            </Col>
            <Col span={12}>
              <Input
                type="color"
                value={color}
                onChange={(e) =>
                  dispatch({
                    type: SET_DATA_LABELS_COLOR,
                    payload: { value: e.target.value, path: colorObj.path },
                    chart: 'shared',
                  })
                }
              />
            </Col>
          </Row>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Size</label>
            </Col>
            <Col span={12}>
              <Input
                type="number"
                value={fontSize}
                onChange={(e) =>
                  dispatch({
                    type: SET_DATA_LABELS_SIZE,
                    payload: { value: e.target.value, path: fontSizeObj.path },
                    chart: 'shared',
                  })
                }
              />
            </Col>
          </Row>
          <Row gutter={[0, 12]}>
            <Col span={12}>
              <label htmlFor="">Format</label>
            </Col>
            <Col span={12}>
              <Input
                type="text"
                value={format}
                onChange={(e) =>
                  dispatch({
                    type: SET_DATA_LABELS_FORMAT,
                    payload: { value: e.target.value, path: formatObj.path },
                    chart: 'shared',
                  })
                }
              />
            </Col>
          </Row>
          <div className="item-container"></div>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default DataLabels;
