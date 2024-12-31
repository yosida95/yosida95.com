import js from "@eslint/js";
import prettier from "eslint-config-prettier";
import globals from "globals";

export default [
  {
    ignores: ["output/", "content/**/*.njk"],
  },
  js.configs.recommended,
  {
    languageOptions: {
      ecmaVersion: 2022,
      sourceType: "module",
    },
    rules: {
      eqeqeq: ["error", "smart"],
    },
  },
  {
    files: ["*.config.js", "content/_includes/**/*.js"],
    languageOptions: { globals: globals.node },
  },
  {
    files: ["content/_static/**/*.js"],
    languageOptions: { globals: globals.browser },
  },
  prettier,
];
