const path = require('path')
//启用热更新
const webpack = require('webpack')
//插件的基本作用就是生成html文件
var htmlWebpackPlugin = require('html-webpack-plugin');

const { VueLoaderPlugin } = require('vue-loader')

module.exports = {
    entry:path.join(__dirname,'./src/main.js'),
    output:{
        path:path.join(__dirname,'./dist'),
        filename:'bundle.js'
    },
    mode: 'production',
    plugins: [ // 添加plugins节点配置插件
        new htmlWebpackPlugin({
            template: path.resolve(__dirname, 'src/index.html'),//模板路径
            filename: 'index.html'//自动生成的HTML文件的名称
        }),
        new webpack.HotModuleReplacementPlugin(),//new一个热更新
        new VueLoaderPlugin()

    ],
    devServer: {
        open: true,
        port: 3600,
        contentBase: 'src',
        hot: true
    },
    //这个节点用于配置所有的第三方模块加载器
    module: {
        rules: [
            { test: /\.css$/, use: ['style-loader', 'css-loader'] },
            { test: /\.less$/, use: ['style-loader', 'css-loader', 'less-loader'] },
            { test: /\.scss$/, use: ['style-loader', 'css-loader', 'sass-loader'] },
            // limit 给定的值，是图片的大小，单位是 byte， 如果我们引用的 图片，
            // 大于或等于给定的 limit值，则不会被转为base64格式的字符串， 如果 图片小于给定的 limit 值，则会被转为 base64的字符串
            { test: /\.(jpg|png|gif|bmp|jpeg)$/, use: 'url-loader?limit=7631&name=[hash:8]-[name].[ext]' }, // 处理 图片路径的 loader
            { test: /\.(ttf|eot|svg|woff|woff2)$/, use: 'url-loader' }, // 处理 字体文件的 loader 
            { test: /\.js$/, use: 'babel-loader', exclude: /node_modules/ }, // 配置 Babel 来转换高级的ES语法
            { test: /\.vue$/, use: 'vue-loader' } 
           
        ]
    },
    //将展示一条警告,通知你这是体积大的资源。
    performance: {
        hints: false
    }
} 