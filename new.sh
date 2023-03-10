#!/usr/bin/env bash
projectName=$1

lastPrefix="$(find . -name "*_*" -not -path "*/.*" -type d -maxdepth 1 -exec basename {} \; | sort | cut -d "_" -f 1 | tail -n1)"
newPrefix="$((lastPrefix + 1))"
if [[ newPrefix -lt 10 ]]; then
  newPrefix="0${newPrefix}"
fi

projectDirectory="${newPrefix}_${projectName}"

function project_init() {
	mkdir -p "${projectDirectory}/bin"

cat <<EOF > "${projectDirectory}/go.work"
go 1.19

use (
	.
)
EOF

cat <<EOF > "${projectDirectory}/main.go"
package main

import "fmt"

func main() {
  fmt.Println("Yo!")
}
EOF

cat <<EOF > "${projectDirectory}/Makefile"
build: clean
	go build -o bin/${projectName}
clean:
	rm -rf bin/*

EOF

  echo "*" > "${projectDirectory}/bin/.gitignore"

  (
    cd "${projectDirectory}";
    go mod init "${projectName}";
    go mod tidy;
    git add -A
  )
}

project_init