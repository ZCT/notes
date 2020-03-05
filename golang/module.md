他·

相关资料：

​	[go mod proxies 原理](https://roberto.selbach.ca/blog/)

​	[go mod](<https://github.com/gomods/athens>)

export GO111MODULE=on

export GOPROXY="https://go-mod-proxy.xxx.org"



go list 显示某个包有没有安装



```go
replace (
  golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
  golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
  golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

​	

```go
git tag v1.0.0
git push --tags
```





go mod tidy why



如果某个项目想要依赖其他项目的分支代码 ，在这个项目目录下 go get -u  包名@tzc



 go clean -modcache 清理所以已缓存的模块版本数据

go mod graph 查看现有的依赖结构

 go mod download 下载go.mod文件中指明的所有依赖



go get -u all



go get -u code.xxxxx.org/kite/kite



