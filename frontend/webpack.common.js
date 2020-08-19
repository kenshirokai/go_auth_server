const path = require('path');
const htmlWebpackPlugin = require('html-webpack-plugin');
const miniCssExtractPlugin = require('mini-css-extract-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const environment = process.env.NODE_ENV || 'development';

const htmlPlugin = new htmlWebpackPlugin({
    template: './static/index.html',
    filename: 'index.html',
})
const cssPlugin = new miniCssExtractPlugin({
    filename: 'style.css'
})
const extractTextplugin = new ExtractTextPlugin({
    filename: '[name].bundle.js',
    allChunks: true
})

module.exports = {
    entry: './index.tsx',
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
    plugins: [
        htmlPlugin, cssPlugin, extractTextplugin
    ]
}