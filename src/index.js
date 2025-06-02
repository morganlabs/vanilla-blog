import { readFileSync } from "fs";
import APIServer from "./server.js";
import { parse } from "url";

const configJson = readFileSync("config.json");
const { PORT, HOSTNAME } = JSON.parse(configJson);

const server = new APIServer();
server.start(HOSTNAME, PORT);

server.GET("/api", (req, res) => {
  server.returnJSON(res, 200, { alive: true });
});

server.GET("/api/posts", (req, res, database) => {
  const reqUrl = parse(req.url, true);
  const postId = reqUrl.query.id;

  if (postId) {
    const post = database.getPost(postId);
    if (!post) {
      return server.returnError(
        res,
        404,
        `Post with ID ${postId} does not exist.`,
      );
    }

    return server.returnJSON(res, 200, post);
  }

  const posts = Object.fromEntries(database.getPosts());
  server.returnJSON(res, 200, posts);
});
