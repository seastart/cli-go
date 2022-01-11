# cli-go
a simple library to build golang command line (cli / cmd)apps

## concepts
```
./app [command1] [-opt1=1] [-opt2=2] [command2]
```
`app` is the application  
`command` is subcommand  
`opt` is options  
one app may have zero or many subcommands  

## steps (3 step)
- `app := cli.NewCliApp`
- `app.AddCommand`
- `app.Run`

## examples
- [build app with no subcommand](./examples/nocommand/main.go)
```
./main
./main -start=2
```
- [build app with subcommands](./examples/commands/main.go)
```
./main
./main test -start=2
./main list -page=3
```

## TODO
- support i18n