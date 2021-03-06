// Code generated by protoc-gen-go. DO NOT EDIT.
// source: score.proto

package SC2APIProtocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Score_ScoreType int32

const (
	Score_Curriculum Score_ScoreType = 1
	Score_Melee      Score_ScoreType = 2
)

var Score_ScoreType_name = map[int32]string{
	1: "Curriculum",
	2: "Melee",
}
var Score_ScoreType_value = map[string]int32{
	"Curriculum": 1,
	"Melee":      2,
}

func (x Score_ScoreType) Enum() *Score_ScoreType {
	p := new(Score_ScoreType)
	*p = x
	return p
}
func (x Score_ScoreType) String() string {
	return proto.EnumName(Score_ScoreType_name, int32(x))
}
func (x *Score_ScoreType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Score_ScoreType_value, data, "Score_ScoreType")
	if err != nil {
		return err
	}
	*x = Score_ScoreType(value)
	return nil
}
func (Score_ScoreType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type Score struct {
	ScoreType        *Score_ScoreType `protobuf:"varint,6,opt,name=score_type,json=scoreType,enum=SC2APIProtocol.Score_ScoreType" json:"score_type,omitempty"`
	Score            *int32           `protobuf:"varint,7,opt,name=score" json:"score,omitempty"`
	ScoreDetails     *ScoreDetails    `protobuf:"bytes,8,opt,name=score_details,json=scoreDetails" json:"score_details,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Score) Reset()                    { *m = Score{} }
func (m *Score) String() string            { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()               {}
func (*Score) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Score) GetScoreType() Score_ScoreType {
	if m != nil && m.ScoreType != nil {
		return *m.ScoreType
	}
	return Score_Curriculum
}

func (m *Score) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Score) GetScoreDetails() *ScoreDetails {
	if m != nil {
		return m.ScoreDetails
	}
	return nil
}

type CategoryScoreDetails struct {
	None             *float32 `protobuf:"fixed32,1,opt,name=none" json:"none,omitempty"`
	Army             *float32 `protobuf:"fixed32,2,opt,name=army" json:"army,omitempty"`
	Economy          *float32 `protobuf:"fixed32,3,opt,name=economy" json:"economy,omitempty"`
	Technology       *float32 `protobuf:"fixed32,4,opt,name=technology" json:"technology,omitempty"`
	Upgrade          *float32 `protobuf:"fixed32,5,opt,name=upgrade" json:"upgrade,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CategoryScoreDetails) Reset()                    { *m = CategoryScoreDetails{} }
func (m *CategoryScoreDetails) String() string            { return proto.CompactTextString(m) }
func (*CategoryScoreDetails) ProtoMessage()               {}
func (*CategoryScoreDetails) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CategoryScoreDetails) GetNone() float32 {
	if m != nil && m.None != nil {
		return *m.None
	}
	return 0
}

func (m *CategoryScoreDetails) GetArmy() float32 {
	if m != nil && m.Army != nil {
		return *m.Army
	}
	return 0
}

func (m *CategoryScoreDetails) GetEconomy() float32 {
	if m != nil && m.Economy != nil {
		return *m.Economy
	}
	return 0
}

func (m *CategoryScoreDetails) GetTechnology() float32 {
	if m != nil && m.Technology != nil {
		return *m.Technology
	}
	return 0
}

func (m *CategoryScoreDetails) GetUpgrade() float32 {
	if m != nil && m.Upgrade != nil {
		return *m.Upgrade
	}
	return 0
}

type VitalScoreDetails struct {
	Life             *float32 `protobuf:"fixed32,1,opt,name=life" json:"life,omitempty"`
	Shields          *float32 `protobuf:"fixed32,2,opt,name=shields" json:"shields,omitempty"`
	Energy           *float32 `protobuf:"fixed32,3,opt,name=energy" json:"energy,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *VitalScoreDetails) Reset()                    { *m = VitalScoreDetails{} }
func (m *VitalScoreDetails) String() string            { return proto.CompactTextString(m) }
func (*VitalScoreDetails) ProtoMessage()               {}
func (*VitalScoreDetails) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *VitalScoreDetails) GetLife() float32 {
	if m != nil && m.Life != nil {
		return *m.Life
	}
	return 0
}

func (m *VitalScoreDetails) GetShields() float32 {
	if m != nil && m.Shields != nil {
		return *m.Shields
	}
	return 0
}

func (m *VitalScoreDetails) GetEnergy() float32 {
	if m != nil && m.Energy != nil {
		return *m.Energy
	}
	return 0
}

type ScoreDetails struct {
	IdleProductionTime *float32 `protobuf:"fixed32,1,opt,name=idle_production_time,json=idleProductionTime" json:"idle_production_time,omitempty"`
	IdleWorkerTime     *float32 `protobuf:"fixed32,2,opt,name=idle_worker_time,json=idleWorkerTime" json:"idle_worker_time,omitempty"`
	// Note the "total_value" fields are a combination of minerals, vespene and a human designer guess. Maybe useful as a delta.
	TotalValueUnits      *float32 `protobuf:"fixed32,3,opt,name=total_value_units,json=totalValueUnits" json:"total_value_units,omitempty"`
	TotalValueStructures *float32 `protobuf:"fixed32,4,opt,name=total_value_structures,json=totalValueStructures" json:"total_value_structures,omitempty"`
	// Note the "killed_value" fields are a combination of minerals, vespene and a human designer guess. Maybe useful as a delta.
	// The weighting of the combination and the human designer guess is different (not symmetric) with the "total_value" fields!
	KilledValueUnits       *float32              `protobuf:"fixed32,5,opt,name=killed_value_units,json=killedValueUnits" json:"killed_value_units,omitempty"`
	KilledValueStructures  *float32              `protobuf:"fixed32,6,opt,name=killed_value_structures,json=killedValueStructures" json:"killed_value_structures,omitempty"`
	CollectedMinerals      *float32              `protobuf:"fixed32,7,opt,name=collected_minerals,json=collectedMinerals" json:"collected_minerals,omitempty"`
	CollectedVespene       *float32              `protobuf:"fixed32,8,opt,name=collected_vespene,json=collectedVespene" json:"collected_vespene,omitempty"`
	CollectionRateMinerals *float32              `protobuf:"fixed32,9,opt,name=collection_rate_minerals,json=collectionRateMinerals" json:"collection_rate_minerals,omitempty"`
	CollectionRateVespene  *float32              `protobuf:"fixed32,10,opt,name=collection_rate_vespene,json=collectionRateVespene" json:"collection_rate_vespene,omitempty"`
	SpentMinerals          *float32              `protobuf:"fixed32,11,opt,name=spent_minerals,json=spentMinerals" json:"spent_minerals,omitempty"`
	SpentVespene           *float32              `protobuf:"fixed32,12,opt,name=spent_vespene,json=spentVespene" json:"spent_vespene,omitempty"`
	FoodUsed               *CategoryScoreDetails `protobuf:"bytes,13,opt,name=food_used,json=foodUsed" json:"food_used,omitempty"`
	KilledMinerals         *CategoryScoreDetails `protobuf:"bytes,14,opt,name=killed_minerals,json=killedMinerals" json:"killed_minerals,omitempty"`
	KilledVespene          *CategoryScoreDetails `protobuf:"bytes,15,opt,name=killed_vespene,json=killedVespene" json:"killed_vespene,omitempty"`
	LostMinerals           *CategoryScoreDetails `protobuf:"bytes,16,opt,name=lost_minerals,json=lostMinerals" json:"lost_minerals,omitempty"`
	LostVespene            *CategoryScoreDetails `protobuf:"bytes,17,opt,name=lost_vespene,json=lostVespene" json:"lost_vespene,omitempty"`
	FriendlyFireMinerals   *CategoryScoreDetails `protobuf:"bytes,18,opt,name=friendly_fire_minerals,json=friendlyFireMinerals" json:"friendly_fire_minerals,omitempty"`
	FriendlyFireVespene    *CategoryScoreDetails `protobuf:"bytes,19,opt,name=friendly_fire_vespene,json=friendlyFireVespene" json:"friendly_fire_vespene,omitempty"`
	UsedMinerals           *CategoryScoreDetails `protobuf:"bytes,20,opt,name=used_minerals,json=usedMinerals" json:"used_minerals,omitempty"`
	UsedVespene            *CategoryScoreDetails `protobuf:"bytes,21,opt,name=used_vespene,json=usedVespene" json:"used_vespene,omitempty"`
	TotalUsedMinerals      *CategoryScoreDetails `protobuf:"bytes,22,opt,name=total_used_minerals,json=totalUsedMinerals" json:"total_used_minerals,omitempty"`
	TotalUsedVespene       *CategoryScoreDetails `protobuf:"bytes,23,opt,name=total_used_vespene,json=totalUsedVespene" json:"total_used_vespene,omitempty"`
	TotalDamageDealt       *VitalScoreDetails    `protobuf:"bytes,24,opt,name=total_damage_dealt,json=totalDamageDealt" json:"total_damage_dealt,omitempty"`
	TotalDamageTaken       *VitalScoreDetails    `protobuf:"bytes,25,opt,name=total_damage_taken,json=totalDamageTaken" json:"total_damage_taken,omitempty"`
	TotalHealed            *VitalScoreDetails    `protobuf:"bytes,26,opt,name=total_healed,json=totalHealed" json:"total_healed,omitempty"`
	XXX_unrecognized       []byte                `json:"-"`
}

func (m *ScoreDetails) Reset()                    { *m = ScoreDetails{} }
func (m *ScoreDetails) String() string            { return proto.CompactTextString(m) }
func (*ScoreDetails) ProtoMessage()               {}
func (*ScoreDetails) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *ScoreDetails) GetIdleProductionTime() float32 {
	if m != nil && m.IdleProductionTime != nil {
		return *m.IdleProductionTime
	}
	return 0
}

func (m *ScoreDetails) GetIdleWorkerTime() float32 {
	if m != nil && m.IdleWorkerTime != nil {
		return *m.IdleWorkerTime
	}
	return 0
}

func (m *ScoreDetails) GetTotalValueUnits() float32 {
	if m != nil && m.TotalValueUnits != nil {
		return *m.TotalValueUnits
	}
	return 0
}

func (m *ScoreDetails) GetTotalValueStructures() float32 {
	if m != nil && m.TotalValueStructures != nil {
		return *m.TotalValueStructures
	}
	return 0
}

func (m *ScoreDetails) GetKilledValueUnits() float32 {
	if m != nil && m.KilledValueUnits != nil {
		return *m.KilledValueUnits
	}
	return 0
}

func (m *ScoreDetails) GetKilledValueStructures() float32 {
	if m != nil && m.KilledValueStructures != nil {
		return *m.KilledValueStructures
	}
	return 0
}

func (m *ScoreDetails) GetCollectedMinerals() float32 {
	if m != nil && m.CollectedMinerals != nil {
		return *m.CollectedMinerals
	}
	return 0
}

func (m *ScoreDetails) GetCollectedVespene() float32 {
	if m != nil && m.CollectedVespene != nil {
		return *m.CollectedVespene
	}
	return 0
}

func (m *ScoreDetails) GetCollectionRateMinerals() float32 {
	if m != nil && m.CollectionRateMinerals != nil {
		return *m.CollectionRateMinerals
	}
	return 0
}

func (m *ScoreDetails) GetCollectionRateVespene() float32 {
	if m != nil && m.CollectionRateVespene != nil {
		return *m.CollectionRateVespene
	}
	return 0
}

func (m *ScoreDetails) GetSpentMinerals() float32 {
	if m != nil && m.SpentMinerals != nil {
		return *m.SpentMinerals
	}
	return 0
}

func (m *ScoreDetails) GetSpentVespene() float32 {
	if m != nil && m.SpentVespene != nil {
		return *m.SpentVespene
	}
	return 0
}

func (m *ScoreDetails) GetFoodUsed() *CategoryScoreDetails {
	if m != nil {
		return m.FoodUsed
	}
	return nil
}

func (m *ScoreDetails) GetKilledMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.KilledMinerals
	}
	return nil
}

func (m *ScoreDetails) GetKilledVespene() *CategoryScoreDetails {
	if m != nil {
		return m.KilledVespene
	}
	return nil
}

func (m *ScoreDetails) GetLostMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.LostMinerals
	}
	return nil
}

func (m *ScoreDetails) GetLostVespene() *CategoryScoreDetails {
	if m != nil {
		return m.LostVespene
	}
	return nil
}

func (m *ScoreDetails) GetFriendlyFireMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.FriendlyFireMinerals
	}
	return nil
}

func (m *ScoreDetails) GetFriendlyFireVespene() *CategoryScoreDetails {
	if m != nil {
		return m.FriendlyFireVespene
	}
	return nil
}

func (m *ScoreDetails) GetUsedMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.UsedMinerals
	}
	return nil
}

func (m *ScoreDetails) GetUsedVespene() *CategoryScoreDetails {
	if m != nil {
		return m.UsedVespene
	}
	return nil
}

func (m *ScoreDetails) GetTotalUsedMinerals() *CategoryScoreDetails {
	if m != nil {
		return m.TotalUsedMinerals
	}
	return nil
}

func (m *ScoreDetails) GetTotalUsedVespene() *CategoryScoreDetails {
	if m != nil {
		return m.TotalUsedVespene
	}
	return nil
}

func (m *ScoreDetails) GetTotalDamageDealt() *VitalScoreDetails {
	if m != nil {
		return m.TotalDamageDealt
	}
	return nil
}

func (m *ScoreDetails) GetTotalDamageTaken() *VitalScoreDetails {
	if m != nil {
		return m.TotalDamageTaken
	}
	return nil
}

func (m *ScoreDetails) GetTotalHealed() *VitalScoreDetails {
	if m != nil {
		return m.TotalHealed
	}
	return nil
}

func init() {
	proto.RegisterType((*Score)(nil), "SC2APIProtocol.Score")
	proto.RegisterType((*CategoryScoreDetails)(nil), "SC2APIProtocol.CategoryScoreDetails")
	proto.RegisterType((*VitalScoreDetails)(nil), "SC2APIProtocol.VitalScoreDetails")
	proto.RegisterType((*ScoreDetails)(nil), "SC2APIProtocol.ScoreDetails")
	proto.RegisterEnum("SC2APIProtocol.Score_ScoreType", Score_ScoreType_name, Score_ScoreType_value)
}

func init() { proto.RegisterFile("score.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0xc7, 0x65, 0x36, 0x24, 0xe1, 0x00, 0x0e, 0x4c, 0x08, 0xf1, 0xae, 0x56, 0xbb, 0x2c, 0xfb,
	0x21, 0xb4, 0xbb, 0x8d, 0xaa, 0xa8, 0x8a, 0x7a, 0x55, 0x29, 0x0a, 0x6a, 0x1b, 0x55, 0x51, 0x23,
	0x87, 0xa4, 0x1f, 0x37, 0x96, 0x65, 0x1f, 0x88, 0x95, 0xc1, 0x83, 0xc6, 0xe3, 0x54, 0xbc, 0x46,
	0x6f, 0xfb, 0x42, 0x7d, 0xac, 0x6a, 0x66, 0x3c, 0x66, 0x20, 0xb9, 0xa8, 0x7b, 0x83, 0xe6, 0x7c,
	0xfc, 0x7f, 0xe7, 0xcf, 0x61, 0x8c, 0xa1, 0x99, 0x45, 0x8c, 0xe3, 0xd1, 0x82, 0x33, 0xc1, 0x88,
	0x7b, 0x75, 0x76, 0x7c, 0x7a, 0x79, 0x7e, 0x29, 0x83, 0x88, 0xd1, 0xe1, 0x57, 0x07, 0xea, 0x57,
	0xb2, 0x4e, 0x5e, 0x00, 0xa8, 0xc6, 0x40, 0x2c, 0x17, 0xe8, 0x6d, 0x0f, 0x9c, 0x91, 0x7b, 0xfc,
	0xfb, 0xd1, 0x7a, 0xfb, 0x91, 0x6a, 0xd5, 0x9f, 0x93, 0xe5, 0x02, 0xfd, 0x46, 0x66, 0x8e, 0xa4,
	0x07, 0x75, 0x15, 0x78, 0x3b, 0x03, 0x67, 0x54, 0xf7, 0x75, 0x40, 0x4e, 0xa1, 0xad, 0xa9, 0x31,
	0x8a, 0x30, 0xa1, 0x99, 0xb7, 0x3b, 0x70, 0x46, 0xcd, 0xe3, 0x5f, 0x1f, 0x05, 0x8f, 0x75, 0x8f,
	0xdf, 0xca, 0xac, 0x68, 0xf8, 0x0f, 0x34, 0xca, 0x81, 0xc4, 0x05, 0x38, 0xcb, 0x39, 0x4f, 0xa2,
	0x9c, 0xe6, 0xf3, 0x8e, 0x43, 0x1a, 0x50, 0xbf, 0x40, 0x8a, 0xd8, 0xa9, 0x0d, 0x3f, 0x3b, 0xd0,
	0x3b, 0x0b, 0x05, 0xce, 0x18, 0x5f, 0xda, 0x38, 0x42, 0x60, 0x2b, 0x65, 0x29, 0x7a, 0xce, 0xc0,
	0x19, 0xd5, 0x7c, 0x75, 0x96, 0xb9, 0x90, 0xcf, 0x97, 0x5e, 0x4d, 0xe7, 0xe4, 0x99, 0x78, 0xb0,
	0x83, 0x11, 0x4b, 0xd9, 0x7c, 0xe9, 0xfd, 0xa4, 0xd2, 0x26, 0x24, 0xbf, 0x01, 0x08, 0x8c, 0x6e,
	0x53, 0x46, 0xd9, 0x6c, 0xe9, 0x6d, 0xa9, 0xa2, 0x95, 0x91, 0xca, 0x7c, 0x31, 0xe3, 0x61, 0x8c,
	0x5e, 0x5d, 0x2b, 0x8b, 0x70, 0xf8, 0x01, 0xba, 0x37, 0x89, 0x08, 0xe9, 0xa6, 0x21, 0x9a, 0x4c,
	0x4b, 0x43, 0xf2, 0x2c, 0x11, 0xd9, 0x6d, 0x82, 0x34, 0xce, 0x0a, 0x4f, 0x26, 0x24, 0x7d, 0xd8,
	0xc6, 0x14, 0xf9, 0xcc, 0xb8, 0x2a, 0xa2, 0xe1, 0x97, 0x36, 0xb4, 0xd6, 0xb0, 0x4f, 0xa1, 0x97,
	0xc4, 0x14, 0x83, 0x05, 0x67, 0x71, 0x1e, 0x89, 0x84, 0xa5, 0x81, 0x48, 0xe6, 0x66, 0x0c, 0x91,
	0xb5, 0xcb, 0xb2, 0x34, 0x49, 0xe6, 0x48, 0x46, 0xd0, 0x51, 0x8a, 0x4f, 0x8c, 0xdf, 0x21, 0xd7,
	0xdd, 0x7a, 0xba, 0x2b, 0xf3, 0xef, 0x54, 0x5a, 0x75, 0xfe, 0x0b, 0x5d, 0xc1, 0x44, 0x48, 0x83,
	0xfb, 0x90, 0xe6, 0x18, 0xe4, 0x69, 0x22, 0xb2, 0xc2, 0xcf, 0x9e, 0x2a, 0xdc, 0xc8, 0xfc, 0xb5,
	0x4c, 0x93, 0x67, 0xd0, 0xb7, 0x7b, 0x33, 0xc1, 0xf3, 0x48, 0xe4, 0x1c, 0xb3, 0x62, 0x73, 0xbd,
	0x95, 0xe0, 0xaa, 0xac, 0x91, 0xff, 0x81, 0xdc, 0x25, 0x94, 0x62, 0xbc, 0x36, 0x42, 0xaf, 0xb3,
	0xa3, 0x2b, 0xd6, 0x8c, 0x13, 0x38, 0x5c, 0xeb, 0xb6, 0x86, 0x6c, 0x2b, 0xc9, 0x81, 0x25, 0xb1,
	0xa6, 0x3c, 0x01, 0x12, 0x31, 0x4a, 0x31, 0x12, 0x18, 0x07, 0xf3, 0x24, 0x45, 0x1e, 0xd2, 0x4c,
	0x5d, 0xd9, 0x9a, 0xdf, 0x2d, 0x2b, 0x17, 0x45, 0x81, 0xfc, 0x07, 0xab, 0x64, 0x70, 0x8f, 0xd9,
	0x02, 0x53, 0x54, 0x57, 0xb8, 0xe6, 0x77, 0xca, 0xc2, 0x8d, 0xce, 0x93, 0xe7, 0xe0, 0x15, 0x39,
	0xb9, 0x7a, 0x1e, 0x0a, 0x5c, 0x4d, 0x68, 0x28, 0x4d, 0x7f, 0x55, 0xf7, 0x43, 0x81, 0xe5, 0x98,
	0x13, 0x38, 0xdc, 0x54, 0x9a, 0x61, 0xa0, 0xbf, 0xcd, 0xba, 0xd0, 0x4c, 0xfc, 0x1b, 0x5c, 0x79,
	0x10, 0xab, 0x39, 0x4d, 0xd5, 0xde, 0x56, 0xd9, 0x12, 0xff, 0x27, 0xe8, 0x44, 0x09, 0x6d, 0xa9,
	0xae, 0x96, 0x4a, 0x1a, 0xd6, 0x29, 0x34, 0xa6, 0x8c, 0xc5, 0x41, 0x9e, 0x61, 0xec, 0xb5, 0xd5,
	0x53, 0xfa, 0xd7, 0xe6, 0x53, 0xfa, 0xd8, 0xe3, 0xe5, 0xef, 0x4a, 0xd9, 0x75, 0x86, 0x31, 0xb9,
	0x80, 0xbd, 0xe2, 0x47, 0x29, 0xfd, 0xb8, 0x15, 0x40, 0xae, 0x16, 0x97, 0xb6, 0xdf, 0x80, 0x6b,
	0x7e, 0xe3, 0xc2, 0xf7, 0x5e, 0x05, 0x5a, 0xbb, 0xb8, 0x00, 0xc5, 0xd7, 0x3b, 0x87, 0x36, 0x65,
	0x99, 0xb5, 0xa9, 0x4e, 0x05, 0x56, 0x4b, 0x4a, 0x4b, 0x5f, 0xaf, 0x40, 0xc5, 0xa5, 0xab, 0x6e,
	0x05, 0x52, 0x53, 0x2a, 0x8d, 0xa7, 0x8f, 0xd0, 0x9f, 0xf2, 0x04, 0xd3, 0x98, 0x2e, 0x83, 0x69,
	0xc2, 0xad, 0xeb, 0x42, 0x2a, 0x20, 0x7b, 0x86, 0xf1, 0x32, 0xe1, 0xab, 0x2b, 0xf5, 0x1e, 0x0e,
	0xd6, 0xd9, 0xc6, 0xed, 0x7e, 0x05, 0xf4, 0xbe, 0x8d, 0xb6, 0x36, 0x29, 0xef, 0xc8, 0xca, 0x6c,
	0xaf, 0xca, 0x26, 0xa5, 0xd4, 0xde, 0xa4, 0x42, 0x19, 0x6f, 0x07, 0x55, 0x36, 0x29, 0x95, 0xc6,
	0xd3, 0x04, 0xf6, 0xf5, 0x5f, 0xce, 0xba, 0xb3, 0x7e, 0x05, 0x9e, 0xfe, 0x7f, 0xbb, 0xb6, 0xed,
	0xf9, 0x40, 0x2c, 0xaa, 0x31, 0x79, 0x58, 0x01, 0xda, 0x29, 0xa1, 0xc6, 0xe9, 0x5b, 0xc3, 0x8c,
	0xc3, 0x79, 0x38, 0x93, 0xef, 0xc5, 0x90, 0x0a, 0xcf, 0x53, 0xcc, 0x3f, 0x36, 0x99, 0x0f, 0x5e,
	0x1d, 0x05, 0x70, 0xac, 0xb4, 0x63, 0x29, 0x7d, 0x00, 0x14, 0xe1, 0x1d, 0xa6, 0xde, 0xcf, 0x3f,
	0x02, 0x9c, 0x48, 0x29, 0x19, 0x43, 0x4b, 0x03, 0x6f, 0x31, 0xa4, 0x18, 0x7b, 0xbf, 0x7c, 0x2f,
	0xaa, 0xa9, 0x64, 0xaf, 0x95, 0xea, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x28, 0xaa, 0x9c,
	0x76, 0x08, 0x00, 0x00,
}
