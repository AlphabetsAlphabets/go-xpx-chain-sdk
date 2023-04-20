// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package transactions

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CreateLiquidityProviderBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsCreateLiquidityProviderBuffer(buf []byte, offset flatbuffers.UOffsetT) *CreateLiquidityProviderBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CreateLiquidityProviderBuffer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCreateLiquidityProviderBuffer(buf []byte, offset flatbuffers.UOffsetT) *CreateLiquidityProviderBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CreateLiquidityProviderBuffer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CreateLiquidityProviderBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CreateLiquidityProviderBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CreateLiquidityProviderBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *CreateLiquidityProviderBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CreateLiquidityProviderBuffer) MutateSignature(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CreateLiquidityProviderBuffer) MutateSigner(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) Version() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *CreateLiquidityProviderBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *CreateLiquidityProviderBuffer) MaxFee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MaxFeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateMaxFee(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateDeadline(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) ProviderMosaicId(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) ProviderMosaicIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateProviderMosaicId(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) CurrencyDeposit(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) CurrencyDepositLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateCurrencyDeposit(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) InitialMosaicsMinting(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) InitialMosaicsMintingLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateInitialMosaicsMinting(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) SlashingPeriod() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateSlashingPeriod(n uint32) bool {
	return rcv._tab.MutateUint32Slot(24, n)
}

func (rcv *CreateLiquidityProviderBuffer) WindowSize() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateWindowSize(n uint16) bool {
	return rcv._tab.MutateUint16Slot(26, n)
}

func (rcv *CreateLiquidityProviderBuffer) SlashingAccount(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) SlashingAccountLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) SlashingAccountBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CreateLiquidityProviderBuffer) MutateSlashingAccount(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *CreateLiquidityProviderBuffer) Alpha() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateAlpha(n uint32) bool {
	return rcv._tab.MutateUint32Slot(30, n)
}

func (rcv *CreateLiquidityProviderBuffer) Beta() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CreateLiquidityProviderBuffer) MutateBeta(n uint32) bool {
	return rcv._tab.MutateUint32Slot(32, n)
}

func CreateLiquidityProviderBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(15)
}
func CreateLiquidityProviderBufferAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(0, size, 0)
}
func CreateLiquidityProviderBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}
func CreateLiquidityProviderBufferStartSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func CreateLiquidityProviderBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}
func CreateLiquidityProviderBufferStartSignerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func CreateLiquidityProviderBufferAddVersion(builder *flatbuffers.Builder, version uint32) {
	builder.PrependUint32Slot(3, version, 0)
}
func CreateLiquidityProviderBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}
func CreateLiquidityProviderBufferAddMaxFee(builder *flatbuffers.Builder, maxFee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(maxFee), 0)
}
func CreateLiquidityProviderBufferStartMaxFeeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CreateLiquidityProviderBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}
func CreateLiquidityProviderBufferStartDeadlineVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CreateLiquidityProviderBufferAddProviderMosaicId(builder *flatbuffers.Builder, providerMosaicId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(providerMosaicId), 0)
}
func CreateLiquidityProviderBufferStartProviderMosaicIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CreateLiquidityProviderBufferAddCurrencyDeposit(builder *flatbuffers.Builder, currencyDeposit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(currencyDeposit), 0)
}
func CreateLiquidityProviderBufferStartCurrencyDepositVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CreateLiquidityProviderBufferAddInitialMosaicsMinting(builder *flatbuffers.Builder, initialMosaicsMinting flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(initialMosaicsMinting), 0)
}
func CreateLiquidityProviderBufferStartInitialMosaicsMintingVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CreateLiquidityProviderBufferAddSlashingPeriod(builder *flatbuffers.Builder, slashingPeriod uint32) {
	builder.PrependUint32Slot(10, slashingPeriod, 0)
}
func CreateLiquidityProviderBufferAddWindowSize(builder *flatbuffers.Builder, windowSize uint16) {
	builder.PrependUint16Slot(11, windowSize, 0)
}
func CreateLiquidityProviderBufferAddSlashingAccount(builder *flatbuffers.Builder, slashingAccount flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(slashingAccount), 0)
}
func CreateLiquidityProviderBufferStartSlashingAccountVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func CreateLiquidityProviderBufferAddAlpha(builder *flatbuffers.Builder, alpha uint32) {
	builder.PrependUint32Slot(13, alpha, 0)
}
func CreateLiquidityProviderBufferAddBeta(builder *flatbuffers.Builder, beta uint32) {
	builder.PrependUint32Slot(14, beta, 0)
}
func CreateLiquidityProviderBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
