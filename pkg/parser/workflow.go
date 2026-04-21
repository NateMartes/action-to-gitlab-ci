package parser

import (
	workflowParser "github.com/rhysd/actionlint"
)

// Takes a byte slice which containes a YAML Github Workflow and converts it into a Workflow struct
func UnmarshalWorkflow(inputWorkflow []byte) (*workflowParser.Workflow, error) {
	workflow, errs := workflowParser.Parse(inputWorkflow)
	if len(errs) != 0 {
		return workflow, errs[0]
	} else {
		return workflow, nil
	}
}