import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import typescript from 'rollup-plugin-typescript2';

export default {
  input: 'pkg/web/src/index.tsx',
  output: { file: 'pkg/web/public/app.js', format: 'iife'},
  plugins: [commonjs(), resolve(), typescript()],
  treeshake: true,
};
