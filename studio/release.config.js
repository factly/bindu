module.exports = {
  branches: 'master',
  repositoryUrl: 'https://github.com/factly/bindu-studio',
  plugins: [
    '@semantic-release/commit-analyzer',
    '@semantic-release/release-notes-generator',
    '@semantic-release/github',
  ],
};
