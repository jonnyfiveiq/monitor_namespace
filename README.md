monitor_namespace
========

## ***FILL THIS OUT WITH A USEFUL DESCRIPTION OF THIS REPO***

## Building

`make`

'kubectel create -f regcd.yaml adds quay registry to k8s using local docker.json'

'kubectel create -f https://github.com/jonnyfiveiq/monitor_namespace/blob/master/serviceaccount-rolebinding.yaml'

`docker build -t quay.io/jonnyfiveiq/final:latest . -f build/Dockerfile`

`kubectl run quaylatest --image=quay.io/jonnyfiveiq/final:latest --image-pull-policy=Never`

`kubectl get pods`

`kubectl logs quaylatest-584fb685c5-p2kg7`


## Running

`./bin/monitor_namespace`

## License
Copyright (c) 2019 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
