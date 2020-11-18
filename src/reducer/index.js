import { combineReducers } from 'redux';
import templates from './templates.js';
import organisation from './organisation.js';
import categories from './categories';
import tags from './tags';

export default combineReducers({
  templates,
  organisation,
  categories,
  tags,
});
