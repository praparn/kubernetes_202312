# Changelog

### 1.7.0

Images:

* registry.k8s.io/ingress-nginx/controller:v1.7.0@sha256:7612338342a1e7b8090bef78f2a04fffcadd548ccaabe8a47bf7758ff549a5f7
* registry.k8s.io/ingress-nginx/controller-chroot:v1.7.0@sha256:e84ef3b44c8efeefd8b0aa08770a886bfea1f04c53b61b4ba9a7204e9f1a7edc

### All changes:

* kick off 1.7.0 build (#9775)
* Update exposing-tcp-udp-services.md (#9777)
* feat: OpenTelemetry module integration (#9062)
* drop k8s 1.23 support (#9772)
* Fix canary-weight-total annotation ignored in rule backends (#9729)
* fix: controller psp's volume config (#9740)
* Fix several Helm YAML issues with extraModules and extraInitContainers (#9709)
* docs(helm): fix value key in readme for enabling certManager (#9640)
* updated digest and sha for e2e-test-echo (#9760)
* updated digest and sha for e2e-test-fastcgi-helloserver (#9759)
* updated digest and sha for opentelemetry (#9758)
* updated digest and sha for e2e-test-cfssl (#9757)
* updated kube-webhook-certgen digest and tags (#9756)
* updated nginx-error digest and tags (#9755)
* added upgrade ginkgo documentation for contributors (#9753)
* changes Makefile of echo folder to trigger code-build (#9754)
* Chart: Drop `controller.headers`, rework DH param secret. (#9659)
* updated NGINX_BASE image with latest tag (#9747)
* Deployment/DaemonSet: Label pods using `ingress-nginx.labels`. (#9732)
* bumped ginkgo to v2.9.0 (#9722)
* HPA: autoscaling/v2beta1 deprecated, bump apiVersion to v2 for defaultBackend (#9731)
* update to golang 1.20 (#9690)
* Indent values.yaml using 2 instead of 4 spaces (#9656)
* fix some comments (#9688)
* migrate mitchellh/hashstructure to v2 (#9651)
* changed v1.6.3 to v1.6.4 on deploy docs (#9647)
* controller: Don't panic when ready condition in a endpointslice is missing (#9550)
* Rework Ginkgo usage (#9522)
* code clean for fsnotify (#9571)
* Optimize the document for readability (#9551)
* sets.String is deprecated: use generic Set instead. new ways: s1 := Set[string]{} s2 := New[string]() (#9589)
* Adjust the import package order and use http library variables (#9587)
* Optimize the judgment mode to remove redundant transformations (#9588)
* Fix rewrite example (#9633)
* remove tests and regex path checks (#9626)
* Fix incorrect annotation name in upstream hashing configuration (#9617)
* Release docs for Controller v1.6.3 and Helm v4.5.0 (#9614)

### Dependency updates:

* Bump aquasecurity/trivy-action from 0.8.0 to 0.9.2 (#9767)
* Bump k8s.io/component-base from 0.26.2 to 0.26.3 (#9764)
* Bump actions/dependency-review-action from 3.0.3 to 3.0.4 (#9766)
* Bump actions/add-to-project from 0.4.0 to 0.4.1 (#9765)
* Bump actions/dependency-review-action from 3.0.2 to 3.0.3 (#9727)
* Bump github.com/prometheus/common from 0.41.0 to 0.42.0 (#9724)
* Bump golang.org/x/crypto from 0.6.0 to 0.7.0 (#9723)
* Bump actions/download-artifact from 3.0.1 to 3.0.2 (#9721)
* Bump goreleaser/goreleaser-action from 4.1.0 to 4.2.0 (#9718)
* Bump actions/upload-artifact from 3.1.1 to 3.1.2 (#9717)
* Bump docker/setup-buildx-action from 2.2.1 to 2.5.0 (#9719)
* Bump helm/chart-releaser-action from 1.4.1 to 1.5.0 (#9720)
* Bump github.com/onsi/ginkgo/v2 from 2.6.1 to 2.9.0 (#9695)
* Bump k8s.io/klog/v2 from 2.90.0 to 2.90.1 (#9694)
* Bump golang.org/x/crypto in /magefiles (#9691)
* Bump k8s.io/component-base from 0.26.1 to 0.26.2 (#9696)
* Bump github.com/prometheus/common from 0.40.0 to 0.41.0 (#9698)
* Bump sigs.k8s.io/controller-runtime from 0.14.2 to 0.14.5 (#9697)
* Bump golang.org/x/net in /magefiles (#9692)
* Bump golang.org/x/sys in /images/custom-error-pages/rootfs (#9671)
* Bump github.com/stretchr/testify from 1.8.1 to 1.8.2 (#9675)
* Bump github.com/prometheus/common from 0.39.0 to 0.40.0 (#9653)
* Bump golang.org/x/net from 0.6.0 to 0.7.0 (#9646)
* Bump golang.org/x/net in /images/kube-webhook-certgen/rootfs (#9645)
* Bump google.golang.org/grpc from 1.52.3 to 1.53.0 (#9610)
* Bump github.com/prometheus/client_golang (#9630)
* Bump golang.org/x/crypto from 0.5.0 to 0.6.0 (#9609)

**Full Changelog**: https://github.com/kubernetes/ingress-nginx/compare/controller-controller-v1.6.3...controller-controller-v1.7.0
