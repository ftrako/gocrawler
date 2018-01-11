爬虫
==============

集成各类型数据的爬取，暂时仅支持单机爬虫

数据类型
-----------

ParserType_None             = 0   // none

ParserType_AndroidWandoujia = 1   // 豌豆荚

ParserType_AndroidAnzhi     = 2   // 安智

ParserType_IosAppStore      = 100 // 苹果应用商店

ParserType_IosAnn9          = 101 // ann9爬的ios网站

ParserTypeApe51             = 200 // 51ape网站

ParserType_FileXuexi111     = 201 // 爬文件

ParserType_FileDowncc       = 202

ParserType_FileGdajie       = 203

ParserType_FileJava1234     = 204

ParserType_FilePdfzj        = 205

运行方式
-----------

1、#gocrawler 2 restart 表示重新运行爬类型为2的数据

2、#gocrawler 2 表示继续上一次爬操作

说明
-----------

1、build.sh编译成linux执行文件及所需资源文件，目测在centos7上运行ok

2、ssh登录linux，执行#nohup gocrawler 2 restart & 表示后台运行并支持关闭ssh客户端

3、表创建语句参考bean下的注释语句，比如bean.AppBean