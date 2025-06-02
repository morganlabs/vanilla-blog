import { readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default class Database {
  path = join(__dirname, "./database.json");
  data = {};

  constructor() {
    const dataJson = readFileSync(this.path);
    const data = JSON.parse(dataJson);
    this.data.posts = new Map(Object.entries(data.posts));
  }

  getPosts() {
    const posts = this.data.posts;
    return new Map(
      Array.from(posts, ([id, value]) => {
        const { content, ...rest } = value;
        return [id, rest];
      }),
    );
  }

  getPost(id) {
    const post = this.data.posts.get(id);

    if (!post) return { status: 404 };
    return post;
  }
}
