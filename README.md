# containerd-image-verifier-sigstore

> A shim between containerd's ImageVerifier and Sigstore's policy-tester

## Requirements

- containerd <v2.0
- policy-tester <v0.12.0

## Configuring containerd

containerd must be configured to something like

``` toml
[plugins]
  [plugins."io.containerd.image-verifier.v1.bindir"]
    bin_dir = "/opt/containerd/image-verifier/bin"
    max_verifiers = 10
    per_verifier_timeout = "10s"
```

where `containerd-image-verifier-sigstore exists` in the defined `bin_dir` above.

see docs: https://github.com/containerd/containerd/blob/main/docs/image-verification.md

## Testing locally

Build policy-tester

``` bash
cd ~/src/github.com/sigstore/policy-controller
go build -o bin/policy-tester ./cmd/tester
```

Run the test policy

``` bash
cd ~/src/github.com/BobyMCbobs/containerd-image-verifier-sigstore

export CONTAINERD_IMAGE_VERIFIER_SIGSTORE_EXEC=/Users/calebwoodbine/src/github.com/sigstore/policy-controller/bin/policy-tester
export CONTAINERD_IMAGE_VERIFIER_SIGSTORE_POLICY_PATH=./test-policy.yaml 

go build -o bin/ .

# valid for test-policy.yaml
bin/containerd-image-verifier-sigstore -name registry.gitlab.com/flattrack/flattrack:latest
bin/containerd-image-verifier-sigstore -name registry.k8s.io/pause:3.9
bin/containerd-image-verifier-sigstore -name quay.io/jetstack/cert-manager-controller:v1.17.0
```

## License

Copyright 2025 Caleb Woodbine.
This project is licensed under the [AGPL-3.0](http://www.gnu.org/licenses/agpl-3.0.html) and is [Free Software](https://www.gnu.org/philosophy/free-sw.en.html).
This program comes with absolutely no warranty.
