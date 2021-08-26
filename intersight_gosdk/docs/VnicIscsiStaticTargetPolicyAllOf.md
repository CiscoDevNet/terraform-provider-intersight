# VnicIscsiStaticTargetPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.IscsiStaticTargetPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.IscsiStaticTargetPolicy"]
**IpAddress** | Pointer to **string** | The IPv4 address assigned to the iSCSI target. | [optional] 
**Lun** | Pointer to [**NullableVnicLun**](VnicLun.md) |  | [optional] 
**Port** | Pointer to **int64** | The port associated with the iSCSI target. | [optional] 
**TargetName** | Pointer to **string** | Qualified Name (IQN) or Extended Unique Identifier (EUI) name of the iSCSI target. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewVnicIscsiStaticTargetPolicyAllOf

`func NewVnicIscsiStaticTargetPolicyAllOf(classId string, objectType string, ) *VnicIscsiStaticTargetPolicyAllOf`

NewVnicIscsiStaticTargetPolicyAllOf instantiates a new VnicIscsiStaticTargetPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiStaticTargetPolicyAllOfWithDefaults

`func NewVnicIscsiStaticTargetPolicyAllOfWithDefaults() *VnicIscsiStaticTargetPolicyAllOf`

NewVnicIscsiStaticTargetPolicyAllOfWithDefaults instantiates a new VnicIscsiStaticTargetPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddress

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VnicIscsiStaticTargetPolicyAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLun

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetLun() VnicLun`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetLunOk() (*VnicLun, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetLun(v VnicLun)`

SetLun sets Lun field to given value.

### HasLun

`func (o *VnicIscsiStaticTargetPolicyAllOf) HasLun() bool`

HasLun returns a boolean if a field has been set.

### SetLunNil

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetLunNil(b bool)`

 SetLunNil sets the value for Lun to be an explicit nil

### UnsetLun
`func (o *VnicIscsiStaticTargetPolicyAllOf) UnsetLun()`

UnsetLun ensures that no value is present for Lun, not even an explicit nil
### GetPort

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VnicIscsiStaticTargetPolicyAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTargetName

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *VnicIscsiStaticTargetPolicyAllOf) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicIscsiStaticTargetPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicIscsiStaticTargetPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicIscsiStaticTargetPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


