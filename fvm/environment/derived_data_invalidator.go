package environment

import (
	"github.com/onflow/cadence/runtime/common"

	"github.com/onflow/flow-go/fvm/storage/derived"
	"github.com/onflow/flow-go/fvm/storage/state"
	"github.com/onflow/flow-go/model/flow"
)

type ContractUpdate struct {
	Location common.AddressLocation
	Code     []byte
}

type ContractUpdates struct {
	Updates   []common.AddressLocation
	Deploys   []common.AddressLocation
	Deletions []common.AddressLocation
}

func (u ContractUpdates) Any() bool {
	return len(u.Updates) > 0 || len(u.Deploys) > 0 || len(u.Deletions) > 0
}

type DerivedDataInvalidator struct {
	ContractUpdates

	MeterParamOverridesUpdated bool
}

var _ derived.TransactionInvalidator = DerivedDataInvalidator{}

// TODO(patrick): extract contractKeys from executionSnapshot
func NewDerivedDataInvalidator(
	contractUpdates ContractUpdates,
	serviceAddress flow.Address,
	executionSnapshot *state.ExecutionSnapshot,
) DerivedDataInvalidator {
	return DerivedDataInvalidator{
		ContractUpdates: contractUpdates,
		MeterParamOverridesUpdated: meterParamOverridesUpdated(
			serviceAddress,
			executionSnapshot),
	}
}

func meterParamOverridesUpdated(
	serviceAddress flow.Address,
	executionSnapshot *state.ExecutionSnapshot,
) bool {
	serviceAccount := string(serviceAddress.Bytes())
	storageDomain := common.PathDomainStorage.Identifier()

	for registerId := range executionSnapshot.WriteSet {
		// The meter param override values are stored in the service account.
		if registerId.Owner != serviceAccount {
			continue
		}

		// NOTE: This condition is empirically generated by running the
		// MeterParamOverridesComputer to capture touched registers.
		//
		// The paramater settings are stored as regular fields in the service
		// account.  In general, each account's regular fields are stored in
		// ordered map known only to cadence.  Cadence encodes this map into
		// bytes and split the bytes into slab chunks before storing the slabs
		// into the ledger.  Hence any changes to the stabs indicate changes
		// the ordered map.
		//
		// The meter param overrides use storageDomain as input, so any
		// changes to it must also invalidate the values.
		if registerId.Key == storageDomain || registerId.IsSlabIndex() {
			return true
		}
	}

	return false
}

func (invalidator DerivedDataInvalidator) ProgramInvalidator() derived.ProgramInvalidator {
	return ProgramInvalidator{invalidator}
}

func (invalidator DerivedDataInvalidator) MeterParamOverridesInvalidator() derived.MeterParamOverridesInvalidator {
	return MeterParamOverridesInvalidator{invalidator}
}

type ProgramInvalidator struct {
	DerivedDataInvalidator
}

func (invalidator ProgramInvalidator) ShouldInvalidateEntries() bool {
	return invalidator.MeterParamOverridesUpdated ||
		invalidator.ContractUpdates.Any()
}

func (invalidator ProgramInvalidator) ShouldInvalidateEntry(
	location common.AddressLocation,
	program *derived.Program,
	snapshot *state.ExecutionSnapshot,
) bool {
	if invalidator.MeterParamOverridesUpdated {
		// if meter parameters changed we need to invalidate all programs
		return true
	}

	// invalidate all programs depending on any of the contracts that were
	// updated. A program has itself listed as a dependency, so that this
	// simpler.
	for _, loc := range invalidator.ContractUpdates.Updates {
		ok := program.Dependencies.ContainsLocation(loc)
		if ok {
			return true
		}
	}

	// In case a contract was deployed or removed from an address,
	// we need to invalidate all programs depending on that address.
	for _, loc := range invalidator.ContractUpdates.Deploys {
		ok := program.Dependencies.ContainsAddress(loc.Address)
		if ok {
			return true
		}
	}
	for _, loc := range invalidator.ContractUpdates.Deletions {
		ok := program.Dependencies.ContainsAddress(loc.Address)
		if ok {
			return true
		}
	}

	return false
}

type MeterParamOverridesInvalidator struct {
	DerivedDataInvalidator
}

func (invalidator MeterParamOverridesInvalidator) ShouldInvalidateEntries() bool {
	return invalidator.MeterParamOverridesUpdated
}

func (invalidator MeterParamOverridesInvalidator) ShouldInvalidateEntry(
	_ struct{},
	_ derived.MeterParamOverrides,
	_ *state.ExecutionSnapshot,
) bool {
	return invalidator.MeterParamOverridesUpdated
}
