#!/bin/bash
set -e -x

go build ./match
go build ./cli/test_matching
go build ./cli/generate_patients
