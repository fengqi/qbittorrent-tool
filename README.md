# qbittorrent-tool
qbittorrent 辅助工具


# 怎么使用

根据平台下载对应的 release 文件，复制 example.config.json 为 config.json，修改 qbittorrent 的配置信息以及自定义标签、分类映射。

使用`设置-下载-运行外部程序`或定时任务运行。

# 1. 自动设置标签

根据`tracker`地址给种子设置标签，如 `https://tracker.baidu.com/announce.php` 则设置标签为 `baidu.com`。

每次处理1000条没有标签的种子，处理完自动退出，如果种子有其他的标签，不会则不会被处理。

如果需要自定义标签，需要在 config.json 配置好对应关系，已经设置过标签的种子可在 webui 删除，等待再次处理。

# 2. 自动设置分类

根据种子保存路径设置分类，需要在 `config.json` 配置好对应关系，获取不到对应关系的种子不会被处理。
