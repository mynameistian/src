package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TimerNotify struct {
	taskID      string
	taskName    string
	taskContext string
}

type CreateTaskRsp struct {
	jobId       string //创建异常为空
	taskContext string // 未知
	times       int    //执行次数  -1 不限制
}
type ResponseInfo struct {
	code        int32  //响应状态码 正常:200 异常:非200
	description string //异常描述信息
}
type TimerDispatcher struct {
	JobInfoList map[string]CreateTaskRsp
}

func randJobId() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("JOBID%d", rand.Intn(100)) //获取
}

//初始化一个类 空间为1024
func NewTimerDispatcher() *TimerDispatcher {
	return &TimerDispatcher{
		JobInfoList: make(map[string]CreateTaskRsp, 1024),
	}
}

//添加一个延时任务
func (timerDispatcher *TimerDispatcher) CreateOnceTask(JobInfo JobInfo) (err ResponseInfo) {
	JobInfo.jobId = randJobId()
	JobInfo.times = 1
	go JobInfo.Timer()
	err.code = 200
	err.description = ""
	return
}

//添加一个定时任务
func (timerDispatcher *TimerDispatcher) CreateContinuedTask(JobInfo JobInfo) (createTaskRsp CreateTaskRsp) {

	JobInfo.jobId = randJobId()
	JobInfo.times = -1
	timerDispatcher.AddJobId(JobInfo)
	go JobInfo.Timer()
	createTaskRsp.jobId = JobInfo.jobId
	return
}

//取消定时任务
func (timerDispatcher *TimerDispatcher) RemoveContinuedTask(createTaskRsp CreateTaskRsp) (responseInfo ResponseInfo) {

	timerDispatcher.DeleteJobId(createTaskRsp.jobId)

	//从任务列表中删除任务
	responseInfo.code = 200
	responseInfo.description = ""
	return
}

//从Map中删除任务信息
func (timerDispatcher *TimerDispatcher) DeleteJobId(jobId string) {
	delete(timerDispatcher.JobInfoList, jobId)
}

//在map中添加任务信息
func (timerDispatcher *TimerDispatcher) AddJobId(jobInfo JobInfo) {
	var createTaskRspOb CreateTaskRsp
	createTaskRspOb.jobId = jobInfo.jobId
	createTaskRspOb.taskContext = jobInfo.taskContext
	createTaskRspOb.times = jobInfo.times
	timerDispatcher.JobInfoList[jobInfo.jobId] = createTaskRspOb
}

//执行者
func (timerDispatcher *TimerDispatcher) Practitioner() {

	//读取消息管道
	timerNotifyOb, ok := <-numChan
	if ok {
		return
	}

	if JobInfo, ok := timerDispatcher.JobInfoList[timerNotifyOb.taskID]; ok {
		return
	} else {
		if JobInfo.times != 1 {
			times := JobInfo.times
			times--
			JobInfo.times = times
			timerDispatcher.JobInfoList[timerNotifyOb.taskID] = JobInfo
		} else {
			timerDispatcher.DeleteJobId(timerNotifyOb.taskID)
			//执行任务函数
			//exec()
			return
		}
	}
	//执行任务函数
	//exec()

}

var numChan chan TimerNotify

func main() {

	var jobInfoOb JobInfo
	numChan = make(chan TimerNotify, 1024)
	timer := NewTimerDispatcher()

	timer.CreateOnceTask(jobInfoOb)
	timer.CreateContinuedTask(jobInfoOb)

	var createTaskRspOb CreateTaskRsp
	createTaskRspOb.jobId = "XXXXXX"
	rs := timer.RemoveContinuedTask(createTaskRspOb)
	fmt.Sprintln(rs)
}
