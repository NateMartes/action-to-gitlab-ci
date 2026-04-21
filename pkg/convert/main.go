package convert

import (
	gitlabTypes "github.com/NateMartes/action-to-gitlab-ci/pkg/types/gitlab"
	workflowTypes "github.com/rhysd/actionlint"
)

// Convert a Github Workflow to a Gitlab CI struct
func ConvertWorkflowToGitlabCI(val *workflowTypes.Workflow) (*gitlabTypes.GitlabCI, error) {
	
	out := &gitlabTypes.GitlabCI{}
	var err error = nil

	stages, err := ConvertJobsToStages(val.Jobs)
	if err != nil {
		return nil, err
	}
	out.Stages = stages
	
	return out, nil
}

// Convert a Github Job map to a Gitlab Stage map
func ConvertJobsToStages(val map[string]*workflowTypes.Job) (map[string]*gitlabTypes.Stage, error) {
	
	out := map[string]*gitlabTypes.Stage{}
	
	for name, job := range val {
		stage, err := ConvertJobToStage(job)
		if err != nil {
			return nil, err
		}
		out[name] = stage
	}
	
	return out, nil
}

// Convert a Github Job to a Gitlab Stage
func ConvertJobToStage(val *workflowTypes.Job) (*gitlabTypes.Stage, error) {
	return &gitlabTypes.Stage{}, nil
}