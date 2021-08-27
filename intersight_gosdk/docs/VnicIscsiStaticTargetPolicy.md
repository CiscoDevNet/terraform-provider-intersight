# VnicIscsiStaticTargetPolicy

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

### NewVnicIscsiStaticTargetPolicy

`func NewVnicIscsiStaticTargetPolicy(classId string, objectType string, ) *VnicIscsiStaticTargetPolicy`

NewVnicIscsiStaticTargetPolicy instantiates a new VnicIscsiStaticTargetPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicIscsiStaticTargetPolicyWithDefaults

`func NewVnicIscsiStaticTargetPolicyWithDefaults() *VnicIscsiStaticTargetPolicy`

NewVnicIscsiStaticTargetPolicyWithDefaults instantiates a new VnicIscsiStaticTargetPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicIscsiStaticTargetPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicIscsiStaticTargetPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicIscsiStaticTargetPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicIscsiStaticTargetPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicIscsiStaticTargetPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicIscsiStaticTargetPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpAddress

`func (o *VnicIscsiStaticTargetPolicy) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VnicIscsiStaticTargetPolicy) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VnicIscsiStaticTargetPolicy) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VnicIscsiStaticTargetPolicy) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLun

`func (o *VnicIscsiStaticTargetPolicy) GetLun() VnicLun`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *VnicIscsiStaticTargetPolicy) GetLunOk() (*VnicLun, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *VnicIscsiStaticTargetPolicy) SetLun(v VnicLun)`

SetLun sets Lun field to given value.

### HasLun

`func (o *VnicIscsiStaticTargetPolicy) HasLun() bool`

HasLun returns a boolean if a field has been set.

### SetLunNil

`func (o *VnicIscsiStaticTargetPolicy) SetLunNil(b bool)`

 SetLunNil sets the value for Lun to be an explicit nil

### UnsetLun
`func (o *VnicIscsiStaticTargetPolicy) UnsetLun()`

UnsetLun ensures that no value is present for Lun, not even an explicit nil
### GetPort

`func (o *VnicIscsiStaticTargetPolicy) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VnicIscsiStaticTargetPolicy) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VnicIscsiStaticTargetPolicy) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VnicIscsiStaticTargetPolicy) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTargetName

`func (o *VnicIscsiStaticTargetPolicy) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *VnicIscsiStaticTargetPolicy) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *VnicIscsiStaticTargetPolicy) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *VnicIscsiStaticTargetPolicy) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetOrganization

`func (o *VnicIscsiStaticTargetPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VnicIscsiStaticTargetPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VnicIscsiStaticTargetPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VnicIscsiStaticTargetPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


