import React from 'react';
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

import {
  SET_TITLE,
  SET_WIDTH,
  SET_HEIGHT,
  SET_BACKGROUND,
} from '../../constants/chart_properties.js';

function Dimensions(props) {
  const spec = useSelector((state) => state.chart.spec);

  const titleObj = props.properties.find((d) => d.prop === 'title');
  const widthObj = props.properties.find((d) => d.prop === 'width');
  const heightObj = props.properties.find((d) => d.prop === 'height');
  const backgroundObj = props.properties.find((d) => d.prop === 'background');

  const title = getValueFromNestedPath(spec, titleObj.path);
  const width = getValueFromNestedPath(spec, widthObj.path);
  const height = getValueFromNestedPath(spec, heightObj.path);
  const background = getValueFromNestedPath(spec, backgroundObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Title</label>
        </Col>
        <Col span={12}>
          <Input
            value={title}
            placeholder="title"
            type="text"
            onChange={(e) =>
              dispatch({
                type: SET_TITLE,
                payload: { value: e.target.value, path: titleObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Width</label>
        </Col>
        <Col span={12}>
          <Input
            value={width}
            placeholder="width"
            type="number"
            onChange={(e) =>
              dispatch({
                type: SET_WIDTH,
                payload: { value: e.target.value, path: widthObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Height</label>
        </Col>
        <Col span={12}>
          <Input
            value={height}
            placeholder="height"
            type="number"
            onChange={(e) =>
              dispatch({
                type: SET_HEIGHT,
                payload: { value: e.target.value, path: heightObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Background</label>
        </Col>
        <Col span={12}>
          <Input
            value={background}
            type="color"
            onChange={(e) =>
              dispatch({
                type: SET_BACKGROUND,
                payload: { value: e.target.value, path: backgroundObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
    </div>
  );
}

export default Dimensions;
