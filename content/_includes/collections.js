import strftime from "strftime";

const collator = Intl.Collator("ja", { sensitivity: "base" });

export function categories(collections) {
  /** @type {{name: string, url: string, pages: any[]}[]} */
  const categories = [];
  collections.getFilteredByTag("posts").forEach((e) => {
    const category = e.data.postCategory?.trim();
    if (category == null) {
      throw new Error(`postCategory not set: ${e.inputPath}`);
    }
    const index = categories.find((e) => e.name === category);
    if (index == null) {
      categories.push({
        name: category,
        url: `/categories/${encodeURIComponent(category)}/`,
        pages: [e],
      });
      return;
    }
    index.pages.push(e);
  });
  categories.sort(({ name: a }, { name: b }) => collator.compare(a, b));
  return categories;
}

export function categoryByName(collections) {
  return Object.fromEntries(categories(collections).map((e) => [e.name, e]));
}

export function tags(collections) {
  /** @type {{name: string, url: string, pages: any[]}[]} */
  const tags = [];
  collections.getFilteredByTag("posts").forEach((e) => {
    const tag = e.data.postTags;
    if (!Array.isArray(tag)) {
      if (tag != null) {
        throw new Error(`postTags must be an array: ${e.inputPath}`);
      }
      return;
    }
    tag.forEach((tag) => {
      const index = tags.find((e) => e.name === tag);
      if (index == null) {
        tags.push({
          name: tag,
          url: `/tags/${encodeURIComponent(tag)}/`,
          pages: [e],
        });
        return;
      }
      index.pages.push(e);
    });
  });
  tags.sort(({ name: a }, { name: b }) => collator.compare(a, b));
  return tags;
}

export function tagByName(collections) {
  return Object.fromEntries(tags(collections).map((e) => [e.name, e]));
}

export function byYear(collections) {
  /** @type {Array<{
   *    title: string,
   *    url: string,
   *    year: number,
   *    months: Array<{
   *      title: string,
   *      url: string,
   *      year: number,
   *      month: number,
   *      pages: any[]
   *      days: Array<{
   *        title: string,
   *        url: string,
   *        year: number,
   *        month: number,
   *        day: number,
   *        pages: any[]
   *      }>
   *    }>,
   *    pages: any[]
   *  }>} */
  const out = [];
  collections.getFilteredByTag("posts").forEach((e) => {
    const year = e.page.date.getFullYear();
    let byYear = out.find((e) => e.year === year);
    if (byYear == null) {
      byYear = {
        title: `${year}`,
        url: `/${year}/`,
        year: year,
        months: [],
        pages: [],
      };
      out.push(byYear);
    }
    byYear.pages.push(e);

    const month = e.page.date.getMonth() + 1;
    let byMonth = byYear.months.find((e) => e.month === month);
    if (byMonth == null) {
      byMonth = {
        title: strftime("%b", e.page.date),
        url: `/${year}/${padZero(month)}/`,
        year: year,
        month: month,
        pages: [],
        days: [],
      };
      byYear.months.push(byMonth);
    }
    byMonth.pages.push(e);

    const day = e.page.date.getDate();
    let byDay = byMonth.days.find((e) => e.day === day);
    if (byDay == null) {
      byDay = {
        title: strftime("%b", e.page.date),
        url: `/${year}/${padZero(month)}/${padZero(day)}/`,
        year: year,
        month: month,
        day: day,
        pages: [],
      };
      byMonth.days.push(byDay);
    }
    byDay.pages.push(e);
  });
  return out;
}

function padZero(n, digits = 2) {
  return n.toString().padStart(digits, "0");
}

export function byMonth(collections) {
  return byYear(collections).flatMap((e) => e.months);
}

export function byDay(collections) {
  return byMonth(collections).flatMap((e) => e.days);
}

export function redirects(collections) {
  return collections
    .getAllSorted()
    .filter(({ data: { aliases } }) => Array.isArray(aliases) && aliases.length > 0)
    .flatMap(({ url: to, data: { title, aliases } }) =>
      aliases.map((from) => ({ from, to, title })),
    );
}
