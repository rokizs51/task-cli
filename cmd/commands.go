package cmd

import (
	"strconv"
	"task-cli/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add [description]",
	Short: "add new task",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args[]string)  {
		err := utils.WriteTask("task.json", args[0])
		if err != nil {
			print(err.Error())
		}
	},
}

var updateCmd = &cobra.Command{
	Use: "update [id] [description]",
	Short: "update new task",
	Args: cobra.ExactArgs(2),
	Run: func (cmd *cobra.Command, args[]string)  {
		id , err := strconv.Atoi(args[0])
		if err != nil {
			println("invalid task ID")
			return
		}
		err = utils.UpdateTask("task.json", id, args[1])
		if err != nil {
			print(err.Error())
		}
	},
}

var markProgress = &cobra.Command{
	Use: "mark-in-progress [id]",
	Short: "update task status to in-progress",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args[]string)  {
		id , err := strconv.Atoi(args[0])
		if err != nil {
			println("invalid task ID")
			return
		}
		err = utils.UpdateTaskStatus("task.json", id, "in-progress")
		if err != nil {
			print(err.Error())
		}
	},
}

var markDone = &cobra.Command{
	Use: "mark-done [id]",
	Short: "update task status to done",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args[]string)  {
		id , err := strconv.Atoi(args[0])
		if err != nil {
			println("invalid task ID")
			return
		}
		err = utils.UpdateTaskStatus("task.json", id, "done")
		if err != nil {
			print(err.Error())
		}
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete [id]",
	Short: "delete task",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			println("invalid task ID")
			return
		}
		err = utils.RemoveTask("task.json", id)
		if err != nil {
			print(err.Error())
		}
	},
}


var listCmd = &cobra.Command{
      Use:   "list [status]",
      Short: "list all tasks",
      Args:  cobra.MaximumNArgs(1), // status is optional
      Run: func(cmd *cobra.Command, args []string) {
          status := ""
          if len(args) > 0 {
              status = args[0] // "done", "todo", "in-progress"
          }
          utils.PrintTask(status)
      },
  }

func init(){
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(markProgress)
	rootCmd.AddCommand(markDone)
	rootCmd.AddCommand(deleteCmd)
	
	
}