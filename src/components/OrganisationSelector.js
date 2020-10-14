import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { Select, Avatar } from 'antd';
import { setSelectedOrganisation, getOrganisations } from '../actions/organisation.js';

const { Option } = Select;

function OrganisationSelector() {
  const { details, selected } = useSelector((state) => state.organisation);
  const dispatch = useDispatch();

  React.useEffect(() => {
    dispatch(getOrganisations());
  }, [dispatch]);

  const handleOrganisationChange = (organisation) => {
    dispatch(setSelectedOrganisation(organisation));
  };

  const DEFAULT_IMAGE = 'https://www.tibs.org.tw/images/default.jpg';

  return (
    <Select
      style={{ width: '200px' }}
      value={selected}
      onChange={handleOrganisationChange}
      bordered={false}
    >
      {Object.values(details).map((organisation) => (
        <Option key={'organisation-' + organisation.id} value={organisation.id}>
          <Avatar
            size="small"
            src={organisation.logo ? organisation.logo.url?.raw : DEFAULT_IMAGE}
          />{' '}
          {organisation.name}
        </Option>
      ))}
    </Select>
  );
}

export default OrganisationSelector;
