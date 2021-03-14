import { combineReducers } from 'redux';
import templates from './templates.js';
import organisation from './organisation.js';
import categories from './categories';
import tags from './tags';
import charts from './charts';
import settings from './settings';
import spaces from './spaces';
import media from './media';
import notifications from './notification';

export default combineReducers({
  spaces,
  templates,
  organisation,
  categories,
  tags,
  charts,
  settings,
  media,
  notifications,
});
