# CloudAwsOrganizationalUnitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsOrganizationalUnit"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsOrganizationalUnit"]
**Identity** | Pointer to **string** | The identity of this organization. This entity is not manipulated by users. It aids in uniquely identifying the organization object. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the organizational unit. | [optional] [readonly] 
**ParentOrg** | Pointer to [**CloudAwsOrganizationalUnitRelationship**](cloud.AwsOrganizationalUnit.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsOrganizationalUnitAllOf

`func NewCloudAwsOrganizationalUnitAllOf(classId string, objectType string, ) *CloudAwsOrganizationalUnitAllOf`

NewCloudAwsOrganizationalUnitAllOf instantiates a new CloudAwsOrganizationalUnitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsOrganizationalUnitAllOfWithDefaults

`func NewCloudAwsOrganizationalUnitAllOfWithDefaults() *CloudAwsOrganizationalUnitAllOf`

NewCloudAwsOrganizationalUnitAllOfWithDefaults instantiates a new CloudAwsOrganizationalUnitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsOrganizationalUnitAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsOrganizationalUnitAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsOrganizationalUnitAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsOrganizationalUnitAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsOrganizationalUnitAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsOrganizationalUnitAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *CloudAwsOrganizationalUnitAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudAwsOrganizationalUnitAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudAwsOrganizationalUnitAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudAwsOrganizationalUnitAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudAwsOrganizationalUnitAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAwsOrganizationalUnitAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAwsOrganizationalUnitAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAwsOrganizationalUnitAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentOrg

`func (o *CloudAwsOrganizationalUnitAllOf) GetParentOrg() CloudAwsOrganizationalUnitRelationship`

GetParentOrg returns the ParentOrg field if non-nil, zero value otherwise.

### GetParentOrgOk

`func (o *CloudAwsOrganizationalUnitAllOf) GetParentOrgOk() (*CloudAwsOrganizationalUnitRelationship, bool)`

GetParentOrgOk returns a tuple with the ParentOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrg

`func (o *CloudAwsOrganizationalUnitAllOf) SetParentOrg(v CloudAwsOrganizationalUnitRelationship)`

SetParentOrg sets ParentOrg field to given value.

### HasParentOrg

`func (o *CloudAwsOrganizationalUnitAllOf) HasParentOrg() bool`

HasParentOrg returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CloudAwsOrganizationalUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CloudAwsOrganizationalUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CloudAwsOrganizationalUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CloudAwsOrganizationalUnitAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


