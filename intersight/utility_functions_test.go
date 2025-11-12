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

func TestParseVlanRange(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
		hasError bool
	}{
		{
			name:     "Single VLAN",
			input:    "1140",
			expected: []int{1140},
			hasError: false,
		},
		{
			name:     "VLAN Range",
			input:    "1140-1142",
			expected: []int{1140, 1141, 1142},
			hasError: false,
		},
		{
			name:     "Single VLAN Range",
			input:    "5-7",
			expected: []int{5, 6, 7},
			hasError: false,
		},
		{
			name:     "Same start and end",
			input:    "10-10",
			expected: []int{10},
			hasError: false,
		},
		{
			name:     "Invalid range - start greater than end",
			input:    "10-5",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid format - multiple hyphens",
			input:    "1-5-10",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid format - non-numeric",
			input:    "abc-def",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid single VLAN",
			input:    "abc",
			expected: nil,
			hasError: true,
		},
		{
			name:     "VLAN range with start below valid range",
			input:    "0-5",
			expected: nil,
			hasError: true,
		},
		{
			name:     "VLAN range with end above valid range",
			input:    "4090-4094",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Valid edge case - maximum VLAN",
			input:    "4093",
			expected: []int{4093},
			hasError: false,
		},
		{
			name:     "Range with negative start",
			input:    "-5-10",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseVlanRange(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input '%s', but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input '%s': %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("For input '%s', expected %v, got %v", tt.input, tt.expected, result)
				}
			}
		})
	}
}

func TestParseVlanString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
		hasError bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []int{},
			hasError: false,
		},
		{
			name:     "Single VLAN",
			input:    "1140",
			expected: []int{1140},
			hasError: false,
		},
		{
			name:     "Multiple comma-separated VLANs",
			input:    "1,5,7,10",
			expected: []int{1, 5, 7, 10},
			hasError: false,
		},
		{
			name:     "VLAN range with hyphen",
			input:    "1140-1142",
			expected: []int{1140, 1141, 1142},
			hasError: false,
		},
		{
			name:     "Mixed comma and hyphen separated",
			input:    "1,5-7,10",
			expected: []int{1, 5, 6, 7, 10},
			hasError: false,
		},
		{
			name:     "Unsorted input should be sorted",
			input:    "10,1,5,3",
			expected: []int{1, 3, 5, 10},
			hasError: false,
		},
		{
			name:     "Duplicates should be removed",
			input:    "1,2,2,3,1",
			expected: []int{1, 2, 3},
			hasError: false,
		},
		{
			name:     "Complex mixed ranges with duplicates",
			input:    "1,3-5,4,7-9,8",
			expected: []int{1, 3, 4, 5, 7, 8, 9},
			hasError: false,
		},
		{
			name:     "Spaces should be handled",
			input:    " 1 , 5-7 , 10 ",
			expected: []int{1, 5, 6, 7, 10},
			hasError: false,
		},
		{
			name:     "Invalid VLAN in range",
			input:    "1,abc,5",
			expected: nil,
			hasError: true,
		},
		{
			name:     "Invalid range format",
			input:    "1,5-7-9,10",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseVlanString(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error for input '%s', but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input '%s': %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("For input '%s', expected %v, got %v", tt.input, tt.expected, result)
				}
			}
		})
	}
}

