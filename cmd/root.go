package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

 var rootCmd = &cobra.Command{
      Use:   "task-cli",
      Short: "A simple task CLI",
      Long:  "A CLI application for managing tasks",
  }

   func Execute() {
      if err := rootCmd.Execute(); err != nil {
          os.Exit(1)
      }
  }