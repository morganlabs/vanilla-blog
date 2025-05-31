const { readFileSync } = require("fs");

const blogPostsJson = readFileSync("blogPosts.json");
const blogPosts = JSON.parse(blogPostsJson);
console.log(blogPosts);
