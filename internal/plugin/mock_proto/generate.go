// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:generate go tool go.uber.org/mock/mockgen -destination mock.go github.com/hashicorp/terraform/internal/tfplugin5 ProviderClient,ProvisionerClient,Provisioner_ProvisionResourceClient,Provisioner_ProvisionResourceServer,Provider_InvokeActionClient

package mock_tfplugin5
