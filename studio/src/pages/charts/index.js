import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import Display from './display.js';
import ChartOption from './options.js';
import { saveAs } from 'file-saver';
import './index.css';
import _ from 'lodash';

import {
  Card,
  Tooltip,
  Button,
  Input,
  Form,
  Modal,
  Dropdown,
  Menu,
  Popover,
  Typography,
  List,
  Tabs,
} from 'antd';
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
import updateFormData from '../../utils/updateFormData.js';

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

  const [uploadDataFile, setUploadDataFile] = useState('');
  const [showUploadDataModal, setShowUploadDataModal] = useState(false);
  const [showSampleDataModal, setShowSampleDataModal] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [showOptions, setShowOptions] = useState(true);
  const [chartName, setChartName] = useState('Untitled');
  const [view, setView] = useState(null);
  const [isDataView, setDataView] = useState(false);
  const [values, setValues] = useState([]);
  const [columns, setColumns] = useState([]);
  const space_slug = useSelector((state) => state.spaces.details[state.spaces.selected]?.slug);
  const template = useSelector(({ templates }) => templates.details[templateId]);
  const dataViewContainer = React.useRef(null);
  const displayRef = React.useRef(null);
  const containerRef = React.useRef(null);

  const spec = form.getFieldValue();

  const handleVegaLiteUpload = (dataDetails) => {
    let values = form.getFieldValue();
    // Keep only one of values and url. If url exists, then remove values.
    _.unset(values, ['data', 'values']);
    _.set(values, ['data', 'url'], dataDetails.url.raw);
    form.setFieldsValue(values);
    fetch(dataDetails.url.raw)
      .then((res) => res.json())
      .then((newValues) => {
        const newColumns = Object.keys(newValues[0]).map((d) => {
          return {
            title: d,
            dataIndex: d,
          };
        });

        setColumns(newColumns);
        setValues(newValues);
      });
  };

  const handleVegaUpload = (dataDetails) => {
    let values = form.getFieldValue();
    const dataObjIndex = values.data.findIndex((dataObj) => dataObj.name === uploadDataFile);
    if (dataObjIndex === -1) return;
    _.unset(values, ['data', dataObjIndex, 'values']);
    _.set(values, ['data', dataObjIndex, 'url'], dataDetails.url.raw);
    form.setFieldsValue(values);
    setUploadDataFile('');
  };

  const onDataUpload = (dataDetails) => {
    let values = form.getFieldValue();
    switch (values.mode) {
      case 'vega': {
        handleVegaUpload(dataDetails);
        return;
      }
      case 'vega-lite': {
        handleVegaLiteUpload(dataDetails);
        return;
      }
      default: {
        return;
      }
    }
  };

  React.useEffect(() => {
    dispatch(collapseSider());
  }, [template]);

  React.useEffect(() => {
    if (values.length > 0 && columns.length > 0) {
      return;
    }
    const spec = form.getFieldValue();
    switch (spec.mode) {
      case 'vega': {
        setDataForVega(spec.data);
        break;
      }
      case 'vega-lite': {
        setDataForVegaLite(spec.data);
        break;
      }
      default: {
        break;
      }
    }
  }, [isDataView]);

  React.useEffect(() => {
    if (data && data.id) {
      form.setFieldsValue({
        ...data.config,
        categories: data.categories.map((category) => category.id),
        tags: data.tags.map((tag) => tag.id),
      });
      setChartName(data.title);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);

  const saveFileFromUrl = (url) => {
    saveAs(url, url.split('/').pop());
  };

  const saveFileFromValues = (values) => {
    const blob = new Blob([JSON.stringify(values)], {
      type: 'application/json;charset=utf-8',
    });
    saveAs(blob, 'sample.json');
  };

  const downloadData = (data) => {
    const url = data.url;
    const values = data.values;
    if (url) {
      saveFileFromUrl(url);
    } else if (values) {
      saveFileFromValues(values);
    }
  };

  const downloadSampleData = () => {
    const spec = form.getFieldValue();
    switch (spec.mode) {
      case 'vega': {
        setShowSampleDataModal(true);
        return;
      }
      case 'vega-lite': {
        const data = form.getFieldValue('data');
        downloadData(data);
        return;
      }
      default: {
        return;
      }
    }
  };

  const uploadData = (name) => {
    setUploadDataFile(name);
    setShowModal(true);
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

  const constructData = async (data) => {
    // If spec contains values, remove it and push it to minio. Then, set url of that file in spec
    let url = data.url;
    const path =
      space_slug +
      '/' +
      new Date().getFullYear() +
      '/' +
      new Date().getMonth() +
      '/' +
      Date.now().toString() +
      '_';

    if (data.values) {
      // make a json file out of values
      if (data.url) {
        // replace file in minio at location `data.url`

        url = await updateFormData(
          data,
          data.url.includes('http://localhost:9000/dega/')
            ? data.url.replace('http://localhost:9000/dega/', '')
            : path + chartName,
        );
      } else {
        // upload file to minio
        url = await updateFormData(data, path + chartName);
      }
      // send uploaded file url in api
    }
    _.set(data, 'url', url);
    _.unset(data, 'values');
    return data;
  };

  const constructDataForVega = async (data) => {
    return await Promise.all(
      data.map(async (dataObj) => {
        if (!(dataObj.url || dataObj.values)) return dataObj;
        return await constructData(dataObj);
      }),
    );
  };

  const constructDataForVegaLite = async (data) => {
    return await constructData(data);
  };

  const handleSaveChart = async (e) => {
    const spec = form.getFieldValue();
    const data = form.getFieldValue('data');
    let updatedData = data;
    switch (spec.mode) {
      case 'vega': {
        updatedData = await constructDataForVega(data);
        break;
      }
      case 'vega-lite': {
        updatedData = await constructDataForVegaLite(data);
        break;
      }
      default: {
        break;
      }
    }
    console.log('handleSaveChart', { spec, updatedData });
    _.set(spec, ['data'], updatedData);
    saveChart(e, spec);
  };

  const saveChart = async (e, formData) => {
    const { tags, categories, ...values } = formData;
    const imageBlob = await view?.toImageURL('png', 1);
    // const svg = await view.toSVG();

    onSubmit({
      title: chartName,
      data_url: '', // TODO: handle this for vega and vega-lite
      config: { ...values },
      featured_medium: imageBlob,
      category_ids: categories,
      tag_ids: tags,
      template_id: data.id ? data.template_id : templateId,
      status: e.key,
      is_public: e.key === 'publish',
      published_date: e.key === 'publish' ? new Date() : null,
    });
  };

  const vegaLiteDataChange = ({ fromRow, toRow, updated }) => {
    const updatedValues = [...values];
    for (let i = fromRow; i <= toRow; i++) {
      updatedValues[i] = { ...updatedValues[i], ...updated };
    }
    setValues(updatedValues);

    let formData = form.getFieldValue();
    _.set(formData, ['data', 'values'], updatedValues);
    form.setFieldsValue(formData);
  };

  const vegaDataChange = ({ fromRow, toRow, updated }, tabIndex) => {
    const updatedValues = [...values];
    for (let i = fromRow; i <= toRow; i++) {
      updatedValues[tabIndex][i] = { ...updatedValues[tabIndex][i], ...updated };
    }
    setValues(updatedValues);

    const data = spec.data;
    const updatedValuesClone = [...updatedValues];
    const updatedDataList = data.map((dataObj) => {
      if (!dataObj.url && !dataObj.values) {
        return dataObj;
      }
      const nextValues = updatedValuesClone.shift();
      if (dataObj.url) {
        _.unset(dataObj, ['url']);
      }
      _.set(dataObj, ['values'], nextValues);
      return dataObj;
    });

    let formData = form.getFieldValue();
    _.set(formData, ['data'], updatedDataList);
    form.setFieldsValue(formData);
  };

  const onDataChange = ({ fromRow, toRow, updated }, tabIndex) => {
    try {
      switch (spec.mode) {
        case 'vega': {
          vegaDataChange({ fromRow, toRow, updated }, tabIndex);
          break;
        }
        case 'vega-lite': {
          vegaLiteDataChange({ fromRow, toRow, updated });
          break;
        }
        default: {
          break;
        }
      }
    } catch (error) {
      console.error(error);
    }
  };

  const handleUploadDataClick = () => {
    switch (spec.mode) {
      case 'vega': {
        setShowUploadDataModal(true);
        return;
      }
      case 'vega-lite': {
        setShowModal(true);
        return;
      }
      default: {
        return;
      }
    }
  };

  const IconSize = 20;
  let actions = [
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
      Component: (
        <Dropdown
          overlay={
            <div style={{ boxShadow: '0px 0px 6px 1px #999' }}>
              <Menu onClick={handleSaveChart}>
                <Menu.Item key="draft">Draft</Menu.Item>
                <Menu.Item key="publish">Publish</Menu.Item>
              </Menu>
            </div>
          }
        >
          <SaveOutlined style={{ fontSize: IconSize }} onClick={() => setShowOptions(false)} />
        </Dropdown>
      ),
    },
    {
      name: 'Upload',
      Component: <UploadOutlined style={{ fontSize: IconSize }} onClick={handleUploadDataClick} />,
    },
    {
      name: isDataView ? 'Chart' : 'Data',
      Component: isDataView ? (
        <AreaChartOutlined
          style={{ fontSize: IconSize }}
          onClick={() => setDataView(!isDataView)}
        />
      ) : (
        <DatabaseOutlined style={{ fontSize: IconSize }} onClick={() => setDataView(!isDataView)} />
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

  if (data.id) {
    actions = [
      {
        name: '',
        Component: (
          <Popover
            content={
              <div style={{ width: 300, height: 'auto' }}>
                <Typography.Link
                  ellipsis
                  href={window.BINDU_CHART_VISUALIZATION_URL + '/' + data.id}
                  target="_blank"
                  rel="noreferrer"
                  style={{
                    overflow: 'hidden',
                    textOverflow: 'ellipsis',
                    display: 'block',
                  }}
                >
                  {window.BINDU_CHART_VISUALIZATION_URL + '/' + data.id}
                </Typography.Link>
                <br />
                <Typography.Text strong>Embed:</Typography.Text>
                <Typography.Text
                  copyable
                  ellipsis={{ rows: 1 }}
                  style={{
                    border: '1px solid',
                    padding: 4,
                    overflow: 'auto',
                    width: '100%',
                  }}
                >
                  {`<div class="factly-embed" data-src=${data.id}><script src="http://localhost:7002/resources/embed.js"></script></div>`}
                </Typography.Text>
              </div>
            }
            title="Export and publish"
            trigger="click"
          >
            <Button style={{ marginBottom: 5 }}>Export</Button>
          </Popover>
        ),
      },
      ...actions,
    ];
  }

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

  const getData = async (data) => {
    if (data?.values) {
      const newColumns = Object.keys(data.values[0]).map((d) => {
        return {
          name: d,
          key: d,
        };
      });

      return [newColumns, data.values];
    } else if (data?.url) {
      const url = data.url;
      const res = await fetch(url);
      const newValues = await res.json();

      const newColumns = Object.keys(newValues[0]).map((d) => {
        return {
          name: d,
          key: d,
        };
      });

      return [newColumns, newValues];
    } else {
      return [[], []];
    }
  };

  const setDataForVegaLite = async (data) => {
    const [newColumns, newValues] = await getData(data);
    if (!columns.length) setColumns(newColumns);
    if (!values.length) setValues(newValues);
  };

  const setDataForVega = async (data) => {
    const promises = data
      ?.filter((dataObj) => dataObj.url || dataObj.values)
      .map(async (dataObj) => {
        const [newColumns, newValues] = await getData(dataObj);
        setValues([...values, newValues]);
        setColumns([...columns, newColumns]);
      });
    await Promise.all(promises);
    return;
  };

  const { width: containerWidth = 0, height: containerHeight = 0 } =
    containerRef?.current?.getBoundingClientRect() || {};
  const { width: displayWidth = 0, height: displayHeight = 0 } =
    displayRef?.current?.getBoundingClientRect() || {};

  return (
    <>
      <div ref={containerRef} className="chart-area-container">
        <Form form={form} layout="horizontal">
          <Card
            title={<TitleComponent chartName={chartName} setChartName={setChartName} />}
            extra={actionsList}
            bodyStyle={{ overflow: 'auto', display: 'flex', padding: '0px', height: '80vh' }}
          >
            <div className="display-container" ref={displayRef}>
              <Form.Item noStyle shouldUpdate={true}>
                {(form) => {
                  return (
                    <Display
                      spec={form.getFieldValue()}
                      mode={form.getFieldValue().mode}
                      setView={setView}
                    />
                  );
                }}
              </Form.Item>
            </div>
            {isDataView ? (
              <div ref={dataViewContainer} className="data-view-container">
                {spec.mode === 'vega' ? (
                  <Tabs tabPosition={'left'}>
                    {spec.data
                      ?.filter((dataObj) => dataObj.url || dataObj.values)
                      .map((dataObj, index) => (
                        <Tabs.TabPane tab={dataObj.name} key={index}>
                          <DataViewer
                            columns={columns[index] || []}
                            dataSource={values[index] || []}
                            onDataChange={onDataChange}
                            tableWidth={containerWidth - displayWidth}
                            tableHeight={displayHeight - 40}
                            tabIndex={index}
                          />
                        </Tabs.TabPane>
                      ))}
                  </Tabs>
                ) : (
                  <DataViewer
                    columns={columns}
                    dataSource={values}
                    onDataChange={onDataChange}
                    tableWidth={containerWidth - displayWidth}
                    tableHeight={displayHeight - 40}
                  />
                )}
              </div>
            ) : (
              <div className="option-container">
                <ChartOption form={form} templateId={data.template_id} isEdit={!!data.id} />
              </div>
            )}
          </Card>
        </Form>
      </div>
      <Modal
        title="Upload Dataset"
        visible={showModal}
        onCancel={() => setShowModal(false)}
        onOk={() => setShowModal(false)}
        okText="Done"
      >
        <UppyUploader onUpload={onDataUpload} />
      </Modal>
      {spec.mode === 'vega' && (
        <Modal
          title="Download Sample Data"
          visible={showSampleDataModal}
          onCancel={() => setShowSampleDataModal(false)}
          onOk={() => setShowSampleDataModal(false)}
          okText="Done"
        >
          <List
            size="large"
            bordered
            dataSource={spec.data?.filter((dataObj) => dataObj.url || dataObj.values)}
            renderItem={(dataObj) => (
              <List.Item actions={[<a onClick={() => downloadData(dataObj)}>Download</a>]}>
                {dataObj.name}
              </List.Item>
            )}
          />
        </Modal>
      )}
      {spec.mode === 'vega' && (
        <Modal
          title="Download Sample Data"
          visible={showUploadDataModal}
          onCancel={() => setShowUploadDataModal(false)}
          onOk={() => setShowUploadDataModal(false)}
          okText="Done"
        >
          <List
            size="large"
            bordered
            dataSource={spec.data?.filter((dataObj) => dataObj.url || dataObj.values)}
            renderItem={(dataObj) => (
              <List.Item actions={[<a onClick={() => uploadData(dataObj.name)}>Upload</a>]}>
                {dataObj.name}
              </List.Item>
            )}
          />
        </Modal>
      )}
    </>
  );
}

export default Chart;
