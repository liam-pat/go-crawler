# crawler by Golang

## 依赖

   1. elastic v7 (https://github.com/olivere/elastic)
   2. golang.org/x (gopm get -g -v golang.org/x/text && gopm get -g -v golang.org/x/net)
   3. Elasticseach (使用 docker 运行，https://hub.docker.com/_/elasticsearch)
    
   > 使用 elastic v7 的时候注意，可能会发生以下的问题 https://github.com/olivere/elastic/issues/1098
   
   > 我的解决方案是将 go upgrade to 1.13 就解决了这个问题，若有人可以明确知道是什么导致这个问题，可以留言哟

## 结构

 * crawler_distributed  `分布式-RPC-实现代码`
 
 * engine `定义是否单进程或者多携程进行爬取`
 
 * fetch `获取网站的 HTML（定义 cookie 等信息）`
 
 * model `定义需要保存到 ELasticSearch 的数据结构`
 
 * persist `定义保存数据的 saver`
 
 * scheduler `定义调度器，如队列等`
 
 * site `爬取每个站点的页面爬取结构`
 
 

 
 