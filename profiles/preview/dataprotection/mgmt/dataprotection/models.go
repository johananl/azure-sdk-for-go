// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package dataprotection

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/dataprotection/mgmt/2021-01-01/dataprotection"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AbsoluteMarker = original.AbsoluteMarker

const (
	AllBackup    AbsoluteMarker = original.AllBackup
	FirstOfDay   AbsoluteMarker = original.FirstOfDay
	FirstOfMonth AbsoluteMarker = original.FirstOfMonth
	FirstOfWeek  AbsoluteMarker = original.FirstOfWeek
	FirstOfYear  AbsoluteMarker = original.FirstOfYear
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type CurrentProtectionState = original.CurrentProtectionState

const (
	BackupSchedulesSuspended    CurrentProtectionState = original.BackupSchedulesSuspended
	ConfiguringProtection       CurrentProtectionState = original.ConfiguringProtection
	ConfiguringProtectionFailed CurrentProtectionState = original.ConfiguringProtectionFailed
	Invalid                     CurrentProtectionState = original.Invalid
	NotProtected                CurrentProtectionState = original.NotProtected
	ProtectionConfigured        CurrentProtectionState = original.ProtectionConfigured
	ProtectionError             CurrentProtectionState = original.ProtectionError
	ProtectionStopped           CurrentProtectionState = original.ProtectionStopped
	RetentionSchedulesSuspended CurrentProtectionState = original.RetentionSchedulesSuspended
	SoftDeleted                 CurrentProtectionState = original.SoftDeleted
	SoftDeleting                CurrentProtectionState = original.SoftDeleting
	UpdatingProtection          CurrentProtectionState = original.UpdatingProtection
)

type DataStoreTypes = original.DataStoreTypes

const (
	ArchiveStore     DataStoreTypes = original.ArchiveStore
	OperationalStore DataStoreTypes = original.OperationalStore
	VaultStore       DataStoreTypes = original.VaultStore
)

type DayOfWeek = original.DayOfWeek

const (
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
)

type FeatureSupportStatus = original.FeatureSupportStatus

const (
	FeatureSupportStatusAlphaPreview       FeatureSupportStatus = original.FeatureSupportStatusAlphaPreview
	FeatureSupportStatusGenerallyAvailable FeatureSupportStatus = original.FeatureSupportStatusGenerallyAvailable
	FeatureSupportStatusInvalid            FeatureSupportStatus = original.FeatureSupportStatusInvalid
	FeatureSupportStatusNotSupported       FeatureSupportStatus = original.FeatureSupportStatusNotSupported
	FeatureSupportStatusPrivatePreview     FeatureSupportStatus = original.FeatureSupportStatusPrivatePreview
	FeatureSupportStatusPublicPreview      FeatureSupportStatus = original.FeatureSupportStatusPublicPreview
)

type FeatureType = original.FeatureType

const (
	FeatureTypeDataSourceType FeatureType = original.FeatureTypeDataSourceType
	FeatureTypeInvalid        FeatureType = original.FeatureTypeInvalid
)

type Month = original.Month

const (
	April     Month = original.April
	August    Month = original.August
	December  Month = original.December
	February  Month = original.February
	January   Month = original.January
	July      Month = original.July
	June      Month = original.June
	March     Month = original.March
	May       Month = original.May
	November  Month = original.November
	October   Month = original.October
	September Month = original.September
)

type ObjectType = original.ObjectType

const (
	ObjectTypeAzureBackupDiscreteRecoveryPoint ObjectType = original.ObjectTypeAzureBackupDiscreteRecoveryPoint
	ObjectTypeAzureBackupRecoveryPoint         ObjectType = original.ObjectTypeAzureBackupRecoveryPoint
)

type ObjectTypeBasicAzureBackupRestoreRequest = original.ObjectTypeBasicAzureBackupRestoreRequest

const (
	ObjectTypeAzureBackupRecoveryPointBasedRestoreRequest ObjectTypeBasicAzureBackupRestoreRequest = original.ObjectTypeAzureBackupRecoveryPointBasedRestoreRequest
	ObjectTypeAzureBackupRecoveryTimeBasedRestoreRequest  ObjectTypeBasicAzureBackupRestoreRequest = original.ObjectTypeAzureBackupRecoveryTimeBasedRestoreRequest
	ObjectTypeAzureBackupRestoreRequest                   ObjectTypeBasicAzureBackupRestoreRequest = original.ObjectTypeAzureBackupRestoreRequest
	ObjectTypeAzureBackupRestoreWithRehydrationRequest    ObjectTypeBasicAzureBackupRestoreRequest = original.ObjectTypeAzureBackupRestoreWithRehydrationRequest
)

type ObjectTypeBasicBackupCriteria = original.ObjectTypeBasicBackupCriteria

const (
	ObjectTypeBackupCriteria              ObjectTypeBasicBackupCriteria = original.ObjectTypeBackupCriteria
	ObjectTypeScheduleBasedBackupCriteria ObjectTypeBasicBackupCriteria = original.ObjectTypeScheduleBasedBackupCriteria
)

type ObjectTypeBasicBackupParameters = original.ObjectTypeBasicBackupParameters

const (
	ObjectTypeAzureBackupParams ObjectTypeBasicBackupParameters = original.ObjectTypeAzureBackupParams
	ObjectTypeBackupParameters  ObjectTypeBasicBackupParameters = original.ObjectTypeBackupParameters
)

type ObjectTypeBasicBaseBackupPolicy = original.ObjectTypeBasicBaseBackupPolicy

const (
	ObjectTypeBackupPolicy     ObjectTypeBasicBaseBackupPolicy = original.ObjectTypeBackupPolicy
	ObjectTypeBaseBackupPolicy ObjectTypeBasicBaseBackupPolicy = original.ObjectTypeBaseBackupPolicy
)

type ObjectTypeBasicBasePolicyRule = original.ObjectTypeBasicBasePolicyRule

const (
	ObjectTypeAzureBackupRule    ObjectTypeBasicBasePolicyRule = original.ObjectTypeAzureBackupRule
	ObjectTypeAzureRetentionRule ObjectTypeBasicBasePolicyRule = original.ObjectTypeAzureRetentionRule
	ObjectTypeBasePolicyRule     ObjectTypeBasicBasePolicyRule = original.ObjectTypeBasePolicyRule
)

type ObjectTypeBasicCopyOption = original.ObjectTypeBasicCopyOption

const (
	ObjectTypeCopyOnExpiryOption  ObjectTypeBasicCopyOption = original.ObjectTypeCopyOnExpiryOption
	ObjectTypeCopyOption          ObjectTypeBasicCopyOption = original.ObjectTypeCopyOption
	ObjectTypeCustomCopyOption    ObjectTypeBasicCopyOption = original.ObjectTypeCustomCopyOption
	ObjectTypeImmediateCopyOption ObjectTypeBasicCopyOption = original.ObjectTypeImmediateCopyOption
)

type ObjectTypeBasicDataStoreParameters = original.ObjectTypeBasicDataStoreParameters

const (
	ObjectTypeAzureOperationalStoreParameters ObjectTypeBasicDataStoreParameters = original.ObjectTypeAzureOperationalStoreParameters
	ObjectTypeDataStoreParameters             ObjectTypeBasicDataStoreParameters = original.ObjectTypeDataStoreParameters
)

type ObjectTypeBasicDeleteOption = original.ObjectTypeBasicDeleteOption

const (
	ObjectTypeAbsoluteDeleteOption ObjectTypeBasicDeleteOption = original.ObjectTypeAbsoluteDeleteOption
	ObjectTypeDeleteOption         ObjectTypeBasicDeleteOption = original.ObjectTypeDeleteOption
)

type ObjectTypeBasicFeatureValidationRequestBase = original.ObjectTypeBasicFeatureValidationRequestBase

const (
	ObjectTypeFeatureValidationRequest     ObjectTypeBasicFeatureValidationRequestBase = original.ObjectTypeFeatureValidationRequest
	ObjectTypeFeatureValidationRequestBase ObjectTypeBasicFeatureValidationRequestBase = original.ObjectTypeFeatureValidationRequestBase
)

type ObjectTypeBasicFeatureValidationResponseBase = original.ObjectTypeBasicFeatureValidationResponseBase

const (
	ObjectTypeFeatureValidationResponse     ObjectTypeBasicFeatureValidationResponseBase = original.ObjectTypeFeatureValidationResponse
	ObjectTypeFeatureValidationResponseBase ObjectTypeBasicFeatureValidationResponseBase = original.ObjectTypeFeatureValidationResponseBase
)

type ObjectTypeBasicItemLevelRestoreCriteria = original.ObjectTypeBasicItemLevelRestoreCriteria

const (
	ObjectTypeItemLevelRestoreCriteria           ObjectTypeBasicItemLevelRestoreCriteria = original.ObjectTypeItemLevelRestoreCriteria
	ObjectTypeRangeBasedItemLevelRestoreCriteria ObjectTypeBasicItemLevelRestoreCriteria = original.ObjectTypeRangeBasedItemLevelRestoreCriteria
)

type ObjectTypeBasicRestoreTargetInfoBase = original.ObjectTypeBasicRestoreTargetInfoBase

const (
	ObjectTypeItemLevelRestoreTargetInfo ObjectTypeBasicRestoreTargetInfoBase = original.ObjectTypeItemLevelRestoreTargetInfo
	ObjectTypeRestoreFilesTargetInfo     ObjectTypeBasicRestoreTargetInfoBase = original.ObjectTypeRestoreFilesTargetInfo
	ObjectTypeRestoreTargetInfo          ObjectTypeBasicRestoreTargetInfoBase = original.ObjectTypeRestoreTargetInfo
	ObjectTypeRestoreTargetInfoBase      ObjectTypeBasicRestoreTargetInfoBase = original.ObjectTypeRestoreTargetInfoBase
)

type ObjectTypeBasicTriggerContext = original.ObjectTypeBasicTriggerContext

const (
	ObjectTypeAdhocBasedTriggerContext    ObjectTypeBasicTriggerContext = original.ObjectTypeAdhocBasedTriggerContext
	ObjectTypeScheduleBasedTriggerContext ObjectTypeBasicTriggerContext = original.ObjectTypeScheduleBasedTriggerContext
	ObjectTypeTriggerContext              ObjectTypeBasicTriggerContext = original.ObjectTypeTriggerContext
)

type ProvisioningState = original.ProvisioningState

const (
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
	Unknown      ProvisioningState = original.Unknown
	Updating     ProvisioningState = original.Updating
)

type RehydrationPriority = original.RehydrationPriority

const (
	RehydrationPriorityHigh     RehydrationPriority = original.RehydrationPriorityHigh
	RehydrationPriorityInvalid  RehydrationPriority = original.RehydrationPriorityInvalid
	RehydrationPriorityStandard RehydrationPriority = original.RehydrationPriorityStandard
)

type RehydrationStatus = original.RehydrationStatus

const (
	COMPLETED        RehydrationStatus = original.COMPLETED
	CREATEINPROGRESS RehydrationStatus = original.CREATEINPROGRESS
	DELETED          RehydrationStatus = original.DELETED
	DELETEINPROGRESS RehydrationStatus = original.DELETEINPROGRESS
	FAILED           RehydrationStatus = original.FAILED
)

type RestoreSourceDataStoreType = original.RestoreSourceDataStoreType

const (
	RestoreSourceDataStoreTypeArchiveStore     RestoreSourceDataStoreType = original.RestoreSourceDataStoreTypeArchiveStore
	RestoreSourceDataStoreTypeOperationalStore RestoreSourceDataStoreType = original.RestoreSourceDataStoreTypeOperationalStore
	RestoreSourceDataStoreTypeVaultStore       RestoreSourceDataStoreType = original.RestoreSourceDataStoreTypeVaultStore
)

type RestoreTargetLocationType = original.RestoreTargetLocationType

const (
	RestoreTargetLocationTypeAzureBlobs RestoreTargetLocationType = original.RestoreTargetLocationTypeAzureBlobs
	RestoreTargetLocationTypeAzureFiles RestoreTargetLocationType = original.RestoreTargetLocationTypeAzureFiles
	RestoreTargetLocationTypeInvalid    RestoreTargetLocationType = original.RestoreTargetLocationTypeInvalid
)

type SourceDataStoreType = original.SourceDataStoreType

const (
	SourceDataStoreTypeArchiveStore  SourceDataStoreType = original.SourceDataStoreTypeArchiveStore
	SourceDataStoreTypeSnapshotStore SourceDataStoreType = original.SourceDataStoreTypeSnapshotStore
	SourceDataStoreTypeVaultStore    SourceDataStoreType = original.SourceDataStoreTypeVaultStore
)

type Status = original.Status

const (
	StatusConfiguringProtection       Status = original.StatusConfiguringProtection
	StatusConfiguringProtectionFailed Status = original.StatusConfiguringProtectionFailed
	StatusProtectionConfigured        Status = original.StatusProtectionConfigured
	StatusProtectionStopped           Status = original.StatusProtectionStopped
	StatusSoftDeleted                 Status = original.StatusSoftDeleted
	StatusSoftDeleting                Status = original.StatusSoftDeleting
)

type StorageSettingStoreTypes = original.StorageSettingStoreTypes

const (
	StorageSettingStoreTypesArchiveStore  StorageSettingStoreTypes = original.StorageSettingStoreTypesArchiveStore
	StorageSettingStoreTypesSnapshotStore StorageSettingStoreTypes = original.StorageSettingStoreTypesSnapshotStore
	StorageSettingStoreTypesVaultStore    StorageSettingStoreTypes = original.StorageSettingStoreTypesVaultStore
)

type StorageSettingTypes = original.StorageSettingTypes

const (
	GeoRedundant     StorageSettingTypes = original.GeoRedundant
	LocallyRedundant StorageSettingTypes = original.LocallyRedundant
)

type WeekNumber = original.WeekNumber

const (
	First  WeekNumber = original.First
	Fourth WeekNumber = original.Fourth
	Last   WeekNumber = original.Last
	Second WeekNumber = original.Second
	Third  WeekNumber = original.Third
)

type AbsoluteDeleteOption = original.AbsoluteDeleteOption
type AdHocBackupRuleOptions = original.AdHocBackupRuleOptions
type AdhocBackupTriggerOption = original.AdhocBackupTriggerOption
type AdhocBasedTaggingCriteria = original.AdhocBasedTaggingCriteria
type AdhocBasedTriggerContext = original.AdhocBasedTriggerContext
type AzureBackupDiscreteRecoveryPoint = original.AzureBackupDiscreteRecoveryPoint
type AzureBackupFindRestorableTimeRangesRequest = original.AzureBackupFindRestorableTimeRangesRequest
type AzureBackupFindRestorableTimeRangesRequestResource = original.AzureBackupFindRestorableTimeRangesRequestResource
type AzureBackupFindRestorableTimeRangesResponse = original.AzureBackupFindRestorableTimeRangesResponse
type AzureBackupFindRestorableTimeRangesResponseResource = original.AzureBackupFindRestorableTimeRangesResponseResource
type AzureBackupJob = original.AzureBackupJob
type AzureBackupJobResource = original.AzureBackupJobResource
type AzureBackupJobResourceList = original.AzureBackupJobResourceList
type AzureBackupJobResourceListIterator = original.AzureBackupJobResourceListIterator
type AzureBackupJobResourceListPage = original.AzureBackupJobResourceListPage
type AzureBackupParams = original.AzureBackupParams
type AzureBackupRecoveryPoint = original.AzureBackupRecoveryPoint
type AzureBackupRecoveryPointBasedRestoreRequest = original.AzureBackupRecoveryPointBasedRestoreRequest
type AzureBackupRecoveryPointResource = original.AzureBackupRecoveryPointResource
type AzureBackupRecoveryPointResourceList = original.AzureBackupRecoveryPointResourceList
type AzureBackupRecoveryPointResourceListIterator = original.AzureBackupRecoveryPointResourceListIterator
type AzureBackupRecoveryPointResourceListPage = original.AzureBackupRecoveryPointResourceListPage
type AzureBackupRecoveryTimeBasedRestoreRequest = original.AzureBackupRecoveryTimeBasedRestoreRequest
type AzureBackupRehydrationRequest = original.AzureBackupRehydrationRequest
type AzureBackupRestoreRequest = original.AzureBackupRestoreRequest
type AzureBackupRestoreWithRehydrationRequest = original.AzureBackupRestoreWithRehydrationRequest
type AzureBackupRule = original.AzureBackupRule
type AzureOperationalStoreParameters = original.AzureOperationalStoreParameters
type AzureRetentionRule = original.AzureRetentionRule
type BackupCriteria = original.BackupCriteria
type BackupInstance = original.BackupInstance
type BackupInstanceResource = original.BackupInstanceResource
type BackupInstanceResourceList = original.BackupInstanceResourceList
type BackupInstanceResourceListIterator = original.BackupInstanceResourceListIterator
type BackupInstanceResourceListPage = original.BackupInstanceResourceListPage
type BackupInstancesAdhocBackupFuture = original.BackupInstancesAdhocBackupFuture
type BackupInstancesClient = original.BackupInstancesClient
type BackupInstancesCreateOrUpdateFuture = original.BackupInstancesCreateOrUpdateFuture
type BackupInstancesDeleteFuture = original.BackupInstancesDeleteFuture
type BackupInstancesTriggerRehydrateFuture = original.BackupInstancesTriggerRehydrateFuture
type BackupInstancesTriggerRestoreFuture = original.BackupInstancesTriggerRestoreFuture
type BackupInstancesValidateForBackupFuture = original.BackupInstancesValidateForBackupFuture
type BackupInstancesValidateRestoreFuture = original.BackupInstancesValidateRestoreFuture
type BackupParameters = original.BackupParameters
type BackupPoliciesClient = original.BackupPoliciesClient
type BackupPolicy = original.BackupPolicy
type BackupSchedule = original.BackupSchedule
type BackupVault = original.BackupVault
type BackupVaultResource = original.BackupVaultResource
type BackupVaultResourceList = original.BackupVaultResourceList
type BackupVaultResourceListIterator = original.BackupVaultResourceListIterator
type BackupVaultResourceListPage = original.BackupVaultResourceListPage
type BackupVaultsClient = original.BackupVaultsClient
type BackupVaultsCreateOrUpdateFuture = original.BackupVaultsCreateOrUpdateFuture
type BackupVaultsPatchFuture = original.BackupVaultsPatchFuture
type BaseBackupPolicy = original.BaseBackupPolicy
type BaseBackupPolicyResource = original.BaseBackupPolicyResource
type BaseBackupPolicyResourceList = original.BaseBackupPolicyResourceList
type BaseBackupPolicyResourceListIterator = original.BaseBackupPolicyResourceListIterator
type BaseBackupPolicyResourceListPage = original.BaseBackupPolicyResourceListPage
type BaseClient = original.BaseClient
type BasePolicyRule = original.BasePolicyRule
type BasicAzureBackupRecoveryPoint = original.BasicAzureBackupRecoveryPoint
type BasicAzureBackupRecoveryPointBasedRestoreRequest = original.BasicAzureBackupRecoveryPointBasedRestoreRequest
type BasicAzureBackupRestoreRequest = original.BasicAzureBackupRestoreRequest
type BasicBackupCriteria = original.BasicBackupCriteria
type BasicBackupParameters = original.BasicBackupParameters
type BasicBaseBackupPolicy = original.BasicBaseBackupPolicy
type BasicBasePolicyRule = original.BasicBasePolicyRule
type BasicCopyOption = original.BasicCopyOption
type BasicDataStoreParameters = original.BasicDataStoreParameters
type BasicDeleteOption = original.BasicDeleteOption
type BasicFeatureValidationRequestBase = original.BasicFeatureValidationRequestBase
type BasicFeatureValidationResponseBase = original.BasicFeatureValidationResponseBase
type BasicItemLevelRestoreCriteria = original.BasicItemLevelRestoreCriteria
type BasicRestoreTargetInfoBase = original.BasicRestoreTargetInfoBase
type BasicTriggerContext = original.BasicTriggerContext
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ClientDiscoveryDisplay = original.ClientDiscoveryDisplay
type ClientDiscoveryForLogSpecification = original.ClientDiscoveryForLogSpecification
type ClientDiscoveryForProperties = original.ClientDiscoveryForProperties
type ClientDiscoveryForServiceSpecification = original.ClientDiscoveryForServiceSpecification
type ClientDiscoveryResponse = original.ClientDiscoveryResponse
type ClientDiscoveryResponseIterator = original.ClientDiscoveryResponseIterator
type ClientDiscoveryResponsePage = original.ClientDiscoveryResponsePage
type ClientDiscoveryValueForSingleAPI = original.ClientDiscoveryValueForSingleAPI
type CloudError = original.CloudError
type CopyOnExpiryOption = original.CopyOnExpiryOption
type CopyOption = original.CopyOption
type CustomCopyOption = original.CustomCopyOption
type DataStoreInfoBase = original.DataStoreInfoBase
type DataStoreParameters = original.DataStoreParameters
type Datasource = original.Datasource
type DatasourceSet = original.DatasourceSet
type Day = original.Day
type DeleteOption = original.DeleteOption
type DppIdentityDetails = original.DppIdentityDetails
type DppResource = original.DppResource
type DppResourceList = original.DppResourceList
type DppTrackedResource = original.DppTrackedResource
type DppTrackedResourceList = original.DppTrackedResourceList
type DppWorkerRequest = original.DppWorkerRequest
type Error = original.Error
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ExportJobsClient = original.ExportJobsClient
type ExportJobsOperationResultClient = original.ExportJobsOperationResultClient
type ExportJobsResult = original.ExportJobsResult
type ExportJobsTriggerFuture = original.ExportJobsTriggerFuture
type FeatureValidationRequest = original.FeatureValidationRequest
type FeatureValidationRequestBase = original.FeatureValidationRequestBase
type FeatureValidationResponse = original.FeatureValidationResponse
type FeatureValidationResponseBase = original.FeatureValidationResponseBase
type FeatureValidationResponseBaseModel = original.FeatureValidationResponseBaseModel
type FindRestorableTimeRangesClient = original.FindRestorableTimeRangesClient
type ImmediateCopyOption = original.ImmediateCopyOption
type InnerError = original.InnerError
type ItemLevelRestoreCriteria = original.ItemLevelRestoreCriteria
type ItemLevelRestoreTargetInfo = original.ItemLevelRestoreTargetInfo
type JobClient = original.JobClient
type JobExtendedInfo = original.JobExtendedInfo
type JobSubTask = original.JobSubTask
type JobsClient = original.JobsClient
type OperationExtendedInfo = original.OperationExtendedInfo
type OperationJobExtendedInfo = original.OperationJobExtendedInfo
type OperationResource = original.OperationResource
type OperationResultClient = original.OperationResultClient
type OperationsClient = original.OperationsClient
type PatchResourceRequestInput = original.PatchResourceRequestInput
type PolicyInfo = original.PolicyInfo
type PolicyParameters = original.PolicyParameters
type ProtectionStatusDetails = original.ProtectionStatusDetails
type RangeBasedItemLevelRestoreCriteria = original.RangeBasedItemLevelRestoreCriteria
type RecoveryPointClient = original.RecoveryPointClient
type RecoveryPointDataStoreDetails = original.RecoveryPointDataStoreDetails
type RecoveryPointsClient = original.RecoveryPointsClient
type RecoveryPointsFilters = original.RecoveryPointsFilters
type RestorableTimeRange = original.RestorableTimeRange
type RestoreFilesTargetInfo = original.RestoreFilesTargetInfo
type RestoreJobRecoveryPointDetails = original.RestoreJobRecoveryPointDetails
type RestoreTargetInfo = original.RestoreTargetInfo
type RestoreTargetInfoBase = original.RestoreTargetInfoBase
type RetentionTag = original.RetentionTag
type ScheduleBasedBackupCriteria = original.ScheduleBasedBackupCriteria
type ScheduleBasedTriggerContext = original.ScheduleBasedTriggerContext
type SourceLifeCycle = original.SourceLifeCycle
type StorageSetting = original.StorageSetting
type SupportedFeature = original.SupportedFeature
type SystemData = original.SystemData
type TaggingCriteria = original.TaggingCriteria
type TargetCopySetting = original.TargetCopySetting
type TargetDetails = original.TargetDetails
type TriggerBackupRequest = original.TriggerBackupRequest
type TriggerContext = original.TriggerContext
type UserFacingError = original.UserFacingError
type ValidateForBackupRequest = original.ValidateForBackupRequest
type ValidateRestoreRequestObject = original.ValidateRestoreRequestObject

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAzureBackupJobResourceListIterator(page AzureBackupJobResourceListPage) AzureBackupJobResourceListIterator {
	return original.NewAzureBackupJobResourceListIterator(page)
}
func NewAzureBackupJobResourceListPage(cur AzureBackupJobResourceList, getNextPage func(context.Context, AzureBackupJobResourceList) (AzureBackupJobResourceList, error)) AzureBackupJobResourceListPage {
	return original.NewAzureBackupJobResourceListPage(cur, getNextPage)
}
func NewAzureBackupRecoveryPointResourceListIterator(page AzureBackupRecoveryPointResourceListPage) AzureBackupRecoveryPointResourceListIterator {
	return original.NewAzureBackupRecoveryPointResourceListIterator(page)
}
func NewAzureBackupRecoveryPointResourceListPage(cur AzureBackupRecoveryPointResourceList, getNextPage func(context.Context, AzureBackupRecoveryPointResourceList) (AzureBackupRecoveryPointResourceList, error)) AzureBackupRecoveryPointResourceListPage {
	return original.NewAzureBackupRecoveryPointResourceListPage(cur, getNextPage)
}
func NewBackupInstanceResourceListIterator(page BackupInstanceResourceListPage) BackupInstanceResourceListIterator {
	return original.NewBackupInstanceResourceListIterator(page)
}
func NewBackupInstanceResourceListPage(cur BackupInstanceResourceList, getNextPage func(context.Context, BackupInstanceResourceList) (BackupInstanceResourceList, error)) BackupInstanceResourceListPage {
	return original.NewBackupInstanceResourceListPage(cur, getNextPage)
}
func NewBackupInstancesClient(subscriptionID string) BackupInstancesClient {
	return original.NewBackupInstancesClient(subscriptionID)
}
func NewBackupInstancesClientWithBaseURI(baseURI string, subscriptionID string) BackupInstancesClient {
	return original.NewBackupInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupPoliciesClient(subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClient(subscriptionID)
}
func NewBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupVaultResourceListIterator(page BackupVaultResourceListPage) BackupVaultResourceListIterator {
	return original.NewBackupVaultResourceListIterator(page)
}
func NewBackupVaultResourceListPage(cur BackupVaultResourceList, getNextPage func(context.Context, BackupVaultResourceList) (BackupVaultResourceList, error)) BackupVaultResourceListPage {
	return original.NewBackupVaultResourceListPage(cur, getNextPage)
}
func NewBackupVaultsClient(subscriptionID string) BackupVaultsClient {
	return original.NewBackupVaultsClient(subscriptionID)
}
func NewBackupVaultsClientWithBaseURI(baseURI string, subscriptionID string) BackupVaultsClient {
	return original.NewBackupVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBaseBackupPolicyResourceListIterator(page BaseBackupPolicyResourceListPage) BaseBackupPolicyResourceListIterator {
	return original.NewBaseBackupPolicyResourceListIterator(page)
}
func NewBaseBackupPolicyResourceListPage(cur BaseBackupPolicyResourceList, getNextPage func(context.Context, BaseBackupPolicyResourceList) (BaseBackupPolicyResourceList, error)) BaseBackupPolicyResourceListPage {
	return original.NewBaseBackupPolicyResourceListPage(cur, getNextPage)
}
func NewClientDiscoveryResponseIterator(page ClientDiscoveryResponsePage) ClientDiscoveryResponseIterator {
	return original.NewClientDiscoveryResponseIterator(page)
}
func NewClientDiscoveryResponsePage(cur ClientDiscoveryResponse, getNextPage func(context.Context, ClientDiscoveryResponse) (ClientDiscoveryResponse, error)) ClientDiscoveryResponsePage {
	return original.NewClientDiscoveryResponsePage(cur, getNextPage)
}
func NewExportJobsClient(subscriptionID string) ExportJobsClient {
	return original.NewExportJobsClient(subscriptionID)
}
func NewExportJobsClientWithBaseURI(baseURI string, subscriptionID string) ExportJobsClient {
	return original.NewExportJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExportJobsOperationResultClient(subscriptionID string) ExportJobsOperationResultClient {
	return original.NewExportJobsOperationResultClient(subscriptionID)
}
func NewExportJobsOperationResultClientWithBaseURI(baseURI string, subscriptionID string) ExportJobsOperationResultClient {
	return original.NewExportJobsOperationResultClientWithBaseURI(baseURI, subscriptionID)
}
func NewFindRestorableTimeRangesClient(subscriptionID string) FindRestorableTimeRangesClient {
	return original.NewFindRestorableTimeRangesClient(subscriptionID)
}
func NewFindRestorableTimeRangesClientWithBaseURI(baseURI string, subscriptionID string) FindRestorableTimeRangesClient {
	return original.NewFindRestorableTimeRangesClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobClient(subscriptionID string) JobClient {
	return original.NewJobClient(subscriptionID)
}
func NewJobClientWithBaseURI(baseURI string, subscriptionID string) JobClient {
	return original.NewJobClientWithBaseURI(baseURI, subscriptionID)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationResultClient(subscriptionID string) OperationResultClient {
	return original.NewOperationResultClient(subscriptionID)
}
func NewOperationResultClientWithBaseURI(baseURI string, subscriptionID string) OperationResultClient {
	return original.NewOperationResultClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecoveryPointClient(subscriptionID string) RecoveryPointClient {
	return original.NewRecoveryPointClient(subscriptionID)
}
func NewRecoveryPointClientWithBaseURI(baseURI string, subscriptionID string) RecoveryPointClient {
	return original.NewRecoveryPointClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecoveryPointsClient(subscriptionID string) RecoveryPointsClient {
	return original.NewRecoveryPointsClient(subscriptionID)
}
func NewRecoveryPointsClientWithBaseURI(baseURI string, subscriptionID string) RecoveryPointsClient {
	return original.NewRecoveryPointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAbsoluteMarkerValues() []AbsoluteMarker {
	return original.PossibleAbsoluteMarkerValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleCurrentProtectionStateValues() []CurrentProtectionState {
	return original.PossibleCurrentProtectionStateValues()
}
func PossibleDataStoreTypesValues() []DataStoreTypes {
	return original.PossibleDataStoreTypesValues()
}
func PossibleDayOfWeekValues() []DayOfWeek {
	return original.PossibleDayOfWeekValues()
}
func PossibleFeatureSupportStatusValues() []FeatureSupportStatus {
	return original.PossibleFeatureSupportStatusValues()
}
func PossibleFeatureTypeValues() []FeatureType {
	return original.PossibleFeatureTypeValues()
}
func PossibleMonthValues() []Month {
	return original.PossibleMonthValues()
}
func PossibleObjectTypeBasicAzureBackupRestoreRequestValues() []ObjectTypeBasicAzureBackupRestoreRequest {
	return original.PossibleObjectTypeBasicAzureBackupRestoreRequestValues()
}
func PossibleObjectTypeBasicBackupCriteriaValues() []ObjectTypeBasicBackupCriteria {
	return original.PossibleObjectTypeBasicBackupCriteriaValues()
}
func PossibleObjectTypeBasicBackupParametersValues() []ObjectTypeBasicBackupParameters {
	return original.PossibleObjectTypeBasicBackupParametersValues()
}
func PossibleObjectTypeBasicBaseBackupPolicyValues() []ObjectTypeBasicBaseBackupPolicy {
	return original.PossibleObjectTypeBasicBaseBackupPolicyValues()
}
func PossibleObjectTypeBasicBasePolicyRuleValues() []ObjectTypeBasicBasePolicyRule {
	return original.PossibleObjectTypeBasicBasePolicyRuleValues()
}
func PossibleObjectTypeBasicCopyOptionValues() []ObjectTypeBasicCopyOption {
	return original.PossibleObjectTypeBasicCopyOptionValues()
}
func PossibleObjectTypeBasicDataStoreParametersValues() []ObjectTypeBasicDataStoreParameters {
	return original.PossibleObjectTypeBasicDataStoreParametersValues()
}
func PossibleObjectTypeBasicDeleteOptionValues() []ObjectTypeBasicDeleteOption {
	return original.PossibleObjectTypeBasicDeleteOptionValues()
}
func PossibleObjectTypeBasicFeatureValidationRequestBaseValues() []ObjectTypeBasicFeatureValidationRequestBase {
	return original.PossibleObjectTypeBasicFeatureValidationRequestBaseValues()
}
func PossibleObjectTypeBasicFeatureValidationResponseBaseValues() []ObjectTypeBasicFeatureValidationResponseBase {
	return original.PossibleObjectTypeBasicFeatureValidationResponseBaseValues()
}
func PossibleObjectTypeBasicItemLevelRestoreCriteriaValues() []ObjectTypeBasicItemLevelRestoreCriteria {
	return original.PossibleObjectTypeBasicItemLevelRestoreCriteriaValues()
}
func PossibleObjectTypeBasicRestoreTargetInfoBaseValues() []ObjectTypeBasicRestoreTargetInfoBase {
	return original.PossibleObjectTypeBasicRestoreTargetInfoBaseValues()
}
func PossibleObjectTypeBasicTriggerContextValues() []ObjectTypeBasicTriggerContext {
	return original.PossibleObjectTypeBasicTriggerContextValues()
}
func PossibleObjectTypeValues() []ObjectType {
	return original.PossibleObjectTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRehydrationPriorityValues() []RehydrationPriority {
	return original.PossibleRehydrationPriorityValues()
}
func PossibleRehydrationStatusValues() []RehydrationStatus {
	return original.PossibleRehydrationStatusValues()
}
func PossibleRestoreSourceDataStoreTypeValues() []RestoreSourceDataStoreType {
	return original.PossibleRestoreSourceDataStoreTypeValues()
}
func PossibleRestoreTargetLocationTypeValues() []RestoreTargetLocationType {
	return original.PossibleRestoreTargetLocationTypeValues()
}
func PossibleSourceDataStoreTypeValues() []SourceDataStoreType {
	return original.PossibleSourceDataStoreTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleStorageSettingStoreTypesValues() []StorageSettingStoreTypes {
	return original.PossibleStorageSettingStoreTypesValues()
}
func PossibleStorageSettingTypesValues() []StorageSettingTypes {
	return original.PossibleStorageSettingTypesValues()
}
func PossibleWeekNumberValues() []WeekNumber {
	return original.PossibleWeekNumberValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}