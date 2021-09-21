const path = require("path")

module.exports = {
  outputDir: path.resolve(
    __dirname,
    process.env.NODE_ENV === "production" ? "../dist" : "./dist"
  )
}
