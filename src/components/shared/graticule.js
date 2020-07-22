import React from 'react';
import { Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

function Graticule(props) {
  const spec = useSelector((state) => state.chart.spec);

  const longSepObj = props.properties.find((d) => d.prop === 'long-sep');
  const latSepObj = props.properties.find((d) => d.prop === 'lat-sep');
  const colorObj = props.properties.find((d) => d.prop === 'color');
  const widthObj = props.properties.find((d) => d.prop === 'width');
  const opacityObj = props.properties.find((d) => d.prop === 'opacity');
  const strokeDashObj = props.properties.find((d) => d.prop === 'dash');

  const longSep = getValueFromNestedPath(spec, longSepObj.path);
  const latSep = getValueFromNestedPath(spec, latSepObj.path);
  const color = getValueFromNestedPath(spec, colorObj.path);
  const width = getValueFromNestedPath(spec, widthObj.path);
  const opacity = getValueFromNestedPath(spec, opacityObj.path);
  const dash = getValueFromNestedPath(spec, strokeDashObj.path);

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Longitude Separation</label>
        </Col>
        <Col span={12}>
          <Input
            value={longSep}
            placeholder="Long Sep"
            type="number"
            min={0}
            onChange={(e) =>
              dispatch({
                type: 'set-map-graticule-long',
                payload: { value: e.target.value, path: longSepObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Latitude Separation</label>
        </Col>
        <Col span={12}>
          <Input
            value={latSep}
            placeholder="Lat Sep"
            type="number"
            min={0}
            onChange={(e) =>
              dispatch({
                type: 'set-map-graticule-lat',
                payload: { value: e.target.value, path: latSepObj.path },
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
            min={0}
            type="number"
            onChange={(e) =>
              dispatch({
                type: 'set-map-graticule-width',
                payload: { value: e.target.value, path: widthObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Dashed</label>
        </Col>
        <Col span={12}>
          <Input
            value={dash}
            placeholder="dash"
            min={0}
            type="number"
            onChange={(e) =>
              dispatch({
                type: 'set-map-graticule-dash',
                payload: { value: e.target.value, path: strokeDashObj.path },
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
            value={opacity}
            placeholder="opacity"
            type="number"
            min={0}
            max={1}
            step={0.05}
            onChange={(e) =>
              dispatch({
                type: 'set-map-graticule-opacity',
                payload: { value: e.target.value, path: opacityObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
      <Row gutter={[0, 12]}>
        <Col span={12}>
          <label htmlFor="">Color</label>
        </Col>
        <Col span={12}>
          <Input
            value={color}
            type="color"
            onChange={(e) =>
              dispatch({
                type: 'set-map-graticule-color',
                payload: { value: e.target.value, path: colorObj.path },
                chart: 'shared',
              })
            }
          />
        </Col>
      </Row>
    </div>
  );
}

export default Graticule;
