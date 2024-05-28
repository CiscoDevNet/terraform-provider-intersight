# NiatelemetryNexusCloudAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NexusCloudAccount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NexusCloudAccount"]
**AciCount** | Pointer to **int64** | Count of ACI-type site devices. | [optional] 
**NxosCount** | Pointer to **int64** | Count of NXOS-type site devices. | [optional] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNexusCloudAccount

`func NewNiatelemetryNexusCloudAccount(classId string, objectType string, ) *NiatelemetryNexusCloudAccount`

NewNiatelemetryNexusCloudAccount instantiates a new NiatelemetryNexusCloudAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNexusCloudAccountWithDefaults

`func NewNiatelemetryNexusCloudAccountWithDefaults() *NiatelemetryNexusCloudAccount`

NewNiatelemetryNexusCloudAccountWithDefaults instantiates a new NiatelemetryNexusCloudAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNexusCloudAccount) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNexusCloudAccount) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNexusCloudAccount) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNexusCloudAccount) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNexusCloudAccount) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNexusCloudAccount) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAciCount

`func (o *NiatelemetryNexusCloudAccount) GetAciCount() int64`

GetAciCount returns the AciCount field if non-nil, zero value otherwise.

### GetAciCountOk

`func (o *NiatelemetryNexusCloudAccount) GetAciCountOk() (*int64, bool)`

GetAciCountOk returns a tuple with the AciCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAciCount

`func (o *NiatelemetryNexusCloudAccount) SetAciCount(v int64)`

SetAciCount sets AciCount field to given value.

### HasAciCount

`func (o *NiatelemetryNexusCloudAccount) HasAciCount() bool`

HasAciCount returns a boolean if a field has been set.

### GetNxosCount

`func (o *NiatelemetryNexusCloudAccount) GetNxosCount() int64`

GetNxosCount returns the NxosCount field if non-nil, zero value otherwise.

### GetNxosCountOk

`func (o *NiatelemetryNexusCloudAccount) GetNxosCountOk() (*int64, bool)`

GetNxosCountOk returns a tuple with the NxosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosCount

`func (o *NiatelemetryNexusCloudAccount) SetNxosCount(v int64)`

SetNxosCount sets NxosCount field to given value.

### HasNxosCount

`func (o *NiatelemetryNexusCloudAccount) HasNxosCount() bool`

HasNxosCount returns a boolean if a field has been set.

### GetAccount

`func (o *NiatelemetryNexusCloudAccount) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NiatelemetryNexusCloudAccount) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NiatelemetryNexusCloudAccount) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NiatelemetryNexusCloudAccount) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *NiatelemetryNexusCloudAccount) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *NiatelemetryNexusCloudAccount) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


