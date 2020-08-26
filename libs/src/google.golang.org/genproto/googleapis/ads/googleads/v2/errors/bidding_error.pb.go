// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/bidding_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible bidding errors.
type BiddingErrorEnum_BiddingError int32

const (
	// Enum unspecified.
	BiddingErrorEnum_UNSPECIFIED BiddingErrorEnum_BiddingError = 0
	// The received error code is not known in this version.
	BiddingErrorEnum_UNKNOWN BiddingErrorEnum_BiddingError = 1
	// Cannot transition to new bidding strategy.
	BiddingErrorEnum_BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED BiddingErrorEnum_BiddingError = 2
	// Cannot attach bidding strategy to campaign.
	BiddingErrorEnum_CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN BiddingErrorEnum_BiddingError = 7
	// Bidding strategy is not supported or cannot be used as anonymous.
	BiddingErrorEnum_INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 10
	// The type does not match the named strategy's type.
	BiddingErrorEnum_INVALID_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 14
	// The bid is invalid.
	BiddingErrorEnum_INVALID_BID BiddingErrorEnum_BiddingError = 17
	// Bidding strategy is not available for the account type.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE BiddingErrorEnum_BiddingError = 18
	// Conversion tracking is not enabled for the campaign for VBB transition.
	BiddingErrorEnum_CONVERSION_TRACKING_NOT_ENABLED BiddingErrorEnum_BiddingError = 19
	// Not enough conversions tracked for VBB transitions.
	BiddingErrorEnum_NOT_ENOUGH_CONVERSIONS BiddingErrorEnum_BiddingError = 20
	// Campaign can not be created with given bidding strategy. It can be
	// transitioned to the strategy, once eligible.
	BiddingErrorEnum_CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 21
	// Cannot target content network only as campaign uses Page One Promoted
	// bidding strategy.
	BiddingErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 23
	// Budget Optimizer and Target Spend bidding strategies are not supported
	// for campaigns with AdSchedule targeting.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE BiddingErrorEnum_BiddingError = 24
	// Pay per conversion is not available to all the customer, only few
	// whitelisted customers can use this.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER BiddingErrorEnum_BiddingError = 25
	// Pay per conversion is not allowed with Target CPA.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA BiddingErrorEnum_BiddingError = 26
	// Cannot set bidding strategy to Manual CPM for search network only
	// campaigns.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS BiddingErrorEnum_BiddingError = 27
	// The bidding strategy is not supported for use in drafts or experiments.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS BiddingErrorEnum_BiddingError = 28
	// Bidding strategy type does not support product type ad group criterion.
	BiddingErrorEnum_BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION BiddingErrorEnum_BiddingError = 29
	// Bid amount is too small.
	BiddingErrorEnum_BID_TOO_SMALL BiddingErrorEnum_BiddingError = 30
	// Bid amount is too big.
	BiddingErrorEnum_BID_TOO_BIG BiddingErrorEnum_BiddingError = 31
	// Bid has too many fractional digit precision.
	BiddingErrorEnum_BID_TOO_MANY_FRACTIONAL_DIGITS BiddingErrorEnum_BiddingError = 32
	// Invalid domain name specified.
	BiddingErrorEnum_INVALID_DOMAIN_NAME BiddingErrorEnum_BiddingError = 33
	// The field is not compatible with the payment mode.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_PAYMENT_MODE BiddingErrorEnum_BiddingError = 34
	// The field is not compatible with the budget type.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_BUDGET_TYPE BiddingErrorEnum_BiddingError = 35
	// The field is not compatible with the bidding strategy type.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 36
)

var BiddingErrorEnum_BiddingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED",
	7:  "CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN",
	10: "INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE",
	14: "INVALID_BIDDING_STRATEGY_TYPE",
	17: "INVALID_BID",
	18: "BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
	19: "CONVERSION_TRACKING_NOT_ENABLED",
	20: "NOT_ENOUGH_CONVERSIONS",
	21: "CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY",
	23: "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY",
	24: "BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE",
	25: "PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER",
	26: "PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA",
	27: "BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS",
	28: "BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS",
	29: "BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION",
	30: "BID_TOO_SMALL",
	31: "BID_TOO_BIG",
	32: "BID_TOO_MANY_FRACTIONAL_DIGITS",
	33: "INVALID_DOMAIN_NAME",
	34: "NOT_COMPATIBLE_WITH_PAYMENT_MODE",
	35: "NOT_COMPATIBLE_WITH_BUDGET_TYPE",
	36: "NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE",
}

