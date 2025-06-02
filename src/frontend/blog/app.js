const blogPostsContainer = document.getElementById("blog-posts");

fetch("http://127.0.0.1:1234/api/posts").then(async (posts) => {
  const blogPosts = await posts.json();
  console.log(blogPosts);

  for (const blogPost of blogPosts) {
    makeBlogPost(blogPost);
  }
});

function makeBlogPost(post) {
  const newPost = document.createElement("article");
  const date = formatDate(new Date(post.publishedAt));
  newPost.classList = ["blog-post"];
  newPost.innerHTML = `
     <a class="button" href="./posts?id=${post.id}">
       <span class="date">${date}</span>
       <h2>${post.title}</h2>
       <p class="summary">${post.summary}</p>
     </a>
    `;

  blogPostsContainer.appendChild(newPost);
}

function formatDate(dateObj) {
  const months = {
    0: "January",
    1: "February",
    2: "March",
    3: "April",
    4: "May",
    5: "June",
    6: "July",
    7: "August",
    8: "September",
    9: "October",
    10: "November",
    11: "December",
  };

  const date = dateObj.getDate();
  const dateSuffix =
    date === 1 ? "st" : date === "2" ? "nd" : date === "3" ? "rd" : "th";
  const month = months[dateObj.getMonth()];
  const year = dateObj.getFullYear();
  const hour = dateObj.getHours();
  const minute = dateObj.getMinutes();
  return `${date}${dateSuffix} ${month}, ${year} at ${hour}:${minute}`;
}
