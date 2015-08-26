// Code generated by protoc-gen-gogo.
// source: evacuation.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import fmt "fmt"
import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EvacuationResponse struct {
	Error         *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	KeepContainer bool   `protobuf:"varint,2,opt,name=keep_container" json:"keep_container"`
}

func (m *EvacuationResponse) Reset()      { *m = EvacuationResponse{} }
func (*EvacuationResponse) ProtoMessage() {}

func (m *EvacuationResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *EvacuationResponse) GetKeepContainer() bool {
	if m != nil {
		return m.KeepContainer
	}
	return false
}

type EvacuateClaimedActualLRPRequest struct {
	ActualLrpKey         *ActualLRPKey         `protobuf:"bytes,1,opt,name=actual_lrp_key" json:"actual_lrp_key,omitempty"`
	ActualLrpInstanceKey *ActualLRPInstanceKey `protobuf:"bytes,2,opt,name=actual_lrp_instance_key" json:"actual_lrp_instance_key,omitempty"`
}

func (m *EvacuateClaimedActualLRPRequest) Reset()      { *m = EvacuateClaimedActualLRPRequest{} }
func (*EvacuateClaimedActualLRPRequest) ProtoMessage() {}

func (m *EvacuateClaimedActualLRPRequest) GetActualLrpKey() *ActualLRPKey {
	if m != nil {
		return m.ActualLrpKey
	}
	return nil
}

func (m *EvacuateClaimedActualLRPRequest) GetActualLrpInstanceKey() *ActualLRPInstanceKey {
	if m != nil {
		return m.ActualLrpInstanceKey
	}
	return nil
}

type EvacuateRunningActualLRPRequest struct {
	ActualLrpKey         *ActualLRPKey         `protobuf:"bytes,1,opt,name=actual_lrp_key" json:"actual_lrp_key,omitempty"`
	ActualLrpInstanceKey *ActualLRPInstanceKey `protobuf:"bytes,2,opt,name=actual_lrp_instance_key" json:"actual_lrp_instance_key,omitempty"`
	ActualLrpNetInfo     *ActualLRPNetInfo     `protobuf:"bytes,3,opt,name=actual_lrp_net_info" json:"actual_lrp_net_info,omitempty"`
	Ttl                  uint64                `protobuf:"varint,4,opt,name=ttl" json:"ttl"`
}

func (m *EvacuateRunningActualLRPRequest) Reset()      { *m = EvacuateRunningActualLRPRequest{} }
func (*EvacuateRunningActualLRPRequest) ProtoMessage() {}

func (m *EvacuateRunningActualLRPRequest) GetActualLrpKey() *ActualLRPKey {
	if m != nil {
		return m.ActualLrpKey
	}
	return nil
}

func (m *EvacuateRunningActualLRPRequest) GetActualLrpInstanceKey() *ActualLRPInstanceKey {
	if m != nil {
		return m.ActualLrpInstanceKey
	}
	return nil
}

func (m *EvacuateRunningActualLRPRequest) GetActualLrpNetInfo() *ActualLRPNetInfo {
	if m != nil {
		return m.ActualLrpNetInfo
	}
	return nil
}

func (m *EvacuateRunningActualLRPRequest) GetTtl() uint64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type EvacuateStoppedActualLRPRequest struct {
	ActualLrpKey         *ActualLRPKey         `protobuf:"bytes,1,opt,name=actual_lrp_key" json:"actual_lrp_key,omitempty"`
	ActualLrpInstanceKey *ActualLRPInstanceKey `protobuf:"bytes,2,opt,name=actual_lrp_instance_key" json:"actual_lrp_instance_key,omitempty"`
}

func (m *EvacuateStoppedActualLRPRequest) Reset()      { *m = EvacuateStoppedActualLRPRequest{} }
func (*EvacuateStoppedActualLRPRequest) ProtoMessage() {}

func (m *EvacuateStoppedActualLRPRequest) GetActualLrpKey() *ActualLRPKey {
	if m != nil {
		return m.ActualLrpKey
	}
	return nil
}

func (m *EvacuateStoppedActualLRPRequest) GetActualLrpInstanceKey() *ActualLRPInstanceKey {
	if m != nil {
		return m.ActualLrpInstanceKey
	}
	return nil
}

type EvacuateCrashedActualLRPRequest struct {
	ActualLrpKey         *ActualLRPKey         `protobuf:"bytes,1,opt,name=actual_lrp_key" json:"actual_lrp_key,omitempty"`
	ActualLrpInstanceKey *ActualLRPInstanceKey `protobuf:"bytes,2,opt,name=actual_lrp_instance_key" json:"actual_lrp_instance_key,omitempty"`
	ErrorMessage         string                `protobuf:"bytes,3,opt,name=error_message" json:"error_message"`
}

