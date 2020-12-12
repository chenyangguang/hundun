# Hundun
大千世界，起于开天辟地的混沌之中。数据就隐身在这些混沌的背后。
## Require
```
go get -u github.com/gocolly/colly/...

``` 
[gocolly官网](http://go-colly.org/docs)

## go (gocolly) 爬取网上感兴趣的数据

 - 爬取腾讯招聘官网数据
```
go run main.go
```

![hr](./img/hr-crawl.png "hr")

- 爬取深圳房源信息
[深圳房源概况](./sz-house/summary.go "深圳房源")

```
  cd sz-house
  go run summary.go
```
- 爬取 在线小说
《斗罗大陆》全章， 该作品已经改编中国动漫，原著作者 唐家三少

```
cd xshuyaya
go run main.go 
```
或者编译 unix 下编译
```
go build -o xshuyaya 
./xshuyaya
```