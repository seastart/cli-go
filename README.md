# cli-go
a simple library to build golang command line (cli / cmd)apps

## concepts
```
./app [command1] [-opt1=1] [-opt2=2] [command2]
```
`app` is the application  
`command` is sub command  
`opt` is options  
one app may have zero or many commands  

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

## TODO
- support i18n