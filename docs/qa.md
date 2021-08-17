## 问题
### 安装 goctl 时报错
报错内容如下：

```
github.com/tal-tech/go-zero/tools/goctl imports
        github.com/tal-tech/go-zero/tools/goctl/model/sql/command imports
        github.com/tal-tech/go-zero/tools/goctl/model/sql/gen imports
        github.com/tal-tech/go-zero/tools/goctl/model/sql/parser imports
        github.com/zeromicro/ddl-parser/parser imports
        github.com/antlr/antlr4/runtime/Go/antlr: ambiguous import: found package github.com/antlr/antlr4/runtime/Go/antlr in multiple modules:
        github.com/antlr/antlr4 v0.0.0-20210105212045-464bcbc32de2 (C:\Users\suhanyujie\go\pkg\mod\github.com\antlr\antlr4@v0.0.0-20210105212045-464bcbc32de2\runtime\Go\antlr)
        github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20210803070921-b358b509191a (C:\Users\suhanyujie\go\pkg\mod\github.com\antlr\antlr4\runtime\!go\antlr@v0.0.0-20210803070921-b358b509191a)
```

在网上查了下，大致是因为多个地方引入了包 github.com/antlr/antlr4/runtime/Go/antlr，单版本不一致造成的。手动不好解决，还是直接下载 goctl 源码，本地编译即可。

```shell
git clone https://github.com/tal-tech/go-zero.git
cd go-zero/tools/goctl
# GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o goctl.exe goctl.go
GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=0818'" -o goctl.exe goctl.go
```

`main.BuildTime=0818` 中的 0818 是自己手填的一个版本标识。然后，将生成的 goctl.exe 文件拷贝到你的 GOROOT 中 `cp goctl.exe $GOROOT/bin/`。
