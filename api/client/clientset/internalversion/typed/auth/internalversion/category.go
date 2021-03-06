/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	auth "tkestack.io/tke/api/auth"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
)

// CategoriesGetter has a method to return a CategoryInterface.
// A group's client should implement this interface.
type CategoriesGetter interface {
	Categories() CategoryInterface
}

// CategoryInterface has methods to work with Category resources.
type CategoryInterface interface {
	Create(*auth.Category) (*auth.Category, error)
	Update(*auth.Category) (*auth.Category, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*auth.Category, error)
	List(opts v1.ListOptions) (*auth.CategoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.Category, err error)
	CategoryExpansion
}

// categories implements CategoryInterface
type categories struct {
	client rest.Interface
}

// newCategories returns a Categories
func newCategories(c *AuthClient) *categories {
	return &categories{
		client: c.RESTClient(),
	}
}

// Get takes name of the category, and returns the corresponding category object, and an error if there is any.
func (c *categories) Get(name string, options v1.GetOptions) (result *auth.Category, err error) {
	result = &auth.Category{}
	err = c.client.Get().
		Resource("categories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Categories that match those selectors.
func (c *categories) List(opts v1.ListOptions) (result *auth.CategoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &auth.CategoryList{}
	err = c.client.Get().
		Resource("categories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested categories.
func (c *categories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("categories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a category and creates it.  Returns the server's representation of the category, and an error, if there is any.
func (c *categories) Create(category *auth.Category) (result *auth.Category, err error) {
	result = &auth.Category{}
	err = c.client.Post().
		Resource("categories").
		Body(category).
		Do().
		Into(result)
	return
}

// Update takes the representation of a category and updates it. Returns the server's representation of the category, and an error, if there is any.
func (c *categories) Update(category *auth.Category) (result *auth.Category, err error) {
	result = &auth.Category{}
	err = c.client.Put().
		Resource("categories").
		Name(category.Name).
		Body(category).
		Do().
		Into(result)
	return
}

// Delete takes name of the category and deletes it. Returns an error if one occurs.
func (c *categories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("categories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *categories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("categories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched category.
func (c *categories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.Category, err error) {
	result = &auth.Category{}
	err = c.client.Patch(pt).
		Resource("categories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
