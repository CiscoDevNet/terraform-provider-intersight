# CapabilityServerSchemaDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerSchemaDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerSchemaDescriptor"]
**LocatorLedName** | Pointer to **string** | Redfish property name for the server. | [optional] 
**RedfishSchema** | Pointer to **string** | Redfish version information for the server. | [optional] 

## Methods

### NewCapabilityServerSchemaDescriptor

`func NewCapabilityServerSchemaDescriptor(classId string, objectType string, ) *CapabilityServerSchemaDescriptor`

NewCapabilityServerSchemaDescriptor instantiates a new CapabilityServerSchemaDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerSchemaDescriptorWithDefaults

`func NewCapabilityServerSchemaDescriptorWithDefaults() *CapabilityServerSchemaDescriptor`

NewCapabilityServerSchemaDescriptorWithDefaults instantiates a new CapabilityServerSchemaDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerSchemaDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerSchemaDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerSchemaDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerSchemaDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerSchemaDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerSchemaDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLocatorLedName

`func (o *CapabilityServerSchemaDescriptor) GetLocatorLedName() string`

GetLocatorLedName returns the LocatorLedName field if non-nil, zero value otherwise.

### GetLocatorLedNameOk

`func (o *CapabilityServerSchemaDescriptor) GetLocatorLedNameOk() (*string, bool)`

GetLocatorLedNameOk returns a tuple with the LocatorLedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLedName

`func (o *CapabilityServerSchemaDescriptor) SetLocatorLedName(v string)`

SetLocatorLedName sets LocatorLedName field to given value.

### HasLocatorLedName

`func (o *CapabilityServerSchemaDescriptor) HasLocatorLedName() bool`

HasLocatorLedName returns a boolean if a field has been set.

### GetRedfishSchema

`func (o *CapabilityServerSchemaDescriptor) GetRedfishSchema() string`

GetRedfishSchema returns the RedfishSchema field if non-nil, zero value otherwise.

### GetRedfishSchemaOk

`func (o *CapabilityServerSchemaDescriptor) GetRedfishSchemaOk() (*string, bool)`

GetRedfishSchemaOk returns a tuple with the RedfishSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedfishSchema

`func (o *CapabilityServerSchemaDescriptor) SetRedfishSchema(v string)`

SetRedfishSchema sets RedfishSchema field to given value.

### HasRedfishSchema

`func (o *CapabilityServerSchemaDescriptor) HasRedfishSchema() bool`

HasRedfishSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


