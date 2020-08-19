module.exports = {
  branches: 'master',
  repositoryUrl: 'https://github.com/factly/data-portal-admin',
  plugins: [
    '@semantic-release/commit-analyzer',
    '@semantic-release/release-notes-generator',
    '@semantic-release/github',
  ],
};
