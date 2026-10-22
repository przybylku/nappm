package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

/*
 * TODO:
 * - SKILL.md to download for agents to manage node processes.
 * - Interactive cli tool for humans for reading output from node state.
 * - Show hardware usage of node processes.
 */

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c := exec.CommandContext(ctx, "node", "server.js")
	read, err := c.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	c.Start()

	// TODO: Add PID tracking, writing buffer to log file. Assign pid,name,time,uuid to sqlite db for later killing or reading state.

	buf := bufio.NewReader(read)
	for {
		line, _, _ := buf.ReadLine()
		fmt.Println(string(line))
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
