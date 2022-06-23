# cli-go
[中文帮助](README_CN.md)
a simple library to build golang command line (cli / cmd)apps
```
go get github.com/seastart/cli-go
```

## concepts
```
./app [-main_opt1=1] [command] [-cmd_opt1=1] [-cmd_opt2=2] [subcommands/args]
```
`app` is the application  
`command` is sub command, no prefix `-`  
`opt` is options, start with `-`  
one app may have some main options, such as environment config, like main_opt1 of above example  
one app may have some commands, like command of above example  
one command may have some command options, like cmd_opt1 cmd_opt2 of above example  
one command may have some subcommands(arguments), like subcommands of above example  

## steps (3 step)
- `app := cli.NewCliApp`
- `app.AddCommand`
- `app.Run`

## examples
- [build app with no sub command](./examples/nocommand/main.go)
```
./main
./main -start=2
```
- [build app with sub commands and options](./examples/commands/main.go)
```
./main
./main test -start=2
./main list -page=3
```
- [build app with sub command and sub commands](./examples/commandcommands/main.go)
```
./main
./main test live
./main test -start=2 live
```
- [build app with main options and sub commands](./examples/combine/main.go)
```
// set default env and then run list command
./main -env=prod list -page=3
```

## default help command
each app have a default help command
```
./main help
```

## TODO
- support i18n