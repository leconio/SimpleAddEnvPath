# SimpleAddEnvPath
Go语言实现，快速添加环境变量工具，适用于Mac和Linux

# 使用方式
从release中下载mac版本的可执行文件，或者自行编译。

```
addPath go 添加当前目录到gopath

addPath go -s [script path]添加当前目录到script path脚本中gopath

addPath go -s [script path] -p [path]添加path到script path脚本中gopath

addPath sys 和上面相同
```

`script path` 为脚本地址，例如zshrc bashrc

`path` 为要添加的路径
