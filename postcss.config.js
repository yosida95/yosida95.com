import autoprefixer from "autoprefixer";
import cssnano from "cssnano";
import postcssImport from "postcss-import";

const isProduction = process.env.NODE_ENV === "production";

/** @type {import('postcss-load-config').Config} */
export default {
  map: !isProduction,
  plugins: [postcssImport(), autoprefixer, cssnano({ preset: "default" })],
};
