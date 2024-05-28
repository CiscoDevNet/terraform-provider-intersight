# StorageNetAppSnapMirrorRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppSnapMirrorRelationship"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppSnapMirrorRelationship"]
**DestinationPath** | Pointer to **string** | Path to the destination endpoint of a SnapMirror relationship. Examples: ONTAP FlexVol/FlexGroup - svm1:volume1; ONTAP SVM - svm1: ; ONTAP Consistency Group - svm1:/cg/cg_name. | [optional] [readonly] 
**Healthy** | Pointer to **string** | Whether the relationship is healthy or not. | [optional] [readonly] 
**LagTime** | Pointer to **string** | Time since the exported Snapshot copy was created. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Name of the NetApp SnapMirror policy. | [optional] [readonly] 
**PolicyType** | Pointer to **string** | SnapMirror policy type can be async, sync, or continuous. | [optional] [readonly] 
**PolicyUuid** | Pointer to **string** | Uuid of the NetApp SnapMirror policy. | [optional] [readonly] 
**SourcePath** | Pointer to **string** | Path to the source endpoint of a SnapMirror relationship. Examples: ONTAP FlexVol/FlexGroup - svm1:volume1; ONTAP SVM - svm1: ; ONTAP Consistency Group - svm1:/cg/cg_name. | [optional] [readonly] 
**State** | Pointer to **string** | State of the relationship. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uuid of the NetApp SnapMirror relationship. | [optional] [readonly] 
**Array** | Pointer to [**NullableStorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**DestinationTenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableStorageNetAppBaseSnapMirrorPolicyRelationship**](StorageNetAppBaseSnapMirrorPolicyRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppSnapMirrorRelationship

`func NewStorageNetAppSnapMirrorRelationship(classId string, objectType string, ) *StorageNetAppSnapMirrorRelationship`

NewStorageNetAppSnapMirrorRelationship instantiates a new StorageNetAppSnapMirrorRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppSnapMirrorRelationshipWithDefaults

`func NewStorageNetAppSnapMirrorRelationshipWithDefaults() *StorageNetAppSnapMirrorRelationship`

NewStorageNetAppSnapMirrorRelationshipWithDefaults instantiates a new StorageNetAppSnapMirrorRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppSnapMirrorRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppSnapMirrorRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppSnapMirrorRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppSnapMirrorRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppSnapMirrorRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppSnapMirrorRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationPath

`func (o *StorageNetAppSnapMirrorRelationship) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *StorageNetAppSnapMirrorRelationship) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *StorageNetAppSnapMirrorRelationship) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *StorageNetAppSnapMirrorRelationship) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetHealthy

`func (o *StorageNetAppSnapMirrorRelationship) GetHealthy() string`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *StorageNetAppSnapMirrorRelationship) GetHealthyOk() (*string, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *StorageNetAppSnapMirrorRelationship) SetHealthy(v string)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *StorageNetAppSnapMirrorRelationship) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetLagTime

`func (o *StorageNetAppSnapMirrorRelationship) GetLagTime() string`

GetLagTime returns the LagTime field if non-nil, zero value otherwise.

### GetLagTimeOk

`func (o *StorageNetAppSnapMirrorRelationship) GetLagTimeOk() (*string, bool)`

GetLagTimeOk returns a tuple with the LagTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagTime

`func (o *StorageNetAppSnapMirrorRelationship) SetLagTime(v string)`

SetLagTime sets LagTime field to given value.

### HasLagTime

`func (o *StorageNetAppSnapMirrorRelationship) HasLagTime() bool`

HasLagTime returns a boolean if a field has been set.

### GetPolicyName

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *StorageNetAppSnapMirrorRelationship) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *StorageNetAppSnapMirrorRelationship) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyType

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *StorageNetAppSnapMirrorRelationship) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *StorageNetAppSnapMirrorRelationship) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPolicyUuid

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyUuid() string`

GetPolicyUuid returns the PolicyUuid field if non-nil, zero value otherwise.

### GetPolicyUuidOk

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyUuidOk() (*string, bool)`

GetPolicyUuidOk returns a tuple with the PolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUuid

`func (o *StorageNetAppSnapMirrorRelationship) SetPolicyUuid(v string)`

SetPolicyUuid sets PolicyUuid field to given value.

### HasPolicyUuid

`func (o *StorageNetAppSnapMirrorRelationship) HasPolicyUuid() bool`

HasPolicyUuid returns a boolean if a field has been set.

### GetSourcePath

`func (o *StorageNetAppSnapMirrorRelationship) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *StorageNetAppSnapMirrorRelationship) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *StorageNetAppSnapMirrorRelationship) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *StorageNetAppSnapMirrorRelationship) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppSnapMirrorRelationship) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppSnapMirrorRelationship) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppSnapMirrorRelationship) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppSnapMirrorRelationship) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppSnapMirrorRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppSnapMirrorRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppSnapMirrorRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppSnapMirrorRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppSnapMirrorRelationship) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppSnapMirrorRelationship) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppSnapMirrorRelationship) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppSnapMirrorRelationship) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageNetAppSnapMirrorRelationship) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageNetAppSnapMirrorRelationship) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetDestinationTenant

`func (o *StorageNetAppSnapMirrorRelationship) GetDestinationTenant() StorageNetAppStorageVmRelationship`

GetDestinationTenant returns the DestinationTenant field if non-nil, zero value otherwise.

### GetDestinationTenantOk

`func (o *StorageNetAppSnapMirrorRelationship) GetDestinationTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetDestinationTenantOk returns a tuple with the DestinationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTenant

`func (o *StorageNetAppSnapMirrorRelationship) SetDestinationTenant(v StorageNetAppStorageVmRelationship)`

SetDestinationTenant sets DestinationTenant field to given value.

### HasDestinationTenant

`func (o *StorageNetAppSnapMirrorRelationship) HasDestinationTenant() bool`

HasDestinationTenant returns a boolean if a field has been set.

### SetDestinationTenantNil

`func (o *StorageNetAppSnapMirrorRelationship) SetDestinationTenantNil(b bool)`

 SetDestinationTenantNil sets the value for DestinationTenant to be an explicit nil

### UnsetDestinationTenant
`func (o *StorageNetAppSnapMirrorRelationship) UnsetDestinationTenant()`

UnsetDestinationTenant ensures that no value is present for DestinationTenant, not even an explicit nil
### GetPolicy

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicy() StorageNetAppBaseSnapMirrorPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StorageNetAppSnapMirrorRelationship) GetPolicyOk() (*StorageNetAppBaseSnapMirrorPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StorageNetAppSnapMirrorRelationship) SetPolicy(v StorageNetAppBaseSnapMirrorPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StorageNetAppSnapMirrorRelationship) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *StorageNetAppSnapMirrorRelationship) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *StorageNetAppSnapMirrorRelationship) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


