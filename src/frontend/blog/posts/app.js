const urlParams = new URLSearchParams(window.location.search);
const id = urlParams.get("id");

const date = document.getElementById("date");
const title = document.getElementById("title");
const summary = document.getElementById("summary");
const content = document.getElementById("content");

fetch(`http://127.0.0.1:1234/api/posts?id=${id}`).then(async (postRes) => {
  const post = await postRes.json();
  const dateObj = new Date(post.publishedAt);
  const dateStr = formatDate(dateObj);

  date.innerText = dateStr;
  title.innerText = post.title;
  summary.innerText = post.summary;
  content.innerText = post.content;
});

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
