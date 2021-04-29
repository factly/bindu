import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import Display from './display.js';
import ChartOption from './options.js';
import { saveAs } from 'file-saver';
import './index.css';
import _ from 'lodash';
import SplitPane from 'react-split-pane';

import { Card, Tooltip, Button, Input, Form, Modal, Dropdown, Menu } from 'antd';
import {
  SaveOutlined,
  SettingOutlined,
  EditOutlined,
  UploadOutlined,
  MenuOutlined,
  DatabaseOutlined,
  AreaChartOutlined,
} from '@ant-design/icons';

import DataViewer from './data_viewer.js';
import UppyUploader from '../../components/uppy';
import { b64toBlob } from '../../utils/file';
import { useParams } from 'react-router';
import { collapseSider } from '../../actions/settings.js';

const IconSize = 20;

const TitleComponent = ({ chartName, setChartName }) => {
  const [isChartNameEditable, setChartNameEditable] = useState(false);

  if (isChartNameEditable) {
    return (
      <div className="chart-name-editable-container">
        <Input
          onPressEnter={() => setChartNameEditable(false)}
          value={chartName}
          onChange={(e) => setChartName(e.target.value)}
        />{' '}
        <Button
          style={{ padding: '4px 0px' }}
          size="medium"
          onClick={() => setChartNameEditable(false)}
          type="primary"
        >
          Save
        </Button>{' '}
      </div>
    );
  } else {
    return (
      <div className="chart-name-container">
        <label className="chart-name">{chartName}</label>
        <EditOutlined style={{ fontSize: IconSize }} onClick={() => setChartNameEditable(true)} />
      </div>
    );
  }
};

