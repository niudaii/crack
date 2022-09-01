package runner

import (
	"fmt"
	"github.com/niudaii/crack/internal/utils"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
	"strings"
)

type Options struct {
	Input      string
	InputFile  string
	Module     string
	User       string
	Pass       string
	UserFile   string
	PassFile   string
	CrackAll   bool
	Delay      int
	Threads    int
	Timeout    int
	OutputFile string
	NoColor    bool
	Silent     bool
	Debug      bool

	Targets  []string
	UserDict []string
	PassDict []string
}

func ParseOptions() *Options {
	options := &Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`Cracker`)

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&options.Input, "input", "i", "", "crack service input(example: -i '127.0.0.1:3306', -i '127.0.0.1:3307|mysql')"),
		flagSet.StringVarP(&options.InputFile, "input-file", "f", "", "crack service file(example: -f 'xxx.txt')"),
		flagSet.StringVarP(&options.Module, "module", "m", "all", "choose module to crack(ftp,ssh,wmi,mssql,oracle,mysql,rdp,postgres,redis,memcached,mongodb)"),
		flagSet.StringVar(&options.User, "user", "", "user(example: -user 'admin,root')"),
		flagSet.StringVar(&options.Pass, "pass", "", "pass(example: -pass 'admin,root')"),
		flagSet.StringVar(&options.UserFile, "user-file", "", "user file(example: -user-file 'user.txt')"),
		flagSet.StringVar(&options.PassFile, "pass-file", "", "pass file(example: -pass-file 'pass.txt')"),
		flagSet.BoolVarP(&options.CrackAll, "crack-all", "", false, "crack all user:pass"),
	)

	flagSet.CreateGroup("config", "Config",
		flagSet.IntVar(&options.Threads, "threads", 1, "number of threads"),
		flagSet.IntVar(&options.Delay, "delay", 0, "delay between requests in seconds (0 to disable)"),
		flagSet.IntVar(&options.Timeout, "timeout", 10, "timeout in seconds"),
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

func (options *Options) configureOutput() {
	if options.NoColor {
		gologger.DefaultLogger.SetFormatter(formatter.NewCLI(true))
	}

	if options.Debug {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	}

	if options.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}

	gologger.DefaultLogger.SetWriter(utils.NewCLI(options.OutputFile))
}

func (options *Options) validateOptions() error {
	if options.Input == "" && options.InputFile == "" {
		return fmt.Errorf("no service input provided")
	}
	if options.Debug && options.Silent {
		return fmt.Errorf("both debug and silent mode specified")
	}
	if options.Delay < 0 {
		return fmt.Errorf("delay can't be negative")
	}

	return nil
}

func (options *Options) configureOptions() error {
	var err error
	var lines []string
	if options.Input != "" {
		options.Targets = append(options.Targets, options.Input)
	} else {
		lines, err = utils.ReadLines(options.InputFile)
		if err != nil {
			return err
		}
		options.Targets = append(options.Targets, lines...)
	}

	if options.User != "" {
		options.UserDict = strings.Split(options.User, ",")
	}
	if options.Pass != "" {
		options.PassDict = strings.Split(options.Pass, ",")
	}
	if options.UserFile != "" {
		if options.UserDict, err = utils.ReadLines(options.UserFile); err != nil {
			return err
		}
	}
	if options.PassFile != "" {
		if options.PassDict, err = utils.ReadLines(options.PassFile); err != nil {
			return err
		}
	}

	options.Targets = utils.RemoveDuplicate(options.Targets)
	options.UserDict = utils.RemoveDuplicate(options.UserDict)
	options.PassDict = utils.RemoveDuplicate(options.PassDict)
	gologger.Debug().Msgf("%+v", options)

	return nil
}
