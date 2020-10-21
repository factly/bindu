import React from 'react';
import { Input, Select, Form, Checkbox, InputNumber } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';
import {
  SET_LINE_DOTS,
  SET_LINE_DOT_SHAPE,
  SET_LINE_DOT_SIZE,
  SET_LINE_DOTS_HOLLOW,
} from '../../constants/dots.js';
import _ from 'lodash';
import CheckboxGroup from 'antd/lib/checkbox/Group';

const { Option } = Select;

function Dots(props) {
  const { form } = props;
  const [enable, setEnable] = React.useState(false);

  const markObj = props.properties.find((d) => d.prop === 'mark');

  const dispatch = useDispatch();
  let values = form.getFieldValue([]);
  console.log({ point: _.get(values, [...markObj.path]) });

  const handleEnable = (checked) => {
    if (!checked) {
      let values = form.getFieldValue();
      _.unset(values, [...markObj.path, 'point']);
      form.setFieldsValue(values);
    }
    setEnable(checked);
  };

  return (
    <div className="property-container">
      <Checkbox
        onChange={(e) => {
          const checked = e.target.checked;
          handleEnable(checked);
          dispatch({
            type: SET_LINE_DOTS,
            payload: { value: e.target.checked, path: markObj.path },
            chart: 'shared',
          });
        }}
      >
        Enable
      </Checkbox>

      {enable ? (
        <React.Fragment>
          <Form.Item
            name={[...markObj.path, 'point', 'shape']}
            initialValue={'circle'}
            label="Symbol"
          >
            <Select
              onChange={(value) =>
                dispatch({
                  type: SET_LINE_DOT_SHAPE,
                  payload: { value: value, path: markObj.path },
                  chart: 'shared',
                })
              }
            >
              <Option value="circle">Circle</Option>
              <Option value="square">Square</Option>
              <Option value="cross">Cross</Option>
              <Option value="diamond">Diamond</Option>
              <Option value="triangle-up">Triangle Up</Option>
              <Option value="triangle-down">Triangle Down</Option>
              <Option value="triangle-right">Triangle Right</Option>
              <Option value="triangle-left">Triangle Left</Option>
            </Select>
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'point', 'size']}
            initialValue={40}
            label="Symbol Size"
          >
            <InputNumber
              placeholder="Symbol Size"
              onChange={(value) =>
                dispatch({
                  type: SET_LINE_DOT_SIZE,
                  payload: { value, path: markObj.path },
                  chart: 'shared',
                })
              }
            />
          </Form.Item>

          <Form.Item
            name={[...markObj.path, 'point', 'filled']}
            initialValue={true}
            valuePropName="checked"
            label="Filled"
          >
            <Checkbox />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default Dots;
