// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.21.12
// source: bank/service.proto

package bank

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_bank_service_proto protoreflect.FileDescriptor

var file_bank_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x17, 0x62, 0x61, 0x6e, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x62, 0x61, 0x6e, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x81, 0x03, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x15, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x7a, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x11, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x45, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x68, 0x53, 0x4c,
	0x45, 0x50, 0x41, 0x4d, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_bank_service_proto_goTypes = []any{
	(*CurrentBalanceRequest)(nil),  // 0: bank.CurrentBalanceRequest
	(*ExchangeRateRequest)(nil),    // 1: bank.ExchangeRateRequest
	(*Transaction)(nil),            // 2: bank.Transaction
	(*TransferRequest)(nil),        // 3: bank.TransferRequest
	(*AccountRequest)(nil),         // 4: bank.AccountRequest
	(*CurrentBalanceResponse)(nil), // 5: bank.CurrentBalanceResponse
	(*ExchangeRateResponse)(nil),   // 6: bank.ExchangeRateResponse
	(*TransactionSummary)(nil),     // 7: bank.TransactionSummary
	(*TransferResponse)(nil),       // 8: bank.TransferResponse
	(*AccountResponse)(nil),        // 9: bank.AccountResponse
}
var file_bank_service_proto_depIdxs = []int32{
	0, // 0: bank.BankService.GetCurrentBalance:input_type -> bank.CurrentBalanceRequest
	1, // 1: bank.BankService.FetchExchangeRates:input_type -> bank.ExchangeRateRequest
	2, // 2: bank.BankService.SummarizeTransactions:input_type -> bank.Transaction
	3, // 3: bank.BankService.TransferMultiple:input_type -> bank.TransferRequest
	4, // 4: bank.BankService.CreateAccount:input_type -> bank.AccountRequest
	5, // 5: bank.BankService.GetCurrentBalance:output_type -> bank.CurrentBalanceResponse
	6, // 6: bank.BankService.FetchExchangeRates:output_type -> bank.ExchangeRateResponse
	7, // 7: bank.BankService.SummarizeTransactions:output_type -> bank.TransactionSummary
	8, // 8: bank.BankService.TransferMultiple:output_type -> bank.TransferResponse
	9, // 9: bank.BankService.CreateAccount:output_type -> bank.AccountResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bank_service_proto_init() }
func file_bank_service_proto_init() {
	if File_bank_service_proto != nil {
		return
	}
	file_bank_type_account_proto_init()
	file_bank_type_exchange_proto_init()
	file_bank_type_transaction_proto_init()
	file_bank_type_transfer_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_service_proto_goTypes,
		DependencyIndexes: file_bank_service_proto_depIdxs,
	}.Build()
	File_bank_service_proto = out.File
	file_bank_service_proto_rawDesc = nil
	file_bank_service_proto_goTypes = nil
	file_bank_service_proto_depIdxs = nil
}