func TestCompareVlanRanges(t *testing.T) {
	tests := []struct {
		name     string
		oldVlans string
		newVlans string
		expected bool
	}{
		{
			name:     "Both empty",
			oldVlans: "",
			newVlans: "",
			expected: true,
		},
		{
			name:     "Identical strings",
			oldVlans: "1140,1141",
			newVlans: "1140,1141",
			expected: true,
		},
		{
			name:     "Comma separated vs hyphen range - matching",
			oldVlans: "1140,1141",
			newVlans: "1140-1141",
			expected: true,
		},
		{
			name:     "Multiple ranges matching",
			oldVlans: "1,5,7,10",
			newVlans: "1,5-7,10",
			expected: false, // 5-7 expands to 5,6,7, but old only has 1,5,7,10
		},
		{
			name:     "Correct multiple ranges matching",
			oldVlans: "1,5,6,7,10",
			newVlans: "1,5-7,10",
			expected: true,
		},
		{
			name:     "New range added - should not match",
			oldVlans: "1,2,3,4",
			newVlans: "1,2,3,4-5",
			expected: false, // 4-5 expands to 4,5, but old only has 1,2,3,4
		},
		{
			name:     "Correct new range - should match",
			oldVlans: "1,2,3,4,5",
			newVlans: "1,2,3,4-5",
			expected: true,
		},
		{
			name:     "New intermediate range added",
			oldVlans: "1,2,5,6",
			newVlans: "1,2,3-4,5,6",
			expected: false, // 3-4 adds new VLANs 3,4
		},
		{
			name:     "Different order but same VLANs",
			oldVlans: "10,1,5,3",
			newVlans: "1,3,5,10",
			expected: true,
		},
		{
			name:     "Complex matching ranges",
			oldVlans: "1,3,4,5,7,8,9",
			newVlans: "1,3-5,7-9",
			expected: true,
		},
		{
			name:     "Complex non-matching ranges",
			oldVlans: "1,3,5,7,9",
			newVlans: "1,3-5,7-9",
			expected: false, // 3-5 expands to 3,4,5 and 7-9 to 7,8,9, but old doesn't have 4,8
		},
		{
			name:     "Single VLAN vs single range",
			oldVlans: "100",
			newVlans: "100-100",
			expected: true,
		},
		{
			name:     "Empty old vs non-empty new",
			oldVlans: "",
			newVlans: "1,2,3",
			expected: false,
		},
		{
			name:     "Non-empty old vs empty new",
			oldVlans: "1,2,3",
			newVlans: "",
			expected: false,
		},
		// Reverse test cases (hyphen to comma)
		{
			name:     "Hyphen range vs comma separated - matching",
			oldVlans: "1140-1141",
			newVlans: "1140,1141",
			expected: true,
		},
		{
			name:     "Hyphen range vs comma separated - matching complex",
			oldVlans: "1,5-7,10",
			newVlans: "1,5,6,7,10",
			expected: true,
		},
		{
			name:     "Hyphen range vs comma separated - not matching",
			oldVlans: "1,5-7,10",
			newVlans: "1,5,7,10",
			expected: false, // old has 1,5,6,7,10 but new missing 6
		},
		{
			name:     "Single range vs single VLAN",
			oldVlans: "100-100",
			newVlans: "100",
			expected: true,
		},
		// Complex reverse test cases
		{
			name:     "Multiple ranges vs comma list - matching",
			oldVlans: "1-3,5-7,10-12",
			newVlans: "1,2,3,5,6,7,10,11,12",
			expected: true,
		},
		{
			name:     "Multiple ranges vs comma list - not matching (missing VLAN)",
			oldVlans: "1-3,5-7,10-12",
			newVlans: "1,2,5,6,7,10,11,12",
			expected: false, // new missing 3
		},
		{
			name:     "Mixed ranges vs comma - matching",
			oldVlans: "1,3-5,8,10-12",
			newVlans: "1,3,4,5,8,10,11,12",
			expected: true,
		},
		{
			name:     "Mixed ranges vs comma - not matching (extra VLAN in old)",
			oldVlans: "1,3-5,8,10-13",
			newVlans: "1,3,4,5,8,10,11,12",
			expected: false, // old has 13, new doesn't
		},
		{
			name:     "Large range vs comma expansion - matching",
			oldVlans: "100-105",
			newVlans: "100,101,102,103,104,105",
			expected: true,
		},
		{
			name:     "Large range vs comma expansion - not matching (gap in new)",
			oldVlans: "100-105",
			newVlans: "100,101,103,104,105",
			expected: false, // new missing 102
		},
		{
			name:     "Overlapping ranges vs comma - matching",
			oldVlans: "1-5,3-7",
			newVlans: "1,2,3,4,5,6,7",
			expected: true,
		},
		{
			name:     "Overlapping ranges vs comma - not matching (extra in new)",
			oldVlans: "1-5,3-7",
			newVlans: "1,2,3,4,5,6,7,8",
			expected: false, // new has extra 8
		},
		{
			name:     "Complex mixed scenario - matching",
			oldVlans: "1-2,5,7-9,12,15-17",
			newVlans: "1,2,5,7,8,9,12,15,16,17",
			expected: true,
		},
		{
			name:     "Complex mixed scenario - not matching (gap in new)",
			oldVlans: "1-2,5,7-9,12,15-17",
			newVlans: "1,2,5,7,9,12,15,16,17",
			expected: false, // new missing 8
		},
		{
			name:     "Consecutive ranges vs comma - matching",
			oldVlans: "10-12,13-15,16-18",
			newVlans: "10,11,12,13,14,15,16,17,18",
			expected: true,
		},
		{
			name:     "Consecutive ranges vs comma - not matching (extra in old)",
			oldVlans: "10-12,13-15,16-19",
			newVlans: "10,11,12,13,14,15,16,17,18",
			expected: false, // old has 19, new doesn't
		},
		{
			name:     "Single VLAN mixed with ranges vs comma - matching",
			oldVlans: "1,3-5,7,9-11,13",
			newVlans: "1,3,4,5,7,9,10,11,13",
			expected: true,
		},
		{
			name:     "Single VLAN mixed with ranges vs comma - not matching",
			oldVlans: "1,3-5,7,9-11,13",
			newVlans: "1,3,4,7,9,10,11,13",
			expected: false, // new missing 5
		},
		{
			name:     "Non-contiguous ranges vs comma - matching",
			oldVlans: "20-22,25-27,30-32",
			newVlans: "20,21,22,25,26,27,30,31,32",
			expected: true,
		},
		{
			name:     "Non-contiguous ranges vs comma - not matching (reordered)",
			oldVlans: "20-22,25-27,30-32",
			newVlans: "20,21,25,26,27,30,31,32",
			expected: false, // new missing 22
		},
		{
			name:     "Edge case: single digit ranges vs comma - matching",
			oldVlans: "1-1,3-3,5-5",
			newVlans: "1,3,5",
			expected: true,
		},
		{
			name:     "Complex unordered ranges vs ordered comma - matching",
			oldVlans: "50-52,10-12,30-32",
			newVlans: "10,11,12,30,31,32,50,51,52",
			expected: true,
		},
		{
			name:     "Sparse ranges vs dense comma - not matching",
			oldVlans: "100-102,105-107,110-112",
			newVlans: "100,101,102,103,104,105,106,107,108,109,110,111,112",
			expected: false, // new has extra 103,104,108,109
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareVlanRanges(tt.oldVlans, tt.newVlans)
			if result != tt.expected {
				t.Errorf("For compareVlanRanges('%s', '%s'), expected %v, got %v",
					tt.oldVlans, tt.newVlans, tt.expected, result)
			}
		})
	}
}

