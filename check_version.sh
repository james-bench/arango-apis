#!/usr/bin/env bash

#
# DISCLAIMER
#
# Copyright 2020 ArangoDB GmbH, Cologne, Germany
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Copyright holder is ArangoDB GmbH, Cologne, Germany
#
# Author Gergely Brautigam
#

# Exit script if you try to use an uninitialized variable.
set -o nounset

# Exit script if a statement returns a non-true return value.
set -o errexit

# Use the error status of the first failure, rather than that of the last item in a pipeline.
set -o pipefail

ignore_string="ignore-version-check:"

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