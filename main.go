package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
)

var (
	DefaultProgramPath   = "sigstore-policy-tester"
	DefaultPolicyDirPath = "/etc/containers/sigstore/"
)

func main() {
	log.Println("containerd-image-verifier-sigstore")

	programPath := DefaultProgramPath
	if val, ok := os.LookupEnv("CONTAINERD_IMAGE_VERIFIER_SIGSTORE_EXEC"); ok {
		programPath = val
	}
	policyDirPath := DefaultPolicyDirPath
	if val, ok := os.LookupEnv("CONTAINERD_IMAGE_VERIFIER_SIGSTORE_POLICY_DIR_PATH"); ok {
		policyDirPath = val
	}
	name := flag.String("name", "", "the image to test")
	if name == nil {
		log.Fatalln("Error: must provide name")
	}
	flag.Parse()

	dirEntries, err := os.ReadDir(policyDirPath)
	if err != nil {
		log.Fatalf("error: failed to read directory: %v\n", err)
	}
	if len(dirEntries) == 0 {
		log.Fatalf("error: no policy files found\n")
	}

	var wg sync.WaitGroup
	errBufs := map[string]*bytes.Buffer{}
	for _, dirEntry := range dirEntries {
		wg.Add(1)

		go func() {
			defer wg.Done()
			policyPath := path.Join(policyDirPath, dirEntry.Name())
			var stdoutb, errb bytes.Buffer
			cmd := exec.Command(programPath, "-policy="+policyPath, "-image="+*name)
			cmd.Stdout = &stdoutb
			cmd.Stderr = &errb
			if err := cmd.Run(); err != nil {
				errBufs[policyPath] = &errb
				errb.WriteString(err.Error())
			} else {
				log.Println(stdoutb.String())
				os.Exit(0)
			}
		}()
	}
	wg.Wait()
	if len(errBufs) > 0 {
		for run, buf := range errBufs {
			log.Println(run, "stderr")
			log.Println(buf.String())
		}
		os.Exit(1)
	}
}
