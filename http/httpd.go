package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os/exec"
	"log"
	"sync"
	"io/ioutil"
)

func checkerr(err error)  {
	if err != nil {
		print("Went worng")
	}
}

func main() {
	//http.HandleFunc("/", hello)
	//log.Fatal(http.ListenAndServe(":8080", nil))*/
	wg := new(sync.WaitGroup)
	wg.Add(3)
	//signer,_ := ssh.ParsePrivateKey([]byte("serveradm1"))
	buf,err := ioutil.ReadFile("serveradm1.ppk")
	checkerr(err)
	key, err := ssh.ParsePrivateKey(buf)
	checkerr(err)
	config := &ssh.ClientConfig{
		User: "serveradm",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	if _, err := ssh.Dial("tcp", "10.138.32.75:22", config); err == nil {
		out, err := exec.Command("df").Output()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Sprint("%s", out)
		wg.Done()
	}

}
