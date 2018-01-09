爬虫
==============

集成各类型数据的爬取，暂时仅支持单机爬虫

数据类型
-----------

  ParserTypeWandoujia = iota // 0 豌豆荚
  
  ParserTypeAppStore         // 苹果应用商店
  
  ParserTypeApe51            // 51ape网站
  
  ParserTypeAnn9             // ann9爬的ios网站
  
  ParserTypeFile             // 爬文件

运行方式
-----------

1、#gocrawler 2 restart 表示重新运行爬类型为2的数据

2、#gocrawler 2 表示继续上一次爬操作

说明
-----------

1、build.sh编译成linux执行文件及所需资源文件，目测在centos7上运行ok

2、ssh登录linux，执行#nohup gocrawler 2 restart & 表示后台运行并支持关闭ssh客户端

3、表创建语句参考bean下的注释语句，比如bean.AppBean