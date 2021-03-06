/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package capabilities

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

const (
	applicationTypeName = "Application"

	// ApplicationV1_1 is the capabilties string for standard new non-backwards compatible fabric v1.1 application capabilities.
	ApplicationV1_1 = "V1_1"

	// ApplicationV1_2 is the capabilties string for standard new non-backwards compatible fabric v1.2 application capabilities.
	ApplicationV1_2 = "V1_2"

	// ApplicationPvtDataExperimental is the capabilties string for private data using the experimental feature of collections/sideDB.
	ApplicationPvtDataExperimental = "V1_1_PVTDATA_EXPERIMENTAL"

	// ApplicationResourcesTreeExperimental is the capabilties string for private data using the experimental feature of collections/sideDB.
	ApplicationResourcesTreeExperimental = "V1_1_RESOURCETREE_EXPERIMENTAL"

	// ApplicationChaincodeLifecycleExperimental is the capabilties string for improved chaincode lifecycle
	// which is targetted to be a real feature for v1.2, but which is being included only as experimental
	// during development.  This string will hopefully be removed prior to the release of v1.2 and
	// Will be enabled along with ApplicationV1_2
	ApplicationChaincodeLifecycleExperimental = "V1_2_CHAINCODE_LIFECYCLE_EXPERIMENTAL"
)

// ApplicationProvider provides capabilities information for application level config.
type ApplicationProvider struct {
	*registry
	v11                          bool
	v12                          bool
	v11PvtDataExperimental       bool
	v11ResourcesTreeExperimental bool
	v12LifecycleExperimental     bool
}

// NewApplicationProvider creates a application capabilities provider.
func NewApplicationProvider(capabilities map[string]*cb.Capability) *ApplicationProvider {
	ap := &ApplicationProvider{}
	ap.registry = newRegistry(ap, capabilities)
	_, ap.v11 = capabilities[ApplicationV1_1]
	_, ap.v12 = capabilities[ApplicationV1_2]
	_, ap.v11PvtDataExperimental = capabilities[ApplicationPvtDataExperimental]
	_, ap.v11ResourcesTreeExperimental = capabilities[ApplicationResourcesTreeExperimental]
	_, ap.v12LifecycleExperimental = capabilities[ApplicationChaincodeLifecycleExperimental]
	return ap
}

// Type returns a descriptive string for logging purposes.
func (ap *ApplicationProvider) Type() string {
	return applicationTypeName
}

// ResourcesTree returns whether the experimental resources tree transaction processing should be enabled.
func (ap *ApplicationProvider) ResourcesTree() bool {
	return ap.v11ResourcesTreeExperimental
}

// ForbidDuplicateTXIdInBlock specifies whether two transactions with the same TXId are permitted
// in the same block or whether we mark the second one as TxValidationCode_DUPLICATE_TXID
func (ap *ApplicationProvider) ForbidDuplicateTXIdInBlock() bool {
	return ap.v11 || ap.v12
}

// PrivateChannelData returns true if support for private channel data (a.k.a. collections) is enabled.
func (ap *ApplicationProvider) PrivateChannelData() bool {
	return ap.v11PvtDataExperimental
}

// V1_1Validation returns true is this channel is configured to perform stricter validation
// of transactions (as introduced in v1.1).
func (ap *ApplicationProvider) V1_1Validation() bool {
	return ap.v11 || ap.v12
}

// V1_2Validation returns true if this channel is configured to perform stricter validation
// of transactions (as introduced in v1.2).
func (ap *ApplicationProvider) V1_2Validation() bool {
	return ap.v12
}

// MetadataLifecycle indicates whether the peer should use the deprecated and problematic
// v1.0/v1.1 lifecycle, or whether it should use the newer per channel peer local chaincode
// metadata package approach planned for release with Fabric v1.2
func (ap *ApplicationProvider) MetadataLifecycle() bool {
	return ap.v12LifecycleExperimental
}
