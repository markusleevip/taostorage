# TaoStorage
### 本项目是由Go语言开发，是[TaoAlbum](https://github.com/markusleevip/TaoAlbum-android)的后端实现，实现手机相册存储到私有服务器的功能。

## 前提条件
### 安装TaoDB
项目地址:[TaoDB](https://github.com/markusleevip/taodb)
在output目标已经提供Windows X64平台的可运行版本。TaoDB的默认端口：7398
TaoStorage的默认端口：8000 ，TaoAlbum连接TaoStorage的端口即8000，如：你的服务器IP是192.168.1.100 ，TaoAlbum连接的完整地址为http://192.168.1.100:8000

-----------
	go get github.com/markusleevip/taodb
	cd taodb
	./build.sh
	./taodbd -dbPath=/data/taodb -addr=:7398


##  编译运行本项目 

-----------
    go get github.com/markusleevip/taostorage
    cd taostorage/main
    go build
    ./main 
    #如果是windows系统 请执行main.exe
		
    
		
## 运行截图 
1.运行TaoDB
<image src="./output/images/taodbd.png" /> 
2.运行TaoStorage
<image src="./output/images/taostorage.png" /> 

## Main方法
<image src="./output/images/main.png" /> 





