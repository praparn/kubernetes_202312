/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log

import (
	"testing"

	api "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

func buildIngress() *networking.Ingress {
	defaultBackend := networking.IngressBackend{
		Service: &networking.IngressServiceBackend{
			Name: "default-backend",
			Port: networking.ServiceBackendPort{
				Number: 80,
			},
		},
	}

	return &networking.Ingress{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "foo",
			Namespace: api.NamespaceDefault,
		},
		Spec: networking.IngressSpec{
			DefaultBackend: &networking.IngressBackend{
				Service: &networking.IngressServiceBackend{
					Name: "default-backend",
					Port: networking.ServiceBackendPort{
						Number: 80,
					},
				},
			},
			Rules: []networking.IngressRule{
				{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path:    "/foo",
									Backend: defaultBackend,
								},
							},
						},
					},
				},
			},
		},
	}
}

func TestIngressAccessLogConfig(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(enableAccessLogAnnotation)] = "false"
	ing.SetAnnotations(data)

	log, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	nginxLogs, ok := log.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}

	if nginxLogs.Access {
		t.Errorf("expected access be disabled but is enabled")
	}
}

func TestIngressRewriteLogConfig(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(enableRewriteLogAnnotation)] = "true"
	ing.SetAnnotations(data)

	log, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("unexpected error parsing annotations %v", err)
	}
	nginxLogs, ok := log.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}

	if !nginxLogs.Rewrite {
		t.Errorf("expected rewrite log to be enabled but it is disabled")
	}
}

func TestInvalidBoolConfig(t *testing.T) {
	ing := buildIngress()

	data := map[string]string{}
	data[parser.GetAnnotationWithPrefix(enableRewriteLogAnnotation)] = "blo"
	ing.SetAnnotations(data)

	log, err := NewParser(&resolver.Mock{}).Parse(ing)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	nginxLogs, ok := log.(*Config)
	if !ok {
		t.Errorf("expected a Config type")
	}

	if !nginxLogs.Access {
		t.Errorf("expected access log to be enabled due to invalid config, but it is disabled")
	}
}
