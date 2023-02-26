//help, bagaimana ini passing argumennya?

package main

import (
	"log"
	"os"
	"os/exec"
)

func main(){
	output := &exec.Cmd{
		Path :	"/usr/bin/mpv",
		Args :	[]string {"/usr/bin/mpv","https://b1world.beritasatumedia.com/Beritasatu/B1World_manifest.m3u8"},
		Stdout: os.Stdout,
        Stderr: os.Stderr,
    }
    err := output.Start()
    if err!=nil {
    	log.Fatal(err)
    }
}