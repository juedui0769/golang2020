package ch06

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func pipe1Demo() {
	useBufferIo := false
	fmt.Println("Run command `echo -n \"My first command from golang.\"`: ")
	cmd0 := exec.Command("echo", "-n", "My first command from golang.")
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}
	if !useBufferIo {
		var outputBuf0 bytes.Buffer
		for {
			tempOutput := make([]byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempOutput[:n])
			}
		}
		fmt.Printf("%s\n", outputBuf0.String())
	} else {
		outputBuf0 := bufio.NewReader(stdout0)
		output0, _, err := outputBuf0.ReadLine()
		if err != nil {
			fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
			return
		}
		fmt.Printf("%s\n", string(output0))
	}
}
