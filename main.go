package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"nappm/internal"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

/*
 * TODO:
 * - SKILL.md to download for agents to manage node processes.
 * - Interactive cli tool for humans for reading output from node state.
 * - Show hardware usage of node processes.
 */

type Process struct {
	id   int
	name string
	pid  int
}

func main() {
	internal.InitDB()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	c := exec.CommandContext(ctx, "node", "server.js")
	c.Stderr = os.Stderr
	c.Cancel = func() error {
		return c.Process.Signal(syscall.SIGTERM)
	}

	read, err := c.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Start(); err != nil {
		log.Fatalf("start node: %v", err)
	}

	if _, err := internal.CreateProcess("node server.js", c.Process.Pid); err != nil {
		log.Printf("save process: %v", err)
	}

	// TODO: Add PID tracking, writing buffer to log file. Assign pid,name,time,uuid to sqlite db for later killing or reading state.

	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("read node output: %v", err)
	}

	if err := c.Wait(); err != nil && ctx.Err() == nil {
		log.Printf("node exited: %v", err)
	}
}

// List all running processes

// Read new output from buffer.

// Manage saved process.

// Manage processes for agents (cli tool)

/**
 * Function for whole os procesess.
 */

/** func GetProcessesList() (list *processes, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	procs, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return nil, err
	}

	list = &processes{}

	for _, p := range procs {
		name, err := p.NameWithContext(ctx)
		if err != nil {
			continue
		}
		proces := &proces{name, int(p.Pid)}
		list.processes = append(list.processes, *proces)
		fmt.Println(" Name: ", name)
	}

	return list, nil

}
*/
