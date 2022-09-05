package crack

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/niudaii/crack/internal/utils"
	"github.com/niudaii/crack/pkg/crack/plugins"
	"github.com/projectdiscovery/gologger"
	"strings"
	"sync"
	"time"
)

type Runner struct {
	threads  int
	timeout  int
	delay    int
	crackAll bool
	silent   bool
}

func NewRunner(threads, timeout, delay int, crackAll, silent bool) (*Runner, error) {
	return &Runner{
		threads:  threads,
		timeout:  timeout,
		delay:    delay,
		crackAll: crackAll,
		silent:   silent,
	}, nil
}

type Result struct {
	Addr     string
	Protocol string
	UserPass string
}

type IpAddr struct {
	Ip       string
	Port     int
	Protocol string
}

func (r *Runner) Run(addrs []*IpAddr, userDict []string, passDict []string) (results []*Result) {
	for _, addr := range addrs {
		results = append(results, r.Crack(addr, userDict, passDict)...)
	}
	return
}

func (r *Runner) Crack(addr *IpAddr, userDict []string, passDict []string) (results []*Result) {
	gologger.Info().Msgf("开始爆破: %v:%v %v", addr.Ip, addr.Port, addr.Protocol)

	var tasks []plugins.Service
	var taskHash string
	taskHashMap := map[string]bool{}
	// GenTask
	if len(userDict) == 0 {
		userDict = UserMap[addr.Protocol]
	}
	if len(passDict) == 0 {
		passDict = append(passDict, TemplatePass...)
		passDict = append(passDict, CommonPass...)
	}
	for _, user := range userDict {
		for _, pass := range passDict {
			// 替换{user}
			pass = strings.ReplaceAll(pass, "{user}", user)
			// 任务去重
			taskHash = utils.Md5(fmt.Sprintf("%v%v%v%v%v", addr.Ip, addr.Port, addr.Protocol, user, pass))
			if taskHashMap[taskHash] {
				continue
			}
			taskHashMap[taskHash] = true
			tasks = append(tasks, plugins.Service{
				Ip:       addr.Ip,
				Port:     addr.Port,
				Protocol: addr.Protocol,
				User:     user,
				Pass:     pass,
				Timeout:  r.timeout,
			})
		}
	}
	// RunTask
	stopHashMap := map[string]bool{}
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	taskChan := make(chan plugins.Service, r.threads)
	for i := 0; i < r.threads; i++ {
		go func() {
			for task := range taskChan {
				addrStr := fmt.Sprintf("%v:%v", addr.Ip, addr.Port)
				userPass := fmt.Sprintf("%v:%v", task.User, task.Pass)
				addrHash := utils.Md5(addrStr)
				// 判断是否已经停止爆破
				mutex.Lock()
				if stopHashMap[addrHash] {
					wg.Done()
					mutex.Unlock()
					continue
				}
				mutex.Unlock()
				gologger.Debug().Msgf("[trying] %v", userPass)
				scanFunc := plugins.ScanFuncMap[task.Protocol]
				resp := scanFunc(&task)
				switch resp {
				case plugins.CrackSuccess:
					if !r.crackAll {
						mutex.Lock()
						stopHashMap[addrHash] = true
						mutex.Unlock()
					}
					gologger.Silent().Msgf("%v -> %v %v", addr.Protocol, addrStr, userPass)
					results = append(results, &Result{
						Addr:     addrStr,
						Protocol: addr.Protocol,
						UserPass: userPass,
					})
				case plugins.CrackError:
					mutex.Lock()
					stopHashMap[addrHash] = true
					mutex.Unlock()
				case plugins.CrackFail:
				}
				if r.delay > 0 {
					time.Sleep(time.Duration(r.delay) * time.Second)
				}
				wg.Done()
			}
		}()
	}

	if r.silent {
		for _, task := range tasks {
			wg.Add(1)
			taskChan <- task
		}
	} else {
		bar := pb.StartNew(len(tasks))
		for _, task := range tasks {
			bar.Increment()
			wg.Add(1)
			taskChan <- task
		}
		bar.Finish()
	}
	close(taskChan)
	wg.Wait()

	gologger.Info().Msgf("爆破结束")

	return
}
