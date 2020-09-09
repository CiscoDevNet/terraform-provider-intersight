# AssetTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]AssetConnection**](asset.Connection.md) |  | [optional] 
**Services** | Pointer to [**[]AssetService**](asset.Service.md) |  | [optional] 
**Status** | Pointer to **string** | Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target. * &#x60;&#x60; - The device registered with Intersight but subsequently did not establish a persistent websocket connection. * &#x60;Connected&#x60; - The device&#39;s connection to Intersight has been established and is active. * &#x60;NotConnected&#x60; - The device&#39;s connection to Intersight has been disconnected. * &#x60;ClaimInProgress&#x60; - Claim of the device is in progress. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. | [optional] [readonly] [default to ""]
**StatusErrorReason** | Pointer to **string** | StatusErrorReason provides additional context for the Status. | [optional] [readonly] 
**TargetType** | Pointer to **string** | The type of the managed target. For example a UCS Server or Vmware Vcenter target. * &#x60;&#x60; - The device reported an empty or unrecognized platform type. * &#x60;APIC&#x60; - An Application Policy Infrastructure Controller cluster. * &#x60;DCNM&#x60; - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * &#x60;UCSFI&#x60; - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * &#x60;UCSFIISM&#x60; - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * &#x60;IMC&#x60; - A standalone UCS Server Integrated Management Controller. * &#x60;IMCM4&#x60; - A standalone UCS M4 Server. * &#x60;IMCM5&#x60; - A standalone UCS M5 server. * &#x60;UCSIOM&#x60; - An UCS Chassis IO module. * &#x60;HX&#x60; - A HyperFlex storage controller. * &#x60;HyperFlexAP&#x60; - A HyperFlex Application Platform. * &#x60;UCSD&#x60; - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * &#x60;IntersightAppliance&#x60; - Intersight on-premise appliance. * &#x60;PureStorageFlashArray&#x60; - A Pure Storage FlashArray device. * &#x60;NetAppOntap&#x60; - A NetApp ONTAP storage system. * &#x60;EmcScaleIo&#x60; - An EMC ScaleIO storage system. * &#x60;EmcVmax&#x60; - An EMC VMAX storage system. * &#x60;EmcVplex&#x60; - An EMC VPLEX storage system. * &#x60;EmcXtremIo&#x60; - An EMC XtremIO storage system. * &#x60;VmwareVcenter&#x60; - A VMware vCenter device that manages Virtual Machines. * &#x60;MicrosoftHyperV&#x60; - A Microsoft HyperV system that manages Virtual Machines. * &#x60;AppDynamics&#x60; - An AppDynamics controller that monitors applications. * &#x60;Dynatrace&#x60; - A Dynatrace controller that monitors applications. * &#x60;MicrosoftSqlServer&#x60; - A Microsoft SQL database server. * &#x60;Kubernetes&#x60; - A Kubernetes cluster that runs containerized applications. * &#x60;MicrosoftAzure&#x60; - A Microsoft Azure target. * &#x60;ServiceEngine&#x60; - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * &#x60;IMCBlade&#x60; - An Intersight managed UCS Blade Server. | [optional] [default to ""]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Assist** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetTargetAllOf

`func NewAssetTargetAllOf() *AssetTargetAllOf`

NewAssetTargetAllOf instantiates a new AssetTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTargetAllOfWithDefaults

`func NewAssetTargetAllOfWithDefaults() *AssetTargetAllOf`

NewAssetTargetAllOfWithDefaults instantiates a new AssetTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *AssetTargetAllOf) GetConnections() []AssetConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *AssetTargetAllOf) GetConnectionsOk() (*[]AssetConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *AssetTargetAllOf) SetConnections(v []AssetConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *AssetTargetAllOf) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetServices

`func (o *AssetTargetAllOf) GetServices() []AssetService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AssetTargetAllOf) GetServicesOk() (*[]AssetService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AssetTargetAllOf) SetServices(v []AssetService)`

SetServices sets Services field to given value.

### HasServices

`func (o *AssetTargetAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStatus

`func (o *AssetTargetAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetTargetAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetTargetAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetTargetAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusErrorReason

`func (o *AssetTargetAllOf) GetStatusErrorReason() string`

GetStatusErrorReason returns the StatusErrorReason field if non-nil, zero value otherwise.

### GetStatusErrorReasonOk

`func (o *AssetTargetAllOf) GetStatusErrorReasonOk() (*string, bool)`

GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusErrorReason

`func (o *AssetTargetAllOf) SetStatusErrorReason(v string)`

SetStatusErrorReason sets StatusErrorReason field to given value.

### HasStatusErrorReason

`func (o *AssetTargetAllOf) HasStatusErrorReason() bool`

HasStatusErrorReason returns a boolean if a field has been set.

### GetTargetType

`func (o *AssetTargetAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AssetTargetAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AssetTargetAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AssetTargetAllOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetAccount

`func (o *AssetTargetAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetTargetAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetTargetAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetTargetAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssist

`func (o *AssetTargetAllOf) GetAssist() AssetDeviceRegistrationRelationship`

GetAssist returns the Assist field if non-nil, zero value otherwise.

### GetAssistOk

`func (o *AssetTargetAllOf) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool)`

GetAssistOk returns a tuple with the Assist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssist

`func (o *AssetTargetAllOf) SetAssist(v AssetDeviceRegistrationRelationship)`

SetAssist sets Assist field to given value.

### HasAssist

`func (o *AssetTargetAllOf) HasAssist() bool`

HasAssist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


