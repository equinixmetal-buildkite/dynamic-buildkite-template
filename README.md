# dynamic-buildkite-plugin

## Building Project
### Docker Image
Run this command to build a docker image
```
docker build -t dynamic-buildkite-plugin:<version> .
```
Mention desired version
### RPM package
To create an RPM package, run this:
```
make rpm
```
### Deb package
To create a Deb package, run this:
```
make deb
```

# Usage
Here's how you can generate the buildkite template
```
$ go run main.go
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/docker-build#v1.1.0:
          build-args:
            - NAME=REPO_NAME
          push: true
      - ssh://git@github.com/equinixmetal/ssm-buildkite-plugin#v1.0.4:
          parameters:
            COSIGN_KEY_SECRET : test-secret
            COSIGN_PASSWORD : passwd
      - equinixmetal-buildkite/cosign#main:
          image: ghcr.io/my-project/my-image:latest
          keyless : false
          keyed-config:
            key: cosign.key
          cosign-version: main
      - equinixmetal-buildkite/trivy#v1.18.3:
          exit-code: 0
          timeout : "5m0s"
          severity: "HIGH,CRITICAL"
          ignore-unfixed: true
          security-checks: "vuln,config"
          skip-files: ""
          skip-dirs: ""
          image-ref: ""
      - equinixmetal-buildkite/docker-metadata#v1.0.0:
          images:
          - "my-org/my-image"
          - "image2"
          extra_tags:
          - "latest"
          - "tag2"
```
## Configuration and Overrides
* Configurations are stored in `conf.yaml` and it has default values.
* Configurations from the file `conf.yaml` can be overridden by command line flags by using the yaml configuration path as below:
```
$ go run main.go --overrides plugins.trivy.skip-files="x.txt,y.txt" --overrides plugins.cosign.keyless=false
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/cosign#v0.1.0:
          image: ghcr.io/my-project/my-image:latest
          keyless : false
          keyed-config:
            key: sample-key
          cosign-version: v0.1.0
      - equinixmetal-buildkite/trivy#v1.18.3:
          timeout : 5m0s
          severity: HIGH,CRITICAL
          ignore-unfixed: true
          security-checks: vuln,config
          skip-files: 'x.txt,y.txt'
```
```
$ go run main.go --overrides plugins.trivy.skip-files="x.txt,y.txt" --overrides plugins.cosign.keyless=true
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/docker-build#v1.1.0:
          build-args:
            - NAME=REPO_NAME
          push: true
      - equinixmetal-buildkite/cosign#main:
          image: ghcr.io/my-project/my-image:latest
          keyless-config:
            fulcio-url: 
            rekor-url: 
          cosign-version: main
      - equinixmetal-buildkite/trivy#v1.18.3:
          exit-code: 0
          timeout : "5m0s"
          severity: "HIGH,CRITICAL"
          ignore-unfixed: true
          security-checks: "vuln,config"
          skip-files: "x.txt,y.txt"
          skip-dirs: ""
          image-ref: ""
      - equinixmetal-buildkite/docker-metadata#v1.0.0:
          images:
          - "my-org/my-image"
          - "image2"
          extra_tags:
          - "latest"
          - "tag2"
```
### Default conf.yaml for example
```
plugins:
  trivy:
    exit-code: 0
    timeout: 5m0s
    severity: HIGH,CRITICAL
    ignore-unfixed: true
    security-checks: vuln,config
    skip-files: ""
    skip-dirs: ""
    image-ref: ""
    version: ""
    helm-overrides-files: ""

```

Execute this command to run through a docker run
```
$ docker run --mount type=bind,source=${PWD}/conf.yaml,target=/go/src/workspace/dynamic-buildkite-template/conf.yaml ghcr.io/equinixmetal-buildkite/dynamic-buildkite-template:latest

output:
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/trivy#v1.18.3:
          exit-code: 0
          timeout : "5m0s"
          severity: "HIGH,CRITICAL"
          ignore-unfixed: true
          security-checks: "vuln,config"
          skip-files: ""
          skip-dirs: ""
          image-ref: ""
```
If you notice you can provide multiple `--overrides` flags and this would in turn collate to a `map[string]string` being passed to the program. The keys in override are in the yaml path format. So for a given config override you can check the path hierarchy in the `conf.yaml` and mention the override accordingly.

For long term config changes, it's suggested to update the `conf.yaml` file itself.