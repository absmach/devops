#!/bin/bash
# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

kubectl -n mf create secret tls nginx-internal-mtls-tls \
    --key nginx-internal-mtls.key \
    --cert nginx-internal-mtls.crt

kubectl -n mf create secret generic nginx-internal-mtls-intermediate-crt \
    --from-file=intermediate.crt=nginx-internal-mtls-intermediate.crt
