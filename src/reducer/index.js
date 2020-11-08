import { combineReducers } from 'redux';
import templates from './templates.js';
import organisation from './organisation.js';
import categories from './categories';
import tags from './tags';
import charts from './charts';

export default combineReducers({
  templates,
  organisation,
  categories,
  tags,
  charts,
});
