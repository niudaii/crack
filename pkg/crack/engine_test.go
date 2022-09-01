package crack

import (
	"testing"
)

func TestCrackAll(t *testing.T) {
	/*
		=== RUN   TestCrackAll
		=== RUN   TestCrackAll/false
		[INF] 开始爆破: 127.0.0.1:3306 mysql
		[INF] success root:123456
		4 / 4 [-----------------------------------------------------------------------------------------------] 100.00% ? p/s
		=== RUN   TestCrackAll/true
		[INF] 开始爆破: 127.0.0.1:3306 mysql
		[INF] success root:123456
		[INF] success test_user:test2022@
		4 / 4 [-----------------------------------------------------------------------------------------------] 100.00% ? p/s
		--- PASS: TestCrackAll (0.02s)
		    --- PASS: TestCrackAll/false (0.01s)
		    --- PASS: TestCrackAll/true (0.01s)
		PASS
		ok  	crack/pkg/crack	0.036s
	*/
	tests := map[string]*Engine{
		"false": {
			threads:  2,
			timeout:  10,
			crackAll: false,
		},
		"true": {
			threads:  2,
			timeout:  10,
			crackAll: true,
		},
	}
	addrs := []*IpAddr{
		{
			Ip:       "127.0.0.1",
			Port:     3306,
			Protocol: "mysql",
		},
	}
	userDict := []string{"root", "test_user"}
	passDict := []string{"123456", "test2022@"}
	for name, engine := range tests {
		t.Run(name, func(t *testing.T) {
			engine.Run(addrs, userDict, passDict)
		})
	}
}