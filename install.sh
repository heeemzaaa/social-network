#!/bin/bash

cd ./backend 
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
cd $HOME
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
export GOBIN=$GOPATH/bin

# to run this and have effects we should use source install.sh 

cd Desktop/social-network