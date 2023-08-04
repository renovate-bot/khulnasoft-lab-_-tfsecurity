#!/bin/bash

pushd ../tfsecurity.wiki
git pull
git add .
git commit -a -m "Updating links and wiki entries"
git push
popd