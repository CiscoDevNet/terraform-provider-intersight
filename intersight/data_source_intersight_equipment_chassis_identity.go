package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceEquipmentChassisIdentity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEquipmentChassisIdentityRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"admin_action": {
				Description: "Updated by UI/API to trigger specific chassis action type.\n* `None` - No operation value for maintenance actions on an equipment.\n* `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight.\n* `Recommission` - Recommission the equipment.\n* `Reack` - Reacknowledge the equipment and discover it again.\n* `Remove` - Remove the equipment permanently from Intersight management.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"admin_action_state": {
				Description: "The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed.\n* `None` - Nil value when no action has been triggered by the user.\n* `Applied` - User configured settings are in applied state.\n* `Applying` - User settings are being applied on the target server.\n* `Failed` - User configured settings could not be applied.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"chassis": {
				Description: "A reference to a equipmentChassis resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"device_registration": {
				Description: "A reference to a assetDeviceRegistration resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
				Computed: true,
			},
			"identifier": {
				Description: "Numeric Identifier assigned by the management system to the equipment.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"io_card_identity_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"io_card_moid": {
							Description: "MO Reference to equipmentIoCard MO in inventory service.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"module_id": {
							Description: "IOM/MUX Module ID connected to the FI.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"network_element_moid": {
							Description: "MO Reference to networkElement MO in inventory service.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"switch_id": {
							Description: "Switch ID to which IOM is connected, ID can be either 1 or 2, equalent to A or B.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"nr_lifecycle": {
				Description: "The equipment's lifecycle status.\n* `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment.\n* `Active` - Default Lifecycle State for a physical entity.\n* `Decommissioned` - Decommission Lifecycle state.\n* `DecommissionInProgress` - Decommission Inprogress Lifecycle state.\n* `RecommissionInProgress` - Recommission Inprogress Lifecycle state.\n* `OperationFailed` - Failed Operation Lifecycle state.\n* `ReackInProgress` - ReackInProgress Lifecycle state.\n* `RemoveInProgress` - RemoveInProgress Lifecycle state.\n* `Discovered` - Discovered Lifecycle state.\n* `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state.\n* `DiscoveryFailed` - DiscoveryFailed Lifecycle state.\n* `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"model": {
				Description: "The vendor provided model name for the equipment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Description: "The serial number of the equipment.",
				Type:        schema.TypeString,
				Optional:    true,
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
			"vendor": {
				Description: "The manufacturer of the equipment.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceEquipmentChassisIdentityRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = &models.EquipmentChassisIdentity{}
	if v, ok := d.GetOk("admin_action"); ok {
		x := (v.(string))
		o.SetAdminAction(x)
	}
	if v, ok := d.GetOk("admin_action_state"); ok {
		x := (v.(string))
		o.SetAdminActionState(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("identifier"); ok {
		x := int64(v.(int))
		o.SetIdentifier(x)
	}
	if v, ok := d.GetOk("nr_lifecycle"); ok {
		x := (v.(string))
		o.SetLifecycle(x)
	}
	if v, ok := d.GetOk("model"); ok {
		x := (v.(string))
		o.SetModel(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("serial"); ok {
		x := (v.(string))
		o.SetSerial(x)
	}
	if v, ok := d.GetOk("vendor"); ok {
		x := (v.(string))
		o.SetVendor(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.EquipmentApi.GetEquipmentChassisIdentityList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.EquipmentChassisIdentityList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to EquipmentChassisIdentity: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.EquipmentChassisIdentity{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("admin_action", (s.GetAdminAction())); err != nil {
				return fmt.Errorf("error occurred while setting property AdminAction: %+v", err)
			}
			if err := d.Set("admin_action_state", (s.GetAdminActionState())); err != nil {
				return fmt.Errorf("error occurred while setting property AdminActionState: %+v", err)
			}

			if err := d.Set("chassis", flattenMapEquipmentChassisRelationship(s.GetChassis(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Chassis: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}

			if err := d.Set("device_registration", flattenMapAssetDeviceRegistrationRelationship(s.GetDeviceRegistration(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property DeviceRegistration: %+v", err)
			}
			if err := d.Set("identifier", (s.GetIdentifier())); err != nil {
				return fmt.Errorf("error occurred while setting property Identifier: %+v", err)
			}

			if err := d.Set("io_card_identity_list", flattenListEquipmentIoCardIdentity(s.GetIoCardIdentityList(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property IoCardIdentityList: %+v", err)
			}
			if err := d.Set("nr_lifecycle", (s.GetLifecycle())); err != nil {
				return fmt.Errorf("error occurred while setting property Lifecycle: %+v", err)
			}
			if err := d.Set("model", (s.GetModel())); err != nil {
				return fmt.Errorf("error occurred while setting property Model: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}
			if err := d.Set("serial", (s.GetSerial())); err != nil {
				return fmt.Errorf("error occurred while setting property Serial: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("vendor", (s.GetVendor())); err != nil {
				return fmt.Errorf("error occurred while setting property Vendor: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
