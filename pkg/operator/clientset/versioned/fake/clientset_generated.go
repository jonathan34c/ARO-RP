// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"

	clientset "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned"
	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned/typed/aro.openshift.io/v1alpha1"
	fakearov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned/typed/aro.openshift.io/v1alpha1/fake"
	previewv1alpha1 "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned/typed/preview.aro.openshift.io/v1alpha1"
	fakepreviewv1alpha1 "github.com/Azure/ARO-RP/pkg/operator/clientset/versioned/typed/preview.aro.openshift.io/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// AroV1alpha1 retrieves the AroV1alpha1Client
func (c *Clientset) AroV1alpha1() arov1alpha1.AroV1alpha1Interface {
	return &fakearov1alpha1.FakeAroV1alpha1{Fake: &c.Fake}
}

// PreviewV1alpha1 retrieves the PreviewV1alpha1Client
func (c *Clientset) PreviewV1alpha1() previewv1alpha1.PreviewV1alpha1Interface {
	return &fakepreviewv1alpha1.FakePreviewV1alpha1{Fake: &c.Fake}
}
