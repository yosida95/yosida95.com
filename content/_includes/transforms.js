import { parser } from "posthtml-parser";
import { render } from "posthtml-render";

import site from "../_data/site.js";

function isInternalLink(href) {
  const baseUrl = new URL(site.url);
  const url = new URL(href, baseUrl);
  if (url.hostname === baseUrl.hostname) {
    return true;
  }
  if (site.externalLinks?.prohibited?.some((e) => isSubdomain(url.hostname, e))) {
    throw new Error(`not allowed: ${url.hostname}`);
  }
  return site.externalLinks?.allowed?.some((e) => isSubdomain(url.hostname, e));
}

function isSubdomain(child, parent) {
  return child === parent || child.endsWith(`.${parent}`);
}

function convert(e) {
  switch (typeof e) {
    case "string":
    case "number":
    case "undefined":
      return e;
    case "object": {
      if (Array.isArray(e)) {
        // Node[]
        return e.map(convert);
      }
      // NodeTag
      if (e.tag !== "a" || e.attrs?.href == null) {
        return {
          ...e,
          content: convert(e.content),
        };
      }
      if (
        isInternalLink(e.attrs.href) ||
        e.attrs.target != null ||
        Object.hasOwn(e.attrs, "data-ignore-external")
      ) {
        return {
          ...e,
          attrs: { ...e.attrs, "data-ignore-external": undefined },
          content: convert(e.content),
        };
      }
      return {
        ...e,
        attrs: { ...e.attrs, target: "_blank" },
        content: convert(e.content),
      };
    }
    default:
      throw new Error("unexpected anchor content");
  }
}

export function external(content) {
  const tree = parser(content, {
    recognizeNoValueAttribute: true,
  });
  return render(convert(tree), {
    closingSingleTag: "slash",
  });
}
