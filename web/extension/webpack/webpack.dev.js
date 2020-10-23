const { merge } = require("webpack-merge");
const common = require("./webpack.common.js");
const { NamedModulesPlugin } = require("webpack");

module.exports = merge(common, {
    mode: "development",

    devtool: "source-map",

    watch: true,
    watchOptions: {
        aggregateTimeout: 600,
        ignored: ["node_modules/**"],
    },

    devServer: {
        hot: true,
    },

    plugins: [
        new NamedModulesPlugin(),
    ],
});
