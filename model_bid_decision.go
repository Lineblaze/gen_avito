/*
Tender Management API

API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка). 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BidDecision Решение по предложению
type BidDecision string

// List of bidDecision
const (
	BidDecisionAPPROVED BidDecision = "Approved"
	BidDecisionREJECTED BidDecision = "Rejected"
)

// All allowed values of BidDecision enum
var AllowedBidDecisionEnumValues = []BidDecision{
	"Approved",
	"Rejected",
}

func (v *BidDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BidDecision(value)
	for _, existing := range AllowedBidDecisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BidDecision", value)
}

// NewBidDecisionFromValue returns a pointer to a valid BidDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBidDecisionFromValue(v string) (*BidDecision, error) {
	ev := BidDecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BidDecision: valid values are %v", v, AllowedBidDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BidDecision) IsValid() bool {
	for _, existing := range AllowedBidDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bidDecision value
func (v BidDecision) Ptr() *BidDecision {
	return &v
}

type NullableBidDecision struct {
	value *BidDecision
	isSet bool
}

func (v NullableBidDecision) Get() *BidDecision {
	return v.value
}

func (v *NullableBidDecision) Set(val *BidDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableBidDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableBidDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidDecision(val *BidDecision) *NullableBidDecision {
	return &NullableBidDecision{value: val, isSet: true}
}

func (v NullableBidDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBidDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

