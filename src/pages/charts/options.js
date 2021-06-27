import React, { useEffect } from 'react';
import { Collapse } from 'antd';
import { useParams } from 'react-router-dom';

import Line from '../../components/shared/area_lines.js';
import Bars from '../../components/shared/bars.js';
import ChartProperties from '../../components/shared/chart_properties.js';
import Colors from '../../components/shared/colors.js';
import DataLabels from '../../components/shared/data_labels.js';
import Dots from '../../components/shared/dots.js';
import Facet from '../../components/shared/facet.js';
import Legend from '../../components/shared/legend.js';
import Lines from '../../components/shared/lines.js';
import LegendLabel from '../../components/shared/legend_label.js';
import PieDataLabels from '../../components/shared/pie_data_labels.js';
import Segment from '../../components/shared/segment.js';
import XAxis from '../../components/shared/x_axis.js';
import YAxis from '../../components/shared/y_axis.js';
import TreeMap from '../../components/shared/tree_map.js';
import RegionsLayer from '../../components/shared/regions_layer.js';
import GraticuleLayer from '../../components/shared/graticule.js';
import ZoomLayer from '../../components/shared/zoom.js';

import Categories from '../../components/categories';
import Tags from '../../components/tags';
import { useDispatch, useSelector } from 'react-redux';
import { getTemplate } from '../../actions/templates';

const { Panel } = Collapse;

const PropertiesComponentMap = {
  ChartProperties,
  Colors,
  XAxis,
  YAxis,
  DataLabels,
  Dots,
  Line,
  Legend,
  LegendLabel,
  Segment,
  Facet,
  Bars,
  Lines,
  PieDataLabels,
  TreeMap,
  RegionsLayer,
  GraticuleLayer,
  ZoomLayer,
};

function OptionComponent(props) {
  const dispatch = useDispatch();

  let { templateId } = useParams();
  if (!templateId) templateId = props.templateId;
  const component = useSelector(({ templates }) => templates.details[templateId]);

  useEffect(() => {
    if (templateId && !component) dispatch(getTemplate(templateId));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    if (!props.isEdit) form.setFieldsValue({ ...component?.spec, mode: component?.mode });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [component]);

  const { form } = props;

  return !component ? null : (
    <>
      <Collapse className="option-item-collapse">
        {component.properties.$components.map((d, i) => {
          const PropertiesComponent = PropertiesComponentMap[d.Component];
          return PropertiesComponent ? (
            <Panel className="option-item-panel" header={d.name} key={i}>
              <PropertiesComponent properties={d.properties} form={form} />
            </Panel>
          ) : null;
        })}
      </Collapse>
      <br />
      <Collapse defaultActiveKey={['tags', 'categories']} className="option-item-collapse">
        <Panel className="option-item-panel" header={'Tags'} key={'tags'}>
          <Tags form={form} />
        </Panel>
        <Panel className="option-item-panel" header={'Categories'} key={'categories'}>
          <Categories form={form} multiple />
        </Panel>
      </Collapse>
    </>
  );
}

export default OptionComponent;
