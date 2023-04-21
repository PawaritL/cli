#!/usr/bin/env bash

set -e

latest=$(gh release list | grep -v snapshot | head -1 | awk '{print $1}')
echo "Downloading ${latest}..."

mkdir "${latest}"
pushd "${latest}"
gh release download "${latest}" -p "*.zip"
popd

version=$(sed -e 's/v//' <<<"$latest")
sed -e "s/\$version/$version/g" ./index.html > $latest/index.html

ls -l "${latest}"
export AWS_PROFILE=bricks
s3cmd sync -v "${latest}" s3://databricks-bricks/
