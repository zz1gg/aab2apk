package internal

import (
	"aab2apk/pkg"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var aabtoapk []byte
var err error
var cmd *exec.Cmd

func ExtractAPK(target, outputname string) {
	bundletoolresult := pkg.CheckISFile("bundletool.jar")

	if bundletoolresult == false {
		log.Println("The directory does not exist bundletool！Plz go to https://github.com/google/bundletool/releases!")
	}
	apksname := outputname + "s"
	apksresult := pkg.CheckISFile(apksname)
	if apksresult == true {
		log.Println("The directory already exists ", apksname, "!")
	} else {

		aabtoapkcommand := "java -jar bundletool.jar build-apks --bundle=" + target + " --output=" + outputname + "s" + " --mode=universal"
		//log.Println("execute the command：" , aabtoapkcommand)
		switch runtime.GOOS {
		case "darwin":
			log.Println("Operating system is Mac")
			cmd = exec.Command("/bin/sh", "-c", aabtoapkcommand)
			if aabtoapk, err = cmd.Output(); err != nil {
				log.Println(err)
				//os.Exit(1)
			}
			log.Println(string(aabtoapk))
		case "windows":
			log.Println("Operating system is Windows")
			cmd := exec.Command("cmd", "/c", aabtoapkcommand)
			if aabtoapk, err = cmd.Output(); err != nil {
				log.Println(err)
				//os.Exit(1)
			}
		case "linux":
			log.Println("Operating system is Linux")
			cmd = exec.Command("/bin/sh", "-c", aabtoapkcommand)
			if aabtoapk, err = cmd.Output(); err != nil {
				log.Println(err)
				//os.Exit(1)
			}

		default:
			fmt.Println("Unrecognized operating system!")
			//cmd = exec.Command("/bin/sh", "-c", aabtoapkcommand)
			//fmt.Println(cmd) //
			//time.Sleep(time.Duration(2)*time.Second)
			apksresult2 := pkg.CheckISFile(apksname)
			if apksresult2 == true {
				log.Println("generate ", apksname, " successfully！")
			}
		}
		DeCompress(apksname, "./")
		universalapkresult := pkg.CheckISFile("universal.apk")
		if universalapkresult == true {
			log.Println("Generate universal.apk successfully!")
			err1 := os.Rename("universal.apk", outputname)
			if err1 != nil {
				panic(err1)
			} else {
				os.Remove("toc.pb")
				delapksname := outputname + "s"
				os.Remove(delapksname)
				log.Println("Successful!", target, " to ", outputname)
				//log.Println(AndroidBinInfo(outputname))
				APPInfo(outputname, target)
			}

		} else {
			log.Println("There is no universal.apk!")
		}
	}

}
