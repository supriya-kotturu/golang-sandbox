package basics

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id       string
	duration int
	cancelCh chan int
}

type Tasks = []Task

func processTask(wg *sync.WaitGroup, task Task) {
	defer func() {
		wg.Done()
	}()

	start := time.Now()

	for i := 0; i < task.duration; i++ {
		select {
		case <-task.cancelCh:
			fmt.Print("Task: ", task.id, " : CANCELLED\n")
			return
		case <-time.After(5 * time.Second):
			fmt.Print("Processing...", task.id, "...Current: ", time.Now().Second(), "...Elapsed: ", time.Since(start), "\n")
		}
	}

	fmt.Print("Task: ", task.id, " : completed\n")
}

func getTasks() Tasks {
	var totalTasks int

	fmt.Println("Enter total tasks to be generated: ")
	fmt.Scanf("%d ", &totalTasks)

	tasks := make(Tasks, totalTasks)

	for i := 0; i < totalTasks; i++ {
		fmt.Println("Enter task id: ")
		fmt.Scanf("%s ", &tasks[i].id)

		fmt.Println("Enter task duration: ")
		fmt.Scanf("%d ", &tasks[i].duration)

		cancelCh := make(chan int, 1)
		tasks[i].cancelCh = cancelCh
	}

	return tasks

}

func getCancelTaskId(wg *sync.WaitGroup, tasks Tasks) {
	defer wg.Done()

	var cancelTaskId string

	fmt.Println("Enter task id to cancel: ")
	fmt.Scanf("%s ", &cancelTaskId)
	fmt.Println("Cancelling task: ", cancelTaskId)

	for _, task := range tasks {
		if task.id == cancelTaskId {
			task.cancelCh <- 1
			close(task.cancelCh)
			return
		}
	}

}

func GenerateTasks() {
	wg := sync.WaitGroup{}

	tasks := getTasks()

	wg.Add(len(tasks) + 1)
	fmt.Println(tasks)

	for _, task := range tasks {
		go processTask(&wg, task)
		time.Sleep(time.Second * 3)
	}

	go getCancelTaskId(&wg, tasks)

	fmt.Println("Waiting for tasks to complete...")
	wg.Wait()
}
