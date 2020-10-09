import React from 'react';
import { Input, Form, InputNumber } from 'antd';
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
      <Form.Item name={titleObj.path} label="Title">
        <Input
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
      </Form.Item>

      <Form.Item name={widthObj.path} label="Width">
        <InputNumber
          placeholder="width"
          onChange={(value) => {
            dispatch({
              type: SET_WIDTH,
              payload: { value, path: widthObj.path },
              value: value,
              chart: 'shared',
            });
          }}
        />
      </Form.Item>
      <Form.Item name={heightObj.path} label="Height">
        <InputNumber
          placeholder="height"
          onChange={(value) =>
            dispatch({
              type: SET_HEIGHT,
              payload: { value: value, path: heightObj.path },
              value: value,
              chart: 'shared',
            })
          }
        />
      </Form.Item>
      <Form.Item name={backgroundObj.path} label="Background">
        <Input
          type="color"
          onChange={(e) =>
            dispatch({
              type: SET_BACKGROUND,
              payload: { value: e.target.value, path: backgroundObj.path },
              value: e.target.value,
              chart: 'shared',
            })
          }
        />
      </Form.Item>
    </div>
  );
}

export default Dimensions;
