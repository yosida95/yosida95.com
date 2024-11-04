import MarkdownIt from "markdown-it";
import markdownItAttrs from "markdown-it-attrs";
import markdownItDeflist from "markdown-it-deflist";

export const markdownIt = setupMarkdownIt(new MarkdownIt());

/**
 * @param {MarkdownIt} markdownIt
 * @returns {MarkdownIt}
 */
export default function setupMarkdownIt(markdownIt) {
  return markdownIt
    .set({ html: true, xhtmlOut: true })
    .use(markdownItAttrs, {
      leftDelimiter: "{",
      rightDelimiter: "}",
    })
    .use(markdownItDeflist);
}
