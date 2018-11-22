package main

import (
	"syscall"
	"unsafe"
	"os/exec"
	"bytes"
	"fmt"
)

var(
	modnetapi32 = NewLazySystemDLL("netapi32.dll")
	procNetUserGetInfo = modnetapi32.NewProc("NetUserGetInfo")
)

func NetUserGetInfo(serverName *uint16, userName *uint16, level uint32, buf **byte) (neterr error) {
	r0, _, _ := syscall.Syscall6(procNetUserGetInfo.Addr(), 4, uintptr(unsafe.Pointer(serverName)), uintptr(unsafe.Pointer(userName)), uintptr(level), uintptr(unsafe.Pointer(buf)), 0, 0)
	if r0 != 0 {
		neterr = syscall.Errno(r0)
	}
	return
}

func exec_shell(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

func main() {
	username := "administrator"
	password := "WanNian123"
	command := "net user " + username + " " + password
	fmt.Println(command)
}