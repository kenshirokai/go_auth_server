const { merge } = require('webpack-merge');
const common = require('./webpack.common');
const path = require('path')

module.exports = merge(common, {
    mode: 'development',
    devtool: 'inline-source-map',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: 'bundle.js',
        chunkFilename: '[name].bundle.js'
    },
    devServer: {
        port: 3000,
        open: true,
        contentBase: path.resolve(__dirname, 'dist')
    }
})