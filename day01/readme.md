# GO 基础之环境搭建 day01

### mac 版本

- 下载编译器
- 配置 path
    ```go
    # vim ~/.bash_profile
    # export PATH= go bin包的位置:$PATH
    # source ~/.bash_profile 
    ```
-  其他配置
   - 创建任意目录存放go的所有代码
        ```go
        /user/shdeng/GoProjects/  # 此目录下创建 bin,pkg,src 三个文件夹
        
        - bin
        - pkg
        - src  # 存放编写所有go代码和依赖
             - crm   # 项目
                  - app.go
             - luffcity  # 项目
                  - app.go
        ```
    - 配置 go的其他环境变量
        ```go
        export GOROOT=/usr/local/go  # 编译器安装目录
        export GOPATH=/Users/shdeng/GoProjects  # 放置 go 代码的相关目录, 指的 '上面' GoProjects 路径
        export GOBIN=/Users/shdeng/GoProjects/bin # go 编译文件的目录 , 指的 '上面' bin 的文件路径，非go/bin/go路径
        ```
- 编写go代码
    ```tree
     GoProjects/
  ├── bin
  ├── pkg
  └── src
  ├── GoRoad
  │   └── day01
  │       └── readme.md
  └── crm
  └── app.go
  ```
- 执行代码
    ```go
    # 方式一 编译+运行
      1。进入 项目目录
      2。执行 ： go run app.go
  
    # 方式二 先编译 ， 后运行 。 生成一个可执行文件，在运行
      1。 go build -o xx # 执行编译 xx为任意名称
      2。 ./xx  # 运行
  
    # 方式三 go install 
      1. 在 项目目录下 执行 ： go install  
      2. 自动在  /user/shdeng/GoProjects/bin 目录下生成一个 可执行程序
  
    # 方式四 代码打包
       go install 生成的包文件 放置到 pkg 墓库
  
    ```


### windows版本
`https://github.com/WuPeiqi/go_course/blob/master/day01%20%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA/day01%20%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA.md`