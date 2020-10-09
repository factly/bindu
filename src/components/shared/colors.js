import React from 'react';
import { Input, Form } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import { getValueFromNestedPath } from '../../utils/index.js';

import { SET_COLOR } from '../../constants/colors.js';

function Colors(props) {
  const spec = useSelector((state) => state.chart.spec);
  const colorObj = props.properties.find((d) => d.prop === 'color');

  let colors = getValueFromNestedPath(spec, colorObj.path);

  if (colorObj.type === 'string') {
    colors = [colors];
  }

  const getName = (index) => {
    if (colorObj.type === 'string') {
      return colorObj.path;
    }

    return [...colorObj.path, index];
  };

  const dispatch = useDispatch();
  const payload = {
    path: colorObj.path,
    type: colorObj.type,
  };

  return (
    <div className="property-container">
      {colors &&
        colors.map((d, i) => (
          <Form.Item name={getName(i)} label="Colors">
            <Input
              type="color"
              key={i}
              onChange={(e) =>
                dispatch({
                  type: SET_COLOR,
                  payload: { index: i, value: e.target.value, ...payload },
                  chart: 'shared',
                })
              }
            ></Input>
          </Form.Item>
        ))}
    </div>
  );
}

export default Colors;
