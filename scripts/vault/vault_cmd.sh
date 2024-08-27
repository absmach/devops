#!/usr/bin/bash
# Copyright (c) Abstract Machines
# SPDX-License-Identifier: Apache-2.0

vault() {
    kubectl exec magistrala-vault-0 -n mg -- vault "$@"
}
