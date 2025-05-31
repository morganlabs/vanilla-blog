const blogPosts = [
  {
    id: "63V93W",
    publishedAt: new Date("2025-05-01 11:32:08"),
    title: "JavaScript in 2025: What’s New and What’s Next?",
    summary:
      "Explore the latest features added to JavaScript in recent ECMAScript updates, including new syntax improvements, performance boosts, and evolving best practices. This post breaks down how these changes impact web development and what developers should expect in the coming year.",
  },
  {
    id: "6D11SJ",
    publishedAt: new Date("2025-04-14 15:36:06"),
    title: "Rust vs. Go: A Modern Systems Programming Showdown",
    summary:
      "This article compares Rust and Go, highlighting their strengths and weaknesses in systems programming, concurrency, and safety. Readers will learn which language suits different project types and how each is shaping the future of backend and infrastructure development.",
  },
  {
    id: "1HQZNZ",
    publishedAt: new Date("2025-03-17 05:20:02"),
    title: "Mastering GDScript for Godot 4: Essential Tips for Game Developers",
    summary:
      "A practical guide to GDScript’s new features in Godot 4, including improved syntax, performance enhancements, and integration with the engine’s latest tools. The post offers actionable tips for both beginners and experienced developers looking to maximize productivity in game development.",
  },
  {
    id: "PSA20V",
    publishedAt: new Date("2025-04-08 10:19:47"),
    title:
      "From JavaScript to Rust: Migrating a Web App for Performance and Safety",
    summary:
      "Follow a step-by-step journey of migrating a critical web application component from JavaScript to Rust using WebAssembly. The post covers the motivations, challenges, and measurable performance gains, offering advice for teams considering a similar transition.",
  },
  {
    id: "86M1IM",
    publishedAt: new Date("2025-04-16 19:18:58"),
    title:
      "Building High-Performance APIs with Go: Best Practices and Pitfalls",
    summary:
      "Discover how Go’s simplicity and concurrency model make it ideal for building scalable APIs. This article outlines best practices for structuring Go projects, handling errors, and optimizing performance, along with common mistakes to avoid for robust backend development.",
  },
].sort((a, b) => b.publishedAt - a.publishedAt);

const blogPostsContainer = document.getElementById("blog-posts");

function makeBlogPost(post) {
  const newPost = document.createElement("article");
  const date = formatDate(post.publishedAt);
  console.log(date);
  newPost.classList = ["blog-post"];
  newPost.innerHTML = `
     <a class="button" href="./posts/${post.id}.html">
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

for (const blogPost of blogPosts) {
  makeBlogPost(blogPost);
}
