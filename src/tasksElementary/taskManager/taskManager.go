package taskManager

import (
	"tasksElementary/task1"
	"tasksElementary/task5"
	"errors"
	"tasksElementary/task2"
)
var allTask = make(map[string]func ([]byte)(string, error))
func init(){
	allTask["task1"] = task1.Run
	allTask["task2"] = task2.Run
	allTask["task5"] = task5.Run
}
func RunTask(key string, data []byte) (string , error){

	task, ok := allTask[key]
	if !ok{
		return "", errors.New("Can't find this task")
	}
	return task(data)
}