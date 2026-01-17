package main

import "task-cli/cmd"
var FILE = "task.json"



func main(){
	// err := utils.RemoveTask(FILE, 0)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// task, err := utils.ReadFile(file)
	// if err != nil {
	// 	fmt.Println("error retrieving data")
	// }
	// fmt.Println(task)

	cmd.Execute()
	// _ = utils.PrintTask("")
}