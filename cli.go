package cli

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

// app command -opt1=1 -opt2=2
// 原理，使用flag包解析
// xx 这是command，可以通过Arg抓取
// -xx 这是option，可以flag parse抓取
// 注意 -1会被认为是option，如果强制认为-1是arg，用 -- -1

// app
type CliApp struct {
	*Command // the self root command
}

// 成功执行
const CODE_SUCCESS = 0
const CODE_ERROR_COMMON = 100

// command执行
type Handler func(cmd *Command, remaincmds []string) (err error)

// command
type Command struct {
	name    string             // 命令名
	desc    string             // 命令说明
	options map[string]*Option // 子参数
	fs      *flag.FlagSet
	pre     Handler             // before run
	handler Handler             // command handler
	cmds    map[string]*Command // sub commands
	pcmd    *Command            // parent command
	app     *CliApp             // app
}

// option
type Option struct {
	Name string      // 参数名
	Dft  interface{} // 参数默认值，需通过默认值来指定参数类型
	Desc string      // 参数说明
	val  interface{} // 参数解析出的值
}

type Val struct {
	val interface{} // 参数解析出的值
}

func (v Val) String() string {
	if v.val == nil {
		return ""
	}

	return fmt.Sprintf("%v", reflect.ValueOf(v.val).Elem())
}

func (v Val) Int() int {
	if v.val == nil {
		return 0
	}
	return *(v.val.(*int))
}

func (v Val) Bool() bool {
	if v.val == nil {
		return false
	}
	return *(v.val.(*bool))
}

func (v Val) Duration() time.Duration {
	if v.val == nil {
		return 0
	}
	return *(v.val.(*time.Duration))
}

// 实例化一个描述为desc，根参数为opts的cli app
func NewCliApp(desc string, opts ...*Option) *CliApp {
	return NewCliWholeApp(desc, nil, opts...)
}

// 实例化一个描述为desc，根参数为opts，且没有子命令的cli app
func NewCliWholeApp(desc string, handler Handler, opts ...*Option) *CliApp {
	app := &CliApp{
		Command: NewCommand("", desc, handler, opts...),
	}
	app.Command.app = app
	app.AddCommandN("help", "help [subcommand] 查看子命令帮助或全部帮助", func(cmd *Command, remaincmds []string) (err error) {
		// 从rootcmd开始朝招
		tcmd := app.Command
		// 遍历展示子命令的具体帮助
		for _, remaincmd := range remaincmds {
			// help过滤参数
			// ./cli help live -env=qa start
			if remaincmd[:1] == "-" {
				continue
			}
			if subcmd, exists := tcmd.cmds[remaincmd]; !exists {
				app.Warningf("%s不存在子命令%s", tcmd.name, remaincmd)
				break
			} else {
				tcmd = subcmd
			}
		}
		tcmd.ShowHelp()
		return
	})
	return app
}

// 解析并运行
func (app *CliApp) Run(args ...string) (err error) {
	// 不用默认的parse，自己用os args自动区分有没有command的情况
	if len(args) == 0 {
		args = os.Args[1:]
	}
	// 从rootcmd开始parse
	cmd := app.Command
	err = cmd.run(args...)

	code := CODE_SUCCESS
	if err != nil {
		code = CODE_ERROR_COMMON
	}
	if code != CODE_SUCCESS {
		app.Errorf("执行错误:%v", err)
	} else {
		// app.Successf("执行成功")
	}
	return
}

// 常见一个command
func NewCommand(name string, desc string, handler Handler, opts ...*Option) *Command {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	options := map[string]*Option{}
	for _, opt := range opts {
		options[opt.Name] = opt
		if opt.Dft == nil {
			opt.Dft = ""
		}
		opt.val = opt.Dft
		switch val := opt.val.(type) {
		case bool:
			opt.val = &val
			fs.BoolVar(&val, opt.Name, val, opt.Desc)
		case int:
			opt.val = &val
			fs.IntVar(&val, opt.Name, val, opt.Desc)
		case string:
			opt.val = &val
			fs.StringVar(&val, opt.Name, val, opt.Desc)
		case time.Duration:
			opt.val = &val
			fs.DurationVar(&val, opt.Name, val, opt.Desc)
		default:
			panic("目前仅支持bool int string time.Duration四种option类型")
		}
	}
	subcmd := &Command{
		name:    name,
		desc:    desc,
		options: options,
		fs:      fs,
		handler: handler,
		cmds:    map[string]*Command{},
	}
	// 如果command自身没有响应，加上默认显示帮助的响应
	if handler == nil {
		subcmd.handler = func(cmd *Command, remaincmds []string) (err error) {
			subcmd.ShowHelp()
			return
		}
	}
	fs.Usage = subcmd.fsusage
	return subcmd
}

