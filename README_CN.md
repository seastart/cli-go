# cli-go
一个用来编写golang命令行(cli / cmd)应用的库
```
go get github.com/seastart/cli-go
```

## 概念
```
./app [-main_opt1=1] [command] [-cmd_opt1=1] [-cmd_opt2=2] [subcommands/args]
```
`app`是程序自身  
`command`是子命令，没有前导`-`  
`opt`是参数配置，有前导`-`  

一个应用可能有一些主参数配置，如env环境配置，如上面的main_opt1  
一个应用可能有一些子命令，如上面的command  
一个子命令可能有一些参数配置，如上面的cmd_opt1 cmd_opt2  
一个子命令可能又有一些孙子命令，如上面的subcommands  

## 步骤 (3 步)
- `app := cli.NewCliApp`
- `app.AddCommand`
- `app.Run`

## 例子
- [编译没有子命令的应用](./examples/nocommand/main.go)
```
./main
./main -start=2
```
- [编写有子命令以及子命令参数配置的应用](./examples/commands/main.go)
```
./main
./main test -start=2
./main list -page=3
```
- [编写有子命令以及孙子命令的应用](./examples/commandcommands/main.go)
```
./main
./main test live
./main test -start=2 live
```
- [编写有主参数配置以及子命令的应用](./examples/combine/main.go)
```
// set default env and then run list command
./main -env=prod list -page=3
```

## 默认的help子命令
每一个应用都自带一个默认的help子命令
```
./main help
```

## TODO
- support i18n