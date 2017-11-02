package main

import (
	"fmt"
	"github.com/robfig/cron"
	"cron/utils"
	"cron/models"
	"time"
)

func main() {
	spec := "10 * * * * *"
	c := cron.New()
	c.AddFunc(spec, cronDetectorUpdateStatus)
	c.Start()
	select {}
}
func cronDetectorUpdateStatus() {
	fmt.Println("每10分钟监测刷新一次探测器状态")
	mdb, mSession := utils.GetMgoDbSession()
	defer mSession.Close()
	ds, _ := models.GetDetectors(mdb)
	for _, d := range ds {
		if time.Now().Unix() - d.UpdatedAt.Unix() > 60 {
			d.UpdateByStatus(mdb)
		}
	}
}