#!/bin/bash

set -eu

trivy_plugin="v1.18.0"
shell_plugin="v1.3.0"
docker_build_plugin="v5.3.0"

echo "steps:"

echo "  - command: ls"
echo "    plugins:"
echo "      - equinixmetal-buildkite/trivy#${trivy_plugin}"

echo "  - label: \":sparkles: SHELL CHECK\""
echo "    plugins:"
echo "      - shellcheck#${shell_plugin}:"
echo "          files: script.sh"
