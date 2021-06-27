import React from 'react';
import TagEditForm from './components/TagForm';
import { useDispatch, useSelector } from 'react-redux';
import { Skeleton } from 'antd';
import { updateTag, getTag } from '../../actions/tags';
import { useHistory } from 'react-router-dom';
import { useParams } from 'react-router-dom';

function EditTag() {
  const history = useHistory();
  const { id } = useParams();

  const dispatch = useDispatch();
  const { tag, loading } = useSelector((state) => {
    return {
      tag: state.tags.details[id] ? state.tags.details[id] : null,
      loading: state.tags.loading,
    };
  });

  React.useEffect(() => {
    dispatch(getTag(id));
  }, [dispatch, id]);

  if (loading) return <Skeleton />;

  const onUpdate = (values) => {
    dispatch(updateTag({ ...tag, ...values }));
    history.push('/tags');
  };
  return <TagEditForm data={tag} onCreate={onUpdate} />;
}

export default EditTag;
