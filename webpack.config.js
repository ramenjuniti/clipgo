const path = require("path");

module.exports = {
  mode: "production",
  target: "node",
  entry: {
    contents: path.join(__dirname, "src", "contents", "js", "contents.js"),
    background: [
      path.join(__dirname, "src", "background", "js", "wasm_exec.js"),
      path.join(__dirname, "src", "background", "js", "init_go.js"),
      path.join(__dirname, "src", "background", "js", "background.js")
    ]
  },
  output: {
    path: path.join(__dirname, "build"),
    filename: "[name].bundle.js"
  }
};
