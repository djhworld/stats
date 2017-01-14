all:
	gb build all

vendor:
	gb vendor fetch -tag 0.2.0 github.com/montanaflynn/stats 
	gb vendor fetch -tag v1.19.1 github.com/urfave/cli

clean:
	rm -rf pkg
	rm -rf bin/
