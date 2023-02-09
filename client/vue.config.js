const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/server": {
        target: "http://localhost:8000", //Server port
        // ws: true,
        // changeOrigin: true, If backend API is not running on the localhost, we need to make this flag true
        changeOrigin: false,
        logLevel: "debug",
        pathRewrite: {
          "/server": "",
        },
      },
    },
  },
});
