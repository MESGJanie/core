package cmd

import (
	"github.com/spf13/cobra"
)

// Workflow is the root command related to workflows
var Workflow = &cobra.Command{
	Use:               "workflow",
	Short:             "Manage your workflows",
	DisableAutoGenTag: true,
}

// TODO this command is disabled for now waiting for the Workflow feature to be implemented
func init() {

	// Workflow.AddCommand(cmdWorkflow.Pause)
	// Workflow.AddCommand(cmdWorkflow.Resume)
	// Workflow.AddCommand(cmdWorkflow.Deploy)
	// Workflow.AddCommand(cmdWorkflow.Test)
	// Workflow.AddCommand(cmdWorkflow.Validate)
	// Workflow.AddCommand(cmdWorkflow.List)
	// Workflow.AddCommand(cmdWorkflow.Kill)
	// Workflow.AddCommand(cmdWorkflow.Topup)
	// Workflow.AddCommand(cmdWorkflow.Log)
	// Workflow.AddCommand(cmdWorkflow.Detail)

	// RootCmd.AddCommand(Workflow)
}
