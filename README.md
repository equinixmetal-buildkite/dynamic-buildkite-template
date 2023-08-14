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
Here's how you can generate the trivy plugin template
```
$ ./dynamic-buildkite-template trivy --version=v1.18.2 --skip-files="cosign.key"
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/trivy#v1.18.2:
          timeout : 5m0s
          severity: HIGH,CRITICAL
          ignore-unfixed: true
          security-checks: vuln,config
          skip-files: 'cosign.key'
```
## Configuration and Overrides
* Configurations are stored in `conf.yaml` and it has default values.
* Configurations from the file `conf.yaml` can be overridden by command line flags as this example:
  Using default configs
  ```
  $ go run main.go trivy
  steps:
    - command: ls
      plugins:
        - equinixmetal-buildkite/trivy#v1.18.2:
            timeout : 5m0s
            severity: HIGH,CRITICAL
            ignore-unfixed: true
            security-checks: vuln,config
  ```

  Using command line flags to override timeout
  ```
  $ go run main.go trivy --timeout=7m15s
  steps:
    - command: ls
      plugins:
        - equinixmetal-buildkite/trivy#v1.18.2:
            timeout : 7m15s
            severity: HIGH,CRITICAL
            ignore-unfixed: true
            security-checks: vuln,config
  ```