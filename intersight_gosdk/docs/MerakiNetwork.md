# MerakiNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meraki.Network"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meraki.Network"]
**Name** | Pointer to **string** | The Meraki network name seamlessly uniting devices for effortless connectivity. | [optional] 
**NetworkId** | Pointer to **string** | The unique Meraki network id. | [optional] 
**NetworkTags** | Pointer to **[]string** |  | [optional] 
**ProductTypes** | Pointer to **[]string** |  | [optional] 
**Organization** | Pointer to [**NullableMerakiOrganizationRelationship**](MerakiOrganizationRelationship.md) |  | [optional] 

## Methods

### NewMerakiNetwork

`func NewMerakiNetwork(classId string, objectType string, ) *MerakiNetwork`

NewMerakiNetwork instantiates a new MerakiNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerakiNetworkWithDefaults

`func NewMerakiNetworkWithDefaults() *MerakiNetwork`

NewMerakiNetworkWithDefaults instantiates a new MerakiNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MerakiNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MerakiNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MerakiNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MerakiNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MerakiNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MerakiNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *MerakiNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerakiNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerakiNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerakiNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *MerakiNetwork) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *MerakiNetwork) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *MerakiNetwork) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *MerakiNetwork) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkTags

`func (o *MerakiNetwork) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *MerakiNetwork) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *MerakiNetwork) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *MerakiNetwork) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### SetNetworkTagsNil

`func (o *MerakiNetwork) SetNetworkTagsNil(b bool)`

 SetNetworkTagsNil sets the value for NetworkTags to be an explicit nil

### UnsetNetworkTags
`func (o *MerakiNetwork) UnsetNetworkTags()`

UnsetNetworkTags ensures that no value is present for NetworkTags, not even an explicit nil
### GetProductTypes

`func (o *MerakiNetwork) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *MerakiNetwork) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *MerakiNetwork) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *MerakiNetwork) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### SetProductTypesNil

`func (o *MerakiNetwork) SetProductTypesNil(b bool)`

 SetProductTypesNil sets the value for ProductTypes to be an explicit nil

### UnsetProductTypes
`func (o *MerakiNetwork) UnsetProductTypes()`

UnsetProductTypes ensures that no value is present for ProductTypes, not even an explicit nil
### GetOrganization

`func (o *MerakiNetwork) GetOrganization() MerakiOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MerakiNetwork) GetOrganizationOk() (*MerakiOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MerakiNetwork) SetOrganization(v MerakiOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MerakiNetwork) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *MerakiNetwork) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *MerakiNetwork) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


