// tasks文件夹下为所有定时任务，每个任务一个文件，由main文件统一管理
package tasks

import "github.com/bamzi/jobrunner"

func RunTasks() {
	// optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Start()
	jobrunner.Schedule("@every 20s", AddRate{})
	jobrunner.Schedule("@every 60s", UpdateRate{})
}
