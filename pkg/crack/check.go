package crack

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/projectdiscovery/gologger"
	"net"
	"sync"
	"time"
)

// CheckAlive 存活检测
func (e *Engine) CheckAlive(addrs []*IpAddr) (results []*IpAddr) {
	if len(addrs) == 0 {
		gologger.Info().Msgf("当前目标为空")
		return
	}
	gologger.Info().Msgf("当前时间: %v", time.Now().Format("2006-01-02 15:04:05"))
	gologger.Info().Msgf("存活探测")
	// RunTask
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	taskChan := make(chan *IpAddr, e.threads*2)
	for i := 0; i < e.threads; i++ {
		go func() {
			for task := range taskChan {
				if e.conn(task) {
					mutex.Lock()
					results = append(results, task)
					mutex.Unlock()
				}
				wg.Done()
			}
		}()
	}

	if e.silent {
		for _, task := range addrs {
			wg.Add(1)
			taskChan <- task
		}
		close(taskChan)
		wg.Wait()
	} else {
		bar := pb.StartNew(len(addrs))
		for _, task := range addrs {
			bar.Increment()
			wg.Add(1)
			taskChan <- task
		}
		close(taskChan)
		wg.Wait()
		bar.Finish()
	}

	gologger.Info().Msgf("存活数量: %v", len(results))
	return
}

// conn 建立tcp连接
func (e *Engine) conn(ipAddr *IpAddr) (alive bool) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ipAddr.Ip, ipAddr.Port), time.Duration(e.timeout)*time.Second)
	if err == nil {
		alive = true
	}
	return
}
