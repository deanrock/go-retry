package main

import "os"

func main() {
	cli := &CLI{
		inStream:  os.Stdin,
		outStream: os.Stdout,
		errStream: os.Stderr,
	}

	command = RealCommand{
		inStream:  cli.inStream,
		outStream: cli.outStream,
		errStream: cli.errStream,
	}

	os.Exit(cli.Run(os.Args))
}
