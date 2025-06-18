type TaskManager struct {
    h *TaskHeap
    priorityMap map[int] int  // taskID : priority
    taskToUser map[int] int // taskID : userID
}


func Constructor(tasks [][]int) TaskManager {
    t := TaskManager{ 
        h: &TaskHeap{},
        priorityMap: map[int] int {},
        taskToUser: map[int] int {},
    }
    heap.Init(t.h)
    for _, task := range tasks {
        heap.Push(t.h, taskType{task[1], task[2]})
        t.priorityMap[task[1]] = task[2]
        t.taskToUser[task[1]] = task[0]
    }

    return t
}


func (this *TaskManager) Add(userId int, taskId int, priority int)  {
    heap.Push(this.h, taskType{taskId, priority})
    this.priorityMap[taskId] = priority
    this.taskToUser[taskId] = userId
}


func (this *TaskManager) Edit(taskId int, newPriority int)  {
    this.Add(this.taskToUser[taskId], taskId, newPriority)
}


func (this *TaskManager) Rmv(taskId int)  {
    delete(this.priorityMap, taskId)
}


func (this *TaskManager) ExecTop() int {
    for this.h.Len() > 0 {
        pop := heap.Pop(this.h).(taskType)

        if val, ok := this.priorityMap[pop.taskID]; ok {
            if val == pop.priority {
                this.Rmv(pop.taskID)
                return this.taskToUser[pop.taskID]
            }
        }
    }

    return -1
}

type taskType struct {
    taskID, priority int
}

type TaskHeap []taskType

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { 
    if h[i].priority == h[j].priority {
        return h[i].taskID > h[j].taskID
    }
    return h[i].priority > h[j].priority
}
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x any) {*h = append(*h, x.(taskType))}
func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
