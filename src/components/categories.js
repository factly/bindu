import React from 'react';
import { Button, Form, Select } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import { getCategories, addCategory } from '../actions/categories';

function Categories(props) {
  const dispatch = useDispatch();
  const [searchText, setSearchText] = React.useState('');
  const categories = useSelector(({ categories }) => Object.values(categories.details));

  React.useEffect(() => {
    dispatch(getCategories());
  }, []);

  const onCreate = () => {
    dispatch(addCategory({ name: searchText }))
      .then(() => {
        setSearchText('');
        dispatch(getCategories());
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div className="property-container">
      <Form.Item name={'categories'} label="Categories">
        <Select
          showSearch
          mode="multiple"
          placeholder="select categories"
          type="text"
          onSelect={() => setSearchText('')}
          onSearch={setSearchText}
          notFoundContent={
            <Button
              block
              type="dashed"
              style={{
                whiteSpace: 'nowrap',
                overflow: 'hidden',
                textOverflow: 'ellipsis',
              }}
              onClick={onCreate}
            >
              Create {searchText}
            </Button>
          }
        >
          {categories.map((category) => (
            <Select.Option key={category.id} value={category.name}>
              {category.name}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
    </div>
  );
}

export default Categories;
