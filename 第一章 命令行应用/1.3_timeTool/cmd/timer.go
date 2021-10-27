package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"timeTool/internal/timer"
)

var (
	NowTimeCmd = cobra.Command{
		Use:   "now",
		Short: "获取当前时间",
		Long:  "获取当前时间",
		Run: func(cmd *cobra.Command, args []string) {
			nowTime := timer.GetNowTime()
			// 第一个为格式化输出，第二个为时间戳---其值为自 UTC 1970 年 1 月 1 日起经过的秒数
			log.Printf("输出结果: %s, %d", nowTime.Format("2006-1-2 15:04:05"), nowTime.Unix())
		},
	}
	TimerCmd = cobra.Command{
		Use:   "time",
		Short: "时间格式处理",
		Long:  "时间格式处理",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	startTime        string
	duration         string
	CalculateTimeCmd = cobra.Command{
		Use:   "cal",
		Short: "计算所需时间",
		Long:  "计算所需时间",
		Run: func(cmd *cobra.Command, args []string) {
			var currentTime time.Time
			var timeFormat = "2006-1-2 15:04:05"
			if startTime == "" {
				currentTime = timer.GetNowTime()
			} else {
				var err error
				space := strings.Count(startTime, " ")
				if space == 0 {
					timeFormat = "2006-1-2"
				} else if space == 1 {
					timeFormat = "2006-1-2 15:04:05"
				}
				currentTime, err = time.Parse(timeFormat, startTime)
				if err != nil {
					fmt.Println(err)
					t, _ := strconv.Atoi(startTime)
					currentTime = time.Unix(int64(t), 0)
				}
			}
			resultTime, err := timer.GetCalculateTime(currentTime, duration)
			if err != nil {
				log.Fatalf("Time calculated err: %v", err)
			}

			log.Printf("输出结果: %s, %d", resultTime.Format(timeFormat), resultTime.Unix())

		},
	}
)

func init() {
	TimerCmd.AddCommand(&NowTimeCmd)
	TimerCmd.AddCommand(&CalculateTimeCmd)

	CalculateTimeCmd.Flags().StringVarP(&startTime, "startTime", "s", "", "开始时间")
	CalculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "时间间隔")
}
