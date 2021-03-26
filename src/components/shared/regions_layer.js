import React from 'react';
import { Select, Input, Row, Col } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { getValueFromNestedPath } from '../../utils/index.js';

const { Option } = Select;
function RegionsLayer(props) {
  const spec = useSelector((state) => state.chart.spec);

  const strokeObj = props.properties.find((d) => d.prop === 'stroke-color');
  const strokeWidthObj = props.properties.find((d) => d.prop === 'stroke-width');
  const strokeOpacityObj = props.properties.find((d) => d.prop === 'stroke-opacity');
  const colorSchemeObj = props.properties.find((d) => d.prop === 'color-scheme');

  let stroke;
  if (strokeObj) {
    stroke = getValueFromNestedPath(spec, strokeObj.path);
  }

  let strokeWidth;
  if (strokeWidthObj) {
    strokeWidth = getValueFromNestedPath(spec, strokeWidthObj.path);
  }

  let strokeOpacity;
  if (strokeOpacityObj) {
    strokeOpacity = getValueFromNestedPath(spec, strokeOpacityObj.path);
  }

  let colorScheme;
  if (colorSchemeObj) {
    colorScheme = getValueFromNestedPath(spec, colorSchemeObj.path);
  }

  const dispatch = useDispatch();

  return (
    <div className="property-container">
      {strokeWidthObj ? (
        <Row gutter={[0, 12]}>
          <Col span={12}>
            <label htmlFor="">Stroke Width</label>
          </Col>
          <Col span={12}>
            <Input
              placeholder="stroke width"
              min={0}
              value={strokeWidth}
              type="number"
              onChange={(e) =>
                dispatch({
                  type: 'set-map-stroke-width',
                  payload: { value: e.target.value, path: strokeWidthObj.path },
                  chart: 'shared',
                })
              }
            />
          </Col>
        </Row>
      ) : null}
      {strokeObj ? (
        <Row gutter={[0, 12]}>
          <Col span={12}>
            <label htmlFor="">Stroke Color</label>
          </Col>
          <Col span={12}>
            <Input
              value={stroke}
              type="color"
              onChange={(e) =>
                dispatch({
                  type: 'set-map-stroke-color',
                  payload: { value: e.target.value, path: strokeObj.path },
                  chart: 'shared',
                })
              }
            />
          </Col>
        </Row>
      ) : null}
      {strokeOpacityObj ? (
        <Row gutter={[0, 12]}>
          <Col span={12}>
            <label htmlFor="">Stroke Opacity</label>
          </Col>
          <Col span={12}>
            <Input
              value={strokeOpacity}
              min={0}
              max={1}
              step={0.05}
              type="number"
              onChange={(e) =>
                dispatch({
                  type: 'set-map-stroke-opacity',
                  payload: { value: e.target.value, path: strokeOpacityObj.path },
                  chart: 'shared',
                })
              }
            />
          </Col>
        </Row>
      ) : null}

      {colorSchemeObj ? (
        <Row gutter={[0, 12]}>
          <Col span={12}>
            <label htmlFor="">Color Scheme</label>
          </Col>
          <Col span={12}>
            <Select
              value={colorScheme}
              onChange={(value) =>
                dispatch({
                  type: 'set-map-color-scheme',
                  payload: { value: value, path: colorSchemeObj.path },
                  chart: 'shared',
                })
              }
            >
              <Option value="top">Top</Option>
              <Option value="blues"> Blues </Option>
              <Option value="tealblues"> Tealblues </Option>
              <Option value="teals"> Teals </Option>
              <Option value="greens"> Greens </Option>
              <Option value="browns"> Browns </Option>
              <Option value="oranges"> Oranges </Option>
              <Option value="reds"> Reds </Option>
              <Option value="purples"> Purples </Option>
              <Option value="warmgreys"> Warmgreys </Option>
              <Option value="greys"> Greys </Option>
              <Option value="viridis"> Viridis </Option>
              <Option value="magma"> Magma </Option>
              <Option value="inferno"> Inferno </Option>
              <Option value="plasma"> Plasma </Option>
              <Option value="bluegreen"> Bluegreen </Option>
              <Option value="bluepurple"> Bluepurple </Option>
              <Option value="goldgreen"> Goldgreen </Option>
              <Option value="goldorange"> Goldorange </Option>
              <Option value="goldred"> Goldred </Option>
              <Option value="greenblue"> Greenblue </Option>
              <Option value="orangered"> Orangered </Option>
              <Option value="purplebluegreen"> Purple Blue Green </Option>
              <Option value="purpleblue"> Purpleblue </Option>
              <Option value="purplered"> Purplered </Option>
              <Option value="redpurple"> Redpurple </Option>
              <Option value="yellowgreenblue"> Yellowgreenblue </Option>
              <Option value="yellowgreen"> Yellowgreen </Option>
              <Option value="yelloworangebrown"> Yelloworangebrown </Option>
              <Option value="yelloworangered"> Yelloworangered </Option>
              <Option value="darkblue"> Darkblue </Option>
              <Option value="darkgold"> Darkgold </Option>
              <Option value="darkgreen"> Darkgreen </Option>
              <Option value="darkmulti"> Darkmulti </Option>
              <Option value="darkred"> Darkred </Option>
              <Option value="lightgreyred"> Lightgreyred </Option>
              <Option value="lightgreyteal"> Lightgreyteal </Option>
              <Option value="lightmulti"> Lightmulti </Option>
              <Option value="lightorange"> Lightorange </Option>
              <Option value="lighttealblue"> Lighttealblue </Option>
              <Option value="blueorange"> Blueorange </Option>
              <Option value="brownbluegreen"> Brownbluegreen </Option>
              <Option value="purplegreen"> Purplegreen </Option>
              <Option value="pinkyellowgreen"> Pinkyellowgreen </Option>
              <Option value="purpleorange"> Purpleorange </Option>
              <Option value="redblue"> Redblue </Option>
              <Option value="redgrey"> Redgrey </Option>
              <Option value="redyellowblue"> Redyellowblue </Option>
              <Option value="redyellowgreen"> Redyellowgreen </Option>
              <Option value="spectral"> Spectral </Option>
              <Option value="rainbow"> Rainbow </Option>
              <Option value="sinebow"> Sinebow </Option>
            </Select>
          </Col>
        </Row>
      ) : null}
    </div>
  );
}

export default RegionsLayer;
