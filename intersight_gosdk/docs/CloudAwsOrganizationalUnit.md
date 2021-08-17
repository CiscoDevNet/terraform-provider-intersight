# CloudAwsOrganizationalUnit

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

### NewCloudAwsOrganizationalUnit

`func NewCloudAwsOrganizationalUnit(classId string, objectType string, ) *CloudAwsOrganizationalUnit`

NewCloudAwsOrganizationalUnit instantiates a new CloudAwsOrganizationalUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsOrganizationalUnitWithDefaults

`func NewCloudAwsOrganizationalUnitWithDefaults() *CloudAwsOrganizationalUnit`

NewCloudAwsOrganizationalUnitWithDefaults instantiates a new CloudAwsOrganizationalUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsOrganizationalUnit) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsOrganizationalUnit) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsOrganizationalUnit) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsOrganizationalUnit) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsOrganizationalUnit) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsOrganizationalUnit) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIdentity

`func (o *CloudAwsOrganizationalUnit) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudAwsOrganizationalUnit) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudAwsOrganizationalUnit) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudAwsOrganizationalUnit) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetName

`func (o *CloudAwsOrganizationalUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAwsOrganizationalUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAwsOrganizationalUnit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAwsOrganizationalUnit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentOrg

`func (o *CloudAwsOrganizationalUnit) GetParentOrg() CloudAwsOrganizationalUnitRelationship`

GetParentOrg returns the ParentOrg field if non-nil, zero value otherwise.

### GetParentOrgOk

`func (o *CloudAwsOrganizationalUnit) GetParentOrgOk() (*CloudAwsOrganizationalUnitRelationship, bool)`

GetParentOrgOk returns a tuple with the ParentOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrg

`func (o *CloudAwsOrganizationalUnit) SetParentOrg(v CloudAwsOrganizationalUnitRelationship)`

SetParentOrg sets ParentOrg field to given value.

### HasParentOrg

`func (o *CloudAwsOrganizationalUnit) HasParentOrg() bool`

HasParentOrg returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CloudAwsOrganizationalUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CloudAwsOrganizationalUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CloudAwsOrganizationalUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CloudAwsOrganizationalUnit) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


