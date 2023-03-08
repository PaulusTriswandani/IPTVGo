package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"encoding/csv"
	"strings"
	"runtime"
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

func searchChannelInCSVFile(channel string) string{
	csvFile, err := os.Open("iptv.csv")
	if err != nil {
    	println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        println(err)
    }
    for _, line := range csvLines {
    	if strings.Compare(channel,line[0])==0 {
        	return (line[1])
    	}
    }
    return ""
}


func main(){
	channel := flag.String("channel","BeritaSatu","Channel Name that you want to watch")
	flag.Parse()

	switch currOS := runtime.GOOS; currOS {
	case "darwin":
		println("Mac OS")
	case "linux":
		executeExternalProgram("/usr/bin/mpv",searchChannelInCSVFile(*channel))
	}
}