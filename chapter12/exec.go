//启动外部命令和程序
package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	//参数说明:1-要运行的进程,2-选项或者参数,3-含有系统环境基本信息的结构体
	fmt.Printf("call  ls -al\n")
	pid, error := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if error != nil {
		fmt.Printf("Error %v starting process!\n", error)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)

	//$PATH被当作字符串处理了。。。
	fmt.Printf("call echo $PATH\n")
	pid, error = os.StartProcess("/bin/echo", []string{"echo", "$PATH"}, procAttr)
	if error != nil {
		fmt.Printf("Error %v starting process!\n", error)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)

	fmt.Printf("call pwd\n")
	pid, error = os.StartProcess("/bin/pwd", []string{"pwd"}, procAttr)
	if error != nil {
		fmt.Printf("Error %v starting process!\n", error)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)
}
