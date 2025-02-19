#!/bin/bash
# Copyright (c) Abstract Machines
# SPDX-License-Identifier: Apache-2.0

kubectl -n mg create secret tls magistrala-server \
    --key magistrala-server.key \
    --cert magistrala-server.crt

kubectl -n mg create secret generic ca \
    --from-file=ca.crt
