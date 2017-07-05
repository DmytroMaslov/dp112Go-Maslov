package taskManager

import (
	"tasksElementary/task1"
	"tasksElementary/task5"
	"errors"
	"tasksElementary/task2"
)
func RunTask(key string, data []byte) (string , error){

allTask := map[string]func ([]byte)(string, error){
	"task1": task1.Run,
	"task2": task2.Run,
	"task5": task5.Run,
}
	task, ok := allTask[key]
	if !ok{
		return "", errors.New("Can't find this task")
	}
	return task(data)
}