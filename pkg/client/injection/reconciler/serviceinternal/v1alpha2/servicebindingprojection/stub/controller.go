/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by injection-gen. DO NOT EDIT.

package servicebindingprojection

import (
	context "context"

	servicebindingprojection "github.com/vmware-labs/service-bindings/pkg/client/injection/informers/serviceinternal/v1alpha2/servicebindingprojection"
	v1alpha2servicebindingprojection "github.com/vmware-labs/service-bindings/pkg/client/injection/reconciler/serviceinternal/v1alpha2/servicebindingprojection"
	configmap "knative.dev/pkg/configmap"
	controller "knative.dev/pkg/controller"
	logging "knative.dev/pkg/logging"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// NewController creates a Reconciler for ServiceBindingProjection and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	servicebindingprojectionInformer := servicebindingprojection.Get(ctx)

	// TODO: setup additional informers here.

	r := &Reconciler{}
	impl := v1alpha2servicebindingprojection.NewImpl(ctx, r)

	logger.Info("Setting up event handlers.")

	servicebindingprojectionInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	// TODO: add additional informer event handlers here.

	return impl
}
