import React from 'react';
import { Form, Checkbox, InputNumber } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import {
  SET_PIE_DATA_LABELS,
  SET_PIE_DATA_LABELS_SIZE,
  SET_PIE_DATA_LABELS_POSITION,
} from '../../constants/pie_data_labels.js';

function DataLabels(props) {
  const [enable, setEnable] = React.useState(false);

  const spec = useSelector((state) => state.chart.spec);
  const layersLength = spec.layer.length;

  const dispatch = useDispatch();

  const handleEnable = (checked) => {
    if (checked) {
      console.log('checked');
      props.addLayer(({ encoding, mark }) => ({
        mark: { type: 'text', fontSize: 12, radius: mark.outerRadius + 10 },
        encoding: {
          theta: { ...encoding.theta, stack: true },
          text: { field: encoding.color.field, type: encoding.color.type },
          color: { ...encoding.color, legend: null },
        },
      }));
    } else {
      props.removeLayer();
    }
    setEnable(checked);
  };

  return (
    <div className="property-container">
      <Checkbox
        onChange={(e) => {
          const checked = e.target.checked;
          handleEnable(checked);
          dispatch({ type: SET_PIE_DATA_LABELS, value: e.target.checked, chart: 'shared' });
        }}
      >
        Enable
      </Checkbox>
      {enable ? (
        <React.Fragment>
          <Form.Item name={['layer', 1, 'mark', 'fontSize']} label="Size">
            <InputNumber
              onChange={(value) =>
                dispatch({
                  type: SET_PIE_DATA_LABELS_SIZE,
                  value: value,
                  chart: 'shared',
                })
              }
            />
          </Form.Item>

          <Form.Item name={['layer', 1, 'mark', 'radius']} label="Position( from center)">
            <InputNumber
              onChange={(value) =>
                dispatch({
                  type: SET_PIE_DATA_LABELS_POSITION,
                  value: value,
                  chart: 'shared',
                })
              }
            />
          </Form.Item>
        </React.Fragment>
      ) : null}
    </div>
  );
}

export default DataLabels;
