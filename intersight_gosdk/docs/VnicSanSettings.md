# VnicSanSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.SanSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.SanSettings"]
**OldInfo** | Pointer to [**NullableVnicSanSettingsOldInfo**](VnicSanSettingsOldInfo.md) |  | [optional] 
**Wwnn** | Pointer to **string** | The WWNN address that is assigned to the server node based on the wwn pool that has been assigned in the SAN Connectivity Policy. | [optional] [readonly] 
**Profile** | Pointer to [**NullablePolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 
**WwnnLease** | Pointer to [**NullableFcpoolLeaseRelationship**](FcpoolLeaseRelationship.md) |  | [optional] 

## Methods

### NewVnicSanSettings

`func NewVnicSanSettings(classId string, objectType string, ) *VnicSanSettings`

NewVnicSanSettings instantiates a new VnicSanSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicSanSettingsWithDefaults

`func NewVnicSanSettingsWithDefaults() *VnicSanSettings`

NewVnicSanSettingsWithDefaults instantiates a new VnicSanSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicSanSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicSanSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicSanSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicSanSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicSanSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicSanSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOldInfo

`func (o *VnicSanSettings) GetOldInfo() VnicSanSettingsOldInfo`

GetOldInfo returns the OldInfo field if non-nil, zero value otherwise.

### GetOldInfoOk

`func (o *VnicSanSettings) GetOldInfoOk() (*VnicSanSettingsOldInfo, bool)`

GetOldInfoOk returns a tuple with the OldInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldInfo

`func (o *VnicSanSettings) SetOldInfo(v VnicSanSettingsOldInfo)`

SetOldInfo sets OldInfo field to given value.

### HasOldInfo

`func (o *VnicSanSettings) HasOldInfo() bool`

HasOldInfo returns a boolean if a field has been set.

### SetOldInfoNil

`func (o *VnicSanSettings) SetOldInfoNil(b bool)`

 SetOldInfoNil sets the value for OldInfo to be an explicit nil

### UnsetOldInfo
`func (o *VnicSanSettings) UnsetOldInfo()`

UnsetOldInfo ensures that no value is present for OldInfo, not even an explicit nil
### GetWwnn

`func (o *VnicSanSettings) GetWwnn() string`

GetWwnn returns the Wwnn field if non-nil, zero value otherwise.

### GetWwnnOk

`func (o *VnicSanSettings) GetWwnnOk() (*string, bool)`

GetWwnnOk returns a tuple with the Wwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnn

`func (o *VnicSanSettings) SetWwnn(v string)`

SetWwnn sets Wwnn field to given value.

### HasWwnn

`func (o *VnicSanSettings) HasWwnn() bool`

HasWwnn returns a boolean if a field has been set.

### GetProfile

`func (o *VnicSanSettings) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicSanSettings) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicSanSettings) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicSanSettings) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *VnicSanSettings) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *VnicSanSettings) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetWwnnLease

`func (o *VnicSanSettings) GetWwnnLease() FcpoolLeaseRelationship`

GetWwnnLease returns the WwnnLease field if non-nil, zero value otherwise.

### GetWwnnLeaseOk

`func (o *VnicSanSettings) GetWwnnLeaseOk() (*FcpoolLeaseRelationship, bool)`

GetWwnnLeaseOk returns a tuple with the WwnnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnnLease

`func (o *VnicSanSettings) SetWwnnLease(v FcpoolLeaseRelationship)`

SetWwnnLease sets WwnnLease field to given value.

### HasWwnnLease

`func (o *VnicSanSettings) HasWwnnLease() bool`

HasWwnnLease returns a boolean if a field has been set.

### SetWwnnLeaseNil

`func (o *VnicSanSettings) SetWwnnLeaseNil(b bool)`

 SetWwnnLeaseNil sets the value for WwnnLease to be an explicit nil

### UnsetWwnnLease
`func (o *VnicSanSettings) UnsetWwnnLease()`

UnsetWwnnLease ensures that no value is present for WwnnLease, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


