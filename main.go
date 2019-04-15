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

// Ulimit do ...
func Ulimit(){
	
}
func main(){

}