package main

import (
	"fmt"
	"time"
)

type JobInfo struct {
	jobId       string        //任务ID
	taskName    string        //任务描述
	taskContent string        //任务内容
	seconds     time.Duration //秒
	NotifyTopic string        //业务的Topic
	taskContext string        // 未知
	times       int           //执行次数  -1 不限制
}

func (jobInfo *JobInfo) Timer() {
	var timerNotifyOb TimerNotify

	timerNotifyOb.taskName = jobInfo.taskName
	timerNotifyOb.taskID = jobInfo.jobId
	timerNotifyOb.taskContext = jobInfo.taskContext
	time.Sleep(jobInfo.seconds)
	fmt.Println("执行任务")
	//向消息管道写数

	numChan <- timerNotifyOb
}
