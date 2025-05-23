/*
Copyright 2020 The Flux authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
	scheme "github.com/fluxcd/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// MetricTemplatesGetter has a method to return a MetricTemplateInterface.
// A group's client should implement this interface.
type MetricTemplatesGetter interface {
	MetricTemplates(namespace string) MetricTemplateInterface
}

// MetricTemplateInterface has methods to work with MetricTemplate resources.
type MetricTemplateInterface interface {
	Create(ctx context.Context, metricTemplate *v1beta1.MetricTemplate, opts v1.CreateOptions) (*v1beta1.MetricTemplate, error)
	Update(ctx context.Context, metricTemplate *v1beta1.MetricTemplate, opts v1.UpdateOptions) (*v1beta1.MetricTemplate, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, metricTemplate *v1beta1.MetricTemplate, opts v1.UpdateOptions) (*v1beta1.MetricTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.MetricTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.MetricTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MetricTemplate, err error)
	MetricTemplateExpansion
}

// metricTemplates implements MetricTemplateInterface
type metricTemplates struct {
	*gentype.ClientWithList[*v1beta1.MetricTemplate, *v1beta1.MetricTemplateList]
}

// newMetricTemplates returns a MetricTemplates
func newMetricTemplates(c *FlaggerV1beta1Client, namespace string) *metricTemplates {
	return &metricTemplates{
		gentype.NewClientWithList[*v1beta1.MetricTemplate, *v1beta1.MetricTemplateList](
			"metrictemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.MetricTemplate { return &v1beta1.MetricTemplate{} },
			func() *v1beta1.MetricTemplateList { return &v1beta1.MetricTemplateList{} }),
	}
}
