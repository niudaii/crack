package runner

import (
	"github.com/niudaii/crack/pkg/crack"
	"github.com/projectdiscovery/gologger"
)

type Runner struct {
	options *Options
	engine  *crack.Engine
}

func NewRunner(options *Options) (*Runner, error) {
	runner := &Runner{
		options: options,
		engine:  crack.NewEngine(options.Threads, options.Timeout, options.Delay, options.CrackAll, options.Silent),
	}
	return runner, nil
}

func (r *Runner) Run() {
	addrs := crack.ParseTargets(r.options.Targets)
	addrs = crack.FilterModule(addrs, r.options.Module)
	addrs = r.engine.CheckAlive(addrs)
	results := r.engine.Run(addrs, r.options.UserDict, r.options.PassDict)
	if len(results) > 0 {
		gologger.Info().Msgf("爆破成功: %v", len(results))
		for _, result := range results {
			gologger.Info().Msgf("%v -> %v %v", result.Protocol, result.Addr, result.UserPass)
		}
	}
}