func (m *EvacuateCrashedActualLRPRequest) Reset()      { *m = EvacuateCrashedActualLRPRequest{} }
func (*EvacuateCrashedActualLRPRequest) ProtoMessage() {}

func (m *EvacuateCrashedActualLRPRequest) GetActualLrpKey() *ActualLRPKey {
	if m != nil {
		return m.ActualLrpKey
	}
	return nil
}

func (m *EvacuateCrashedActualLRPRequest) GetActualLrpInstanceKey() *ActualLRPInstanceKey {
	if m != nil {
		return m.ActualLrpInstanceKey
	}
	return nil
}

func (m *EvacuateCrashedActualLRPRequest) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type RemoveEvacuatingActualLRPRequest struct {
	ActualLrpKey         *ActualLRPKey         `protobuf:"bytes,1,opt,name=actual_lrp_key" json:"actual_lrp_key,omitempty"`
	ActualLrpInstanceKey *ActualLRPInstanceKey `protobuf:"bytes,2,opt,name=actual_lrp_instance_key" json:"actual_lrp_instance_key,omitempty"`
}

func (m *RemoveEvacuatingActualLRPRequest) Reset()      { *m = RemoveEvacuatingActualLRPRequest{} }
func (*RemoveEvacuatingActualLRPRequest) ProtoMessage() {}

func (m *RemoveEvacuatingActualLRPRequest) GetActualLrpKey() *ActualLRPKey {
	if m != nil {
		return m.ActualLrpKey
	}
	return nil
}

func (m *RemoveEvacuatingActualLRPRequest) GetActualLrpInstanceKey() *ActualLRPInstanceKey {
	if m != nil {
		return m.ActualLrpInstanceKey
	}
	return nil
}

type RemoveEvacuatingActualLRPResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *RemoveEvacuatingActualLRPResponse) Reset()      { *m = RemoveEvacuatingActualLRPResponse{} }
func (*RemoveEvacuatingActualLRPResponse) ProtoMessage() {}

func (m *RemoveEvacuatingActualLRPResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (this *EvacuationResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.EvacuationResponse{` +
		`Error:` + fmt.Sprintf("%#v", this.Error),
		`KeepContainer:` + fmt.Sprintf("%#v", this.KeepContainer) + `}`}, ", ")
	return s
}
func (this *EvacuateClaimedActualLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.EvacuateClaimedActualLRPRequest{` +
		`ActualLrpKey:` + fmt.Sprintf("%#v", this.ActualLrpKey),
		`ActualLrpInstanceKey:` + fmt.Sprintf("%#v", this.ActualLrpInstanceKey) + `}`}, ", ")
	return s
}
func (this *EvacuateRunningActualLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.EvacuateRunningActualLRPRequest{` +
		`ActualLrpKey:` + fmt.Sprintf("%#v", this.ActualLrpKey),
		`ActualLrpInstanceKey:` + fmt.Sprintf("%#v", this.ActualLrpInstanceKey),
		`ActualLrpNetInfo:` + fmt.Sprintf("%#v", this.ActualLrpNetInfo),
		`Ttl:` + fmt.Sprintf("%#v", this.Ttl) + `}`}, ", ")
	return s
}
func (this *EvacuateStoppedActualLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.EvacuateStoppedActualLRPRequest{` +
		`ActualLrpKey:` + fmt.Sprintf("%#v", this.ActualLrpKey),
		`ActualLrpInstanceKey:` + fmt.Sprintf("%#v", this.ActualLrpInstanceKey) + `}`}, ", ")
	return s
}
func (this *EvacuateCrashedActualLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.EvacuateCrashedActualLRPRequest{` +
		`ActualLrpKey:` + fmt.Sprintf("%#v", this.ActualLrpKey),
		`ActualLrpInstanceKey:` + fmt.Sprintf("%#v", this.ActualLrpInstanceKey),
		`ErrorMessage:` + fmt.Sprintf("%#v", this.ErrorMessage) + `}`}, ", ")
	return s
}
func (this *RemoveEvacuatingActualLRPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.RemoveEvacuatingActualLRPRequest{` +
		`ActualLrpKey:` + fmt.Sprintf("%#v", this.ActualLrpKey),
		`ActualLrpInstanceKey:` + fmt.Sprintf("%#v", this.ActualLrpInstanceKey) + `}`}, ", ")
	return s
}
func (this *RemoveEvacuatingActualLRPResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.RemoveEvacuatingActualLRPResponse{` +
		`Error:` + fmt.Sprintf("%#v", this.Error) + `}`}, ", ")
	return s
}
func valueToGoStringEvacuation(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringEvacuation(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *EvacuationResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EvacuationResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.Error.Size()))
		n1, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	data[i] = 0x10
	i++
	if m.KeepContainer {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func (m *EvacuateClaimedActualLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EvacuateClaimedActualLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpKey.Size()))
		n2, err := m.ActualLrpKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ActualLrpInstanceKey != nil {
		data[i] = 0x12
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpInstanceKey.Size()))
		n3, err := m.ActualLrpInstanceKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *EvacuateRunningActualLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EvacuateRunningActualLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpKey.Size()))
		n4, err := m.ActualLrpKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.ActualLrpInstanceKey != nil {
		data[i] = 0x12
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpInstanceKey.Size()))
		n5, err := m.ActualLrpInstanceKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.ActualLrpNetInfo != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpNetInfo.Size()))
		n6, err := m.ActualLrpNetInfo.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	data[i] = 0x20
	i++
	i = encodeVarintEvacuation(data, i, uint64(m.Ttl))
	return i, nil
}

func (m *EvacuateStoppedActualLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EvacuateStoppedActualLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpKey.Size()))
		n7, err := m.ActualLrpKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.ActualLrpInstanceKey != nil {
		data[i] = 0x12
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpInstanceKey.Size()))
		n8, err := m.ActualLrpInstanceKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}

func (m *EvacuateCrashedActualLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EvacuateCrashedActualLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpKey.Size()))
		n9, err := m.ActualLrpKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.ActualLrpInstanceKey != nil {
		data[i] = 0x12
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpInstanceKey.Size()))
		n10, err := m.ActualLrpInstanceKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	data[i] = 0x1a
	i++
	i = encodeVarintEvacuation(data, i, uint64(len(m.ErrorMessage)))
	i += copy(data[i:], m.ErrorMessage)
	return i, nil
}

func (m *RemoveEvacuatingActualLRPRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoveEvacuatingActualLRPRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpKey.Size()))
		n11, err := m.ActualLrpKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	if m.ActualLrpInstanceKey != nil {
		data[i] = 0x12
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.ActualLrpInstanceKey.Size()))
		n12, err := m.ActualLrpInstanceKey.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	return i, nil
}

func (m *RemoveEvacuatingActualLRPResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RemoveEvacuatingActualLRPResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintEvacuation(data, i, uint64(m.Error.Size()))
		n13, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	return i, nil
}

func encodeFixed64Evacuation(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Evacuation(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEvacuation(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *EvacuationResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	n += 2
	return n
}

func (m *EvacuateClaimedActualLRPRequest) Size() (n int) {
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		l = m.ActualLrpKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	if m.ActualLrpInstanceKey != nil {
		l = m.ActualLrpInstanceKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	return n
}

func (m *EvacuateRunningActualLRPRequest) Size() (n int) {
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		l = m.ActualLrpKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	if m.ActualLrpInstanceKey != nil {
		l = m.ActualLrpInstanceKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	if m.ActualLrpNetInfo != nil {
		l = m.ActualLrpNetInfo.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	n += 1 + sovEvacuation(uint64(m.Ttl))
	return n
}

func (m *EvacuateStoppedActualLRPRequest) Size() (n int) {
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		l = m.ActualLrpKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	if m.ActualLrpInstanceKey != nil {
		l = m.ActualLrpInstanceKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	return n
}

func (m *EvacuateCrashedActualLRPRequest) Size() (n int) {
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		l = m.ActualLrpKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	if m.ActualLrpInstanceKey != nil {
		l = m.ActualLrpInstanceKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	l = len(m.ErrorMessage)
	n += 1 + l + sovEvacuation(uint64(l))
	return n
}

func (m *RemoveEvacuatingActualLRPRequest) Size() (n int) {
	var l int
	_ = l
	if m.ActualLrpKey != nil {
		l = m.ActualLrpKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	if m.ActualLrpInstanceKey != nil {
		l = m.ActualLrpInstanceKey.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	return n
}

func (m *RemoveEvacuatingActualLRPResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovEvacuation(uint64(l))
	}
	return n
}

func sovEvacuation(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvacuation(x uint64) (n int) {
	return sovEvacuation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *EvacuationResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EvacuationResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`KeepContainer:` + fmt.Sprintf("%v", this.KeepContainer) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EvacuateClaimedActualLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EvacuateClaimedActualLRPRequest{`,
		`ActualLrpKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpKey), "ActualLRPKey", "ActualLRPKey", 1) + `,`,
		`ActualLrpInstanceKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpInstanceKey), "ActualLRPInstanceKey", "ActualLRPInstanceKey", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EvacuateRunningActualLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EvacuateRunningActualLRPRequest{`,
		`ActualLrpKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpKey), "ActualLRPKey", "ActualLRPKey", 1) + `,`,
		`ActualLrpInstanceKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpInstanceKey), "ActualLRPInstanceKey", "ActualLRPInstanceKey", 1) + `,`,
		`ActualLrpNetInfo:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpNetInfo), "ActualLRPNetInfo", "ActualLRPNetInfo", 1) + `,`,
		`Ttl:` + fmt.Sprintf("%v", this.Ttl) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EvacuateStoppedActualLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EvacuateStoppedActualLRPRequest{`,
		`ActualLrpKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpKey), "ActualLRPKey", "ActualLRPKey", 1) + `,`,
		`ActualLrpInstanceKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpInstanceKey), "ActualLRPInstanceKey", "ActualLRPInstanceKey", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EvacuateCrashedActualLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EvacuateCrashedActualLRPRequest{`,
		`ActualLrpKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpKey), "ActualLRPKey", "ActualLRPKey", 1) + `,`,
		`ActualLrpInstanceKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpInstanceKey), "ActualLRPInstanceKey", "ActualLRPInstanceKey", 1) + `,`,
		`ErrorMessage:` + fmt.Sprintf("%v", this.ErrorMessage) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RemoveEvacuatingActualLRPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RemoveEvacuatingActualLRPRequest{`,
		`ActualLrpKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpKey), "ActualLRPKey", "ActualLRPKey", 1) + `,`,
		`ActualLrpInstanceKey:` + strings.Replace(fmt.Sprintf("%v", this.ActualLrpInstanceKey), "ActualLRPInstanceKey", "ActualLRPInstanceKey", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RemoveEvacuatingActualLRPResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RemoveEvacuatingActualLRPResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEvacuation(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *EvacuationResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeepContainer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.KeepContainer = bool(v != 0)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *EvacuateClaimedActualLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpKey == nil {
				m.ActualLrpKey = &ActualLRPKey{}
			}
			if err := m.ActualLrpKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpInstanceKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpInstanceKey == nil {
				m.ActualLrpInstanceKey = &ActualLRPInstanceKey{}
			}
			if err := m.ActualLrpInstanceKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *EvacuateRunningActualLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpKey == nil {
				m.ActualLrpKey = &ActualLRPKey{}
			}
			if err := m.ActualLrpKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpInstanceKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpInstanceKey == nil {
				m.ActualLrpInstanceKey = &ActualLRPInstanceKey{}
			}
			if err := m.ActualLrpInstanceKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpNetInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpNetInfo == nil {
				m.ActualLrpNetInfo = &ActualLRPNetInfo{}
			}
			if err := m.ActualLrpNetInfo.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Ttl |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *EvacuateStoppedActualLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpKey == nil {
				m.ActualLrpKey = &ActualLRPKey{}
			}
			if err := m.ActualLrpKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpInstanceKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpInstanceKey == nil {
				m.ActualLrpInstanceKey = &ActualLRPInstanceKey{}
			}
			if err := m.ActualLrpInstanceKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *EvacuateCrashedActualLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpKey == nil {
				m.ActualLrpKey = &ActualLRPKey{}
			}
			if err := m.ActualLrpKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpInstanceKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpInstanceKey == nil {
				m.ActualLrpInstanceKey = &ActualLRPInstanceKey{}
			}
			if err := m.ActualLrpInstanceKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *RemoveEvacuatingActualLRPRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpKey == nil {
				m.ActualLrpKey = &ActualLRPKey{}
			}
			if err := m.ActualLrpKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActualLrpInstanceKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActualLrpInstanceKey == nil {
				m.ActualLrpInstanceKey = &ActualLRPInstanceKey{}
			}
			if err := m.ActualLrpInstanceKey.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *RemoveEvacuatingActualLRPResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthEvacuation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipEvacuation(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvacuation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipEvacuation(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEvacuation
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEvacuation(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEvacuation = fmt.Errorf("proto: negative length found during unmarshaling")
)
