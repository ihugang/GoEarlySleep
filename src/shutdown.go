package main

import (
	"math/rand"
	"time"

	. "github.com/CodyGuo/win"
)

// 服务主体
func doIt() {
	for {
		now := time.Now()
		hour := now.Hour()
		minute := now.Minute()
		if (hour == 21 && minute >= 30) || hour >= 23 || hour <= 6 {
			v := rand.Intn(10)
			if v == minute%10 {
				shutdown()
			}

		}

		if hour < 21 {
			time.Sleep(60 * time.Second)
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}

// 关机
func shutdown() {
	getPrivileges()
	ExitWindowsEx(EWX_SHUTDOWN, 0)
}

// 获取权限
func getPrivileges() {
	var hToken HANDLE
	var tkp TOKEN_PRIVILEGES

	OpenProcessToken(GetCurrentProcess(), TOKEN_ADJUST_PRIVILEGES|TOKEN_QUERY, &hToken)
	LookupPrivilegeValueA(nil, StringToBytePtr(SE_SHUTDOWN_NAME), &tkp.Privileges[0].Luid)
	tkp.PrivilegeCount = 1
	tkp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED
	AdjustTokenPrivileges(hToken, false, &tkp, 0, nil, nil)
}
