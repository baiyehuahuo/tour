package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"tour/internal/timer"
)

var (
	calculateTime string
	duration      string
	nowTimeCmd    = &cobra.Command{
		Use:   "now",
		Short: "获取当前时间",
		Long:  "获取当前时间",
		Run: func(cmd *cobra.Command, args []string) {
			nowTime := timer.GetNowTime()
			log.Printf("当前时间: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		},
	}
	calculateTimeCmd = &cobra.Command{
		Use:   "calc",
		Short: "计算所需时间",
		Long:  "计算所需时间",
		Run: func(cmd *cobra.Command, args []string) {
			var calcTime time.Time
			var layout = "2006-01-02 15:04:05"
			if calculateTime == "" {
				calcTime = time.Now()
			} else {
				var err error
				if strings.Count(calculateTime, " ") == 0 {
					layout = "2006-01-02"
				}
				calcTime, err = time.Parse(layout, calculateTime)
				if err != nil {
					t, _ := strconv.Atoi(calculateTime)
					calcTime = time.Unix(int64(t), 0)
					layout = "2006-01-02 15:04:05"
				}
			}
			ansTime, err := timer.GetCalculateTime(calcTime, duration)
			if err != nil {
				log.Fatalf("timer.GetCalculateTime: %s", err)
			}
			log.Printf("计算时间: %s, %d", ansTime.Format(layout), ansTime.Unix())
		},
	}
	timeCmd = &cobra.Command{
		Use:   "time",
		Short: "时间格式处理",
		Long:  "时间格式处理",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要被计算的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "增加的时间， 单位是 ns us ms s m h")
}

/*
go run . time now
go run . time calc -c="2001-01-02 22:22:22" -d="1h"
go run . time calc -c="2001-01-02 22:22:22" -d="-1h"
go run . time calc -c="2001-01-02 22:22:22" -d="3h"
go run . time calc -c="2001-01-02" -d="24h"
go run . time calc -c="978566400" -d="2h"
*/
