# Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM quay.io/fluentd_elasticsearch/fluentd:v3.0.2

COPY fluentd-es/Gemfile.gardener /Gemfile.gardener

RUN BUILD_DEPS="make build-essential g++ libc6-dev ruby-full" \
    && apt-get purge $BUILD_DEPS \
    || apt-get update \
    && apt-get -y install $BUILD_DEPS \
    && gem install --file Gemfile.gardener \
    && gem cleanup \
    && apt-get purge -y --auto-remove \
                     -o APT::AutoRemove::RecommendsImportant=false \
                     $BUILD_DEPS \
    && apt-get clean
