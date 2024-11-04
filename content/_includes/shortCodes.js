import { html } from "common-tags";

export function footnote(content) {
  return html`<div class="footnote">${content}</div> `;
}

export function fnitem(content, id, ...backrefs) {
  return html`<aside id="${id}" role="doc-footnote">
    ${content}${backrefs.map((e) => html`&nbsp;<a href="#${e}">↩</a>`)}
  </aside>`;
}
