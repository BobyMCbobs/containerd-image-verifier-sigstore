package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

var (
	DefaultProgramPath = "sigstore-policy-tester"
	DefaultPolicyPath  = "/etc/containers/sigstore/policy.yaml"
)

func main() {
	programPath := DefaultProgramPath
	if val, ok := os.LookupEnv("CONTAINERD_IMAGE_VERIFIER_SIGSTORE_EXEC"); ok {
		programPath = val
	}
	policyPath := DefaultPolicyPath
	if val, ok := os.LookupEnv("CONTAINERD_IMAGE_VERIFIER_SIGSTORE_POLICY_PATH"); ok {
		policyPath = val
	}
	name := flag.String("name", "", "the image to test")
	if name == nil {
		log.Fatalln("Error: must provide name")
	}
	flag.Parse()

	cmd := exec.Command(programPath, "-policy="+policyPath, "-image="+*name)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		log.Fatal()
	}
}
