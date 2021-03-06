package devops

import (
	"github.com/fsnotify/fsnotify"
	"github.com/wybosys/nnt.logic.go/nnt/config"
	"github.com/wybosys/nnt.logic.go/nnt/core/json"
	"github.com/wybosys/nnt.logic.go/nnt/logger"
	"log"
)

const (
	KEY_PERMISSIONID    = "_permissionid"
	KEY_SKIPPERMISSION  = "_skippermission"
	REDIS_PERMISSIONIDS = 17
)

var (
	current_pid = ""
)

func PID() string {
	return current_pid
}

func flushCurrentPid() {
	cfg, err := json.ReadFile(config.APP_DIR + "/run/permission.cfg")
	if err != nil {
		return
	}

	current_pid, _ = cfg.GetString("id")
}

func permissioncfgWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					//log.Println("modified file:", event.Name)
					flushCurrentPid()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error(err)
			}
		}
	}()

	err = watcher.Add(config.APP_DIR + "/run/permission.cfg")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func init() {
	// 启动刷新permissionid
	permissioncfgWatcher()

	// 刷新最新的PID
	flushCurrentPid()
}
