#!/usr/bin/env bash
protoc -I ~/go/src/fleet_backend_final/proto  --go_out=. --micro_out=. ~/go/src/fleet_backend_final/proto/common.proto
protoc -I ~/go/src/fleet_backend_final/proto  --go_out=. --micro_out=. ~/go/src/fleet_backend_final/proto/customer.proto
protoc -I ~/go/src/fleet_backend_final/proto  --go_out=. --micro_out=. ~/go/src/fleet_backend_final/proto/truck qw.proto