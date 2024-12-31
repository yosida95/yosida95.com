import { isProduction } from "../_includes/env.js";

export default {
  name: "yosida95",
  url: "https://yosida95.com/",
  googleAnalytics: isProduction ? "G-LDLSZKMGS4" : undefined,
  externalLinks: {
    allowed: ["blogmedia.yosida95.com"],
    prohibited: ["localhost", "localdomain", "hq.yosida95.net"],
  },
};
