# TechsupportmanagementTechSupportStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | The name of the Techsupport bundle file. | [optional] 
**Reason** | Pointer to **string** | Reason for techsupport failure, if any. | [optional] 
**RelayReason** | Pointer to **string** | Reason for status relay failure, if any. | [optional] [readonly] 
**RelayStatus** | Pointer to **string** | Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. | [optional] [readonly] 
**RequestTs** | Pointer to [**time.Time**](time.Time.md) | The time at which the techsupport request was initiated. | [optional] 
**Status** | Pointer to **string** | Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. | [optional] 
**TechsupportDownloadUrl** | Pointer to **string** | The Url to download the techsupport file. | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](asset.ClusterMember.Relationship.md) |  | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**OriginResource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**TechSupportRequest** | Pointer to [**TechsupportmanagementTechSupportBundleRelationship**](techsupportmanagement.TechSupportBundle.Relationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementTechSupportStatusAllOf

`func NewTechsupportmanagementTechSupportStatusAllOf() *TechsupportmanagementTechSupportStatusAllOf`

NewTechsupportmanagementTechSupportStatusAllOf instantiates a new TechsupportmanagementTechSupportStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportStatusAllOfWithDefaults

`func NewTechsupportmanagementTechSupportStatusAllOfWithDefaults() *TechsupportmanagementTechSupportStatusAllOf`

NewTechsupportmanagementTechSupportStatusAllOfWithDefaults instantiates a new TechsupportmanagementTechSupportStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetReason

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelayReason

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetRelayReason() string`

GetRelayReason returns the RelayReason field if non-nil, zero value otherwise.

### GetRelayReasonOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetRelayReasonOk() (*string, bool)`

GetRelayReasonOk returns a tuple with the RelayReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayReason

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetRelayReason(v string)`

SetRelayReason sets RelayReason field to given value.

### HasRelayReason

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasRelayReason() bool`

HasRelayReason returns a boolean if a field has been set.

### GetRelayStatus

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetRelayStatus() string`

GetRelayStatus returns the RelayStatus field if non-nil, zero value otherwise.

### GetRelayStatusOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetRelayStatusOk() (*string, bool)`

GetRelayStatusOk returns a tuple with the RelayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayStatus

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetRelayStatus(v string)`

SetRelayStatus sets RelayStatus field to given value.

### HasRelayStatus

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasRelayStatus() bool`

HasRelayStatus returns a boolean if a field has been set.

### GetRequestTs

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetRequestTs() time.Time`

GetRequestTs returns the RequestTs field if non-nil, zero value otherwise.

### GetRequestTsOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetRequestTsOk() (*time.Time, bool)`

GetRequestTsOk returns a tuple with the RequestTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTs

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetRequestTs(v time.Time)`

SetRequestTs sets RequestTs field to given value.

### HasRequestTs

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasRequestTs() bool`

HasRequestTs returns a boolean if a field has been set.

### GetStatus

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetTechsupportDownloadUrl() string`

GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field if non-nil, zero value otherwise.

### GetTechsupportDownloadUrlOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetTechsupportDownloadUrlOk() (*string, bool)`

GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetTechsupportDownloadUrl(v string)`

SetTechsupportDownloadUrl sets TechsupportDownloadUrl field to given value.

### HasTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasTechsupportDownloadUrl() bool`

HasTechsupportDownloadUrl returns a boolean if a field has been set.

### GetClusterMember

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetOriginResource

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetOriginResource() MoBaseMoRelationship`

GetOriginResource returns the OriginResource field if non-nil, zero value otherwise.

### GetOriginResourceOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetOriginResourceOk() (*MoBaseMoRelationship, bool)`

GetOriginResourceOk returns a tuple with the OriginResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginResource

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetOriginResource(v MoBaseMoRelationship)`

SetOriginResource sets OriginResource field to given value.

### HasOriginResource

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasOriginResource() bool`

HasOriginResource returns a boolean if a field has been set.

### GetTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetTechSupportRequest() TechsupportmanagementTechSupportBundleRelationship`

GetTechSupportRequest returns the TechSupportRequest field if non-nil, zero value otherwise.

### GetTechSupportRequestOk

`func (o *TechsupportmanagementTechSupportStatusAllOf) GetTechSupportRequestOk() (*TechsupportmanagementTechSupportBundleRelationship, bool)`

GetTechSupportRequestOk returns a tuple with the TechSupportRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatusAllOf) SetTechSupportRequest(v TechsupportmanagementTechSupportBundleRelationship)`

SetTechSupportRequest sets TechSupportRequest field to given value.

### HasTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatusAllOf) HasTechSupportRequest() bool`

HasTechSupportRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