var BiddingErrorEnum_BiddingError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED":                                     2,
	"CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN":                                  7,
	"INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE":                                     10,
	"INVALID_BIDDING_STRATEGY_TYPE":                                               14,
	"INVALID_BID":                                                                 17,
	"BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                             18,
	"CONVERSION_TRACKING_NOT_ENABLED":                                             19,
	"NOT_ENOUGH_CONVERSIONS":                                                      20,
	"CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY":                                21,
	"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY": 23,
	"BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE":                             24,
	"PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER":                               25,
	"PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA":                              26,
	"BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS":                      27,
	"BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS":                     28,
	"BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION":       29,
	"BID_TOO_SMALL":                             30,
	"BID_TOO_BIG":                               31,
	"BID_TOO_MANY_FRACTIONAL_DIGITS":            32,
	"INVALID_DOMAIN_NAME":                       33,
	"NOT_COMPATIBLE_WITH_PAYMENT_MODE":          34,
	"NOT_COMPATIBLE_WITH_BUDGET_TYPE":           35,
	"NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE": 36,
}

func (x BiddingErrorEnum_BiddingError) String() string {
	return proto.EnumName(BiddingErrorEnum_BiddingError_name, int32(x))
}

func (BiddingErrorEnum_BiddingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_489f8b87c43e0401, []int{0, 0}
}

// Container for enum describing possible bidding errors.
type BiddingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingErrorEnum) Reset()         { *m = BiddingErrorEnum{} }
func (m *BiddingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingErrorEnum) ProtoMessage()    {}
func (*BiddingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_489f8b87c43e0401, []int{0}
}

func (m *BiddingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingErrorEnum.Unmarshal(m, b)
}
func (m *BiddingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingErrorEnum.Marshal(b, m, deterministic)
}
func (m *BiddingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingErrorEnum.Merge(m, src)
}
func (m *BiddingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingErrorEnum.Size(m)
}
func (m *BiddingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.BiddingErrorEnum_BiddingError", BiddingErrorEnum_BiddingError_name, BiddingErrorEnum_BiddingError_value)
	proto.RegisterType((*BiddingErrorEnum)(nil), "google.ads.googleads.v2.errors.BiddingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/bidding_error.proto", fileDescriptor_489f8b87c43e0401)
}

