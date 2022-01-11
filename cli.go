package cli

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/color"
)

// app command -opt1=1 -opt2=2
// 原理，使用flag包解析
// xx 这是command，可以通过Arg抓取
// -xx 这是option，可以flag parse抓取
// 注意 -1会被认为是option，如果强制认为-1是arg，用 -- -1

// app
type CliApp struct {
	Desc     string              // 描述
	commands map[string]*Command // 子命令
}

// command
type Command struct {
	Name    string             // 命令名
	Usage   string             // 命令说明
	options map[string]*Option // 子参数
	fs      *flag.FlagSet
	Handler func(subcmds []string, options map[string]*Option) // command的响应
}

// option
type Option struct {
	Name  string      // 参数名
	Dft   interface{} // 参数默认值
	Usage string      // 参数说明
	val   interface{} // 参数解析出的值
}

// 获取解析出的值
func (opt *Option) GetVal() reflect.Value {
	return reflect.ValueOf(opt.val).Elem()
}

// 实例化一个描述为desc的cli app
func NewCliApp(desc string) *CliApp {
	app := &CliApp{
		Desc:     desc,
		commands: map[string]*Command{},
	}
	app.AddCommand("help", "help [subcommand] 查看子命令帮助或全部帮助", func(subcmds []string, options map[string]*Option) {
		if len(subcmds) > 0 {
			for _, subcmd := range subcmds {
				if cmd := app.commands[subcmd]; cmd == nil {
					fmt.Fprintf(os.Stderr, "%s子命令不存在\n", subcmd)
				} else {
					cmd.fs.Usage()
				}
			}
		} else {
			app.showAllHelpAndExit()
		}
	})
	return app
}

// 创建一个command，可传递name为空表示整个程序就是一个command
func (app *CliApp) AddCommand(name string, usage string, handler func(subcmds []string, options map[string]*Option), opts ...*Option) *Command {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	fs.Usage = func() {
		// 第一行command的usage，剩下的是options
		color.New(color.FgGreen).Fprintf(os.Stderr, "%s:\t", name)
		fmt.Fprintln(os.Stderr, usage)
		fs.PrintDefaults()
		fmt.Fprintln(os.Stderr, "")
	}
	options := map[string]*Option{}
	for _, opt := range opts {
		options[opt.Name] = opt
		opt.val = opt.Dft
		switch val := opt.val.(type) {
		case bool:
			opt.val = &val
			fs.BoolVar(&val, opt.Name, val, opt.Usage)
		case int:
			opt.val = &val
			fs.IntVar(&val, opt.Name, val, opt.Usage)
		case string:
			opt.val = &val
			fs.StringVar(&val, opt.Name, val, opt.Usage)
		default:
			panic("目前仅支持bool int string三种option类型")
		}
	}
	cmd := &Command{
		Name:    name,
		Usage:   usage,
		options: options,
		fs:      fs,
		Handler: handler,
	}
	app.commands[name] = cmd
	return cmd
}

// 获取command
func (app *CliApp) GetCommand(name string) *Command {
	return app.commands[name]
}

// 解析并运行
func (app *CliApp) Run() {
	// 不用默认的parse，自己用os args自动区分有没有command的情况
	// flag.Parse()
	args := os.Args[1:]
	// 第一个是command
	// 如果没有指定命令，有默认命令展示默认命令，否则显示帮助并退出
	// 如果没有注册对应命令，展示错误并退出
	mainCommand := ""
	// 没有command或者第一个arg不是参数
	if len(args) > 0 && strings.IndexRune(args[0], '-') != 0 {
		mainCommand = args[0]
		// 剩余的命令行参数
		args = args[1:]
	}
	cmd := app.commands[mainCommand]
	if cmd == nil {
		if mainCommand != "" {
			fmt.Fprintf(os.Stderr, "未注册%s命令\n", mainCommand)
		}
		app.showAllHelpAndExit()
	}
	// // 解析
	// err := cmd.fs.Parse(args)
	// if err == flag.ErrHelp {
	// 	cmd.fs.Usage()
	// } else if err != nil {
	// 	app.Failf("错误: %s", err.Error())
	// 	cmd.fs.Usage()
	// } else {
	// 	// 运行响应
	// 	cmd.Handler(cmd.fs.Args())
	// }
	// flag设置了exitonerror，会自动打印
	cmd.fs.Parse(args)
	// 运行响应
	cmd.Handler(cmd.fs.Args(), cmd.options)
}

// 展示全部帮助并退出
func (app *CliApp) showAllHelpAndExit() {
	color.New().Fprintln(os.Stderr, app.Desc+"\n")
	var keys []string
	for k := range app.commands {
		keys = append(keys, k)
	}
	// 默认排序区分了大小写
	// sort.Strings(keys)
	sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
	// 打印子command的help
	for _, key := range keys {
		app.commands[key].fs.Usage()
	}
	os.Exit(2)
}

// 展示错误并退出
func (app *CliApp) Failf(format string, info ...interface{}) {
	color.New(color.FgRed).Fprintf(os.Stderr, format+"\n", info...)
	os.Exit(3)
}
