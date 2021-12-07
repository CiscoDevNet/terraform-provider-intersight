# ConnectorTargetChangeMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.TargetChangeMessage"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.TargetChangeMessage"]
**ModStatus** | Pointer to **string** | ModStatus indicates if the change message conveys a creation, modification or deletion of an Target instance. * &#x60;None&#x60; - The &#39;none&#39; operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * &#x60;Created&#x60; - The &#39;create&#39; operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * &#x60;Modified&#x60; - The &#39;update&#39; operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully. * &#x60;Deleted&#x60; - The &#39;delete&#39; operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. | [optional] [default to "None"]
**TargetDetails** | Pointer to **interface{}** | A Json-serialized representation of the &#39;configuration&#39; portion of the Target instance. Ie the representation contains configuration properties like the target&#39;s connectivity information but not operation status. The representation include credential information, encrypted with the RSA public key of the Appliance device connector. Appliance device connector is the sole maintainer of the RSA private key and the only system component which is capable of interpreting the credential. | [optional] 
**TargetMoid** | Pointer to **string** | The Moid identifying the Target instance being created, modified or deleted. | [optional] 

## Methods

### NewConnectorTargetChangeMessage

`func NewConnectorTargetChangeMessage(classId string, objectType string, ) *ConnectorTargetChangeMessage`

NewConnectorTargetChangeMessage instantiates a new ConnectorTargetChangeMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTargetChangeMessageWithDefaults

`func NewConnectorTargetChangeMessageWithDefaults() *ConnectorTargetChangeMessage`

NewConnectorTargetChangeMessageWithDefaults instantiates a new ConnectorTargetChangeMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorTargetChangeMessage) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorTargetChangeMessage) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorTargetChangeMessage) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorTargetChangeMessage) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorTargetChangeMessage) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorTargetChangeMessage) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetModStatus

`func (o *ConnectorTargetChangeMessage) GetModStatus() string`

GetModStatus returns the ModStatus field if non-nil, zero value otherwise.

### GetModStatusOk

`func (o *ConnectorTargetChangeMessage) GetModStatusOk() (*string, bool)`

GetModStatusOk returns a tuple with the ModStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModStatus

`func (o *ConnectorTargetChangeMessage) SetModStatus(v string)`

SetModStatus sets ModStatus field to given value.

### HasModStatus

`func (o *ConnectorTargetChangeMessage) HasModStatus() bool`

HasModStatus returns a boolean if a field has been set.

### GetTargetDetails

`func (o *ConnectorTargetChangeMessage) GetTargetDetails() interface{}`

GetTargetDetails returns the TargetDetails field if non-nil, zero value otherwise.

### GetTargetDetailsOk

`func (o *ConnectorTargetChangeMessage) GetTargetDetailsOk() (*interface{}, bool)`

GetTargetDetailsOk returns a tuple with the TargetDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDetails

`func (o *ConnectorTargetChangeMessage) SetTargetDetails(v interface{})`

SetTargetDetails sets TargetDetails field to given value.

### HasTargetDetails

`func (o *ConnectorTargetChangeMessage) HasTargetDetails() bool`

HasTargetDetails returns a boolean if a field has been set.

### SetTargetDetailsNil

`func (o *ConnectorTargetChangeMessage) SetTargetDetailsNil(b bool)`

 SetTargetDetailsNil sets the value for TargetDetails to be an explicit nil

### UnsetTargetDetails
`func (o *ConnectorTargetChangeMessage) UnsetTargetDetails()`

UnsetTargetDetails ensures that no value is present for TargetDetails, not even an explicit nil
### GetTargetMoid

`func (o *ConnectorTargetChangeMessage) GetTargetMoid() string`

GetTargetMoid returns the TargetMoid field if non-nil, zero value otherwise.

### GetTargetMoidOk

`func (o *ConnectorTargetChangeMessage) GetTargetMoidOk() (*string, bool)`

GetTargetMoidOk returns a tuple with the TargetMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoid

`func (o *ConnectorTargetChangeMessage) SetTargetMoid(v string)`

SetTargetMoid sets TargetMoid field to given value.

### HasTargetMoid

`func (o *ConnectorTargetChangeMessage) HasTargetMoid() bool`

HasTargetMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


