# TaoStorage
### This project was developed by Go language, is [TaoAlbum](https://github.com/markusleevip/TaoAlbum-android)'s back-end implementation, the implementation of mobile album storage to private server .
### 本项目是由Go语言开发，是[TaoAlbum](https://github.com/markusleevip/TaoAlbum-android)的后端实现，实现手机相册存储到私有服务器的功能。

## Preconditions 前提条件
###  Install TaoDb 安装TaoDB
项目地址:[TaoDB](https://github.com/markusleevip/taodb)
在output目标已经提供Windows X64平台的可运行版本。
-----------
	go get github.com/markusleevip/taodb
	cd taodb
	./build.sh
	./taodbd -dbPath=/data/taodb -addr=:7398


## Build & Run 编译&运行

-----------
    go get github.com/markusleevip/taostorage
    cd taostorage/main
    go build
    ./main 
		
## Changelog

### Data:7/17/2019
Add the Browse Album feature

    



