# go-blog

react + golang 前后端分离的博客

### 功能

- [x] 文章展示
- [x] 标签
- [ ] 归档
- [ ] 其他页面
- [x] 单用户登录
- [x] 后台管理
- [x] Hexo兼容
- [ ] RSS

### 使用

1. 安装postgresql，并新建一个数据库
2. 修改`config/config.yaml`配置文件
3. 在NGINX或caddy上配置静态文件目录和反代

#### Markdown格式

Markdown

#### hexo文章的导入

在编辑器内粘贴hexo的`_posts`文件内容就能自动导入，后续会添加一键上传的功能

### License

GPL v3

### 感谢

- hexo-theme-apollo: https://github.com/pinggod/hexo-theme-apollo(https://github.com/pinggod/hexo-theme-apollo)
- ant-desin-pro: [https://pro.ant.design/](https://pro.ant.design/)

- gin: 