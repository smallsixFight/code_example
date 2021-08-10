## 前言
本文为 docker 小白的记录，大佬请路过。

## 背景

前段时间为一个开发好的项目编写 docker compose，以 ubuntu 作为运行环境，编写一个独立的容器，为了在执行 `docker-compose up -d` 后直接运行容器内的程序，我直接添加了一条如下的 `command` ：

```docker
command: ./program
```

然后过了几天，我重新登录服务器的时候，发现 disk 的使用率高达 87%，立马觉得不对劲，赶紧排查，在排查掉程序自身的 log 的问题后，然后就定位到 docker 自带的日志记录上，查看了日志的内容，发现 ORM 框架自带的控制台标准输出被保存到了日志文件中，再就是因为用了定时任务，导致日志增长得特别快，再就是服务器没有对 docker 的日志大小进行限制

基于上述的四个因素导致的问题出现。在解决问题后，决定作下记录。

## 问题复现

下面会通过简单的 go 代码和 Dockerfile 来复现问题。

### 简单的打印程序

下面用 go 写了一个简单的在进行控制台进行标准输出的死循环程序。通过 `go build -o example main.go` 编译生成二进制文件。

```go
func main() {
	os.Create(filepath.Join(filepath.Dir(os.Args[0]), "example_exec_successful.txt"))
	for {
		fmt.Println("test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test_test")
		time.Sleep(time.Millisecond)
	}
}
```

### Dockerfile

```docker
FROM ubuntu:21.04

ENV HOME="/root"

WORKDIR /root/

COPY example /root/

CMD ./example
```

### 生成镜像和运行

```bash
docker build -t example .
docker run -d --name example example
```

### 查看日志

通过 `docker logs -f --tail=100 CONTAINER_ID` 即可看到程序的标准输出的内容被保存到日志中。

## 解决

我大致总结了下面的三种处理解决方法：

- 程序角度：处理好程序的控制台输出，例如我的项目中是由于 GORM 框架在查询不到记录时会对相关的 SQL 进行控制台标准打印输出，我就自定义了 GORM 的日志配置，不去进行控制台输出打印。
- Dockerfile/docker-compose 角度：不要使用 `CMD` 及 `command` 在进行 `docker run` 的时候直接启动程序，除非能保证程序不会有大量无用的控制台标准输出。
- Docker 方面：通过对 docker log 进行配置，限制 log 文件的最大空间，让 docker 自行处理 log，log 的配置在后面给出。

我认为，对 docker log 进行配置限制大小是一定要的，然后在上述的第一点或者第二点进行选择了。

### docker log 配置

通过 `vim /etc/docker/daemon.json` 进行 log 配置，由于该配置文件一开始是不存在的，所以没有的话也不要过多考虑，新建进行配置就好。log 的简单限制配置如下，下面的限制单个容器中的每个日志文件最大为 `1MB`，最多能有 `5` 个日志文件，即每个容器的日志文件总共最大可使用空间为 `5MB`。

```json
{
"log-driver":"json-file",
"log-opts": { "max-size": "1m", "max-file": "5" }
}
```

---

## 示例代码和程序

[example in Github]([https://github.com/smallsixFight/code_example/tree/main/docker_log_example](https://github.com/smallsixFight/code_example/tree/main/docker_log_example))