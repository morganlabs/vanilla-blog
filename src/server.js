import { createServer } from "http";
import { fileURLToPath, parse } from "url";
import Database from "./database/index.js";
import { dirname, extname, join } from "path";
import { lstatSync, readdirSync, readFileSync } from "fs";

const mimeTypes = {
  ".html": "text/html",
  ".css": "text/css",
  ".js": "application/javascript",
  ".png": "image/png",
  ".jpg": "image/jpeg",
  ".jpeg": "image/jpeg",
  ".gif": "image/gif",
};

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default class APIServer {
  constructor() {
    this.apiRoutes = new Map();
    this.database = new Database();

    // Create a server from the NodeJS `http` module
    this.server = createServer((req, res) => {
      // Parse the full URL and queries
      const { pathname, query } = parse(req.url, true);
      const queryParams = new URLSearchParams(query).toString();
      const queryString = queryParams ? `?${queryParams}` : "";

      console.log(`${req.method} ${pathname}`);

      if (pathname.includes("/api/"))
        return this.processAPIRequest(req, res, pathname);
      else if (pathname.endsWith(".html")) {
        const newUrl = pathname.endsWith("index.html")
          ? pathname.replace("index.html", "")
          : pathname.replace(".html", "");
        res.statusCode = 301;
        res.setHeader("Location", `${newUrl}/${queryString}`);
        return res.end();
      } else return this.processRequest(req, res, pathname, queryString);
    });
  }

  processRequest(req, res, pathname, queryString) {
    const localPath = join(__dirname, "frontend", pathname);
    const mimeType = mimeTypes[extname(localPath)];

    try {
      const stats = lstatSync(localPath);

      if (stats.isDirectory()) {
        if (!pathname.endsWith("/")) {
          res.statusCode = 301;
          res.setHeader("Location", `${pathname}/${queryString}`);
          return res.end();
        }

        const indexPath = join(localPath, "index.html");
        const indexStat = lstatSync(indexPath);
        if (!indexStat.isFile()) throw new Error("No Index file");
        const indexFile = readFileSync(indexPath);
        return this.serveFile(res, indexFile, mimeTypes[".html"]);
      }

      const file = readFileSync(localPath);
      return this.serveFile(res, file, mimeType);
    } catch (error) {
      console.error(error);
      return this.returnError(res, 500, error);
    }
  }

  serveFile(res, file, mimeType) {
    res.setHeader("Content-Type", mimeType);
    res.statusCode = 200;
    res.end(file);
  }

  processAPIRequest(req, res, pathname) {
    // Look for a matching existing route
    const callback = this.apiRoutes.get(`${req.method} ${pathname}`);

    // If a match exists, run it's callback function
    if (callback) {
      callback(req, res, this.database);
      return;
    } else {
      this.returnError(res, 404, "This route does not exist");
    }
  }

  returnJSON(res, status, json) {
    // Set type to JSON
    res.setHeader("content-Type", "Application/JSON");

    // Set status code
    res.statusCode = status;

    res.end(JSON.stringify(json));
  }

  returnError(res, status, message) {
    this.returnJSON(res, status, { status, message });
  }

  start(hostname, port) {
    // Start listening on hostname:port
    this.server.listen(port, hostname, () => {
      // Log when succesfully started
      console.log(`Server running on ${hostname}:${port}`);
    });
  }

  registerRoute(method, path, callback) {
    this.apiRoutes.set(`${method} ${path}`, callback);
  }

  GET(path, callback) {
    this.registerRoute("GET", path, callback);
  }
}
