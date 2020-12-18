import Home from '../pages/home/index.js';
import Templates from '../pages/templates/index.js';

// Categories
import Categories from '../pages/categories/index.js';
import CreateCategory from '../pages/categories/CreateCategory';
import EditCategory from '../pages/categories/EditCategory';

// Charts
import Charts from '../pages/charts/saved.js';
import EditChart from '../pages/charts/EditCharts.js';
import CreateChart from '../pages/charts/CreateChart.js';

// Spaces
import Spaces from '../pages/spaces/index.js';
import CreateSpace from '../pages/spaces/CreateSpace';
import EditSpace from '../pages/spaces/EditSpace';

// Tag
import Tags from '../pages/tags/index.js';
import CreateTag from '../pages/tags/CreateTag';
import EditTag from '../pages/tags/EditTag';

const routes = {
  templates: {
    path: '/templates',
    Component: Templates,
  },
  home: {
    path: '/',
    Component: Home,
  },
  charts: {
    path: '/charts',
    Component: Charts,
    title: 'Charts',
    permission: {
      resource: 'charts',
      action: 'get',
    },
  },
  createChart: {
    path: '/charts/create',
    Component: CreateChart,
    title: 'Create Chart',
    permission: {
      resource: 'charts',
      action: 'create',
    },
  },
  editChart: {
    path: '/charts/:id/edit',
    Component: EditChart,
    title: 'Edit Chart',
    permission: {
      resource: 'charts',
      action: 'update',
    },
  },
  categories: {
    path: '/categories',
    Component: Categories,
    title: 'Categories',
    permission: {
      resource: 'categories',
      action: 'get',
    },
  },
  createCategory: {
    path: '/categories/create',
    Component: CreateCategory,
    title: 'Create Category',
    permission: {
      resource: 'categories',
      action: 'create',
    },
  },
  editCategory: {
    path: '/categories/:id/edit',
    Component: EditCategory,
    title: 'Edit Category',
    permission: {
      resource: 'categories',
      action: 'update',
    },
  },
  spaces: {
    path: '/spaces',
    Component: Spaces,
    title: 'Spaces',
  },
  createSpace: {
    path: '/spaces/create',
    Component: CreateSpace,
    title: 'Create Space',
    permission: {
      resource: 'spaces',
      action: 'create',
      isSpace: true,
    },
  },
  editSpace: {
    path: '/spaces/:id/edit',
    Component: EditSpace,
    title: 'Edit Space',
    permission: {
      resource: 'spaces',
      action: 'update',
    },
  },
  tags: {
    path: '/tags',
    Component: Tags,
    title: 'Tags',
    permission: {
      resource: 'tags',
      action: 'get',
    },
  },
  createTag: {
    path: '/tags/create',
    Component: CreateTag,
    title: 'Create Tag',
    permission: {
      resource: 'tags',
      action: 'create',
    },
  },
  editTag: {
    path: '/tags/:id/edit',
    Component: EditTag,
    title: 'Edit Tag',
    permission: {
      resource: 'tags',
      action: 'update',
    },
  },
};

export default routes;
