package utils

import (
	"bufio"
	"fmt"
	"os"
)

// will scan input for full sentences
// the normal scan command only scans until it gets to a space
// effectively only scanning one word, we want to scan the whole input
// in case we have commands with arguments
func ScanInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return "-999"
}

// sets the windows console title to a given string
// func SetConsoleTitle(title string) (int, error) {
// 	handle, err := syscall.LoadLibrary("Kernel32.dll")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer syscall.FreeLibrary(handle)
// 	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
// 	if err != nil {
// 		return 0, err
// 	}
// 	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
// 	return int(r), err
// }

func PrintMsg(msg string) {
	fmt.Printf("system -> %v", msg)
}
