const path = require('path');
const htmlWebpackPlugin = require('html-webpack-plugin');
const miniCssExtractPlugin = require('mini-css-extract-plugin');

const htmlPlugin = new htmlWebpackPlugin({
    template: './static/index.html',
    filename: 'index.html'
})
const cssPlugin = new miniCssExtractPlugin({
    filename: 'style.css'
})

module.exports = {
    mode: 'development',
    entry: './index.tsx',
    output: {
        path: path.resolve(__dirname, '/dist'),
        filename: 'bundle.js'
    },
    devtool: 'inline-source-map',
    module: {
        rules: [
            {
                test: /\.(ts|tsx)$/,
                loader: 'ts-loader'
            },
            {
                test: /\.css$/,
                loader: [
                    miniCssExtractPlugin.loader,
                    'css-loader'
                ]
            }
        ]
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js', '.json']
    },
    devServer: {
        contentBase: path.resolve(__dirname, '/dist'),
        port: 3000,
        open: true
    },
    plugins: [
        htmlPlugin, cssPlugin
    ]
}