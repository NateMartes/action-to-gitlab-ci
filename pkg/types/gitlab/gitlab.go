package gitlab

// A representation of a GitlabCI file
type GitlabCI struct {
	Stages map[string]*Stage
}

// A representation of a GitlabCI stage
type Stage struct {
}