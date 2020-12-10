package main

import (
	"fmt"
	"sync"
	"time"
)

type Process struct {
	Name string
	Time int
}
func (p *Process) Execute(){
	time.Sleep(time.Duration(p.Time) * time.Millisecond)
}
type ProcessManager struct {
	JobQueue []Process
	WaitTime int
}

func (pm *ProcessManager) AddProcess(p Process){
	pm.JobQueue = append(pm.JobQueue, p)
	fmt.Printf("%v with duration %v was added to  work queue\n",p.Name,p.Time)
}

func  (pm *ProcessManager) Start(wg *sync.WaitGroup){
	working := true
	for working{
		if len(pm.JobQueue)>0{
			working = pm.ProcessNextTask()
		}
	}
	wg.Done()
}

func  (pm *ProcessManager) ProcessNextTask() (bool){
	p := pm.JobQueue[0]
	pm.JobQueue = pm.JobQueue[1:]
	p.Execute()
	fmt.Printf("%v is done with duration of %v ms.Wait time is %v\n", p.Name, p.Time,pm.WaitTime)
	pm.WaitTime+= p.Time
	if p.Name == "Off"{
		return false
	}
	return true
}



func main() {
	wg := sync.WaitGroup{}
	pm := ProcessManager{}
	go pm.Start(&wg)
	wg.Add(1)
	for i := 1;i<10;i++{
		pm.AddProcess(Process{
			Name: fmt.Sprintf("P%v",i),
			Time: i*10,
		})
	}
	time.Sleep(time.Millisecond*50)
	for i := 11;i<20;i++{
		pm.AddProcess(Process{
			Name: fmt.Sprintf("P%v",i),
			Time: i*10,
		})
	}
	pm.AddProcess(Process{
		Name: "Off",
		Time: 10,
	})
	wg.Wait()
}


