import { combineReducers } from 'redux';
import templates from './templates.js';
import organisation from './organisation.js';

export default combineReducers({
  templates,
  organisation,
});
