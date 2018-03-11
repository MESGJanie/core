package cmd

import (
	"github.com/mesg-foundation/application/cmd/workflow"

	"github.com/spf13/cobra"
)

// Workflow is the root command related to workflows
var Workflow = &cobra.Command{
	Use:   "workflow",
	Short: "Manage the workflow you create",
}

func init() {

	Workflow.AddCommand(cmdWorkflow.Pause)
	Workflow.AddCommand(cmdWorkflow.Resume)
	Workflow.AddCommand(cmdWorkflow.Deploy)

	RootCmd.AddCommand(Workflow)
}
