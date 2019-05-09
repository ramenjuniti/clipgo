const path = require("path");

module.exports = {
  mode: "production",
  target: "node",
  entry: {
    contents: path.join(__dirname, "src", "contents", "contents.js"),
    background: [
      path.join(__dirname, "src", "background", "wasm_exec.js"),
      path.join(__dirname, "src", "background", "init_go.js"),
      path.join(__dirname, "src", "background", "background.js")
    ]
  },
  output: {
    path: path.join(__dirname, "build"),
    filename: "[name].bundle.js"
  }
};
