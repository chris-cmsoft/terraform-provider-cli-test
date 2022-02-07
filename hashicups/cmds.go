package hashicups

import (
	"fmt"
	"os/exec"
)

func RunCommandInDirWithEnvReturnOutput(name string, dir string, env map[string]string, params ...string) ([]byte, error) {
	cmd := exec.Command(name, params...)
	if dir != "." {
		cmd.Dir = dir
	}
	for key, value := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}
	return cmd.Output()
}

func RunCommandWithEnvReturnOutput(path string, env map[string]string, args ...string) ([]byte, error) {
	cmd := exec.Command(path, args...)
	for key, value := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}
	return cmd.Output()
}