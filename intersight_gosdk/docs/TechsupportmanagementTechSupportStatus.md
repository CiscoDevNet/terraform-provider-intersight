# TechsupportmanagementTechSupportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "techsupportmanagement.TechSupportStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "techsupportmanagement.TechSupportStatus"]
**FileName** | Pointer to **string** | The name of the Techsupport bundle file. | [optional] 
**Reason** | Pointer to **string** | Reason for techsupport failure, if any. | [optional] 
**RelayReason** | Pointer to **string** | Reason for status relay failure, if any. | [optional] [readonly] 
**RelayStatus** | Pointer to **string** | Status of techsupport status relay. Valid values are NoRelay, Pending, Completed, and Failed. | [optional] [readonly] 
**RequestTs** | Pointer to **time.Time** | The time at which the techsupport request was initiated. | [optional] 
**Status** | Pointer to **string** | Status of techsupport collection. Valid values are Pending, CollectionInProgress, CollectionFailed, CollectionComplete, UploadPending, UploadInProgress, UploadPartsComplete, UploadFailed and Completed. The final status will be either CollectionFailed or UploadFailed if there is a failure and Completed if the request completed successfully and the file was uploaded to Intersight Storage Service. All the remaining status values indicates the progress of techsupport collection. | [optional] 
**TechsupportDownloadUrl** | Pointer to **string** | The Url to download the techsupport file. | [optional] 
**ClusterMember** | Pointer to [**AssetClusterMemberRelationship**](AssetClusterMemberRelationship.md) |  | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**OriginResource** | Pointer to [**MoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**TechSupportRequest** | Pointer to [**TechsupportmanagementTechSupportBundleRelationship**](TechsupportmanagementTechSupportBundleRelationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementTechSupportStatus

`func NewTechsupportmanagementTechSupportStatus(classId string, objectType string, ) *TechsupportmanagementTechSupportStatus`

NewTechsupportmanagementTechSupportStatus instantiates a new TechsupportmanagementTechSupportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportStatusWithDefaults

`func NewTechsupportmanagementTechSupportStatusWithDefaults() *TechsupportmanagementTechSupportStatus`

NewTechsupportmanagementTechSupportStatusWithDefaults instantiates a new TechsupportmanagementTechSupportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TechsupportmanagementTechSupportStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TechsupportmanagementTechSupportStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TechsupportmanagementTechSupportStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TechsupportmanagementTechSupportStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TechsupportmanagementTechSupportStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TechsupportmanagementTechSupportStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFileName

`func (o *TechsupportmanagementTechSupportStatus) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *TechsupportmanagementTechSupportStatus) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *TechsupportmanagementTechSupportStatus) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *TechsupportmanagementTechSupportStatus) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetReason

`func (o *TechsupportmanagementTechSupportStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TechsupportmanagementTechSupportStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TechsupportmanagementTechSupportStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TechsupportmanagementTechSupportStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRelayReason

`func (o *TechsupportmanagementTechSupportStatus) GetRelayReason() string`

GetRelayReason returns the RelayReason field if non-nil, zero value otherwise.

### GetRelayReasonOk

`func (o *TechsupportmanagementTechSupportStatus) GetRelayReasonOk() (*string, bool)`

GetRelayReasonOk returns a tuple with the RelayReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayReason

`func (o *TechsupportmanagementTechSupportStatus) SetRelayReason(v string)`

SetRelayReason sets RelayReason field to given value.

### HasRelayReason

`func (o *TechsupportmanagementTechSupportStatus) HasRelayReason() bool`

HasRelayReason returns a boolean if a field has been set.

### GetRelayStatus

`func (o *TechsupportmanagementTechSupportStatus) GetRelayStatus() string`

GetRelayStatus returns the RelayStatus field if non-nil, zero value otherwise.

### GetRelayStatusOk

`func (o *TechsupportmanagementTechSupportStatus) GetRelayStatusOk() (*string, bool)`

GetRelayStatusOk returns a tuple with the RelayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayStatus

`func (o *TechsupportmanagementTechSupportStatus) SetRelayStatus(v string)`

SetRelayStatus sets RelayStatus field to given value.

### HasRelayStatus

`func (o *TechsupportmanagementTechSupportStatus) HasRelayStatus() bool`

HasRelayStatus returns a boolean if a field has been set.

### GetRequestTs

`func (o *TechsupportmanagementTechSupportStatus) GetRequestTs() time.Time`

GetRequestTs returns the RequestTs field if non-nil, zero value otherwise.

### GetRequestTsOk

`func (o *TechsupportmanagementTechSupportStatus) GetRequestTsOk() (*time.Time, bool)`

GetRequestTsOk returns a tuple with the RequestTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTs

`func (o *TechsupportmanagementTechSupportStatus) SetRequestTs(v time.Time)`

SetRequestTs sets RequestTs field to given value.

### HasRequestTs

`func (o *TechsupportmanagementTechSupportStatus) HasRequestTs() bool`

HasRequestTs returns a boolean if a field has been set.

### GetStatus

`func (o *TechsupportmanagementTechSupportStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TechsupportmanagementTechSupportStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TechsupportmanagementTechSupportStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TechsupportmanagementTechSupportStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatus) GetTechsupportDownloadUrl() string`

GetTechsupportDownloadUrl returns the TechsupportDownloadUrl field if non-nil, zero value otherwise.

### GetTechsupportDownloadUrlOk

`func (o *TechsupportmanagementTechSupportStatus) GetTechsupportDownloadUrlOk() (*string, bool)`

GetTechsupportDownloadUrlOk returns a tuple with the TechsupportDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatus) SetTechsupportDownloadUrl(v string)`

SetTechsupportDownloadUrl sets TechsupportDownloadUrl field to given value.

### HasTechsupportDownloadUrl

`func (o *TechsupportmanagementTechSupportStatus) HasTechsupportDownloadUrl() bool`

HasTechsupportDownloadUrl returns a boolean if a field has been set.

### GetClusterMember

`func (o *TechsupportmanagementTechSupportStatus) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *TechsupportmanagementTechSupportStatus) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *TechsupportmanagementTechSupportStatus) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *TechsupportmanagementTechSupportStatus) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatus) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportStatus) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatus) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportStatus) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetOriginResource

`func (o *TechsupportmanagementTechSupportStatus) GetOriginResource() MoBaseMoRelationship`

GetOriginResource returns the OriginResource field if non-nil, zero value otherwise.

### GetOriginResourceOk

`func (o *TechsupportmanagementTechSupportStatus) GetOriginResourceOk() (*MoBaseMoRelationship, bool)`

GetOriginResourceOk returns a tuple with the OriginResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginResource

`func (o *TechsupportmanagementTechSupportStatus) SetOriginResource(v MoBaseMoRelationship)`

SetOriginResource sets OriginResource field to given value.

### HasOriginResource

`func (o *TechsupportmanagementTechSupportStatus) HasOriginResource() bool`

HasOriginResource returns a boolean if a field has been set.

### GetTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatus) GetTechSupportRequest() TechsupportmanagementTechSupportBundleRelationship`

GetTechSupportRequest returns the TechSupportRequest field if non-nil, zero value otherwise.

### GetTechSupportRequestOk

`func (o *TechsupportmanagementTechSupportStatus) GetTechSupportRequestOk() (*TechsupportmanagementTechSupportBundleRelationship, bool)`

GetTechSupportRequestOk returns a tuple with the TechSupportRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatus) SetTechSupportRequest(v TechsupportmanagementTechSupportBundleRelationship)`

SetTechSupportRequest sets TechSupportRequest field to given value.

### HasTechSupportRequest

`func (o *TechsupportmanagementTechSupportStatus) HasTechSupportRequest() bool`

HasTechSupportRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


