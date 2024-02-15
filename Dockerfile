# Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
# All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http:www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.21.4-alpine3.18 AS builder

ARG SSH_PRIVATE_KEY

# Prepare SSH mode for downloading git repositories/dependencies
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa
RUN echo "StrictHostKeyChecking no" >> /root/.ssh/config

# Force git to use SSH over HTTPS to avoid password prompt
RUN apk add git openssh
RUN git config --global --add url."git@wwwin-github.cisco.com:".insteadOf "https://wwwin-github.cisco.com/"

RUN mkdir -p /root/go/src/wwwin-github.cisco.com/awi-infra-guard

WORKDIR /root/go/src/wwwin-github.cisco.com/awi-infra-guard
COPY . .

RUN go build -o awi-infra-guard .

# Second stage: create the runtime image
FROM alpine:3.18.4

WORKDIR /root/
COPY --from=builder /root/go/src/wwwin-github.cisco.com/awi-infra-guard/awi-infra-guard .

# As k8s mounting makes it hard to create a file from a config map
# within the directory with already present files, we create a symlink
# to point to a new empty directory where actual config.yaml will be
# mounted.
RUN ln -s /root/config/config.yaml /root/config.yaml

CMD ["./awi-infra-guard"]
