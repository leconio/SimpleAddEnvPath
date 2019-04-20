# pathgo
Go语言实现，快速添加环境变量工具，适用于Mac和Linux。
可以用来添加PATH变量和GOPATH变量

## 下载方式
```
go get github.com/leconio/pathgo
```

## 使用方式
```
pathgo go 添加当前目录到gopath

pathgo go -s [script path]添加当前目录到script path脚本中gopath

pathgo go -s [script path] -p [path]添加path到script path脚本中gopath

pathgo sys 和上面相同
```

`script path` 为脚本地址，例如zshrc bashrc

`path` 为要添加的路径
