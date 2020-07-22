import React from 'react';
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

function Zoom(props) {
  const spec = useSelector((state) => state.chart.spec);

  const scaleObj = props.properties.find((d) => d.prop === 'scale');
  const rotate0Obj = props.properties.find((d) => d.prop === 'rotate0');
  const rotate1Obj = props.properties.find((d) => d.prop === 'rotate1');
  const rotate2Obj = props.properties.find((d) => d.prop === 'rotate2');
  const center0Obj = props.properties.find((d) => d.prop === 'center0');
  const center1Obj = props.properties.find((d) => d.prop === 'center1');

  const scale = getValueFromNestedPath(spec, scaleObj.path);
  const rotate0 = getValueFromNestedPath(spec, rotate0Obj.path);
  const rotate1 = getValueFromNestedPath(spec, rotate1Obj.path);
  const rotate2 = getValueFromNestedPath(spec, rotate2Obj.path);
  const center0 = getValueFromNestedPath(spec, center0Obj.path);
  const center1 = getValueFromNestedPath(spec, center1Obj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Scale</label>
        </Col>
        <Col span={12}>
          <Input
            value={scale}
            placeholder="scale"
            type="number"
            min={50}
            max={2000}
            onChange={(e) =>
              dispatch({
                type: 'set-map-zoom-scale',
                payload: { value: e.target.value, path: scaleObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Rotate 0</label>
        </Col>
        <Col span={12}>
          <Input
            value={rotate0}
            placeholder="rotate0"
            type="number"
            min={-180}
            max={180}
            onChange={(e) =>
              dispatch({
                type: 'set-map-zoom-rotate0',
                payload: { value: e.target.value, path: rotate0Obj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Rotate 1</label>
        </Col>
        <Col span={12}>
          <Input
            value={rotate1}
            placeholder="rotate1"
            min={-90}
            max={90}
            type="number"
            onChange={(e) =>
              dispatch({
                type: 'set-map-zoom-rotate1',
                payload: { value: e.target.value, path: rotate1Obj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Rotate 2</label>
        </Col>
        <Col span={12}>
          <Input
            value={rotate2}
            placeholder="rotate2"
            min={-90}
            max={90}
            type="number"
            onChange={(e) =>
              dispatch({
                type: 'set-map-zoom-rotate2',
                payload: { value: e.target.value, path: rotate2Obj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Center 0</label>
        </Col>
        <Col span={12}>
          <Input
            value={center0}
            placeholder="center0"
            type="number"
            min={-180}
            max={180}
            onChange={(e) =>
              dispatch({
                type: 'set-map-zoom-center0',
                payload: { value: e.target.value, path: center0Obj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Center 1</label>
        </Col>
        <Col span={12}>
          <Input
            value={center1}
            placeholder="center1"
            type="number"
            min={-180}
            max={180}
            onChange={(e) =>
              dispatch({
                type: 'set-map-zoom-center1',
                payload: { value: e.target.value, path: center1Obj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
    </div>
  );
}

export default Zoom;
