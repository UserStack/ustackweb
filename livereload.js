var livereload = require('livereload'),
    server = livereload.createServer({
    originalPath: "http://domain.com",
    debug: true,
    exts: [
      "html",
      "css",
      "js",
      "png",
      "gif",
      "jpg",
      "php",
      "php5",
      "py",
      "rb",
      "erb",
      "coffee",
      "go"
    ]
});
server.watch(__dirname);
