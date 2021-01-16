module.exports = {
  extends: ['airbnb-typescript', 'prettier'],
  parserOptions: {
    project: './tsconfig.json',
  },
  settings: {
    react: {
      pragma: 'h',
      version: '17',
    },
  },
  rules: {
    '@typescript-eslint/comma-dangle': ['error', {
      'arrays': 'always-multiline',
      'exports': 'always-multiline',
      'functions': 'never',
      'imports': 'always-multiline',
      'objects': 'always-multiline',
    }],
    'react/no-unknown-property': ['error', { ignore: ['class'] }],
  },
};
