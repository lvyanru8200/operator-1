/*


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

package fake

import (
	"context"

	v1beta1 "github.com/VictoriaMetrics/operator/api/victoriametrics/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVMAlertmanagerConfigs implements VMAlertmanagerConfigInterface
type FakeVMAlertmanagerConfigs struct {
	Fake *FakeVictoriametricsV1beta1
	ns   string
}

var vmalertmanagerconfigsResource = schema.GroupVersionResource{Group: "victoriametrics", Version: "v1beta1", Resource: "vmalertmanagerconfigs"}

var vmalertmanagerconfigsKind = schema.GroupVersionKind{Group: "victoriametrics", Version: "v1beta1", Kind: "VMAlertmanagerConfig"}

// Get takes name of the vMAlertmanagerConfig, and returns the corresponding vMAlertmanagerConfig object, and an error if there is any.
func (c *FakeVMAlertmanagerConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMAlertmanagerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vmalertmanagerconfigsResource, c.ns, name), &v1beta1.VMAlertmanagerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAlertmanagerConfig), err
}

// List takes label and field selectors, and returns the list of VMAlertmanagerConfigs that match those selectors.
func (c *FakeVMAlertmanagerConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMAlertmanagerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vmalertmanagerconfigsResource, vmalertmanagerconfigsKind, c.ns, opts), &v1beta1.VMAlertmanagerConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VMAlertmanagerConfigList{ListMeta: obj.(*v1beta1.VMAlertmanagerConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.VMAlertmanagerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vMAlertmanagerConfigs.
func (c *FakeVMAlertmanagerConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vmalertmanagerconfigsResource, c.ns, opts))

}

// Create takes the representation of a vMAlertmanagerConfig and creates it.  Returns the server's representation of the vMAlertmanagerConfig, and an error, if there is any.
func (c *FakeVMAlertmanagerConfigs) Create(ctx context.Context, vMAlertmanagerConfig *v1beta1.VMAlertmanagerConfig, opts v1.CreateOptions) (result *v1beta1.VMAlertmanagerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vmalertmanagerconfigsResource, c.ns, vMAlertmanagerConfig), &v1beta1.VMAlertmanagerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAlertmanagerConfig), err
}

// Update takes the representation of a vMAlertmanagerConfig and updates it. Returns the server's representation of the vMAlertmanagerConfig, and an error, if there is any.
func (c *FakeVMAlertmanagerConfigs) Update(ctx context.Context, vMAlertmanagerConfig *v1beta1.VMAlertmanagerConfig, opts v1.UpdateOptions) (result *v1beta1.VMAlertmanagerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vmalertmanagerconfigsResource, c.ns, vMAlertmanagerConfig), &v1beta1.VMAlertmanagerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAlertmanagerConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVMAlertmanagerConfigs) UpdateStatus(ctx context.Context, vMAlertmanagerConfig *v1beta1.VMAlertmanagerConfig, opts v1.UpdateOptions) (*v1beta1.VMAlertmanagerConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vmalertmanagerconfigsResource, "status", c.ns, vMAlertmanagerConfig), &v1beta1.VMAlertmanagerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAlertmanagerConfig), err
}

// Delete takes name of the vMAlertmanagerConfig and deletes it. Returns an error if one occurs.
func (c *FakeVMAlertmanagerConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vmalertmanagerconfigsResource, c.ns, name, opts), &v1beta1.VMAlertmanagerConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVMAlertmanagerConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vmalertmanagerconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VMAlertmanagerConfigList{})
	return err
}

// Patch applies the patch and returns the patched vMAlertmanagerConfig.
func (c *FakeVMAlertmanagerConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAlertmanagerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vmalertmanagerconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.VMAlertmanagerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAlertmanagerConfig), err
}
