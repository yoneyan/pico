package main

import (
	"log"
	"time"
)

type Job struct {
	Name    string   `yaml:"name"`
	Actions []Action `yaml:"actions"`
}

type Action struct {
	Action  string      `yaml:"action"`
	Command interface{} `yaml:"command"`
}

type PatliteCommand struct {
	Name string `yaml:"name"`
	LED  string `yaml:"led"`
}

func checkOverwritePattern() bool {
	if len(job.Actions) != 0 {
		return false
	}
	switch patternName {
	case "default":
		job = Job{}
	default:
		jobTmp := GetJobName(patternName)
		if jobTmp == nil {
			return false
		}
		job = *jobTmp
	}
	return true
}

func ExecuteJob() {
	if len(job.Actions) != 0 {
		return
	}
	action := job.Actions[0]
	job.Actions = job.Actions[1:]
	switch action.Action {
	case "send_patlite":
		patliteCommands := action.Command.([]PatliteCommand)
		for _, patliteCommand := range patliteCommands {
			patliteConfig := GetPatliteConfig(patliteCommand.Name)
			if patliteConfig == nil {
				continue
			}
			if patliteConfig.IsHttp {
				err := sendPatliteViaHttp(patliteConfig.Host, patliteCommand.LED)
				if err != nil {
					log.Println("Error sending patlite command via HTTP:", err)
				}
			} else {
				//err := sendPatliteViaSocket(patliteConfig.Host, patliteCommand.LED)
				//if err != nil {
				//	log.Println("Error sending patlite command via Socket:", err)
				//}
			}
		}
	case "wait":
		// Wait for the specified duration
		duration := action.Command.(int)
		time.Sleep(time.Duration(duration) * time.Second)
	}

}
