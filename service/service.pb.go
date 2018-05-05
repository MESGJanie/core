// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/service/service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	github.com/mesg-foundation/core/service/service.proto

It has these top-level messages:
	Service
	Task
	Fee
	Event
	Output
	Parameter
	Dependency
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Service struct {
	Name         string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description  string                 `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Visibility   string                 `protobuf:"bytes,3,opt,name=visibility" json:"visibility,omitempty"`
	Publish      string                 `protobuf:"bytes,4,opt,name=publish" json:"publish,omitempty"`
	Tasks        map[string]*Task       `protobuf:"bytes,5,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Events       map[string]*Event      `protobuf:"bytes,6,rep,name=events" json:"events,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Dependencies map[string]*Dependency `protobuf:"bytes,7,rep,name=dependencies" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetVisibility() string {
	if m != nil {
		return m.Visibility
	}
	return ""
}

func (m *Service) GetPublish() string {
	if m != nil {
		return m.Publish
	}
	return ""
}

func (m *Service) GetTasks() map[string]*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Service) GetEvents() map[string]*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Service) GetDependencies() map[string]*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type Task struct {
	Name        string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Verifiable  bool                  `protobuf:"varint,3,opt,name=verifiable" json:"verifiable,omitempty"`
	Payable     bool                  `protobuf:"varint,4,opt,name=payable" json:"payable,omitempty"`
	Fees        *Fee                  `protobuf:"bytes,5,opt,name=fees" json:"fees,omitempty"`
	Inputs      map[string]*Parameter `protobuf:"bytes,6,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outputs     map[string]*Output    `protobuf:"bytes,7,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetVerifiable() bool {
	if m != nil {
		return m.Verifiable
	}
	return false
}

func (m *Task) GetPayable() bool {
	if m != nil {
		return m.Payable
	}
	return false
}

func (m *Task) GetFees() *Fee {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *Task) GetInputs() map[string]*Parameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Task) GetOutputs() map[string]*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type Fee struct {
	Developer string `protobuf:"bytes,1,opt,name=developer" json:"developer,omitempty"`
	Validator string `protobuf:"bytes,2,opt,name=validator" json:"validator,omitempty"`
	Executor  string `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Emittors  string `protobuf:"bytes,4,opt,name=emittors" json:"emittors,omitempty"`
}

func (m *Fee) Reset()                    { *m = Fee{} }
func (m *Fee) String() string            { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()               {}
func (*Fee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Fee) GetDeveloper() string {
	if m != nil {
		return m.Developer
	}
	return ""
}

func (m *Fee) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *Fee) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *Fee) GetEmittors() string {
	if m != nil {
		return m.Emittors
	}
	return ""
}

type Event struct {
	Name        string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Data        map[string]*Parameter `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetData() map[string]*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

type Output struct {
	Name        string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Data        map[string]*Parameter `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Output) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Output) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Output) GetData() map[string]*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

type Parameter struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Optional    bool   `protobuf:"varint,4,opt,name=optional" json:"optional,omitempty"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Parameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

type Dependency struct {
	Image   string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Volumes []string `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
	Ports   []string `protobuf:"bytes,3,rep,name=ports" json:"ports,omitempty"`
	Command string   `protobuf:"bytes,4,opt,name=command" json:"command,omitempty"`
}

func (m *Dependency) Reset()                    { *m = Dependency{} }
func (m *Dependency) String() string            { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()               {}
func (*Dependency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Dependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Dependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Dependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Dependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*Service)(nil), "service.Service")
	proto.RegisterType((*Task)(nil), "service.Task")
	proto.RegisterType((*Fee)(nil), "service.Fee")
	proto.RegisterType((*Event)(nil), "service.Event")
	proto.RegisterType((*Output)(nil), "service.Output")
	proto.RegisterType((*Parameter)(nil), "service.Parameter")
	proto.RegisterType((*Dependency)(nil), "service.Dependency")
}

func init() {
	proto.RegisterFile("github.com/mesg-foundation/core/service/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x6a, 0x27, 0x69, 0x26, 0xe5, 0x6f, 0xe1, 0xe0, 0x86, 0x0a, 0x45, 0x01, 0xa4, 0x22,
	0xd1, 0x54, 0x2d, 0x45, 0x42, 0x9c, 0xdb, 0xa2, 0xaa, 0x42, 0xa0, 0xd0, 0x17, 0xd8, 0xd8, 0xd3,
	0x76, 0x55, 0xdb, 0x6b, 0xec, 0xb5, 0x85, 0xdf, 0x83, 0x2b, 0xef, 0x81, 0x78, 0x3a, 0xb4, 0xb3,
	0xeb, 0x8d, 0x4d, 0x73, 0x6a, 0xc5, 0x29, 0x3b, 0xf3, 0xcd, 0xf7, 0xed, 0xa7, 0x99, 0xd9, 0x18,
	0xde, 0x5f, 0x09, 0x75, 0x5d, 0x2e, 0xe7, 0xa1, 0x4c, 0xf6, 0x13, 0x2c, 0xae, 0xf6, 0x2e, 0x65,
	0x99, 0x46, 0x5c, 0x09, 0x99, 0xee, 0x87, 0x32, 0xc7, 0xfd, 0x02, 0xf3, 0x4a, 0x84, 0xee, 0x77,
	0x9e, 0xe5, 0x52, 0x49, 0x36, 0xb4, 0xe1, 0xec, 0xa7, 0x0f, 0xc3, 0x6f, 0xe6, 0xcc, 0x18, 0xf8,
	0x29, 0x4f, 0x30, 0xe8, 0x4d, 0x7b, 0xbb, 0xa3, 0x05, 0x9d, 0xd9, 0x14, 0xc6, 0x11, 0x16, 0x61,
	0x2e, 0x32, 0x2d, 0x19, 0x6c, 0x10, 0xd4, 0x4e, 0xb1, 0x17, 0x00, 0x95, 0x28, 0xc4, 0x52, 0xc4,
	0x42, 0xd5, 0x81, 0x47, 0x05, 0xad, 0x0c, 0x0b, 0x60, 0x98, 0x95, 0xcb, 0x58, 0x14, 0xd7, 0x81,
	0x4f, 0x60, 0x13, 0xb2, 0x03, 0xe8, 0x2b, 0x5e, 0xdc, 0x14, 0x41, 0x7f, 0xea, 0xed, 0x8e, 0x0f,
	0x9f, 0xcf, 0x1b, 0x8f, 0xd6, 0xd0, 0xfc, 0x42, 0xa3, 0x27, 0xa9, 0xca, 0xeb, 0x85, 0xa9, 0x64,
	0x47, 0x30, 0xc0, 0x0a, 0x53, 0x55, 0x04, 0x03, 0xe2, 0xec, 0xdc, 0xe2, 0x9c, 0x10, 0x6c, 0x48,
	0xb6, 0x96, 0x9d, 0xc2, 0x56, 0x84, 0x19, 0xa6, 0x11, 0xa6, 0xa1, 0xc0, 0x22, 0x18, 0x12, 0x77,
	0x76, 0x8b, 0x7b, 0xdc, 0x2a, 0x32, 0x0a, 0x1d, 0xde, 0xe4, 0x13, 0xc0, 0xca, 0x12, 0x7b, 0x0c,
	0xde, 0x0d, 0xd6, 0xb6, 0x5b, 0xfa, 0xc8, 0x5e, 0x42, 0xbf, 0xe2, 0x71, 0x89, 0xd4, 0xa6, 0xf1,
	0xe1, 0x03, 0x77, 0x81, 0x66, 0x2d, 0x0c, 0xf6, 0x71, 0xe3, 0x43, 0x6f, 0x72, 0x06, 0xe3, 0x96,
	0xcf, 0x35, 0x4a, 0xaf, 0xba, 0x4a, 0x0f, 0x9d, 0x12, 0xd1, 0xda, 0x52, 0x17, 0xf0, 0xe4, 0x96,
	0xed, 0x35, 0x82, 0x6f, 0xba, 0x82, 0x4f, 0x9d, 0xa0, 0x23, 0xd7, 0x2d, 0xd5, 0xd9, 0x2f, 0x0f,
	0x7c, 0x6d, 0xfa, 0x1e, 0x3b, 0x81, 0xb9, 0xb8, 0x14, 0x7c, 0x19, 0x23, 0xed, 0xc4, 0xe6, 0xa2,
	0x95, 0xa1, 0x9d, 0xe0, 0x35, 0x81, 0x3e, 0x81, 0x4d, 0xc8, 0xa6, 0xe0, 0x5f, 0x22, 0xea, 0x95,
	0xd0, 0x36, 0xb7, 0x9c, 0xcd, 0x53, 0xc4, 0x05, 0x21, 0xec, 0x00, 0x06, 0x22, 0xcd, 0x4a, 0xb7,
	0x02, 0xdb, 0x9d, 0x2e, 0xcf, 0xcf, 0x08, 0xb3, 0xf3, 0x37, 0x85, 0xec, 0x08, 0x86, 0xb2, 0x54,
	0xc4, 0x31, 0xa3, 0x9f, 0x74, 0x39, 0x5f, 0x0c, 0x68, 0x48, 0x4d, 0xe9, 0xe4, 0x33, 0x8c, 0x5b,
	0x62, 0x6b, 0x7a, 0xba, 0xdb, 0xed, 0x29, 0x73, 0xa2, 0x5f, 0x79, 0xce, 0x13, 0x54, 0x98, 0xb7,
	0x07, 0x75, 0x0e, 0x5b, 0xed, 0x7b, 0xd6, 0xe8, 0xbd, 0xee, 0xea, 0x3d, 0x72, 0x7a, 0x86, 0xd7,
	0x9e, 0x4f, 0x0d, 0xde, 0x29, 0x22, 0xdb, 0x81, 0x51, 0x84, 0x15, 0xc6, 0x32, 0xc3, 0xdc, 0x2a,
	0xad, 0x12, 0x1a, 0xad, 0x78, 0x2c, 0x22, 0xae, 0x64, 0x6e, 0xa7, 0xb4, 0x4a, 0xb0, 0x09, 0x6c,
	0xe2, 0x0f, 0x0c, 0x4b, 0x0d, 0x9a, 0x57, 0xeb, 0x62, 0xc2, 0x12, 0xa1, 0x94, 0xcc, 0x0b, 0xfb,
	0x68, 0x5d, 0x3c, 0xfb, 0xdd, 0x83, 0x3e, 0x6d, 0xe1, 0x1d, 0x77, 0xe3, 0x2d, 0xf8, 0x11, 0x57,
	0x3c, 0xf0, 0x68, 0x12, 0x41, 0x77, 0xb3, 0xe7, 0xc7, 0x5c, 0x71, 0x33, 0x07, 0xaa, 0x9a, 0x9c,
	0xc3, 0xc8, 0xa5, 0xee, 0x3b, 0x82, 0xd9, 0x9f, 0x1e, 0x0c, 0x4c, 0x2f, 0xef, 0xe8, 0x7d, 0xaf,
	0xe3, 0x7d, 0xfb, 0x9f, 0x01, 0xfd, 0x5f, 0xf3, 0xdf, 0x61, 0xe4, 0xf2, 0x77, 0xb4, 0xcf, 0xc0,
	0x57, 0x75, 0x86, 0x76, 0xdc, 0x74, 0xd6, 0xa3, 0x96, 0x84, 0xf2, 0xd8, 0xbe, 0x45, 0x17, 0xcf,
	0x62, 0x80, 0xd5, 0xdf, 0x03, 0x7b, 0x06, 0x7d, 0x91, 0xf0, 0xab, 0xe6, 0x52, 0x13, 0xe8, 0xa7,
	0x5c, 0xc9, 0xb8, 0x4c, 0xb0, 0x08, 0x36, 0xa6, 0x9e, 0xfe, 0x7b, 0xb7, 0xa1, 0xae, 0xcf, 0x64,
	0xae, 0x0a, 0xea, 0xd6, 0x68, 0x61, 0x02, 0x5d, 0x1f, 0xca, 0x24, 0xe1, 0x69, 0xd4, 0x7c, 0x0e,
	0x6c, 0xb8, 0x1c, 0xd0, 0xa7, 0xe9, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x15, 0xab,
	0x95, 0xd3, 0x06, 0x00, 0x00,
}