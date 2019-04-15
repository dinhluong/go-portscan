package main

import (
	"fmt"
	"strconv"
	"context"
	"os/exec"
	"net"
	"time"
	"sync"
	"strings"
	"golang.org/x/sync/semaphore"
)
// PortScanners is struct of scanner, lock is number of limit goroutine
type PortScanners struc{
	ip string
	// The lock will act as a threshold that will limit the number of go routines that will be running at any given time
	lock *semaphore.Weighted
}

// Ulimit does check max threads in OS...
func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	
	s := strings.TrimSpace(string(out))
	
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	
	return i
}

func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}

func main() {
	ps := &PortScanner{
		ip:   "127.0.0.1",
		lock: semaphore.NewWeighted(Ulimit()),
	}
	ps.Start(1, 65535, 500*time.Millisecond)
}