import React from 'react';
import { Select, Input, Form, InputNumber } from 'antd';

const { Option } = Select;
function RegionsLayer(props) {
  const strokeObj = props.properties.find((d) => d.prop === 'stroke-color');
  const strokeWidthObj = props.properties.find((d) => d.prop === 'stroke-width');
  const strokeOpacityObj = props.properties.find((d) => d.prop === 'stroke-opacity');
  const colorSchemeObj = props.properties.find((d) => d.prop === 'color-scheme');

  return (
    <div className="property-container">
      {strokeWidthObj ? (
        <Form.Item name={strokeWidthObj.path} label="Stroke Width">
          <InputNumber
            formatter={(value) => parseInt(value) || 0}
            parser={(value) => parseInt(value) || 0}
            min={0}
            placeholder="stroke width"
          />
        </Form.Item>
      ) : null}
      {strokeObj ? (
        <Form.Item name={strokeObj.path} label="Stroke Color">
          <Input type="color"></Input>
        </Form.Item>
      ) : null}
      {strokeOpacityObj ? (
        <Form.Item name={strokeOpacityObj.path} label="Stroke Opacity">
          <InputNumber
            formatter={(value) => parseInt(value) || 0}
            parser={(value) => parseInt(value) || 0}
            min={0}
            max={1}
            step={0.05}
          />
        </Form.Item>
      ) : null}

      {colorSchemeObj ? (
        <Form.Item name={colorSchemeObj.path} label="Color Scheme">
          <Select>
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
        </Form.Item>
      ) : null}
    </div>
  );
}

export default RegionsLayer;
