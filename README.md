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
$ ./dynamic-buildkite-template
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/cosign#v0.1.0:
          image: ghcr.io/my-project/my-image:latest
          keyless-config:
            fulcio-url: https://fulcio.sigstore.dev
            rekor-url: https://rekor.sigstore.dev
          cosign-version: v0.1.0
      - equinixmetal-buildkite/trivy#v1.18.3:
          timeout : 5m0s
          severity: HIGH,CRITICAL
          ignore-unfixed: true
          security-checks: vuln,config
```
## Configuration and Overrides
* Configurations are stored in `resources/config/conf.yaml` and it has default values.
* Configurations from the file `resources/config/conf.yaml` can be overridden by command line flags by using the yaml configuration path as below:
```
$ ./dynamic-buildkite-template --overrides plugins.trivy.skip-files="x.txt,y.txt" --overrides plugins.cosign.keyless=false
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
If you notice you can provide multiple `--overrides` flags and this would in turn collate to a `map[string]string` being passed to the program. The keys in override are in the yaml path format. So for a given config override you can check the path hierarchy in the `conf.yaml` and mention the override accordingly.

For long term config changes, it's suggested to update the `conf.yaml` file itself.