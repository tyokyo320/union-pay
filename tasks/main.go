// tasks文件夹下为所有定时任务，每个任务一个文件，由main文件统一管理
package tasks

import (
	"encoding/json"
	"fmt"
	"time"
	"union-pay/global"
	"union-pay/repository"

	"github.com/bamzi/jobrunner"
)

func RunTasks() {
	// optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Start()
	jobrunner.Schedule("TZ=Asia/Tokyo */1 * * * *", boloRate{})
	// At minute o past every hour
	jobrunner.Schedule("TZ=Asia/Tokyo 0 * * * *", AddRate{})
	// At 23:00
	jobrunner.Schedule("TZ=Asia/Tokyo 0 23 * * *", UpdateRate{})

	// tasks := jobrunner.Entries()
	// for _, v := range tasks {
	// 	fmt.Println(v.Job)
	// }
}

type boloRate struct {
	// filtered
}

func (e boloRate) Run() {
	fmt.Println("run bolo")
	j, err := json.Marshal(map[string]interface{}{
		"rate": time.Now().String(),
		"date": "2020-01-18",
	})
	if err != nil {
		global.ErrorLogger.Println("[tasks add]Json marshal went wrong")
		return
	}

	// 并添加至缓存中
	var redisRepo *repository.RateCacheRepository = repository.NewRateCacheRepository(global.REDIS)
	redisRepo.Create("latest", string(j))
}