function Chart({ data = {}, onSubmit }) {
  const dispatch = useDispatch();
  const { templateId } = useParams();
  const [form] = Form.useForm();

  const [spec, setSpec] = useState({});
  const [showModal, setShowModal] = useState(false);
  const [showOptions, setShowOptions] = useState(true);
  const [chartName, setChartName] = useState('Untitled');
  const [view, setView] = useState(null);
  const [isDataView, setDataView] = useState(false);
  const [values, setValues] = useState([]);
  const [columns, setColumns] = useState([]);

  const splitContainer = React.useRef(null);

  const onDataUpload = (dataDetails) => {
    let values = form.getFieldValue();
    _.unset(values, ['data', 'values']);
    _.set(values, ['data', 'url'], dataDetails.url.raw);
    form.setFieldsValue(values);
  };
  React.useEffect(() => {
    dispatch(collapseSider());
    onValuesChange();
  }, []);

  React.useEffect(() => {
    if (data && data.id) {
      form.setFieldsValue({
        ...data.config,
        categories: data.categories.map((category) => category.id),
        tags: data.tags.map((tag) => tag.id),
      });
      setChartName(data.title);
    }
  }, [data]);

  const downloadSampleData = () => {
    const url = form.getFieldValue(['data', 'url']);
    const values = form.getFieldValue(['data', 'values']);
    if (url) {
      saveAs(url, url.split('/').pop());
    } else if (values) {
      const blob = new Blob([JSON.stringify(values)], { type: 'application/json;charset=utf-8' });
      saveAs(blob, 'sample.json');
    }
  };

  const downloadImage = async (e) => {
    const ext = e.key;
    const data = await view?.toImageURL(ext, 1);
    let blob;
    if (ext === 'svg') {
      blob = data;
    } else {
      blob = b64toBlob(data.split(',')[1], 'image/' + ext);
    }
    saveAs(blob, `${chartName}.${ext}`);
  };

  const saveChart = async () => {
    const { tags, categories, ...values } = form.getFieldValue();
    const imageBlob = await view?.toImageURL('png', 1);

    onSubmit({
      title: chartName,
      data_url: values.data.url,
      config: values,
      featured_medium: imageBlob,
      category_ids: categories,
      tag_ids: tags,
      template_id: Number(templateId),
    });
  };

  const onValuesChange = () => {
    setSpec(form.getFieldValue());
  };

  const IconSize = 20;
  const actions = [
    {
      name: 'Customize',
      Component: (
        <SettingOutlined
          style={{ fontSize: IconSize }}
          onClick={() => setShowOptions(!showOptions)}
        />
      ),
    },
    {
      name: 'Save',
      Component: <SaveOutlined style={{ fontSize: IconSize }} onClick={saveChart} />,
    },
    {
      name: 'Upload',
      Component: (
        <UploadOutlined style={{ fontSize: IconSize }} onClick={() => setShowModal(true)} />
      ),
    },
    {
      name: isDataView ? 'Chart' : 'Data',
      Component: isDataView ? (
        <AreaChartOutlined onClick={() => setDataView(!isDataView)} />
      ) : (
        <DatabaseOutlined onClick={() => setDataView(!isDataView)} />
      ),
    },
    {
      name: 'Options',
      Component: (
        <Dropdown
          overlay={
            <div style={{ boxShadow: '0px 0px 6px 1px #999' }}>
              <Menu onClick={downloadImage}>
                <Menu.Item key="svg">Download Image (SVG)</Menu.Item>
                <Menu.Item key="png">Download Image (PNG)</Menu.Item>
              </Menu>
              <Menu.Divider />
              <Menu onClick={downloadSampleData}>
                <Menu.Item>Download Sample</Menu.Item>
              </Menu>
            </div>
          }
        >
          <Button style={{ marginBottom: 5 }}>
            <MenuOutlined />
          </Button>
        </Dropdown>
      ),
    },
  ];

  const actionsList = (
    <div className="extra-actions-container">
      <ul style={{ display: 'flex', alignItems: 'center' }}>
        {actions.map((item) => (
          <li key={item.name}>
            <Tooltip title={item.name}>{item.Component}</Tooltip>
          </li>
        ))}
      </ul>
    </div>
  );

  let SplitView;
  if (isDataView) {
    let headerRow = {};
    if (spec?.data?.values) {
      const newColumns = Object.keys(spec.data.values[0]).map((d) => {
        headerRow[d] = d;
        return {
          title: d,
          dataIndex: d,
        };
      });

      if (!columns.length) setColumns(newColumns);
      if (!values.length) setValues(spec.data.values);
    } else if (spec?.data?.url) {
      const url = spec.data.url;
      fetch(url)
        .then((res) => res.json())
        .then((newValues) => {
          const newColumns = Object.keys(newValues[0]).map((d) => {
            headerRow[d] = d;
            return {
              title: d,
              dataIndex: d,
            };
          });

          if (!columns.length) setColumns(newColumns);
          if (!values.length) setValues(newValues);
        });
    }

    const { width, height } = splitContainer.current.pane1.getBoundingClientRect();

    SplitView = (
      <SplitPane
        ref={splitContainer}
        pane1Style={{ width: '70%' }}
        style={{ height: 'calc(100% - 48px)' }}
        split="vertical"
      >
        <DataViewer
          columns={columns}
          dataSource={values}
          scroll={{
            y: height - 55,
            x: width,
          }}
        />
        <SplitPane pane1Style={{ height: '50%', height: 'inherit' }} split="horizontal">
          {/* <DataComponent /> */}
          <Display spec={spec} setView={setView} />
        </SplitPane>
      </SplitPane>
    );
  } else {
    SplitView = (
      <SplitPane
        ref={splitContainer}
        pane1Style={{ width: showOptions ? '70%' : '100%', height: 'inherit' }}
        style={{ height: 'calc(100% - 48px)' }}
        split="vertical"
      >
        <Display spec={spec} setView={setView} />
        <SplitPane
          pane1Style={{
            height: 'inherit',
            overflow: 'auto',
            flexDirection: 'column',
            right: showOptions ? '0' : '-400px',
          }}
          split="horizontal"
        >
          <ChartOption form={form} templateId={data ? data.template_id : 0} />
          {/* <div className="extra-options" style={{ padding: '12px' }}>
            <ChartMeta />
          </div> */}
        </SplitPane>
      </SplitPane>
    );
  }

  return (
    <>
      <Form form={form} layout="horizontal" onValuesChange={onValuesChange}>
        <Card
          title={<TitleComponent chartName={chartName} setChartName={setChartName} />}
          extra={actionsList}
          bodyStyle={{ overflow: 'auto', display: 'flex', padding: '0px', height: '80vh' }}
        >
          {SplitView}
        </Card>
      </Form>
      <Modal
        title="Upload Dataset"
        visible={showModal}
        onOk={() => setShowModal(false)}
        okText="Done"
      >
        <UppyUploader onUpload={onDataUpload} />
      </Modal>
    </>
  );
}

export default Chart;
