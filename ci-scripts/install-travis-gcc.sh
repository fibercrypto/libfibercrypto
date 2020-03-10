#!/usr/bin/env bash

if [[ "$TRAVIS_OS_NAME" == "linux" ]]; then
  sudo apt-get update
  sudo apt-get install -y libgl1-mesa-dev gcc g++ build-essential libqt5multimedia5 libpulse-mainloop-glib0
  go get -u github.com/therecipe/qt/cmd/...
  (qtsetup -test=false | true)
fi

if [[ "$TRAVIS_OS_NAME" == "osx" ]]; then
  echo 'Updating packages database'
  brew update
  echo 'Install QT'
  export GO111MODULE=off; go get github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup -test=false
fi

cd $TRAVIS_BUILD_DIR
