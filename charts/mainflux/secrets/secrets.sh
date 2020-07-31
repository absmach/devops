#!/bin/bash
# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

kubectl -n mf create secret tls mainflux-server \
    --key mainflux-server.key \
    --cert mainflux-server.crt

kubectl -n mf create secret generic ca \
    --from-file=ca.crt
