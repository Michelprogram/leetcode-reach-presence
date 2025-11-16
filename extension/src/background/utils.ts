export const setTitle = (title: string, difficulty: string) =>
  `${title.replace(/^\d+\. /, "")} (${difficulty})`;
