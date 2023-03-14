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