// 设置执行前执行
func (cmd *Command) SetPreRun(handler Handler) {
	cmd.pre = handler
}

// 添加子command
func (cmd *Command) AddCommand(subcmds ...*Command) {
	for _, subcmd := range subcmds {
		subcmd.pcmd = cmd
		subcmd.app = cmd.app
		cmd.cmds[subcmd.name] = subcmd
	}
}

// 创建并返回一个子command，handler里options为解析后的参数(-开头)
func (cmd *Command) AddCommandN(name string, desc string, handler Handler, opts ...*Option) *Command {
	subcmd := NewCommand(name, desc, handler, opts...)
	cmd.AddCommand(subcmd)
	return subcmd
}

// 重写fs的usage
func (cmd *Command) fsusage() {
	cmd.usage(true)
}

func (cmd *Command) usage(showname bool) {
	// 第一行command的name:desc，剩下的是options
	if showname && cmd.name != "" {
		color.New(color.FgGreen).Fprintf(os.Stderr, "%s: ", cmd.name)
	}
	fmt.Fprintln(os.Stderr, cmd.desc)
	cmd.fs.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
}

// 获取子command
func (cmd *Command) SubCommand(name string) *Command {
	return cmd.cmds[name]
}

// 获取父command
func (cmd *Command) ParentCommand() *Command {
	return cmd.pcmd
}

// 获取app
func (cmd *Command) App() *CliApp {
	return cmd.app
}

// 是否有option
func (cmd *Command) HasOpt(name string) bool {
	_, ok := cmd.options[name]
	return ok
}

// 获取option值
func (cmd *Command) OptVal(name string) (val Val) {
	val, _ = cmd.OptValE(name)
	return
}

// 获取option值
func (cmd *Command) OptValE(name string) (val Val, err error) {
	if opt, ok := cmd.options[name]; ok {
		return Val{val: opt.val}, nil
	} else {
		err = fmt.Errorf("%s不存在参数%s", cmd.name, name)
	}
	return
}

// 运行command
func (cmd *Command) run(args ...string) (err error) {
	err = cmd.fs.Parse(args)
	if err != nil {
		return
	}
	// 先执行自身pre
	if cmd.pre != nil {
		err = cmd.pre(cmd, cmd.fs.Args())
		if err != nil {
			return
		}
	}
	// 有子cmd并且存在，遍历执行；否则执行自身
	if len(cmd.fs.Args()) > 0 {
		subcmdname := cmd.fs.Args()[0]
		subargs := cmd.fs.Args()[1:]
		if subcmd, exists := cmd.cmds[subcmdname]; exists {
			cmd = subcmd
			return cmd.run(subargs...)
		} else {
			// cmd.app.Warningf("%s不存在子命令%s", cmd.name, subcmdname)
		}
	}
	// 执行自身
	return cmd.handler(cmd, cmd.fs.Args())
}

// 展示帮助
func (cmd *Command) ShowHelp() {
	// 不显示自己的name
	cmd.usage(false)
	var keys []string
	for k := range cmd.cmds {
		keys = append(keys, k)
	}
	// 默认排序区分了大小写
	// sort.Strings(keys)
	sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
	// 打印子command的help
	for _, key := range keys {
		cmd.cmds[key].fs.Usage()
	}
}

// 以状态码code(0代表成功)退出并展示信息
func (app *CliApp) Exitf(code int, format string, info ...interface{}) {
	if code == 0 {
		app.Successf(format, info...)
	} else {
		app.Errorf(format, info...)
	}
	os.Exit(code)
}

// 展示错误信息到stderr
func (app *CliApp) Errorf(format string, info ...interface{}) {
	color.New(color.FgRed).Fprintf(os.Stderr, format+"\n", info...)
}

// 展示成功信息到stdout
func (app *CliApp) Successf(format string, info ...interface{}) {
	color.New(color.FgGreen).Fprintf(os.Stdout, format+"\n", info...)
}

// 展示告警信息到stdout
func (app *CliApp) Warningf(format string, info ...interface{}) {
	color.New(color.FgYellow).Fprintf(os.Stdout, format+"\n", info...)
}

// 展示提示信息到stdout
func (app *CliApp) Infof(format string, info ...interface{}) {
	color.New(color.FgWhite).Fprintf(os.Stdout, format+"\n", info...)
}
