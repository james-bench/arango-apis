#!/usr/bin/env bash

# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail

check_file() {
    fullpath=$1
    file=`basename ${fullpath}`
    dir=`dirname ${fullpath}`
    version_file="${dir}/version.go"

    pattern="*.proto"
    if [[ ! $file == $pattern ]]; then
        return 0
    fi

    echo "Found proto change in ${dir}. Checking if there is a version change as well..."

    if [[ ! -f "${version_file}" ]]; then
        echo "${version_file} file does not exist for corresponding proto file."
        exit 1
    fi

    status=`git status ${version_file}`
    if [[ ! "${status}" == *"modified"* ]]; then
        echo "no version change found in ${version_file} but proto file was changed."
        exit 1
    fi
    echo "done. All good."
}

# Get last changes and see if the list contains the corresponding version go file.
git diff-tree --no-commit-id --name-only -r HEAD |
    while read line ; do
        check_file $file
    done
