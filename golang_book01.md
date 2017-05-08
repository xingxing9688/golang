1.go环境准备
   下载https://golang.org/dl/ 
   添加环境变量
   vim /etc/profile or 用户家目录
   #golang
   export GOROOT=/usr/local/go
   export GOPATH=/usr/local/go/packages
   export PATH=$PATH:$GOROOT/bin 
   
   使用go  env 可以查看变量      
   
  1.2 go文件夹作用
     api 文件夹,存放依照go 版本顺序的API增量列表文件.API包含公开的变量、常量、函数等。 
     bin 文件夹,用于存放主要的标准命令文件，包括go godoc 和gofmt 
     blog 文件夹,用于存放官方博客中的所有文章，这些文章都是Mardown 格式的。 
     doc 文件夹,用于存放标准库的HTML格式的程序文档，我们可以通过godoc命令启动一个web程序展现这些文档。
     lib 文件夹,用于存放一些特殊的库文件
     misc文件夹,用于存放一些辅助类的说明和工具。
     pkg文件夹,用于存放安装go标准库后的所有归档文件。
     src文件夹,用于存放go自身、Go标准工具以及标准库的所有源码文件。
     test文件夹,存放用来测试和验证go本身的相关文件
  1.3 go 环境变量设置 
      GOROOT指是go的根目录/usr/local/go 
      GOPATH go install 安装标准库的目录
  1.4 go 标准命令概述
      build:用于编译指定的代码包或Go语言源码文件。
      clean:用于清除因执行其他go命令而遗留下来的临时目录和文件
      doc:用于显示go语言代码包以及程序实体的文档。
      env:用于打印go语言相关的环境信息
      fix:用于修正指定代码包的源码文件中包含的过时语法和代码调用。
      fmt:用于格式化指定代码包中的go源码文件。实际上，它是通过gofmt命令来实现功能的。 
      generate:用于识别指定代码包中源中的go:generate注释,并执行其携带的任意命令，该命令独立于go语言标准的编译和安装体系，如果你有需要解析的go:generate注释，就单独
      运行它。 

      get:用于下载、编译并安装指定的代码包及依赖包.
      install: 用于编译并安装指定的代码包及其依赖包.
      list:用于显示指定代码包的信息.
      run:用于编译并运行指定的命令源码文件,不想生成可执行文件而直接运行命令源码文件.
      test:用于测试指定的代码包，前提是该代码包中必须存在测试源码文件.




      




































2. 
