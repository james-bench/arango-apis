#!/usr/bin/env bash

# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail

ignore_string="ignore-ci:"

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

    status=`git diff-tree --no-commit-id --name-only -r HEAD ${version_file}`
    if [[ -z "${status}" ]]; then
        echo "no version change found in ${version_file} but proto file was changed."
        exit 1
    fi
    echo "done. All good."
}

echo "Running check for version.go consistency..."
echo "To ignore a commit, add '${ignore_string}' at the begin of he commit message."

commit_message=`git log -1 --pretty=%B`
if [[ $commit_message == $ignore_string* ]]; then
    echo "Ignore directive found in commit message."
    exit 0
fi

# Get last commit and check to see if it contained a version change as well.
# This is convenient and will check changes on every commit.
git diff-tree --no-commit-id --name-only -r HEAD |
    while read line ; do
        check_file $line
    done

echo "Finished."