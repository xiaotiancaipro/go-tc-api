package main

import "go-tc-api/cmd"

func main() {
	GoTCAPICli := cmd.Root{}.Init()
	GoTCAPICli.Execute()
}
