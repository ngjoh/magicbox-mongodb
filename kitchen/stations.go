package kitchen

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/viper"
)

type Options struct {
	Dir string
	Env []string
}

func StartSession(kitchen string, stationId string) (string, int, error) {
	root := viper.GetString("KITCHENROOT")
	sessionRoot := path.Join(root, ".sessions")

	dir := path.Join(sessionRoot, kitchen, stationId)
	os.MkdirAll(dir, os.ModePerm)

	pid, err := Spawn("bash", Options{Dir: dir}, "nohup", "pwsh")
	if err != nil {
		return dir, pid, err
	}
	return dir, pid, nil

}

func Spawn(program string, options Options, args ...string) (int, error,
) {
	cmd := exec.Command(program, args...)
	if options.Dir != "" {
		cmd.Dir = options.Dir
	}
	if options.Env != nil {
		cmd.Env = options.Env
	}
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		return 0, errors.New(fmt.Sprint(err))
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		return 0, errors.New(fmt.Sprint(err))
	}
	log.Println("Spawned", cmd.Process.Pid)

	return cmd.Process.Pid, nil
}
