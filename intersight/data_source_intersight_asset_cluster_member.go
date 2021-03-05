package intersight

import (
	"context"
	"log"
	"reflect"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssetClusterMember() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAssetClusterMemberRead,
		Schema: map[string]*schema.Schema{
			"api_version": {
				Description: "The version of the connector API, describes the capability of the connector's framework.\nIf the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"app_partition_number": {
				Description: "The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"connection_id": {
				Description: "The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_reason": {
				Description: "If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_status": {
				Description: "The status of the persistent connection between the device connector and Intersight.\n* `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.\n* `Connected` - Intersight is able to establish a connection to the target and initiate management activities.\n* `NotConnected` - Intersight is unable to establish a connection to the target.\n* `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.\n* `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.\n* `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connection_status_last_change_time": {
				Description: "The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"connector_version": {
				Description: "The version of the device connector running on the managed device.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_external_ip_address": {
				Description: "The IP Address of the managed device as seen from Intersight at the time of registration.\nThis could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"leadership": {
				Description: "The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection.\n* `Unknown` - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight.\n* `Primary` - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role.\n* `Secondary` - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"locked_leader": {
				Description: "Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"member_identity": {
				Description: "The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parent_cluster_member_identity": {
				Description: "The member idenity of the cluster member through which this device is connected if applicable.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"proxy_app": {
				Description: "The name of the app which will proxy the messages to the device connector.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"additional_properties": {
					Type:             schema.TypeString,
					Optional:         true,
					DiffSuppressFunc: SuppressDiffAdditionProps,
				},
					"api_version": {
						Description: "The version of the connector API, describes the capability of the connector's framework.\nIf the version is lower than the current minimum supported version defined in the service managing the connection, the device connector will be connected with limited capabilities until the device connector is upgraded to a fully supported version. For example if a device connector that was released without delta inventory capabilities registers and connects to Intersight, inventory collection may be disabled until it has been upgraded.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"app_partition_number": {
						Description: "The partition number corresponding to the instance of the Proxy App which is managing the web-socket to the device connector.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"connection_id": {
						Description: "The unique identifier for the current connection. The identifier persists across network connectivity loss and is reset on device connector process restart or platform administrator toggle of the Intersight connectivity. The connectionId can be used by services that need to interact with stateful plugins running in the device connector process. For example if a service schedules an inventory in a devices job scheduler plugin at registration it is not necessary to reschedule the job if the device loses network connectivity due to an Intersight service upgrade or intermittent network issues in the devices datacenter.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"connection_reason": {
						Description: "If 'connectionStatus' is not equal to Connected, connectionReason provides further details about why the device is not connected with Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"connection_status": {
						Description: "The status of the persistent connection between the device connector and Intersight.\n* `` - The target details have been persisted but Intersight has not yet attempted to connect to the target.\n* `Connected` - Intersight is able to establish a connection to the target and initiate management activities.\n* `NotConnected` - Intersight is unable to establish a connection to the target.\n* `ClaimInProgress` - Claim of the target is in progress. A connection to the target has not been fully established.\n* `Unclaimed` - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.\n* `Claimed` - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"connection_status_last_change_time": {
						Description: "The last time at which the 'connectionStatus' property value changed. If connectionStatus is Connected, this time can be interpreted as the starting time since which a persistent connection has been maintained between Intersight and Device Connector. If connectionStatus is NotConnected, this time can be interpreted as the last time the device connector was connected with Intersight.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"connector_version": {
						Description: "The version of the device connector running on the managed device.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"device": {
						Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"device_external_ip_address": {
						Description: "The IP Address of the managed device as seen from Intersight at the time of registration.\nThis could be the IP address of the managed device's interface which has a route to the internet or a NAT IP addresss when the managed device is deployed in a private network.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"leadership": {
						Description: "The current leadershipstate of this member. Updated by the device connector on failover or leadership change. If a member is elected as Primary within the cluster this connection will be the same as the DeviceRegistration connection. E.g a message addressed to the DeviceRegistration will be forwarded to the ClusterMember connection.\n* `Unknown` - The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight.\n* `Primary` - The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role.\n* `Secondary` - The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"locked_leader": {
						Description: "Devices lock their leadership on failure to heartbeat with peers. Value acts as a third party tie breaker in election between nodes. Intersight enforces that only one cluster member exists with this value set to true.",
						Type:        schema.TypeBool,
						Optional:    true,
					},
					"member_identity": {
						Description: "The unique identity of the member within the cluster. The identity is retrieved from the platform and reported by the device connector at connection time.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"parent_cluster_member_identity": {
						Description: "The member idenity of the cluster member through which this device is connected if applicable.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"proxy_app": {
						Description: "The name of the app which will proxy the messages to the device connector.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"sudi": {
						Description: "The SUDI status read from the Trust Anchor Module on a device.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"pid": {
									Description: "The device model (PID) extracted from the X.509 SUDI Leaf Certificate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"serial_number": {
									Description: "The device SerialNumber extracted from the X.509 SUDI Leaf Certiicate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"signature": {
									Description: "The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certifiate.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"status": {
									Description: "The validation status of the device.\n* `DeviceStatusUnknown` - SUDI validation is done on the establishment of a connection. Before a device connects or after it disconnects, the SUDI validation status is set to this value.\n* `Verified` - The device returned a valid PID, Serial Number, Status and X.509 Leaf Certificate. The certificate signing chain was validated.\n* `CertificateValidationFailed` - Validation of the certificate signing chain failed.\n* `UnsupportedFirmware` - The firmware version of the Cisco IMC that is installed does not contain the SUDI APIs needed to perform validation.\n* `UnsupportedHardware` - The device is a model that does not contain a Trust Anchor Module (TAM) and thus cannot be validated.\n* `DeviceNotResponding` - An request was sent to the device, but no response was received.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"sudi_certificate": {
									Description: "The X.509 SUDI Leaf Certificate from the Trust Anchor Module. The certificate is serialized in PEM format (Base64 encoded DER certificate).",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"issuer": {
												Description: "The X.509 distinguished name of the issuer of this certificate.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Computed:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"common_name": {
															Description: "A required component that identifies a person or an object.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"country": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"locality": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"organization": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"organizational_unit": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"state": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"pem_certificate": {
												Description: "The base64 encoded certificate in PEM format.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"sha256_fingerprint": {
												Description: "The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"signature_algorithm": {
												Description: "Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"subject": {
												Description: "The X.509 distinguished name of the subject of this certificate.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Computed:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"common_name": {
															Description: "A required component that identifies a person or an object.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"country": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"locality": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
															Computed:    true,
														},
														"organization": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"organizational_unit": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"state": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
													},
												},
											},
										},
									},
									Computed: true,
								},
							},
						},
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceAssetClusterMemberRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.AssetClusterMember{}
	if v, ok := d.GetOk("api_version"); ok {
		x := int64(v.(int))
		o.SetApiVersion(x)
	}
	if v, ok := d.GetOk("app_partition_number"); ok {
		x := int64(v.(int))
		o.SetAppPartitionNumber(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("connection_id"); ok {
		x := (v.(string))
		o.SetConnectionId(x)
	}
	if v, ok := d.GetOk("connection_reason"); ok {
		x := (v.(string))
		o.SetConnectionReason(x)
	}
	if v, ok := d.GetOk("connection_status"); ok {
		x := (v.(string))
		o.SetConnectionStatus(x)
	}
	if v, ok := d.GetOk("connection_status_last_change_time"); ok {
		x, _ := time.Parse(v.(string), time.RFC1123)
		o.SetConnectionStatusLastChangeTime(x)
	}
	if v, ok := d.GetOk("connector_version"); ok {
		x := (v.(string))
		o.SetConnectorVersion(x)
	}
	if v, ok := d.GetOk("device_external_ip_address"); ok {
		x := (v.(string))
		o.SetDeviceExternalIpAddress(x)
	}
	if v, ok := d.GetOk("leadership"); ok {
		x := (v.(string))
		o.SetLeadership(x)
	}
	if v, ok := d.GetOk("locked_leader"); ok {
		x := (v.(bool))
		o.SetLockedLeader(x)
	}
	if v, ok := d.GetOk("member_identity"); ok {
		x := (v.(string))
		o.SetMemberIdentity(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("parent_cluster_member_identity"); ok {
		x := (v.(string))
		o.SetParentClusterMemberIdentity(x)
	}
	if v, ok := d.GetOk("proxy_app"); ok {
		x := (v.(string))
		o.SetProxyApp(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of AssetClusterMember object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.AssetApi.GetAssetClusterMemberList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of AssetClusterMember: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.AssetClusterMemberList.GetCount()
	var i int32
	var assetClusterMemberResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.AssetApi.GetAssetClusterMemberList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching AssetClusterMember: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.AssetClusterMemberList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for AssetClusterMember data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["api_version"] = (s.GetApiVersion())
				temp["app_partition_number"] = (s.GetAppPartitionNumber())
				temp["class_id"] = (s.GetClassId())
				temp["connection_id"] = (s.GetConnectionId())
				temp["connection_reason"] = (s.GetConnectionReason())
				temp["connection_status"] = (s.GetConnectionStatus())

				temp["connection_status_last_change_time"] = (s.GetConnectionStatusLastChangeTime()).String()
				temp["connector_version"] = (s.GetConnectorVersion())

				temp["device"] = flattenMapAssetDeviceRegistrationRelationship(s.GetDevice(), d)
				temp["device_external_ip_address"] = (s.GetDeviceExternalIpAddress())
				temp["leadership"] = (s.GetLeadership())
				temp["locked_leader"] = (s.GetLockedLeader())
				temp["member_identity"] = (s.GetMemberIdentity())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["parent_cluster_member_identity"] = (s.GetParentClusterMemberIdentity())
				temp["proxy_app"] = (s.GetProxyApp())

				temp["sudi"] = flattenMapAssetSudiInfo(s.GetSudi(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				assetClusterMemberResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(assetClusterMemberResults))
	if err := d.Set("results", assetClusterMemberResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(assetClusterMemberResults[0]["moid"].(string))
	return de
}
