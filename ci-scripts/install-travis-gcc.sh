#!/usr/bin/env bash

if [[ "$TRAVIS_OS_NAME" == "linux" ]]; then
  sudo apt-get update
  sudo apt-get install -y  gcc g++ build-essential libgl1-mesa-dev libfontconfig1-dev libfreetype6-dev libx11-dev libxext-dev libxfixes-dev libxi-dev libxrender-dev libxcb1-dev libx11-xcb-dev libxcb-glx0-dev libxkbcommon-x11-dev
  go get -u -v github.com/therecipe/qt/cmd/...
  (qtsetup -docker -test=false | true)
fi

if [[ "$TRAVIS_OS_NAME" == "osx" ]]; then
  echo 'Updating packages database'
  brew update
  echo 'Install QT'
  brew install qt5
  go get -u -v github.com/therecipe/qt/cmd/...
fi

cd $TRAVIS_BUILD_DIR
