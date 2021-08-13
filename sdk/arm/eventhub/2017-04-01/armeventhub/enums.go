// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{
		AccessRightsListen,
		AccessRightsManage,
		AccessRightsSend,
	}
}

func (c AccessRights) ToPtr() *AccessRights {
	return &c
}

// DefaultAction - Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

func (c DefaultAction) ToPtr() *DefaultAction {
	return &c
}

// EncodingCaptureDescription - Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in New API Version
type EncodingCaptureDescription string

const (
	EncodingCaptureDescriptionAvro        EncodingCaptureDescription = "Avro"
	EncodingCaptureDescriptionAvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{
		EncodingCaptureDescriptionAvro,
		EncodingCaptureDescriptionAvroDeflate,
	}
}

func (c EncodingCaptureDescription) ToPtr() *EncodingCaptureDescription {
	return &c
}

// EntityStatus - Enumerates the possible values for the status of the Event Hub.
type EntityStatus string

const (
	EntityStatusActive          EntityStatus = "Active"
	EntityStatusDisabled        EntityStatus = "Disabled"
	EntityStatusRestoring       EntityStatus = "Restoring"
	EntityStatusSendDisabled    EntityStatus = "SendDisabled"
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	EntityStatusCreating        EntityStatus = "Creating"
	EntityStatusDeleting        EntityStatus = "Deleting"
	EntityStatusRenaming        EntityStatus = "Renaming"
	EntityStatusUnknown         EntityStatus = "Unknown"
)

func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{
		EntityStatusActive,
		EntityStatusDisabled,
		EntityStatusRestoring,
		EntityStatusSendDisabled,
		EntityStatusReceiveDisabled,
		EntityStatusCreating,
		EntityStatusDeleting,
		EntityStatusRenaming,
		EntityStatusUnknown,
	}
}

func (c EntityStatus) ToPtr() *EntityStatus {
	return &c
}

// KeyType - The access key to regenerate.
type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimaryKey,
		KeyTypeSecondaryKey,
	}
}

func (c KeyType) ToPtr() *KeyType {
	return &c
}

// ProvisioningStateDR - Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
type ProvisioningStateDR string

const (
	ProvisioningStateDRAccepted  ProvisioningStateDR = "Accepted"
	ProvisioningStateDRSucceeded ProvisioningStateDR = "Succeeded"
	ProvisioningStateDRFailed    ProvisioningStateDR = "Failed"
)

func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{
		ProvisioningStateDRAccepted,
		ProvisioningStateDRSucceeded,
		ProvisioningStateDRFailed,
	}
}

func (c ProvisioningStateDR) ToPtr() *ProvisioningStateDR {
	return &c
}

// RoleDisasterRecovery - role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
type RoleDisasterRecovery string

const (
	RoleDisasterRecoveryPrimary               RoleDisasterRecovery = "Primary"
	RoleDisasterRecoveryPrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	RoleDisasterRecoverySecondary             RoleDisasterRecovery = "Secondary"
)

func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{
		RoleDisasterRecoveryPrimary,
		RoleDisasterRecoveryPrimaryNotReplicating,
		RoleDisasterRecoverySecondary,
	}
}

func (c RoleDisasterRecovery) ToPtr() *RoleDisasterRecovery {
	return &c
}

// SkuName - Name of this SKU.
type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNameStandard SkuName = "Standard"
)

func PossibleSkuNameValues() []SkuName {
	return []SkuName{
		SkuNameBasic,
		SkuNameStandard,
	}
}

func (c SkuName) ToPtr() *SkuName {
	return &c
}

// SkuTier - The billing tier of this particular SKU.
type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{
		SkuTierBasic,
		SkuTierStandard,
	}
}

func (c SkuTier) ToPtr() *SkuTier {
	return &c
}

// UnavailableReason - Specifies the reason for the unavailability of the service.
type UnavailableReason string

const (
	UnavailableReasonNone                                  UnavailableReason = "None"
	UnavailableReasonInvalidName                           UnavailableReason = "InvalidName"
	UnavailableReasonSubscriptionIsDisabled                UnavailableReason = "SubscriptionIsDisabled"
	UnavailableReasonNameInUse                             UnavailableReason = "NameInUse"
	UnavailableReasonNameInLockdown                        UnavailableReason = "NameInLockdown"
	UnavailableReasonTooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{
		UnavailableReasonNone,
		UnavailableReasonInvalidName,
		UnavailableReasonSubscriptionIsDisabled,
		UnavailableReasonNameInUse,
		UnavailableReasonNameInLockdown,
		UnavailableReasonTooManyNamespaceInCurrentSubscription,
	}
}

func (c UnavailableReason) ToPtr() *UnavailableReason {
	return &c
}