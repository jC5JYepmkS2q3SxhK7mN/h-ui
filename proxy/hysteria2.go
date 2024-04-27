package proxy

import (
	"h-ui/util"
	"os/exec"
	"sync"
)

var mutexHysteria2 sync.Mutex
var cmdHysteria2 exec.Cmd

type hysteria2Process struct {
	Process
	port       string
	binPath    string
	configPath string
}

func NewHysteria2Instance(port string, binPath string, configPath string) *hysteria2Process {
	return &hysteria2Process{Process{mutex: &mutexHysteria2, cmd: &cmdHysteria2}, port, binPath, configPath}
}

func (h *hysteria2Process) StartHysteria2() error {
	if err := h.Start(h.binPath, "-c", h.configPath, "server"); err != nil {
		_ = util.RemoveFile(h.configPath)
		return err
	}
	return nil
}

func (h *hysteria2Process) StopHysteria2() error {
	if err := h.Stop(); err != nil {
		return err
	}
	_ = util.RemoveFile(h.configPath)
	return nil
}

func (h *hysteria2Process) RestartHysteria2() error {
	if err := h.Stop(); err != nil {
		return err
	}
	if err := h.StartHysteria2(); err != nil {
		return err
	}
	return nil
}
