package hack

import (
	//"io/ioutil"
	"os"
	"os/exec"
	"fmt"
)

func Hack(){
	//ioutil.WriteFile("/course/cs666/tabin/hack",[]byte("#!/bin/bash\n/home/whoami\n"),0777)
	fmt.Println("Running inserted command!")
	cmd := exec.Command("/home/whoami")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
