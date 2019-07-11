# gracehttp
golang http 服务优雅重启组件

## 版本要求
Require Go 1.8+

## 使用示例
```bash
go build ./demo/main.go && ./main
curl http://localhost:1991/slow?ms=1
curl http://localhost:1991/slow?ms=20000
kill -SIGUSR2 pid
curl http://localhost:1991/slow?ms=1
```

## 特别说明
本项目克隆于[github.com/tabalt/gracehttp](https://github.com/tabalt/gracehttp)，在此对原作者表示感谢。

## License
[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0.html)