// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policystore.proto

package policystore

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

// An enumeration which allows us to specify what type of sharing is to be
// defined for the specified sensor type. The default value is `SHARE` which
// implies sharing the data at full resolution. If this type is specified, it
// is an error if either of `buckets` or `interval` is also supplied.
type Operation_Action int32

const (
	Operation_UNKNOWN    Operation_Action = 0
	Operation_SHARE      Operation_Action = 1
	Operation_BIN        Operation_Action = 2
	Operation_MOVING_AVG Operation_Action = 3
)

var Operation_Action_name = map[int32]string{
	0: "UNKNOWN",
	1: "SHARE",
	2: "BIN",
	3: "MOVING_AVG",
}
var Operation_Action_value = map[string]int32{
	"UNKNOWN":    0,
	"SHARE":      1,
	"BIN":        2,
	"MOVING_AVG": 3,
}

func (x Operation_Action) String() string {
	return proto.EnumName(Operation_Action_name, int32(x))
}
func (Operation_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{0, 0}
}

// Operation is a message used to describe an operation that may be applied to
// a specific data type published by a SmartCitizen device. The message contains
// two required fields: the sensor_id (this is the type of data we are entitling
// over), and a specified operation to be performed on that sensor type. This
// can be one of three actions: to share the sensor without modification, to
// apply a binning algorithm to the data so we output a bucketed value, or a
// moving average calculated dynamically for incoming values.
//
// If an operation specifies an Action type of `BIN`, then the optional
// `buckets` parameter is required, similarly if an action type of `MOVING_AVG`
// is specified, then `interval` is a required field.
type Operation struct {
	// The unique id of the sensor type for which this specific entitlement is
	// defined. This is a required field.
	SensorId uint32 `protobuf:"varint,1,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	// The specific action this operation defines for the sensor type. This is a
	// required field.
	Action Operation_Action `protobuf:"varint,2,opt,name=action,proto3,enum=decode.iot.policystore.Operation_Action" json:"action,omitempty"`
	// The bins attribute is used to specify the the bins into which incoming
	// values should be classified. Each element in the list is the upper
	// inclusive bound of a bin. The values submitted must be sorted in strictly
	// increasing order. There is no need to add a highest bin with +Inf bound, it
	// will be added implicitly. This field is optional unless an Action of `BIN`
	// has been requested, in which case it is required. It is an error to send
	// values for this attribute unless the value of Action is `BIN`.
	Bins []float64 `protobuf:"fixed64,3,rep,packed,name=bins,proto3" json:"bins,omitempty"`
	// This attribute is used to control the entitlement in the case for which we
	// have specified an action type representing a moving average. It represents
	// the interval in seconds over which the moving average should be calculated,
	// e.g. for a 15 minute moving average the value supplied here would be 900.
	// This field is optional unless an Action of `MOVING_AVG` has been specified,
	// in which case it is required. It is an error to send a value for this
	// attribute unless the value of Action is `MOVING_AVG`.
	Interval             uint32   `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{0}
}
func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (dst *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(dst, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetSensorId() uint32 {
	if m != nil {
		return m.SensorId
	}
	return 0
}

func (m *Operation) GetAction() Operation_Action {
	if m != nil {
		return m.Action
	}
	return Operation_UNKNOWN
}

func (m *Operation) GetBins() []float64 {
	if m != nil {
		return m.Bins
	}
	return nil
}

func (m *Operation) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// CreateEntitlementPolicyRequest is a message sent to the policy registration
// service to create a new entitlement policy. An entitlement policy is a
// collection of one or more "Operations". A single Operation specifies an
// functional transformation to be performed on a single data channel being
// published by a SmartCitizen device. The policy as a whole is comprised of
// one or more Operations.
type CreateEntitlementPolicyRequest struct {
	// This attribute is used to attach a human friendly label to the policy
	// suitable for presenting to the end user in the DECODE wallet. This is a
	// required field.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// The list of operations we wish to create for the policy. This field is
	// required, and it is required that the client supplies at least one
	// Operation.
	Operations []*Operation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	// This attribute contains an identifier of an authorizable attribute required
	// for taking part in the Coconut protocol.
	AuthorizableAttributeId string `protobuf:"bytes,4,opt,name=authorizable_attribute_id,json=authorizableAttributeId,proto3" json:"authorizable_attribute_id,omitempty"`
	// This attribute contains a reference to the credential issuer service the
	// end user must interact with in order to obtain blind credentials when
	// taking part in the Coconut protocol.
	CredentialIssuerEndpointUrl string `protobuf:"bytes,5,opt,name=credential_issuer_endpoint_url,json=credentialIssuerEndpointUrl,proto3" json:"credential_issuer_endpoint_url,omitempty"`
	// Multi language description of the community for which this policy applies
	// consisting of a map where the map keys are a standard language code, and
	// the value is the description in that language.
	Descriptions         map[string]string `protobuf:"bytes,6,rep,name=descriptions,proto3" json:"descriptions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateEntitlementPolicyRequest) Reset()         { *m = CreateEntitlementPolicyRequest{} }
func (m *CreateEntitlementPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntitlementPolicyRequest) ProtoMessage()    {}
func (*CreateEntitlementPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{1}
}
func (m *CreateEntitlementPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntitlementPolicyRequest.Unmarshal(m, b)
}
func (m *CreateEntitlementPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntitlementPolicyRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEntitlementPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntitlementPolicyRequest.Merge(dst, src)
}
func (m *CreateEntitlementPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntitlementPolicyRequest.Size(m)
}
func (m *CreateEntitlementPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntitlementPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntitlementPolicyRequest proto.InternalMessageInfo

func (m *CreateEntitlementPolicyRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *CreateEntitlementPolicyRequest) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *CreateEntitlementPolicyRequest) GetAuthorizableAttributeId() string {
	if m != nil {
		return m.AuthorizableAttributeId
	}
	return ""
}

func (m *CreateEntitlementPolicyRequest) GetCredentialIssuerEndpointUrl() string {
	if m != nil {
		return m.CredentialIssuerEndpointUrl
	}
	return ""
}

func (m *CreateEntitlementPolicyRequest) GetDescriptions() map[string]string {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

// CreateEntitlementPolicyResponse is a message returned by the service after a
// policy has been created. The message simply contains an identifier for the
// policy, as well as a token that the caller must protect.
type CreateEntitlementPolicyResponse struct {
	// This attribute contains a unique identifier for the community that can be
	// used for later requests to either join a community to a specific device, or
	// to delete the policy and so prevent new instances being applied to devices.
	CommunityId string `protobuf:"bytes,3,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
	// This attribute contains a secret generated by the service that is
	// associated with the policy. This token is required to be presented by a
	// caller when deleting a policy, so must be treated as confidential by the
	// caller.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntitlementPolicyResponse) Reset()         { *m = CreateEntitlementPolicyResponse{} }
func (m *CreateEntitlementPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEntitlementPolicyResponse) ProtoMessage()    {}
func (*CreateEntitlementPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{2}
}
func (m *CreateEntitlementPolicyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntitlementPolicyResponse.Unmarshal(m, b)
}
func (m *CreateEntitlementPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntitlementPolicyResponse.Marshal(b, m, deterministic)
}
func (dst *CreateEntitlementPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntitlementPolicyResponse.Merge(dst, src)
}
func (m *CreateEntitlementPolicyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEntitlementPolicyResponse.Size(m)
}
func (m *CreateEntitlementPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntitlementPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntitlementPolicyResponse proto.InternalMessageInfo

func (m *CreateEntitlementPolicyResponse) GetCommunityId() string {
	if m != nil {
		return m.CommunityId
	}
	return ""
}

func (m *CreateEntitlementPolicyResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DeleteEntitlementPolicyRequest is a message that can be sent to the
// registration service in order to delete an existing policy.
//
// Deleting a policy does not affect any already existing streams configured for
// the policy, it just stops any new instances of this policy being applied to
// other devices.
type DeleteEntitlementPolicyRequest struct {
	// This attribute contains the unique community identifier returned when
	// creating the policy. This is a required field.
	CommunityId string `protobuf:"bytes,3,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
	// This attribute contains the token returned to the creator when they
	// created the policy, and must match the value stored within the
	// PolicyStore. This is a required field.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntitlementPolicyRequest) Reset()         { *m = DeleteEntitlementPolicyRequest{} }
func (m *DeleteEntitlementPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEntitlementPolicyRequest) ProtoMessage()    {}
func (*DeleteEntitlementPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{3}
}
func (m *DeleteEntitlementPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntitlementPolicyRequest.Unmarshal(m, b)
}
func (m *DeleteEntitlementPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntitlementPolicyRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteEntitlementPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntitlementPolicyRequest.Merge(dst, src)
}
func (m *DeleteEntitlementPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEntitlementPolicyRequest.Size(m)
}
func (m *DeleteEntitlementPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntitlementPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntitlementPolicyRequest proto.InternalMessageInfo

func (m *DeleteEntitlementPolicyRequest) GetCommunityId() string {
	if m != nil {
		return m.CommunityId
	}
	return ""
}

func (m *DeleteEntitlementPolicyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// DeleteEntitlementPolicyResponse is a placeholder response returned from a
// delete request. Currently empty, but reserved for any fields identified for
// future iterations.
type DeleteEntitlementPolicyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntitlementPolicyResponse) Reset()         { *m = DeleteEntitlementPolicyResponse{} }
func (m *DeleteEntitlementPolicyResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEntitlementPolicyResponse) ProtoMessage()    {}
func (*DeleteEntitlementPolicyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{4}
}
func (m *DeleteEntitlementPolicyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntitlementPolicyResponse.Unmarshal(m, b)
}
func (m *DeleteEntitlementPolicyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntitlementPolicyResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteEntitlementPolicyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntitlementPolicyResponse.Merge(dst, src)
}
func (m *DeleteEntitlementPolicyResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEntitlementPolicyResponse.Size(m)
}
func (m *DeleteEntitlementPolicyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntitlementPolicyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntitlementPolicyResponse proto.InternalMessageInfo

// ListEntitlementPoliciesRequest is the message sent to the service in order
// to receive a list of currently defined entitlement policies. Currently this
// message is empty as we simply return a list of all known policies, but this
// message may be extended should a need be identified to paginate through
// policies, or apply any search or filtering techniques.
type ListEntitlementPoliciesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEntitlementPoliciesRequest) Reset()         { *m = ListEntitlementPoliciesRequest{} }
func (m *ListEntitlementPoliciesRequest) String() string { return proto.CompactTextString(m) }
func (*ListEntitlementPoliciesRequest) ProtoMessage()    {}
func (*ListEntitlementPoliciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{5}
}
func (m *ListEntitlementPoliciesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitlementPoliciesRequest.Unmarshal(m, b)
}
func (m *ListEntitlementPoliciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitlementPoliciesRequest.Marshal(b, m, deterministic)
}
func (dst *ListEntitlementPoliciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitlementPoliciesRequest.Merge(dst, src)
}
func (m *ListEntitlementPoliciesRequest) XXX_Size() int {
	return xxx_messageInfo_ListEntitlementPoliciesRequest.Size(m)
}
func (m *ListEntitlementPoliciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitlementPoliciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitlementPoliciesRequest proto.InternalMessageInfo

// ListEntitlementPoliciesResponse is the response to the method call to list
// policies. It simply returns a list of all currently registered and
// non-deleted policies. This is intended to be able to be fed to the DECODE
// wallet in order to allow participant to choose which entitlements to apply to
// their devices.
type ListEntitlementPoliciesResponse struct {
	// This attribute contains the list of all policies currently available on
	// the device registration service.
	Policies             []*ListEntitlementPoliciesResponse_Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *ListEntitlementPoliciesResponse) Reset()         { *m = ListEntitlementPoliciesResponse{} }
func (m *ListEntitlementPoliciesResponse) String() string { return proto.CompactTextString(m) }
func (*ListEntitlementPoliciesResponse) ProtoMessage()    {}
func (*ListEntitlementPoliciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{6}
}
func (m *ListEntitlementPoliciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitlementPoliciesResponse.Unmarshal(m, b)
}
func (m *ListEntitlementPoliciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitlementPoliciesResponse.Marshal(b, m, deterministic)
}
func (dst *ListEntitlementPoliciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitlementPoliciesResponse.Merge(dst, src)
}
func (m *ListEntitlementPoliciesResponse) XXX_Size() int {
	return xxx_messageInfo_ListEntitlementPoliciesResponse.Size(m)
}
func (m *ListEntitlementPoliciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitlementPoliciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitlementPoliciesResponse proto.InternalMessageInfo

func (m *ListEntitlementPoliciesResponse) GetPolicies() []*ListEntitlementPoliciesResponse_Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy is a nested type used to be able to cleanly return a list of
// Policies within a single response. Each Policy instance contains the id of
// the policy, the list of entitlements defined by the policy, as well as the
// policy's public key.
type ListEntitlementPoliciesResponse_Policy struct {
	// This attribute contains the unique identifier of the policy.
	CommunityId string `protobuf:"bytes,7,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
	// This attribute contains a human friendly label describing the policy
	// suitable for rendering in the DECODE wallet
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// This field contains a list of the operations that define the policy.
	Operations []*Operation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	// This attribute contains the public key of the policy. This public key
	// attribute is the label applied to the bucket within the datastore which
	// will be how data can be downloaded for the entitlement policy.
	PublicKey string `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// This attribute contains an identifier of an authorizable attribute required
	// for taking part in the Coconut protocol.
	AuthorizableAttributeId string `protobuf:"bytes,5,opt,name=authorizable_attribute_id,json=authorizableAttributeId,proto3" json:"authorizable_attribute_id,omitempty"`
	// This attribute contains a reference to the credential issuer service the
	// end user must interact with in order to obtain blind credentials when
	// taking part in the Coconut protocol.
	CredentialIssuerEndpointUrl string `protobuf:"bytes,6,opt,name=credential_issuer_endpoint_url,json=credentialIssuerEndpointUrl,proto3" json:"credential_issuer_endpoint_url,omitempty"`
	// Multi language description of the community for which this policy applies
	// consisting of a map where the map keys are a standard language code, and
	// the value is the description in that language.
	Descriptions         map[string]string `protobuf:"bytes,8,rep,name=descriptions,proto3" json:"descriptions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListEntitlementPoliciesResponse_Policy) Reset() {
	*m = ListEntitlementPoliciesResponse_Policy{}
}
func (m *ListEntitlementPoliciesResponse_Policy) String() string { return proto.CompactTextString(m) }
func (*ListEntitlementPoliciesResponse_Policy) ProtoMessage()    {}
func (*ListEntitlementPoliciesResponse_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_policystore_61f3d1595e27eebe, []int{6, 0}
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Unmarshal(m, b)
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Marshal(b, m, deterministic)
}
func (dst *ListEntitlementPoliciesResponse_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Merge(dst, src)
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_Size() int {
	return xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.Size(m)
}
func (m *ListEntitlementPoliciesResponse_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEntitlementPoliciesResponse_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ListEntitlementPoliciesResponse_Policy proto.InternalMessageInfo

func (m *ListEntitlementPoliciesResponse_Policy) GetCommunityId() string {
	if m != nil {
		return m.CommunityId
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListEntitlementPoliciesResponse_Policy) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetAuthorizableAttributeId() string {
	if m != nil {
		return m.AuthorizableAttributeId
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetCredentialIssuerEndpointUrl() string {
	if m != nil {
		return m.CredentialIssuerEndpointUrl
	}
	return ""
}

func (m *ListEntitlementPoliciesResponse_Policy) GetDescriptions() map[string]string {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

func init() {
	proto.RegisterType((*Operation)(nil), "decode.iot.policystore.Operation")
	proto.RegisterType((*CreateEntitlementPolicyRequest)(nil), "decode.iot.policystore.CreateEntitlementPolicyRequest")
	proto.RegisterMapType((map[string]string)(nil), "decode.iot.policystore.CreateEntitlementPolicyRequest.DescriptionsEntry")
	proto.RegisterType((*CreateEntitlementPolicyResponse)(nil), "decode.iot.policystore.CreateEntitlementPolicyResponse")
	proto.RegisterType((*DeleteEntitlementPolicyRequest)(nil), "decode.iot.policystore.DeleteEntitlementPolicyRequest")
	proto.RegisterType((*DeleteEntitlementPolicyResponse)(nil), "decode.iot.policystore.DeleteEntitlementPolicyResponse")
	proto.RegisterType((*ListEntitlementPoliciesRequest)(nil), "decode.iot.policystore.ListEntitlementPoliciesRequest")
	proto.RegisterType((*ListEntitlementPoliciesResponse)(nil), "decode.iot.policystore.ListEntitlementPoliciesResponse")
	proto.RegisterType((*ListEntitlementPoliciesResponse_Policy)(nil), "decode.iot.policystore.ListEntitlementPoliciesResponse.Policy")
	proto.RegisterMapType((map[string]string)(nil), "decode.iot.policystore.ListEntitlementPoliciesResponse.Policy.DescriptionsEntry")
	proto.RegisterEnum("decode.iot.policystore.Operation_Action", Operation_Action_name, Operation_Action_value)
}

func init() { proto.RegisterFile("policystore.proto", fileDescriptor_policystore_61f3d1595e27eebe) }

var fileDescriptor_policystore_61f3d1595e27eebe = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xae, 0xe3, 0x24, 0x24, 0x13, 0x40, 0x61, 0x55, 0x95, 0x34, 0xa8, 0x21, 0xf8, 0x94, 0x93,
	0x0f, 0x54, 0xea, 0x0f, 0x87, 0xb6, 0x01, 0x22, 0x48, 0x69, 0x03, 0x32, 0x82, 0x4a, 0x5c, 0x2c,
	0xc7, 0x1e, 0xa9, 0xdb, 0x2c, 0xbb, 0xee, 0x7a, 0x8d, 0x94, 0x3e, 0x42, 0x9f, 0xa7, 0xb7, 0x3e,
	0x47, 0xdf, 0x82, 0x87, 0xa8, 0xbc, 0x76, 0x20, 0x05, 0x9c, 0xa8, 0xa5, 0xed, 0xcd, 0x3b, 0xfe,
	0xfc, 0xcd, 0x37, 0xdf, 0xce, 0x17, 0x05, 0x56, 0x42, 0xc1, 0xa8, 0x3f, 0x8e, 0x94, 0x90, 0x68,
	0x87, 0x52, 0x28, 0x41, 0x1e, 0x05, 0xe8, 0x8b, 0x00, 0x6d, 0x2a, 0x94, 0x3d, 0xf5, 0xd6, 0xfa,
	0x61, 0x40, 0xf5, 0x30, 0x44, 0xe9, 0x29, 0x2a, 0x38, 0x59, 0x83, 0x6a, 0x84, 0x3c, 0x12, 0xd2,
	0xa5, 0x41, 0xc3, 0x68, 0x1b, 0x9d, 0x25, 0xa7, 0x92, 0x16, 0xfa, 0x01, 0x79, 0x03, 0x65, 0xcf,
	0x4f, 0x60, 0x8d, 0x42, 0xdb, 0xe8, 0x2c, 0x6f, 0x76, 0xec, 0xbb, 0x39, 0xed, 0x2b, 0x3e, 0xbb,
	0xab, 0xf1, 0x4e, 0xf6, 0x1d, 0x21, 0x50, 0x1c, 0x52, 0x1e, 0x35, 0xcc, 0xb6, 0xd9, 0x31, 0x1c,
	0xfd, 0x4c, 0x9a, 0x50, 0xa1, 0x5c, 0xa1, 0xbc, 0xf0, 0x58, 0xa3, 0x98, 0x76, 0x9c, 0x9c, 0xad,
	0x97, 0x50, 0x4e, 0x19, 0x48, 0x0d, 0x16, 0x4e, 0x06, 0x07, 0x83, 0xc3, 0x0f, 0x83, 0xfa, 0x03,
	0x52, 0x85, 0xd2, 0xf1, 0x7e, 0xd7, 0xe9, 0xd5, 0x0d, 0xb2, 0x00, 0xe6, 0x76, 0x7f, 0x50, 0x2f,
	0x90, 0x65, 0x80, 0xf7, 0x87, 0xa7, 0xfd, 0xc1, 0x9e, 0xdb, 0x3d, 0xdd, 0xab, 0x9b, 0xd6, 0x77,
	0x13, 0x5a, 0x3b, 0x12, 0x3d, 0x85, 0x3d, 0xae, 0xa8, 0x62, 0x78, 0x8e, 0x5c, 0x1d, 0x69, 0x91,
	0x0e, 0x7e, 0x8e, 0x31, 0x52, 0xe4, 0x21, 0x94, 0x98, 0x37, 0x44, 0xa6, 0xc7, 0xa9, 0x3a, 0xe9,
	0x81, 0x74, 0x01, 0xc4, 0x44, 0x7f, 0xaa, 0xb4, 0xb6, 0xb9, 0x31, 0x77, 0x52, 0x67, 0xea, 0x23,
	0xb2, 0x05, 0x8f, 0xbd, 0x58, 0x7d, 0x14, 0x92, 0x7e, 0xf1, 0x86, 0x0c, 0x5d, 0x4f, 0x29, 0x49,
	0x87, 0xb1, 0xc2, 0xc4, 0xd5, 0xa2, 0x6e, 0xb6, 0x3a, 0x0d, 0xe8, 0x4e, 0xde, 0xf7, 0x03, 0xb2,
	0x03, 0x2d, 0x5f, 0x62, 0x80, 0x5c, 0x51, 0x8f, 0xb9, 0x34, 0x8a, 0x62, 0x94, 0x2e, 0xf2, 0x20,
	0x14, 0x94, 0x2b, 0x37, 0x96, 0xac, 0x51, 0xd2, 0x04, 0x6b, 0xd7, 0xa8, 0xbe, 0x06, 0xf5, 0x32,
	0xcc, 0x89, 0x64, 0x84, 0xc1, 0x62, 0x80, 0x91, 0x2f, 0x69, 0x98, 0x4e, 0x51, 0xd6, 0x53, 0xec,
	0xe7, 0x4d, 0x31, 0xdb, 0x27, 0x7b, 0x77, 0x8a, 0xaa, 0xc7, 0x95, 0x1c, 0x3b, 0xbf, 0xb0, 0x37,
	0x5f, 0xc3, 0xca, 0x2d, 0x08, 0xa9, 0x83, 0x39, 0xc2, 0xb1, 0xde, 0xa1, 0xaa, 0x93, 0x3c, 0x26,
	0x76, 0x5f, 0x78, 0x2c, 0xc6, 0x89, 0xdd, 0xfa, 0xb0, 0x55, 0x78, 0x61, 0xbc, 0x2d, 0x56, 0x8c,
	0x7a, 0xc1, 0x81, 0x30, 0x1e, 0x32, 0xea, 0xbb, 0x23, 0x1c, 0x5b, 0x23, 0x58, 0xcf, 0x15, 0x15,
	0x85, 0x82, 0x47, 0x48, 0x36, 0x60, 0xd1, 0x17, 0xe7, 0xe7, 0x31, 0xa7, 0x6a, 0x9c, 0xf8, 0x6a,
	0x6a, 0xd6, 0xda, 0x55, 0xad, 0x1f, 0x24, 0x1d, 0x95, 0x18, 0x21, 0x9f, 0x74, 0xd4, 0x87, 0xac,
	0x5b, 0x35, 0x35, 0xc0, 0xa5, 0x81, 0xf5, 0x09, 0x5a, 0xbb, 0xc8, 0x70, 0xc6, 0xa6, 0xfc, 0xbd,
	0x5e, 0x1b, 0xb0, 0x9e, 0xdb, 0x2b, 0x1d, 0xcc, 0x6a, 0x43, 0xeb, 0x1d, 0x8d, 0xd4, 0x4d, 0x00,
	0xc5, 0x28, 0x93, 0x63, 0x5d, 0x16, 0x61, 0x3d, 0x17, 0x92, 0xd9, 0x73, 0x06, 0x95, 0x30, 0xab,
	0x35, 0x0c, 0x7d, 0xfd, 0xaf, 0xf2, 0xae, 0x7f, 0x0e, 0x95, 0x9d, 0xe9, 0xbb, 0xe2, 0x6b, 0x5e,
	0x9a, 0x50, 0x4e, 0x8b, 0xb7, 0x9c, 0x59, 0xb8, 0xd3, 0x99, 0x7f, 0x13, 0xb3, 0x27, 0x30, 0xb5,
	0x32, 0x59, 0xae, 0xaa, 0x69, 0xe5, 0x00, 0xc7, 0xb3, 0x53, 0x58, 0xba, 0x6f, 0x0a, 0xcb, 0xf3,
	0x53, 0xa8, 0x6e, 0xa4, 0xb0, 0xa2, 0x87, 0x3c, 0xba, 0xdf, 0x35, 0xfc, 0xaf, 0x34, 0x5e, 0xef,
	0xec, 0xe6, 0x37, 0x13, 0x6a, 0x69, 0xf3, 0xe3, 0x44, 0x28, 0xf9, 0x6a, 0xc0, 0x6a, 0x4e, 0x3a,
	0xc9, 0xb3, 0x3f, 0xfb, 0x8d, 0x69, 0x3e, 0xff, 0xed, 0xef, 0xb2, 0x3d, 0x4f, 0xc4, 0xe4, 0x24,
	0x2a, 0x5f, 0xcc, 0xec, 0xb8, 0xe7, 0x8b, 0x99, 0x13, 0x5d, 0x2d, 0x26, 0xe7, 0x1a, 0xf3, 0xc5,
	0xcc, 0x0e, 0x7b, 0xbe, 0x98, 0x39, 0xfb, 0xb2, 0xbd, 0x74, 0x56, 0x9b, 0x82, 0x0f, 0xcb, 0xfa,
	0x7f, 0xc0, 0xd3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0xea, 0x9c, 0xf3, 0x1c, 0x08, 0x00,
	0x00,
}
