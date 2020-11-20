package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIppoolShadowPool() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIppoolShadowPoolRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"assigned": {
				Description: "Number of IDs that are currently assigned.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"assignment_order": {
				Description: "Assignment order decides the order in which the next identifier is allocated.\n* `sequential` - Identifiers are assigned in a sequential order.\n* `default` - Assignment order is decided by the system.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_block_heads": {
				Description: "An array of relationships to ippoolShadowBlock resources.",
				Type:        schema.TypeList,
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
			"ip_v4_blocks": {
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
						"from": {
							Description: "First IPv4 address of the block.",
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
						"size": {
							Description: "Number of identifiers this block can hold.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"to": {
							Description: "Last IPv4 address of the block.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"ip_v4_config": {
				Description: "Netmask, Gateway and DNS settings for IPv4 addresses.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"gateway": {
							Description: "IP address of the default IPv4 gateway.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"netmask": {
							Description: "A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address.",
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
						"primary_dns": {
							Description: "IP Address of the primary Domain Name System (DNS) server.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"secondary_dns": {
							Description: "IP Address of the secondary Domain Name System (DNS) server.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"ip_v6_blocks": {
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
						"from": {
							Description: "First IPv6 address of the block.",
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
						"size": {
							Description: "Number of identifiers this block can hold.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"to": {
							Description: "Last IPv6 address of the block.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"ip_v6_config": {
				Description: "Netmask, Gateway and DNS settings for IPv6 addresses.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"gateway": {
							Description: "IP address of the default IPv6 gateway.",
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
						"prefix": {
							Description: "A prefix length which masks the  IP address and divides the IP address into network address and host address.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"primary_dns": {
							Description: "IP Address of the primary Domain Name System (DNS) server.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"secondary_dns": {
							Description: "IP Address of the secondary Domain Name System (DNS) server.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"pool": {
				Description: "A reference to a ippoolPool resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"size": {
				Description: "Total number of identifiers in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
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
			"v4_assigned": {
				Description: "Number of IPv4 addresses currently in use.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v4_size": {
				Description: "Number of IPv4 addresses in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v6_assigned": {
				Description: "Number of IPv6 addresses currently in use.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"v6_size": {
				Description: "Number of IPv6 addresses in this pool.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"vrf": {
				Description: "A reference to a vrfVrf resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
		},
	}
}

func dataSourceIppoolShadowPoolRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewIppoolShadowPoolWithDefaults()
	if v, ok := d.GetOk("assigned"); ok {
		x := int64(v.(int))
		o.SetAssigned(x)
	}
	if v, ok := d.GetOk("assignment_order"); ok {
		x := (v.(string))
		o.SetAssignmentOrder(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("size"); ok {
		x := int64(v.(int))
		o.SetSize(x)
	}
	if v, ok := d.GetOk("v4_assigned"); ok {
		x := int64(v.(int))
		o.SetV4Assigned(x)
	}
	if v, ok := d.GetOk("v4_size"); ok {
		x := int64(v.(int))
		o.SetV4Size(x)
	}
	if v, ok := d.GetOk("v6_assigned"); ok {
		x := int64(v.(int))
		o.SetV6Assigned(x)
	}
	if v, ok := d.GetOk("v6_size"); ok {
		x := int64(v.(int))
		o.SetV6Size(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	res, _, err := conn.ApiClient.IppoolApi.GetIppoolShadowPoolList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("error occurred while sending request %+v", err)
	}

	x, err := res.MarshalJSON()
	if err != nil {
		return fmt.Errorf("error occurred while marshalling response: %+v", err)
	}
	var s = &models.IppoolShadowPoolList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return fmt.Errorf("error occurred while unmarshalling response to IppoolShadowPool: %+v", err)
	}
	result := s.GetResults()
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewIppoolShadowPoolWithDefaults()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return fmt.Errorf("error occurred while unmarshalling result at index %+v: %+v", i, err)
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return fmt.Errorf("error occurred while setting property AdditionalProperties: %+v", err)
			}
			if err := d.Set("assigned", (s.GetAssigned())); err != nil {
				return fmt.Errorf("error occurred while setting property Assigned: %+v", err)
			}
			if err := d.Set("assignment_order", (s.GetAssignmentOrder())); err != nil {
				return fmt.Errorf("error occurred while setting property AssignmentOrder: %+v", err)
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return fmt.Errorf("error occurred while setting property ClassId: %+v", err)
			}
			if err := d.Set("description", (s.GetDescription())); err != nil {
				return fmt.Errorf("error occurred while setting property Description: %+v", err)
			}

			if err := d.Set("ip_block_heads", flattenListIppoolShadowBlockRelationship(s.GetIpBlockHeads(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property IpBlockHeads: %+v", err)
			}

			if err := d.Set("ip_v4_blocks", flattenListIppoolIpV4Block(s.GetIpV4Blocks(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property IpV4Blocks: %+v", err)
			}

			if err := d.Set("ip_v4_config", flattenMapIppoolIpV4Config(s.GetIpV4Config(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property IpV4Config: %+v", err)
			}

			if err := d.Set("ip_v6_blocks", flattenListIppoolIpV6Block(s.GetIpV6Blocks(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property IpV6Blocks: %+v", err)
			}

			if err := d.Set("ip_v6_config", flattenMapIppoolIpV6Config(s.GetIpV6Config(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property IpV6Config: %+v", err)
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return fmt.Errorf("error occurred while setting property Moid: %+v", err)
			}
			if err := d.Set("name", (s.GetName())); err != nil {
				return fmt.Errorf("error occurred while setting property Name: %+v", err)
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return fmt.Errorf("error occurred while setting property ObjectType: %+v", err)
			}

			if err := d.Set("pool", flattenMapIppoolPoolRelationship(s.GetPool(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Pool: %+v", err)
			}
			if err := d.Set("size", (s.GetSize())); err != nil {
				return fmt.Errorf("error occurred while setting property Size: %+v", err)
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Tags: %+v", err)
			}
			if err := d.Set("v4_assigned", (s.GetV4Assigned())); err != nil {
				return fmt.Errorf("error occurred while setting property V4Assigned: %+v", err)
			}
			if err := d.Set("v4_size", (s.GetV4Size())); err != nil {
				return fmt.Errorf("error occurred while setting property V4Size: %+v", err)
			}
			if err := d.Set("v6_assigned", (s.GetV6Assigned())); err != nil {
				return fmt.Errorf("error occurred while setting property V6Assigned: %+v", err)
			}
			if err := d.Set("v6_size", (s.GetV6Size())); err != nil {
				return fmt.Errorf("error occurred while setting property V6Size: %+v", err)
			}

			if err := d.Set("vrf", flattenMapVrfVrfRelationship(s.GetVrf(), d)); err != nil {
				return fmt.Errorf("error occurred while setting property Vrf: %+v", err)
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
