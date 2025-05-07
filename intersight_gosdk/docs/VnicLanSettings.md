# VnicLanSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.LanSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.LanSettings"]
**IqnName** | Pointer to **string** | Qualified Name (IQN) assigned to iSCSI vNICs in a Fabric Interconnect domain. | [optional] [readonly] 
**IqnLease** | Pointer to [**NullableIqnpoolLeaseRelationship**](IqnpoolLeaseRelationship.md) |  | [optional] 
**Profile** | Pointer to [**NullablePolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) |  | [optional] 

## Methods

### NewVnicLanSettings

`func NewVnicLanSettings(classId string, objectType string, ) *VnicLanSettings`

NewVnicLanSettings instantiates a new VnicLanSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicLanSettingsWithDefaults

`func NewVnicLanSettingsWithDefaults() *VnicLanSettings`

NewVnicLanSettingsWithDefaults instantiates a new VnicLanSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicLanSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicLanSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicLanSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicLanSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicLanSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicLanSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIqnName

`func (o *VnicLanSettings) GetIqnName() string`

GetIqnName returns the IqnName field if non-nil, zero value otherwise.

### GetIqnNameOk

`func (o *VnicLanSettings) GetIqnNameOk() (*string, bool)`

GetIqnNameOk returns a tuple with the IqnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnName

`func (o *VnicLanSettings) SetIqnName(v string)`

SetIqnName sets IqnName field to given value.

### HasIqnName

`func (o *VnicLanSettings) HasIqnName() bool`

HasIqnName returns a boolean if a field has been set.

### GetIqnLease

`func (o *VnicLanSettings) GetIqnLease() IqnpoolLeaseRelationship`

GetIqnLease returns the IqnLease field if non-nil, zero value otherwise.

### GetIqnLeaseOk

`func (o *VnicLanSettings) GetIqnLeaseOk() (*IqnpoolLeaseRelationship, bool)`

GetIqnLeaseOk returns a tuple with the IqnLease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqnLease

`func (o *VnicLanSettings) SetIqnLease(v IqnpoolLeaseRelationship)`

SetIqnLease sets IqnLease field to given value.

### HasIqnLease

`func (o *VnicLanSettings) HasIqnLease() bool`

HasIqnLease returns a boolean if a field has been set.

### SetIqnLeaseNil

`func (o *VnicLanSettings) SetIqnLeaseNil(b bool)`

 SetIqnLeaseNil sets the value for IqnLease to be an explicit nil

### UnsetIqnLease
`func (o *VnicLanSettings) UnsetIqnLease()`

UnsetIqnLease ensures that no value is present for IqnLease, not even an explicit nil
### GetProfile

`func (o *VnicLanSettings) GetProfile() PolicyAbstractConfigProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *VnicLanSettings) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *VnicLanSettings) SetProfile(v PolicyAbstractConfigProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *VnicLanSettings) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *VnicLanSettings) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *VnicLanSettings) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


