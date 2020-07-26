import React from 'react';
import { Input, Select, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

const { Option } = Select;

function TreeMap(props) {
  const spec = useSelector((state) => state.chart.spec);
  const layoutObj = props.properties.find((d) => d.prop === 'layout');
  const layout = getValueFromNestedPath(spec, layoutObj.path);

  const aspectRatioObj = props.properties.find((d) => d.prop === 'aspect_ratio');
  const aspect_ratio = getValueFromNestedPath(spec, aspectRatioObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Layout Mode</label>
        </Col>
        <Col span={12}>
          <Select
            value={layout}
            onChange={(value) =>
              dispatch({
                type: 'set-tree-map-layout',
                payload: { value: value, path: layoutObj.path },
                chart: 'shared',
              })
            }
          >
            <Option value="squarify">Squarify</Option>
            <Option value="resquarify">Resquarify</Option>
            <Option value="binary">Binary</Option>
            <Option value="slice">Slice</Option>
            <Option value="dice">Dice</Option>
            <Option value="slicedice">Slicedice</Option>
          </Select>
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Aspect Ratio</label>
        </Col>
        <Col span={12}>
          <Input
            value={aspect_ratio}
            placeholder="Aspect Ratio"
            type="number"
            onChange={(e) =>
              dispatch({
                type: 'set-tree-map-aspect-ratio',
                payload: { value: e.target.value, path: aspectRatioObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
    </div>
  );
}

export default TreeMap;
