import React from 'react';
import { Input, Row, Col } from 'antd';
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

  const dispatch = useDispatch();
  const payload = {
    path: colorObj.path,
    type: colorObj.type,
  };

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Colors</label>
        </Col>
        <Col span={12}>
          {colors &&
            colors.map((d, i) => (
              <Input
                type="color"
                value={d}
                key={i}
                onChange={(e) =>
                  dispatch({
                    type: SET_COLOR,
                    payload: { index: i, value: e.target.value, ...payload },
                    chart: 'shared',
                  })
                }
              ></Input>
            ))}
        </Col>
      </Row>
    </div>
  );
}

export default Colors;
