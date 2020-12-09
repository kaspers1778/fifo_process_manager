# fifo_process_manager
My student ID is IP8102,so by formula ```2 mod 16 + 1``` i got 3rd variant which is **FIFO**.

FIFO is the most simple strategic planning of processes and fields in that the resource is transferred to the process that has gone before all the others. If the process has been consumed by the ready-made processes, the process control block will come to the tail of the heart. The middle hour for a FIFO strategy is to often finish the great and keep in order the proper processes in the middle of ready-made processes.

In this realisation of Process Manager were used 2 structures:
```
type Process struct {
	Name string
	Time int
}
```
Process contains of 2 fields:
* Name - is the name of process,usualy it just to indetify process(P1,P2...) but also we use name to finish processing(Off);
* Time - is time duration what takes this process;
```
type ProcessManager struct {
	JobQueue []Process
	WaitTime int
}
```
ProcessManager contains of 2 fields:
* JobQueue  - slice of Process objects to queue new ones and dequeue in ***FIFO*** order;
* WaitTime - its time this ProcessManager is processing tasks;
