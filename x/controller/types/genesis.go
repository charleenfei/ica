package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CmpDataList:              []CmpData{},
		CmpControllerRequestList: []CmpControllerRequest{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in cmpData
	cmpDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.CmpDataList {
		index := string(CmpDataKey(elem.Index))
		if _, ok := cmpDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for cmpData")
		}
		cmpDataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in cmpControllerRequest
	cmpControllerRequestIndexMap := make(map[string]struct{})

	for _, elem := range gs.CmpControllerRequestList {
		index := string(CmpControllerRequestKey(elem.Index))
		if _, ok := cmpControllerRequestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for cmpControllerRequest")
		}
		cmpControllerRequestIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
