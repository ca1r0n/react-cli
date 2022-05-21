const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require("copy-webpack-plugin");
const ESLintPlugin = require("eslint-webpack-plugin")


module.exports = (env, args) => {
    return {
        entry: './src/index.tsx',
        output: {
            filename: 'bundle.js',
            path: path.resolve(__dirname, 'build'),
        },
        mode: args,
        module: {
            rules: [
                {
                    test: /\.(ts|js)x?$/i,
                    use: {
                        loader: "ts-loader",
                    }
                },
                {
                    test: /\.s[ac]ss$/i,
                    use: [
                        "style-loader",
                        {
                            loader: 'css-loader',
                            options: {
                                sourceMap: true,
                                url: false,
                            }
                        },
                        'postcss-loader',
                        {
                            loader: "sass-loader",
                            options: {
                                additionalData: '@import "variables"; ' +
                                    '@import "mixins";',
                                sassOptions: {
                                    includePaths: ['src/styles/global']
                                },
                                sourceMap: true,
                            }
                        }
                    ],
                },
            ],
        },
        resolve: {
            extensions: ['.tsx', '.ts', '.js'],
        },
        devServer: {
            compress: true,
            port: 8080,
            open: true,
            headers: {
                'Access-Control-Allow-Origin': '*'
            }
        },
        plugins: [
            // for html
            new HtmlWebpackPlugin({
                template: path.join(__dirname, 'src', 'index.html')
            }),
            // for assets
            new CopyPlugin({
                patterns: [
                    {
                        from: "./src/assets",
                        to: "./assets"
                    }
                ]
            }),
            new ESLintPlugin({
                extensions: ["js", "jsx", "ts", "tsx"],
                failOnError: false,
                emitWarning: true,
            }),
        ]
    }
};