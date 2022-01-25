package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1   string
	cpu2   string
	ready1 []string
	ready2 []string
	ready3 []string
	io1    []string
	io2    []string
	io3    []string
	io4    []string
	cpu1Process   string
	cpu2Process   string
	ready1Process []string
	ready2Process []string
	ready3Process []string
	io1Process    []string
	io2Process    []string
	io3Process    []string
	io4Process    []string
	r1 int
	r2 int
	r3 int
)

func VariableData() {
	ready1Process = make([]string, 10)
	ready2Process = make([]string, 10)
	ready3Process = make([]string, 10)
	cpu1 = ""
	cpu2 = ""
	ready1 = make([]string, 10)
	ready2 = make([]string, 10)
	ready3 = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
	cpu1Process = ""
	cpu2Process = ""
	io1Process = make([]string, 10)
	io2Process = make([]string, 10)
	io3Process = make([]string, 10)
	io4Process = make([]string, 10)
	r1 = 0
	r2 = 0
	r3 = 0
}

func showProcess() {
	fmt.Printf("CPU 1  --> %s \n", cpu1)
	fmt.Printf("CPU 2  --> %s \n", cpu2)
	fmt.Printf("Ready 1 --> ")
	for i := range ready1 {
		fmt.Printf("%s ", ready1[i])
	}
	fmt.Printf("\nReady 2 --> ")
	for i := range ready2 {
		fmt.Printf("%s ", ready2[i])
	}
	fmt.Printf("\nReady 3 --> ")
	for i := range ready3 {
		fmt.Printf("%s ", ready3[i])
	}
	fmt.Printf("\nI/O 1 --> ")
	for i := range io1 {
		fmt.Printf("%s", io1[i])
	}
	fmt.Printf("\nI/O 2 --> ")
	for i := range io2 {
		fmt.Printf("%s", io2[i])
	}
	fmt.Printf("\nI/O 3 --> ")
	for i := range io3 {
		fmt.Printf("%s", io3[i])
	}
	fmt.Printf("\nI/O 4 --> ")
	for i := range io4 {
		fmt.Printf("%s", io4[i])
	}
	fmt.Printf("\nr1=%d     r2=%d     r3=%d", r1, r2, r3)
	fmt.Printf("\nCommand --> ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func insertQueue(q []string, data string, qplt []string, dataPiority string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			qplt[i] = dataPiority
			break
		}
	}
}

func newProcess(p string, dataPiority string) {
	if cpu1 == "" {
		cpu1 = p
		cpu1Process = dataPiority
		addPior(dataPiority)
	} else if cpu2 == "" {
		cpu2 = p
		cpu2Process = dataPiority
		addPior(dataPiority)
	} else {
		if dataPiority == "1" {
			insertQueue(ready1, p, ready1Process, dataPiority)
		} else if dataPiority == "2" {
			insertQueue(ready2, p, ready2Process, dataPiority)
		} else if dataPiority == "3" {
			insertQueue(ready3, p, ready3Process, dataPiority)
		}
	}
}
func deleteQueue(q []string, dataPiority []string) (string, string) {
	result := q[0]
	resultp := dataPiority[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
		dataPiority[i-1] = dataPiority[i]
	}
	q[9] = ""
	dataPiority[9] = ""
	return result, resultp
}
func expire(cpuName string) {
	if cpuName == "cpu1" {
		dataPiority := cpu1Process
		CheckExpireCpu1(dataPiority)
	} else if cpuName == "cpu2" {
		dataPiority := cpu2Process
		CheckExpireCpu2(dataPiority)
	}
	newQueue := ""
	newPiority := ""
	if r1 < 3 && ready1[0] != "" {
		newQueue, newPiority = deleteQueue(ready1, ready1Process)
	} else if r2 < 3 && ready2[0] != "" {
		newQueue, newPiority = deleteQueue(ready2, ready2Process)
		if r2 < 2 {
			checkPriority()
		}
	} else if r3 < 3 && ready3[0] != "" {
		newQueue, newPiority = deleteQueue(ready3, ready3Process)
		checkPriority()
	}
	addPior(newPiority)

	if newQueue == "" {
		return
	}

	if cpuName == "cpu1" {
		cpu1 = newQueue
		cpu1Process = newPiority
	} else if cpuName == "cpu2" {
		cpu2 = newQueue
		cpu2Process = newPiority
	}
}
func CheckExpireCpu1(inputPiorCpu1 string){
	if inputPiorCpu1 == "1" {
		insertQueue(ready1, cpu1, ready1Process, cpu1Process)
	} else if inputPiorCpu1 == "2" {
		insertQueue(ready2, cpu1, ready2Process, cpu1Process)
	} else if inputPiorCpu1 == "3" {
		insertQueue(ready3, cpu1, ready3Process, cpu1Process)
	}
}
func CheckExpireCpu2(inputPiorCpu2 string){
	if inputPiorCpu2 == "1" {
		insertQueue(ready1, cpu2, ready1Process, cpu2Process)
	} else if inputPiorCpu2 == "2" {
		insertQueue(ready2, cpu2, ready2Process, cpu2Process)
	} else if inputPiorCpu2 == "3" {
		insertQueue(ready3, cpu2, ready3Process, cpu2Process)
	}
}
func use_ioS(ioName string, cpuName string) {
	switch ioName {
	case "1":
		io_cpu(io1, io1Process, cpuName)
	case "2":
		io_cpu(io2, io2Process, cpuName)
	case "3":
		io_cpu(io3, io3Process, cpuName)
	case "4":
		io_cpu(io4, io4Process, cpuName)
	default:
		return
	}
}

