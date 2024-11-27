package accountV2Migration

import (
	_ "embed"

	"github.com/onflow/cadence/common"
	cadenceErrors "github.com/onflow/cadence/errors"
	"github.com/onflow/cadence/interpreter"
	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/sema"
	"github.com/onflow/cadence/stdlib"

	"github.com/onflow/flow-go/fvm/errors"
	"github.com/onflow/flow-go/model/flow"
)

//go:embed AccountV2Migration.cdc
var ContractCode []byte

const ContractName = "AccountV2Migration"

const scheduleAccountV2MigrationFunctionName = "scheduleAccountV2Migration"

// scheduleAccountV2MigrationType is the type of the `scheduleAccountV2Migration` function.
// This defines the signature as `func(addressStartIndex: UInt64, count: UInt64): Bool`
var scheduleAccountV2MigrationType = &sema.FunctionType{
	Parameters: []sema.Parameter{
		{
			Identifier:     "addressStartIndex",
			TypeAnnotation: sema.UInt64TypeAnnotation,
		},
		{
			Identifier:     "count",
			TypeAnnotation: sema.UInt64TypeAnnotation,
		},
	},
	ReturnTypeAnnotation: sema.NewTypeAnnotation(sema.BoolType),
}

func DeclareScheduleAccountV2MigrationFunction(environment runtime.Environment, chainID flow.ChainID) {

	functionType := scheduleAccountV2MigrationType

	functionValue := stdlib.StandardLibraryValue{
		Name: scheduleAccountV2MigrationFunctionName,
		Type: functionType,
		Kind: common.DeclarationKindFunction,
		Value: interpreter.NewUnmeteredStaticHostFunctionValue(
			functionType,
			func(invocation interpreter.Invocation) interpreter.Value {
				inter := invocation.Interpreter

				// Get interpreter storage

				storage := inter.Storage()

				runtimeStorage, ok := storage.(*runtime.Storage)
				if !ok {
					panic(cadenceErrors.NewUnexpectedError("interpreter storage is not a runtime.Storage"))
				}

				// Check the number of arguments

				actualArgumentCount := len(invocation.Arguments)
				expectedArgumentCount := len(functionType.Parameters)

				if actualArgumentCount != expectedArgumentCount {
					panic(errors.NewInvalidArgumentErrorf(
						"incorrect number of arguments: got %d, expected %d",
						actualArgumentCount,
						expectedArgumentCount,
					))
				}

				// Get addressStartIndex argument

				firstArgument := invocation.Arguments[0]
				addressStartIndexValue, ok := firstArgument.(interpreter.UInt64Value)
				if !ok {
					panic(errors.NewInvalidArgumentErrorf(
						"incorrect type for argument 0: got `%s`, expected `%s`",
						firstArgument.StaticType(inter),
						sema.UInt64Type,
					))
				}
				addressStartIndex := uint64(addressStartIndexValue)

				// Get count argument

				secondArgument := invocation.Arguments[1]
				countValue, ok := secondArgument.(interpreter.UInt64Value)
				if !ok {
					panic(errors.NewInvalidArgumentErrorf(
						"incorrect type for argument 1: got `%s`, expected `%s`",
						secondArgument.StaticType(inter),
						sema.UInt64Type,
					))
				}
				count := uint64(countValue)

				// Schedule the account V2 migration for addresses

				addressGenerator := chainID.Chain().NewAddressGeneratorAtIndex(addressStartIndex)
				for i := uint64(0); i < count; i++ {
					address, err := addressGenerator.NextAddress()
					if err != nil {
						panic(err)
					}

					if !runtimeStorage.ScheduleV2Migration(common.Address(address)) {
						return interpreter.FalseValue
					}
				}

				return interpreter.TrueValue
			},
		),
	}

	// TODO: restrict, but requires to be set during bootstrapping
	//sc := systemcontracts.SystemContractsForChain(chainID)
	//
	//accountV2MigrationLocation := common.NewAddressLocation(
	//	nil,
	//	common.Address(sc.AccountV2Migration.Address),
	//	ContractName,
	//)

	environment.DeclareValue(
		functionValue,
		// TODO: accountV2MigrationLocation,
		nil,
	)
}
