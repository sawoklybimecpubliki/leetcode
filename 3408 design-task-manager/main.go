package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	userID   int
	taskID   int
	priority int
}

type TaskHeap []Task

func (h TaskHeap) Len() int {
	return len(h)
}

func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h TaskHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].taskID > h[j].taskID
	}
	return h[i].priority > h[j].priority
}

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(Task))
}

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TaskManager struct {
	h      TaskHeap
	record map[int]Task
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		h:      TaskHeap{},
		record: make(map[int]Task),
	}
	heap.Init(&tm.h)
	for _, t := range tasks {
		tm.Add(t[0], t[1], t[2])
	}
	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := Task{
		userID:   userId,
		taskID:   taskId,
		priority: priority,
	}
	heap.Push(&this.h, task)
	this.record[taskId] = task
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	old := this.record[taskId]
	updated := Task{
		userID:   old.userID,
		taskID:   taskId,
		priority: newPriority,
	}
	heap.Push(&this.h, updated)
	this.record[taskId] = updated
}

func (this *TaskManager) Rmv(taskId int) {
	delete(this.record, taskId)
}

func (this *TaskManager) ExecTop() int {
	for this.h.Len() > 0 {
		top := heap.Pop(&this.h).(Task)
		latest, ok := this.record[top.taskID]
		if !ok {
			continue
		}
		if latest.priority != top.priority {
			continue
		}
		delete(this.record, top.taskID)
		return latest.userID
	}
	return -1
}

func main() {
	tm := Constructor([][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}})
	tm.Add(4, 104, 5)
	tm.Edit(102, 9)
	fmt.Println(tm.ExecTop())
	tm.Rmv(101)
	tm.Add(50, 101, 8)
	fmt.Println(tm.ExecTop())
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
