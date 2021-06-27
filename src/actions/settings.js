import { TOGGLE_SIDER, COLLAPSE_SIDER } from '../constants/settings';

export const toggleSider = () => {
  return (dispatch) => {
    dispatch({
      type: TOGGLE_SIDER,
    });
  };
};

export const collapseSider = () => {
  return (dispatch) => {
    dispatch({
      type: COLLAPSE_SIDER,
    });
  };
};
