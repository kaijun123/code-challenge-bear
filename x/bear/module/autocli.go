package bear

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "bear/api/bear/bear"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowBear",
					Use:            "show-bear [id]",
					Short:          "Query show-bear",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateBear",
					Use:            "create-bear [role] [background] [clothes] [weapon]",
					Short:          "Send a create-bear tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "role"}, {ProtoField: "background"}, {ProtoField: "clothes"}, {ProtoField: "weapon"}},
				},
				{
					RpcMethod:      "CreateBear",
					Use:            "create-bear [role] [background] [clothes] [weapon]",
					Short:          "Send a create-bear tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "role"}, {ProtoField: "background"}, {ProtoField: "clothes"}, {ProtoField: "weapon"}},
				},
				{
					RpcMethod:      "CreateBear",
					Use:            "create-bear [role] [background] [clothes] [weapon]",
					Short:          "Send a create-bear tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "role"}, {ProtoField: "background"}, {ProtoField: "clothes"}, {ProtoField: "weapon"}},
				},
				{
					RpcMethod:      "CreateBear",
					Use:            "create-bear [role] [background] [clothes] [weapon]",
					Short:          "Send a create-bear tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "role"}, {ProtoField: "background"}, {ProtoField: "clothes"}, {ProtoField: "weapon"}},
				},
				{
					RpcMethod:      "UpdateBear",
					Use:            "update-bear [role] [background] [clothes] [weapon] [id]",
					Short:          "Send a update-bear tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "role"}, {ProtoField: "background"}, {ProtoField: "clothes"}, {ProtoField: "weapon"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteBear",
					Use:            "delete-bear [id]",
					Short:          "Send a delete-bear tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
