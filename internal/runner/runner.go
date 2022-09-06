package runner

import (
	"fmt"
	"github.com/niudaii/crack/pkg/crack"
	"github.com/projectdiscovery/gologger"
	"time"
)

type Runner struct {
	options     *Options
	crackRunner *crack.Runner
}

func NewRunner(options *Options) (*Runner, error) {
	crackOptions := &crack.Options{
		Threads:  options.Threads,
		Timeout:  options.Timeout,
		Delay:    options.Delay,
		CrackAll: options.CrackAll,
		Silent:   options.Silent,
	}
	crackRunner, err := crack.NewRunner(crackOptions)
	if err != nil {
		return nil, fmt.Errorf("crack.NewRunner() err, %v", err)
	}
	return &Runner{
		options:     options,
		crackRunner: crackRunner,
	}, nil
}

func (r *Runner) Run() {
	start := time.Now()
	gologger.Info().Msgf("当前时间: %v", start.Format("2006-01-02 15:04:05"))
	// 解析目标
	addrs := crack.ParseTargets(r.options.Targets)
	addrs = crack.FilterModule(addrs, r.options.Module)
	if len(addrs) == 0 {
		gologger.Info().Msgf("目标为空")
		return
	}
	// 存活探测
	gologger.Info().Msgf("存活探测")
	addrs = r.crackRunner.CheckAlive(addrs)
	gologger.Info().Msgf("存活数量: %v", len(addrs))
	// 服务爆破
	results := r.crackRunner.Run(addrs, r.options.UserDict, r.options.PassDict)
	if len(results) > 0 {
		gologger.Info().Msgf("爆破成功: %v", len(results))
		for _, result := range results {
			gologger.Print().Msgf("%v -> %v %v", result.Protocol, result.Addr, result.UserPass)
		}
	}
	// 运行时间
	gologger.Info().Msgf("运行时间: %v", time.Since(start))
}
