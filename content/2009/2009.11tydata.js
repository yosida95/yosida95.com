import { isProduction } from "../_includes/env.js";

export default {
  date: isProduction ? "git Created" : undefined,
  layout: "layouts/article.njk",
  tags: ["posts"],
};
