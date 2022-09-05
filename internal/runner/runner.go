package runner

import (
	"fmt"
	"github.com/niudaii/crack/pkg/crack"
	"github.com/projectdiscovery/gologger"
)

type Runner struct {
	options     *Options
	crackRunner *crack.Runner
}

func NewRunner(options *Options) (*Runner, error) {
	crackRunner, err := crack.NewRunner(options.Threads, options.Timeout, options.Delay, options.CrackAll, options.Silent)
	if err != nil {
		return nil, fmt.Errorf("NewRunner err, %v", err)
	}
	runner := &Runner{
		options:     options,
		crackRunner: crackRunner,
	}
	return runner, nil
}

func (r *Runner) Run() {
	addrs := crack.ParseTargets(r.options.Targets)
	addrs = crack.FilterModule(addrs, r.options.Module)
	addrs = r.crackRunner.CheckAlive(addrs)
	results := r.crackRunner.Run(addrs, r.options.UserDict, r.options.PassDict)
	if len(results) > 0 {
		gologger.Info().Msgf("爆破成功: %v", len(results))
		for _, result := range results {
			gologger.Print().Msgf("%v -> %v %v", result.Protocol, result.Addr, result.UserPass)
		}
	}
}
