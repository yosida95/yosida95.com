import { html } from "common-tags";

export default class Redirect {
  data() {
    return {
      pagination: {
        data: "collections.redirects",
        size: 1,
        alias: "redirect",
      },
      eleventyExcludeFromCollections: true,
      eleventyComputed: {
        title: ({ redirect }) => redirect.title,
        permalink: ({ redirect }) => redirect.from,
      },
    };
  }

  render({ site, redirect, title }) {
    const absoluteUrl = this.htmlBaseUrl(redirect.to, site.url);
    return html`
      <!doctype html>
      <html lang="ja">
        <head>
          <meta charset="utf-8" />
          <link rel="canonical" href="${absoluteUrl}" />
          <meta http-equiv="refresh" content="0; url=${absoluteUrl}" />
          <meta name="robots" content="noindex" />
          <title>${title}</title>
        </head>
        <body>
          <a href="${redirect.to}">Moved Permanently</a>.
        </body>
      </html>
    `;
  }
}
