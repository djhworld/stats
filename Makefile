install:
	go install

deps:
	gvt fetch -tag 0.2.0 github.com/montanaflynn/stats
	gvt fetch -tag v1.19.1 github.com/urfave/cli
