#!/usr/bin/env bash

set -e;

CURRENT_PATH=$(cd "$(dirname "${0}")" || exit 1; pwd)

cd "${CURRENT_PATH}"/.. || exit 1

export PATH="./build/bin:$PATH"

PLUGIN="godefault"

MODULE="github.com/yu31/protoc-plugin-default"

# Delete old file first.
/bin/rm -fr ./xgo/tests/pb/*/*.pb.go;
#/bin/rm -fr ./xgo/tests/cases/tags/*/*.pb.go;

# generate code for module file.
for file in tests/proto/module/*; do
  if [ -d "${file}" ]; then
    continue
  fi
  protoc -I=. --go_opt=module="${MODULE}" --"${PLUGIN}"_opt=module="${MODULE}" --go_out=./ --"${PLUGIN}"_out=./  "${file}"

#  protoc -I=. --go_out=./ --"${PLUGIN}"_out=./  "${file}"
#  protoc -I=. -I=./xgo --go_opt=paths=source_relative --"${PLUGIN}"_opt=paths=source_relative --go_out=. --"${PLUGIN}"_out=. "${file}"
done

for file in tests/proto/cases/*/*; do
  if [ -d "${file}" ]; then
    for f in "${file}"/*; do
      protoc -I=. --go_out=./ --"${PLUGIN}"_out=./  "${f}"
    done
  else
    protoc -I=. --go_out=./ --"${PLUGIN}"_out=./  "${file}";
  fi

#  protoc -I=. --go_out=./ --"${PLUGIN}"_out=./  "${file}"
#  protoc -I=. -I=./xgo --go_opt=paths=source_relative --"${PLUGIN}"_opt=paths=source_relative --go_out=. --"${PLUGIN}"_out=. "${file}"
done

# generate test file for test custom method func name.
for file in tests/proto/customfn/*; do
  protoc -I=. --${PLUGIN}_opt="fn=InitField" --go_out=./ --"${PLUGIN}"_out=./  "${file}"
done

