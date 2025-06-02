import { existsSync, readFileSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";
import sqlite3 from "sqlite3";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default class Database {
  // data = {};

  dataPath = join(__dirname, "./data.json");
  path = join(__dirname, "./database.db");

  constructor() {
    // const dataJson = readFileSync(this.path);
    // const data = JSON.parse(dataJson);
    // this.data.posts = new Map(Object.entries(data.posts));
    this.initialise();
  }

  initialise() {
    const isNew = !existsSync(this.path);
    this.sqlite = new sqlite3.Database(
      this.path,
      sqlite3.OPEN_READWRITE | sqlite3.OPEN_CREATE,
      (err) => {
        if (err) {
          throw new Error(err);
        }
        if (isNew) {
          console.log("Created new SQLite3 Database");
        }
        console.log("Connected to SQLite3 Database");
      },
    );

    if (!isNew) return;

    this.sqlite.serialize(() => {
      this.sqlite.run(`
      CREATE TABLE posts(
        publishedAt text,
        title       text,
        summary     text,
        content     text
      ); 
    `);

      const jsonData = JSON.parse(readFileSync(this.dataPath));

      for (const { publishedAt, title, summary, content } of jsonData.posts) {
        this.sqlite.run(
          `INSERT INTO posts VALUES ("${publishedAt}", "${title}", "${summary}", "${content}");`,
        );
      }

      // this.sqlite.each(
      //   "SELECT rowid AS id, title FROM posts",
      //   (err, { id, title }) => {
      //     console.log(`${id}: ${title}`);
      //   },
      // );
    });
  }

  getPosts(callback) {
    this.sqlite.all(
      "SELECT rowid as id, publishedAt, title, summary FROM posts",
      (err, posts) => {
        if (err) return callback(err);
        return callback(null, posts);
      },
    );
  }

  getPost(id, callback) {
    this.sqlite.get(`SELECT * FROM posts WHERE rowid = ${id}`, (err, post) => {
      if (err) return callback(err);
      return callback(null, post);
    });
  }
}
