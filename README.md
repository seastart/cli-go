# cli-go
a simple library to build golang command line (cli / cmd)apps
```
go get github.com/seastart/cli-go
```

## concepts
```
./app [-main_opt1=1] [command] [-cmd_opt1=1] [-cmd_opt2=2] [subcommands/args]
```
`app` is the application  
`command` is sub command  
`opt` is options
one app may have some main options   
one app may have some commands  
one command may have some command options  
one command main have some subcommands(arguments)   

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
- [build app with sub commands](./examples/commands/main.go)
```
./main
./main test -start=2
./main list -page=3
```
- [build app with sub command with subcommands](./examples/commandcommands/main.go)
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

## TODO
- support i18n