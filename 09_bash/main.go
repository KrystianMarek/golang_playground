package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync"
)

type ControlKey byte
type ControlSequence [2]ControlKey

/*
91|65 - arrow up
91|66 - arrow down
91|68 - arrow left
91|87 - arrow right
*/

func (c ControlKey) equals(b byte) bool {
	return byte(c) == b
}

func (c ControlSequence) equals(b []byte) bool {
	return c[0] == ControlKey(b[0]) && c[0] == ControlKey(b[1])
}

// https://www.physics.udel.edu/~watson/scen103/ascii.html
const (
	ctrlA ControlKey = 1
	ctrlB ControlKey = 2
)

func asciiHelper() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Bytes()
		for b := range msg {
			if ctrlA.equals(msg[b]) {
				fmt.Print("Ctrl-A")
			} else if ctrlB.equals(msg[b]) {
				fmt.Print("Ctrl-B")
			} else {

				fmt.Printf("%s(%d)", string(msg[b]), int(msg[b]))
			}
		}
		fmt.Println("")
	}
}

type Screen struct {
	initCommand []string
	cmdStdin    io.WriteCloser
	cmdLog      []string
	outLog      []string
	Rx          chan string
	Tx          chan string
	Terminal    *Terminal
}

func (s *Screen) run() {
	command := exec.Command("bash")
	//command := exec.Command("docker", "run", "-it", "ubuntu:latest", "bin/bash")
	s.cmdStdin, _ = command.StdinPipe()
	cmdStdout, _ := command.StdoutPipe()

	//stdinScanner := bufio.NewScanner(cmdStdin)
	stdoutScanner := bufio.NewScanner(cmdStdout)
	go func() {
		//log.Println("Start: s.run.func.A")
		stdoutScanner.Scan()
		for stdoutScanner.Scan() {
			s.outLog = append(s.outLog, stdoutScanner.Text())
			s.Tx <- stdoutScanner.Text()
		}
	}()

	go func() {
		//log.Println("Start: s.run.func.B")
		for rx := range s.Rx {
			//log.Printf("s.run.func: %v\n", rx)
			fmt.Fprintln(s.cmdStdin, rx)
		}
	}()

	s.Terminal.wg.Add(1)

	go func() {
		//log.Println("Start: s.run.func.C")
		err := command.Run()
		if err != nil {
			//log.Println(command)
			panic(err)
		}
		command.Wait()
	}()
}

func (s *Screen) close() {
	close(s.Rx)
	close(s.Tx)
}

func NewScreen(t *Terminal) *Screen {
	screen := Screen{
		Terminal: t,
		Rx:       make(chan string, 1),
		Tx:       make(chan string, 1),
	}
	screen.run()

	return &screen
}

type Terminal struct {
	stdin        *os.File
	stdout       *os.File
	Screens      map[int]*Screen
	Prompt       string
	Cmd          string
	ActiveScreen int
	wg           sync.WaitGroup
}

func (t *Terminal) String() string {
	return fmt.Sprintf("(%d)%s", t.ActiveScreen, t.Prompt)
}

func (t *Terminal) updateActiveScreen(ck ControlKey, commandLine []byte) {
	for byte := range commandLine {
		if ck.equals(commandLine[byte]) && len(commandLine) > byte+1 {
			screenNumber, err := strconv.Atoi(string(commandLine[byte+1]))
			if err == nil {
				t.ActiveScreen = screenNumber
				_, ok := t.Screens[t.ActiveScreen]
				if !ok {
					t.Screens[t.ActiveScreen] = NewScreen(t)
					//ToDo: FixME!
					go func() {
						for out := range t.Screens[t.ActiveScreen].Tx {
							fmt.Println(out)
						}
					}()
				}
			}
		}
	}
}

func (s *Screen) runCommand(commandLine string) {
	go func() {
		//log.Println("s.runCommand: " + commandLine)
		s.Rx <- commandLine
	}()
}

func (t *Terminal) Run() {
	go func() {
		stdinScanner := bufio.NewScanner(t.stdin)
		fmt.Print(t)
		for stdinScanner.Scan() {
			t.updateActiveScreen(ctrlA, stdinScanner.Bytes())
			fmt.Print(t)
			t.Screens[t.ActiveScreen].runCommand(stdinScanner.Text())
		}
	}()

	//ToDo: FixME!
	go func() {
		for out := range t.Screens[t.ActiveScreen].Tx {
			fmt.Println(out)
		}
	}()
}

func (t *Terminal) Close() {
	log.Println("Closing...")

	for k, _ := range t.Screens {
		t.Screens[k].close()
	}

	t.wg.Done()

	t.stdin.Close()
	t.stdout.Close()
	log.Println("Closed")
}

func NewTerminal(Stdin *os.File, Stdout *os.File) *Terminal {
	t := Terminal{
		stdin:        Stdin,
		stdout:       Stdout,
		Prompt:       ">",
		Cmd:          "",
		ActiveScreen: 1,
		Screens:      make(map[int]*Screen),
	}

	t.Screens[t.ActiveScreen] = NewScreen(&t)

	t.Run()

	return &t
}

func main() {
	//asciiHelper()
	terminal := NewTerminal(os.Stdin, os.Stdout)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	terminal.Close()
}
