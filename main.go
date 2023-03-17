package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Targets []Target `yaml:"targets"`
}

type Target struct {
	Name       string `yaml:"name"`
	Nameserver string `yaml:"nameserver"`
	RecordType string `yaml:"recordtype"`
}

func main() {
	InitTmux()
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <config.yaml>\n", os.Args[0])
		os.Exit(1)
	}

	config := ReadConfig(os.Args[1])
	for _, target := range config.Targets {
		TmuxCommand("split-pane", "watch dig @"+target.Nameserver+" "+target.Name+" "+target.RecordType+"")
		TmuxCommand("select-layout", "tiled")
	}
	TmuxCommand("select-pane", "-t", "0")
	TmuxCommand("kill-pane")
	TmuxCommand("select-layout", "tiled")

	AttachTmux()
}

func ReadConfig(filename string) (config Config) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Can't read file %s: %v ", filename, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse yaml file %s: %v", filename, err)
	}

	return
}

func InitTmux() {
	TmuxCommand("kill-session")
	TmuxCommand("new-session", "-d")
}

func AttachTmux() {
	shell := os.Getenv("SHELL")
	cmd := exec.Command(shell, "-c", "tmux attach -t dns-watcher")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
	TmuxCommand("kill-session")
}

func TmuxCommand(command string, args ...string) {
	fullCommand := append([]string{}, command, "-t", "dns-watcher")
	fullCommand = append(fullCommand, args...)
	fmt.Println("tmux " + strings.Join(fullCommand, " "))
	cmd := exec.Command("tmux", fullCommand...)
	cmd.Run()
}
