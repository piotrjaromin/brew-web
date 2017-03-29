const webpack = require('webpack');
const path = require('path');

const BUILD_DIR = path.resolve(__dirname, 'public/dist');
const APP_DIR = path.resolve(__dirname, 'web');

const config = {
    devtool: 'cheap-module-source-map',
    entry: {
        app: APP_DIR + '/app.js',
        vendor: ["react", "react-dom"]
    },
    output: {
        path: BUILD_DIR,
        filename: 'app.js'
    },
    module: {
        loaders: [
            {test: /\.js?/, include: APP_DIR, loader: 'babel-loader', query: {presets:['es2015','react']}},
            {test: /\.css$/, loader: 'style-loader!css-loader'},
            {test: /\.eot(\?v=\d+\.\d+\.\d+)?$/, loader: "file-loader?name=fonts/[name].[ext]"},
            {test: /\.scss$/, loaders: ["style-loader", "css-loader?sourceMap", "sass-loader?sourceMap"]},
            {test: /\.(woff|woff2)$/, loader: "url-loader?name=fonts/[name].[ext]&prefix=font/&limit=5000"},
            {
                test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
                loader: "url-loader?name=fonts/[name].[ext]&limit=10000&mimetype=application/octet-stream"
            },
            {test: /\.svg(\?v=\d+\.\d+\.\d+)?$/, loader: "url-loader?name=fonts/[name].[ext]&limit=10000&mimetype=image/svg+xml"}
        ]
    },
    plugins: [
        new webpack.optimize.CommonsChunkPlugin({names: "vendor", filename: "vendor.js"}),
        new webpack.optimize.UglifyJsPlugin({
            compress: {warnings: false},
            comments: false,
            sourceMap: true,
            mangle: true,
            minimize: false
        }),
        new webpack.DefinePlugin({
            'process.env': {
                'NODE_ENV': JSON.stringify('production')
            }
        }),
        new webpack.optimize.AggressiveMergingPlugin()
    ]
};

module.exports = config;