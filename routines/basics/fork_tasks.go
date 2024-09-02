package basics

import (
	"fmt"
	"sync"
	"time"
)

type SubTask struct {
	id       string
	duration int
}

type ForkTask struct {
	id         string
	doneCh     chan bool
	subTasks   []SubTask
	subTasksWg *sync.WaitGroup
}

func (t *ForkTask) getSubTasks() {
	var totalSubTasks int

	fmt.Println("Enter total number of subtasks: ")
	fmt.Scanf("%d ", &totalSubTasks)

	t.subTasks = make([]SubTask, totalSubTasks)

	for idx := 0; idx < totalSubTasks; idx++ {
		fmt.Println("Enter subtask id: ")
		fmt.Scanf("%s ", &t.subTasks[idx].id)

		fmt.Println("Enter subtask duration: ")
		fmt.Scanf("%d ", &t.subTasks[idx].duration)
	}

	fmt.Println("Subtasks: ", t.subTasks)
}

func (s *SubTask) processSubTask(wg *sync.WaitGroup, taskId string, start time.Time) {
	defer wg.Done()

	fmt.Println("PROCESSING SUB_TASK: ", s.id)

	dummyCh := make(chan bool)

	for idx := 0; idx < s.duration; idx++ {
		select {
		case <-dummyCh:
			fmt.Println("unreachable code")
			return
		case <-time.After(time.Second * 5):
			fmt.Println("Processing: ", taskId, ":", s.id, "...Current Sec: ", time.Now().Second(), "...Elapsed Time: ", time.Since(start))
		}
	}

	fmt.Println("COMPLETED SUB_TASK: ", s.id)
}

func getForkTasks() []ForkTask {
	fmt.Println("GET_TASKS")

	var totalTasks int

	fmt.Println("Enter total number of task: ")
	fmt.Scanf("%d ", &totalTasks)

	tasks := make([]ForkTask, totalTasks)

	for idx := 0; idx < totalTasks; idx++ {
		fmt.Println("Enter task id: ")
		fmt.Scanf("%s ", &tasks[idx].id)

		tasks[idx].doneCh = make(chan bool)
		tasks[idx].subTasksWg = &sync.WaitGroup{}
		tasks[idx].getSubTasks()
	}

	return tasks
}

func (t *ForkTask) processTask() {
	fmt.Println("PROCESSING TASK: ", t.id)
	start := time.Now()

	t.subTasksWg.Add(len(t.subTasks))

	for _, subTask := range t.subTasks {
		a := subTask
		go a.processSubTask(t.subTasksWg, t.id, start)
	}

	t.subTasksWg.Wait()
	t.doneCh <- true

	fmt.Println("COMPLETED TASK: ", t.id)
}

func ForkTasks() {
	tasks := getForkTasks()

	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(wg *sync.WaitGroup, task ForkTask) {
			defer wg.Done()

			task.processTask()
		}(wg, task)

		<-task.doneCh
	}

	wg.Wait()
	fmt.Println("ALL TASKS COMPLETED")
}
