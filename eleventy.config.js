import { HtmlBasePlugin } from "@11ty/eleventy";
import eleventyRss from "@11ty/eleventy-plugin-rss";
import syntaxHighlight from "@11ty/eleventy-plugin-syntaxhighlight";
import htmlnano from "htmlnano";
import strftime from "strftime";

import * as collections from "./content/_includes/collections.js";
import * as shortCodes from "./content/_includes/shortCodes.js";
import * as transforms from "./content/_includes/transforms.js";
import { isProduction } from "./content/_includes/env.js";
import setupMarkdownIt from "./content/_includes/setupMarkdownIt.js";

/** @param {import("@11ty/eleventy").UserConfig} eleventyConfig */
export default (eleventyConfig) => {
  eleventyConfig.amendLibrary("md", (mdLib) => setupMarkdownIt(mdLib));

  eleventyConfig.addGlobalData("permalink", () => {
    return (data) => `${data.page.filePathStem}.${data.page.outputFileExtension}`;
  });

  eleventyConfig.addPlugin(HtmlBasePlugin);
  eleventyConfig.addPlugin(eleventyRss);
  eleventyConfig.addPlugin(syntaxHighlight, {
    errorOnInvalidLanguage: true,
  });

  eleventyConfig.setLayoutResolution(false);
  eleventyConfig.setNunjucksEnvironmentOptions({
    throwOnUndefined: true,
  });

  eleventyConfig.addPassthroughCopy({
    "content/_static/authorized_keys": "_static/authorized_keys",
    "content/_static/favicon.ico": "favicon.ico",
    "content/_static/icon_200x200.png": "_static/icon_200x200.png",
    "content/_static/robots.txt": "robots.txt",
    "content/_static/uBlacklist.txt": "_static/uBlacklist.txt",
    "content/_static/css/font": "_static/css/font",
    "content/_static/img": "_static/img",
    "content/_static/js": "_static/js",
  });

  eleventyConfig.addFilter("categoryUrl", function (name) {
    const slug = eleventyConfig.getFilter("slugify").call(this, name);
    return `/categories/${slug}/`;
  });
  eleventyConfig.addFilter("firstN", (collection, count) => collection.slice(0, count));
  eleventyConfig.addFilter("strftime", (data, format) => strftime(format, data));
  eleventyConfig.addFilter("tagUrl", function (name) {
    const slug = eleventyConfig.getFilter("slugify").call(this, name);
    return `/tags/${slug}/`;
  });

  eleventyConfig.addPairedShortcode("footnote", shortCodes.footnote);
  eleventyConfig.addPairedShortcode("fnitem", shortCodes.fnitem);

  eleventyConfig.addCollection("categories", collections.categories);
  eleventyConfig.addCollection("tags", collections.tags);
  eleventyConfig.addCollection("redirects", collections.redirects);
  eleventyConfig.addCollection("byYear", collections.byYear);
  eleventyConfig.addCollection("byMonth", collections.byMonth);
  eleventyConfig.addCollection("byDay", collections.byDay);

  if (isProduction) {
    eleventyConfig.addTransform("html-minifier", async function (content) {
      if (this.page.outputPath.endsWith(".html")) {
        return htmlnano
          .process(
            content,
            {
              removeComments: false,
              minifyCss: false,
              minifyJs: false,
            },
            undefined,
            { closingSingleTag: "slash" },
          )
          .then((result) => result.html);
      }
      return content;
    });
  }

  eleventyConfig.addTransform("external", function (content) {
    if (this.page.outputPath.endsWith(".html")) {
      return transforms.external(content);
    }
    return content;
  });

  return {
    dir: {
      input: "content",
      output: "dist",
    },
    markdownTemplateEngine: "njk",
    htmlTemplateEngine: "njk",
  };
};
