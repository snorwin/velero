/*
Copyright the Velero contributors.

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
	velerov1 "github.com/heptio/velero/pkg/apis/velero/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResticRepositories implements ResticRepositoryInterface
type FakeResticRepositories struct {
	Fake *FakeVeleroV1
	ns   string
}

var resticrepositoriesResource = schema.GroupVersionResource{Group: "velero.io", Version: "v1", Resource: "resticrepositories"}

var resticrepositoriesKind = schema.GroupVersionKind{Group: "velero.io", Version: "v1", Kind: "ResticRepository"}

// Get takes name of the resticRepository, and returns the corresponding resticRepository object, and an error if there is any.
func (c *FakeResticRepositories) Get(name string, options v1.GetOptions) (result *velerov1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resticrepositoriesResource, c.ns, name), &velerov1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.ResticRepository), err
}

// List takes label and field selectors, and returns the list of ResticRepositories that match those selectors.
func (c *FakeResticRepositories) List(opts v1.ListOptions) (result *velerov1.ResticRepositoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resticrepositoriesResource, resticrepositoriesKind, c.ns, opts), &velerov1.ResticRepositoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &velerov1.ResticRepositoryList{ListMeta: obj.(*velerov1.ResticRepositoryList).ListMeta}
	for _, item := range obj.(*velerov1.ResticRepositoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resticRepositories.
func (c *FakeResticRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resticrepositoriesResource, c.ns, opts))

}

// Create takes the representation of a resticRepository and creates it.  Returns the server's representation of the resticRepository, and an error, if there is any.
func (c *FakeResticRepositories) Create(resticRepository *velerov1.ResticRepository) (result *velerov1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resticrepositoriesResource, c.ns, resticRepository), &velerov1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.ResticRepository), err
}

// Update takes the representation of a resticRepository and updates it. Returns the server's representation of the resticRepository, and an error, if there is any.
func (c *FakeResticRepositories) Update(resticRepository *velerov1.ResticRepository) (result *velerov1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resticrepositoriesResource, c.ns, resticRepository), &velerov1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.ResticRepository), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResticRepositories) UpdateStatus(resticRepository *velerov1.ResticRepository) (*velerov1.ResticRepository, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resticrepositoriesResource, "status", c.ns, resticRepository), &velerov1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.ResticRepository), err
}

// Delete takes name of the resticRepository and deletes it. Returns an error if one occurs.
func (c *FakeResticRepositories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resticrepositoriesResource, c.ns, name), &velerov1.ResticRepository{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResticRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resticrepositoriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &velerov1.ResticRepositoryList{})
	return err
}

// Patch applies the patch and returns the patched resticRepository.
func (c *FakeResticRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *velerov1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resticrepositoriesResource, c.ns, name, pt, data, subresources...), &velerov1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.ResticRepository), err
}
