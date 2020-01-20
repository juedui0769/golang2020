
<time.geekbang.org>

极客时间上的golang视频课程

-----

### concurrent_map

- `concurrent_map`是老师在讲课时提到的一个线程安全的字段。github地址如下：
- <https://github.com/easierway/concurrent_map>
- 可以使用`go get`安装，使用`go get`时要去除`https://`只使用`github.com/easierway/concurrent_map`否则会报错。
- 另，下载的源码会放在`GOPATH`的第一个路径中。

```
go get github.com/easierway/concurrent_map
```

```
C:\Users\Wxg>echo %GOPATH%
D:\golang\lib\;D:\wxg104_Go\
```

- 如上，我有设置两个路径，下载的第三方源码就放在`D:\golang\lib\`这个目录下。

### GoLand

- 目前我的工作目录在`D:\wxg104_Go\`, 从github下载的第三方包在`D:\golang\lib\`下
- 然后，导入报错，但是重启GoLand可以解决。目前没有找到更好的办法。





