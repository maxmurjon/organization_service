// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: supplier.proto

package organization_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	StoreId     string `protobuf:"bytes,4,opt,name=storeId,proto3" json:"storeId,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Active      string `protobuf:"bytes,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

func (x *Supplier) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *Supplier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Supplier) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Supplier) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Supplier) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *Supplier) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Supplier) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

type CreateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	StoreId     string `protobuf:"bytes,3,opt,name=storeId,proto3" json:"storeId,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Active      string `protobuf:"bytes,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *CreateSupplierRequest) Reset() {
	*x = CreateSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierRequest) ProtoMessage() {}

func (x *CreateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierRequest.ProtoReflect.Descriptor instead.
func (*CreateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSupplierRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateSupplierRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateSupplierRequest) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *CreateSupplierRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateSupplierRequest) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

type GetListSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *GetListSupplierRequest) Reset() {
	*x = GetListSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSupplierRequest) ProtoMessage() {}

func (x *GetListSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSupplierRequest.ProtoReflect.Descriptor instead.
func (*GetListSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *GetListSupplierRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetListSupplierRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GetSuppliersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suppliers []*Supplier `protobuf:"bytes,1,rep,name=suppliers,proto3" json:"suppliers,omitempty"`
	Count     int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetSuppliersListResponse) Reset() {
	*x = GetSuppliersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuppliersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuppliersListResponse) ProtoMessage() {}

func (x *GetSuppliersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuppliersListResponse.ProtoReflect.Descriptor instead.
func (*GetSuppliersListResponse) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{3}
}

func (x *GetSuppliersListResponse) GetSuppliers() []*Supplier {
	if x != nil {
		return x.Suppliers
	}
	return nil
}

func (x *GetSuppliersListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier *Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier,omitempty"`
}

func (x *UpdateSupplierRequest) Reset() {
	*x = UpdateSupplierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierRequest) ProtoMessage() {}

func (x *UpdateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierRequest.ProtoReflect.Descriptor instead.
func (*UpdateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSupplierRequest) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

var File_supplier_proto protoreflect.FileDescriptor

var file_supplier_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x22, 0xa5, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x52, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6e, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x09, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supplier_proto_rawDescOnce sync.Once
	file_supplier_proto_rawDescData = file_supplier_proto_rawDesc
)

func file_supplier_proto_rawDescGZIP() []byte {
	file_supplier_proto_rawDescOnce.Do(func() {
		file_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_supplier_proto_rawDescData)
	})
	return file_supplier_proto_rawDescData
}

var file_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_supplier_proto_goTypes = []interface{}{
	(*Supplier)(nil),                 // 0: organization_service.Supplier
	(*CreateSupplierRequest)(nil),    // 1: organization_service.CreateSupplierRequest
	(*GetListSupplierRequest)(nil),   // 2: organization_service.GetListSupplierRequest
	(*GetSuppliersListResponse)(nil), // 3: organization_service.GetSuppliersListResponse
	(*UpdateSupplierRequest)(nil),    // 4: organization_service.UpdateSupplierRequest
}
var file_supplier_proto_depIdxs = []int32{
	0, // 0: organization_service.GetSuppliersListResponse.suppliers:type_name -> organization_service.Supplier
	0, // 1: organization_service.UpdateSupplierRequest.supplier:type_name -> organization_service.Supplier
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_supplier_proto_init() }
func file_supplier_proto_init() {
	if File_supplier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_supplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplierRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_supplier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSupplierRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_supplier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuppliersListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_supplier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplierRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_supplier_proto_goTypes,
		DependencyIndexes: file_supplier_proto_depIdxs,
		MessageInfos:      file_supplier_proto_msgTypes,
	}.Build()
	File_supplier_proto = out.File
	file_supplier_proto_rawDesc = nil
	file_supplier_proto_goTypes = nil
	file_supplier_proto_depIdxs = nil
}
