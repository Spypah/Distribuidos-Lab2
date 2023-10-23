// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/main.proto

package proto

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

type EstadoPersona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *EstadoPersona) Reset() {
	*x = EstadoPersona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstadoPersona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoPersona) ProtoMessage() {}

func (x *EstadoPersona) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoPersona.ProtoReflect.Descriptor instead.
func (*EstadoPersona) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

func (x *EstadoPersona) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *EstadoPersona) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *EstadoPersona) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Estado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Estado) Reset() {
	*x = Estado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estado) ProtoMessage() {}

func (x *Estado) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estado.ProtoReflect.Descriptor instead.
func (*Estado) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

func (x *Estado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type IdPersona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdPersona) Reset() {
	*x = IdPersona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdPersona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdPersona) ProtoMessage() {}

func (x *IdPersona) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdPersona.ProtoReflect.Descriptor instead.
func (*IdPersona) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

func (x *IdPersona) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListaPersonas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Personas []*NombrePersona `protobuf:"bytes,1,rep,name=personas,proto3" json:"personas,omitempty"` // ojito aqui
}

func (x *ListaPersonas) Reset() {
	*x = ListaPersonas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListaPersonas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaPersonas) ProtoMessage() {}

func (x *ListaPersonas) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaPersonas.ProtoReflect.Descriptor instead.
func (*ListaPersonas) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{3}
}

func (x *ListaPersonas) GetPersonas() []*NombrePersona {
	if x != nil {
		return x.Personas
	}
	return nil
}

type Persona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
	Id       int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Persona) Reset() {
	*x = Persona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Persona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Persona) ProtoMessage() {}

func (x *Persona) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Persona.ProtoReflect.Descriptor instead.
func (*Persona) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *Persona) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Persona) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *Persona) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NombrePersona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *NombrePersona) Reset() {
	*x = NombrePersona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NombrePersona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NombrePersona) ProtoMessage() {}

func (x *NombrePersona) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NombrePersona.ProtoReflect.Descriptor instead.
func (*NombrePersona) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{5}
}

func (x *NombrePersona) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *NombrePersona) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_main_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{6}
}

func (x *Ok) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_proto_main_proto protoreflect.FileDescriptor

var file_proto_main_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x5b, 0x0a, 0x0d, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x52, 0x08, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x73, 0x22, 0x4d, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x0d, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22, 0x14, 0x0a, 0x02, 0x4f, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x32, 0xd8, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x10,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x1a, 0x08, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x6b, 0x22,
	0x00, 0x12, 0x2a, 0x0a, 0x0d, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x4e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x1a, 0x08, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0d, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x1a,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x61, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData = file_proto_main_proto_rawDesc
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_proto_rawDescData)
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_main_proto_goTypes = []interface{}{
	(*EstadoPersona)(nil), // 0: grpc.EstadoPersona
	(*Estado)(nil),        // 1: grpc.Estado
	(*IdPersona)(nil),     // 2: grpc.IdPersona
	(*ListaPersonas)(nil), // 3: grpc.ListaPersonas
	(*Persona)(nil),       // 4: grpc.Persona
	(*NombrePersona)(nil), // 5: grpc.NombrePersona
	(*Ok)(nil),            // 6: grpc.Ok
}
var file_proto_main_proto_depIdxs = []int32{
	5, // 0: grpc.ListaPersonas.personas:type_name -> grpc.NombrePersona
	0, // 1: grpc.Greeter.NotificarPersona:input_type -> grpc.EstadoPersona
	4, // 2: grpc.Greeter.GuardarNombre:input_type -> grpc.Persona
	2, // 3: grpc.Greeter.ObtenerNombre:input_type -> grpc.IdPersona
	1, // 4: grpc.Greeter.ObtenerLista:input_type -> grpc.Estado
	6, // 5: grpc.Greeter.NotificarPersona:output_type -> grpc.Ok
	6, // 6: grpc.Greeter.GuardarNombre:output_type -> grpc.Ok
	5, // 7: grpc.Greeter.ObtenerNombre:output_type -> grpc.NombrePersona
	3, // 8: grpc.Greeter.ObtenerLista:output_type -> grpc.ListaPersonas
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstadoPersona); i {
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
		file_proto_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estado); i {
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
		file_proto_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdPersona); i {
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
		file_proto_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListaPersonas); i {
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
		file_proto_main_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Persona); i {
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
		file_proto_main_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NombrePersona); i {
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
		file_proto_main_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
			RawDescriptor: file_proto_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_rawDesc = nil
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}