# Hundun
大千世界，起于开天辟地的混沌之中。数据就隐身在这些混沌的背后。
## Require
```
go get -u github.com/gocolly/colly/...

``` 
[gocolly官网](http://go-colly.org/docs) 


#### 爬取腾讯招聘官网数据
```
go run main.go
```

![hr](./img/hr-crawl.png "hr")

#### 爬取深圳房源信息
[深圳房源概况](./sz-house/summary.go "深圳房源")
```
cd sz-house
go run summary.go
```

#### 爬取 在线小说
[史莱克七怪](https://www.xshuyaya.com/read/5/52735.html), 该作品已经改编中国动漫，原著作者 唐家三少。 [斗罗大陆爬虫脚本](./xshuyaya/main.go)。

```
cd xshuyaya
go run main.go 
```

或者编译 Unix 系统下编译运行
```
go build -o xshuyaya 
./xshuyaya
```