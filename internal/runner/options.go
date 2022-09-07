package runner

import (
	"encoding/json"
	"fmt"
	"github.com/niudaii/crack/internal/utils"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
	"strings"
)

type Options struct {
	// input
	Input     string
	InputFile string
	Module    string
	User      string
	Pass      string
	UserFile  string
	PassFile  string
	// config
	Threads  int
	Timeout  int
	Delay    int
	CrackAll bool
	// output
	OutputFile string
	NoColor    bool
	// debug
	Silent bool
	Debug  bool

	Targets  []string
	UserDict []string
	PassDict []string
}

func ParseOptions() *Options {
	options := &Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`Service cracker`)

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&options.Input, "input", "i", "", "crack service input(example: -i '127.0.0.1:3306', -i '127.0.0.1:3307|mysql')"),
		flagSet.StringVarP(&options.InputFile, "input-file", "f", "", "crack services file(example: -f 'xxx.txt')"),
		flagSet.StringVarP(&options.Module, "module", "m", "all", "choose one module to crack(ftp,ssh,wmi,mssql,oracle,mysql,rdp,postgres,redis,memcached,mongodb)"),
		flagSet.StringVar(&options.User, "user", "", "user(example: -user 'admin,root')"),
		flagSet.StringVar(&options.Pass, "pass", "", "pass(example: -pass 'admin,root')"),
		flagSet.StringVar(&options.UserFile, "user-file", "", "user file(example: -user-file 'user.txt')"),
		flagSet.StringVar(&options.PassFile, "pass-file", "", "pass file(example: -pass-file 'pass.txt')"),
	)

	flagSet.CreateGroup("config", "Config",
		flagSet.IntVar(&options.Threads, "threads", 1, "number of threads"),
		flagSet.IntVar(&options.Timeout, "timeout", 10, "timeout in seconds"),
		flagSet.IntVar(&options.Delay, "delay", 0, "delay between requests in seconds (0 to disable)"),
		flagSet.BoolVarP(&options.CrackAll, "crack-all", "", false, "crack all user:pass"),
	)

	flagSet.CreateGroup("output", "Output",
		flagSet.StringVarP(&options.OutputFile, "output", "o", "crack.txt", "output file to write found results"),
		flagSet.BoolVarP(&options.NoColor, "no-color", "nc", false, "disable colors in output"),
	)

	flagSet.CreateGroup("debug", "Debug",
		flagSet.BoolVar(&options.Silent, "silent", false, "show only results in output"),
		flagSet.BoolVar(&options.Debug, "debug", false, "show debug output"),
	)

	if err := flagSet.Parse(); err != nil {
		gologger.Fatal().Msgf("Program exiting: %v", err)
	}

	options.configureOutput()

	showBanner()

	if err := options.validateOptions(); err != nil {
		gologger.Fatal().Msgf("Program exiting: %v", err)
	}

	if err := options.configureOptions(); err != nil {
		gologger.Fatal().Msgf("Program exiting: %v", err)
	}

	return options
}

// configureOutput 配置输出
func (o *Options) configureOutput() {
	if o.NoColor {
		gologger.DefaultLogger.SetFormatter(formatter.NewCLI(true))
	}

	if o.Debug {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	}

	if o.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}

	gologger.DefaultLogger.SetWriter(utils.NewCLI(o.OutputFile))
}

// validateOptions 验证选项
func (o *Options) validateOptions() error {
	if o.Input == "" && o.InputFile == "" {
		return fmt.Errorf("no service input provided")
	}
	if o.Debug && o.Silent {
		return fmt.Errorf("both debug and silent mode specified")
	}
	if o.Delay < 0 {
		return fmt.Errorf("delay can't be negative")
	}

	return nil
}

// configureOptions 配置选项
func (o *Options) configureOptions() error {
	var err error
	if o.Input != "" {
		o.Targets = append(o.Targets, o.Input)
	} else {
		var lines []string
		lines, err = utils.ReadLines(o.InputFile)
		if err != nil {
			return err
		}
		o.Targets = append(o.Targets, lines...)
	}
	if o.User != "" {
		o.UserDict = strings.Split(o.User, ",")
	}
	if o.Pass != "" {
		o.PassDict = strings.Split(o.Pass, ",")
	}
	if o.UserFile != "" {
		if o.UserDict, err = utils.ReadLines(o.UserFile); err != nil {
			return err
		}
	}
	if o.PassFile != "" {
		if o.PassDict, err = utils.ReadLines(o.PassFile); err != nil {
			return err
		}
	}
	// 去重
	o.Targets = utils.RemoveDuplicate(o.Targets)
	o.UserDict = utils.RemoveDuplicate(o.UserDict)
	o.PassDict = utils.RemoveDuplicate(o.PassDict)
	// 打印配置
	opt, _ := json.Marshal(o)
	gologger.Debug().Msgf("当前配置: %v", string(opt))

	return nil
}