func TestSuppressDiffVlanRanges(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		oldValue string
		newValue string
		expected bool
	}{
		{
			name:     "Both empty should suppress diff",
			key:      "allowed_vlans",
			oldValue: "",
			newValue: "",
			expected: true,
		},
		{
			name:     "Identical values should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1140,1141",
			newValue: "1140,1141",
			expected: true,
		},
		{
			name:     "Comma vs hyphen range matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1140,1141",
			newValue: "1140-1141",
			expected: true,
		},
		{
			name:     "Non-matching ranges should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "1,2,3,4",
			newValue: "1,2,3,4-5",
			expected: false,
		},
		{
			name:     "Complex matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1,5,6,7,10",
			newValue: "1,5-7,10",
			expected: true,
		},
		{
			name:     "Different VLANs should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "1,2,5,6",
			newValue: "1,2,3-4,5,6",
			expected: false,
		},
		// Complex reverse scenarios for SuppressDiff
		{
			name:     "Hyphen to comma matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1140-1141",
			newValue: "1140,1141",
			expected: true,
		},
		{
			name:     "Multiple ranges to comma matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1-3,5-7,10-12",
			newValue: "1,2,3,5,6,7,10,11,12",
			expected: true,
		},
		{
			name:     "Multiple ranges to comma not matching should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "1-3,5-7,10-12",
			newValue: "1,2,5,6,7,10,11,12",
			expected: false, // new missing 3
		},
		{
			name:     "Large range to comma expansion matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "100-105",
			newValue: "100,101,102,103,104,105",
			expected: true,
		},
		{
			name:     "Large range to comma expansion not matching should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "100-105",
			newValue: "100,101,103,104,105",
			expected: false, // new missing 102
		},
		{
			name:     "Mixed ranges to comma matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1,3-5,8,10-12",
			newValue: "1,3,4,5,8,10,11,12",
			expected: true,
		},
		{
			name:     "Mixed ranges to comma not matching should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "1,3-5,8,10-13",
			newValue: "1,3,4,5,8,10,11,12",
			expected: false, // old has 13, new doesn't
		},
		{
			name:     "Overlapping ranges to comma matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1-5,3-7",
			newValue: "1,2,3,4,5,6,7",
			expected: true,
		},
		{
			name:     "Complex mixed scenario matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "1-2,5,7-9,12,15-17",
			newValue: "1,2,5,7,8,9,12,15,16,17",
			expected: true,
		},
		{
			name:     "Complex mixed scenario not matching should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "1-2,5,7-9,12,15-17",
			newValue: "1,2,5,7,9,12,15,16,17",
			expected: false, // new missing 8
		},
		{
			name:     "Bidirectional complex matching should suppress diff",
			key:      "allowed_vlans",
			oldValue: "20-22,25-27,30-32",
			newValue: "20,21,22,25,26,27,30,31,32",
			expected: true,
		},
		{
			name:     "Bidirectional complex not matching should not suppress diff",
			key:      "allowed_vlans",
			oldValue: "20-22,25-27,30-32",
			newValue: "20,21,25,26,27,30,31,32",
			expected: false, // new missing 22
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuppressDiffVlanRanges(tt.key, tt.oldValue, tt.newValue, nil)
			if result != tt.expected {
				t.Errorf("For SuppressDiffVlanRanges('%s', '%s', '%s'), expected %v, got %v",
					tt.key, tt.oldValue, tt.newValue, tt.expected, result)
			}
		})
	}
}
