#!/bin/bash
cd ./backend
go run main.go 
cd ../frontend
npm install
npm run dev
