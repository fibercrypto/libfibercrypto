#!/usr/bin/env bash

if [[ "$TRAVIS_OS_NAME" == "linux" ]]; then
  sudo apt-get update
  sudo apt-get install  gcc g++ -y
  sudo apt-get -y install build-essential libgl1-mesa-dev --install-suggests
  go get -u -v github.com/therecipe/qt/cmd/...
  (qtsetup -test=false | true)
fi

if [[ "$TRAVIS_OS_NAME" == "osx" ]]; then
  echo 'Updating packages database'
  brew update

  echo 'Install QT'
  brew install qt5
  go get -u -v github.com/therecipe/qt/cmd/...
  qtsetup -test=false || true
fi

cd $TRAVIS_BUILD_DIR