func io_cpu(io []string, iop []string, cpu string) {
	if cpu == "cpu1" {
		insertQueue(io, cpu1, iop, cpu1Process)
		cpu1 = ""
		cpu1Process = ""
	} else if cpu == "cpu2" {
		insertQueue(io, cpu2, iop, cpu2Process)
		cpu2 = ""
		cpu2Process = ""
	}
	expire(cpu)
}
func terminate(cpuName string) {
	if cpuName == "cpu1" {
		if r1 < 3 && ready1[0] != "" {
			cpu1, cpu1Process = deleteQueue(ready1, ready1Process)
		} else if r2 < 3 && ready2[0] != "" {
			cpu1, cpu1Process = deleteQueue(ready2, ready2Process)
		} else if r3 < 3 && ready3[0] != "" {
			cpu1, cpu1Process = deleteQueue(ready3, ready3Process)
		} else if ready1[0] == "" && ready2[0] == "" && ready3[0] == "" {
			cpu1 = ""
			cpu1Process = ""
		}
		checkPriority()
		addPior(cpu1Process)
	} else if cpuName == "cpu2" {
		if r1 < 3 && ready1[0] != "" {
			cpu2, cpu2Process = deleteQueue(ready1, ready1Process)
		} else if r2 < 3 && ready2[0] != "" {
			cpu2, cpu2Process = deleteQueue(ready2, ready2Process)
		} else if r3 < 3 && ready3[0] != "" {
			cpu2, cpu2Process = deleteQueue(ready3, ready3Process)
		} else if ready1[0] == "" && ready2[0] == "" && ready3[0] == "" {
			cpu2 = ""
			cpu2Process = ""
		}
		checkPriority()
		addPior(cpu2Process)
	}
}
func use_ioSx(ioName string) {
	fq := ""
	dataPiority := ""
	switch ioName {
	case "1":
		fq, dataPiority = deleteQueue(io1, io1Process)
	case "2":
		fq, dataPiority = deleteQueue(io2, io2Process)
	case "3":
		fq, dataPiority = deleteQueue(io3, io3Process)
	case "4":
		fq, dataPiority = deleteQueue(io4, io4Process)
	default:
		return
	}
	if fq == "" {
		return
	}

	if cpu1 == "" {
		cpu1 = fq
		cpu1Process = dataPiority
		addPior(dataPiority)
	} else if cpu2 == "" {
		cpu2 = fq
		cpu2Process = dataPiority
		addPior(dataPiority)
	} else {
		if dataPiority == "1" {
			insertQueue(ready1, fq, ready1Process, dataPiority)
		} else if dataPiority == "2" {
			insertQueue(ready2, fq, ready2Process, dataPiority)
		} else if dataPiority == "3" {
			insertQueue(ready3, fq, ready3Process, dataPiority)
		}
	}
}
func addPior(p string) {
	if p == "1" {
		r1++
	} else if p == "2" {
		r2++
	} else if p == "3" {
		r3++
	}
}
func checkPriority() {
	if r1 == 3 {
		r1 = 0
		if r2 == 3 {
			r2 = 0
		}
	} else if r2 == 3 {
		r2 = 0
	} else if r3 == 3 {
		r3 = 0
	}
}
func main() {
	VariableData()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
			//	newProcess(commandx[i])
			if i % 2 != 0 {
				newProcess(commandx[i], commandx[i+1])
			}
				
			}
		case "terminate":
			terminate(commandx[1])
		case "expire":
			expire(commandx[1])
		case "io":
			use_ioS(commandx[1], commandx[2])

		case "iox":
			use_ioSx(commandx[1])

		default:
			fmt.Printf("\nInput Error \n")
		}
	}

}
