package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func executeExternalProgram(program string, argument1 string){
	output := &exec.Cmd{
		Path :	program,
		Args :	[]string {program,argument1},
		Stdout: os.Stdout,
        Stderr: os.Stderr,
    }
    err := output.Start()
    if err!=nil {
    	log.Fatal(err)
    }

    err = output.Wait()
    if err != nil {
    	log.Fatal(err)
    }
}

func main(){
	program := flag.String("channel","BeritaSatu","Channel Name that you want to watch")
	executeExternalProgram("/usr/bin/mpv","https://b1world.beritasatumedia.com/Beritasatu/B1World_manifest.m3u8")
}