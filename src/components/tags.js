import React from 'react';
import { Button, Form, Select, Empty } from 'antd';
import { useDispatch, useSelector } from 'react-redux';

import { getTags, addTag } from '../actions/tags';

function Tags(props) {
  const dispatch = useDispatch();
  const [searchText, setSearchText] = React.useState('');
  const { tags } = useSelector(({ tags }) => ({
    tags: Object.values(tags.details),
  }));

  React.useEffect(() => {
    dispatch(getTags());
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const onCreate = () => {
    if (!searchText.trim()) return;
    dispatch(addTag({ name: searchText }))
      .then(() => {
        setSearchText('');
        dispatch(getTags());
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div className="property-container">
      <Form.Item name={'tags'}>
        <Select
          showSearch
          mode="multiple"
          placeholder="select tags"
          type="text"
          onSelect={() => setSearchText('')}
          onSearch={setSearchText}
          filterOption={(input, option) =>
            option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
          }
          notFoundContent={
            searchText.trim() ? (
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
                Create a tag '{searchText}'
              </Button>
            ) : (
              <Empty
                image={Empty.PRESENTED_IMAGE_SIMPLE}
                description={'No tags available. Type something to create new tag'}
              />
            )
          }
        >
          {tags.map((tag) => (
            <Select.Option key={tag.id} value={tag.id}>
              {tag.name}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
    </div>
  );
}

export default Tags;
