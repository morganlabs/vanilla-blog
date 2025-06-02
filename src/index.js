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
    database.getPost(postId, (err, post) => {
      if (err)
        return server.returnError(
          res,
          404,
          `Post with ID ${postId} does not exist.`,
        );
      return server.returnJSON(res, 200, post);
    });
    return;
  }

  database.getPosts((err, posts) => {
    if (err) return server.returnError(err);
    server.returnJSON(res, 200, posts);
  });
});
