// tasks文件夹下为所有定时任务，每个任务一个文件，由main文件统一管理
package tasks

import (
	"github.com/bamzi/jobrunner"
)

func RunTasks() {
	// optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Start()
	// At minute o past every hour
	jobrunner.Schedule("TZ=Asia/Tokyo 0 */1 * * *", AddRate{})
	// At 23:00
	jobrunner.Schedule("TZ=Asia/Tokyo 0 23 * * *", UpdateRate{})

	// tasks := jobrunner.Entries()
	// for _, v := range tasks {
	// 	fmt.Println(v.Job)
	// }
}
