// Code generated by "enumer -type InstallPhase -output zz_generated_installphase_enumer.go"; DO NOT EDIT.

//
package admin

import (
	"fmt"
)

const _InstallPhaseName = "InstallPhaseBootstrapInstallPhaseRemoveBootstrap"

var _InstallPhaseIndex = [...]uint8{0, 21, 48}

func (i InstallPhase) String() string {
	if i < 0 || i >= InstallPhase(len(_InstallPhaseIndex)-1) {
		return fmt.Sprintf("InstallPhase(%d)", i)
	}
	return _InstallPhaseName[_InstallPhaseIndex[i]:_InstallPhaseIndex[i+1]]
}

var _InstallPhaseValues = []InstallPhase{0, 1}

var _InstallPhaseNameToValueMap = map[string]InstallPhase{
	_InstallPhaseName[0:21]:  0,
	_InstallPhaseName[21:48]: 1,
}

// InstallPhaseString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func InstallPhaseString(s string) (InstallPhase, error) {
	if val, ok := _InstallPhaseNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to InstallPhase values", s)
}

// InstallPhaseValues returns all values of the enum
func InstallPhaseValues() []InstallPhase {
	return _InstallPhaseValues
}

// IsAInstallPhase returns "true" if the value is listed in the enum definition. "false" otherwise
func (i InstallPhase) IsAInstallPhase() bool {
	for _, v := range _InstallPhaseValues {
		if i == v {
			return true
		}
	}
	return false
}
