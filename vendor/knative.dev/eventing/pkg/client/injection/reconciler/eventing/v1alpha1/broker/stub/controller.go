/*
Copyright 2020 The Knative Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package broker

import (
	context "context"

	cache "k8s.io/client-go/tools/cache"
	broker "knative.dev/eventing/pkg/client/injection/informers/eventing/v1alpha1/broker"
	v1alpha1broker "knative.dev/eventing/pkg/client/injection/reconciler/eventing/v1alpha1/broker"
	configmap "knative.dev/pkg/configmap"
	controller "knative.dev/pkg/controller"
	logging "knative.dev/pkg/logging"
	reconciler "knative.dev/pkg/reconciler"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// NewController creates a Reconciler for Broker and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	brokerInformer := broker.Get(ctx)

	classValue := "default" // TODO: update this to the appropriate value.
	classFilter := reconciler.AnnotationFilterFunc(v1alpha1broker.ClassAnnotationKey, classValue, false /*allowUnset*/)

	// TODO: setup additional informers here.
	// TODO: remember to use the classFilter from above to filter appropriately.

	r := &Reconciler{}
	impl := v1alpha1broker.NewImpl(ctx, r, classValue)

	logger.Info("Setting up event handlers.")

	brokerInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: classFilter,
		Handler:    controller.HandleAll(impl.Enqueue),
	})

	// TODO: add additional informer event handlers here.

	return impl
}
