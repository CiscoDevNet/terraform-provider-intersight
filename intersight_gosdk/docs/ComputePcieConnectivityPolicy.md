# ComputePcieConnectivityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.PcieConnectivityPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.PcieConnectivityPolicy"]
**PcieZones** | Pointer to [**[]ComputePcieZone**](ComputePcieZone.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]PolicyAbstractConfigProfileRelationship**](PolicyAbstractConfigProfileRelationship.md) | An array of relationships to policyAbstractConfigProfile resources. | [optional] 

## Methods

### NewComputePcieConnectivityPolicy

`func NewComputePcieConnectivityPolicy(classId string, objectType string, ) *ComputePcieConnectivityPolicy`

NewComputePcieConnectivityPolicy instantiates a new ComputePcieConnectivityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePcieConnectivityPolicyWithDefaults

`func NewComputePcieConnectivityPolicyWithDefaults() *ComputePcieConnectivityPolicy`

NewComputePcieConnectivityPolicyWithDefaults instantiates a new ComputePcieConnectivityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputePcieConnectivityPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputePcieConnectivityPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputePcieConnectivityPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputePcieConnectivityPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputePcieConnectivityPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputePcieConnectivityPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPcieZones

`func (o *ComputePcieConnectivityPolicy) GetPcieZones() []ComputePcieZone`

GetPcieZones returns the PcieZones field if non-nil, zero value otherwise.

### GetPcieZonesOk

`func (o *ComputePcieConnectivityPolicy) GetPcieZonesOk() (*[]ComputePcieZone, bool)`

GetPcieZonesOk returns a tuple with the PcieZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieZones

`func (o *ComputePcieConnectivityPolicy) SetPcieZones(v []ComputePcieZone)`

SetPcieZones sets PcieZones field to given value.

### HasPcieZones

`func (o *ComputePcieConnectivityPolicy) HasPcieZones() bool`

HasPcieZones returns a boolean if a field has been set.

### SetPcieZonesNil

`func (o *ComputePcieConnectivityPolicy) SetPcieZonesNil(b bool)`

 SetPcieZonesNil sets the value for PcieZones to be an explicit nil

### UnsetPcieZones
`func (o *ComputePcieConnectivityPolicy) UnsetPcieZones()`

UnsetPcieZones ensures that no value is present for PcieZones, not even an explicit nil
### GetOrganization

`func (o *ComputePcieConnectivityPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ComputePcieConnectivityPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ComputePcieConnectivityPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ComputePcieConnectivityPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *ComputePcieConnectivityPolicy) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *ComputePcieConnectivityPolicy) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetProfiles

`func (o *ComputePcieConnectivityPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *ComputePcieConnectivityPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *ComputePcieConnectivityPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *ComputePcieConnectivityPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *ComputePcieConnectivityPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *ComputePcieConnectivityPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


