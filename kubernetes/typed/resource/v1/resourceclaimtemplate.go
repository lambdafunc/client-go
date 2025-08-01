/*
Copyright The Kubernetes Authors.

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

package v1

import (
	context "context"

	resourcev1 "k8s.io/api/resource/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsresourcev1 "k8s.io/client-go/applyconfigurations/resource/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ResourceClaimTemplatesGetter has a method to return a ResourceClaimTemplateInterface.
// A group's client should implement this interface.
type ResourceClaimTemplatesGetter interface {
	ResourceClaimTemplates(namespace string) ResourceClaimTemplateInterface
}

// ResourceClaimTemplateInterface has methods to work with ResourceClaimTemplate resources.
type ResourceClaimTemplateInterface interface {
	Create(ctx context.Context, resourceClaimTemplate *resourcev1.ResourceClaimTemplate, opts metav1.CreateOptions) (*resourcev1.ResourceClaimTemplate, error)
	Update(ctx context.Context, resourceClaimTemplate *resourcev1.ResourceClaimTemplate, opts metav1.UpdateOptions) (*resourcev1.ResourceClaimTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*resourcev1.ResourceClaimTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*resourcev1.ResourceClaimTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *resourcev1.ResourceClaimTemplate, err error)
	Apply(ctx context.Context, resourceClaimTemplate *applyconfigurationsresourcev1.ResourceClaimTemplateApplyConfiguration, opts metav1.ApplyOptions) (result *resourcev1.ResourceClaimTemplate, err error)
	ResourceClaimTemplateExpansion
}

// resourceClaimTemplates implements ResourceClaimTemplateInterface
type resourceClaimTemplates struct {
	*gentype.ClientWithListAndApply[*resourcev1.ResourceClaimTemplate, *resourcev1.ResourceClaimTemplateList, *applyconfigurationsresourcev1.ResourceClaimTemplateApplyConfiguration]
}

// newResourceClaimTemplates returns a ResourceClaimTemplates
func newResourceClaimTemplates(c *ResourceV1Client, namespace string) *resourceClaimTemplates {
	return &resourceClaimTemplates{
		gentype.NewClientWithListAndApply[*resourcev1.ResourceClaimTemplate, *resourcev1.ResourceClaimTemplateList, *applyconfigurationsresourcev1.ResourceClaimTemplateApplyConfiguration](
			"resourceclaimtemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *resourcev1.ResourceClaimTemplate { return &resourcev1.ResourceClaimTemplate{} },
			func() *resourcev1.ResourceClaimTemplateList { return &resourcev1.ResourceClaimTemplateList{} },
			gentype.PrefersProtobuf[*resourcev1.ResourceClaimTemplate](),
		),
	}
}
