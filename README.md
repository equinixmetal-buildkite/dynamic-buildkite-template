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
$ ./dynamic-buildkite-template trivy --version=v1.18.2 --skip-dirs="dir"
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/trivy#v1.18.2:
          timeout : 15m
          severity: HIGH,CRITICAL
          ignore-unfixed: true
          security-checks: vuln,config
          skip-dirs: 'dir'
```