var fileDescriptor_489f8b87c43e0401 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x8f, 0x1b, 0x35,
	0x14, 0x65, 0x17, 0x89, 0x56, 0x2e, 0x1f, 0x5e, 0x17, 0x28, 0x2c, 0xed, 0xb6, 0x4d, 0x2b, 0x21,
	0x3e, 0x3a, 0x81, 0x54, 0x02, 0x69, 0xfa, 0x74, 0x67, 0xec, 0x4c, 0xac, 0xcc, 0xd8, 0x96, 0xed,
	0x49, 0x08, 0x8a, 0x64, 0xa5, 0x64, 0x15, 0x45, 0x6a, 0x33, 0xab, 0xcc, 0xd2, 0x9f, 0xc3, 0x03,
	0x8f, 0xbc, 0xf2, 0x2f, 0xf8, 0x29, 0x88, 0x1f, 0x81, 0x3c, 0xce, 0x84, 0x48, 0x99, 0xa5, 0x4f,
	0xb9, 0xb9, 0x3a, 0xe7, 0xd8, 0xe7, 0xce, 0xf1, 0x45, 0x83, 0x55, 0x55, 0xad, 0x5e, 0x5d, 0xf6,
	0x17, 0xcb, 0xba, 0x1f, 0x4a, 0x5f, 0xbd, 0x19, 0xf4, 0x2f, 0xb7, 0xdb, 0x6a, 0x5b, 0xf7, 0x5f,
	0xae, 0x97, 0xcb, 0xf5, 0x66, 0xe5, 0x9a, 0xbf, 0xd1, 0xd5, 0xb6, 0xba, 0xae, 0xc8, 0x45, 0x00,
	0x46, 0x8b, 0x65, 0x1d, 0xed, 0x39, 0xd1, 0x9b, 0x41, 0x14, 0x38, 0xe7, 0xf7, 0x5b, 0xcd, 0xab,
	0x75, 0x7f, 0xb1, 0xd9, 0x54, 0xd7, 0x8b, 0xeb, 0x75, 0xb5, 0xa9, 0x03, 0xbb, 0xf7, 0xe7, 0x6d,
	0x84, 0x93, 0xa0, 0xca, 0x3c, 0x9e, 0x6d, 0x7e, 0x7d, 0xdd, 0xfb, 0xed, 0x36, 0x7a, 0xff, 0xb0,
	0x49, 0x3e, 0x42, 0x77, 0x4a, 0x61, 0x14, 0x4b, 0xf9, 0x90, 0x33, 0x8a, 0xdf, 0x21, 0x77, 0xd0,
	0xad, 0x52, 0x8c, 0x85, 0x9c, 0x0a, 0x7c, 0x42, 0xbe, 0x41, 0x5f, 0x26, 0x9c, 0x52, 0x2e, 0x32,
	0x67, 0xac, 0x06, 0xcb, 0xb2, 0x99, 0xb3, 0x1a, 0x84, 0xe1, 0x96, 0x4b, 0xe1, 0x84, 0xb4, 0x0e,
	0xf2, 0x5c, 0x4e, 0x19, 0xc5, 0xa7, 0x24, 0x42, 0x5f, 0xa7, 0x20, 0x9a, 0x9e, 0xb5, 0x90, 0x8e,
	0xdc, 0x31, 0x55, 0xba, 0x14, 0x0a, 0x05, 0x3c, 0x13, 0xf8, 0x96, 0x17, 0xe7, 0x62, 0x02, 0x39,
	0xa7, 0x0e, 0x84, 0x14, 0xb3, 0x42, 0x96, 0xa6, 0x83, 0x33, 0x53, 0x0c, 0x23, 0xf2, 0x18, 0x3d,
	0x68, 0xc1, 0xdd, 0x90, 0x0f, 0xbd, 0x95, 0x03, 0x08, 0x3e, 0x23, 0xcf, 0x51, 0xff, 0x08, 0xdb,
	0x5c, 0x6f, 0x02, 0x3c, 0x87, 0x24, 0x67, 0x6e, 0x28, 0xb5, 0x83, 0x34, 0x95, 0xa5, 0xb0, 0x41,
	0x85, 0x90, 0x27, 0xe8, 0x61, 0x2a, 0xc5, 0x84, 0x69, 0xe3, 0x1d, 0x5a, 0x0d, 0xe9, 0xd8, 0x0b,
	0x78, 0x1e, 0x13, 0x9e, 0x44, 0xf1, 0x5d, 0x72, 0x8e, 0x3e, 0x0d, 0x0d, 0x59, 0x66, 0x23, 0xf7,
	0x1f, 0xde, 0xe0, 0x8f, 0xc9, 0x77, 0xe8, 0xdb, 0xdd, 0x18, 0x52, 0xcd, 0xc0, 0xb2, 0xbd, 0x65,
	0x37, 0xe5, 0xf6, 0x78, 0x28, 0xf8, 0x13, 0x22, 0xd1, 0x78, 0xc7, 0xb0, 0xa0, 0x33, 0x66, 0xbd,
	0xa0, 0x65, 0xc2, 0x3a, 0xc1, 0xec, 0x54, 0xea, 0xb1, 0x93, 0x22, 0x9f, 0x05, 0xf6, 0x5e, 0x2b,
	0x67, 0x13, 0x96, 0x3b, 0x25, 0xd5, 0xb1, 0xe0, 0xbd, 0x1b, 0x8d, 0x9b, 0x52, 0x29, 0xa9, 0x2d,
	0xa3, 0x41, 0x0c, 0xa8, 0x33, 0xe9, 0x88, 0xd1, 0x32, 0x67, 0xf8, 0x33, 0xf2, 0x3d, 0x7a, 0xa6,
	0x60, 0xe6, 0x14, 0xd3, 0x07, 0x86, 0x3a, 0xe6, 0x95, 0x96, 0xc6, 0xca, 0x82, 0x69, 0xfc, 0x39,
	0x19, 0xa0, 0xe8, 0x26, 0x4a, 0x48, 0x45, 0x38, 0xa7, 0x75, 0xa5, 0x00, 0x9f, 0x93, 0x18, 0xfd,
	0xd0, 0xfd, 0x51, 0x76, 0x0c, 0x7f, 0x84, 0x61, 0xa0, 0xd3, 0x51, 0xb0, 0xdd, 0x3a, 0x36, 0xf8,
	0x0b, 0xf2, 0x02, 0xfd, 0xf8, 0x16, 0x5f, 0x5c, 0x38, 0xaa, 0x61, 0x68, 0x8d, 0x93, 0xda, 0xb1,
	0x9f, 0x14, 0xd3, 0xbc, 0x60, 0xc2, 0x1a, 0x7c, 0x9f, 0x70, 0xc4, 0x3a, 0x93, 0xe3, 0xa8, 0x64,
	0xe6, 0x50, 0xc6, 0x29, 0x2d, 0x69, 0x99, 0x86, 0x44, 0x38, 0xa0, 0x99, 0x96, 0xa5, 0x72, 0xa9,
	0xe6, 0x96, 0x69, 0x2e, 0x05, 0x7e, 0x40, 0xce, 0xd0, 0x07, 0x09, 0xa7, 0xce, 0x4a, 0xe9, 0x4c,
	0x01, 0x79, 0x8e, 0x2f, 0x7c, 0xf8, 0xda, 0x56, 0xc2, 0x33, 0xfc, 0x90, 0xf4, 0xd0, 0x45, 0xdb,
	0x28, 0x40, 0xcc, 0xdc, 0x50, 0x43, 0xea, 0x1f, 0x0d, 0xe4, 0x8e, 0xf2, 0x8c, 0x5b, 0x83, 0x1f,
	0x91, 0x7b, 0xe8, 0x6e, 0x9b, 0x58, 0x2a, 0x0b, 0xe0, 0xc2, 0x09, 0x28, 0x18, 0x7e, 0x4c, 0x9e,
	0xa2, 0x47, 0x4d, 0x80, 0x64, 0xa1, 0xc0, 0x72, 0x3f, 0xf9, 0x66, 0x90, 0x0a, 0x66, 0xde, 0x8d,
	0x2b, 0x24, 0x65, 0xb8, 0xe7, 0xa3, 0xda, 0x85, 0x4a, 0x4a, 0xea, 0xc7, 0xdd, 0xe4, 0xf9, 0x09,
	0x79, 0x86, 0xbe, 0xea, 0x04, 0x75, 0x3e, 0xa2, 0xa7, 0xc9, 0x3f, 0x27, 0xa8, 0xf7, 0x4b, 0xf5,
	0x3a, 0xfa, 0xff, 0xd5, 0x93, 0x9c, 0x1d, 0x2e, 0x11, 0xe5, 0xf7, 0x8d, 0x3a, 0xf9, 0x99, 0xee,
	0x48, 0xab, 0xea, 0xd5, 0x62, 0xb3, 0x8a, 0xaa, 0xed, 0xaa, 0xbf, 0xba, 0xdc, 0x34, 0xdb, 0xa8,
	0xdd, 0x79, 0x57, 0xeb, 0xfa, 0xa6, 0x15, 0xf8, 0x22, 0xfc, 0xfc, 0x7e, 0xfa, 0x6e, 0x06, 0xf0,
	0xc7, 0xe9, 0x45, 0x16, 0xc4, 0x60, 0x59, 0x47, 0xa1, 0xf4, 0xd5, 0x64, 0x10, 0x35, 0x47, 0xd6,
	0x7f, 0xb5, 0x80, 0x39, 0x2c, 0xeb, 0xf9, 0x1e, 0x30, 0x9f, 0x0c, 0xe6, 0x01, 0xf0, 0xf7, 0x69,
	0x2f, 0x74, 0xe3, 0x18, 0x96, 0x75, 0x1c, 0xef, 0x21, 0x71, 0x3c, 0x19, 0xc4, 0x71, 0x00, 0xbd,
	0x7c, 0xaf, 0xb9, 0xdd, 0xf3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16, 0xfc, 0x7c, 0x7a, 0x9f,
	0x05, 0x00, 0x00,
}
