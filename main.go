package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
)

type tokenInfo struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Status     struct {
		Token               string    `json:"token"`
		ExpirationTimestamp time.Time `json:"expirationTimestamp"`
	} `json:"status"`
}

func main() {
	flag.Parse()
	cu, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	tokenFile := filepath.Join(cu.HomeDir, ".oci", "token-cache.json")
	f, err := os.Open(tokenFile)
	if err == nil {
		defer f.Close()
		var ti tokenInfo
		err = json.NewDecoder(f).Decode(&ti)
		if err != nil {
			log.Fatal(err)
		}
		if time.Now().Add(time.Second * 30).Before(ti.Status.ExpirationTimestamp) {
			err = json.NewEncoder(os.Stdout).Encode(ti)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	var buf bytes.Buffer
	cmd := exec.Command(flag.Arg(0), flag.Args()[1:]...)
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				os.Exit(s.ExitStatus())
			} else {
				panic(errors.New("Unimplemented for system where exec.ExitError.Sys() is not syscall.WaitStatus."))
			}
		}
	}
	err = ioutil.WriteFile(tokenFile, buf.Bytes(), 0600)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, &buf)
	if err != nil {
		log.Fatal(err)
	}
}
