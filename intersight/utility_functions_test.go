package intersight

import (
	"reflect"
	"testing"
)

func TestReorderPolicyBucket(t *testing.T) {
	tests := []struct {
		name     string
		old      []interface{}
		new      []interface{}
		expected []interface{}
	}{
		{
			name: "No changes",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Add new policy",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Remove policy",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Reorder policies",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "snmp1",
					"object_type": "snmp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Empty old array",
			old:  []interface{}{},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
		{
			name: "Empty new array",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new:      []interface{}{},
			expected: nil,
		},
		{
			name: "Modify existing policy",
			old: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			new: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "modified",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"moid":        "ntp1",
					"object_type": "ntp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "modified",
				},
				map[string]interface{}{
					"moid":        "smtp1",
					"object_type": "smtp.Policy",
					"class_id":    "mo.MoRef",
					"selector":    "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReorderPolicyBucket(tt.old, tt.new)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReorderPolicyBucket() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRecursiveValueCheck(t *testing.T) {
	const (
		oldM  = "oldM"
		key   = "key"
		value = "value"
	)
	tests := []struct {
		name   string
		input  map[string]interface{}
		output bool
	}{
		{
			name: "No change in simple values",
			input: map[string]interface{}{
				"oldM":  map[string]interface{}{"object_type": "snmp.User"},
				"key":   "object_type",
				"value": "snmp.User",
			},
			output: true,
		},
		{
			name: "Different simple values",
			input: map[string]interface{}{
				"oldM":  map[string]interface{}{"object_type": "snmp.User"},
				"key":   "object_type",
				"value": "snmp.Trap",
			},
			output: false,
		},
		{
			name: "Complex map comparison",
			input: map[string]interface{}{
				"oldM": map[string]interface{}{
					"GpuControllersRange": map[string]interface{}{
						"ConditionType": "RANGE",
						"MaxValue":      0,
						"MinValue":      0,
						"ObjectType":    "GpuControllersRangeFilter",
					}},
				"key": "GpuControllersRange",
				"value": map[string]interface{}{
					"ConditionType": "RANGE",
					"MaxValue":      0,
					"MinValue":      0,
					"ObjectType":    "GpuControllersRangeFilter",
				},
			},
			output: true,
		},
		{
			name: "List comparison",
			input: map[string]interface{}{
				"oldM": map[string]interface{}{
					"ChassisAndSlotIdRange": []interface{}{
						map[string]interface{}{
							"ChassisIdRange": map[string]interface{}{
								"ClassId":       "resource.ChassisIdRangeFilter",
								"ConditionType": "RANGE",
								"MaxValue":      5,
								"MinValue":      1,
								"ObjectType":    "resource.ChassisIdRangeFilter",
							},
							"ObjectType": "resource.ChassisAndSlotQualification",
							"SlotIdRanges": []interface{}{
								map[string]interface{}{
									"ClassId":       "resource.SlotIdRangeFilter",
									"ConditionType": "RANGE",
									"MaxValue":      3,
									"MinValue":      1,
									"ObjectType":    "resource.SlotIdRangeFilter",
								},
								map[string]interface{}{
									"ClassId":       "resource.SlotIdRangeFilter",
									"ConditionType": "RANGE",
									"MaxValue":      6,
									"MinValue":      5,
									"ObjectType":    "resource.SlotIdRangeFilter",
								},
							},
						},
					},
				},
				"key": "ChassisAndSlotIdRange",
				"value": []interface{}{
					map[string]interface{}{
						"ChassisIdRange": map[string]interface{}{
							"ClassId":       "resource.ChassisIdRangeFilter",
							"ConditionType": "RANGE",
							"MaxValue":      5,
							"MinValue":      1,
							"ObjectType":    "resource.ChassisIdRangeFilter",
						},
						"ObjectType": "resource.ChassisAndSlotQualification",
						"SlotIdRanges": []interface{}{
							map[string]interface{}{
								"ClassId":       "resource.SlotIdRangeFilter",
								"ConditionType": "RANGE",
								"MaxValue":      3,
								"MinValue":      1,
								"ObjectType":    "resource.SlotIdRangeFilter",
							},
							map[string]interface{}{
								"ClassId":       "resource.SlotIdRangeFilter",
								"ConditionType": "RANGE",
								"MaxValue":      6,
								"MinValue":      5,
								"ObjectType":    "resource.SlotIdRangeFilter",
							},
						},
					},
				},
			},
			output: true,
		},
		{
			name: "List mismatch",
			input: map[string]interface{}{
				"oldM": map[string]interface{}{
					"SlotIdRanges": []interface{}{
						map[string]interface{}{
							"ClassId":       "resource.SlotIdRangeFilter",
							"ConditionType": "RANGE",
							"MaxValue":      3,
							"MinValue":      1,
							"ObjectType":    "resource.SlotIdRangeFilter",
						},
						map[string]interface{}{
							"ClassId":       "resource.SlotIdRangeFilter",
							"ConditionType": "RANGE",
							"MaxValue":      7,
							"MinValue":      4,
							"ObjectType":    "resource.SlotIdRangeFilter",
						},
					},
				},
				"key": "SlotIdRanges",
				"value": []interface{}{
					map[string]interface{}{
						"ClassId":       "resource.SlotIdRangeFilter",
						"ConditionType": "RANGE",
						"MaxValue":      3,
						"MinValue":      3,
						"ObjectType":    "resource.SlotIdRangeFilter",
					},
					map[string]interface{}{
						"ClassId":       "resource.SlotIdRangeFilter",
						"ConditionType": "RANGE",
						"MaxValue":      8,
						"MinValue":      4,
						"ObjectType":    "resource.SlotIdRangeFilter",
					},
				},
			},
			output: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := recursiveValueCheck(
				tt.input["oldM"].(map[string]interface{}),
				tt.input["key"].(string),
				tt.input["value"])
			if result != tt.output {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.output, result)
			}
		})
	}
}
