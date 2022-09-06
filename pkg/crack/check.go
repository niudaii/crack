package crack

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"net"
	"sync"
	"time"
)

// CheckAlive 存活检测
func (e *Runner) CheckAlive(addrs []*IpAddr) (results []*IpAddr) {
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
	} else {
		bar := pb.StartNew(len(addrs))
		for _, task := range addrs {
			bar.Increment()
			wg.Add(1)
			taskChan <- task
		}
		close(taskChan)
	}
	close(taskChan)
	wg.Wait()

	return
}

// conn 建立tcp连接
func (e *Runner) conn(ipAddr *IpAddr) (alive bool) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ipAddr.Ip, ipAddr.Port), time.Duration(e.timeout)*time.Second)
	if err == nil {
		alive = true
	}
	return
}
