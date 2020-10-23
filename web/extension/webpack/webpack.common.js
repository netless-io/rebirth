const paths = require("./paths");
const nodeExternals = require('webpack-node-externals');
const { NoEmitOnErrorsPlugin, NamedModulesPlugin } = require("webpack");
const ForkTsCheckerWebpackPlugin = require("fork-ts-checker-webpack-plugin");

module.exports = {
    entry: {
        "background": paths.backgroundFile,
        "content_script": paths.contentFile,
        "injected": paths.injectedFile,
    },
    target: "web",

    node: {
        __filename: true,
        __dirname: true
    },

    module: {
        rules: [
            {
                test: /\.ts?$/,
                use: [
                    {
                        loader: "babel-loader",
                    },
                    {
                        loader: "eslint-loader",
                        options: {
                            fix: true,
                        },
                    },
                ],
                exclude: /node_modules/,
            },
        ],
    },

    plugins: [
        new NamedModulesPlugin(),
        new NoEmitOnErrorsPlugin(),
        new ForkTsCheckerWebpackPlugin({
            typescript: {
                configFile: paths.tsConfig,
                diagnosticOptions: {
                    semantic: true,
                    syntactic: true,
                    declaration: true,
                },
            },
        }),
    ],

    externals: [nodeExternals()],

    resolve: {
        extensions: [".ts", ".js"],
        alias: {
            "@": paths.appSrc,
        },
    },
    output: {
        path: paths.dist,
    },
};
