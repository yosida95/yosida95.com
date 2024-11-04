import { isProduction } from "./_includes/env.js";

export default {
  description:
    "ここで述べる内容は私個人の見解に基づくものであり、私の雇用者や私が所属する団体とは一切の関係がありません。",
  date: isProduction ? "git Last Modified" : undefined,
  thumbnail: "/_static/icon_200x200.png",
  // ogpDefaultImage,
  // ogpImage: ogpDefaultImage,
};
