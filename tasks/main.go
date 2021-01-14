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

	err := jobrunner.Schedule("TZ=Asia/Tokyo */1 * * * *", boloRate{})
	if err != nil {
		panic(err)
	}

	// At minute o past every hour
	err = jobrunner.Schedule("TZ=Asia/Tokyo 0 * * * *", AddRate{})
	if err != nil {
		panic(err)
	}

	// At 23:00
	err = jobrunner.Schedule("TZ=Asia/Tokyo 0 23 * * *", UpdateRate{})
	if err != nil {
		panic(err)
	}

	tasks := jobrunner.Entries()
	for _, v := range tasks {
		fmt.Println(v.Job)
	}
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
