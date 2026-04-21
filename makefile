build:
	go build -o action-to-gitlab-ci cmd/action-to-gitlab-ci/main.go
    
install:
	go install .
    
clean:
	rm action-to-gitlab-ci