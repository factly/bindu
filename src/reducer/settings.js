import defaultSettings from '../config/proSettings';
import { TOGGLE_SIDER, COLLAPSE_SIDER } from '../constants/settings';

const initialState = {
  ...defaultSettings,
  sider: {
    collapsed: true,
  },
};

export default function settings(state = initialState, action = {}) {
  switch (action.type) {
    case TOGGLE_SIDER:
      return {
        ...state,
        ...{
          sider: {
            collapsed: !state.sider.collapsed,
          },
        },
      };
    case COLLAPSE_SIDER:
      return {
        ...state,
        ...{
          sider: {
            collapsed: true,
          },
        },
      };
    default:
      return state;
  }
}
