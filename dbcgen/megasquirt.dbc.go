// Package megasquirtcan provides primitives for encoding and decoding megasquirt CAN messages.
//
// Source: dbc/megasquirt.dbc
package megasquirtcan

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/candebug"
	"go.einride.tech/can/pkg/canrunner"
	"go.einride.tech/can/pkg/cantext"
	"go.einride.tech/can/pkg/descriptor"
	"go.einride.tech/can/pkg/generated"
	"go.einride.tech/can/pkg/socketcan"
)

// prevent unused imports
var (
	_ = context.Background
	_ = fmt.Print
	_ = net.Dial
	_ = http.Error
	_ = sync.Mutex{}
	_ = time.Now
	_ = socketcan.Dial
	_ = candebug.ServeMessagesHTTP
	_ = canrunner.Run
)

// Generated code. DO NOT EDIT.
// megasquirt_gp0Reader provides read access to a megasquirt_gp0 message.
type megasquirt_gp0Reader interface {
	// seconds returns the value of the seconds signal.
	seconds() uint16
	// pw1 returns the physical value of the pw1 signal.
	pw1() float64
	// pw2 returns the physical value of the pw2 signal.
	pw2() float64
	// rpm returns the value of the rpm signal.
	rpm() uint16
}

// megasquirt_gp0Writer provides write access to a megasquirt_gp0 message.
type megasquirt_gp0Writer interface {
	// CopyFrom copies all values from megasquirt_gp0Reader.
	CopyFrom(megasquirt_gp0Reader) *megasquirt_gp0
	// Setseconds sets the value of the seconds signal.
	Setseconds(uint16) *megasquirt_gp0
	// Setpw1 sets the physical value of the pw1 signal.
	Setpw1(float64) *megasquirt_gp0
	// Setpw2 sets the physical value of the pw2 signal.
	Setpw2(float64) *megasquirt_gp0
	// Setrpm sets the value of the rpm signal.
	Setrpm(uint16) *megasquirt_gp0
}

type megasquirt_gp0 struct {
	xxx_seconds uint16
	xxx_pw1     uint16
	xxx_pw2     uint16
	xxx_rpm     uint16
}

func Newmegasquirt_gp0() *megasquirt_gp0 {
	m := &megasquirt_gp0{}
	m.Reset()
	return m
}

func (m *megasquirt_gp0) Reset() {
	m.xxx_seconds = 0
	m.xxx_pw1 = 0
	m.xxx_pw2 = 0
	m.xxx_rpm = 0
}

func (m *megasquirt_gp0) CopyFrom(o megasquirt_gp0Reader) *megasquirt_gp0 {
	m.xxx_seconds = o.seconds()
	m.Setpw1(o.pw1())
	m.Setpw2(o.pw2())
	m.xxx_rpm = o.rpm()
	return m
}

// Descriptor returns the megasquirt_gp0 descriptor.
func (m *megasquirt_gp0) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp0.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp0) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp0) seconds() uint16 {
	return m.xxx_seconds
}

func (m *megasquirt_gp0) Setseconds(v uint16) *megasquirt_gp0 {
	m.xxx_seconds = uint16(Messages().megasquirt_gp0.seconds.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp0) pw1() float64 {
	return Messages().megasquirt_gp0.pw1.ToPhysical(float64(m.xxx_pw1))
}

func (m *megasquirt_gp0) Setpw1(v float64) *megasquirt_gp0 {
	m.xxx_pw1 = uint16(Messages().megasquirt_gp0.pw1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp0) pw2() float64 {
	return Messages().megasquirt_gp0.pw2.ToPhysical(float64(m.xxx_pw2))
}

func (m *megasquirt_gp0) Setpw2(v float64) *megasquirt_gp0 {
	m.xxx_pw2 = uint16(Messages().megasquirt_gp0.pw2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp0) rpm() uint16 {
	return m.xxx_rpm
}

func (m *megasquirt_gp0) Setrpm(v uint16) *megasquirt_gp0 {
	m.xxx_rpm = uint16(Messages().megasquirt_gp0.rpm.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp0) Frame() can.Frame {
	md := Messages().megasquirt_gp0
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.seconds.MarshalUnsigned(&f.Data, uint64(m.xxx_seconds))
	md.pw1.MarshalUnsigned(&f.Data, uint64(m.xxx_pw1))
	md.pw2.MarshalUnsigned(&f.Data, uint64(m.xxx_pw2))
	md.rpm.MarshalUnsigned(&f.Data, uint64(m.xxx_rpm))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp0) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp0) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp0
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp0: expects ID 1520 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp0: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp0: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp0: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_seconds = uint16(md.seconds.UnmarshalUnsigned(f.Data))
	m.xxx_pw1 = uint16(md.pw1.UnmarshalUnsigned(f.Data))
	m.xxx_pw2 = uint16(md.pw2.UnmarshalUnsigned(f.Data))
	m.xxx_rpm = uint16(md.rpm.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp1Reader provides read access to a megasquirt_gp1 message.
type megasquirt_gp1Reader interface {
	// adv_deg returns the physical value of the adv_deg signal.
	adv_deg() float64
	// squirt returns the value of the squirt signal.
	squirt() uint8
	// engine returns the value of the engine signal.
	engine() uint8
	// afrtgt1 returns the physical value of the afrtgt1 signal.
	afrtgt1() float64
	// afrtgt2 returns the physical value of the afrtgt2 signal.
	afrtgt2() float64
	// wbo2_en1 returns the value of the wbo2_en1 signal.
	wbo2_en1() uint8
	// wbo2_en2 returns the value of the wbo2_en2 signal.
	wbo2_en2() uint8
}

// megasquirt_gp1Writer provides write access to a megasquirt_gp1 message.
type megasquirt_gp1Writer interface {
	// CopyFrom copies all values from megasquirt_gp1Reader.
	CopyFrom(megasquirt_gp1Reader) *megasquirt_gp1
	// Setadv_deg sets the physical value of the adv_deg signal.
	Setadv_deg(float64) *megasquirt_gp1
	// Setsquirt sets the value of the squirt signal.
	Setsquirt(uint8) *megasquirt_gp1
	// Setengine sets the value of the engine signal.
	Setengine(uint8) *megasquirt_gp1
	// Setafrtgt1 sets the physical value of the afrtgt1 signal.
	Setafrtgt1(float64) *megasquirt_gp1
	// Setafrtgt2 sets the physical value of the afrtgt2 signal.
	Setafrtgt2(float64) *megasquirt_gp1
	// Setwbo2_en1 sets the value of the wbo2_en1 signal.
	Setwbo2_en1(uint8) *megasquirt_gp1
	// Setwbo2_en2 sets the value of the wbo2_en2 signal.
	Setwbo2_en2(uint8) *megasquirt_gp1
}

type megasquirt_gp1 struct {
	xxx_adv_deg  int16
	xxx_squirt   uint8
	xxx_engine   uint8
	xxx_afrtgt1  uint8
	xxx_afrtgt2  uint8
	xxx_wbo2_en1 uint8
	xxx_wbo2_en2 uint8
}

func Newmegasquirt_gp1() *megasquirt_gp1 {
	m := &megasquirt_gp1{}
	m.Reset()
	return m
}

func (m *megasquirt_gp1) Reset() {
	m.xxx_adv_deg = 0
	m.xxx_squirt = 0
	m.xxx_engine = 0
	m.xxx_afrtgt1 = 0
	m.xxx_afrtgt2 = 0
	m.xxx_wbo2_en1 = 0
	m.xxx_wbo2_en2 = 0
}

func (m *megasquirt_gp1) CopyFrom(o megasquirt_gp1Reader) *megasquirt_gp1 {
	m.Setadv_deg(o.adv_deg())
	m.xxx_squirt = o.squirt()
	m.xxx_engine = o.engine()
	m.Setafrtgt1(o.afrtgt1())
	m.Setafrtgt2(o.afrtgt2())
	m.xxx_wbo2_en1 = o.wbo2_en1()
	m.xxx_wbo2_en2 = o.wbo2_en2()
	return m
}

// Descriptor returns the megasquirt_gp1 descriptor.
func (m *megasquirt_gp1) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp1.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp1) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp1) adv_deg() float64 {
	return Messages().megasquirt_gp1.adv_deg.ToPhysical(float64(m.xxx_adv_deg))
}

func (m *megasquirt_gp1) Setadv_deg(v float64) *megasquirt_gp1 {
	m.xxx_adv_deg = int16(Messages().megasquirt_gp1.adv_deg.FromPhysical(v))
	return m
}

func (m *megasquirt_gp1) squirt() uint8 {
	return m.xxx_squirt
}

func (m *megasquirt_gp1) Setsquirt(v uint8) *megasquirt_gp1 {
	m.xxx_squirt = uint8(Messages().megasquirt_gp1.squirt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp1) engine() uint8 {
	return m.xxx_engine
}

func (m *megasquirt_gp1) Setengine(v uint8) *megasquirt_gp1 {
	m.xxx_engine = uint8(Messages().megasquirt_gp1.engine.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp1) afrtgt1() float64 {
	return Messages().megasquirt_gp1.afrtgt1.ToPhysical(float64(m.xxx_afrtgt1))
}

func (m *megasquirt_gp1) Setafrtgt1(v float64) *megasquirt_gp1 {
	m.xxx_afrtgt1 = uint8(Messages().megasquirt_gp1.afrtgt1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp1) afrtgt2() float64 {
	return Messages().megasquirt_gp1.afrtgt2.ToPhysical(float64(m.xxx_afrtgt2))
}

func (m *megasquirt_gp1) Setafrtgt2(v float64) *megasquirt_gp1 {
	m.xxx_afrtgt2 = uint8(Messages().megasquirt_gp1.afrtgt2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp1) wbo2_en1() uint8 {
	return m.xxx_wbo2_en1
}

func (m *megasquirt_gp1) Setwbo2_en1(v uint8) *megasquirt_gp1 {
	m.xxx_wbo2_en1 = uint8(Messages().megasquirt_gp1.wbo2_en1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp1) wbo2_en2() uint8 {
	return m.xxx_wbo2_en2
}

func (m *megasquirt_gp1) Setwbo2_en2(v uint8) *megasquirt_gp1 {
	m.xxx_wbo2_en2 = uint8(Messages().megasquirt_gp1.wbo2_en2.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp1) Frame() can.Frame {
	md := Messages().megasquirt_gp1
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.adv_deg.MarshalSigned(&f.Data, int64(m.xxx_adv_deg))
	md.squirt.MarshalUnsigned(&f.Data, uint64(m.xxx_squirt))
	md.engine.MarshalUnsigned(&f.Data, uint64(m.xxx_engine))
	md.afrtgt1.MarshalUnsigned(&f.Data, uint64(m.xxx_afrtgt1))
	md.afrtgt2.MarshalUnsigned(&f.Data, uint64(m.xxx_afrtgt2))
	md.wbo2_en1.MarshalUnsigned(&f.Data, uint64(m.xxx_wbo2_en1))
	md.wbo2_en2.MarshalUnsigned(&f.Data, uint64(m.xxx_wbo2_en2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp1) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp1) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp1
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp1: expects ID 1521 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp1: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp1: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp1: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_adv_deg = int16(md.adv_deg.UnmarshalSigned(f.Data))
	m.xxx_squirt = uint8(md.squirt.UnmarshalUnsigned(f.Data))
	m.xxx_engine = uint8(md.engine.UnmarshalUnsigned(f.Data))
	m.xxx_afrtgt1 = uint8(md.afrtgt1.UnmarshalUnsigned(f.Data))
	m.xxx_afrtgt2 = uint8(md.afrtgt2.UnmarshalUnsigned(f.Data))
	m.xxx_wbo2_en1 = uint8(md.wbo2_en1.UnmarshalUnsigned(f.Data))
	m.xxx_wbo2_en2 = uint8(md.wbo2_en2.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp2Reader provides read access to a megasquirt_gp2 message.
type megasquirt_gp2Reader interface {
	// baro returns the physical value of the baro signal.
	baro() float64
	// map1 returns the physical value of the map1 signal.
	map1() float64
	// mat returns the physical value of the mat signal.
	mat() float64
	// clt returns the physical value of the clt signal.
	clt() float64
}

// megasquirt_gp2Writer provides write access to a megasquirt_gp2 message.
type megasquirt_gp2Writer interface {
	// CopyFrom copies all values from megasquirt_gp2Reader.
	CopyFrom(megasquirt_gp2Reader) *megasquirt_gp2
	// Setbaro sets the physical value of the baro signal.
	Setbaro(float64) *megasquirt_gp2
	// Setmap1 sets the physical value of the map1 signal.
	Setmap1(float64) *megasquirt_gp2
	// Setmat sets the physical value of the mat signal.
	Setmat(float64) *megasquirt_gp2
	// Setclt sets the physical value of the clt signal.
	Setclt(float64) *megasquirt_gp2
}

type megasquirt_gp2 struct {
	xxx_baro int16
	xxx_map1 int16
	xxx_mat  int16
	xxx_clt  int16
}

func Newmegasquirt_gp2() *megasquirt_gp2 {
	m := &megasquirt_gp2{}
	m.Reset()
	return m
}

func (m *megasquirt_gp2) Reset() {
	m.xxx_baro = 0
	m.xxx_map1 = 0
	m.xxx_mat = 0
	m.xxx_clt = 0
}

func (m *megasquirt_gp2) CopyFrom(o megasquirt_gp2Reader) *megasquirt_gp2 {
	m.Setbaro(o.baro())
	m.Setmap1(o.map1())
	m.Setmat(o.mat())
	m.Setclt(o.clt())
	return m
}

// Descriptor returns the megasquirt_gp2 descriptor.
func (m *megasquirt_gp2) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp2.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp2) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp2) baro() float64 {
	return Messages().megasquirt_gp2.baro.ToPhysical(float64(m.xxx_baro))
}

func (m *megasquirt_gp2) Setbaro(v float64) *megasquirt_gp2 {
	m.xxx_baro = int16(Messages().megasquirt_gp2.baro.FromPhysical(v))
	return m
}

func (m *megasquirt_gp2) map1() float64 {
	return Messages().megasquirt_gp2.map1.ToPhysical(float64(m.xxx_map1))
}

func (m *megasquirt_gp2) Setmap1(v float64) *megasquirt_gp2 {
	m.xxx_map1 = int16(Messages().megasquirt_gp2.map1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp2) mat() float64 {
	return Messages().megasquirt_gp2.mat.ToPhysical(float64(m.xxx_mat))
}

func (m *megasquirt_gp2) Setmat(v float64) *megasquirt_gp2 {
	m.xxx_mat = int16(Messages().megasquirt_gp2.mat.FromPhysical(v))
	return m
}

func (m *megasquirt_gp2) clt() float64 {
	return Messages().megasquirt_gp2.clt.ToPhysical(float64(m.xxx_clt))
}

func (m *megasquirt_gp2) Setclt(v float64) *megasquirt_gp2 {
	m.xxx_clt = int16(Messages().megasquirt_gp2.clt.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp2) Frame() can.Frame {
	md := Messages().megasquirt_gp2
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.baro.MarshalSigned(&f.Data, int64(m.xxx_baro))
	md.map1.MarshalSigned(&f.Data, int64(m.xxx_map1))
	md.mat.MarshalSigned(&f.Data, int64(m.xxx_mat))
	md.clt.MarshalSigned(&f.Data, int64(m.xxx_clt))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp2) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp2) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp2
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp2: expects ID 1522 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp2: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp2: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp2: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_baro = int16(md.baro.UnmarshalSigned(f.Data))
	m.xxx_map1 = int16(md.map1.UnmarshalSigned(f.Data))
	m.xxx_mat = int16(md.mat.UnmarshalSigned(f.Data))
	m.xxx_clt = int16(md.clt.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp3Reader provides read access to a megasquirt_gp3 message.
type megasquirt_gp3Reader interface {
	// tps returns the physical value of the tps signal.
	tps() float64
	// batt returns the physical value of the batt signal.
	batt() float64
	// afr1_old returns the physical value of the afr1_old signal.
	afr1_old() float64
	// afr2_old returns the physical value of the afr2_old signal.
	afr2_old() float64
}

// megasquirt_gp3Writer provides write access to a megasquirt_gp3 message.
type megasquirt_gp3Writer interface {
	// CopyFrom copies all values from megasquirt_gp3Reader.
	CopyFrom(megasquirt_gp3Reader) *megasquirt_gp3
	// Settps sets the physical value of the tps signal.
	Settps(float64) *megasquirt_gp3
	// Setbatt sets the physical value of the batt signal.
	Setbatt(float64) *megasquirt_gp3
	// Setafr1_old sets the physical value of the afr1_old signal.
	Setafr1_old(float64) *megasquirt_gp3
	// Setafr2_old sets the physical value of the afr2_old signal.
	Setafr2_old(float64) *megasquirt_gp3
}

type megasquirt_gp3 struct {
	xxx_tps      int16
	xxx_batt     int16
	xxx_afr1_old int16
	xxx_afr2_old int16
}

func Newmegasquirt_gp3() *megasquirt_gp3 {
	m := &megasquirt_gp3{}
	m.Reset()
	return m
}

func (m *megasquirt_gp3) Reset() {
	m.xxx_tps = 0
	m.xxx_batt = 0
	m.xxx_afr1_old = 0
	m.xxx_afr2_old = 0
}

func (m *megasquirt_gp3) CopyFrom(o megasquirt_gp3Reader) *megasquirt_gp3 {
	m.Settps(o.tps())
	m.Setbatt(o.batt())
	m.Setafr1_old(o.afr1_old())
	m.Setafr2_old(o.afr2_old())
	return m
}

// Descriptor returns the megasquirt_gp3 descriptor.
func (m *megasquirt_gp3) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp3.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp3) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp3) tps() float64 {
	return Messages().megasquirt_gp3.tps.ToPhysical(float64(m.xxx_tps))
}

func (m *megasquirt_gp3) Settps(v float64) *megasquirt_gp3 {
	m.xxx_tps = int16(Messages().megasquirt_gp3.tps.FromPhysical(v))
	return m
}

func (m *megasquirt_gp3) batt() float64 {
	return Messages().megasquirt_gp3.batt.ToPhysical(float64(m.xxx_batt))
}

func (m *megasquirt_gp3) Setbatt(v float64) *megasquirt_gp3 {
	m.xxx_batt = int16(Messages().megasquirt_gp3.batt.FromPhysical(v))
	return m
}

func (m *megasquirt_gp3) afr1_old() float64 {
	return Messages().megasquirt_gp3.afr1_old.ToPhysical(float64(m.xxx_afr1_old))
}

func (m *megasquirt_gp3) Setafr1_old(v float64) *megasquirt_gp3 {
	m.xxx_afr1_old = int16(Messages().megasquirt_gp3.afr1_old.FromPhysical(v))
	return m
}

func (m *megasquirt_gp3) afr2_old() float64 {
	return Messages().megasquirt_gp3.afr2_old.ToPhysical(float64(m.xxx_afr2_old))
}

func (m *megasquirt_gp3) Setafr2_old(v float64) *megasquirt_gp3 {
	m.xxx_afr2_old = int16(Messages().megasquirt_gp3.afr2_old.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp3) Frame() can.Frame {
	md := Messages().megasquirt_gp3
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.tps.MarshalSigned(&f.Data, int64(m.xxx_tps))
	md.batt.MarshalSigned(&f.Data, int64(m.xxx_batt))
	md.afr1_old.MarshalSigned(&f.Data, int64(m.xxx_afr1_old))
	md.afr2_old.MarshalSigned(&f.Data, int64(m.xxx_afr2_old))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp3) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp3) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp3
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp3: expects ID 1523 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp3: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp3: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp3: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_tps = int16(md.tps.UnmarshalSigned(f.Data))
	m.xxx_batt = int16(md.batt.UnmarshalSigned(f.Data))
	m.xxx_afr1_old = int16(md.afr1_old.UnmarshalSigned(f.Data))
	m.xxx_afr2_old = int16(md.afr2_old.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp4Reader provides read access to a megasquirt_gp4 message.
type megasquirt_gp4Reader interface {
	// knock returns the physical value of the knock signal.
	knock() float64
	// egocor1 returns the physical value of the egocor1 signal.
	egocor1() float64
	// egocor2 returns the physical value of the egocor2 signal.
	egocor2() float64
	// aircor returns the physical value of the aircor signal.
	aircor() float64
}

// megasquirt_gp4Writer provides write access to a megasquirt_gp4 message.
type megasquirt_gp4Writer interface {
	// CopyFrom copies all values from megasquirt_gp4Reader.
	CopyFrom(megasquirt_gp4Reader) *megasquirt_gp4
	// Setknock sets the physical value of the knock signal.
	Setknock(float64) *megasquirt_gp4
	// Setegocor1 sets the physical value of the egocor1 signal.
	Setegocor1(float64) *megasquirt_gp4
	// Setegocor2 sets the physical value of the egocor2 signal.
	Setegocor2(float64) *megasquirt_gp4
	// Setaircor sets the physical value of the aircor signal.
	Setaircor(float64) *megasquirt_gp4
}

type megasquirt_gp4 struct {
	xxx_knock   int16
	xxx_egocor1 int16
	xxx_egocor2 int16
	xxx_aircor  int16
}

func Newmegasquirt_gp4() *megasquirt_gp4 {
	m := &megasquirt_gp4{}
	m.Reset()
	return m
}

func (m *megasquirt_gp4) Reset() {
	m.xxx_knock = 0
	m.xxx_egocor1 = 0
	m.xxx_egocor2 = 0
	m.xxx_aircor = 0
}

func (m *megasquirt_gp4) CopyFrom(o megasquirt_gp4Reader) *megasquirt_gp4 {
	m.Setknock(o.knock())
	m.Setegocor1(o.egocor1())
	m.Setegocor2(o.egocor2())
	m.Setaircor(o.aircor())
	return m
}

// Descriptor returns the megasquirt_gp4 descriptor.
func (m *megasquirt_gp4) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp4.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp4) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp4) knock() float64 {
	return Messages().megasquirt_gp4.knock.ToPhysical(float64(m.xxx_knock))
}

func (m *megasquirt_gp4) Setknock(v float64) *megasquirt_gp4 {
	m.xxx_knock = int16(Messages().megasquirt_gp4.knock.FromPhysical(v))
	return m
}

func (m *megasquirt_gp4) egocor1() float64 {
	return Messages().megasquirt_gp4.egocor1.ToPhysical(float64(m.xxx_egocor1))
}

func (m *megasquirt_gp4) Setegocor1(v float64) *megasquirt_gp4 {
	m.xxx_egocor1 = int16(Messages().megasquirt_gp4.egocor1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp4) egocor2() float64 {
	return Messages().megasquirt_gp4.egocor2.ToPhysical(float64(m.xxx_egocor2))
}

func (m *megasquirt_gp4) Setegocor2(v float64) *megasquirt_gp4 {
	m.xxx_egocor2 = int16(Messages().megasquirt_gp4.egocor2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp4) aircor() float64 {
	return Messages().megasquirt_gp4.aircor.ToPhysical(float64(m.xxx_aircor))
}

func (m *megasquirt_gp4) Setaircor(v float64) *megasquirt_gp4 {
	m.xxx_aircor = int16(Messages().megasquirt_gp4.aircor.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp4) Frame() can.Frame {
	md := Messages().megasquirt_gp4
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.knock.MarshalSigned(&f.Data, int64(m.xxx_knock))
	md.egocor1.MarshalSigned(&f.Data, int64(m.xxx_egocor1))
	md.egocor2.MarshalSigned(&f.Data, int64(m.xxx_egocor2))
	md.aircor.MarshalSigned(&f.Data, int64(m.xxx_aircor))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp4) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp4) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp4
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp4: expects ID 1524 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp4: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp4: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp4: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_knock = int16(md.knock.UnmarshalSigned(f.Data))
	m.xxx_egocor1 = int16(md.egocor1.UnmarshalSigned(f.Data))
	m.xxx_egocor2 = int16(md.egocor2.UnmarshalSigned(f.Data))
	m.xxx_aircor = int16(md.aircor.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp5Reader provides read access to a megasquirt_gp5 message.
type megasquirt_gp5Reader interface {
	// warmcor returns the physical value of the warmcor signal.
	warmcor() float64
	// tpsaccel returns the physical value of the tpsaccel signal.
	tpsaccel() float64
	// tpsfuelcut returns the physical value of the tpsfuelcut signal.
	tpsfuelcut() float64
	// barocor returns the physical value of the barocor signal.
	barocor() float64
}

// megasquirt_gp5Writer provides write access to a megasquirt_gp5 message.
type megasquirt_gp5Writer interface {
	// CopyFrom copies all values from megasquirt_gp5Reader.
	CopyFrom(megasquirt_gp5Reader) *megasquirt_gp5
	// Setwarmcor sets the physical value of the warmcor signal.
	Setwarmcor(float64) *megasquirt_gp5
	// Settpsaccel sets the physical value of the tpsaccel signal.
	Settpsaccel(float64) *megasquirt_gp5
	// Settpsfuelcut sets the physical value of the tpsfuelcut signal.
	Settpsfuelcut(float64) *megasquirt_gp5
	// Setbarocor sets the physical value of the barocor signal.
	Setbarocor(float64) *megasquirt_gp5
}

type megasquirt_gp5 struct {
	xxx_warmcor    int16
	xxx_tpsaccel   int16
	xxx_tpsfuelcut int16
	xxx_barocor    int16
}

func Newmegasquirt_gp5() *megasquirt_gp5 {
	m := &megasquirt_gp5{}
	m.Reset()
	return m
}

func (m *megasquirt_gp5) Reset() {
	m.xxx_warmcor = 0
	m.xxx_tpsaccel = 0
	m.xxx_tpsfuelcut = 0
	m.xxx_barocor = 0
}

func (m *megasquirt_gp5) CopyFrom(o megasquirt_gp5Reader) *megasquirt_gp5 {
	m.Setwarmcor(o.warmcor())
	m.Settpsaccel(o.tpsaccel())
	m.Settpsfuelcut(o.tpsfuelcut())
	m.Setbarocor(o.barocor())
	return m
}

// Descriptor returns the megasquirt_gp5 descriptor.
func (m *megasquirt_gp5) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp5.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp5) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp5) warmcor() float64 {
	return Messages().megasquirt_gp5.warmcor.ToPhysical(float64(m.xxx_warmcor))
}

func (m *megasquirt_gp5) Setwarmcor(v float64) *megasquirt_gp5 {
	m.xxx_warmcor = int16(Messages().megasquirt_gp5.warmcor.FromPhysical(v))
	return m
}

func (m *megasquirt_gp5) tpsaccel() float64 {
	return Messages().megasquirt_gp5.tpsaccel.ToPhysical(float64(m.xxx_tpsaccel))
}

func (m *megasquirt_gp5) Settpsaccel(v float64) *megasquirt_gp5 {
	m.xxx_tpsaccel = int16(Messages().megasquirt_gp5.tpsaccel.FromPhysical(v))
	return m
}

func (m *megasquirt_gp5) tpsfuelcut() float64 {
	return Messages().megasquirt_gp5.tpsfuelcut.ToPhysical(float64(m.xxx_tpsfuelcut))
}

func (m *megasquirt_gp5) Settpsfuelcut(v float64) *megasquirt_gp5 {
	m.xxx_tpsfuelcut = int16(Messages().megasquirt_gp5.tpsfuelcut.FromPhysical(v))
	return m
}

func (m *megasquirt_gp5) barocor() float64 {
	return Messages().megasquirt_gp5.barocor.ToPhysical(float64(m.xxx_barocor))
}

func (m *megasquirt_gp5) Setbarocor(v float64) *megasquirt_gp5 {
	m.xxx_barocor = int16(Messages().megasquirt_gp5.barocor.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp5) Frame() can.Frame {
	md := Messages().megasquirt_gp5
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.warmcor.MarshalSigned(&f.Data, int64(m.xxx_warmcor))
	md.tpsaccel.MarshalSigned(&f.Data, int64(m.xxx_tpsaccel))
	md.tpsfuelcut.MarshalSigned(&f.Data, int64(m.xxx_tpsfuelcut))
	md.barocor.MarshalSigned(&f.Data, int64(m.xxx_barocor))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp5) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp5) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp5
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp5: expects ID 1525 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp5: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp5: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp5: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_warmcor = int16(md.warmcor.UnmarshalSigned(f.Data))
	m.xxx_tpsaccel = int16(md.tpsaccel.UnmarshalSigned(f.Data))
	m.xxx_tpsfuelcut = int16(md.tpsfuelcut.UnmarshalSigned(f.Data))
	m.xxx_barocor = int16(md.barocor.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp6Reader provides read access to a megasquirt_gp6 message.
type megasquirt_gp6Reader interface {
	// totalcor returns the physical value of the totalcor signal.
	totalcor() float64
	// ve1 returns the physical value of the ve1 signal.
	ve1() float64
	// ve2 returns the physical value of the ve2 signal.
	ve2() float64
	// iacstep returns the value of the iacstep signal.
	iacstep() int16
}

// megasquirt_gp6Writer provides write access to a megasquirt_gp6 message.
type megasquirt_gp6Writer interface {
	// CopyFrom copies all values from megasquirt_gp6Reader.
	CopyFrom(megasquirt_gp6Reader) *megasquirt_gp6
	// Settotalcor sets the physical value of the totalcor signal.
	Settotalcor(float64) *megasquirt_gp6
	// Setve1 sets the physical value of the ve1 signal.
	Setve1(float64) *megasquirt_gp6
	// Setve2 sets the physical value of the ve2 signal.
	Setve2(float64) *megasquirt_gp6
	// Setiacstep sets the value of the iacstep signal.
	Setiacstep(int16) *megasquirt_gp6
}

type megasquirt_gp6 struct {
	xxx_totalcor int16
	xxx_ve1      int16
	xxx_ve2      int16
	xxx_iacstep  int16
}

func Newmegasquirt_gp6() *megasquirt_gp6 {
	m := &megasquirt_gp6{}
	m.Reset()
	return m
}

func (m *megasquirt_gp6) Reset() {
	m.xxx_totalcor = 0
	m.xxx_ve1 = 0
	m.xxx_ve2 = 0
	m.xxx_iacstep = 0
}

func (m *megasquirt_gp6) CopyFrom(o megasquirt_gp6Reader) *megasquirt_gp6 {
	m.Settotalcor(o.totalcor())
	m.Setve1(o.ve1())
	m.Setve2(o.ve2())
	m.xxx_iacstep = o.iacstep()
	return m
}

// Descriptor returns the megasquirt_gp6 descriptor.
func (m *megasquirt_gp6) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp6.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp6) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp6) totalcor() float64 {
	return Messages().megasquirt_gp6.totalcor.ToPhysical(float64(m.xxx_totalcor))
}

func (m *megasquirt_gp6) Settotalcor(v float64) *megasquirt_gp6 {
	m.xxx_totalcor = int16(Messages().megasquirt_gp6.totalcor.FromPhysical(v))
	return m
}

func (m *megasquirt_gp6) ve1() float64 {
	return Messages().megasquirt_gp6.ve1.ToPhysical(float64(m.xxx_ve1))
}

func (m *megasquirt_gp6) Setve1(v float64) *megasquirt_gp6 {
	m.xxx_ve1 = int16(Messages().megasquirt_gp6.ve1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp6) ve2() float64 {
	return Messages().megasquirt_gp6.ve2.ToPhysical(float64(m.xxx_ve2))
}

func (m *megasquirt_gp6) Setve2(v float64) *megasquirt_gp6 {
	m.xxx_ve2 = int16(Messages().megasquirt_gp6.ve2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp6) iacstep() int16 {
	return m.xxx_iacstep
}

func (m *megasquirt_gp6) Setiacstep(v int16) *megasquirt_gp6 {
	m.xxx_iacstep = int16(Messages().megasquirt_gp6.iacstep.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp6) Frame() can.Frame {
	md := Messages().megasquirt_gp6
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.totalcor.MarshalSigned(&f.Data, int64(m.xxx_totalcor))
	md.ve1.MarshalSigned(&f.Data, int64(m.xxx_ve1))
	md.ve2.MarshalSigned(&f.Data, int64(m.xxx_ve2))
	md.iacstep.MarshalSigned(&f.Data, int64(m.xxx_iacstep))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp6) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp6) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp6
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp6: expects ID 1526 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp6: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp6: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp6: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_totalcor = int16(md.totalcor.UnmarshalSigned(f.Data))
	m.xxx_ve1 = int16(md.ve1.UnmarshalSigned(f.Data))
	m.xxx_ve2 = int16(md.ve2.UnmarshalSigned(f.Data))
	m.xxx_iacstep = int16(md.iacstep.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp7Reader provides read access to a megasquirt_gp7 message.
type megasquirt_gp7Reader interface {
	// cold_adv_deg returns the physical value of the cold_adv_deg signal.
	cold_adv_deg() float64
	// TPSdot returns the physical value of the TPSdot signal.
	TPSdot() float64
	// MAPdot returns the value of the MAPdot signal.
	MAPdot() int16
	// RPMdot returns the physical value of the RPMdot signal.
	RPMdot() float64
}

// megasquirt_gp7Writer provides write access to a megasquirt_gp7 message.
type megasquirt_gp7Writer interface {
	// CopyFrom copies all values from megasquirt_gp7Reader.
	CopyFrom(megasquirt_gp7Reader) *megasquirt_gp7
	// Setcold_adv_deg sets the physical value of the cold_adv_deg signal.
	Setcold_adv_deg(float64) *megasquirt_gp7
	// SetTPSdot sets the physical value of the TPSdot signal.
	SetTPSdot(float64) *megasquirt_gp7
	// SetMAPdot sets the value of the MAPdot signal.
	SetMAPdot(int16) *megasquirt_gp7
	// SetRPMdot sets the physical value of the RPMdot signal.
	SetRPMdot(float64) *megasquirt_gp7
}

type megasquirt_gp7 struct {
	xxx_cold_adv_deg int16
	xxx_TPSdot       int16
	xxx_MAPdot       int16
	xxx_RPMdot       int16
}

func Newmegasquirt_gp7() *megasquirt_gp7 {
	m := &megasquirt_gp7{}
	m.Reset()
	return m
}

func (m *megasquirt_gp7) Reset() {
	m.xxx_cold_adv_deg = 0
	m.xxx_TPSdot = 0
	m.xxx_MAPdot = 0
	m.xxx_RPMdot = 0
}

func (m *megasquirt_gp7) CopyFrom(o megasquirt_gp7Reader) *megasquirt_gp7 {
	m.Setcold_adv_deg(o.cold_adv_deg())
	m.SetTPSdot(o.TPSdot())
	m.xxx_MAPdot = o.MAPdot()
	m.SetRPMdot(o.RPMdot())
	return m
}

// Descriptor returns the megasquirt_gp7 descriptor.
func (m *megasquirt_gp7) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp7.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp7) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp7) cold_adv_deg() float64 {
	return Messages().megasquirt_gp7.cold_adv_deg.ToPhysical(float64(m.xxx_cold_adv_deg))
}

func (m *megasquirt_gp7) Setcold_adv_deg(v float64) *megasquirt_gp7 {
	m.xxx_cold_adv_deg = int16(Messages().megasquirt_gp7.cold_adv_deg.FromPhysical(v))
	return m
}

func (m *megasquirt_gp7) TPSdot() float64 {
	return Messages().megasquirt_gp7.TPSdot.ToPhysical(float64(m.xxx_TPSdot))
}

func (m *megasquirt_gp7) SetTPSdot(v float64) *megasquirt_gp7 {
	m.xxx_TPSdot = int16(Messages().megasquirt_gp7.TPSdot.FromPhysical(v))
	return m
}

func (m *megasquirt_gp7) MAPdot() int16 {
	return m.xxx_MAPdot
}

func (m *megasquirt_gp7) SetMAPdot(v int16) *megasquirt_gp7 {
	m.xxx_MAPdot = int16(Messages().megasquirt_gp7.MAPdot.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp7) RPMdot() float64 {
	return Messages().megasquirt_gp7.RPMdot.ToPhysical(float64(m.xxx_RPMdot))
}

func (m *megasquirt_gp7) SetRPMdot(v float64) *megasquirt_gp7 {
	m.xxx_RPMdot = int16(Messages().megasquirt_gp7.RPMdot.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp7) Frame() can.Frame {
	md := Messages().megasquirt_gp7
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.cold_adv_deg.MarshalSigned(&f.Data, int64(m.xxx_cold_adv_deg))
	md.TPSdot.MarshalSigned(&f.Data, int64(m.xxx_TPSdot))
	md.MAPdot.MarshalSigned(&f.Data, int64(m.xxx_MAPdot))
	md.RPMdot.MarshalSigned(&f.Data, int64(m.xxx_RPMdot))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp7) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp7) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp7
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp7: expects ID 1527 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp7: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp7: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp7: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_cold_adv_deg = int16(md.cold_adv_deg.UnmarshalSigned(f.Data))
	m.xxx_TPSdot = int16(md.TPSdot.UnmarshalSigned(f.Data))
	m.xxx_MAPdot = int16(md.MAPdot.UnmarshalSigned(f.Data))
	m.xxx_RPMdot = int16(md.RPMdot.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp8Reader provides read access to a megasquirt_gp8 message.
type megasquirt_gp8Reader interface {
	// MAFload returns the physical value of the MAFload signal.
	MAFload() float64
	// fuelload returns the physical value of the fuelload signal.
	fuelload() float64
	// fuelcor returns the physical value of the fuelcor signal.
	fuelcor() float64
	// MAF returns the physical value of the MAF signal.
	MAF() float64
}

// megasquirt_gp8Writer provides write access to a megasquirt_gp8 message.
type megasquirt_gp8Writer interface {
	// CopyFrom copies all values from megasquirt_gp8Reader.
	CopyFrom(megasquirt_gp8Reader) *megasquirt_gp8
	// SetMAFload sets the physical value of the MAFload signal.
	SetMAFload(float64) *megasquirt_gp8
	// Setfuelload sets the physical value of the fuelload signal.
	Setfuelload(float64) *megasquirt_gp8
	// Setfuelcor sets the physical value of the fuelcor signal.
	Setfuelcor(float64) *megasquirt_gp8
	// SetMAF sets the physical value of the MAF signal.
	SetMAF(float64) *megasquirt_gp8
}

type megasquirt_gp8 struct {
	xxx_MAFload  int16
	xxx_fuelload int16
	xxx_fuelcor  int16
	xxx_MAF      int16
}

func Newmegasquirt_gp8() *megasquirt_gp8 {
	m := &megasquirt_gp8{}
	m.Reset()
	return m
}

func (m *megasquirt_gp8) Reset() {
	m.xxx_MAFload = 0
	m.xxx_fuelload = 0
	m.xxx_fuelcor = 0
	m.xxx_MAF = 0
}

func (m *megasquirt_gp8) CopyFrom(o megasquirt_gp8Reader) *megasquirt_gp8 {
	m.SetMAFload(o.MAFload())
	m.Setfuelload(o.fuelload())
	m.Setfuelcor(o.fuelcor())
	m.SetMAF(o.MAF())
	return m
}

// Descriptor returns the megasquirt_gp8 descriptor.
func (m *megasquirt_gp8) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp8.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp8) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp8) MAFload() float64 {
	return Messages().megasquirt_gp8.MAFload.ToPhysical(float64(m.xxx_MAFload))
}

func (m *megasquirt_gp8) SetMAFload(v float64) *megasquirt_gp8 {
	m.xxx_MAFload = int16(Messages().megasquirt_gp8.MAFload.FromPhysical(v))
	return m
}

func (m *megasquirt_gp8) fuelload() float64 {
	return Messages().megasquirt_gp8.fuelload.ToPhysical(float64(m.xxx_fuelload))
}

func (m *megasquirt_gp8) Setfuelload(v float64) *megasquirt_gp8 {
	m.xxx_fuelload = int16(Messages().megasquirt_gp8.fuelload.FromPhysical(v))
	return m
}

func (m *megasquirt_gp8) fuelcor() float64 {
	return Messages().megasquirt_gp8.fuelcor.ToPhysical(float64(m.xxx_fuelcor))
}

func (m *megasquirt_gp8) Setfuelcor(v float64) *megasquirt_gp8 {
	m.xxx_fuelcor = int16(Messages().megasquirt_gp8.fuelcor.FromPhysical(v))
	return m
}

func (m *megasquirt_gp8) MAF() float64 {
	return Messages().megasquirt_gp8.MAF.ToPhysical(float64(m.xxx_MAF))
}

func (m *megasquirt_gp8) SetMAF(v float64) *megasquirt_gp8 {
	m.xxx_MAF = int16(Messages().megasquirt_gp8.MAF.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp8) Frame() can.Frame {
	md := Messages().megasquirt_gp8
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.MAFload.MarshalSigned(&f.Data, int64(m.xxx_MAFload))
	md.fuelload.MarshalSigned(&f.Data, int64(m.xxx_fuelload))
	md.fuelcor.MarshalSigned(&f.Data, int64(m.xxx_fuelcor))
	md.MAF.MarshalSigned(&f.Data, int64(m.xxx_MAF))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp8) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp8) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp8
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp8: expects ID 1528 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp8: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp8: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp8: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_MAFload = int16(md.MAFload.UnmarshalSigned(f.Data))
	m.xxx_fuelload = int16(md.fuelload.UnmarshalSigned(f.Data))
	m.xxx_fuelcor = int16(md.fuelcor.UnmarshalSigned(f.Data))
	m.xxx_MAF = int16(md.MAF.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp9Reader provides read access to a megasquirt_gp9 message.
type megasquirt_gp9Reader interface {
	// egoV1 returns the physical value of the egoV1 signal.
	egoV1() float64
	// egoV2 returns the physical value of the egoV2 signal.
	egoV2() float64
	// dwell returns the physical value of the dwell signal.
	dwell() float64
	// dwell_trl returns the physical value of the dwell_trl signal.
	dwell_trl() float64
}

// megasquirt_gp9Writer provides write access to a megasquirt_gp9 message.
type megasquirt_gp9Writer interface {
	// CopyFrom copies all values from megasquirt_gp9Reader.
	CopyFrom(megasquirt_gp9Reader) *megasquirt_gp9
	// SetegoV1 sets the physical value of the egoV1 signal.
	SetegoV1(float64) *megasquirt_gp9
	// SetegoV2 sets the physical value of the egoV2 signal.
	SetegoV2(float64) *megasquirt_gp9
	// Setdwell sets the physical value of the dwell signal.
	Setdwell(float64) *megasquirt_gp9
	// Setdwell_trl sets the physical value of the dwell_trl signal.
	Setdwell_trl(float64) *megasquirt_gp9
}

type megasquirt_gp9 struct {
	xxx_egoV1     int16
	xxx_egoV2     int16
	xxx_dwell     uint16
	xxx_dwell_trl uint16
}

func Newmegasquirt_gp9() *megasquirt_gp9 {
	m := &megasquirt_gp9{}
	m.Reset()
	return m
}

func (m *megasquirt_gp9) Reset() {
	m.xxx_egoV1 = 0
	m.xxx_egoV2 = 0
	m.xxx_dwell = 0
	m.xxx_dwell_trl = 0
}

func (m *megasquirt_gp9) CopyFrom(o megasquirt_gp9Reader) *megasquirt_gp9 {
	m.SetegoV1(o.egoV1())
	m.SetegoV2(o.egoV2())
	m.Setdwell(o.dwell())
	m.Setdwell_trl(o.dwell_trl())
	return m
}

// Descriptor returns the megasquirt_gp9 descriptor.
func (m *megasquirt_gp9) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp9.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp9) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp9) egoV1() float64 {
	return Messages().megasquirt_gp9.egoV1.ToPhysical(float64(m.xxx_egoV1))
}

func (m *megasquirt_gp9) SetegoV1(v float64) *megasquirt_gp9 {
	m.xxx_egoV1 = int16(Messages().megasquirt_gp9.egoV1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp9) egoV2() float64 {
	return Messages().megasquirt_gp9.egoV2.ToPhysical(float64(m.xxx_egoV2))
}

func (m *megasquirt_gp9) SetegoV2(v float64) *megasquirt_gp9 {
	m.xxx_egoV2 = int16(Messages().megasquirt_gp9.egoV2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp9) dwell() float64 {
	return Messages().megasquirt_gp9.dwell.ToPhysical(float64(m.xxx_dwell))
}

func (m *megasquirt_gp9) Setdwell(v float64) *megasquirt_gp9 {
	m.xxx_dwell = uint16(Messages().megasquirt_gp9.dwell.FromPhysical(v))
	return m
}

func (m *megasquirt_gp9) dwell_trl() float64 {
	return Messages().megasquirt_gp9.dwell_trl.ToPhysical(float64(m.xxx_dwell_trl))
}

func (m *megasquirt_gp9) Setdwell_trl(v float64) *megasquirt_gp9 {
	m.xxx_dwell_trl = uint16(Messages().megasquirt_gp9.dwell_trl.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp9) Frame() can.Frame {
	md := Messages().megasquirt_gp9
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.egoV1.MarshalSigned(&f.Data, int64(m.xxx_egoV1))
	md.egoV2.MarshalSigned(&f.Data, int64(m.xxx_egoV2))
	md.dwell.MarshalUnsigned(&f.Data, uint64(m.xxx_dwell))
	md.dwell_trl.MarshalUnsigned(&f.Data, uint64(m.xxx_dwell_trl))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp9) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp9) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp9
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp9: expects ID 1529 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp9: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp9: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp9: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_egoV1 = int16(md.egoV1.UnmarshalSigned(f.Data))
	m.xxx_egoV2 = int16(md.egoV2.UnmarshalSigned(f.Data))
	m.xxx_dwell = uint16(md.dwell.UnmarshalUnsigned(f.Data))
	m.xxx_dwell_trl = uint16(md.dwell_trl.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp10Reader provides read access to a megasquirt_gp10 message.
type megasquirt_gp10Reader interface {
	// status1 returns the value of the status1 signal.
	status1() uint8
	// status2 returns the value of the status2 signal.
	status2() uint8
	// status3 returns the value of the status3 signal.
	status3() uint8
	// status4 returns the value of the status4 signal.
	status4() uint8
	// status5 returns the value of the status5 signal.
	status5() int16
	// status6 returns the value of the status6 signal.
	status6() uint8
	// status7 returns the value of the status7 signal.
	status7() uint8
}

// megasquirt_gp10Writer provides write access to a megasquirt_gp10 message.
type megasquirt_gp10Writer interface {
	// CopyFrom copies all values from megasquirt_gp10Reader.
	CopyFrom(megasquirt_gp10Reader) *megasquirt_gp10
	// Setstatus1 sets the value of the status1 signal.
	Setstatus1(uint8) *megasquirt_gp10
	// Setstatus2 sets the value of the status2 signal.
	Setstatus2(uint8) *megasquirt_gp10
	// Setstatus3 sets the value of the status3 signal.
	Setstatus3(uint8) *megasquirt_gp10
	// Setstatus4 sets the value of the status4 signal.
	Setstatus4(uint8) *megasquirt_gp10
	// Setstatus5 sets the value of the status5 signal.
	Setstatus5(int16) *megasquirt_gp10
	// Setstatus6 sets the value of the status6 signal.
	Setstatus6(uint8) *megasquirt_gp10
	// Setstatus7 sets the value of the status7 signal.
	Setstatus7(uint8) *megasquirt_gp10
}

type megasquirt_gp10 struct {
	xxx_status1 uint8
	xxx_status2 uint8
	xxx_status3 uint8
	xxx_status4 uint8
	xxx_status5 int16
	xxx_status6 uint8
	xxx_status7 uint8
}

func Newmegasquirt_gp10() *megasquirt_gp10 {
	m := &megasquirt_gp10{}
	m.Reset()
	return m
}

func (m *megasquirt_gp10) Reset() {
	m.xxx_status1 = 0
	m.xxx_status2 = 0
	m.xxx_status3 = 0
	m.xxx_status4 = 0
	m.xxx_status5 = 0
	m.xxx_status6 = 0
	m.xxx_status7 = 0
}

func (m *megasquirt_gp10) CopyFrom(o megasquirt_gp10Reader) *megasquirt_gp10 {
	m.xxx_status1 = o.status1()
	m.xxx_status2 = o.status2()
	m.xxx_status3 = o.status3()
	m.xxx_status4 = o.status4()
	m.xxx_status5 = o.status5()
	m.xxx_status6 = o.status6()
	m.xxx_status7 = o.status7()
	return m
}

// Descriptor returns the megasquirt_gp10 descriptor.
func (m *megasquirt_gp10) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp10.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp10) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp10) status1() uint8 {
	return m.xxx_status1
}

func (m *megasquirt_gp10) Setstatus1(v uint8) *megasquirt_gp10 {
	m.xxx_status1 = uint8(Messages().megasquirt_gp10.status1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp10) status2() uint8 {
	return m.xxx_status2
}

func (m *megasquirt_gp10) Setstatus2(v uint8) *megasquirt_gp10 {
	m.xxx_status2 = uint8(Messages().megasquirt_gp10.status2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp10) status3() uint8 {
	return m.xxx_status3
}

func (m *megasquirt_gp10) Setstatus3(v uint8) *megasquirt_gp10 {
	m.xxx_status3 = uint8(Messages().megasquirt_gp10.status3.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp10) status4() uint8 {
	return m.xxx_status4
}

func (m *megasquirt_gp10) Setstatus4(v uint8) *megasquirt_gp10 {
	m.xxx_status4 = uint8(Messages().megasquirt_gp10.status4.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp10) status5() int16 {
	return m.xxx_status5
}

func (m *megasquirt_gp10) Setstatus5(v int16) *megasquirt_gp10 {
	m.xxx_status5 = int16(Messages().megasquirt_gp10.status5.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp10) status6() uint8 {
	return m.xxx_status6
}

func (m *megasquirt_gp10) Setstatus6(v uint8) *megasquirt_gp10 {
	m.xxx_status6 = uint8(Messages().megasquirt_gp10.status6.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp10) status7() uint8 {
	return m.xxx_status7
}

func (m *megasquirt_gp10) Setstatus7(v uint8) *megasquirt_gp10 {
	m.xxx_status7 = uint8(Messages().megasquirt_gp10.status7.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp10) Frame() can.Frame {
	md := Messages().megasquirt_gp10
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.status1.MarshalUnsigned(&f.Data, uint64(m.xxx_status1))
	md.status2.MarshalUnsigned(&f.Data, uint64(m.xxx_status2))
	md.status3.MarshalUnsigned(&f.Data, uint64(m.xxx_status3))
	md.status4.MarshalUnsigned(&f.Data, uint64(m.xxx_status4))
	md.status5.MarshalSigned(&f.Data, int64(m.xxx_status5))
	md.status6.MarshalUnsigned(&f.Data, uint64(m.xxx_status6))
	md.status7.MarshalUnsigned(&f.Data, uint64(m.xxx_status7))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp10) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp10) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp10
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp10: expects ID 1530 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp10: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp10: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp10: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_status1 = uint8(md.status1.UnmarshalUnsigned(f.Data))
	m.xxx_status2 = uint8(md.status2.UnmarshalUnsigned(f.Data))
	m.xxx_status3 = uint8(md.status3.UnmarshalUnsigned(f.Data))
	m.xxx_status4 = uint8(md.status4.UnmarshalUnsigned(f.Data))
	m.xxx_status5 = int16(md.status5.UnmarshalSigned(f.Data))
	m.xxx_status6 = uint8(md.status6.UnmarshalUnsigned(f.Data))
	m.xxx_status7 = uint8(md.status7.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp11Reader provides read access to a megasquirt_gp11 message.
type megasquirt_gp11Reader interface {
	// fuelload2 returns the physical value of the fuelload2 signal.
	fuelload2() float64
	// ignload returns the physical value of the ignload signal.
	ignload() float64
	// ignload2 returns the physical value of the ignload2 signal.
	ignload2() float64
	// airtemp returns the physical value of the airtemp signal.
	airtemp() float64
}

// megasquirt_gp11Writer provides write access to a megasquirt_gp11 message.
type megasquirt_gp11Writer interface {
	// CopyFrom copies all values from megasquirt_gp11Reader.
	CopyFrom(megasquirt_gp11Reader) *megasquirt_gp11
	// Setfuelload2 sets the physical value of the fuelload2 signal.
	Setfuelload2(float64) *megasquirt_gp11
	// Setignload sets the physical value of the ignload signal.
	Setignload(float64) *megasquirt_gp11
	// Setignload2 sets the physical value of the ignload2 signal.
	Setignload2(float64) *megasquirt_gp11
	// Setairtemp sets the physical value of the airtemp signal.
	Setairtemp(float64) *megasquirt_gp11
}

type megasquirt_gp11 struct {
	xxx_fuelload2 int16
	xxx_ignload   int16
	xxx_ignload2  int16
	xxx_airtemp   int16
}

func Newmegasquirt_gp11() *megasquirt_gp11 {
	m := &megasquirt_gp11{}
	m.Reset()
	return m
}

func (m *megasquirt_gp11) Reset() {
	m.xxx_fuelload2 = 0
	m.xxx_ignload = 0
	m.xxx_ignload2 = 0
	m.xxx_airtemp = 0
}

func (m *megasquirt_gp11) CopyFrom(o megasquirt_gp11Reader) *megasquirt_gp11 {
	m.Setfuelload2(o.fuelload2())
	m.Setignload(o.ignload())
	m.Setignload2(o.ignload2())
	m.Setairtemp(o.airtemp())
	return m
}

// Descriptor returns the megasquirt_gp11 descriptor.
func (m *megasquirt_gp11) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp11.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp11) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp11) fuelload2() float64 {
	return Messages().megasquirt_gp11.fuelload2.ToPhysical(float64(m.xxx_fuelload2))
}

func (m *megasquirt_gp11) Setfuelload2(v float64) *megasquirt_gp11 {
	m.xxx_fuelload2 = int16(Messages().megasquirt_gp11.fuelload2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp11) ignload() float64 {
	return Messages().megasquirt_gp11.ignload.ToPhysical(float64(m.xxx_ignload))
}

func (m *megasquirt_gp11) Setignload(v float64) *megasquirt_gp11 {
	m.xxx_ignload = int16(Messages().megasquirt_gp11.ignload.FromPhysical(v))
	return m
}

func (m *megasquirt_gp11) ignload2() float64 {
	return Messages().megasquirt_gp11.ignload2.ToPhysical(float64(m.xxx_ignload2))
}

func (m *megasquirt_gp11) Setignload2(v float64) *megasquirt_gp11 {
	m.xxx_ignload2 = int16(Messages().megasquirt_gp11.ignload2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp11) airtemp() float64 {
	return Messages().megasquirt_gp11.airtemp.ToPhysical(float64(m.xxx_airtemp))
}

func (m *megasquirt_gp11) Setairtemp(v float64) *megasquirt_gp11 {
	m.xxx_airtemp = int16(Messages().megasquirt_gp11.airtemp.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp11) Frame() can.Frame {
	md := Messages().megasquirt_gp11
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.fuelload2.MarshalSigned(&f.Data, int64(m.xxx_fuelload2))
	md.ignload.MarshalSigned(&f.Data, int64(m.xxx_ignload))
	md.ignload2.MarshalSigned(&f.Data, int64(m.xxx_ignload2))
	md.airtemp.MarshalSigned(&f.Data, int64(m.xxx_airtemp))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp11) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp11) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp11
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp11: expects ID 1531 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp11: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp11: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp11: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_fuelload2 = int16(md.fuelload2.UnmarshalSigned(f.Data))
	m.xxx_ignload = int16(md.ignload.UnmarshalSigned(f.Data))
	m.xxx_ignload2 = int16(md.ignload2.UnmarshalSigned(f.Data))
	m.xxx_airtemp = int16(md.airtemp.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp12Reader provides read access to a megasquirt_gp12 message.
type megasquirt_gp12Reader interface {
	// wallfuel1 returns the physical value of the wallfuel1 signal.
	wallfuel1() float64
	// wallfuel2 returns the physical value of the wallfuel2 signal.
	wallfuel2() float64
}

// megasquirt_gp12Writer provides write access to a megasquirt_gp12 message.
type megasquirt_gp12Writer interface {
	// CopyFrom copies all values from megasquirt_gp12Reader.
	CopyFrom(megasquirt_gp12Reader) *megasquirt_gp12
	// Setwallfuel1 sets the physical value of the wallfuel1 signal.
	Setwallfuel1(float64) *megasquirt_gp12
	// Setwallfuel2 sets the physical value of the wallfuel2 signal.
	Setwallfuel2(float64) *megasquirt_gp12
}

type megasquirt_gp12 struct {
	xxx_wallfuel1 int32
	xxx_wallfuel2 int32
}

func Newmegasquirt_gp12() *megasquirt_gp12 {
	m := &megasquirt_gp12{}
	m.Reset()
	return m
}

func (m *megasquirt_gp12) Reset() {
	m.xxx_wallfuel1 = 0
	m.xxx_wallfuel2 = 0
}

func (m *megasquirt_gp12) CopyFrom(o megasquirt_gp12Reader) *megasquirt_gp12 {
	m.Setwallfuel1(o.wallfuel1())
	m.Setwallfuel2(o.wallfuel2())
	return m
}

// Descriptor returns the megasquirt_gp12 descriptor.
func (m *megasquirt_gp12) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp12.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp12) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp12) wallfuel1() float64 {
	return Messages().megasquirt_gp12.wallfuel1.ToPhysical(float64(m.xxx_wallfuel1))
}

func (m *megasquirt_gp12) Setwallfuel1(v float64) *megasquirt_gp12 {
	m.xxx_wallfuel1 = int32(Messages().megasquirt_gp12.wallfuel1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp12) wallfuel2() float64 {
	return Messages().megasquirt_gp12.wallfuel2.ToPhysical(float64(m.xxx_wallfuel2))
}

func (m *megasquirt_gp12) Setwallfuel2(v float64) *megasquirt_gp12 {
	m.xxx_wallfuel2 = int32(Messages().megasquirt_gp12.wallfuel2.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp12) Frame() can.Frame {
	md := Messages().megasquirt_gp12
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.wallfuel1.MarshalSigned(&f.Data, int64(m.xxx_wallfuel1))
	md.wallfuel2.MarshalSigned(&f.Data, int64(m.xxx_wallfuel2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp12) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp12) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp12
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp12: expects ID 1532 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp12: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp12: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp12: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_wallfuel1 = int32(md.wallfuel1.UnmarshalSigned(f.Data))
	m.xxx_wallfuel2 = int32(md.wallfuel2.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp13Reader provides read access to a megasquirt_gp13 message.
type megasquirt_gp13Reader interface {
	// sensors1 returns the physical value of the sensors1 signal.
	sensors1() float64
	// sensors2 returns the physical value of the sensors2 signal.
	sensors2() float64
	// sensors3 returns the physical value of the sensors3 signal.
	sensors3() float64
	// sensors4 returns the physical value of the sensors4 signal.
	sensors4() float64
}

// megasquirt_gp13Writer provides write access to a megasquirt_gp13 message.
type megasquirt_gp13Writer interface {
	// CopyFrom copies all values from megasquirt_gp13Reader.
	CopyFrom(megasquirt_gp13Reader) *megasquirt_gp13
	// Setsensors1 sets the physical value of the sensors1 signal.
	Setsensors1(float64) *megasquirt_gp13
	// Setsensors2 sets the physical value of the sensors2 signal.
	Setsensors2(float64) *megasquirt_gp13
	// Setsensors3 sets the physical value of the sensors3 signal.
	Setsensors3(float64) *megasquirt_gp13
	// Setsensors4 sets the physical value of the sensors4 signal.
	Setsensors4(float64) *megasquirt_gp13
}

type megasquirt_gp13 struct {
	xxx_sensors1 int16
	xxx_sensors2 int16
	xxx_sensors3 int16
	xxx_sensors4 int16
}

func Newmegasquirt_gp13() *megasquirt_gp13 {
	m := &megasquirt_gp13{}
	m.Reset()
	return m
}

func (m *megasquirt_gp13) Reset() {
	m.xxx_sensors1 = 0
	m.xxx_sensors2 = 0
	m.xxx_sensors3 = 0
	m.xxx_sensors4 = 0
}

func (m *megasquirt_gp13) CopyFrom(o megasquirt_gp13Reader) *megasquirt_gp13 {
	m.Setsensors1(o.sensors1())
	m.Setsensors2(o.sensors2())
	m.Setsensors3(o.sensors3())
	m.Setsensors4(o.sensors4())
	return m
}

// Descriptor returns the megasquirt_gp13 descriptor.
func (m *megasquirt_gp13) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp13.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp13) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp13) sensors1() float64 {
	return Messages().megasquirt_gp13.sensors1.ToPhysical(float64(m.xxx_sensors1))
}

func (m *megasquirt_gp13) Setsensors1(v float64) *megasquirt_gp13 {
	m.xxx_sensors1 = int16(Messages().megasquirt_gp13.sensors1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp13) sensors2() float64 {
	return Messages().megasquirt_gp13.sensors2.ToPhysical(float64(m.xxx_sensors2))
}

func (m *megasquirt_gp13) Setsensors2(v float64) *megasquirt_gp13 {
	m.xxx_sensors2 = int16(Messages().megasquirt_gp13.sensors2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp13) sensors3() float64 {
	return Messages().megasquirt_gp13.sensors3.ToPhysical(float64(m.xxx_sensors3))
}

func (m *megasquirt_gp13) Setsensors3(v float64) *megasquirt_gp13 {
	m.xxx_sensors3 = int16(Messages().megasquirt_gp13.sensors3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp13) sensors4() float64 {
	return Messages().megasquirt_gp13.sensors4.ToPhysical(float64(m.xxx_sensors4))
}

func (m *megasquirt_gp13) Setsensors4(v float64) *megasquirt_gp13 {
	m.xxx_sensors4 = int16(Messages().megasquirt_gp13.sensors4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp13) Frame() can.Frame {
	md := Messages().megasquirt_gp13
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.sensors1.MarshalSigned(&f.Data, int64(m.xxx_sensors1))
	md.sensors2.MarshalSigned(&f.Data, int64(m.xxx_sensors2))
	md.sensors3.MarshalSigned(&f.Data, int64(m.xxx_sensors3))
	md.sensors4.MarshalSigned(&f.Data, int64(m.xxx_sensors4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp13) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp13) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp13
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp13: expects ID 1533 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp13: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp13: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp13: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_sensors1 = int16(md.sensors1.UnmarshalSigned(f.Data))
	m.xxx_sensors2 = int16(md.sensors2.UnmarshalSigned(f.Data))
	m.xxx_sensors3 = int16(md.sensors3.UnmarshalSigned(f.Data))
	m.xxx_sensors4 = int16(md.sensors4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp14Reader provides read access to a megasquirt_gp14 message.
type megasquirt_gp14Reader interface {
	// sensors5 returns the physical value of the sensors5 signal.
	sensors5() float64
	// sensors6 returns the physical value of the sensors6 signal.
	sensors6() float64
	// sensors7 returns the physical value of the sensors7 signal.
	sensors7() float64
	// sensors8 returns the physical value of the sensors8 signal.
	sensors8() float64
}

// megasquirt_gp14Writer provides write access to a megasquirt_gp14 message.
type megasquirt_gp14Writer interface {
	// CopyFrom copies all values from megasquirt_gp14Reader.
	CopyFrom(megasquirt_gp14Reader) *megasquirt_gp14
	// Setsensors5 sets the physical value of the sensors5 signal.
	Setsensors5(float64) *megasquirt_gp14
	// Setsensors6 sets the physical value of the sensors6 signal.
	Setsensors6(float64) *megasquirt_gp14
	// Setsensors7 sets the physical value of the sensors7 signal.
	Setsensors7(float64) *megasquirt_gp14
	// Setsensors8 sets the physical value of the sensors8 signal.
	Setsensors8(float64) *megasquirt_gp14
}

type megasquirt_gp14 struct {
	xxx_sensors5 int16
	xxx_sensors6 int16
	xxx_sensors7 int16
	xxx_sensors8 int16
}

func Newmegasquirt_gp14() *megasquirt_gp14 {
	m := &megasquirt_gp14{}
	m.Reset()
	return m
}

func (m *megasquirt_gp14) Reset() {
	m.xxx_sensors5 = 0
	m.xxx_sensors6 = 0
	m.xxx_sensors7 = 0
	m.xxx_sensors8 = 0
}

func (m *megasquirt_gp14) CopyFrom(o megasquirt_gp14Reader) *megasquirt_gp14 {
	m.Setsensors5(o.sensors5())
	m.Setsensors6(o.sensors6())
	m.Setsensors7(o.sensors7())
	m.Setsensors8(o.sensors8())
	return m
}

// Descriptor returns the megasquirt_gp14 descriptor.
func (m *megasquirt_gp14) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp14.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp14) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp14) sensors5() float64 {
	return Messages().megasquirt_gp14.sensors5.ToPhysical(float64(m.xxx_sensors5))
}

func (m *megasquirt_gp14) Setsensors5(v float64) *megasquirt_gp14 {
	m.xxx_sensors5 = int16(Messages().megasquirt_gp14.sensors5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp14) sensors6() float64 {
	return Messages().megasquirt_gp14.sensors6.ToPhysical(float64(m.xxx_sensors6))
}

func (m *megasquirt_gp14) Setsensors6(v float64) *megasquirt_gp14 {
	m.xxx_sensors6 = int16(Messages().megasquirt_gp14.sensors6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp14) sensors7() float64 {
	return Messages().megasquirt_gp14.sensors7.ToPhysical(float64(m.xxx_sensors7))
}

func (m *megasquirt_gp14) Setsensors7(v float64) *megasquirt_gp14 {
	m.xxx_sensors7 = int16(Messages().megasquirt_gp14.sensors7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp14) sensors8() float64 {
	return Messages().megasquirt_gp14.sensors8.ToPhysical(float64(m.xxx_sensors8))
}

func (m *megasquirt_gp14) Setsensors8(v float64) *megasquirt_gp14 {
	m.xxx_sensors8 = int16(Messages().megasquirt_gp14.sensors8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp14) Frame() can.Frame {
	md := Messages().megasquirt_gp14
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.sensors5.MarshalSigned(&f.Data, int64(m.xxx_sensors5))
	md.sensors6.MarshalSigned(&f.Data, int64(m.xxx_sensors6))
	md.sensors7.MarshalSigned(&f.Data, int64(m.xxx_sensors7))
	md.sensors8.MarshalSigned(&f.Data, int64(m.xxx_sensors8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp14) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp14) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp14
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp14: expects ID 1534 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp14: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp14: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp14: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_sensors5 = int16(md.sensors5.UnmarshalSigned(f.Data))
	m.xxx_sensors6 = int16(md.sensors6.UnmarshalSigned(f.Data))
	m.xxx_sensors7 = int16(md.sensors7.UnmarshalSigned(f.Data))
	m.xxx_sensors8 = int16(md.sensors8.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp15Reader provides read access to a megasquirt_gp15 message.
type megasquirt_gp15Reader interface {
	// sensors9 returns the physical value of the sensors9 signal.
	sensors9() float64
	// sensors10 returns the physical value of the sensors10 signal.
	sensors10() float64
	// sensors11 returns the physical value of the sensors11 signal.
	sensors11() float64
	// sensors12 returns the physical value of the sensors12 signal.
	sensors12() float64
}

// megasquirt_gp15Writer provides write access to a megasquirt_gp15 message.
type megasquirt_gp15Writer interface {
	// CopyFrom copies all values from megasquirt_gp15Reader.
	CopyFrom(megasquirt_gp15Reader) *megasquirt_gp15
	// Setsensors9 sets the physical value of the sensors9 signal.
	Setsensors9(float64) *megasquirt_gp15
	// Setsensors10 sets the physical value of the sensors10 signal.
	Setsensors10(float64) *megasquirt_gp15
	// Setsensors11 sets the physical value of the sensors11 signal.
	Setsensors11(float64) *megasquirt_gp15
	// Setsensors12 sets the physical value of the sensors12 signal.
	Setsensors12(float64) *megasquirt_gp15
}

type megasquirt_gp15 struct {
	xxx_sensors9  int16
	xxx_sensors10 int16
	xxx_sensors11 int16
	xxx_sensors12 int16
}

func Newmegasquirt_gp15() *megasquirt_gp15 {
	m := &megasquirt_gp15{}
	m.Reset()
	return m
}

func (m *megasquirt_gp15) Reset() {
	m.xxx_sensors9 = 0
	m.xxx_sensors10 = 0
	m.xxx_sensors11 = 0
	m.xxx_sensors12 = 0
}

func (m *megasquirt_gp15) CopyFrom(o megasquirt_gp15Reader) *megasquirt_gp15 {
	m.Setsensors9(o.sensors9())
	m.Setsensors10(o.sensors10())
	m.Setsensors11(o.sensors11())
	m.Setsensors12(o.sensors12())
	return m
}

// Descriptor returns the megasquirt_gp15 descriptor.
func (m *megasquirt_gp15) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp15.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp15) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp15) sensors9() float64 {
	return Messages().megasquirt_gp15.sensors9.ToPhysical(float64(m.xxx_sensors9))
}

func (m *megasquirt_gp15) Setsensors9(v float64) *megasquirt_gp15 {
	m.xxx_sensors9 = int16(Messages().megasquirt_gp15.sensors9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp15) sensors10() float64 {
	return Messages().megasquirt_gp15.sensors10.ToPhysical(float64(m.xxx_sensors10))
}

func (m *megasquirt_gp15) Setsensors10(v float64) *megasquirt_gp15 {
	m.xxx_sensors10 = int16(Messages().megasquirt_gp15.sensors10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp15) sensors11() float64 {
	return Messages().megasquirt_gp15.sensors11.ToPhysical(float64(m.xxx_sensors11))
}

func (m *megasquirt_gp15) Setsensors11(v float64) *megasquirt_gp15 {
	m.xxx_sensors11 = int16(Messages().megasquirt_gp15.sensors11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp15) sensors12() float64 {
	return Messages().megasquirt_gp15.sensors12.ToPhysical(float64(m.xxx_sensors12))
}

func (m *megasquirt_gp15) Setsensors12(v float64) *megasquirt_gp15 {
	m.xxx_sensors12 = int16(Messages().megasquirt_gp15.sensors12.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp15) Frame() can.Frame {
	md := Messages().megasquirt_gp15
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.sensors9.MarshalSigned(&f.Data, int64(m.xxx_sensors9))
	md.sensors10.MarshalSigned(&f.Data, int64(m.xxx_sensors10))
	md.sensors11.MarshalSigned(&f.Data, int64(m.xxx_sensors11))
	md.sensors12.MarshalSigned(&f.Data, int64(m.xxx_sensors12))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp15) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp15) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp15
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp15: expects ID 1535 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp15: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp15: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp15: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_sensors9 = int16(md.sensors9.UnmarshalSigned(f.Data))
	m.xxx_sensors10 = int16(md.sensors10.UnmarshalSigned(f.Data))
	m.xxx_sensors11 = int16(md.sensors11.UnmarshalSigned(f.Data))
	m.xxx_sensors12 = int16(md.sensors12.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp16Reader provides read access to a megasquirt_gp16 message.
type megasquirt_gp16Reader interface {
	// sensors13 returns the physical value of the sensors13 signal.
	sensors13() float64
	// sensors14 returns the physical value of the sensors14 signal.
	sensors14() float64
	// sensors15 returns the physical value of the sensors15 signal.
	sensors15() float64
	// sensors16 returns the physical value of the sensors16 signal.
	sensors16() float64
}

// megasquirt_gp16Writer provides write access to a megasquirt_gp16 message.
type megasquirt_gp16Writer interface {
	// CopyFrom copies all values from megasquirt_gp16Reader.
	CopyFrom(megasquirt_gp16Reader) *megasquirt_gp16
	// Setsensors13 sets the physical value of the sensors13 signal.
	Setsensors13(float64) *megasquirt_gp16
	// Setsensors14 sets the physical value of the sensors14 signal.
	Setsensors14(float64) *megasquirt_gp16
	// Setsensors15 sets the physical value of the sensors15 signal.
	Setsensors15(float64) *megasquirt_gp16
	// Setsensors16 sets the physical value of the sensors16 signal.
	Setsensors16(float64) *megasquirt_gp16
}

type megasquirt_gp16 struct {
	xxx_sensors13 int16
	xxx_sensors14 int16
	xxx_sensors15 int16
	xxx_sensors16 int16
}

func Newmegasquirt_gp16() *megasquirt_gp16 {
	m := &megasquirt_gp16{}
	m.Reset()
	return m
}

func (m *megasquirt_gp16) Reset() {
	m.xxx_sensors13 = 0
	m.xxx_sensors14 = 0
	m.xxx_sensors15 = 0
	m.xxx_sensors16 = 0
}

func (m *megasquirt_gp16) CopyFrom(o megasquirt_gp16Reader) *megasquirt_gp16 {
	m.Setsensors13(o.sensors13())
	m.Setsensors14(o.sensors14())
	m.Setsensors15(o.sensors15())
	m.Setsensors16(o.sensors16())
	return m
}

// Descriptor returns the megasquirt_gp16 descriptor.
func (m *megasquirt_gp16) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp16.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp16) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp16) sensors13() float64 {
	return Messages().megasquirt_gp16.sensors13.ToPhysical(float64(m.xxx_sensors13))
}

func (m *megasquirt_gp16) Setsensors13(v float64) *megasquirt_gp16 {
	m.xxx_sensors13 = int16(Messages().megasquirt_gp16.sensors13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp16) sensors14() float64 {
	return Messages().megasquirt_gp16.sensors14.ToPhysical(float64(m.xxx_sensors14))
}

func (m *megasquirt_gp16) Setsensors14(v float64) *megasquirt_gp16 {
	m.xxx_sensors14 = int16(Messages().megasquirt_gp16.sensors14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp16) sensors15() float64 {
	return Messages().megasquirt_gp16.sensors15.ToPhysical(float64(m.xxx_sensors15))
}

func (m *megasquirt_gp16) Setsensors15(v float64) *megasquirt_gp16 {
	m.xxx_sensors15 = int16(Messages().megasquirt_gp16.sensors15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp16) sensors16() float64 {
	return Messages().megasquirt_gp16.sensors16.ToPhysical(float64(m.xxx_sensors16))
}

func (m *megasquirt_gp16) Setsensors16(v float64) *megasquirt_gp16 {
	m.xxx_sensors16 = int16(Messages().megasquirt_gp16.sensors16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp16) Frame() can.Frame {
	md := Messages().megasquirt_gp16
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.sensors13.MarshalSigned(&f.Data, int64(m.xxx_sensors13))
	md.sensors14.MarshalSigned(&f.Data, int64(m.xxx_sensors14))
	md.sensors15.MarshalSigned(&f.Data, int64(m.xxx_sensors15))
	md.sensors16.MarshalSigned(&f.Data, int64(m.xxx_sensors16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp16) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp16) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp16
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp16: expects ID 1536 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp16: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp16: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp16: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_sensors13 = int16(md.sensors13.UnmarshalSigned(f.Data))
	m.xxx_sensors14 = int16(md.sensors14.UnmarshalSigned(f.Data))
	m.xxx_sensors15 = int16(md.sensors15.UnmarshalSigned(f.Data))
	m.xxx_sensors16 = int16(md.sensors16.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp17Reader provides read access to a megasquirt_gp17 message.
type megasquirt_gp17Reader interface {
	// boost_targ_1 returns the physical value of the boost_targ_1 signal.
	boost_targ_1() float64
	// boost_targ_2 returns the physical value of the boost_targ_2 signal.
	boost_targ_2() float64
	// boostduty returns the value of the boostduty signal.
	boostduty() uint8
	// boostduty2 returns the value of the boostduty2 signal.
	boostduty2() uint8
	// maf_volts returns the physical value of the maf_volts signal.
	maf_volts() float64
}

// megasquirt_gp17Writer provides write access to a megasquirt_gp17 message.
type megasquirt_gp17Writer interface {
	// CopyFrom copies all values from megasquirt_gp17Reader.
	CopyFrom(megasquirt_gp17Reader) *megasquirt_gp17
	// Setboost_targ_1 sets the physical value of the boost_targ_1 signal.
	Setboost_targ_1(float64) *megasquirt_gp17
	// Setboost_targ_2 sets the physical value of the boost_targ_2 signal.
	Setboost_targ_2(float64) *megasquirt_gp17
	// Setboostduty sets the value of the boostduty signal.
	Setboostduty(uint8) *megasquirt_gp17
	// Setboostduty2 sets the value of the boostduty2 signal.
	Setboostduty2(uint8) *megasquirt_gp17
	// Setmaf_volts sets the physical value of the maf_volts signal.
	Setmaf_volts(float64) *megasquirt_gp17
}

type megasquirt_gp17 struct {
	xxx_boost_targ_1 uint16
	xxx_boost_targ_2 uint16
	xxx_boostduty    uint8
	xxx_boostduty2   uint8
	xxx_maf_volts    int16
}

func Newmegasquirt_gp17() *megasquirt_gp17 {
	m := &megasquirt_gp17{}
	m.Reset()
	return m
}

func (m *megasquirt_gp17) Reset() {
	m.xxx_boost_targ_1 = 0
	m.xxx_boost_targ_2 = 0
	m.xxx_boostduty = 0
	m.xxx_boostduty2 = 0
	m.xxx_maf_volts = 0
}

func (m *megasquirt_gp17) CopyFrom(o megasquirt_gp17Reader) *megasquirt_gp17 {
	m.Setboost_targ_1(o.boost_targ_1())
	m.Setboost_targ_2(o.boost_targ_2())
	m.xxx_boostduty = o.boostduty()
	m.xxx_boostduty2 = o.boostduty2()
	m.Setmaf_volts(o.maf_volts())
	return m
}

// Descriptor returns the megasquirt_gp17 descriptor.
func (m *megasquirt_gp17) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp17.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp17) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp17) boost_targ_1() float64 {
	return Messages().megasquirt_gp17.boost_targ_1.ToPhysical(float64(m.xxx_boost_targ_1))
}

func (m *megasquirt_gp17) Setboost_targ_1(v float64) *megasquirt_gp17 {
	m.xxx_boost_targ_1 = uint16(Messages().megasquirt_gp17.boost_targ_1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp17) boost_targ_2() float64 {
	return Messages().megasquirt_gp17.boost_targ_2.ToPhysical(float64(m.xxx_boost_targ_2))
}

func (m *megasquirt_gp17) Setboost_targ_2(v float64) *megasquirt_gp17 {
	m.xxx_boost_targ_2 = uint16(Messages().megasquirt_gp17.boost_targ_2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp17) boostduty() uint8 {
	return m.xxx_boostduty
}

func (m *megasquirt_gp17) Setboostduty(v uint8) *megasquirt_gp17 {
	m.xxx_boostduty = uint8(Messages().megasquirt_gp17.boostduty.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp17) boostduty2() uint8 {
	return m.xxx_boostduty2
}

func (m *megasquirt_gp17) Setboostduty2(v uint8) *megasquirt_gp17 {
	m.xxx_boostduty2 = uint8(Messages().megasquirt_gp17.boostduty2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp17) maf_volts() float64 {
	return Messages().megasquirt_gp17.maf_volts.ToPhysical(float64(m.xxx_maf_volts))
}

func (m *megasquirt_gp17) Setmaf_volts(v float64) *megasquirt_gp17 {
	m.xxx_maf_volts = int16(Messages().megasquirt_gp17.maf_volts.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp17) Frame() can.Frame {
	md := Messages().megasquirt_gp17
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.boost_targ_1.MarshalUnsigned(&f.Data, uint64(m.xxx_boost_targ_1))
	md.boost_targ_2.MarshalUnsigned(&f.Data, uint64(m.xxx_boost_targ_2))
	md.boostduty.MarshalUnsigned(&f.Data, uint64(m.xxx_boostduty))
	md.boostduty2.MarshalUnsigned(&f.Data, uint64(m.xxx_boostduty2))
	md.maf_volts.MarshalSigned(&f.Data, int64(m.xxx_maf_volts))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp17) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp17) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp17
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp17: expects ID 1537 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp17: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp17: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp17: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_boost_targ_1 = uint16(md.boost_targ_1.UnmarshalUnsigned(f.Data))
	m.xxx_boost_targ_2 = uint16(md.boost_targ_2.UnmarshalUnsigned(f.Data))
	m.xxx_boostduty = uint8(md.boostduty.UnmarshalUnsigned(f.Data))
	m.xxx_boostduty2 = uint8(md.boostduty2.UnmarshalUnsigned(f.Data))
	m.xxx_maf_volts = int16(md.maf_volts.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp18Reader provides read access to a megasquirt_gp18 message.
type megasquirt_gp18Reader interface {
	// pwseq1 returns the physical value of the pwseq1 signal.
	pwseq1() float64
	// pwseq2 returns the physical value of the pwseq2 signal.
	pwseq2() float64
	// pwseq3 returns the physical value of the pwseq3 signal.
	pwseq3() float64
	// pwseq4 returns the physical value of the pwseq4 signal.
	pwseq4() float64
}

// megasquirt_gp18Writer provides write access to a megasquirt_gp18 message.
type megasquirt_gp18Writer interface {
	// CopyFrom copies all values from megasquirt_gp18Reader.
	CopyFrom(megasquirt_gp18Reader) *megasquirt_gp18
	// Setpwseq1 sets the physical value of the pwseq1 signal.
	Setpwseq1(float64) *megasquirt_gp18
	// Setpwseq2 sets the physical value of the pwseq2 signal.
	Setpwseq2(float64) *megasquirt_gp18
	// Setpwseq3 sets the physical value of the pwseq3 signal.
	Setpwseq3(float64) *megasquirt_gp18
	// Setpwseq4 sets the physical value of the pwseq4 signal.
	Setpwseq4(float64) *megasquirt_gp18
}

type megasquirt_gp18 struct {
	xxx_pwseq1 int16
	xxx_pwseq2 int16
	xxx_pwseq3 int16
	xxx_pwseq4 int16
}

func Newmegasquirt_gp18() *megasquirt_gp18 {
	m := &megasquirt_gp18{}
	m.Reset()
	return m
}

func (m *megasquirt_gp18) Reset() {
	m.xxx_pwseq1 = 0
	m.xxx_pwseq2 = 0
	m.xxx_pwseq3 = 0
	m.xxx_pwseq4 = 0
}

func (m *megasquirt_gp18) CopyFrom(o megasquirt_gp18Reader) *megasquirt_gp18 {
	m.Setpwseq1(o.pwseq1())
	m.Setpwseq2(o.pwseq2())
	m.Setpwseq3(o.pwseq3())
	m.Setpwseq4(o.pwseq4())
	return m
}

// Descriptor returns the megasquirt_gp18 descriptor.
func (m *megasquirt_gp18) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp18.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp18) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp18) pwseq1() float64 {
	return Messages().megasquirt_gp18.pwseq1.ToPhysical(float64(m.xxx_pwseq1))
}

func (m *megasquirt_gp18) Setpwseq1(v float64) *megasquirt_gp18 {
	m.xxx_pwseq1 = int16(Messages().megasquirt_gp18.pwseq1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp18) pwseq2() float64 {
	return Messages().megasquirt_gp18.pwseq2.ToPhysical(float64(m.xxx_pwseq2))
}

func (m *megasquirt_gp18) Setpwseq2(v float64) *megasquirt_gp18 {
	m.xxx_pwseq2 = int16(Messages().megasquirt_gp18.pwseq2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp18) pwseq3() float64 {
	return Messages().megasquirt_gp18.pwseq3.ToPhysical(float64(m.xxx_pwseq3))
}

func (m *megasquirt_gp18) Setpwseq3(v float64) *megasquirt_gp18 {
	m.xxx_pwseq3 = int16(Messages().megasquirt_gp18.pwseq3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp18) pwseq4() float64 {
	return Messages().megasquirt_gp18.pwseq4.ToPhysical(float64(m.xxx_pwseq4))
}

func (m *megasquirt_gp18) Setpwseq4(v float64) *megasquirt_gp18 {
	m.xxx_pwseq4 = int16(Messages().megasquirt_gp18.pwseq4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp18) Frame() can.Frame {
	md := Messages().megasquirt_gp18
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.pwseq1.MarshalSigned(&f.Data, int64(m.xxx_pwseq1))
	md.pwseq2.MarshalSigned(&f.Data, int64(m.xxx_pwseq2))
	md.pwseq3.MarshalSigned(&f.Data, int64(m.xxx_pwseq3))
	md.pwseq4.MarshalSigned(&f.Data, int64(m.xxx_pwseq4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp18) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp18) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp18
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp18: expects ID 1538 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp18: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp18: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp18: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_pwseq1 = int16(md.pwseq1.UnmarshalSigned(f.Data))
	m.xxx_pwseq2 = int16(md.pwseq2.UnmarshalSigned(f.Data))
	m.xxx_pwseq3 = int16(md.pwseq3.UnmarshalSigned(f.Data))
	m.xxx_pwseq4 = int16(md.pwseq4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp19Reader provides read access to a megasquirt_gp19 message.
type megasquirt_gp19Reader interface {
	// pwseq5 returns the physical value of the pwseq5 signal.
	pwseq5() float64
	// pwseq6 returns the physical value of the pwseq6 signal.
	pwseq6() float64
	// pwseq7 returns the physical value of the pwseq7 signal.
	pwseq7() float64
	// pwseq8 returns the physical value of the pwseq8 signal.
	pwseq8() float64
}

// megasquirt_gp19Writer provides write access to a megasquirt_gp19 message.
type megasquirt_gp19Writer interface {
	// CopyFrom copies all values from megasquirt_gp19Reader.
	CopyFrom(megasquirt_gp19Reader) *megasquirt_gp19
	// Setpwseq5 sets the physical value of the pwseq5 signal.
	Setpwseq5(float64) *megasquirt_gp19
	// Setpwseq6 sets the physical value of the pwseq6 signal.
	Setpwseq6(float64) *megasquirt_gp19
	// Setpwseq7 sets the physical value of the pwseq7 signal.
	Setpwseq7(float64) *megasquirt_gp19
	// Setpwseq8 sets the physical value of the pwseq8 signal.
	Setpwseq8(float64) *megasquirt_gp19
}

type megasquirt_gp19 struct {
	xxx_pwseq5 int16
	xxx_pwseq6 int16
	xxx_pwseq7 int16
	xxx_pwseq8 int16
}

func Newmegasquirt_gp19() *megasquirt_gp19 {
	m := &megasquirt_gp19{}
	m.Reset()
	return m
}

func (m *megasquirt_gp19) Reset() {
	m.xxx_pwseq5 = 0
	m.xxx_pwseq6 = 0
	m.xxx_pwseq7 = 0
	m.xxx_pwseq8 = 0
}

func (m *megasquirt_gp19) CopyFrom(o megasquirt_gp19Reader) *megasquirt_gp19 {
	m.Setpwseq5(o.pwseq5())
	m.Setpwseq6(o.pwseq6())
	m.Setpwseq7(o.pwseq7())
	m.Setpwseq8(o.pwseq8())
	return m
}

// Descriptor returns the megasquirt_gp19 descriptor.
func (m *megasquirt_gp19) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp19.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp19) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp19) pwseq5() float64 {
	return Messages().megasquirt_gp19.pwseq5.ToPhysical(float64(m.xxx_pwseq5))
}

func (m *megasquirt_gp19) Setpwseq5(v float64) *megasquirt_gp19 {
	m.xxx_pwseq5 = int16(Messages().megasquirt_gp19.pwseq5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp19) pwseq6() float64 {
	return Messages().megasquirt_gp19.pwseq6.ToPhysical(float64(m.xxx_pwseq6))
}

func (m *megasquirt_gp19) Setpwseq6(v float64) *megasquirt_gp19 {
	m.xxx_pwseq6 = int16(Messages().megasquirt_gp19.pwseq6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp19) pwseq7() float64 {
	return Messages().megasquirt_gp19.pwseq7.ToPhysical(float64(m.xxx_pwseq7))
}

func (m *megasquirt_gp19) Setpwseq7(v float64) *megasquirt_gp19 {
	m.xxx_pwseq7 = int16(Messages().megasquirt_gp19.pwseq7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp19) pwseq8() float64 {
	return Messages().megasquirt_gp19.pwseq8.ToPhysical(float64(m.xxx_pwseq8))
}

func (m *megasquirt_gp19) Setpwseq8(v float64) *megasquirt_gp19 {
	m.xxx_pwseq8 = int16(Messages().megasquirt_gp19.pwseq8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp19) Frame() can.Frame {
	md := Messages().megasquirt_gp19
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.pwseq5.MarshalSigned(&f.Data, int64(m.xxx_pwseq5))
	md.pwseq6.MarshalSigned(&f.Data, int64(m.xxx_pwseq6))
	md.pwseq7.MarshalSigned(&f.Data, int64(m.xxx_pwseq7))
	md.pwseq8.MarshalSigned(&f.Data, int64(m.xxx_pwseq8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp19) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp19) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp19
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp19: expects ID 1539 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp19: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp19: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp19: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_pwseq5 = int16(md.pwseq5.UnmarshalSigned(f.Data))
	m.xxx_pwseq6 = int16(md.pwseq6.UnmarshalSigned(f.Data))
	m.xxx_pwseq7 = int16(md.pwseq7.UnmarshalSigned(f.Data))
	m.xxx_pwseq8 = int16(md.pwseq8.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp20Reader provides read access to a megasquirt_gp20 message.
type megasquirt_gp20Reader interface {
	// pwseq9 returns the physical value of the pwseq9 signal.
	pwseq9() float64
	// pwseq10 returns the physical value of the pwseq10 signal.
	pwseq10() float64
	// pwseq11 returns the physical value of the pwseq11 signal.
	pwseq11() float64
	// pwseq12 returns the physical value of the pwseq12 signal.
	pwseq12() float64
}

// megasquirt_gp20Writer provides write access to a megasquirt_gp20 message.
type megasquirt_gp20Writer interface {
	// CopyFrom copies all values from megasquirt_gp20Reader.
	CopyFrom(megasquirt_gp20Reader) *megasquirt_gp20
	// Setpwseq9 sets the physical value of the pwseq9 signal.
	Setpwseq9(float64) *megasquirt_gp20
	// Setpwseq10 sets the physical value of the pwseq10 signal.
	Setpwseq10(float64) *megasquirt_gp20
	// Setpwseq11 sets the physical value of the pwseq11 signal.
	Setpwseq11(float64) *megasquirt_gp20
	// Setpwseq12 sets the physical value of the pwseq12 signal.
	Setpwseq12(float64) *megasquirt_gp20
}

type megasquirt_gp20 struct {
	xxx_pwseq9  int16
	xxx_pwseq10 int16
	xxx_pwseq11 int16
	xxx_pwseq12 int16
}

func Newmegasquirt_gp20() *megasquirt_gp20 {
	m := &megasquirt_gp20{}
	m.Reset()
	return m
}

func (m *megasquirt_gp20) Reset() {
	m.xxx_pwseq9 = 0
	m.xxx_pwseq10 = 0
	m.xxx_pwseq11 = 0
	m.xxx_pwseq12 = 0
}

func (m *megasquirt_gp20) CopyFrom(o megasquirt_gp20Reader) *megasquirt_gp20 {
	m.Setpwseq9(o.pwseq9())
	m.Setpwseq10(o.pwseq10())
	m.Setpwseq11(o.pwseq11())
	m.Setpwseq12(o.pwseq12())
	return m
}

// Descriptor returns the megasquirt_gp20 descriptor.
func (m *megasquirt_gp20) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp20.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp20) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp20) pwseq9() float64 {
	return Messages().megasquirt_gp20.pwseq9.ToPhysical(float64(m.xxx_pwseq9))
}

func (m *megasquirt_gp20) Setpwseq9(v float64) *megasquirt_gp20 {
	m.xxx_pwseq9 = int16(Messages().megasquirt_gp20.pwseq9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp20) pwseq10() float64 {
	return Messages().megasquirt_gp20.pwseq10.ToPhysical(float64(m.xxx_pwseq10))
}

func (m *megasquirt_gp20) Setpwseq10(v float64) *megasquirt_gp20 {
	m.xxx_pwseq10 = int16(Messages().megasquirt_gp20.pwseq10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp20) pwseq11() float64 {
	return Messages().megasquirt_gp20.pwseq11.ToPhysical(float64(m.xxx_pwseq11))
}

func (m *megasquirt_gp20) Setpwseq11(v float64) *megasquirt_gp20 {
	m.xxx_pwseq11 = int16(Messages().megasquirt_gp20.pwseq11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp20) pwseq12() float64 {
	return Messages().megasquirt_gp20.pwseq12.ToPhysical(float64(m.xxx_pwseq12))
}

func (m *megasquirt_gp20) Setpwseq12(v float64) *megasquirt_gp20 {
	m.xxx_pwseq12 = int16(Messages().megasquirt_gp20.pwseq12.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp20) Frame() can.Frame {
	md := Messages().megasquirt_gp20
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.pwseq9.MarshalSigned(&f.Data, int64(m.xxx_pwseq9))
	md.pwseq10.MarshalSigned(&f.Data, int64(m.xxx_pwseq10))
	md.pwseq11.MarshalSigned(&f.Data, int64(m.xxx_pwseq11))
	md.pwseq12.MarshalSigned(&f.Data, int64(m.xxx_pwseq12))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp20) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp20) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp20
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp20: expects ID 1540 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp20: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp20: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp20: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_pwseq9 = int16(md.pwseq9.UnmarshalSigned(f.Data))
	m.xxx_pwseq10 = int16(md.pwseq10.UnmarshalSigned(f.Data))
	m.xxx_pwseq11 = int16(md.pwseq11.UnmarshalSigned(f.Data))
	m.xxx_pwseq12 = int16(md.pwseq12.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp21Reader provides read access to a megasquirt_gp21 message.
type megasquirt_gp21Reader interface {
	// pwseq13 returns the physical value of the pwseq13 signal.
	pwseq13() float64
	// pwseq14 returns the physical value of the pwseq14 signal.
	pwseq14() float64
	// pwseq15 returns the physical value of the pwseq15 signal.
	pwseq15() float64
	// pwseq16 returns the physical value of the pwseq16 signal.
	pwseq16() float64
}

// megasquirt_gp21Writer provides write access to a megasquirt_gp21 message.
type megasquirt_gp21Writer interface {
	// CopyFrom copies all values from megasquirt_gp21Reader.
	CopyFrom(megasquirt_gp21Reader) *megasquirt_gp21
	// Setpwseq13 sets the physical value of the pwseq13 signal.
	Setpwseq13(float64) *megasquirt_gp21
	// Setpwseq14 sets the physical value of the pwseq14 signal.
	Setpwseq14(float64) *megasquirt_gp21
	// Setpwseq15 sets the physical value of the pwseq15 signal.
	Setpwseq15(float64) *megasquirt_gp21
	// Setpwseq16 sets the physical value of the pwseq16 signal.
	Setpwseq16(float64) *megasquirt_gp21
}

type megasquirt_gp21 struct {
	xxx_pwseq13 int16
	xxx_pwseq14 int16
	xxx_pwseq15 int16
	xxx_pwseq16 int16
}

func Newmegasquirt_gp21() *megasquirt_gp21 {
	m := &megasquirt_gp21{}
	m.Reset()
	return m
}

func (m *megasquirt_gp21) Reset() {
	m.xxx_pwseq13 = 0
	m.xxx_pwseq14 = 0
	m.xxx_pwseq15 = 0
	m.xxx_pwseq16 = 0
}

func (m *megasquirt_gp21) CopyFrom(o megasquirt_gp21Reader) *megasquirt_gp21 {
	m.Setpwseq13(o.pwseq13())
	m.Setpwseq14(o.pwseq14())
	m.Setpwseq15(o.pwseq15())
	m.Setpwseq16(o.pwseq16())
	return m
}

// Descriptor returns the megasquirt_gp21 descriptor.
func (m *megasquirt_gp21) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp21.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp21) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp21) pwseq13() float64 {
	return Messages().megasquirt_gp21.pwseq13.ToPhysical(float64(m.xxx_pwseq13))
}

func (m *megasquirt_gp21) Setpwseq13(v float64) *megasquirt_gp21 {
	m.xxx_pwseq13 = int16(Messages().megasquirt_gp21.pwseq13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp21) pwseq14() float64 {
	return Messages().megasquirt_gp21.pwseq14.ToPhysical(float64(m.xxx_pwseq14))
}

func (m *megasquirt_gp21) Setpwseq14(v float64) *megasquirt_gp21 {
	m.xxx_pwseq14 = int16(Messages().megasquirt_gp21.pwseq14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp21) pwseq15() float64 {
	return Messages().megasquirt_gp21.pwseq15.ToPhysical(float64(m.xxx_pwseq15))
}

func (m *megasquirt_gp21) Setpwseq15(v float64) *megasquirt_gp21 {
	m.xxx_pwseq15 = int16(Messages().megasquirt_gp21.pwseq15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp21) pwseq16() float64 {
	return Messages().megasquirt_gp21.pwseq16.ToPhysical(float64(m.xxx_pwseq16))
}

func (m *megasquirt_gp21) Setpwseq16(v float64) *megasquirt_gp21 {
	m.xxx_pwseq16 = int16(Messages().megasquirt_gp21.pwseq16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp21) Frame() can.Frame {
	md := Messages().megasquirt_gp21
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.pwseq13.MarshalSigned(&f.Data, int64(m.xxx_pwseq13))
	md.pwseq14.MarshalSigned(&f.Data, int64(m.xxx_pwseq14))
	md.pwseq15.MarshalSigned(&f.Data, int64(m.xxx_pwseq15))
	md.pwseq16.MarshalSigned(&f.Data, int64(m.xxx_pwseq16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp21) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp21) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp21
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp21: expects ID 1541 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp21: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp21: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp21: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_pwseq13 = int16(md.pwseq13.UnmarshalSigned(f.Data))
	m.xxx_pwseq14 = int16(md.pwseq14.UnmarshalSigned(f.Data))
	m.xxx_pwseq15 = int16(md.pwseq15.UnmarshalSigned(f.Data))
	m.xxx_pwseq16 = int16(md.pwseq16.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp22Reader provides read access to a megasquirt_gp22 message.
type megasquirt_gp22Reader interface {
	// egt1 returns the physical value of the egt1 signal.
	egt1() float64
	// egt2 returns the physical value of the egt2 signal.
	egt2() float64
	// egt3 returns the physical value of the egt3 signal.
	egt3() float64
	// egt4 returns the physical value of the egt4 signal.
	egt4() float64
}

// megasquirt_gp22Writer provides write access to a megasquirt_gp22 message.
type megasquirt_gp22Writer interface {
	// CopyFrom copies all values from megasquirt_gp22Reader.
	CopyFrom(megasquirt_gp22Reader) *megasquirt_gp22
	// Setegt1 sets the physical value of the egt1 signal.
	Setegt1(float64) *megasquirt_gp22
	// Setegt2 sets the physical value of the egt2 signal.
	Setegt2(float64) *megasquirt_gp22
	// Setegt3 sets the physical value of the egt3 signal.
	Setegt3(float64) *megasquirt_gp22
	// Setegt4 sets the physical value of the egt4 signal.
	Setegt4(float64) *megasquirt_gp22
}

type megasquirt_gp22 struct {
	xxx_egt1 int16
	xxx_egt2 int16
	xxx_egt3 int16
	xxx_egt4 int16
}

func Newmegasquirt_gp22() *megasquirt_gp22 {
	m := &megasquirt_gp22{}
	m.Reset()
	return m
}

func (m *megasquirt_gp22) Reset() {
	m.xxx_egt1 = 0
	m.xxx_egt2 = 0
	m.xxx_egt3 = 0
	m.xxx_egt4 = 0
}

func (m *megasquirt_gp22) CopyFrom(o megasquirt_gp22Reader) *megasquirt_gp22 {
	m.Setegt1(o.egt1())
	m.Setegt2(o.egt2())
	m.Setegt3(o.egt3())
	m.Setegt4(o.egt4())
	return m
}

// Descriptor returns the megasquirt_gp22 descriptor.
func (m *megasquirt_gp22) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp22.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp22) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp22) egt1() float64 {
	return Messages().megasquirt_gp22.egt1.ToPhysical(float64(m.xxx_egt1))
}

func (m *megasquirt_gp22) Setegt1(v float64) *megasquirt_gp22 {
	m.xxx_egt1 = int16(Messages().megasquirt_gp22.egt1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp22) egt2() float64 {
	return Messages().megasquirt_gp22.egt2.ToPhysical(float64(m.xxx_egt2))
}

func (m *megasquirt_gp22) Setegt2(v float64) *megasquirt_gp22 {
	m.xxx_egt2 = int16(Messages().megasquirt_gp22.egt2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp22) egt3() float64 {
	return Messages().megasquirt_gp22.egt3.ToPhysical(float64(m.xxx_egt3))
}

func (m *megasquirt_gp22) Setegt3(v float64) *megasquirt_gp22 {
	m.xxx_egt3 = int16(Messages().megasquirt_gp22.egt3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp22) egt4() float64 {
	return Messages().megasquirt_gp22.egt4.ToPhysical(float64(m.xxx_egt4))
}

func (m *megasquirt_gp22) Setegt4(v float64) *megasquirt_gp22 {
	m.xxx_egt4 = int16(Messages().megasquirt_gp22.egt4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp22) Frame() can.Frame {
	md := Messages().megasquirt_gp22
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.egt1.MarshalSigned(&f.Data, int64(m.xxx_egt1))
	md.egt2.MarshalSigned(&f.Data, int64(m.xxx_egt2))
	md.egt3.MarshalSigned(&f.Data, int64(m.xxx_egt3))
	md.egt4.MarshalSigned(&f.Data, int64(m.xxx_egt4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp22) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp22) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp22
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp22: expects ID 1542 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp22: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp22: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp22: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_egt1 = int16(md.egt1.UnmarshalSigned(f.Data))
	m.xxx_egt2 = int16(md.egt2.UnmarshalSigned(f.Data))
	m.xxx_egt3 = int16(md.egt3.UnmarshalSigned(f.Data))
	m.xxx_egt4 = int16(md.egt4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp23Reader provides read access to a megasquirt_gp23 message.
type megasquirt_gp23Reader interface {
	// egt5 returns the physical value of the egt5 signal.
	egt5() float64
	// egt6 returns the physical value of the egt6 signal.
	egt6() float64
	// egt7 returns the physical value of the egt7 signal.
	egt7() float64
	// egt8 returns the physical value of the egt8 signal.
	egt8() float64
}

// megasquirt_gp23Writer provides write access to a megasquirt_gp23 message.
type megasquirt_gp23Writer interface {
	// CopyFrom copies all values from megasquirt_gp23Reader.
	CopyFrom(megasquirt_gp23Reader) *megasquirt_gp23
	// Setegt5 sets the physical value of the egt5 signal.
	Setegt5(float64) *megasquirt_gp23
	// Setegt6 sets the physical value of the egt6 signal.
	Setegt6(float64) *megasquirt_gp23
	// Setegt7 sets the physical value of the egt7 signal.
	Setegt7(float64) *megasquirt_gp23
	// Setegt8 sets the physical value of the egt8 signal.
	Setegt8(float64) *megasquirt_gp23
}

type megasquirt_gp23 struct {
	xxx_egt5 int16
	xxx_egt6 int16
	xxx_egt7 int16
	xxx_egt8 int16
}

func Newmegasquirt_gp23() *megasquirt_gp23 {
	m := &megasquirt_gp23{}
	m.Reset()
	return m
}

func (m *megasquirt_gp23) Reset() {
	m.xxx_egt5 = 0
	m.xxx_egt6 = 0
	m.xxx_egt7 = 0
	m.xxx_egt8 = 0
}

func (m *megasquirt_gp23) CopyFrom(o megasquirt_gp23Reader) *megasquirt_gp23 {
	m.Setegt5(o.egt5())
	m.Setegt6(o.egt6())
	m.Setegt7(o.egt7())
	m.Setegt8(o.egt8())
	return m
}

// Descriptor returns the megasquirt_gp23 descriptor.
func (m *megasquirt_gp23) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp23.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp23) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp23) egt5() float64 {
	return Messages().megasquirt_gp23.egt5.ToPhysical(float64(m.xxx_egt5))
}

func (m *megasquirt_gp23) Setegt5(v float64) *megasquirt_gp23 {
	m.xxx_egt5 = int16(Messages().megasquirt_gp23.egt5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp23) egt6() float64 {
	return Messages().megasquirt_gp23.egt6.ToPhysical(float64(m.xxx_egt6))
}

func (m *megasquirt_gp23) Setegt6(v float64) *megasquirt_gp23 {
	m.xxx_egt6 = int16(Messages().megasquirt_gp23.egt6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp23) egt7() float64 {
	return Messages().megasquirt_gp23.egt7.ToPhysical(float64(m.xxx_egt7))
}

func (m *megasquirt_gp23) Setegt7(v float64) *megasquirt_gp23 {
	m.xxx_egt7 = int16(Messages().megasquirt_gp23.egt7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp23) egt8() float64 {
	return Messages().megasquirt_gp23.egt8.ToPhysical(float64(m.xxx_egt8))
}

func (m *megasquirt_gp23) Setegt8(v float64) *megasquirt_gp23 {
	m.xxx_egt8 = int16(Messages().megasquirt_gp23.egt8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp23) Frame() can.Frame {
	md := Messages().megasquirt_gp23
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.egt5.MarshalSigned(&f.Data, int64(m.xxx_egt5))
	md.egt6.MarshalSigned(&f.Data, int64(m.xxx_egt6))
	md.egt7.MarshalSigned(&f.Data, int64(m.xxx_egt7))
	md.egt8.MarshalSigned(&f.Data, int64(m.xxx_egt8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp23) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp23) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp23
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp23: expects ID 1543 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp23: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp23: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp23: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_egt5 = int16(md.egt5.UnmarshalSigned(f.Data))
	m.xxx_egt6 = int16(md.egt6.UnmarshalSigned(f.Data))
	m.xxx_egt7 = int16(md.egt7.UnmarshalSigned(f.Data))
	m.xxx_egt8 = int16(md.egt8.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp24Reader provides read access to a megasquirt_gp24 message.
type megasquirt_gp24Reader interface {
	// egt9 returns the physical value of the egt9 signal.
	egt9() float64
	// egt10 returns the physical value of the egt10 signal.
	egt10() float64
	// egt11 returns the physical value of the egt11 signal.
	egt11() float64
	// egt12 returns the physical value of the egt12 signal.
	egt12() float64
}

// megasquirt_gp24Writer provides write access to a megasquirt_gp24 message.
type megasquirt_gp24Writer interface {
	// CopyFrom copies all values from megasquirt_gp24Reader.
	CopyFrom(megasquirt_gp24Reader) *megasquirt_gp24
	// Setegt9 sets the physical value of the egt9 signal.
	Setegt9(float64) *megasquirt_gp24
	// Setegt10 sets the physical value of the egt10 signal.
	Setegt10(float64) *megasquirt_gp24
	// Setegt11 sets the physical value of the egt11 signal.
	Setegt11(float64) *megasquirt_gp24
	// Setegt12 sets the physical value of the egt12 signal.
	Setegt12(float64) *megasquirt_gp24
}

type megasquirt_gp24 struct {
	xxx_egt9  int16
	xxx_egt10 int16
	xxx_egt11 int16
	xxx_egt12 int16
}

func Newmegasquirt_gp24() *megasquirt_gp24 {
	m := &megasquirt_gp24{}
	m.Reset()
	return m
}

func (m *megasquirt_gp24) Reset() {
	m.xxx_egt9 = 0
	m.xxx_egt10 = 0
	m.xxx_egt11 = 0
	m.xxx_egt12 = 0
}

func (m *megasquirt_gp24) CopyFrom(o megasquirt_gp24Reader) *megasquirt_gp24 {
	m.Setegt9(o.egt9())
	m.Setegt10(o.egt10())
	m.Setegt11(o.egt11())
	m.Setegt12(o.egt12())
	return m
}

// Descriptor returns the megasquirt_gp24 descriptor.
func (m *megasquirt_gp24) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp24.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp24) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp24) egt9() float64 {
	return Messages().megasquirt_gp24.egt9.ToPhysical(float64(m.xxx_egt9))
}

func (m *megasquirt_gp24) Setegt9(v float64) *megasquirt_gp24 {
	m.xxx_egt9 = int16(Messages().megasquirt_gp24.egt9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp24) egt10() float64 {
	return Messages().megasquirt_gp24.egt10.ToPhysical(float64(m.xxx_egt10))
}

func (m *megasquirt_gp24) Setegt10(v float64) *megasquirt_gp24 {
	m.xxx_egt10 = int16(Messages().megasquirt_gp24.egt10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp24) egt11() float64 {
	return Messages().megasquirt_gp24.egt11.ToPhysical(float64(m.xxx_egt11))
}

func (m *megasquirt_gp24) Setegt11(v float64) *megasquirt_gp24 {
	m.xxx_egt11 = int16(Messages().megasquirt_gp24.egt11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp24) egt12() float64 {
	return Messages().megasquirt_gp24.egt12.ToPhysical(float64(m.xxx_egt12))
}

func (m *megasquirt_gp24) Setegt12(v float64) *megasquirt_gp24 {
	m.xxx_egt12 = int16(Messages().megasquirt_gp24.egt12.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp24) Frame() can.Frame {
	md := Messages().megasquirt_gp24
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.egt9.MarshalSigned(&f.Data, int64(m.xxx_egt9))
	md.egt10.MarshalSigned(&f.Data, int64(m.xxx_egt10))
	md.egt11.MarshalSigned(&f.Data, int64(m.xxx_egt11))
	md.egt12.MarshalSigned(&f.Data, int64(m.xxx_egt12))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp24) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp24) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp24
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp24: expects ID 1544 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp24: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp24: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp24: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_egt9 = int16(md.egt9.UnmarshalSigned(f.Data))
	m.xxx_egt10 = int16(md.egt10.UnmarshalSigned(f.Data))
	m.xxx_egt11 = int16(md.egt11.UnmarshalSigned(f.Data))
	m.xxx_egt12 = int16(md.egt12.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp25Reader provides read access to a megasquirt_gp25 message.
type megasquirt_gp25Reader interface {
	// egt13 returns the physical value of the egt13 signal.
	egt13() float64
	// egt14 returns the physical value of the egt14 signal.
	egt14() float64
	// egt15 returns the physical value of the egt15 signal.
	egt15() float64
	// egt16 returns the physical value of the egt16 signal.
	egt16() float64
}

// megasquirt_gp25Writer provides write access to a megasquirt_gp25 message.
type megasquirt_gp25Writer interface {
	// CopyFrom copies all values from megasquirt_gp25Reader.
	CopyFrom(megasquirt_gp25Reader) *megasquirt_gp25
	// Setegt13 sets the physical value of the egt13 signal.
	Setegt13(float64) *megasquirt_gp25
	// Setegt14 sets the physical value of the egt14 signal.
	Setegt14(float64) *megasquirt_gp25
	// Setegt15 sets the physical value of the egt15 signal.
	Setegt15(float64) *megasquirt_gp25
	// Setegt16 sets the physical value of the egt16 signal.
	Setegt16(float64) *megasquirt_gp25
}

type megasquirt_gp25 struct {
	xxx_egt13 int16
	xxx_egt14 int16
	xxx_egt15 int16
	xxx_egt16 int16
}

func Newmegasquirt_gp25() *megasquirt_gp25 {
	m := &megasquirt_gp25{}
	m.Reset()
	return m
}

func (m *megasquirt_gp25) Reset() {
	m.xxx_egt13 = 0
	m.xxx_egt14 = 0
	m.xxx_egt15 = 0
	m.xxx_egt16 = 0
}

func (m *megasquirt_gp25) CopyFrom(o megasquirt_gp25Reader) *megasquirt_gp25 {
	m.Setegt13(o.egt13())
	m.Setegt14(o.egt14())
	m.Setegt15(o.egt15())
	m.Setegt16(o.egt16())
	return m
}

// Descriptor returns the megasquirt_gp25 descriptor.
func (m *megasquirt_gp25) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp25.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp25) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp25) egt13() float64 {
	return Messages().megasquirt_gp25.egt13.ToPhysical(float64(m.xxx_egt13))
}

func (m *megasquirt_gp25) Setegt13(v float64) *megasquirt_gp25 {
	m.xxx_egt13 = int16(Messages().megasquirt_gp25.egt13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp25) egt14() float64 {
	return Messages().megasquirt_gp25.egt14.ToPhysical(float64(m.xxx_egt14))
}

func (m *megasquirt_gp25) Setegt14(v float64) *megasquirt_gp25 {
	m.xxx_egt14 = int16(Messages().megasquirt_gp25.egt14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp25) egt15() float64 {
	return Messages().megasquirt_gp25.egt15.ToPhysical(float64(m.xxx_egt15))
}

func (m *megasquirt_gp25) Setegt15(v float64) *megasquirt_gp25 {
	m.xxx_egt15 = int16(Messages().megasquirt_gp25.egt15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp25) egt16() float64 {
	return Messages().megasquirt_gp25.egt16.ToPhysical(float64(m.xxx_egt16))
}

func (m *megasquirt_gp25) Setegt16(v float64) *megasquirt_gp25 {
	m.xxx_egt16 = int16(Messages().megasquirt_gp25.egt16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp25) Frame() can.Frame {
	md := Messages().megasquirt_gp25
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.egt13.MarshalSigned(&f.Data, int64(m.xxx_egt13))
	md.egt14.MarshalSigned(&f.Data, int64(m.xxx_egt14))
	md.egt15.MarshalSigned(&f.Data, int64(m.xxx_egt15))
	md.egt16.MarshalSigned(&f.Data, int64(m.xxx_egt16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp25) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp25) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp25
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp25: expects ID 1545 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp25: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp25: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp25: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_egt13 = int16(md.egt13.UnmarshalSigned(f.Data))
	m.xxx_egt14 = int16(md.egt14.UnmarshalSigned(f.Data))
	m.xxx_egt15 = int16(md.egt15.UnmarshalSigned(f.Data))
	m.xxx_egt16 = int16(md.egt16.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp26Reader provides read access to a megasquirt_gp26 message.
type megasquirt_gp26Reader interface {
	// nitrous1_duty returns the value of the nitrous1_duty signal.
	nitrous1_duty() uint8
	// nitrous2_duty returns the value of the nitrous2_duty signal.
	nitrous2_duty() uint8
	// nitrous_timer_out returns the physical value of the nitrous_timer_out signal.
	nitrous_timer_out() float64
	// n2o_addfuel returns the physical value of the n2o_addfuel signal.
	n2o_addfuel() float64
	// n2o_retard returns the physical value of the n2o_retard signal.
	n2o_retard() float64
}

// megasquirt_gp26Writer provides write access to a megasquirt_gp26 message.
type megasquirt_gp26Writer interface {
	// CopyFrom copies all values from megasquirt_gp26Reader.
	CopyFrom(megasquirt_gp26Reader) *megasquirt_gp26
	// Setnitrous1_duty sets the value of the nitrous1_duty signal.
	Setnitrous1_duty(uint8) *megasquirt_gp26
	// Setnitrous2_duty sets the value of the nitrous2_duty signal.
	Setnitrous2_duty(uint8) *megasquirt_gp26
	// Setnitrous_timer_out sets the physical value of the nitrous_timer_out signal.
	Setnitrous_timer_out(float64) *megasquirt_gp26
	// Setn2o_addfuel sets the physical value of the n2o_addfuel signal.
	Setn2o_addfuel(float64) *megasquirt_gp26
	// Setn2o_retard sets the physical value of the n2o_retard signal.
	Setn2o_retard(float64) *megasquirt_gp26
}

type megasquirt_gp26 struct {
	xxx_nitrous1_duty     uint8
	xxx_nitrous2_duty     uint8
	xxx_nitrous_timer_out uint16
	xxx_n2o_addfuel       int16
	xxx_n2o_retard        int16
}

func Newmegasquirt_gp26() *megasquirt_gp26 {
	m := &megasquirt_gp26{}
	m.Reset()
	return m
}

func (m *megasquirt_gp26) Reset() {
	m.xxx_nitrous1_duty = 0
	m.xxx_nitrous2_duty = 0
	m.xxx_nitrous_timer_out = 0
	m.xxx_n2o_addfuel = 0
	m.xxx_n2o_retard = 0
}

func (m *megasquirt_gp26) CopyFrom(o megasquirt_gp26Reader) *megasquirt_gp26 {
	m.xxx_nitrous1_duty = o.nitrous1_duty()
	m.xxx_nitrous2_duty = o.nitrous2_duty()
	m.Setnitrous_timer_out(o.nitrous_timer_out())
	m.Setn2o_addfuel(o.n2o_addfuel())
	m.Setn2o_retard(o.n2o_retard())
	return m
}

// Descriptor returns the megasquirt_gp26 descriptor.
func (m *megasquirt_gp26) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp26.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp26) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp26) nitrous1_duty() uint8 {
	return m.xxx_nitrous1_duty
}

func (m *megasquirt_gp26) Setnitrous1_duty(v uint8) *megasquirt_gp26 {
	m.xxx_nitrous1_duty = uint8(Messages().megasquirt_gp26.nitrous1_duty.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp26) nitrous2_duty() uint8 {
	return m.xxx_nitrous2_duty
}

func (m *megasquirt_gp26) Setnitrous2_duty(v uint8) *megasquirt_gp26 {
	m.xxx_nitrous2_duty = uint8(Messages().megasquirt_gp26.nitrous2_duty.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp26) nitrous_timer_out() float64 {
	return Messages().megasquirt_gp26.nitrous_timer_out.ToPhysical(float64(m.xxx_nitrous_timer_out))
}

func (m *megasquirt_gp26) Setnitrous_timer_out(v float64) *megasquirt_gp26 {
	m.xxx_nitrous_timer_out = uint16(Messages().megasquirt_gp26.nitrous_timer_out.FromPhysical(v))
	return m
}

func (m *megasquirt_gp26) n2o_addfuel() float64 {
	return Messages().megasquirt_gp26.n2o_addfuel.ToPhysical(float64(m.xxx_n2o_addfuel))
}

func (m *megasquirt_gp26) Setn2o_addfuel(v float64) *megasquirt_gp26 {
	m.xxx_n2o_addfuel = int16(Messages().megasquirt_gp26.n2o_addfuel.FromPhysical(v))
	return m
}

func (m *megasquirt_gp26) n2o_retard() float64 {
	return Messages().megasquirt_gp26.n2o_retard.ToPhysical(float64(m.xxx_n2o_retard))
}

func (m *megasquirt_gp26) Setn2o_retard(v float64) *megasquirt_gp26 {
	m.xxx_n2o_retard = int16(Messages().megasquirt_gp26.n2o_retard.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp26) Frame() can.Frame {
	md := Messages().megasquirt_gp26
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.nitrous1_duty.MarshalUnsigned(&f.Data, uint64(m.xxx_nitrous1_duty))
	md.nitrous2_duty.MarshalUnsigned(&f.Data, uint64(m.xxx_nitrous2_duty))
	md.nitrous_timer_out.MarshalUnsigned(&f.Data, uint64(m.xxx_nitrous_timer_out))
	md.n2o_addfuel.MarshalSigned(&f.Data, int64(m.xxx_n2o_addfuel))
	md.n2o_retard.MarshalSigned(&f.Data, int64(m.xxx_n2o_retard))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp26) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp26) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp26
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp26: expects ID 1546 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp26: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp26: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp26: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_nitrous1_duty = uint8(md.nitrous1_duty.UnmarshalUnsigned(f.Data))
	m.xxx_nitrous2_duty = uint8(md.nitrous2_duty.UnmarshalUnsigned(f.Data))
	m.xxx_nitrous_timer_out = uint16(md.nitrous_timer_out.UnmarshalUnsigned(f.Data))
	m.xxx_n2o_addfuel = int16(md.n2o_addfuel.UnmarshalSigned(f.Data))
	m.xxx_n2o_retard = int16(md.n2o_retard.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp27Reader provides read access to a megasquirt_gp27 message.
type megasquirt_gp27Reader interface {
	// canpwmin1 returns the value of the canpwmin1 signal.
	canpwmin1() int16
	// canpwmin2 returns the value of the canpwmin2 signal.
	canpwmin2() int16
	// canpwmin3 returns the value of the canpwmin3 signal.
	canpwmin3() int16
	// canpwmin4 returns the value of the canpwmin4 signal.
	canpwmin4() int16
}

// megasquirt_gp27Writer provides write access to a megasquirt_gp27 message.
type megasquirt_gp27Writer interface {
	// CopyFrom copies all values from megasquirt_gp27Reader.
	CopyFrom(megasquirt_gp27Reader) *megasquirt_gp27
	// Setcanpwmin1 sets the value of the canpwmin1 signal.
	Setcanpwmin1(int16) *megasquirt_gp27
	// Setcanpwmin2 sets the value of the canpwmin2 signal.
	Setcanpwmin2(int16) *megasquirt_gp27
	// Setcanpwmin3 sets the value of the canpwmin3 signal.
	Setcanpwmin3(int16) *megasquirt_gp27
	// Setcanpwmin4 sets the value of the canpwmin4 signal.
	Setcanpwmin4(int16) *megasquirt_gp27
}

type megasquirt_gp27 struct {
	xxx_canpwmin1 int16
	xxx_canpwmin2 int16
	xxx_canpwmin3 int16
	xxx_canpwmin4 int16
}

func Newmegasquirt_gp27() *megasquirt_gp27 {
	m := &megasquirt_gp27{}
	m.Reset()
	return m
}

func (m *megasquirt_gp27) Reset() {
	m.xxx_canpwmin1 = 0
	m.xxx_canpwmin2 = 0
	m.xxx_canpwmin3 = 0
	m.xxx_canpwmin4 = 0
}

func (m *megasquirt_gp27) CopyFrom(o megasquirt_gp27Reader) *megasquirt_gp27 {
	m.xxx_canpwmin1 = o.canpwmin1()
	m.xxx_canpwmin2 = o.canpwmin2()
	m.xxx_canpwmin3 = o.canpwmin3()
	m.xxx_canpwmin4 = o.canpwmin4()
	return m
}

// Descriptor returns the megasquirt_gp27 descriptor.
func (m *megasquirt_gp27) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp27.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp27) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp27) canpwmin1() int16 {
	return m.xxx_canpwmin1
}

func (m *megasquirt_gp27) Setcanpwmin1(v int16) *megasquirt_gp27 {
	m.xxx_canpwmin1 = int16(Messages().megasquirt_gp27.canpwmin1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp27) canpwmin2() int16 {
	return m.xxx_canpwmin2
}

func (m *megasquirt_gp27) Setcanpwmin2(v int16) *megasquirt_gp27 {
	m.xxx_canpwmin2 = int16(Messages().megasquirt_gp27.canpwmin2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp27) canpwmin3() int16 {
	return m.xxx_canpwmin3
}

func (m *megasquirt_gp27) Setcanpwmin3(v int16) *megasquirt_gp27 {
	m.xxx_canpwmin3 = int16(Messages().megasquirt_gp27.canpwmin3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp27) canpwmin4() int16 {
	return m.xxx_canpwmin4
}

func (m *megasquirt_gp27) Setcanpwmin4(v int16) *megasquirt_gp27 {
	m.xxx_canpwmin4 = int16(Messages().megasquirt_gp27.canpwmin4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp27) Frame() can.Frame {
	md := Messages().megasquirt_gp27
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.canpwmin1.MarshalSigned(&f.Data, int64(m.xxx_canpwmin1))
	md.canpwmin2.MarshalSigned(&f.Data, int64(m.xxx_canpwmin2))
	md.canpwmin3.MarshalSigned(&f.Data, int64(m.xxx_canpwmin3))
	md.canpwmin4.MarshalSigned(&f.Data, int64(m.xxx_canpwmin4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp27) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp27) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp27
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp27: expects ID 1547 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp27: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp27: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp27: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_canpwmin1 = int16(md.canpwmin1.UnmarshalSigned(f.Data))
	m.xxx_canpwmin2 = int16(md.canpwmin2.UnmarshalSigned(f.Data))
	m.xxx_canpwmin3 = int16(md.canpwmin3.UnmarshalSigned(f.Data))
	m.xxx_canpwmin4 = int16(md.canpwmin4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp28Reader provides read access to a megasquirt_gp28 message.
type megasquirt_gp28Reader interface {
	// cl_idle_targ_rpm returns the value of the cl_idle_targ_rpm signal.
	cl_idle_targ_rpm() uint16
	// tpsadc returns the value of the tpsadc signal.
	tpsadc() int16
	// eaeload returns the physical value of the eaeload signal.
	eaeload() float64
	// afrload returns the physical value of the afrload signal.
	afrload() float64
}

// megasquirt_gp28Writer provides write access to a megasquirt_gp28 message.
type megasquirt_gp28Writer interface {
	// CopyFrom copies all values from megasquirt_gp28Reader.
	CopyFrom(megasquirt_gp28Reader) *megasquirt_gp28
	// Setcl_idle_targ_rpm sets the value of the cl_idle_targ_rpm signal.
	Setcl_idle_targ_rpm(uint16) *megasquirt_gp28
	// Settpsadc sets the value of the tpsadc signal.
	Settpsadc(int16) *megasquirt_gp28
	// Seteaeload sets the physical value of the eaeload signal.
	Seteaeload(float64) *megasquirt_gp28
	// Setafrload sets the physical value of the afrload signal.
	Setafrload(float64) *megasquirt_gp28
}

type megasquirt_gp28 struct {
	xxx_cl_idle_targ_rpm uint16
	xxx_tpsadc           int16
	xxx_eaeload          int16
	xxx_afrload          int16
}

func Newmegasquirt_gp28() *megasquirt_gp28 {
	m := &megasquirt_gp28{}
	m.Reset()
	return m
}

func (m *megasquirt_gp28) Reset() {
	m.xxx_cl_idle_targ_rpm = 0
	m.xxx_tpsadc = 0
	m.xxx_eaeload = 0
	m.xxx_afrload = 0
}

func (m *megasquirt_gp28) CopyFrom(o megasquirt_gp28Reader) *megasquirt_gp28 {
	m.xxx_cl_idle_targ_rpm = o.cl_idle_targ_rpm()
	m.xxx_tpsadc = o.tpsadc()
	m.Seteaeload(o.eaeload())
	m.Setafrload(o.afrload())
	return m
}

// Descriptor returns the megasquirt_gp28 descriptor.
func (m *megasquirt_gp28) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp28.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp28) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp28) cl_idle_targ_rpm() uint16 {
	return m.xxx_cl_idle_targ_rpm
}

func (m *megasquirt_gp28) Setcl_idle_targ_rpm(v uint16) *megasquirt_gp28 {
	m.xxx_cl_idle_targ_rpm = uint16(Messages().megasquirt_gp28.cl_idle_targ_rpm.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp28) tpsadc() int16 {
	return m.xxx_tpsadc
}

func (m *megasquirt_gp28) Settpsadc(v int16) *megasquirt_gp28 {
	m.xxx_tpsadc = int16(Messages().megasquirt_gp28.tpsadc.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp28) eaeload() float64 {
	return Messages().megasquirt_gp28.eaeload.ToPhysical(float64(m.xxx_eaeload))
}

func (m *megasquirt_gp28) Seteaeload(v float64) *megasquirt_gp28 {
	m.xxx_eaeload = int16(Messages().megasquirt_gp28.eaeload.FromPhysical(v))
	return m
}

func (m *megasquirt_gp28) afrload() float64 {
	return Messages().megasquirt_gp28.afrload.ToPhysical(float64(m.xxx_afrload))
}

func (m *megasquirt_gp28) Setafrload(v float64) *megasquirt_gp28 {
	m.xxx_afrload = int16(Messages().megasquirt_gp28.afrload.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp28) Frame() can.Frame {
	md := Messages().megasquirt_gp28
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.cl_idle_targ_rpm.MarshalUnsigned(&f.Data, uint64(m.xxx_cl_idle_targ_rpm))
	md.tpsadc.MarshalSigned(&f.Data, int64(m.xxx_tpsadc))
	md.eaeload.MarshalSigned(&f.Data, int64(m.xxx_eaeload))
	md.afrload.MarshalSigned(&f.Data, int64(m.xxx_afrload))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp28) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp28) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp28
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp28: expects ID 1548 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp28: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp28: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp28: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_cl_idle_targ_rpm = uint16(md.cl_idle_targ_rpm.UnmarshalUnsigned(f.Data))
	m.xxx_tpsadc = int16(md.tpsadc.UnmarshalSigned(f.Data))
	m.xxx_eaeload = int16(md.eaeload.UnmarshalSigned(f.Data))
	m.xxx_afrload = int16(md.afrload.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp29Reader provides read access to a megasquirt_gp29 message.
type megasquirt_gp29Reader interface {
	// EAEfcor1 returns the physical value of the EAEfcor1 signal.
	EAEfcor1() float64
	// EAEfcor2 returns the physical value of the EAEfcor2 signal.
	EAEfcor2() float64
	// VSS1dot returns the physical value of the VSS1dot signal.
	VSS1dot() float64
	// VSS2dot returns the physical value of the VSS2dot signal.
	VSS2dot() float64
}

// megasquirt_gp29Writer provides write access to a megasquirt_gp29 message.
type megasquirt_gp29Writer interface {
	// CopyFrom copies all values from megasquirt_gp29Reader.
	CopyFrom(megasquirt_gp29Reader) *megasquirt_gp29
	// SetEAEfcor1 sets the physical value of the EAEfcor1 signal.
	SetEAEfcor1(float64) *megasquirt_gp29
	// SetEAEfcor2 sets the physical value of the EAEfcor2 signal.
	SetEAEfcor2(float64) *megasquirt_gp29
	// SetVSS1dot sets the physical value of the VSS1dot signal.
	SetVSS1dot(float64) *megasquirt_gp29
	// SetVSS2dot sets the physical value of the VSS2dot signal.
	SetVSS2dot(float64) *megasquirt_gp29
}

type megasquirt_gp29 struct {
	xxx_EAEfcor1 uint16
	xxx_EAEfcor2 uint16
	xxx_VSS1dot  int16
	xxx_VSS2dot  int16
}

func Newmegasquirt_gp29() *megasquirt_gp29 {
	m := &megasquirt_gp29{}
	m.Reset()
	return m
}

func (m *megasquirt_gp29) Reset() {
	m.xxx_EAEfcor1 = 0
	m.xxx_EAEfcor2 = 0
	m.xxx_VSS1dot = 0
	m.xxx_VSS2dot = 0
}

func (m *megasquirt_gp29) CopyFrom(o megasquirt_gp29Reader) *megasquirt_gp29 {
	m.SetEAEfcor1(o.EAEfcor1())
	m.SetEAEfcor2(o.EAEfcor2())
	m.SetVSS1dot(o.VSS1dot())
	m.SetVSS2dot(o.VSS2dot())
	return m
}

// Descriptor returns the megasquirt_gp29 descriptor.
func (m *megasquirt_gp29) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp29.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp29) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp29) EAEfcor1() float64 {
	return Messages().megasquirt_gp29.EAEfcor1.ToPhysical(float64(m.xxx_EAEfcor1))
}

func (m *megasquirt_gp29) SetEAEfcor1(v float64) *megasquirt_gp29 {
	m.xxx_EAEfcor1 = uint16(Messages().megasquirt_gp29.EAEfcor1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp29) EAEfcor2() float64 {
	return Messages().megasquirt_gp29.EAEfcor2.ToPhysical(float64(m.xxx_EAEfcor2))
}

func (m *megasquirt_gp29) SetEAEfcor2(v float64) *megasquirt_gp29 {
	m.xxx_EAEfcor2 = uint16(Messages().megasquirt_gp29.EAEfcor2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp29) VSS1dot() float64 {
	return Messages().megasquirt_gp29.VSS1dot.ToPhysical(float64(m.xxx_VSS1dot))
}

func (m *megasquirt_gp29) SetVSS1dot(v float64) *megasquirt_gp29 {
	m.xxx_VSS1dot = int16(Messages().megasquirt_gp29.VSS1dot.FromPhysical(v))
	return m
}

func (m *megasquirt_gp29) VSS2dot() float64 {
	return Messages().megasquirt_gp29.VSS2dot.ToPhysical(float64(m.xxx_VSS2dot))
}

func (m *megasquirt_gp29) SetVSS2dot(v float64) *megasquirt_gp29 {
	m.xxx_VSS2dot = int16(Messages().megasquirt_gp29.VSS2dot.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp29) Frame() can.Frame {
	md := Messages().megasquirt_gp29
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EAEfcor1.MarshalUnsigned(&f.Data, uint64(m.xxx_EAEfcor1))
	md.EAEfcor2.MarshalUnsigned(&f.Data, uint64(m.xxx_EAEfcor2))
	md.VSS1dot.MarshalSigned(&f.Data, int64(m.xxx_VSS1dot))
	md.VSS2dot.MarshalSigned(&f.Data, int64(m.xxx_VSS2dot))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp29) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp29) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp29
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp29: expects ID 1549 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp29: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp29: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp29: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EAEfcor1 = uint16(md.EAEfcor1.UnmarshalUnsigned(f.Data))
	m.xxx_EAEfcor2 = uint16(md.EAEfcor2.UnmarshalUnsigned(f.Data))
	m.xxx_VSS1dot = int16(md.VSS1dot.UnmarshalSigned(f.Data))
	m.xxx_VSS2dot = int16(md.VSS2dot.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp30Reader provides read access to a megasquirt_gp30 message.
type megasquirt_gp30Reader interface {
	// accelx returns the physical value of the accelx signal.
	accelx() float64
	// accely returns the physical value of the accely signal.
	accely() float64
	// accelz returns the physical value of the accelz signal.
	accelz() float64
	// stream_level returns the value of the stream_level signal.
	stream_level() uint8
	// water_duty returns the value of the water_duty signal.
	water_duty() uint8
}

// megasquirt_gp30Writer provides write access to a megasquirt_gp30 message.
type megasquirt_gp30Writer interface {
	// CopyFrom copies all values from megasquirt_gp30Reader.
	CopyFrom(megasquirt_gp30Reader) *megasquirt_gp30
	// Setaccelx sets the physical value of the accelx signal.
	Setaccelx(float64) *megasquirt_gp30
	// Setaccely sets the physical value of the accely signal.
	Setaccely(float64) *megasquirt_gp30
	// Setaccelz sets the physical value of the accelz signal.
	Setaccelz(float64) *megasquirt_gp30
	// Setstream_level sets the value of the stream_level signal.
	Setstream_level(uint8) *megasquirt_gp30
	// Setwater_duty sets the value of the water_duty signal.
	Setwater_duty(uint8) *megasquirt_gp30
}

type megasquirt_gp30 struct {
	xxx_accelx       int16
	xxx_accely       int16
	xxx_accelz       int16
	xxx_stream_level uint8
	xxx_water_duty   uint8
}

func Newmegasquirt_gp30() *megasquirt_gp30 {
	m := &megasquirt_gp30{}
	m.Reset()
	return m
}

func (m *megasquirt_gp30) Reset() {
	m.xxx_accelx = 0
	m.xxx_accely = 0
	m.xxx_accelz = 0
	m.xxx_stream_level = 0
	m.xxx_water_duty = 0
}

func (m *megasquirt_gp30) CopyFrom(o megasquirt_gp30Reader) *megasquirt_gp30 {
	m.Setaccelx(o.accelx())
	m.Setaccely(o.accely())
	m.Setaccelz(o.accelz())
	m.xxx_stream_level = o.stream_level()
	m.xxx_water_duty = o.water_duty()
	return m
}

// Descriptor returns the megasquirt_gp30 descriptor.
func (m *megasquirt_gp30) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp30.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp30) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp30) accelx() float64 {
	return Messages().megasquirt_gp30.accelx.ToPhysical(float64(m.xxx_accelx))
}

func (m *megasquirt_gp30) Setaccelx(v float64) *megasquirt_gp30 {
	m.xxx_accelx = int16(Messages().megasquirt_gp30.accelx.FromPhysical(v))
	return m
}

func (m *megasquirt_gp30) accely() float64 {
	return Messages().megasquirt_gp30.accely.ToPhysical(float64(m.xxx_accely))
}

func (m *megasquirt_gp30) Setaccely(v float64) *megasquirt_gp30 {
	m.xxx_accely = int16(Messages().megasquirt_gp30.accely.FromPhysical(v))
	return m
}

func (m *megasquirt_gp30) accelz() float64 {
	return Messages().megasquirt_gp30.accelz.ToPhysical(float64(m.xxx_accelz))
}

func (m *megasquirt_gp30) Setaccelz(v float64) *megasquirt_gp30 {
	m.xxx_accelz = int16(Messages().megasquirt_gp30.accelz.FromPhysical(v))
	return m
}

func (m *megasquirt_gp30) stream_level() uint8 {
	return m.xxx_stream_level
}

func (m *megasquirt_gp30) Setstream_level(v uint8) *megasquirt_gp30 {
	m.xxx_stream_level = uint8(Messages().megasquirt_gp30.stream_level.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp30) water_duty() uint8 {
	return m.xxx_water_duty
}

func (m *megasquirt_gp30) Setwater_duty(v uint8) *megasquirt_gp30 {
	m.xxx_water_duty = uint8(Messages().megasquirt_gp30.water_duty.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp30) Frame() can.Frame {
	md := Messages().megasquirt_gp30
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.accelx.MarshalSigned(&f.Data, int64(m.xxx_accelx))
	md.accely.MarshalSigned(&f.Data, int64(m.xxx_accely))
	md.accelz.MarshalSigned(&f.Data, int64(m.xxx_accelz))
	md.stream_level.MarshalUnsigned(&f.Data, uint64(m.xxx_stream_level))
	md.water_duty.MarshalUnsigned(&f.Data, uint64(m.xxx_water_duty))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp30) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp30) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp30
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp30: expects ID 1550 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp30: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp30: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp30: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_accelx = int16(md.accelx.UnmarshalSigned(f.Data))
	m.xxx_accely = int16(md.accely.UnmarshalSigned(f.Data))
	m.xxx_accelz = int16(md.accelz.UnmarshalSigned(f.Data))
	m.xxx_stream_level = uint8(md.stream_level.UnmarshalUnsigned(f.Data))
	m.xxx_water_duty = uint8(md.water_duty.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp31Reader provides read access to a megasquirt_gp31 message.
type megasquirt_gp31Reader interface {
	// AFR1 returns the physical value of the AFR1 signal.
	AFR1() float64
	// AFR2 returns the physical value of the AFR2 signal.
	AFR2() float64
	// AFR3 returns the physical value of the AFR3 signal.
	AFR3() float64
	// AFR4 returns the physical value of the AFR4 signal.
	AFR4() float64
	// AFR5 returns the physical value of the AFR5 signal.
	AFR5() float64
	// AFR6 returns the physical value of the AFR6 signal.
	AFR6() float64
	// AFR7 returns the physical value of the AFR7 signal.
	AFR7() float64
	// AFR8 returns the physical value of the AFR8 signal.
	AFR8() float64
}

// megasquirt_gp31Writer provides write access to a megasquirt_gp31 message.
type megasquirt_gp31Writer interface {
	// CopyFrom copies all values from megasquirt_gp31Reader.
	CopyFrom(megasquirt_gp31Reader) *megasquirt_gp31
	// SetAFR1 sets the physical value of the AFR1 signal.
	SetAFR1(float64) *megasquirt_gp31
	// SetAFR2 sets the physical value of the AFR2 signal.
	SetAFR2(float64) *megasquirt_gp31
	// SetAFR3 sets the physical value of the AFR3 signal.
	SetAFR3(float64) *megasquirt_gp31
	// SetAFR4 sets the physical value of the AFR4 signal.
	SetAFR4(float64) *megasquirt_gp31
	// SetAFR5 sets the physical value of the AFR5 signal.
	SetAFR5(float64) *megasquirt_gp31
	// SetAFR6 sets the physical value of the AFR6 signal.
	SetAFR6(float64) *megasquirt_gp31
	// SetAFR7 sets the physical value of the AFR7 signal.
	SetAFR7(float64) *megasquirt_gp31
	// SetAFR8 sets the physical value of the AFR8 signal.
	SetAFR8(float64) *megasquirt_gp31
}

type megasquirt_gp31 struct {
	xxx_AFR1 uint8
	xxx_AFR2 uint8
	xxx_AFR3 uint8
	xxx_AFR4 uint8
	xxx_AFR5 uint8
	xxx_AFR6 uint8
	xxx_AFR7 uint8
	xxx_AFR8 uint8
}

func Newmegasquirt_gp31() *megasquirt_gp31 {
	m := &megasquirt_gp31{}
	m.Reset()
	return m
}

func (m *megasquirt_gp31) Reset() {
	m.xxx_AFR1 = 0
	m.xxx_AFR2 = 0
	m.xxx_AFR3 = 0
	m.xxx_AFR4 = 0
	m.xxx_AFR5 = 0
	m.xxx_AFR6 = 0
	m.xxx_AFR7 = 0
	m.xxx_AFR8 = 0
}

func (m *megasquirt_gp31) CopyFrom(o megasquirt_gp31Reader) *megasquirt_gp31 {
	m.SetAFR1(o.AFR1())
	m.SetAFR2(o.AFR2())
	m.SetAFR3(o.AFR3())
	m.SetAFR4(o.AFR4())
	m.SetAFR5(o.AFR5())
	m.SetAFR6(o.AFR6())
	m.SetAFR7(o.AFR7())
	m.SetAFR8(o.AFR8())
	return m
}

// Descriptor returns the megasquirt_gp31 descriptor.
func (m *megasquirt_gp31) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp31.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp31) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp31) AFR1() float64 {
	return Messages().megasquirt_gp31.AFR1.ToPhysical(float64(m.xxx_AFR1))
}

func (m *megasquirt_gp31) SetAFR1(v float64) *megasquirt_gp31 {
	m.xxx_AFR1 = uint8(Messages().megasquirt_gp31.AFR1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR2() float64 {
	return Messages().megasquirt_gp31.AFR2.ToPhysical(float64(m.xxx_AFR2))
}

func (m *megasquirt_gp31) SetAFR2(v float64) *megasquirt_gp31 {
	m.xxx_AFR2 = uint8(Messages().megasquirt_gp31.AFR2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR3() float64 {
	return Messages().megasquirt_gp31.AFR3.ToPhysical(float64(m.xxx_AFR3))
}

func (m *megasquirt_gp31) SetAFR3(v float64) *megasquirt_gp31 {
	m.xxx_AFR3 = uint8(Messages().megasquirt_gp31.AFR3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR4() float64 {
	return Messages().megasquirt_gp31.AFR4.ToPhysical(float64(m.xxx_AFR4))
}

func (m *megasquirt_gp31) SetAFR4(v float64) *megasquirt_gp31 {
	m.xxx_AFR4 = uint8(Messages().megasquirt_gp31.AFR4.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR5() float64 {
	return Messages().megasquirt_gp31.AFR5.ToPhysical(float64(m.xxx_AFR5))
}

func (m *megasquirt_gp31) SetAFR5(v float64) *megasquirt_gp31 {
	m.xxx_AFR5 = uint8(Messages().megasquirt_gp31.AFR5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR6() float64 {
	return Messages().megasquirt_gp31.AFR6.ToPhysical(float64(m.xxx_AFR6))
}

func (m *megasquirt_gp31) SetAFR6(v float64) *megasquirt_gp31 {
	m.xxx_AFR6 = uint8(Messages().megasquirt_gp31.AFR6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR7() float64 {
	return Messages().megasquirt_gp31.AFR7.ToPhysical(float64(m.xxx_AFR7))
}

func (m *megasquirt_gp31) SetAFR7(v float64) *megasquirt_gp31 {
	m.xxx_AFR7 = uint8(Messages().megasquirt_gp31.AFR7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp31) AFR8() float64 {
	return Messages().megasquirt_gp31.AFR8.ToPhysical(float64(m.xxx_AFR8))
}

func (m *megasquirt_gp31) SetAFR8(v float64) *megasquirt_gp31 {
	m.xxx_AFR8 = uint8(Messages().megasquirt_gp31.AFR8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp31) Frame() can.Frame {
	md := Messages().megasquirt_gp31
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.AFR1.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR1))
	md.AFR2.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR2))
	md.AFR3.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR3))
	md.AFR4.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR4))
	md.AFR5.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR5))
	md.AFR6.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR6))
	md.AFR7.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR7))
	md.AFR8.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp31) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp31) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp31
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp31: expects ID 1551 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp31: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp31: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp31: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_AFR1 = uint8(md.AFR1.UnmarshalUnsigned(f.Data))
	m.xxx_AFR2 = uint8(md.AFR2.UnmarshalUnsigned(f.Data))
	m.xxx_AFR3 = uint8(md.AFR3.UnmarshalUnsigned(f.Data))
	m.xxx_AFR4 = uint8(md.AFR4.UnmarshalUnsigned(f.Data))
	m.xxx_AFR5 = uint8(md.AFR5.UnmarshalUnsigned(f.Data))
	m.xxx_AFR6 = uint8(md.AFR6.UnmarshalUnsigned(f.Data))
	m.xxx_AFR7 = uint8(md.AFR7.UnmarshalUnsigned(f.Data))
	m.xxx_AFR8 = uint8(md.AFR8.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp32Reader provides read access to a megasquirt_gp32 message.
type megasquirt_gp32Reader interface {
	// AFR9 returns the physical value of the AFR9 signal.
	AFR9() float64
	// AFR10 returns the physical value of the AFR10 signal.
	AFR10() float64
	// AFR11 returns the physical value of the AFR11 signal.
	AFR11() float64
	// AFR12 returns the physical value of the AFR12 signal.
	AFR12() float64
	// AFR13 returns the physical value of the AFR13 signal.
	AFR13() float64
	// AFR14 returns the physical value of the AFR14 signal.
	AFR14() float64
	// AFR15 returns the physical value of the AFR15 signal.
	AFR15() float64
	// AFR16 returns the physical value of the AFR16 signal.
	AFR16() float64
}

// megasquirt_gp32Writer provides write access to a megasquirt_gp32 message.
type megasquirt_gp32Writer interface {
	// CopyFrom copies all values from megasquirt_gp32Reader.
	CopyFrom(megasquirt_gp32Reader) *megasquirt_gp32
	// SetAFR9 sets the physical value of the AFR9 signal.
	SetAFR9(float64) *megasquirt_gp32
	// SetAFR10 sets the physical value of the AFR10 signal.
	SetAFR10(float64) *megasquirt_gp32
	// SetAFR11 sets the physical value of the AFR11 signal.
	SetAFR11(float64) *megasquirt_gp32
	// SetAFR12 sets the physical value of the AFR12 signal.
	SetAFR12(float64) *megasquirt_gp32
	// SetAFR13 sets the physical value of the AFR13 signal.
	SetAFR13(float64) *megasquirt_gp32
	// SetAFR14 sets the physical value of the AFR14 signal.
	SetAFR14(float64) *megasquirt_gp32
	// SetAFR15 sets the physical value of the AFR15 signal.
	SetAFR15(float64) *megasquirt_gp32
	// SetAFR16 sets the physical value of the AFR16 signal.
	SetAFR16(float64) *megasquirt_gp32
}

type megasquirt_gp32 struct {
	xxx_AFR9  uint8
	xxx_AFR10 uint8
	xxx_AFR11 uint8
	xxx_AFR12 uint8
	xxx_AFR13 uint8
	xxx_AFR14 uint8
	xxx_AFR15 uint8
	xxx_AFR16 uint8
}

func Newmegasquirt_gp32() *megasquirt_gp32 {
	m := &megasquirt_gp32{}
	m.Reset()
	return m
}

func (m *megasquirt_gp32) Reset() {
	m.xxx_AFR9 = 0
	m.xxx_AFR10 = 0
	m.xxx_AFR11 = 0
	m.xxx_AFR12 = 0
	m.xxx_AFR13 = 0
	m.xxx_AFR14 = 0
	m.xxx_AFR15 = 0
	m.xxx_AFR16 = 0
}

func (m *megasquirt_gp32) CopyFrom(o megasquirt_gp32Reader) *megasquirt_gp32 {
	m.SetAFR9(o.AFR9())
	m.SetAFR10(o.AFR10())
	m.SetAFR11(o.AFR11())
	m.SetAFR12(o.AFR12())
	m.SetAFR13(o.AFR13())
	m.SetAFR14(o.AFR14())
	m.SetAFR15(o.AFR15())
	m.SetAFR16(o.AFR16())
	return m
}

// Descriptor returns the megasquirt_gp32 descriptor.
func (m *megasquirt_gp32) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp32.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp32) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp32) AFR9() float64 {
	return Messages().megasquirt_gp32.AFR9.ToPhysical(float64(m.xxx_AFR9))
}

func (m *megasquirt_gp32) SetAFR9(v float64) *megasquirt_gp32 {
	m.xxx_AFR9 = uint8(Messages().megasquirt_gp32.AFR9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR10() float64 {
	return Messages().megasquirt_gp32.AFR10.ToPhysical(float64(m.xxx_AFR10))
}

func (m *megasquirt_gp32) SetAFR10(v float64) *megasquirt_gp32 {
	m.xxx_AFR10 = uint8(Messages().megasquirt_gp32.AFR10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR11() float64 {
	return Messages().megasquirt_gp32.AFR11.ToPhysical(float64(m.xxx_AFR11))
}

func (m *megasquirt_gp32) SetAFR11(v float64) *megasquirt_gp32 {
	m.xxx_AFR11 = uint8(Messages().megasquirt_gp32.AFR11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR12() float64 {
	return Messages().megasquirt_gp32.AFR12.ToPhysical(float64(m.xxx_AFR12))
}

func (m *megasquirt_gp32) SetAFR12(v float64) *megasquirt_gp32 {
	m.xxx_AFR12 = uint8(Messages().megasquirt_gp32.AFR12.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR13() float64 {
	return Messages().megasquirt_gp32.AFR13.ToPhysical(float64(m.xxx_AFR13))
}

func (m *megasquirt_gp32) SetAFR13(v float64) *megasquirt_gp32 {
	m.xxx_AFR13 = uint8(Messages().megasquirt_gp32.AFR13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR14() float64 {
	return Messages().megasquirt_gp32.AFR14.ToPhysical(float64(m.xxx_AFR14))
}

func (m *megasquirt_gp32) SetAFR14(v float64) *megasquirt_gp32 {
	m.xxx_AFR14 = uint8(Messages().megasquirt_gp32.AFR14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR15() float64 {
	return Messages().megasquirt_gp32.AFR15.ToPhysical(float64(m.xxx_AFR15))
}

func (m *megasquirt_gp32) SetAFR15(v float64) *megasquirt_gp32 {
	m.xxx_AFR15 = uint8(Messages().megasquirt_gp32.AFR15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp32) AFR16() float64 {
	return Messages().megasquirt_gp32.AFR16.ToPhysical(float64(m.xxx_AFR16))
}

func (m *megasquirt_gp32) SetAFR16(v float64) *megasquirt_gp32 {
	m.xxx_AFR16 = uint8(Messages().megasquirt_gp32.AFR16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp32) Frame() can.Frame {
	md := Messages().megasquirt_gp32
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.AFR9.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR9))
	md.AFR10.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR10))
	md.AFR11.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR11))
	md.AFR12.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR12))
	md.AFR13.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR13))
	md.AFR14.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR14))
	md.AFR15.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR15))
	md.AFR16.MarshalUnsigned(&f.Data, uint64(m.xxx_AFR16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp32) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp32) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp32
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp32: expects ID 1552 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp32: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp32: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp32: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_AFR9 = uint8(md.AFR9.UnmarshalUnsigned(f.Data))
	m.xxx_AFR10 = uint8(md.AFR10.UnmarshalUnsigned(f.Data))
	m.xxx_AFR11 = uint8(md.AFR11.UnmarshalUnsigned(f.Data))
	m.xxx_AFR12 = uint8(md.AFR12.UnmarshalUnsigned(f.Data))
	m.xxx_AFR13 = uint8(md.AFR13.UnmarshalUnsigned(f.Data))
	m.xxx_AFR14 = uint8(md.AFR14.UnmarshalUnsigned(f.Data))
	m.xxx_AFR15 = uint8(md.AFR15.UnmarshalUnsigned(f.Data))
	m.xxx_AFR16 = uint8(md.AFR16.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp33Reader provides read access to a megasquirt_gp33 message.
type megasquirt_gp33Reader interface {
	// duty_pwm1 returns the value of the duty_pwm1 signal.
	duty_pwm1() uint8
	// duty_pwm2 returns the value of the duty_pwm2 signal.
	duty_pwm2() uint8
	// duty_pwm3 returns the value of the duty_pwm3 signal.
	duty_pwm3() uint8
	// duty_pwm4 returns the value of the duty_pwm4 signal.
	duty_pwm4() uint8
	// duty_pwm5 returns the value of the duty_pwm5 signal.
	duty_pwm5() uint8
	// duty_pwm6 returns the value of the duty_pwm6 signal.
	duty_pwm6() uint8
	// gear returns the value of the gear signal.
	gear() uint8
	// status8 returns the value of the status8 signal.
	status8() uint8
}

// megasquirt_gp33Writer provides write access to a megasquirt_gp33 message.
type megasquirt_gp33Writer interface {
	// CopyFrom copies all values from megasquirt_gp33Reader.
	CopyFrom(megasquirt_gp33Reader) *megasquirt_gp33
	// Setduty_pwm1 sets the value of the duty_pwm1 signal.
	Setduty_pwm1(uint8) *megasquirt_gp33
	// Setduty_pwm2 sets the value of the duty_pwm2 signal.
	Setduty_pwm2(uint8) *megasquirt_gp33
	// Setduty_pwm3 sets the value of the duty_pwm3 signal.
	Setduty_pwm3(uint8) *megasquirt_gp33
	// Setduty_pwm4 sets the value of the duty_pwm4 signal.
	Setduty_pwm4(uint8) *megasquirt_gp33
	// Setduty_pwm5 sets the value of the duty_pwm5 signal.
	Setduty_pwm5(uint8) *megasquirt_gp33
	// Setduty_pwm6 sets the value of the duty_pwm6 signal.
	Setduty_pwm6(uint8) *megasquirt_gp33
	// Setgear sets the value of the gear signal.
	Setgear(uint8) *megasquirt_gp33
	// Setstatus8 sets the value of the status8 signal.
	Setstatus8(uint8) *megasquirt_gp33
}

type megasquirt_gp33 struct {
	xxx_duty_pwm1 uint8
	xxx_duty_pwm2 uint8
	xxx_duty_pwm3 uint8
	xxx_duty_pwm4 uint8
	xxx_duty_pwm5 uint8
	xxx_duty_pwm6 uint8
	xxx_gear      uint8
	xxx_status8   uint8
}

func Newmegasquirt_gp33() *megasquirt_gp33 {
	m := &megasquirt_gp33{}
	m.Reset()
	return m
}

func (m *megasquirt_gp33) Reset() {
	m.xxx_duty_pwm1 = 0
	m.xxx_duty_pwm2 = 0
	m.xxx_duty_pwm3 = 0
	m.xxx_duty_pwm4 = 0
	m.xxx_duty_pwm5 = 0
	m.xxx_duty_pwm6 = 0
	m.xxx_gear = 0
	m.xxx_status8 = 0
}

func (m *megasquirt_gp33) CopyFrom(o megasquirt_gp33Reader) *megasquirt_gp33 {
	m.xxx_duty_pwm1 = o.duty_pwm1()
	m.xxx_duty_pwm2 = o.duty_pwm2()
	m.xxx_duty_pwm3 = o.duty_pwm3()
	m.xxx_duty_pwm4 = o.duty_pwm4()
	m.xxx_duty_pwm5 = o.duty_pwm5()
	m.xxx_duty_pwm6 = o.duty_pwm6()
	m.xxx_gear = o.gear()
	m.xxx_status8 = o.status8()
	return m
}

// Descriptor returns the megasquirt_gp33 descriptor.
func (m *megasquirt_gp33) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp33.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp33) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp33) duty_pwm1() uint8 {
	return m.xxx_duty_pwm1
}

func (m *megasquirt_gp33) Setduty_pwm1(v uint8) *megasquirt_gp33 {
	m.xxx_duty_pwm1 = uint8(Messages().megasquirt_gp33.duty_pwm1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) duty_pwm2() uint8 {
	return m.xxx_duty_pwm2
}

func (m *megasquirt_gp33) Setduty_pwm2(v uint8) *megasquirt_gp33 {
	m.xxx_duty_pwm2 = uint8(Messages().megasquirt_gp33.duty_pwm2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) duty_pwm3() uint8 {
	return m.xxx_duty_pwm3
}

func (m *megasquirt_gp33) Setduty_pwm3(v uint8) *megasquirt_gp33 {
	m.xxx_duty_pwm3 = uint8(Messages().megasquirt_gp33.duty_pwm3.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) duty_pwm4() uint8 {
	return m.xxx_duty_pwm4
}

func (m *megasquirt_gp33) Setduty_pwm4(v uint8) *megasquirt_gp33 {
	m.xxx_duty_pwm4 = uint8(Messages().megasquirt_gp33.duty_pwm4.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) duty_pwm5() uint8 {
	return m.xxx_duty_pwm5
}

func (m *megasquirt_gp33) Setduty_pwm5(v uint8) *megasquirt_gp33 {
	m.xxx_duty_pwm5 = uint8(Messages().megasquirt_gp33.duty_pwm5.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) duty_pwm6() uint8 {
	return m.xxx_duty_pwm6
}

func (m *megasquirt_gp33) Setduty_pwm6(v uint8) *megasquirt_gp33 {
	m.xxx_duty_pwm6 = uint8(Messages().megasquirt_gp33.duty_pwm6.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) gear() uint8 {
	return m.xxx_gear
}

func (m *megasquirt_gp33) Setgear(v uint8) *megasquirt_gp33 {
	m.xxx_gear = uint8(Messages().megasquirt_gp33.gear.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp33) status8() uint8 {
	return m.xxx_status8
}

func (m *megasquirt_gp33) Setstatus8(v uint8) *megasquirt_gp33 {
	m.xxx_status8 = uint8(Messages().megasquirt_gp33.status8.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp33) Frame() can.Frame {
	md := Messages().megasquirt_gp33
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.duty_pwm1.MarshalUnsigned(&f.Data, uint64(m.xxx_duty_pwm1))
	md.duty_pwm2.MarshalUnsigned(&f.Data, uint64(m.xxx_duty_pwm2))
	md.duty_pwm3.MarshalUnsigned(&f.Data, uint64(m.xxx_duty_pwm3))
	md.duty_pwm4.MarshalUnsigned(&f.Data, uint64(m.xxx_duty_pwm4))
	md.duty_pwm5.MarshalUnsigned(&f.Data, uint64(m.xxx_duty_pwm5))
	md.duty_pwm6.MarshalUnsigned(&f.Data, uint64(m.xxx_duty_pwm6))
	md.gear.MarshalUnsigned(&f.Data, uint64(m.xxx_gear))
	md.status8.MarshalUnsigned(&f.Data, uint64(m.xxx_status8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp33) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp33) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp33
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp33: expects ID 1553 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp33: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp33: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp33: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_duty_pwm1 = uint8(md.duty_pwm1.UnmarshalUnsigned(f.Data))
	m.xxx_duty_pwm2 = uint8(md.duty_pwm2.UnmarshalUnsigned(f.Data))
	m.xxx_duty_pwm3 = uint8(md.duty_pwm3.UnmarshalUnsigned(f.Data))
	m.xxx_duty_pwm4 = uint8(md.duty_pwm4.UnmarshalUnsigned(f.Data))
	m.xxx_duty_pwm5 = uint8(md.duty_pwm5.UnmarshalUnsigned(f.Data))
	m.xxx_duty_pwm6 = uint8(md.duty_pwm6.UnmarshalUnsigned(f.Data))
	m.xxx_gear = uint8(md.gear.UnmarshalUnsigned(f.Data))
	m.xxx_status8 = uint8(md.status8.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp34Reader provides read access to a megasquirt_gp34 message.
type megasquirt_gp34Reader interface {
	// EGOv1 returns the physical value of the EGOv1 signal.
	EGOv1() float64
	// EGOv2 returns the physical value of the EGOv2 signal.
	EGOv2() float64
	// EGOv3 returns the physical value of the EGOv3 signal.
	EGOv3() float64
	// EGOv4 returns the physical value of the EGOv4 signal.
	EGOv4() float64
}

// megasquirt_gp34Writer provides write access to a megasquirt_gp34 message.
type megasquirt_gp34Writer interface {
	// CopyFrom copies all values from megasquirt_gp34Reader.
	CopyFrom(megasquirt_gp34Reader) *megasquirt_gp34
	// SetEGOv1 sets the physical value of the EGOv1 signal.
	SetEGOv1(float64) *megasquirt_gp34
	// SetEGOv2 sets the physical value of the EGOv2 signal.
	SetEGOv2(float64) *megasquirt_gp34
	// SetEGOv3 sets the physical value of the EGOv3 signal.
	SetEGOv3(float64) *megasquirt_gp34
	// SetEGOv4 sets the physical value of the EGOv4 signal.
	SetEGOv4(float64) *megasquirt_gp34
}

type megasquirt_gp34 struct {
	xxx_EGOv1 int16
	xxx_EGOv2 int16
	xxx_EGOv3 int16
	xxx_EGOv4 int16
}

func Newmegasquirt_gp34() *megasquirt_gp34 {
	m := &megasquirt_gp34{}
	m.Reset()
	return m
}

func (m *megasquirt_gp34) Reset() {
	m.xxx_EGOv1 = 0
	m.xxx_EGOv2 = 0
	m.xxx_EGOv3 = 0
	m.xxx_EGOv4 = 0
}

func (m *megasquirt_gp34) CopyFrom(o megasquirt_gp34Reader) *megasquirt_gp34 {
	m.SetEGOv1(o.EGOv1())
	m.SetEGOv2(o.EGOv2())
	m.SetEGOv3(o.EGOv3())
	m.SetEGOv4(o.EGOv4())
	return m
}

// Descriptor returns the megasquirt_gp34 descriptor.
func (m *megasquirt_gp34) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp34.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp34) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp34) EGOv1() float64 {
	return Messages().megasquirt_gp34.EGOv1.ToPhysical(float64(m.xxx_EGOv1))
}

func (m *megasquirt_gp34) SetEGOv1(v float64) *megasquirt_gp34 {
	m.xxx_EGOv1 = int16(Messages().megasquirt_gp34.EGOv1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp34) EGOv2() float64 {
	return Messages().megasquirt_gp34.EGOv2.ToPhysical(float64(m.xxx_EGOv2))
}

func (m *megasquirt_gp34) SetEGOv2(v float64) *megasquirt_gp34 {
	m.xxx_EGOv2 = int16(Messages().megasquirt_gp34.EGOv2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp34) EGOv3() float64 {
	return Messages().megasquirt_gp34.EGOv3.ToPhysical(float64(m.xxx_EGOv3))
}

func (m *megasquirt_gp34) SetEGOv3(v float64) *megasquirt_gp34 {
	m.xxx_EGOv3 = int16(Messages().megasquirt_gp34.EGOv3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp34) EGOv4() float64 {
	return Messages().megasquirt_gp34.EGOv4.ToPhysical(float64(m.xxx_EGOv4))
}

func (m *megasquirt_gp34) SetEGOv4(v float64) *megasquirt_gp34 {
	m.xxx_EGOv4 = int16(Messages().megasquirt_gp34.EGOv4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp34) Frame() can.Frame {
	md := Messages().megasquirt_gp34
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOv1.MarshalSigned(&f.Data, int64(m.xxx_EGOv1))
	md.EGOv2.MarshalSigned(&f.Data, int64(m.xxx_EGOv2))
	md.EGOv3.MarshalSigned(&f.Data, int64(m.xxx_EGOv3))
	md.EGOv4.MarshalSigned(&f.Data, int64(m.xxx_EGOv4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp34) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp34) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp34
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp34: expects ID 1554 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp34: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp34: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp34: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOv1 = int16(md.EGOv1.UnmarshalSigned(f.Data))
	m.xxx_EGOv2 = int16(md.EGOv2.UnmarshalSigned(f.Data))
	m.xxx_EGOv3 = int16(md.EGOv3.UnmarshalSigned(f.Data))
	m.xxx_EGOv4 = int16(md.EGOv4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp35Reader provides read access to a megasquirt_gp35 message.
type megasquirt_gp35Reader interface {
	// EGOv5 returns the physical value of the EGOv5 signal.
	EGOv5() float64
	// EGOv6 returns the physical value of the EGOv6 signal.
	EGOv6() float64
	// EGOv7 returns the physical value of the EGOv7 signal.
	EGOv7() float64
	// EGOv8 returns the physical value of the EGOv8 signal.
	EGOv8() float64
}

// megasquirt_gp35Writer provides write access to a megasquirt_gp35 message.
type megasquirt_gp35Writer interface {
	// CopyFrom copies all values from megasquirt_gp35Reader.
	CopyFrom(megasquirt_gp35Reader) *megasquirt_gp35
	// SetEGOv5 sets the physical value of the EGOv5 signal.
	SetEGOv5(float64) *megasquirt_gp35
	// SetEGOv6 sets the physical value of the EGOv6 signal.
	SetEGOv6(float64) *megasquirt_gp35
	// SetEGOv7 sets the physical value of the EGOv7 signal.
	SetEGOv7(float64) *megasquirt_gp35
	// SetEGOv8 sets the physical value of the EGOv8 signal.
	SetEGOv8(float64) *megasquirt_gp35
}

type megasquirt_gp35 struct {
	xxx_EGOv5 int16
	xxx_EGOv6 int16
	xxx_EGOv7 int16
	xxx_EGOv8 int16
}

func Newmegasquirt_gp35() *megasquirt_gp35 {
	m := &megasquirt_gp35{}
	m.Reset()
	return m
}

func (m *megasquirt_gp35) Reset() {
	m.xxx_EGOv5 = 0
	m.xxx_EGOv6 = 0
	m.xxx_EGOv7 = 0
	m.xxx_EGOv8 = 0
}

func (m *megasquirt_gp35) CopyFrom(o megasquirt_gp35Reader) *megasquirt_gp35 {
	m.SetEGOv5(o.EGOv5())
	m.SetEGOv6(o.EGOv6())
	m.SetEGOv7(o.EGOv7())
	m.SetEGOv8(o.EGOv8())
	return m
}

// Descriptor returns the megasquirt_gp35 descriptor.
func (m *megasquirt_gp35) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp35.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp35) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp35) EGOv5() float64 {
	return Messages().megasquirt_gp35.EGOv5.ToPhysical(float64(m.xxx_EGOv5))
}

func (m *megasquirt_gp35) SetEGOv5(v float64) *megasquirt_gp35 {
	m.xxx_EGOv5 = int16(Messages().megasquirt_gp35.EGOv5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp35) EGOv6() float64 {
	return Messages().megasquirt_gp35.EGOv6.ToPhysical(float64(m.xxx_EGOv6))
}

func (m *megasquirt_gp35) SetEGOv6(v float64) *megasquirt_gp35 {
	m.xxx_EGOv6 = int16(Messages().megasquirt_gp35.EGOv6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp35) EGOv7() float64 {
	return Messages().megasquirt_gp35.EGOv7.ToPhysical(float64(m.xxx_EGOv7))
}

func (m *megasquirt_gp35) SetEGOv7(v float64) *megasquirt_gp35 {
	m.xxx_EGOv7 = int16(Messages().megasquirt_gp35.EGOv7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp35) EGOv8() float64 {
	return Messages().megasquirt_gp35.EGOv8.ToPhysical(float64(m.xxx_EGOv8))
}

func (m *megasquirt_gp35) SetEGOv8(v float64) *megasquirt_gp35 {
	m.xxx_EGOv8 = int16(Messages().megasquirt_gp35.EGOv8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp35) Frame() can.Frame {
	md := Messages().megasquirt_gp35
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOv5.MarshalSigned(&f.Data, int64(m.xxx_EGOv5))
	md.EGOv6.MarshalSigned(&f.Data, int64(m.xxx_EGOv6))
	md.EGOv7.MarshalSigned(&f.Data, int64(m.xxx_EGOv7))
	md.EGOv8.MarshalSigned(&f.Data, int64(m.xxx_EGOv8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp35) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp35) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp35
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp35: expects ID 1555 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp35: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp35: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp35: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOv5 = int16(md.EGOv5.UnmarshalSigned(f.Data))
	m.xxx_EGOv6 = int16(md.EGOv6.UnmarshalSigned(f.Data))
	m.xxx_EGOv7 = int16(md.EGOv7.UnmarshalSigned(f.Data))
	m.xxx_EGOv8 = int16(md.EGOv8.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp36Reader provides read access to a megasquirt_gp36 message.
type megasquirt_gp36Reader interface {
	// EGOv9 returns the physical value of the EGOv9 signal.
	EGOv9() float64
	// EGOv10 returns the physical value of the EGOv10 signal.
	EGOv10() float64
	// EGOv11 returns the physical value of the EGOv11 signal.
	EGOv11() float64
	// EGOv12 returns the physical value of the EGOv12 signal.
	EGOv12() float64
}

// megasquirt_gp36Writer provides write access to a megasquirt_gp36 message.
type megasquirt_gp36Writer interface {
	// CopyFrom copies all values from megasquirt_gp36Reader.
	CopyFrom(megasquirt_gp36Reader) *megasquirt_gp36
	// SetEGOv9 sets the physical value of the EGOv9 signal.
	SetEGOv9(float64) *megasquirt_gp36
	// SetEGOv10 sets the physical value of the EGOv10 signal.
	SetEGOv10(float64) *megasquirt_gp36
	// SetEGOv11 sets the physical value of the EGOv11 signal.
	SetEGOv11(float64) *megasquirt_gp36
	// SetEGOv12 sets the physical value of the EGOv12 signal.
	SetEGOv12(float64) *megasquirt_gp36
}

type megasquirt_gp36 struct {
	xxx_EGOv9  int16
	xxx_EGOv10 int16
	xxx_EGOv11 int16
	xxx_EGOv12 int16
}

func Newmegasquirt_gp36() *megasquirt_gp36 {
	m := &megasquirt_gp36{}
	m.Reset()
	return m
}

func (m *megasquirt_gp36) Reset() {
	m.xxx_EGOv9 = 0
	m.xxx_EGOv10 = 0
	m.xxx_EGOv11 = 0
	m.xxx_EGOv12 = 0
}

func (m *megasquirt_gp36) CopyFrom(o megasquirt_gp36Reader) *megasquirt_gp36 {
	m.SetEGOv9(o.EGOv9())
	m.SetEGOv10(o.EGOv10())
	m.SetEGOv11(o.EGOv11())
	m.SetEGOv12(o.EGOv12())
	return m
}

// Descriptor returns the megasquirt_gp36 descriptor.
func (m *megasquirt_gp36) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp36.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp36) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp36) EGOv9() float64 {
	return Messages().megasquirt_gp36.EGOv9.ToPhysical(float64(m.xxx_EGOv9))
}

func (m *megasquirt_gp36) SetEGOv9(v float64) *megasquirt_gp36 {
	m.xxx_EGOv9 = int16(Messages().megasquirt_gp36.EGOv9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp36) EGOv10() float64 {
	return Messages().megasquirt_gp36.EGOv10.ToPhysical(float64(m.xxx_EGOv10))
}

func (m *megasquirt_gp36) SetEGOv10(v float64) *megasquirt_gp36 {
	m.xxx_EGOv10 = int16(Messages().megasquirt_gp36.EGOv10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp36) EGOv11() float64 {
	return Messages().megasquirt_gp36.EGOv11.ToPhysical(float64(m.xxx_EGOv11))
}

func (m *megasquirt_gp36) SetEGOv11(v float64) *megasquirt_gp36 {
	m.xxx_EGOv11 = int16(Messages().megasquirt_gp36.EGOv11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp36) EGOv12() float64 {
	return Messages().megasquirt_gp36.EGOv12.ToPhysical(float64(m.xxx_EGOv12))
}

func (m *megasquirt_gp36) SetEGOv12(v float64) *megasquirt_gp36 {
	m.xxx_EGOv12 = int16(Messages().megasquirt_gp36.EGOv12.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp36) Frame() can.Frame {
	md := Messages().megasquirt_gp36
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOv9.MarshalSigned(&f.Data, int64(m.xxx_EGOv9))
	md.EGOv10.MarshalSigned(&f.Data, int64(m.xxx_EGOv10))
	md.EGOv11.MarshalSigned(&f.Data, int64(m.xxx_EGOv11))
	md.EGOv12.MarshalSigned(&f.Data, int64(m.xxx_EGOv12))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp36) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp36) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp36
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp36: expects ID 1556 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp36: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp36: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp36: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOv9 = int16(md.EGOv9.UnmarshalSigned(f.Data))
	m.xxx_EGOv10 = int16(md.EGOv10.UnmarshalSigned(f.Data))
	m.xxx_EGOv11 = int16(md.EGOv11.UnmarshalSigned(f.Data))
	m.xxx_EGOv12 = int16(md.EGOv12.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp37Reader provides read access to a megasquirt_gp37 message.
type megasquirt_gp37Reader interface {
	// EGOv13 returns the physical value of the EGOv13 signal.
	EGOv13() float64
	// EGOv14 returns the physical value of the EGOv14 signal.
	EGOv14() float64
	// EGOv15 returns the physical value of the EGOv15 signal.
	EGOv15() float64
	// EGOv16 returns the physical value of the EGOv16 signal.
	EGOv16() float64
}

// megasquirt_gp37Writer provides write access to a megasquirt_gp37 message.
type megasquirt_gp37Writer interface {
	// CopyFrom copies all values from megasquirt_gp37Reader.
	CopyFrom(megasquirt_gp37Reader) *megasquirt_gp37
	// SetEGOv13 sets the physical value of the EGOv13 signal.
	SetEGOv13(float64) *megasquirt_gp37
	// SetEGOv14 sets the physical value of the EGOv14 signal.
	SetEGOv14(float64) *megasquirt_gp37
	// SetEGOv15 sets the physical value of the EGOv15 signal.
	SetEGOv15(float64) *megasquirt_gp37
	// SetEGOv16 sets the physical value of the EGOv16 signal.
	SetEGOv16(float64) *megasquirt_gp37
}

type megasquirt_gp37 struct {
	xxx_EGOv13 int16
	xxx_EGOv14 int16
	xxx_EGOv15 int16
	xxx_EGOv16 int16
}

func Newmegasquirt_gp37() *megasquirt_gp37 {
	m := &megasquirt_gp37{}
	m.Reset()
	return m
}

func (m *megasquirt_gp37) Reset() {
	m.xxx_EGOv13 = 0
	m.xxx_EGOv14 = 0
	m.xxx_EGOv15 = 0
	m.xxx_EGOv16 = 0
}

func (m *megasquirt_gp37) CopyFrom(o megasquirt_gp37Reader) *megasquirt_gp37 {
	m.SetEGOv13(o.EGOv13())
	m.SetEGOv14(o.EGOv14())
	m.SetEGOv15(o.EGOv15())
	m.SetEGOv16(o.EGOv16())
	return m
}

// Descriptor returns the megasquirt_gp37 descriptor.
func (m *megasquirt_gp37) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp37.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp37) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp37) EGOv13() float64 {
	return Messages().megasquirt_gp37.EGOv13.ToPhysical(float64(m.xxx_EGOv13))
}

func (m *megasquirt_gp37) SetEGOv13(v float64) *megasquirt_gp37 {
	m.xxx_EGOv13 = int16(Messages().megasquirt_gp37.EGOv13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp37) EGOv14() float64 {
	return Messages().megasquirt_gp37.EGOv14.ToPhysical(float64(m.xxx_EGOv14))
}

func (m *megasquirt_gp37) SetEGOv14(v float64) *megasquirt_gp37 {
	m.xxx_EGOv14 = int16(Messages().megasquirt_gp37.EGOv14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp37) EGOv15() float64 {
	return Messages().megasquirt_gp37.EGOv15.ToPhysical(float64(m.xxx_EGOv15))
}

func (m *megasquirt_gp37) SetEGOv15(v float64) *megasquirt_gp37 {
	m.xxx_EGOv15 = int16(Messages().megasquirt_gp37.EGOv15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp37) EGOv16() float64 {
	return Messages().megasquirt_gp37.EGOv16.ToPhysical(float64(m.xxx_EGOv16))
}

func (m *megasquirt_gp37) SetEGOv16(v float64) *megasquirt_gp37 {
	m.xxx_EGOv16 = int16(Messages().megasquirt_gp37.EGOv16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp37) Frame() can.Frame {
	md := Messages().megasquirt_gp37
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOv13.MarshalSigned(&f.Data, int64(m.xxx_EGOv13))
	md.EGOv14.MarshalSigned(&f.Data, int64(m.xxx_EGOv14))
	md.EGOv15.MarshalSigned(&f.Data, int64(m.xxx_EGOv15))
	md.EGOv16.MarshalSigned(&f.Data, int64(m.xxx_EGOv16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp37) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp37) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp37
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp37: expects ID 1557 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp37: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp37: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp37: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOv13 = int16(md.EGOv13.UnmarshalSigned(f.Data))
	m.xxx_EGOv14 = int16(md.EGOv14.UnmarshalSigned(f.Data))
	m.xxx_EGOv15 = int16(md.EGOv15.UnmarshalSigned(f.Data))
	m.xxx_EGOv16 = int16(md.EGOv16.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp38Reader provides read access to a megasquirt_gp38 message.
type megasquirt_gp38Reader interface {
	// EGOcor1 returns the physical value of the EGOcor1 signal.
	EGOcor1() float64
	// EGOcor2 returns the physical value of the EGOcor2 signal.
	EGOcor2() float64
	// EGOcor3 returns the physical value of the EGOcor3 signal.
	EGOcor3() float64
	// EGOcor4 returns the physical value of the EGOcor4 signal.
	EGOcor4() float64
}

// megasquirt_gp38Writer provides write access to a megasquirt_gp38 message.
type megasquirt_gp38Writer interface {
	// CopyFrom copies all values from megasquirt_gp38Reader.
	CopyFrom(megasquirt_gp38Reader) *megasquirt_gp38
	// SetEGOcor1 sets the physical value of the EGOcor1 signal.
	SetEGOcor1(float64) *megasquirt_gp38
	// SetEGOcor2 sets the physical value of the EGOcor2 signal.
	SetEGOcor2(float64) *megasquirt_gp38
	// SetEGOcor3 sets the physical value of the EGOcor3 signal.
	SetEGOcor3(float64) *megasquirt_gp38
	// SetEGOcor4 sets the physical value of the EGOcor4 signal.
	SetEGOcor4(float64) *megasquirt_gp38
}

type megasquirt_gp38 struct {
	xxx_EGOcor1 int16
	xxx_EGOcor2 int16
	xxx_EGOcor3 int16
	xxx_EGOcor4 int16
}

func Newmegasquirt_gp38() *megasquirt_gp38 {
	m := &megasquirt_gp38{}
	m.Reset()
	return m
}

func (m *megasquirt_gp38) Reset() {
	m.xxx_EGOcor1 = 0
	m.xxx_EGOcor2 = 0
	m.xxx_EGOcor3 = 0
	m.xxx_EGOcor4 = 0
}

func (m *megasquirt_gp38) CopyFrom(o megasquirt_gp38Reader) *megasquirt_gp38 {
	m.SetEGOcor1(o.EGOcor1())
	m.SetEGOcor2(o.EGOcor2())
	m.SetEGOcor3(o.EGOcor3())
	m.SetEGOcor4(o.EGOcor4())
	return m
}

// Descriptor returns the megasquirt_gp38 descriptor.
func (m *megasquirt_gp38) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp38.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp38) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp38) EGOcor1() float64 {
	return Messages().megasquirt_gp38.EGOcor1.ToPhysical(float64(m.xxx_EGOcor1))
}

func (m *megasquirt_gp38) SetEGOcor1(v float64) *megasquirt_gp38 {
	m.xxx_EGOcor1 = int16(Messages().megasquirt_gp38.EGOcor1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp38) EGOcor2() float64 {
	return Messages().megasquirt_gp38.EGOcor2.ToPhysical(float64(m.xxx_EGOcor2))
}

func (m *megasquirt_gp38) SetEGOcor2(v float64) *megasquirt_gp38 {
	m.xxx_EGOcor2 = int16(Messages().megasquirt_gp38.EGOcor2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp38) EGOcor3() float64 {
	return Messages().megasquirt_gp38.EGOcor3.ToPhysical(float64(m.xxx_EGOcor3))
}

func (m *megasquirt_gp38) SetEGOcor3(v float64) *megasquirt_gp38 {
	m.xxx_EGOcor3 = int16(Messages().megasquirt_gp38.EGOcor3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp38) EGOcor4() float64 {
	return Messages().megasquirt_gp38.EGOcor4.ToPhysical(float64(m.xxx_EGOcor4))
}

func (m *megasquirt_gp38) SetEGOcor4(v float64) *megasquirt_gp38 {
	m.xxx_EGOcor4 = int16(Messages().megasquirt_gp38.EGOcor4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp38) Frame() can.Frame {
	md := Messages().megasquirt_gp38
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOcor1.MarshalSigned(&f.Data, int64(m.xxx_EGOcor1))
	md.EGOcor2.MarshalSigned(&f.Data, int64(m.xxx_EGOcor2))
	md.EGOcor3.MarshalSigned(&f.Data, int64(m.xxx_EGOcor3))
	md.EGOcor4.MarshalSigned(&f.Data, int64(m.xxx_EGOcor4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp38) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp38) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp38
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp38: expects ID 1558 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp38: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp38: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp38: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOcor1 = int16(md.EGOcor1.UnmarshalSigned(f.Data))
	m.xxx_EGOcor2 = int16(md.EGOcor2.UnmarshalSigned(f.Data))
	m.xxx_EGOcor3 = int16(md.EGOcor3.UnmarshalSigned(f.Data))
	m.xxx_EGOcor4 = int16(md.EGOcor4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp39Reader provides read access to a megasquirt_gp39 message.
type megasquirt_gp39Reader interface {
	// EGOcor5 returns the physical value of the EGOcor5 signal.
	EGOcor5() float64
	// EGOcor6 returns the physical value of the EGOcor6 signal.
	EGOcor6() float64
	// EGOcor7 returns the physical value of the EGOcor7 signal.
	EGOcor7() float64
	// EGOcor8 returns the physical value of the EGOcor8 signal.
	EGOcor8() float64
}

// megasquirt_gp39Writer provides write access to a megasquirt_gp39 message.
type megasquirt_gp39Writer interface {
	// CopyFrom copies all values from megasquirt_gp39Reader.
	CopyFrom(megasquirt_gp39Reader) *megasquirt_gp39
	// SetEGOcor5 sets the physical value of the EGOcor5 signal.
	SetEGOcor5(float64) *megasquirt_gp39
	// SetEGOcor6 sets the physical value of the EGOcor6 signal.
	SetEGOcor6(float64) *megasquirt_gp39
	// SetEGOcor7 sets the physical value of the EGOcor7 signal.
	SetEGOcor7(float64) *megasquirt_gp39
	// SetEGOcor8 sets the physical value of the EGOcor8 signal.
	SetEGOcor8(float64) *megasquirt_gp39
}

type megasquirt_gp39 struct {
	xxx_EGOcor5 int16
	xxx_EGOcor6 int16
	xxx_EGOcor7 int16
	xxx_EGOcor8 int16
}

func Newmegasquirt_gp39() *megasquirt_gp39 {
	m := &megasquirt_gp39{}
	m.Reset()
	return m
}

func (m *megasquirt_gp39) Reset() {
	m.xxx_EGOcor5 = 0
	m.xxx_EGOcor6 = 0
	m.xxx_EGOcor7 = 0
	m.xxx_EGOcor8 = 0
}

func (m *megasquirt_gp39) CopyFrom(o megasquirt_gp39Reader) *megasquirt_gp39 {
	m.SetEGOcor5(o.EGOcor5())
	m.SetEGOcor6(o.EGOcor6())
	m.SetEGOcor7(o.EGOcor7())
	m.SetEGOcor8(o.EGOcor8())
	return m
}

// Descriptor returns the megasquirt_gp39 descriptor.
func (m *megasquirt_gp39) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp39.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp39) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp39) EGOcor5() float64 {
	return Messages().megasquirt_gp39.EGOcor5.ToPhysical(float64(m.xxx_EGOcor5))
}

func (m *megasquirt_gp39) SetEGOcor5(v float64) *megasquirt_gp39 {
	m.xxx_EGOcor5 = int16(Messages().megasquirt_gp39.EGOcor5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp39) EGOcor6() float64 {
	return Messages().megasquirt_gp39.EGOcor6.ToPhysical(float64(m.xxx_EGOcor6))
}

func (m *megasquirt_gp39) SetEGOcor6(v float64) *megasquirt_gp39 {
	m.xxx_EGOcor6 = int16(Messages().megasquirt_gp39.EGOcor6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp39) EGOcor7() float64 {
	return Messages().megasquirt_gp39.EGOcor7.ToPhysical(float64(m.xxx_EGOcor7))
}

func (m *megasquirt_gp39) SetEGOcor7(v float64) *megasquirt_gp39 {
	m.xxx_EGOcor7 = int16(Messages().megasquirt_gp39.EGOcor7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp39) EGOcor8() float64 {
	return Messages().megasquirt_gp39.EGOcor8.ToPhysical(float64(m.xxx_EGOcor8))
}

func (m *megasquirt_gp39) SetEGOcor8(v float64) *megasquirt_gp39 {
	m.xxx_EGOcor8 = int16(Messages().megasquirt_gp39.EGOcor8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp39) Frame() can.Frame {
	md := Messages().megasquirt_gp39
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOcor5.MarshalSigned(&f.Data, int64(m.xxx_EGOcor5))
	md.EGOcor6.MarshalSigned(&f.Data, int64(m.xxx_EGOcor6))
	md.EGOcor7.MarshalSigned(&f.Data, int64(m.xxx_EGOcor7))
	md.EGOcor8.MarshalSigned(&f.Data, int64(m.xxx_EGOcor8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp39) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp39) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp39
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp39: expects ID 1559 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp39: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp39: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp39: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOcor5 = int16(md.EGOcor5.UnmarshalSigned(f.Data))
	m.xxx_EGOcor6 = int16(md.EGOcor6.UnmarshalSigned(f.Data))
	m.xxx_EGOcor7 = int16(md.EGOcor7.UnmarshalSigned(f.Data))
	m.xxx_EGOcor8 = int16(md.EGOcor8.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp40Reader provides read access to a megasquirt_gp40 message.
type megasquirt_gp40Reader interface {
	// EGOcor9 returns the physical value of the EGOcor9 signal.
	EGOcor9() float64
	// EGOcor10 returns the physical value of the EGOcor10 signal.
	EGOcor10() float64
	// EGOcor11 returns the physical value of the EGOcor11 signal.
	EGOcor11() float64
	// EGOcor12 returns the physical value of the EGOcor12 signal.
	EGOcor12() float64
}

// megasquirt_gp40Writer provides write access to a megasquirt_gp40 message.
type megasquirt_gp40Writer interface {
	// CopyFrom copies all values from megasquirt_gp40Reader.
	CopyFrom(megasquirt_gp40Reader) *megasquirt_gp40
	// SetEGOcor9 sets the physical value of the EGOcor9 signal.
	SetEGOcor9(float64) *megasquirt_gp40
	// SetEGOcor10 sets the physical value of the EGOcor10 signal.
	SetEGOcor10(float64) *megasquirt_gp40
	// SetEGOcor11 sets the physical value of the EGOcor11 signal.
	SetEGOcor11(float64) *megasquirt_gp40
	// SetEGOcor12 sets the physical value of the EGOcor12 signal.
	SetEGOcor12(float64) *megasquirt_gp40
}

type megasquirt_gp40 struct {
	xxx_EGOcor9  int16
	xxx_EGOcor10 int16
	xxx_EGOcor11 int16
	xxx_EGOcor12 int16
}

func Newmegasquirt_gp40() *megasquirt_gp40 {
	m := &megasquirt_gp40{}
	m.Reset()
	return m
}

func (m *megasquirt_gp40) Reset() {
	m.xxx_EGOcor9 = 0
	m.xxx_EGOcor10 = 0
	m.xxx_EGOcor11 = 0
	m.xxx_EGOcor12 = 0
}

func (m *megasquirt_gp40) CopyFrom(o megasquirt_gp40Reader) *megasquirt_gp40 {
	m.SetEGOcor9(o.EGOcor9())
	m.SetEGOcor10(o.EGOcor10())
	m.SetEGOcor11(o.EGOcor11())
	m.SetEGOcor12(o.EGOcor12())
	return m
}

// Descriptor returns the megasquirt_gp40 descriptor.
func (m *megasquirt_gp40) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp40.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp40) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp40) EGOcor9() float64 {
	return Messages().megasquirt_gp40.EGOcor9.ToPhysical(float64(m.xxx_EGOcor9))
}

func (m *megasquirt_gp40) SetEGOcor9(v float64) *megasquirt_gp40 {
	m.xxx_EGOcor9 = int16(Messages().megasquirt_gp40.EGOcor9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp40) EGOcor10() float64 {
	return Messages().megasquirt_gp40.EGOcor10.ToPhysical(float64(m.xxx_EGOcor10))
}

func (m *megasquirt_gp40) SetEGOcor10(v float64) *megasquirt_gp40 {
	m.xxx_EGOcor10 = int16(Messages().megasquirt_gp40.EGOcor10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp40) EGOcor11() float64 {
	return Messages().megasquirt_gp40.EGOcor11.ToPhysical(float64(m.xxx_EGOcor11))
}

func (m *megasquirt_gp40) SetEGOcor11(v float64) *megasquirt_gp40 {
	m.xxx_EGOcor11 = int16(Messages().megasquirt_gp40.EGOcor11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp40) EGOcor12() float64 {
	return Messages().megasquirt_gp40.EGOcor12.ToPhysical(float64(m.xxx_EGOcor12))
}

func (m *megasquirt_gp40) SetEGOcor12(v float64) *megasquirt_gp40 {
	m.xxx_EGOcor12 = int16(Messages().megasquirt_gp40.EGOcor12.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp40) Frame() can.Frame {
	md := Messages().megasquirt_gp40
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOcor9.MarshalSigned(&f.Data, int64(m.xxx_EGOcor9))
	md.EGOcor10.MarshalSigned(&f.Data, int64(m.xxx_EGOcor10))
	md.EGOcor11.MarshalSigned(&f.Data, int64(m.xxx_EGOcor11))
	md.EGOcor12.MarshalSigned(&f.Data, int64(m.xxx_EGOcor12))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp40) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp40) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp40
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp40: expects ID 1560 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp40: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp40: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp40: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOcor9 = int16(md.EGOcor9.UnmarshalSigned(f.Data))
	m.xxx_EGOcor10 = int16(md.EGOcor10.UnmarshalSigned(f.Data))
	m.xxx_EGOcor11 = int16(md.EGOcor11.UnmarshalSigned(f.Data))
	m.xxx_EGOcor12 = int16(md.EGOcor12.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp41Reader provides read access to a megasquirt_gp41 message.
type megasquirt_gp41Reader interface {
	// EGOcor13 returns the physical value of the EGOcor13 signal.
	EGOcor13() float64
	// EGOcor14 returns the physical value of the EGOcor14 signal.
	EGOcor14() float64
	// EGOcor15 returns the physical value of the EGOcor15 signal.
	EGOcor15() float64
	// EGOcor16 returns the physical value of the EGOcor16 signal.
	EGOcor16() float64
}

// megasquirt_gp41Writer provides write access to a megasquirt_gp41 message.
type megasquirt_gp41Writer interface {
	// CopyFrom copies all values from megasquirt_gp41Reader.
	CopyFrom(megasquirt_gp41Reader) *megasquirt_gp41
	// SetEGOcor13 sets the physical value of the EGOcor13 signal.
	SetEGOcor13(float64) *megasquirt_gp41
	// SetEGOcor14 sets the physical value of the EGOcor14 signal.
	SetEGOcor14(float64) *megasquirt_gp41
	// SetEGOcor15 sets the physical value of the EGOcor15 signal.
	SetEGOcor15(float64) *megasquirt_gp41
	// SetEGOcor16 sets the physical value of the EGOcor16 signal.
	SetEGOcor16(float64) *megasquirt_gp41
}

type megasquirt_gp41 struct {
	xxx_EGOcor13 int16
	xxx_EGOcor14 int16
	xxx_EGOcor15 int16
	xxx_EGOcor16 int16
}

func Newmegasquirt_gp41() *megasquirt_gp41 {
	m := &megasquirt_gp41{}
	m.Reset()
	return m
}

func (m *megasquirt_gp41) Reset() {
	m.xxx_EGOcor13 = 0
	m.xxx_EGOcor14 = 0
	m.xxx_EGOcor15 = 0
	m.xxx_EGOcor16 = 0
}

func (m *megasquirt_gp41) CopyFrom(o megasquirt_gp41Reader) *megasquirt_gp41 {
	m.SetEGOcor13(o.EGOcor13())
	m.SetEGOcor14(o.EGOcor14())
	m.SetEGOcor15(o.EGOcor15())
	m.SetEGOcor16(o.EGOcor16())
	return m
}

// Descriptor returns the megasquirt_gp41 descriptor.
func (m *megasquirt_gp41) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp41.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp41) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp41) EGOcor13() float64 {
	return Messages().megasquirt_gp41.EGOcor13.ToPhysical(float64(m.xxx_EGOcor13))
}

func (m *megasquirt_gp41) SetEGOcor13(v float64) *megasquirt_gp41 {
	m.xxx_EGOcor13 = int16(Messages().megasquirt_gp41.EGOcor13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp41) EGOcor14() float64 {
	return Messages().megasquirt_gp41.EGOcor14.ToPhysical(float64(m.xxx_EGOcor14))
}

func (m *megasquirt_gp41) SetEGOcor14(v float64) *megasquirt_gp41 {
	m.xxx_EGOcor14 = int16(Messages().megasquirt_gp41.EGOcor14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp41) EGOcor15() float64 {
	return Messages().megasquirt_gp41.EGOcor15.ToPhysical(float64(m.xxx_EGOcor15))
}

func (m *megasquirt_gp41) SetEGOcor15(v float64) *megasquirt_gp41 {
	m.xxx_EGOcor15 = int16(Messages().megasquirt_gp41.EGOcor15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp41) EGOcor16() float64 {
	return Messages().megasquirt_gp41.EGOcor16.ToPhysical(float64(m.xxx_EGOcor16))
}

func (m *megasquirt_gp41) SetEGOcor16(v float64) *megasquirt_gp41 {
	m.xxx_EGOcor16 = int16(Messages().megasquirt_gp41.EGOcor16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp41) Frame() can.Frame {
	md := Messages().megasquirt_gp41
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.EGOcor13.MarshalSigned(&f.Data, int64(m.xxx_EGOcor13))
	md.EGOcor14.MarshalSigned(&f.Data, int64(m.xxx_EGOcor14))
	md.EGOcor15.MarshalSigned(&f.Data, int64(m.xxx_EGOcor15))
	md.EGOcor16.MarshalSigned(&f.Data, int64(m.xxx_EGOcor16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp41) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp41) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp41
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp41: expects ID 1561 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp41: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp41: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp41: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_EGOcor13 = int16(md.EGOcor13.UnmarshalSigned(f.Data))
	m.xxx_EGOcor14 = int16(md.EGOcor14.UnmarshalSigned(f.Data))
	m.xxx_EGOcor15 = int16(md.EGOcor15.UnmarshalSigned(f.Data))
	m.xxx_EGOcor16 = int16(md.EGOcor16.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp42Reader provides read access to a megasquirt_gp42 message.
type megasquirt_gp42Reader interface {
	// VSS1 returns the physical value of the VSS1 signal.
	VSS1() float64
	// VSS2 returns the physical value of the VSS2 signal.
	VSS2() float64
	// VSS3 returns the physical value of the VSS3 signal.
	VSS3() float64
	// VSS4 returns the physical value of the VSS4 signal.
	VSS4() float64
}

// megasquirt_gp42Writer provides write access to a megasquirt_gp42 message.
type megasquirt_gp42Writer interface {
	// CopyFrom copies all values from megasquirt_gp42Reader.
	CopyFrom(megasquirt_gp42Reader) *megasquirt_gp42
	// SetVSS1 sets the physical value of the VSS1 signal.
	SetVSS1(float64) *megasquirt_gp42
	// SetVSS2 sets the physical value of the VSS2 signal.
	SetVSS2(float64) *megasquirt_gp42
	// SetVSS3 sets the physical value of the VSS3 signal.
	SetVSS3(float64) *megasquirt_gp42
	// SetVSS4 sets the physical value of the VSS4 signal.
	SetVSS4(float64) *megasquirt_gp42
}

type megasquirt_gp42 struct {
	xxx_VSS1 uint16
	xxx_VSS2 uint16
	xxx_VSS3 uint16
	xxx_VSS4 uint16
}

func Newmegasquirt_gp42() *megasquirt_gp42 {
	m := &megasquirt_gp42{}
	m.Reset()
	return m
}

func (m *megasquirt_gp42) Reset() {
	m.xxx_VSS1 = 0
	m.xxx_VSS2 = 0
	m.xxx_VSS3 = 0
	m.xxx_VSS4 = 0
}

func (m *megasquirt_gp42) CopyFrom(o megasquirt_gp42Reader) *megasquirt_gp42 {
	m.SetVSS1(o.VSS1())
	m.SetVSS2(o.VSS2())
	m.SetVSS3(o.VSS3())
	m.SetVSS4(o.VSS4())
	return m
}

// Descriptor returns the megasquirt_gp42 descriptor.
func (m *megasquirt_gp42) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp42.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp42) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp42) VSS1() float64 {
	return Messages().megasquirt_gp42.VSS1.ToPhysical(float64(m.xxx_VSS1))
}

func (m *megasquirt_gp42) SetVSS1(v float64) *megasquirt_gp42 {
	m.xxx_VSS1 = uint16(Messages().megasquirt_gp42.VSS1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp42) VSS2() float64 {
	return Messages().megasquirt_gp42.VSS2.ToPhysical(float64(m.xxx_VSS2))
}

func (m *megasquirt_gp42) SetVSS2(v float64) *megasquirt_gp42 {
	m.xxx_VSS2 = uint16(Messages().megasquirt_gp42.VSS2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp42) VSS3() float64 {
	return Messages().megasquirt_gp42.VSS3.ToPhysical(float64(m.xxx_VSS3))
}

func (m *megasquirt_gp42) SetVSS3(v float64) *megasquirt_gp42 {
	m.xxx_VSS3 = uint16(Messages().megasquirt_gp42.VSS3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp42) VSS4() float64 {
	return Messages().megasquirt_gp42.VSS4.ToPhysical(float64(m.xxx_VSS4))
}

func (m *megasquirt_gp42) SetVSS4(v float64) *megasquirt_gp42 {
	m.xxx_VSS4 = uint16(Messages().megasquirt_gp42.VSS4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp42) Frame() can.Frame {
	md := Messages().megasquirt_gp42
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.VSS1.MarshalUnsigned(&f.Data, uint64(m.xxx_VSS1))
	md.VSS2.MarshalUnsigned(&f.Data, uint64(m.xxx_VSS2))
	md.VSS3.MarshalUnsigned(&f.Data, uint64(m.xxx_VSS3))
	md.VSS4.MarshalUnsigned(&f.Data, uint64(m.xxx_VSS4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp42) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp42) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp42
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp42: expects ID 1562 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp42: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp42: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp42: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_VSS1 = uint16(md.VSS1.UnmarshalUnsigned(f.Data))
	m.xxx_VSS2 = uint16(md.VSS2.UnmarshalUnsigned(f.Data))
	m.xxx_VSS3 = uint16(md.VSS3.UnmarshalUnsigned(f.Data))
	m.xxx_VSS4 = uint16(md.VSS4.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp43Reader provides read access to a megasquirt_gp43 message.
type megasquirt_gp43Reader interface {
	// synccnt returns the value of the synccnt signal.
	synccnt() uint8
	// syncreason returns the value of the syncreason signal.
	syncreason() uint8
	// sd_filenum returns the value of the sd_filenum signal.
	sd_filenum() uint16
	// sd_error returns the value of the sd_error signal.
	sd_error() uint8
	// sd_phase returns the value of the sd_phase signal.
	sd_phase() uint8
	// sd_status returns the value of the sd_status signal.
	sd_status() uint8
	// timing_err returns the value of the timing_err signal.
	timing_err() int8
}

// megasquirt_gp43Writer provides write access to a megasquirt_gp43 message.
type megasquirt_gp43Writer interface {
	// CopyFrom copies all values from megasquirt_gp43Reader.
	CopyFrom(megasquirt_gp43Reader) *megasquirt_gp43
	// Setsynccnt sets the value of the synccnt signal.
	Setsynccnt(uint8) *megasquirt_gp43
	// Setsyncreason sets the value of the syncreason signal.
	Setsyncreason(uint8) *megasquirt_gp43
	// Setsd_filenum sets the value of the sd_filenum signal.
	Setsd_filenum(uint16) *megasquirt_gp43
	// Setsd_error sets the value of the sd_error signal.
	Setsd_error(uint8) *megasquirt_gp43
	// Setsd_phase sets the value of the sd_phase signal.
	Setsd_phase(uint8) *megasquirt_gp43
	// Setsd_status sets the value of the sd_status signal.
	Setsd_status(uint8) *megasquirt_gp43
	// Settiming_err sets the value of the timing_err signal.
	Settiming_err(int8) *megasquirt_gp43
}

type megasquirt_gp43 struct {
	xxx_synccnt    uint8
	xxx_syncreason uint8
	xxx_sd_filenum uint16
	xxx_sd_error   uint8
	xxx_sd_phase   uint8
	xxx_sd_status  uint8
	xxx_timing_err int8
}

func Newmegasquirt_gp43() *megasquirt_gp43 {
	m := &megasquirt_gp43{}
	m.Reset()
	return m
}

func (m *megasquirt_gp43) Reset() {
	m.xxx_synccnt = 0
	m.xxx_syncreason = 0
	m.xxx_sd_filenum = 0
	m.xxx_sd_error = 0
	m.xxx_sd_phase = 0
	m.xxx_sd_status = 0
	m.xxx_timing_err = 0
}

func (m *megasquirt_gp43) CopyFrom(o megasquirt_gp43Reader) *megasquirt_gp43 {
	m.xxx_synccnt = o.synccnt()
	m.xxx_syncreason = o.syncreason()
	m.xxx_sd_filenum = o.sd_filenum()
	m.xxx_sd_error = o.sd_error()
	m.xxx_sd_phase = o.sd_phase()
	m.xxx_sd_status = o.sd_status()
	m.xxx_timing_err = o.timing_err()
	return m
}

// Descriptor returns the megasquirt_gp43 descriptor.
func (m *megasquirt_gp43) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp43.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp43) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp43) synccnt() uint8 {
	return m.xxx_synccnt
}

func (m *megasquirt_gp43) Setsynccnt(v uint8) *megasquirt_gp43 {
	m.xxx_synccnt = uint8(Messages().megasquirt_gp43.synccnt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp43) syncreason() uint8 {
	return m.xxx_syncreason
}

func (m *megasquirt_gp43) Setsyncreason(v uint8) *megasquirt_gp43 {
	m.xxx_syncreason = uint8(Messages().megasquirt_gp43.syncreason.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp43) sd_filenum() uint16 {
	return m.xxx_sd_filenum
}

func (m *megasquirt_gp43) Setsd_filenum(v uint16) *megasquirt_gp43 {
	m.xxx_sd_filenum = uint16(Messages().megasquirt_gp43.sd_filenum.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp43) sd_error() uint8 {
	return m.xxx_sd_error
}

func (m *megasquirt_gp43) Setsd_error(v uint8) *megasquirt_gp43 {
	m.xxx_sd_error = uint8(Messages().megasquirt_gp43.sd_error.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp43) sd_phase() uint8 {
	return m.xxx_sd_phase
}

func (m *megasquirt_gp43) Setsd_phase(v uint8) *megasquirt_gp43 {
	m.xxx_sd_phase = uint8(Messages().megasquirt_gp43.sd_phase.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp43) sd_status() uint8 {
	return m.xxx_sd_status
}

func (m *megasquirt_gp43) Setsd_status(v uint8) *megasquirt_gp43 {
	m.xxx_sd_status = uint8(Messages().megasquirt_gp43.sd_status.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp43) timing_err() int8 {
	return m.xxx_timing_err
}

func (m *megasquirt_gp43) Settiming_err(v int8) *megasquirt_gp43 {
	m.xxx_timing_err = int8(Messages().megasquirt_gp43.timing_err.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp43) Frame() can.Frame {
	md := Messages().megasquirt_gp43
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.synccnt.MarshalUnsigned(&f.Data, uint64(m.xxx_synccnt))
	md.syncreason.MarshalUnsigned(&f.Data, uint64(m.xxx_syncreason))
	md.sd_filenum.MarshalUnsigned(&f.Data, uint64(m.xxx_sd_filenum))
	md.sd_error.MarshalUnsigned(&f.Data, uint64(m.xxx_sd_error))
	md.sd_phase.MarshalUnsigned(&f.Data, uint64(m.xxx_sd_phase))
	md.sd_status.MarshalUnsigned(&f.Data, uint64(m.xxx_sd_status))
	md.timing_err.MarshalSigned(&f.Data, int64(m.xxx_timing_err))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp43) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp43) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp43
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp43: expects ID 1563 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp43: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp43: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp43: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_synccnt = uint8(md.synccnt.UnmarshalUnsigned(f.Data))
	m.xxx_syncreason = uint8(md.syncreason.UnmarshalUnsigned(f.Data))
	m.xxx_sd_filenum = uint16(md.sd_filenum.UnmarshalUnsigned(f.Data))
	m.xxx_sd_error = uint8(md.sd_error.UnmarshalUnsigned(f.Data))
	m.xxx_sd_phase = uint8(md.sd_phase.UnmarshalUnsigned(f.Data))
	m.xxx_sd_status = uint8(md.sd_status.UnmarshalUnsigned(f.Data))
	m.xxx_timing_err = int8(md.timing_err.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp44Reader provides read access to a megasquirt_gp44 message.
type megasquirt_gp44Reader interface {
	// vvt_ang1 returns the value of the vvt_ang1 signal.
	vvt_ang1() int16
	// vvt_ang2 returns the value of the vvt_ang2 signal.
	vvt_ang2() int16
	// vvt_ang3 returns the value of the vvt_ang3 signal.
	vvt_ang3() int16
	// vvt_ang4 returns the value of the vvt_ang4 signal.
	vvt_ang4() int16
}

// megasquirt_gp44Writer provides write access to a megasquirt_gp44 message.
type megasquirt_gp44Writer interface {
	// CopyFrom copies all values from megasquirt_gp44Reader.
	CopyFrom(megasquirt_gp44Reader) *megasquirt_gp44
	// Setvvt_ang1 sets the value of the vvt_ang1 signal.
	Setvvt_ang1(int16) *megasquirt_gp44
	// Setvvt_ang2 sets the value of the vvt_ang2 signal.
	Setvvt_ang2(int16) *megasquirt_gp44
	// Setvvt_ang3 sets the value of the vvt_ang3 signal.
	Setvvt_ang3(int16) *megasquirt_gp44
	// Setvvt_ang4 sets the value of the vvt_ang4 signal.
	Setvvt_ang4(int16) *megasquirt_gp44
}

type megasquirt_gp44 struct {
	xxx_vvt_ang1 int16
	xxx_vvt_ang2 int16
	xxx_vvt_ang3 int16
	xxx_vvt_ang4 int16
}

func Newmegasquirt_gp44() *megasquirt_gp44 {
	m := &megasquirt_gp44{}
	m.Reset()
	return m
}

func (m *megasquirt_gp44) Reset() {
	m.xxx_vvt_ang1 = 0
	m.xxx_vvt_ang2 = 0
	m.xxx_vvt_ang3 = 0
	m.xxx_vvt_ang4 = 0
}

func (m *megasquirt_gp44) CopyFrom(o megasquirt_gp44Reader) *megasquirt_gp44 {
	m.xxx_vvt_ang1 = o.vvt_ang1()
	m.xxx_vvt_ang2 = o.vvt_ang2()
	m.xxx_vvt_ang3 = o.vvt_ang3()
	m.xxx_vvt_ang4 = o.vvt_ang4()
	return m
}

// Descriptor returns the megasquirt_gp44 descriptor.
func (m *megasquirt_gp44) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp44.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp44) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp44) vvt_ang1() int16 {
	return m.xxx_vvt_ang1
}

func (m *megasquirt_gp44) Setvvt_ang1(v int16) *megasquirt_gp44 {
	m.xxx_vvt_ang1 = int16(Messages().megasquirt_gp44.vvt_ang1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp44) vvt_ang2() int16 {
	return m.xxx_vvt_ang2
}

func (m *megasquirt_gp44) Setvvt_ang2(v int16) *megasquirt_gp44 {
	m.xxx_vvt_ang2 = int16(Messages().megasquirt_gp44.vvt_ang2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp44) vvt_ang3() int16 {
	return m.xxx_vvt_ang3
}

func (m *megasquirt_gp44) Setvvt_ang3(v int16) *megasquirt_gp44 {
	m.xxx_vvt_ang3 = int16(Messages().megasquirt_gp44.vvt_ang3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp44) vvt_ang4() int16 {
	return m.xxx_vvt_ang4
}

func (m *megasquirt_gp44) Setvvt_ang4(v int16) *megasquirt_gp44 {
	m.xxx_vvt_ang4 = int16(Messages().megasquirt_gp44.vvt_ang4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp44) Frame() can.Frame {
	md := Messages().megasquirt_gp44
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.vvt_ang1.MarshalSigned(&f.Data, int64(m.xxx_vvt_ang1))
	md.vvt_ang2.MarshalSigned(&f.Data, int64(m.xxx_vvt_ang2))
	md.vvt_ang3.MarshalSigned(&f.Data, int64(m.xxx_vvt_ang3))
	md.vvt_ang4.MarshalSigned(&f.Data, int64(m.xxx_vvt_ang4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp44) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp44) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp44
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp44: expects ID 1564 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp44: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp44: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp44: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_vvt_ang1 = int16(md.vvt_ang1.UnmarshalSigned(f.Data))
	m.xxx_vvt_ang2 = int16(md.vvt_ang2.UnmarshalSigned(f.Data))
	m.xxx_vvt_ang3 = int16(md.vvt_ang3.UnmarshalSigned(f.Data))
	m.xxx_vvt_ang4 = int16(md.vvt_ang4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp45Reader provides read access to a megasquirt_gp45 message.
type megasquirt_gp45Reader interface {
	// vvt_target1 returns the value of the vvt_target1 signal.
	vvt_target1() int16
	// vvt_target2 returns the value of the vvt_target2 signal.
	vvt_target2() int16
	// vvt_target3 returns the value of the vvt_target3 signal.
	vvt_target3() int16
	// vvt_target4 returns the value of the vvt_target4 signal.
	vvt_target4() int16
}

// megasquirt_gp45Writer provides write access to a megasquirt_gp45 message.
type megasquirt_gp45Writer interface {
	// CopyFrom copies all values from megasquirt_gp45Reader.
	CopyFrom(megasquirt_gp45Reader) *megasquirt_gp45
	// Setvvt_target1 sets the value of the vvt_target1 signal.
	Setvvt_target1(int16) *megasquirt_gp45
	// Setvvt_target2 sets the value of the vvt_target2 signal.
	Setvvt_target2(int16) *megasquirt_gp45
	// Setvvt_target3 sets the value of the vvt_target3 signal.
	Setvvt_target3(int16) *megasquirt_gp45
	// Setvvt_target4 sets the value of the vvt_target4 signal.
	Setvvt_target4(int16) *megasquirt_gp45
}

type megasquirt_gp45 struct {
	xxx_vvt_target1 int16
	xxx_vvt_target2 int16
	xxx_vvt_target3 int16
	xxx_vvt_target4 int16
}

func Newmegasquirt_gp45() *megasquirt_gp45 {
	m := &megasquirt_gp45{}
	m.Reset()
	return m
}

func (m *megasquirt_gp45) Reset() {
	m.xxx_vvt_target1 = 0
	m.xxx_vvt_target2 = 0
	m.xxx_vvt_target3 = 0
	m.xxx_vvt_target4 = 0
}

func (m *megasquirt_gp45) CopyFrom(o megasquirt_gp45Reader) *megasquirt_gp45 {
	m.xxx_vvt_target1 = o.vvt_target1()
	m.xxx_vvt_target2 = o.vvt_target2()
	m.xxx_vvt_target3 = o.vvt_target3()
	m.xxx_vvt_target4 = o.vvt_target4()
	return m
}

// Descriptor returns the megasquirt_gp45 descriptor.
func (m *megasquirt_gp45) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp45.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp45) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp45) vvt_target1() int16 {
	return m.xxx_vvt_target1
}

func (m *megasquirt_gp45) Setvvt_target1(v int16) *megasquirt_gp45 {
	m.xxx_vvt_target1 = int16(Messages().megasquirt_gp45.vvt_target1.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp45) vvt_target2() int16 {
	return m.xxx_vvt_target2
}

func (m *megasquirt_gp45) Setvvt_target2(v int16) *megasquirt_gp45 {
	m.xxx_vvt_target2 = int16(Messages().megasquirt_gp45.vvt_target2.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp45) vvt_target3() int16 {
	return m.xxx_vvt_target3
}

func (m *megasquirt_gp45) Setvvt_target3(v int16) *megasquirt_gp45 {
	m.xxx_vvt_target3 = int16(Messages().megasquirt_gp45.vvt_target3.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp45) vvt_target4() int16 {
	return m.xxx_vvt_target4
}

func (m *megasquirt_gp45) Setvvt_target4(v int16) *megasquirt_gp45 {
	m.xxx_vvt_target4 = int16(Messages().megasquirt_gp45.vvt_target4.SaturatedCastSigned(int64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp45) Frame() can.Frame {
	md := Messages().megasquirt_gp45
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.vvt_target1.MarshalSigned(&f.Data, int64(m.xxx_vvt_target1))
	md.vvt_target2.MarshalSigned(&f.Data, int64(m.xxx_vvt_target2))
	md.vvt_target3.MarshalSigned(&f.Data, int64(m.xxx_vvt_target3))
	md.vvt_target4.MarshalSigned(&f.Data, int64(m.xxx_vvt_target4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp45) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp45) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp45
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp45: expects ID 1565 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp45: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp45: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp45: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_vvt_target1 = int16(md.vvt_target1.UnmarshalSigned(f.Data))
	m.xxx_vvt_target2 = int16(md.vvt_target2.UnmarshalSigned(f.Data))
	m.xxx_vvt_target3 = int16(md.vvt_target3.UnmarshalSigned(f.Data))
	m.xxx_vvt_target4 = int16(md.vvt_target4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp46Reader provides read access to a megasquirt_gp46 message.
type megasquirt_gp46Reader interface {
	// vvt_duty1 returns the physical value of the vvt_duty1 signal.
	vvt_duty1() float64
	// vvt_duty2 returns the physical value of the vvt_duty2 signal.
	vvt_duty2() float64
	// vvt_duty3 returns the physical value of the vvt_duty3 signal.
	vvt_duty3() float64
	// vvt_duty4 returns the physical value of the vvt_duty4 signal.
	vvt_duty4() float64
	// inj_timing_pri returns the physical value of the inj_timing_pri signal.
	inj_timing_pri() float64
	// inj_timing_sec returns the physical value of the inj_timing_sec signal.
	inj_timing_sec() float64
}

// megasquirt_gp46Writer provides write access to a megasquirt_gp46 message.
type megasquirt_gp46Writer interface {
	// CopyFrom copies all values from megasquirt_gp46Reader.
	CopyFrom(megasquirt_gp46Reader) *megasquirt_gp46
	// Setvvt_duty1 sets the physical value of the vvt_duty1 signal.
	Setvvt_duty1(float64) *megasquirt_gp46
	// Setvvt_duty2 sets the physical value of the vvt_duty2 signal.
	Setvvt_duty2(float64) *megasquirt_gp46
	// Setvvt_duty3 sets the physical value of the vvt_duty3 signal.
	Setvvt_duty3(float64) *megasquirt_gp46
	// Setvvt_duty4 sets the physical value of the vvt_duty4 signal.
	Setvvt_duty4(float64) *megasquirt_gp46
	// Setinj_timing_pri sets the physical value of the inj_timing_pri signal.
	Setinj_timing_pri(float64) *megasquirt_gp46
	// Setinj_timing_sec sets the physical value of the inj_timing_sec signal.
	Setinj_timing_sec(float64) *megasquirt_gp46
}

type megasquirt_gp46 struct {
	xxx_vvt_duty1      uint8
	xxx_vvt_duty2      uint8
	xxx_vvt_duty3      uint8
	xxx_vvt_duty4      uint8
	xxx_inj_timing_pri int16
	xxx_inj_timing_sec int16
}

func Newmegasquirt_gp46() *megasquirt_gp46 {
	m := &megasquirt_gp46{}
	m.Reset()
	return m
}

func (m *megasquirt_gp46) Reset() {
	m.xxx_vvt_duty1 = 0
	m.xxx_vvt_duty2 = 0
	m.xxx_vvt_duty3 = 0
	m.xxx_vvt_duty4 = 0
	m.xxx_inj_timing_pri = 0
	m.xxx_inj_timing_sec = 0
}

func (m *megasquirt_gp46) CopyFrom(o megasquirt_gp46Reader) *megasquirt_gp46 {
	m.Setvvt_duty1(o.vvt_duty1())
	m.Setvvt_duty2(o.vvt_duty2())
	m.Setvvt_duty3(o.vvt_duty3())
	m.Setvvt_duty4(o.vvt_duty4())
	m.Setinj_timing_pri(o.inj_timing_pri())
	m.Setinj_timing_sec(o.inj_timing_sec())
	return m
}

// Descriptor returns the megasquirt_gp46 descriptor.
func (m *megasquirt_gp46) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp46.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp46) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp46) vvt_duty1() float64 {
	return Messages().megasquirt_gp46.vvt_duty1.ToPhysical(float64(m.xxx_vvt_duty1))
}

func (m *megasquirt_gp46) Setvvt_duty1(v float64) *megasquirt_gp46 {
	m.xxx_vvt_duty1 = uint8(Messages().megasquirt_gp46.vvt_duty1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp46) vvt_duty2() float64 {
	return Messages().megasquirt_gp46.vvt_duty2.ToPhysical(float64(m.xxx_vvt_duty2))
}

func (m *megasquirt_gp46) Setvvt_duty2(v float64) *megasquirt_gp46 {
	m.xxx_vvt_duty2 = uint8(Messages().megasquirt_gp46.vvt_duty2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp46) vvt_duty3() float64 {
	return Messages().megasquirt_gp46.vvt_duty3.ToPhysical(float64(m.xxx_vvt_duty3))
}

func (m *megasquirt_gp46) Setvvt_duty3(v float64) *megasquirt_gp46 {
	m.xxx_vvt_duty3 = uint8(Messages().megasquirt_gp46.vvt_duty3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp46) vvt_duty4() float64 {
	return Messages().megasquirt_gp46.vvt_duty4.ToPhysical(float64(m.xxx_vvt_duty4))
}

func (m *megasquirt_gp46) Setvvt_duty4(v float64) *megasquirt_gp46 {
	m.xxx_vvt_duty4 = uint8(Messages().megasquirt_gp46.vvt_duty4.FromPhysical(v))
	return m
}

func (m *megasquirt_gp46) inj_timing_pri() float64 {
	return Messages().megasquirt_gp46.inj_timing_pri.ToPhysical(float64(m.xxx_inj_timing_pri))
}

func (m *megasquirt_gp46) Setinj_timing_pri(v float64) *megasquirt_gp46 {
	m.xxx_inj_timing_pri = int16(Messages().megasquirt_gp46.inj_timing_pri.FromPhysical(v))
	return m
}

func (m *megasquirt_gp46) inj_timing_sec() float64 {
	return Messages().megasquirt_gp46.inj_timing_sec.ToPhysical(float64(m.xxx_inj_timing_sec))
}

func (m *megasquirt_gp46) Setinj_timing_sec(v float64) *megasquirt_gp46 {
	m.xxx_inj_timing_sec = int16(Messages().megasquirt_gp46.inj_timing_sec.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp46) Frame() can.Frame {
	md := Messages().megasquirt_gp46
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.vvt_duty1.MarshalUnsigned(&f.Data, uint64(m.xxx_vvt_duty1))
	md.vvt_duty2.MarshalUnsigned(&f.Data, uint64(m.xxx_vvt_duty2))
	md.vvt_duty3.MarshalUnsigned(&f.Data, uint64(m.xxx_vvt_duty3))
	md.vvt_duty4.MarshalUnsigned(&f.Data, uint64(m.xxx_vvt_duty4))
	md.inj_timing_pri.MarshalSigned(&f.Data, int64(m.xxx_inj_timing_pri))
	md.inj_timing_sec.MarshalSigned(&f.Data, int64(m.xxx_inj_timing_sec))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp46) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp46) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp46
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp46: expects ID 1566 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp46: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp46: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp46: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_vvt_duty1 = uint8(md.vvt_duty1.UnmarshalUnsigned(f.Data))
	m.xxx_vvt_duty2 = uint8(md.vvt_duty2.UnmarshalUnsigned(f.Data))
	m.xxx_vvt_duty3 = uint8(md.vvt_duty3.UnmarshalUnsigned(f.Data))
	m.xxx_vvt_duty4 = uint8(md.vvt_duty4.UnmarshalUnsigned(f.Data))
	m.xxx_inj_timing_pri = int16(md.inj_timing_pri.UnmarshalSigned(f.Data))
	m.xxx_inj_timing_sec = int16(md.inj_timing_sec.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp47Reader provides read access to a megasquirt_gp47 message.
type megasquirt_gp47Reader interface {
	// fuel_pct returns the physical value of the fuel_pct signal.
	fuel_pct() float64
	// tps_accel returns the physical value of the tps_accel signal.
	tps_accel() float64
	// SS1 returns the physical value of the SS1 signal.
	SS1() float64
	// SS2 returns the physical value of the SS2 signal.
	SS2() float64
}

// megasquirt_gp47Writer provides write access to a megasquirt_gp47 message.
type megasquirt_gp47Writer interface {
	// CopyFrom copies all values from megasquirt_gp47Reader.
	CopyFrom(megasquirt_gp47Reader) *megasquirt_gp47
	// Setfuel_pct sets the physical value of the fuel_pct signal.
	Setfuel_pct(float64) *megasquirt_gp47
	// Settps_accel sets the physical value of the tps_accel signal.
	Settps_accel(float64) *megasquirt_gp47
	// SetSS1 sets the physical value of the SS1 signal.
	SetSS1(float64) *megasquirt_gp47
	// SetSS2 sets the physical value of the SS2 signal.
	SetSS2(float64) *megasquirt_gp47
}

type megasquirt_gp47 struct {
	xxx_fuel_pct  int16
	xxx_tps_accel int16
	xxx_SS1       uint16
	xxx_SS2       uint16
}

func Newmegasquirt_gp47() *megasquirt_gp47 {
	m := &megasquirt_gp47{}
	m.Reset()
	return m
}

func (m *megasquirt_gp47) Reset() {
	m.xxx_fuel_pct = 0
	m.xxx_tps_accel = 0
	m.xxx_SS1 = 0
	m.xxx_SS2 = 0
}

func (m *megasquirt_gp47) CopyFrom(o megasquirt_gp47Reader) *megasquirt_gp47 {
	m.Setfuel_pct(o.fuel_pct())
	m.Settps_accel(o.tps_accel())
	m.SetSS1(o.SS1())
	m.SetSS2(o.SS2())
	return m
}

// Descriptor returns the megasquirt_gp47 descriptor.
func (m *megasquirt_gp47) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp47.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp47) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp47) fuel_pct() float64 {
	return Messages().megasquirt_gp47.fuel_pct.ToPhysical(float64(m.xxx_fuel_pct))
}

func (m *megasquirt_gp47) Setfuel_pct(v float64) *megasquirt_gp47 {
	m.xxx_fuel_pct = int16(Messages().megasquirt_gp47.fuel_pct.FromPhysical(v))
	return m
}

func (m *megasquirt_gp47) tps_accel() float64 {
	return Messages().megasquirt_gp47.tps_accel.ToPhysical(float64(m.xxx_tps_accel))
}

func (m *megasquirt_gp47) Settps_accel(v float64) *megasquirt_gp47 {
	m.xxx_tps_accel = int16(Messages().megasquirt_gp47.tps_accel.FromPhysical(v))
	return m
}

func (m *megasquirt_gp47) SS1() float64 {
	return Messages().megasquirt_gp47.SS1.ToPhysical(float64(m.xxx_SS1))
}

func (m *megasquirt_gp47) SetSS1(v float64) *megasquirt_gp47 {
	m.xxx_SS1 = uint16(Messages().megasquirt_gp47.SS1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp47) SS2() float64 {
	return Messages().megasquirt_gp47.SS2.ToPhysical(float64(m.xxx_SS2))
}

func (m *megasquirt_gp47) SetSS2(v float64) *megasquirt_gp47 {
	m.xxx_SS2 = uint16(Messages().megasquirt_gp47.SS2.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp47) Frame() can.Frame {
	md := Messages().megasquirt_gp47
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.fuel_pct.MarshalSigned(&f.Data, int64(m.xxx_fuel_pct))
	md.tps_accel.MarshalSigned(&f.Data, int64(m.xxx_tps_accel))
	md.SS1.MarshalUnsigned(&f.Data, uint64(m.xxx_SS1))
	md.SS2.MarshalUnsigned(&f.Data, uint64(m.xxx_SS2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp47) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp47) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp47
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp47: expects ID 1567 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp47: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp47: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp47: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_fuel_pct = int16(md.fuel_pct.UnmarshalSigned(f.Data))
	m.xxx_tps_accel = int16(md.tps_accel.UnmarshalSigned(f.Data))
	m.xxx_SS1 = uint16(md.SS1.UnmarshalUnsigned(f.Data))
	m.xxx_SS2 = uint16(md.SS2.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp48Reader provides read access to a megasquirt_gp48 message.
type megasquirt_gp48Reader interface {
	// knock_cyl1 returns the physical value of the knock_cyl1 signal.
	knock_cyl1() float64
	// knock_cyl2 returns the physical value of the knock_cyl2 signal.
	knock_cyl2() float64
	// knock_cyl3 returns the physical value of the knock_cyl3 signal.
	knock_cyl3() float64
	// knock_cyl4 returns the physical value of the knock_cyl4 signal.
	knock_cyl4() float64
	// knock_cyl5 returns the physical value of the knock_cyl5 signal.
	knock_cyl5() float64
	// knock_cyl6 returns the physical value of the knock_cyl6 signal.
	knock_cyl6() float64
	// knock_cyl7 returns the physical value of the knock_cyl7 signal.
	knock_cyl7() float64
	// knock_cyl8 returns the physical value of the knock_cyl8 signal.
	knock_cyl8() float64
}

// megasquirt_gp48Writer provides write access to a megasquirt_gp48 message.
type megasquirt_gp48Writer interface {
	// CopyFrom copies all values from megasquirt_gp48Reader.
	CopyFrom(megasquirt_gp48Reader) *megasquirt_gp48
	// Setknock_cyl1 sets the physical value of the knock_cyl1 signal.
	Setknock_cyl1(float64) *megasquirt_gp48
	// Setknock_cyl2 sets the physical value of the knock_cyl2 signal.
	Setknock_cyl2(float64) *megasquirt_gp48
	// Setknock_cyl3 sets the physical value of the knock_cyl3 signal.
	Setknock_cyl3(float64) *megasquirt_gp48
	// Setknock_cyl4 sets the physical value of the knock_cyl4 signal.
	Setknock_cyl4(float64) *megasquirt_gp48
	// Setknock_cyl5 sets the physical value of the knock_cyl5 signal.
	Setknock_cyl5(float64) *megasquirt_gp48
	// Setknock_cyl6 sets the physical value of the knock_cyl6 signal.
	Setknock_cyl6(float64) *megasquirt_gp48
	// Setknock_cyl7 sets the physical value of the knock_cyl7 signal.
	Setknock_cyl7(float64) *megasquirt_gp48
	// Setknock_cyl8 sets the physical value of the knock_cyl8 signal.
	Setknock_cyl8(float64) *megasquirt_gp48
}

type megasquirt_gp48 struct {
	xxx_knock_cyl1 uint8
	xxx_knock_cyl2 uint8
	xxx_knock_cyl3 uint8
	xxx_knock_cyl4 uint8
	xxx_knock_cyl5 uint8
	xxx_knock_cyl6 uint8
	xxx_knock_cyl7 uint8
	xxx_knock_cyl8 uint8
}

func Newmegasquirt_gp48() *megasquirt_gp48 {
	m := &megasquirt_gp48{}
	m.Reset()
	return m
}

func (m *megasquirt_gp48) Reset() {
	m.xxx_knock_cyl1 = 0
	m.xxx_knock_cyl2 = 0
	m.xxx_knock_cyl3 = 0
	m.xxx_knock_cyl4 = 0
	m.xxx_knock_cyl5 = 0
	m.xxx_knock_cyl6 = 0
	m.xxx_knock_cyl7 = 0
	m.xxx_knock_cyl8 = 0
}

func (m *megasquirt_gp48) CopyFrom(o megasquirt_gp48Reader) *megasquirt_gp48 {
	m.Setknock_cyl1(o.knock_cyl1())
	m.Setknock_cyl2(o.knock_cyl2())
	m.Setknock_cyl3(o.knock_cyl3())
	m.Setknock_cyl4(o.knock_cyl4())
	m.Setknock_cyl5(o.knock_cyl5())
	m.Setknock_cyl6(o.knock_cyl6())
	m.Setknock_cyl7(o.knock_cyl7())
	m.Setknock_cyl8(o.knock_cyl8())
	return m
}

// Descriptor returns the megasquirt_gp48 descriptor.
func (m *megasquirt_gp48) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp48.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp48) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp48) knock_cyl1() float64 {
	return Messages().megasquirt_gp48.knock_cyl1.ToPhysical(float64(m.xxx_knock_cyl1))
}

func (m *megasquirt_gp48) Setknock_cyl1(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl1 = uint8(Messages().megasquirt_gp48.knock_cyl1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl2() float64 {
	return Messages().megasquirt_gp48.knock_cyl2.ToPhysical(float64(m.xxx_knock_cyl2))
}

func (m *megasquirt_gp48) Setknock_cyl2(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl2 = uint8(Messages().megasquirt_gp48.knock_cyl2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl3() float64 {
	return Messages().megasquirt_gp48.knock_cyl3.ToPhysical(float64(m.xxx_knock_cyl3))
}

func (m *megasquirt_gp48) Setknock_cyl3(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl3 = uint8(Messages().megasquirt_gp48.knock_cyl3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl4() float64 {
	return Messages().megasquirt_gp48.knock_cyl4.ToPhysical(float64(m.xxx_knock_cyl4))
}

func (m *megasquirt_gp48) Setknock_cyl4(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl4 = uint8(Messages().megasquirt_gp48.knock_cyl4.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl5() float64 {
	return Messages().megasquirt_gp48.knock_cyl5.ToPhysical(float64(m.xxx_knock_cyl5))
}

func (m *megasquirt_gp48) Setknock_cyl5(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl5 = uint8(Messages().megasquirt_gp48.knock_cyl5.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl6() float64 {
	return Messages().megasquirt_gp48.knock_cyl6.ToPhysical(float64(m.xxx_knock_cyl6))
}

func (m *megasquirt_gp48) Setknock_cyl6(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl6 = uint8(Messages().megasquirt_gp48.knock_cyl6.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl7() float64 {
	return Messages().megasquirt_gp48.knock_cyl7.ToPhysical(float64(m.xxx_knock_cyl7))
}

func (m *megasquirt_gp48) Setknock_cyl7(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl7 = uint8(Messages().megasquirt_gp48.knock_cyl7.FromPhysical(v))
	return m
}

func (m *megasquirt_gp48) knock_cyl8() float64 {
	return Messages().megasquirt_gp48.knock_cyl8.ToPhysical(float64(m.xxx_knock_cyl8))
}

func (m *megasquirt_gp48) Setknock_cyl8(v float64) *megasquirt_gp48 {
	m.xxx_knock_cyl8 = uint8(Messages().megasquirt_gp48.knock_cyl8.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp48) Frame() can.Frame {
	md := Messages().megasquirt_gp48
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.knock_cyl1.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl1))
	md.knock_cyl2.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl2))
	md.knock_cyl3.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl3))
	md.knock_cyl4.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl4))
	md.knock_cyl5.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl5))
	md.knock_cyl6.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl6))
	md.knock_cyl7.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl7))
	md.knock_cyl8.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl8))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp48) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp48) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp48
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp48: expects ID 1568 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp48: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp48: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp48: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_knock_cyl1 = uint8(md.knock_cyl1.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl2 = uint8(md.knock_cyl2.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl3 = uint8(md.knock_cyl3.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl4 = uint8(md.knock_cyl4.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl5 = uint8(md.knock_cyl5.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl6 = uint8(md.knock_cyl6.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl7 = uint8(md.knock_cyl7.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl8 = uint8(md.knock_cyl8.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp49Reader provides read access to a megasquirt_gp49 message.
type megasquirt_gp49Reader interface {
	// knock_cyl9 returns the physical value of the knock_cyl9 signal.
	knock_cyl9() float64
	// knock_cyl10 returns the physical value of the knock_cyl10 signal.
	knock_cyl10() float64
	// knock_cyl11 returns the physical value of the knock_cyl11 signal.
	knock_cyl11() float64
	// knock_cyl12 returns the physical value of the knock_cyl12 signal.
	knock_cyl12() float64
	// knock_cyl13 returns the physical value of the knock_cyl13 signal.
	knock_cyl13() float64
	// knock_cyl14 returns the physical value of the knock_cyl14 signal.
	knock_cyl14() float64
	// knock_cyl15 returns the physical value of the knock_cyl15 signal.
	knock_cyl15() float64
	// knock_cyl16 returns the physical value of the knock_cyl16 signal.
	knock_cyl16() float64
}

// megasquirt_gp49Writer provides write access to a megasquirt_gp49 message.
type megasquirt_gp49Writer interface {
	// CopyFrom copies all values from megasquirt_gp49Reader.
	CopyFrom(megasquirt_gp49Reader) *megasquirt_gp49
	// Setknock_cyl9 sets the physical value of the knock_cyl9 signal.
	Setknock_cyl9(float64) *megasquirt_gp49
	// Setknock_cyl10 sets the physical value of the knock_cyl10 signal.
	Setknock_cyl10(float64) *megasquirt_gp49
	// Setknock_cyl11 sets the physical value of the knock_cyl11 signal.
	Setknock_cyl11(float64) *megasquirt_gp49
	// Setknock_cyl12 sets the physical value of the knock_cyl12 signal.
	Setknock_cyl12(float64) *megasquirt_gp49
	// Setknock_cyl13 sets the physical value of the knock_cyl13 signal.
	Setknock_cyl13(float64) *megasquirt_gp49
	// Setknock_cyl14 sets the physical value of the knock_cyl14 signal.
	Setknock_cyl14(float64) *megasquirt_gp49
	// Setknock_cyl15 sets the physical value of the knock_cyl15 signal.
	Setknock_cyl15(float64) *megasquirt_gp49
	// Setknock_cyl16 sets the physical value of the knock_cyl16 signal.
	Setknock_cyl16(float64) *megasquirt_gp49
}

type megasquirt_gp49 struct {
	xxx_knock_cyl9  uint8
	xxx_knock_cyl10 uint8
	xxx_knock_cyl11 uint8
	xxx_knock_cyl12 uint8
	xxx_knock_cyl13 uint8
	xxx_knock_cyl14 uint8
	xxx_knock_cyl15 uint8
	xxx_knock_cyl16 uint8
}

func Newmegasquirt_gp49() *megasquirt_gp49 {
	m := &megasquirt_gp49{}
	m.Reset()
	return m
}

func (m *megasquirt_gp49) Reset() {
	m.xxx_knock_cyl9 = 0
	m.xxx_knock_cyl10 = 0
	m.xxx_knock_cyl11 = 0
	m.xxx_knock_cyl12 = 0
	m.xxx_knock_cyl13 = 0
	m.xxx_knock_cyl14 = 0
	m.xxx_knock_cyl15 = 0
	m.xxx_knock_cyl16 = 0
}

func (m *megasquirt_gp49) CopyFrom(o megasquirt_gp49Reader) *megasquirt_gp49 {
	m.Setknock_cyl9(o.knock_cyl9())
	m.Setknock_cyl10(o.knock_cyl10())
	m.Setknock_cyl11(o.knock_cyl11())
	m.Setknock_cyl12(o.knock_cyl12())
	m.Setknock_cyl13(o.knock_cyl13())
	m.Setknock_cyl14(o.knock_cyl14())
	m.Setknock_cyl15(o.knock_cyl15())
	m.Setknock_cyl16(o.knock_cyl16())
	return m
}

// Descriptor returns the megasquirt_gp49 descriptor.
func (m *megasquirt_gp49) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp49.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp49) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp49) knock_cyl9() float64 {
	return Messages().megasquirt_gp49.knock_cyl9.ToPhysical(float64(m.xxx_knock_cyl9))
}

func (m *megasquirt_gp49) Setknock_cyl9(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl9 = uint8(Messages().megasquirt_gp49.knock_cyl9.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl10() float64 {
	return Messages().megasquirt_gp49.knock_cyl10.ToPhysical(float64(m.xxx_knock_cyl10))
}

func (m *megasquirt_gp49) Setknock_cyl10(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl10 = uint8(Messages().megasquirt_gp49.knock_cyl10.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl11() float64 {
	return Messages().megasquirt_gp49.knock_cyl11.ToPhysical(float64(m.xxx_knock_cyl11))
}

func (m *megasquirt_gp49) Setknock_cyl11(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl11 = uint8(Messages().megasquirt_gp49.knock_cyl11.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl12() float64 {
	return Messages().megasquirt_gp49.knock_cyl12.ToPhysical(float64(m.xxx_knock_cyl12))
}

func (m *megasquirt_gp49) Setknock_cyl12(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl12 = uint8(Messages().megasquirt_gp49.knock_cyl12.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl13() float64 {
	return Messages().megasquirt_gp49.knock_cyl13.ToPhysical(float64(m.xxx_knock_cyl13))
}

func (m *megasquirt_gp49) Setknock_cyl13(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl13 = uint8(Messages().megasquirt_gp49.knock_cyl13.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl14() float64 {
	return Messages().megasquirt_gp49.knock_cyl14.ToPhysical(float64(m.xxx_knock_cyl14))
}

func (m *megasquirt_gp49) Setknock_cyl14(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl14 = uint8(Messages().megasquirt_gp49.knock_cyl14.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl15() float64 {
	return Messages().megasquirt_gp49.knock_cyl15.ToPhysical(float64(m.xxx_knock_cyl15))
}

func (m *megasquirt_gp49) Setknock_cyl15(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl15 = uint8(Messages().megasquirt_gp49.knock_cyl15.FromPhysical(v))
	return m
}

func (m *megasquirt_gp49) knock_cyl16() float64 {
	return Messages().megasquirt_gp49.knock_cyl16.ToPhysical(float64(m.xxx_knock_cyl16))
}

func (m *megasquirt_gp49) Setknock_cyl16(v float64) *megasquirt_gp49 {
	m.xxx_knock_cyl16 = uint8(Messages().megasquirt_gp49.knock_cyl16.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp49) Frame() can.Frame {
	md := Messages().megasquirt_gp49
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.knock_cyl9.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl9))
	md.knock_cyl10.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl10))
	md.knock_cyl11.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl11))
	md.knock_cyl12.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl12))
	md.knock_cyl13.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl13))
	md.knock_cyl14.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl14))
	md.knock_cyl15.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl15))
	md.knock_cyl16.MarshalUnsigned(&f.Data, uint64(m.xxx_knock_cyl16))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp49) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp49) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp49
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp49: expects ID 1569 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp49: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp49: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp49: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_knock_cyl9 = uint8(md.knock_cyl9.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl10 = uint8(md.knock_cyl10.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl11 = uint8(md.knock_cyl11.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl12 = uint8(md.knock_cyl12.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl13 = uint8(md.knock_cyl13.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl14 = uint8(md.knock_cyl14.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl15 = uint8(md.knock_cyl15.UnmarshalUnsigned(f.Data))
	m.xxx_knock_cyl16 = uint8(md.knock_cyl16.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp50Reader provides read access to a megasquirt_gp50 message.
type megasquirt_gp50Reader interface {
	// map_accel returns the physical value of the map_accel signal.
	map_accel() float64
	// total_accel returns the physical value of the total_accel signal.
	total_accel() float64
	// launch_timer returns the physical value of the launch_timer signal.
	launch_timer() float64
	// launch_retard returns the physical value of the launch_retard signal.
	launch_retard() float64
}

// megasquirt_gp50Writer provides write access to a megasquirt_gp50 message.
type megasquirt_gp50Writer interface {
	// CopyFrom copies all values from megasquirt_gp50Reader.
	CopyFrom(megasquirt_gp50Reader) *megasquirt_gp50
	// Setmap_accel sets the physical value of the map_accel signal.
	Setmap_accel(float64) *megasquirt_gp50
	// Settotal_accel sets the physical value of the total_accel signal.
	Settotal_accel(float64) *megasquirt_gp50
	// Setlaunch_timer sets the physical value of the launch_timer signal.
	Setlaunch_timer(float64) *megasquirt_gp50
	// Setlaunch_retard sets the physical value of the launch_retard signal.
	Setlaunch_retard(float64) *megasquirt_gp50
}

type megasquirt_gp50 struct {
	xxx_map_accel     int16
	xxx_total_accel   int16
	xxx_launch_timer  uint16
	xxx_launch_retard int16
}

func Newmegasquirt_gp50() *megasquirt_gp50 {
	m := &megasquirt_gp50{}
	m.Reset()
	return m
}

func (m *megasquirt_gp50) Reset() {
	m.xxx_map_accel = 0
	m.xxx_total_accel = 0
	m.xxx_launch_timer = 0
	m.xxx_launch_retard = 0
}

func (m *megasquirt_gp50) CopyFrom(o megasquirt_gp50Reader) *megasquirt_gp50 {
	m.Setmap_accel(o.map_accel())
	m.Settotal_accel(o.total_accel())
	m.Setlaunch_timer(o.launch_timer())
	m.Setlaunch_retard(o.launch_retard())
	return m
}

// Descriptor returns the megasquirt_gp50 descriptor.
func (m *megasquirt_gp50) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp50.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp50) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp50) map_accel() float64 {
	return Messages().megasquirt_gp50.map_accel.ToPhysical(float64(m.xxx_map_accel))
}

func (m *megasquirt_gp50) Setmap_accel(v float64) *megasquirt_gp50 {
	m.xxx_map_accel = int16(Messages().megasquirt_gp50.map_accel.FromPhysical(v))
	return m
}

func (m *megasquirt_gp50) total_accel() float64 {
	return Messages().megasquirt_gp50.total_accel.ToPhysical(float64(m.xxx_total_accel))
}

func (m *megasquirt_gp50) Settotal_accel(v float64) *megasquirt_gp50 {
	m.xxx_total_accel = int16(Messages().megasquirt_gp50.total_accel.FromPhysical(v))
	return m
}

func (m *megasquirt_gp50) launch_timer() float64 {
	return Messages().megasquirt_gp50.launch_timer.ToPhysical(float64(m.xxx_launch_timer))
}

func (m *megasquirt_gp50) Setlaunch_timer(v float64) *megasquirt_gp50 {
	m.xxx_launch_timer = uint16(Messages().megasquirt_gp50.launch_timer.FromPhysical(v))
	return m
}

func (m *megasquirt_gp50) launch_retard() float64 {
	return Messages().megasquirt_gp50.launch_retard.ToPhysical(float64(m.xxx_launch_retard))
}

func (m *megasquirt_gp50) Setlaunch_retard(v float64) *megasquirt_gp50 {
	m.xxx_launch_retard = int16(Messages().megasquirt_gp50.launch_retard.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp50) Frame() can.Frame {
	md := Messages().megasquirt_gp50
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.map_accel.MarshalSigned(&f.Data, int64(m.xxx_map_accel))
	md.total_accel.MarshalSigned(&f.Data, int64(m.xxx_total_accel))
	md.launch_timer.MarshalUnsigned(&f.Data, uint64(m.xxx_launch_timer))
	md.launch_retard.MarshalSigned(&f.Data, int64(m.xxx_launch_retard))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp50) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp50) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp50
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp50: expects ID 1570 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp50: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp50: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp50: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_map_accel = int16(md.map_accel.UnmarshalSigned(f.Data))
	m.xxx_total_accel = int16(md.total_accel.UnmarshalSigned(f.Data))
	m.xxx_launch_timer = uint16(md.launch_timer.UnmarshalUnsigned(f.Data))
	m.xxx_launch_retard = int16(md.launch_retard.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp51Reader provides read access to a megasquirt_gp51 message.
type megasquirt_gp51Reader interface {
	// porta returns the value of the porta signal.
	porta() uint8
	// portb returns the value of the portb signal.
	portb() uint8
	// porteh returns the value of the porteh signal.
	porteh() uint8
	// portk returns the value of the portk signal.
	portk() uint8
	// portmj returns the value of the portmj signal.
	portmj() uint8
	// portp returns the value of the portp signal.
	portp() uint8
	// portt returns the value of the portt signal.
	portt() uint8
	// cel_errorcode returns the value of the cel_errorcode signal.
	cel_errorcode() uint8
}

// megasquirt_gp51Writer provides write access to a megasquirt_gp51 message.
type megasquirt_gp51Writer interface {
	// CopyFrom copies all values from megasquirt_gp51Reader.
	CopyFrom(megasquirt_gp51Reader) *megasquirt_gp51
	// Setporta sets the value of the porta signal.
	Setporta(uint8) *megasquirt_gp51
	// Setportb sets the value of the portb signal.
	Setportb(uint8) *megasquirt_gp51
	// Setporteh sets the value of the porteh signal.
	Setporteh(uint8) *megasquirt_gp51
	// Setportk sets the value of the portk signal.
	Setportk(uint8) *megasquirt_gp51
	// Setportmj sets the value of the portmj signal.
	Setportmj(uint8) *megasquirt_gp51
	// Setportp sets the value of the portp signal.
	Setportp(uint8) *megasquirt_gp51
	// Setportt sets the value of the portt signal.
	Setportt(uint8) *megasquirt_gp51
	// Setcel_errorcode sets the value of the cel_errorcode signal.
	Setcel_errorcode(uint8) *megasquirt_gp51
}

type megasquirt_gp51 struct {
	xxx_porta         uint8
	xxx_portb         uint8
	xxx_porteh        uint8
	xxx_portk         uint8
	xxx_portmj        uint8
	xxx_portp         uint8
	xxx_portt         uint8
	xxx_cel_errorcode uint8
}

func Newmegasquirt_gp51() *megasquirt_gp51 {
	m := &megasquirt_gp51{}
	m.Reset()
	return m
}

func (m *megasquirt_gp51) Reset() {
	m.xxx_porta = 0
	m.xxx_portb = 0
	m.xxx_porteh = 0
	m.xxx_portk = 0
	m.xxx_portmj = 0
	m.xxx_portp = 0
	m.xxx_portt = 0
	m.xxx_cel_errorcode = 0
}

func (m *megasquirt_gp51) CopyFrom(o megasquirt_gp51Reader) *megasquirt_gp51 {
	m.xxx_porta = o.porta()
	m.xxx_portb = o.portb()
	m.xxx_porteh = o.porteh()
	m.xxx_portk = o.portk()
	m.xxx_portmj = o.portmj()
	m.xxx_portp = o.portp()
	m.xxx_portt = o.portt()
	m.xxx_cel_errorcode = o.cel_errorcode()
	return m
}

// Descriptor returns the megasquirt_gp51 descriptor.
func (m *megasquirt_gp51) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp51.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp51) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp51) porta() uint8 {
	return m.xxx_porta
}

func (m *megasquirt_gp51) Setporta(v uint8) *megasquirt_gp51 {
	m.xxx_porta = uint8(Messages().megasquirt_gp51.porta.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) portb() uint8 {
	return m.xxx_portb
}

func (m *megasquirt_gp51) Setportb(v uint8) *megasquirt_gp51 {
	m.xxx_portb = uint8(Messages().megasquirt_gp51.portb.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) porteh() uint8 {
	return m.xxx_porteh
}

func (m *megasquirt_gp51) Setporteh(v uint8) *megasquirt_gp51 {
	m.xxx_porteh = uint8(Messages().megasquirt_gp51.porteh.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) portk() uint8 {
	return m.xxx_portk
}

func (m *megasquirt_gp51) Setportk(v uint8) *megasquirt_gp51 {
	m.xxx_portk = uint8(Messages().megasquirt_gp51.portk.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) portmj() uint8 {
	return m.xxx_portmj
}

func (m *megasquirt_gp51) Setportmj(v uint8) *megasquirt_gp51 {
	m.xxx_portmj = uint8(Messages().megasquirt_gp51.portmj.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) portp() uint8 {
	return m.xxx_portp
}

func (m *megasquirt_gp51) Setportp(v uint8) *megasquirt_gp51 {
	m.xxx_portp = uint8(Messages().megasquirt_gp51.portp.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) portt() uint8 {
	return m.xxx_portt
}

func (m *megasquirt_gp51) Setportt(v uint8) *megasquirt_gp51 {
	m.xxx_portt = uint8(Messages().megasquirt_gp51.portt.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp51) cel_errorcode() uint8 {
	return m.xxx_cel_errorcode
}

func (m *megasquirt_gp51) Setcel_errorcode(v uint8) *megasquirt_gp51 {
	m.xxx_cel_errorcode = uint8(Messages().megasquirt_gp51.cel_errorcode.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp51) Frame() can.Frame {
	md := Messages().megasquirt_gp51
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.porta.MarshalUnsigned(&f.Data, uint64(m.xxx_porta))
	md.portb.MarshalUnsigned(&f.Data, uint64(m.xxx_portb))
	md.porteh.MarshalUnsigned(&f.Data, uint64(m.xxx_porteh))
	md.portk.MarshalUnsigned(&f.Data, uint64(m.xxx_portk))
	md.portmj.MarshalUnsigned(&f.Data, uint64(m.xxx_portmj))
	md.portp.MarshalUnsigned(&f.Data, uint64(m.xxx_portp))
	md.portt.MarshalUnsigned(&f.Data, uint64(m.xxx_portt))
	md.cel_errorcode.MarshalUnsigned(&f.Data, uint64(m.xxx_cel_errorcode))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp51) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp51) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp51
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp51: expects ID 1571 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp51: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp51: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp51: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_porta = uint8(md.porta.UnmarshalUnsigned(f.Data))
	m.xxx_portb = uint8(md.portb.UnmarshalUnsigned(f.Data))
	m.xxx_porteh = uint8(md.porteh.UnmarshalUnsigned(f.Data))
	m.xxx_portk = uint8(md.portk.UnmarshalUnsigned(f.Data))
	m.xxx_portmj = uint8(md.portmj.UnmarshalUnsigned(f.Data))
	m.xxx_portp = uint8(md.portp.UnmarshalUnsigned(f.Data))
	m.xxx_portt = uint8(md.portt.UnmarshalUnsigned(f.Data))
	m.xxx_cel_errorcode = uint8(md.cel_errorcode.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp52Reader provides read access to a megasquirt_gp52 message.
type megasquirt_gp52Reader interface {
	// canin1 returns the value of the canin1 signal.
	canin1() uint8
	// canin2 returns the value of the canin2 signal.
	canin2() uint8
	// canout returns the value of the canout signal.
	canout() uint8
	// knk_rtd returns the physical value of the knk_rtd signal.
	knk_rtd() float64
	// fuelflow returns the value of the fuelflow signal.
	fuelflow() uint16
	// fuelcons returns the value of the fuelcons signal.
	fuelcons() uint16
}

// megasquirt_gp52Writer provides write access to a megasquirt_gp52 message.
type megasquirt_gp52Writer interface {
	// CopyFrom copies all values from megasquirt_gp52Reader.
	CopyFrom(megasquirt_gp52Reader) *megasquirt_gp52
	// Setcanin1 sets the value of the canin1 signal.
	Setcanin1(uint8) *megasquirt_gp52
	// Setcanin2 sets the value of the canin2 signal.
	Setcanin2(uint8) *megasquirt_gp52
	// Setcanout sets the value of the canout signal.
	Setcanout(uint8) *megasquirt_gp52
	// Setknk_rtd sets the physical value of the knk_rtd signal.
	Setknk_rtd(float64) *megasquirt_gp52
	// Setfuelflow sets the value of the fuelflow signal.
	Setfuelflow(uint16) *megasquirt_gp52
	// Setfuelcons sets the value of the fuelcons signal.
	Setfuelcons(uint16) *megasquirt_gp52
}

type megasquirt_gp52 struct {
	xxx_canin1   uint8
	xxx_canin2   uint8
	xxx_canout   uint8
	xxx_knk_rtd  uint8
	xxx_fuelflow uint16
	xxx_fuelcons uint16
}

func Newmegasquirt_gp52() *megasquirt_gp52 {
	m := &megasquirt_gp52{}
	m.Reset()
	return m
}

func (m *megasquirt_gp52) Reset() {
	m.xxx_canin1 = 0
	m.xxx_canin2 = 0
	m.xxx_canout = 0
	m.xxx_knk_rtd = 0
	m.xxx_fuelflow = 0
	m.xxx_fuelcons = 0
}

func (m *megasquirt_gp52) CopyFrom(o megasquirt_gp52Reader) *megasquirt_gp52 {
	m.xxx_canin1 = o.canin1()
	m.xxx_canin2 = o.canin2()
	m.xxx_canout = o.canout()
	m.Setknk_rtd(o.knk_rtd())
	m.xxx_fuelflow = o.fuelflow()
	m.xxx_fuelcons = o.fuelcons()
	return m
}

// Descriptor returns the megasquirt_gp52 descriptor.
func (m *megasquirt_gp52) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp52.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp52) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp52) canin1() uint8 {
	return m.xxx_canin1
}

func (m *megasquirt_gp52) Setcanin1(v uint8) *megasquirt_gp52 {
	m.xxx_canin1 = uint8(Messages().megasquirt_gp52.canin1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp52) canin2() uint8 {
	return m.xxx_canin2
}

func (m *megasquirt_gp52) Setcanin2(v uint8) *megasquirt_gp52 {
	m.xxx_canin2 = uint8(Messages().megasquirt_gp52.canin2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp52) canout() uint8 {
	return m.xxx_canout
}

func (m *megasquirt_gp52) Setcanout(v uint8) *megasquirt_gp52 {
	m.xxx_canout = uint8(Messages().megasquirt_gp52.canout.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp52) knk_rtd() float64 {
	return Messages().megasquirt_gp52.knk_rtd.ToPhysical(float64(m.xxx_knk_rtd))
}

func (m *megasquirt_gp52) Setknk_rtd(v float64) *megasquirt_gp52 {
	m.xxx_knk_rtd = uint8(Messages().megasquirt_gp52.knk_rtd.FromPhysical(v))
	return m
}

func (m *megasquirt_gp52) fuelflow() uint16 {
	return m.xxx_fuelflow
}

func (m *megasquirt_gp52) Setfuelflow(v uint16) *megasquirt_gp52 {
	m.xxx_fuelflow = uint16(Messages().megasquirt_gp52.fuelflow.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp52) fuelcons() uint16 {
	return m.xxx_fuelcons
}

func (m *megasquirt_gp52) Setfuelcons(v uint16) *megasquirt_gp52 {
	m.xxx_fuelcons = uint16(Messages().megasquirt_gp52.fuelcons.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp52) Frame() can.Frame {
	md := Messages().megasquirt_gp52
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.canin1.MarshalUnsigned(&f.Data, uint64(m.xxx_canin1))
	md.canin2.MarshalUnsigned(&f.Data, uint64(m.xxx_canin2))
	md.canout.MarshalUnsigned(&f.Data, uint64(m.xxx_canout))
	md.knk_rtd.MarshalUnsigned(&f.Data, uint64(m.xxx_knk_rtd))
	md.fuelflow.MarshalUnsigned(&f.Data, uint64(m.xxx_fuelflow))
	md.fuelcons.MarshalUnsigned(&f.Data, uint64(m.xxx_fuelcons))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp52) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp52) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp52
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp52: expects ID 1572 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp52: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp52: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp52: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_canin1 = uint8(md.canin1.UnmarshalUnsigned(f.Data))
	m.xxx_canin2 = uint8(md.canin2.UnmarshalUnsigned(f.Data))
	m.xxx_canout = uint8(md.canout.UnmarshalUnsigned(f.Data))
	m.xxx_knk_rtd = uint8(md.knk_rtd.UnmarshalUnsigned(f.Data))
	m.xxx_fuelflow = uint16(md.fuelflow.UnmarshalUnsigned(f.Data))
	m.xxx_fuelcons = uint16(md.fuelcons.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp53Reader provides read access to a megasquirt_gp53 message.
type megasquirt_gp53Reader interface {
	// fuel_press1 returns the physical value of the fuel_press1 signal.
	fuel_press1() float64
	// fuel_press2 returns the physical value of the fuel_press2 signal.
	fuel_press2() float64
	// fuel_temp1 returns the physical value of the fuel_temp1 signal.
	fuel_temp1() float64
	// fuel_temp2 returns the physical value of the fuel_temp2 signal.
	fuel_temp2() float64
}

// megasquirt_gp53Writer provides write access to a megasquirt_gp53 message.
type megasquirt_gp53Writer interface {
	// CopyFrom copies all values from megasquirt_gp53Reader.
	CopyFrom(megasquirt_gp53Reader) *megasquirt_gp53
	// Setfuel_press1 sets the physical value of the fuel_press1 signal.
	Setfuel_press1(float64) *megasquirt_gp53
	// Setfuel_press2 sets the physical value of the fuel_press2 signal.
	Setfuel_press2(float64) *megasquirt_gp53
	// Setfuel_temp1 sets the physical value of the fuel_temp1 signal.
	Setfuel_temp1(float64) *megasquirt_gp53
	// Setfuel_temp2 sets the physical value of the fuel_temp2 signal.
	Setfuel_temp2(float64) *megasquirt_gp53
}

type megasquirt_gp53 struct {
	xxx_fuel_press1 int16
	xxx_fuel_press2 int16
	xxx_fuel_temp1  int16
	xxx_fuel_temp2  int16
}

func Newmegasquirt_gp53() *megasquirt_gp53 {
	m := &megasquirt_gp53{}
	m.Reset()
	return m
}

func (m *megasquirt_gp53) Reset() {
	m.xxx_fuel_press1 = 0
	m.xxx_fuel_press2 = 0
	m.xxx_fuel_temp1 = 0
	m.xxx_fuel_temp2 = 0
}

func (m *megasquirt_gp53) CopyFrom(o megasquirt_gp53Reader) *megasquirt_gp53 {
	m.Setfuel_press1(o.fuel_press1())
	m.Setfuel_press2(o.fuel_press2())
	m.Setfuel_temp1(o.fuel_temp1())
	m.Setfuel_temp2(o.fuel_temp2())
	return m
}

// Descriptor returns the megasquirt_gp53 descriptor.
func (m *megasquirt_gp53) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp53.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp53) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp53) fuel_press1() float64 {
	return Messages().megasquirt_gp53.fuel_press1.ToPhysical(float64(m.xxx_fuel_press1))
}

func (m *megasquirt_gp53) Setfuel_press1(v float64) *megasquirt_gp53 {
	m.xxx_fuel_press1 = int16(Messages().megasquirt_gp53.fuel_press1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp53) fuel_press2() float64 {
	return Messages().megasquirt_gp53.fuel_press2.ToPhysical(float64(m.xxx_fuel_press2))
}

func (m *megasquirt_gp53) Setfuel_press2(v float64) *megasquirt_gp53 {
	m.xxx_fuel_press2 = int16(Messages().megasquirt_gp53.fuel_press2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp53) fuel_temp1() float64 {
	return Messages().megasquirt_gp53.fuel_temp1.ToPhysical(float64(m.xxx_fuel_temp1))
}

func (m *megasquirt_gp53) Setfuel_temp1(v float64) *megasquirt_gp53 {
	m.xxx_fuel_temp1 = int16(Messages().megasquirt_gp53.fuel_temp1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp53) fuel_temp2() float64 {
	return Messages().megasquirt_gp53.fuel_temp2.ToPhysical(float64(m.xxx_fuel_temp2))
}

func (m *megasquirt_gp53) Setfuel_temp2(v float64) *megasquirt_gp53 {
	m.xxx_fuel_temp2 = int16(Messages().megasquirt_gp53.fuel_temp2.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp53) Frame() can.Frame {
	md := Messages().megasquirt_gp53
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.fuel_press1.MarshalSigned(&f.Data, int64(m.xxx_fuel_press1))
	md.fuel_press2.MarshalSigned(&f.Data, int64(m.xxx_fuel_press2))
	md.fuel_temp1.MarshalSigned(&f.Data, int64(m.xxx_fuel_temp1))
	md.fuel_temp2.MarshalSigned(&f.Data, int64(m.xxx_fuel_temp2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp53) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp53) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp53
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp53: expects ID 1573 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp53: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp53: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp53: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_fuel_press1 = int16(md.fuel_press1.UnmarshalSigned(f.Data))
	m.xxx_fuel_press2 = int16(md.fuel_press2.UnmarshalSigned(f.Data))
	m.xxx_fuel_temp1 = int16(md.fuel_temp1.UnmarshalSigned(f.Data))
	m.xxx_fuel_temp2 = int16(md.fuel_temp2.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp54Reader provides read access to a megasquirt_gp54 message.
type megasquirt_gp54Reader interface {
	// batt_cur returns the physical value of the batt_cur signal.
	batt_cur() float64
	// cel_status returns the value of the cel_status signal.
	cel_status() uint16
	// fp_duty returns the physical value of the fp_duty signal.
	fp_duty() float64
	// alt_duty returns the value of the alt_duty signal.
	alt_duty() uint8
	// load_duty returns the value of the load_duty signal.
	load_duty() uint8
	// alt_targv returns the physical value of the alt_targv signal.
	alt_targv() float64
}

// megasquirt_gp54Writer provides write access to a megasquirt_gp54 message.
type megasquirt_gp54Writer interface {
	// CopyFrom copies all values from megasquirt_gp54Reader.
	CopyFrom(megasquirt_gp54Reader) *megasquirt_gp54
	// Setbatt_cur sets the physical value of the batt_cur signal.
	Setbatt_cur(float64) *megasquirt_gp54
	// Setcel_status sets the value of the cel_status signal.
	Setcel_status(uint16) *megasquirt_gp54
	// Setfp_duty sets the physical value of the fp_duty signal.
	Setfp_duty(float64) *megasquirt_gp54
	// Setalt_duty sets the value of the alt_duty signal.
	Setalt_duty(uint8) *megasquirt_gp54
	// Setload_duty sets the value of the load_duty signal.
	Setload_duty(uint8) *megasquirt_gp54
	// Setalt_targv sets the physical value of the alt_targv signal.
	Setalt_targv(float64) *megasquirt_gp54
}

type megasquirt_gp54 struct {
	xxx_batt_cur   int16
	xxx_cel_status uint16
	xxx_fp_duty    uint8
	xxx_alt_duty   uint8
	xxx_load_duty  uint8
	xxx_alt_targv  uint8
}

func Newmegasquirt_gp54() *megasquirt_gp54 {
	m := &megasquirt_gp54{}
	m.Reset()
	return m
}

func (m *megasquirt_gp54) Reset() {
	m.xxx_batt_cur = 0
	m.xxx_cel_status = 0
	m.xxx_fp_duty = 0
	m.xxx_alt_duty = 0
	m.xxx_load_duty = 0
	m.xxx_alt_targv = 0
}

func (m *megasquirt_gp54) CopyFrom(o megasquirt_gp54Reader) *megasquirt_gp54 {
	m.Setbatt_cur(o.batt_cur())
	m.xxx_cel_status = o.cel_status()
	m.Setfp_duty(o.fp_duty())
	m.xxx_alt_duty = o.alt_duty()
	m.xxx_load_duty = o.load_duty()
	m.Setalt_targv(o.alt_targv())
	return m
}

// Descriptor returns the megasquirt_gp54 descriptor.
func (m *megasquirt_gp54) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp54.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp54) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp54) batt_cur() float64 {
	return Messages().megasquirt_gp54.batt_cur.ToPhysical(float64(m.xxx_batt_cur))
}

func (m *megasquirt_gp54) Setbatt_cur(v float64) *megasquirt_gp54 {
	m.xxx_batt_cur = int16(Messages().megasquirt_gp54.batt_cur.FromPhysical(v))
	return m
}

func (m *megasquirt_gp54) cel_status() uint16 {
	return m.xxx_cel_status
}

func (m *megasquirt_gp54) Setcel_status(v uint16) *megasquirt_gp54 {
	m.xxx_cel_status = uint16(Messages().megasquirt_gp54.cel_status.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp54) fp_duty() float64 {
	return Messages().megasquirt_gp54.fp_duty.ToPhysical(float64(m.xxx_fp_duty))
}

func (m *megasquirt_gp54) Setfp_duty(v float64) *megasquirt_gp54 {
	m.xxx_fp_duty = uint8(Messages().megasquirt_gp54.fp_duty.FromPhysical(v))
	return m
}

func (m *megasquirt_gp54) alt_duty() uint8 {
	return m.xxx_alt_duty
}

func (m *megasquirt_gp54) Setalt_duty(v uint8) *megasquirt_gp54 {
	m.xxx_alt_duty = uint8(Messages().megasquirt_gp54.alt_duty.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp54) load_duty() uint8 {
	return m.xxx_load_duty
}

func (m *megasquirt_gp54) Setload_duty(v uint8) *megasquirt_gp54 {
	m.xxx_load_duty = uint8(Messages().megasquirt_gp54.load_duty.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp54) alt_targv() float64 {
	return Messages().megasquirt_gp54.alt_targv.ToPhysical(float64(m.xxx_alt_targv))
}

func (m *megasquirt_gp54) Setalt_targv(v float64) *megasquirt_gp54 {
	m.xxx_alt_targv = uint8(Messages().megasquirt_gp54.alt_targv.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp54) Frame() can.Frame {
	md := Messages().megasquirt_gp54
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.batt_cur.MarshalSigned(&f.Data, int64(m.xxx_batt_cur))
	md.cel_status.MarshalUnsigned(&f.Data, uint64(m.xxx_cel_status))
	md.fp_duty.MarshalUnsigned(&f.Data, uint64(m.xxx_fp_duty))
	md.alt_duty.MarshalUnsigned(&f.Data, uint64(m.xxx_alt_duty))
	md.load_duty.MarshalUnsigned(&f.Data, uint64(m.xxx_load_duty))
	md.alt_targv.MarshalUnsigned(&f.Data, uint64(m.xxx_alt_targv))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp54) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp54) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp54
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp54: expects ID 1574 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp54: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp54: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp54: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_batt_cur = int16(md.batt_cur.UnmarshalSigned(f.Data))
	m.xxx_cel_status = uint16(md.cel_status.UnmarshalUnsigned(f.Data))
	m.xxx_fp_duty = uint8(md.fp_duty.UnmarshalUnsigned(f.Data))
	m.xxx_alt_duty = uint8(md.alt_duty.UnmarshalUnsigned(f.Data))
	m.xxx_load_duty = uint8(md.load_duty.UnmarshalUnsigned(f.Data))
	m.xxx_alt_targv = uint8(md.alt_targv.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp55Reader provides read access to a megasquirt_gp55 message.
type megasquirt_gp55Reader interface {
	// looptime returns the value of the looptime signal.
	looptime() uint16
	// fueltemp_cor returns the physical value of the fueltemp_cor signal.
	fueltemp_cor() float64
	// fuelpress_cor returns the physical value of the fuelpress_cor signal.
	fuelpress_cor() float64
	// ltt_cor returns the physical value of the ltt_cor signal.
	ltt_cor() float64
	// sp1 returns the value of the sp1 signal.
	sp1() uint8
}

// megasquirt_gp55Writer provides write access to a megasquirt_gp55 message.
type megasquirt_gp55Writer interface {
	// CopyFrom copies all values from megasquirt_gp55Reader.
	CopyFrom(megasquirt_gp55Reader) *megasquirt_gp55
	// Setlooptime sets the value of the looptime signal.
	Setlooptime(uint16) *megasquirt_gp55
	// Setfueltemp_cor sets the physical value of the fueltemp_cor signal.
	Setfueltemp_cor(float64) *megasquirt_gp55
	// Setfuelpress_cor sets the physical value of the fuelpress_cor signal.
	Setfuelpress_cor(float64) *megasquirt_gp55
	// Setltt_cor sets the physical value of the ltt_cor signal.
	Setltt_cor(float64) *megasquirt_gp55
	// Setsp1 sets the value of the sp1 signal.
	Setsp1(uint8) *megasquirt_gp55
}

type megasquirt_gp55 struct {
	xxx_looptime      uint16
	xxx_fueltemp_cor  uint16
	xxx_fuelpress_cor uint16
	xxx_ltt_cor       int8
	xxx_sp1           uint8
}

func Newmegasquirt_gp55() *megasquirt_gp55 {
	m := &megasquirt_gp55{}
	m.Reset()
	return m
}

func (m *megasquirt_gp55) Reset() {
	m.xxx_looptime = 0
	m.xxx_fueltemp_cor = 0
	m.xxx_fuelpress_cor = 0
	m.xxx_ltt_cor = 0
	m.xxx_sp1 = 0
}

func (m *megasquirt_gp55) CopyFrom(o megasquirt_gp55Reader) *megasquirt_gp55 {
	m.xxx_looptime = o.looptime()
	m.Setfueltemp_cor(o.fueltemp_cor())
	m.Setfuelpress_cor(o.fuelpress_cor())
	m.Setltt_cor(o.ltt_cor())
	m.xxx_sp1 = o.sp1()
	return m
}

// Descriptor returns the megasquirt_gp55 descriptor.
func (m *megasquirt_gp55) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp55.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp55) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp55) looptime() uint16 {
	return m.xxx_looptime
}

func (m *megasquirt_gp55) Setlooptime(v uint16) *megasquirt_gp55 {
	m.xxx_looptime = uint16(Messages().megasquirt_gp55.looptime.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp55) fueltemp_cor() float64 {
	return Messages().megasquirt_gp55.fueltemp_cor.ToPhysical(float64(m.xxx_fueltemp_cor))
}

func (m *megasquirt_gp55) Setfueltemp_cor(v float64) *megasquirt_gp55 {
	m.xxx_fueltemp_cor = uint16(Messages().megasquirt_gp55.fueltemp_cor.FromPhysical(v))
	return m
}

func (m *megasquirt_gp55) fuelpress_cor() float64 {
	return Messages().megasquirt_gp55.fuelpress_cor.ToPhysical(float64(m.xxx_fuelpress_cor))
}

func (m *megasquirt_gp55) Setfuelpress_cor(v float64) *megasquirt_gp55 {
	m.xxx_fuelpress_cor = uint16(Messages().megasquirt_gp55.fuelpress_cor.FromPhysical(v))
	return m
}

func (m *megasquirt_gp55) ltt_cor() float64 {
	return Messages().megasquirt_gp55.ltt_cor.ToPhysical(float64(m.xxx_ltt_cor))
}

func (m *megasquirt_gp55) Setltt_cor(v float64) *megasquirt_gp55 {
	m.xxx_ltt_cor = int8(Messages().megasquirt_gp55.ltt_cor.FromPhysical(v))
	return m
}

func (m *megasquirt_gp55) sp1() uint8 {
	return m.xxx_sp1
}

func (m *megasquirt_gp55) Setsp1(v uint8) *megasquirt_gp55 {
	m.xxx_sp1 = uint8(Messages().megasquirt_gp55.sp1.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp55) Frame() can.Frame {
	md := Messages().megasquirt_gp55
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.looptime.MarshalUnsigned(&f.Data, uint64(m.xxx_looptime))
	md.fueltemp_cor.MarshalUnsigned(&f.Data, uint64(m.xxx_fueltemp_cor))
	md.fuelpress_cor.MarshalUnsigned(&f.Data, uint64(m.xxx_fuelpress_cor))
	md.ltt_cor.MarshalSigned(&f.Data, int64(m.xxx_ltt_cor))
	md.sp1.MarshalUnsigned(&f.Data, uint64(m.xxx_sp1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp55) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp55) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp55
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp55: expects ID 1575 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp55: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp55: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp55: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_looptime = uint16(md.looptime.UnmarshalUnsigned(f.Data))
	m.xxx_fueltemp_cor = uint16(md.fueltemp_cor.UnmarshalUnsigned(f.Data))
	m.xxx_fuelpress_cor = uint16(md.fuelpress_cor.UnmarshalUnsigned(f.Data))
	m.xxx_ltt_cor = int8(md.ltt_cor.UnmarshalSigned(f.Data))
	m.xxx_sp1 = uint8(md.sp1.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp56Reader provides read access to a megasquirt_gp56 message.
type megasquirt_gp56Reader interface {
	// tc_retard returns the physical value of the tc_retard signal.
	tc_retard() float64
	// cel_retard returns the physical value of the cel_retard signal.
	cel_retard() float64
	// fc_retard returns the physical value of the fc_retard signal.
	fc_retard() float64
	// als_addfuel returns the physical value of the als_addfuel signal.
	als_addfuel() float64
}

// megasquirt_gp56Writer provides write access to a megasquirt_gp56 message.
type megasquirt_gp56Writer interface {
	// CopyFrom copies all values from megasquirt_gp56Reader.
	CopyFrom(megasquirt_gp56Reader) *megasquirt_gp56
	// Settc_retard sets the physical value of the tc_retard signal.
	Settc_retard(float64) *megasquirt_gp56
	// Setcel_retard sets the physical value of the cel_retard signal.
	Setcel_retard(float64) *megasquirt_gp56
	// Setfc_retard sets the physical value of the fc_retard signal.
	Setfc_retard(float64) *megasquirt_gp56
	// Setals_addfuel sets the physical value of the als_addfuel signal.
	Setals_addfuel(float64) *megasquirt_gp56
}

type megasquirt_gp56 struct {
	xxx_tc_retard   int16
	xxx_cel_retard  int16
	xxx_fc_retard   int16
	xxx_als_addfuel int16
}

func Newmegasquirt_gp56() *megasquirt_gp56 {
	m := &megasquirt_gp56{}
	m.Reset()
	return m
}

func (m *megasquirt_gp56) Reset() {
	m.xxx_tc_retard = 0
	m.xxx_cel_retard = 0
	m.xxx_fc_retard = 0
	m.xxx_als_addfuel = 0
}

func (m *megasquirt_gp56) CopyFrom(o megasquirt_gp56Reader) *megasquirt_gp56 {
	m.Settc_retard(o.tc_retard())
	m.Setcel_retard(o.cel_retard())
	m.Setfc_retard(o.fc_retard())
	m.Setals_addfuel(o.als_addfuel())
	return m
}

// Descriptor returns the megasquirt_gp56 descriptor.
func (m *megasquirt_gp56) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp56.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp56) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp56) tc_retard() float64 {
	return Messages().megasquirt_gp56.tc_retard.ToPhysical(float64(m.xxx_tc_retard))
}

func (m *megasquirt_gp56) Settc_retard(v float64) *megasquirt_gp56 {
	m.xxx_tc_retard = int16(Messages().megasquirt_gp56.tc_retard.FromPhysical(v))
	return m
}

func (m *megasquirt_gp56) cel_retard() float64 {
	return Messages().megasquirt_gp56.cel_retard.ToPhysical(float64(m.xxx_cel_retard))
}

func (m *megasquirt_gp56) Setcel_retard(v float64) *megasquirt_gp56 {
	m.xxx_cel_retard = int16(Messages().megasquirt_gp56.cel_retard.FromPhysical(v))
	return m
}

func (m *megasquirt_gp56) fc_retard() float64 {
	return Messages().megasquirt_gp56.fc_retard.ToPhysical(float64(m.xxx_fc_retard))
}

func (m *megasquirt_gp56) Setfc_retard(v float64) *megasquirt_gp56 {
	m.xxx_fc_retard = int16(Messages().megasquirt_gp56.fc_retard.FromPhysical(v))
	return m
}

func (m *megasquirt_gp56) als_addfuel() float64 {
	return Messages().megasquirt_gp56.als_addfuel.ToPhysical(float64(m.xxx_als_addfuel))
}

func (m *megasquirt_gp56) Setals_addfuel(v float64) *megasquirt_gp56 {
	m.xxx_als_addfuel = int16(Messages().megasquirt_gp56.als_addfuel.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp56) Frame() can.Frame {
	md := Messages().megasquirt_gp56
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.tc_retard.MarshalSigned(&f.Data, int64(m.xxx_tc_retard))
	md.cel_retard.MarshalSigned(&f.Data, int64(m.xxx_cel_retard))
	md.fc_retard.MarshalSigned(&f.Data, int64(m.xxx_fc_retard))
	md.als_addfuel.MarshalSigned(&f.Data, int64(m.xxx_als_addfuel))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp56) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp56) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp56
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp56: expects ID 1576 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp56: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp56: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp56: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_tc_retard = int16(md.tc_retard.UnmarshalSigned(f.Data))
	m.xxx_cel_retard = int16(md.cel_retard.UnmarshalSigned(f.Data))
	m.xxx_fc_retard = int16(md.fc_retard.UnmarshalSigned(f.Data))
	m.xxx_als_addfuel = int16(md.als_addfuel.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp57Reader provides read access to a megasquirt_gp57 message.
type megasquirt_gp57Reader interface {
	// base_advance returns the physical value of the base_advance signal.
	base_advance() float64
	// idle_cor_advance returns the physical value of the idle_cor_advance signal.
	idle_cor_advance() float64
	// mat_retard returns the physical value of the mat_retard signal.
	mat_retard() float64
	// flex_advance returns the physical value of the flex_advance signal.
	flex_advance() float64
}

// megasquirt_gp57Writer provides write access to a megasquirt_gp57 message.
type megasquirt_gp57Writer interface {
	// CopyFrom copies all values from megasquirt_gp57Reader.
	CopyFrom(megasquirt_gp57Reader) *megasquirt_gp57
	// Setbase_advance sets the physical value of the base_advance signal.
	Setbase_advance(float64) *megasquirt_gp57
	// Setidle_cor_advance sets the physical value of the idle_cor_advance signal.
	Setidle_cor_advance(float64) *megasquirt_gp57
	// Setmat_retard sets the physical value of the mat_retard signal.
	Setmat_retard(float64) *megasquirt_gp57
	// Setflex_advance sets the physical value of the flex_advance signal.
	Setflex_advance(float64) *megasquirt_gp57
}

type megasquirt_gp57 struct {
	xxx_base_advance     int16
	xxx_idle_cor_advance int16
	xxx_mat_retard       int16
	xxx_flex_advance     int16
}

func Newmegasquirt_gp57() *megasquirt_gp57 {
	m := &megasquirt_gp57{}
	m.Reset()
	return m
}

func (m *megasquirt_gp57) Reset() {
	m.xxx_base_advance = 0
	m.xxx_idle_cor_advance = 0
	m.xxx_mat_retard = 0
	m.xxx_flex_advance = 0
}

func (m *megasquirt_gp57) CopyFrom(o megasquirt_gp57Reader) *megasquirt_gp57 {
	m.Setbase_advance(o.base_advance())
	m.Setidle_cor_advance(o.idle_cor_advance())
	m.Setmat_retard(o.mat_retard())
	m.Setflex_advance(o.flex_advance())
	return m
}

// Descriptor returns the megasquirt_gp57 descriptor.
func (m *megasquirt_gp57) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp57.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp57) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp57) base_advance() float64 {
	return Messages().megasquirt_gp57.base_advance.ToPhysical(float64(m.xxx_base_advance))
}

func (m *megasquirt_gp57) Setbase_advance(v float64) *megasquirt_gp57 {
	m.xxx_base_advance = int16(Messages().megasquirt_gp57.base_advance.FromPhysical(v))
	return m
}

func (m *megasquirt_gp57) idle_cor_advance() float64 {
	return Messages().megasquirt_gp57.idle_cor_advance.ToPhysical(float64(m.xxx_idle_cor_advance))
}

func (m *megasquirt_gp57) Setidle_cor_advance(v float64) *megasquirt_gp57 {
	m.xxx_idle_cor_advance = int16(Messages().megasquirt_gp57.idle_cor_advance.FromPhysical(v))
	return m
}

func (m *megasquirt_gp57) mat_retard() float64 {
	return Messages().megasquirt_gp57.mat_retard.ToPhysical(float64(m.xxx_mat_retard))
}

func (m *megasquirt_gp57) Setmat_retard(v float64) *megasquirt_gp57 {
	m.xxx_mat_retard = int16(Messages().megasquirt_gp57.mat_retard.FromPhysical(v))
	return m
}

func (m *megasquirt_gp57) flex_advance() float64 {
	return Messages().megasquirt_gp57.flex_advance.ToPhysical(float64(m.xxx_flex_advance))
}

func (m *megasquirt_gp57) Setflex_advance(v float64) *megasquirt_gp57 {
	m.xxx_flex_advance = int16(Messages().megasquirt_gp57.flex_advance.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp57) Frame() can.Frame {
	md := Messages().megasquirt_gp57
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.base_advance.MarshalSigned(&f.Data, int64(m.xxx_base_advance))
	md.idle_cor_advance.MarshalSigned(&f.Data, int64(m.xxx_idle_cor_advance))
	md.mat_retard.MarshalSigned(&f.Data, int64(m.xxx_mat_retard))
	md.flex_advance.MarshalSigned(&f.Data, int64(m.xxx_flex_advance))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp57) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp57) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp57
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp57: expects ID 1577 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp57: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp57: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp57: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_base_advance = int16(md.base_advance.UnmarshalSigned(f.Data))
	m.xxx_idle_cor_advance = int16(md.idle_cor_advance.UnmarshalSigned(f.Data))
	m.xxx_mat_retard = int16(md.mat_retard.UnmarshalSigned(f.Data))
	m.xxx_flex_advance = int16(md.flex_advance.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp58Reader provides read access to a megasquirt_gp58 message.
type megasquirt_gp58Reader interface {
	// adv1 returns the physical value of the adv1 signal.
	adv1() float64
	// adv2 returns the physical value of the adv2 signal.
	adv2() float64
	// adv3 returns the physical value of the adv3 signal.
	adv3() float64
	// adv4 returns the physical value of the adv4 signal.
	adv4() float64
}

// megasquirt_gp58Writer provides write access to a megasquirt_gp58 message.
type megasquirt_gp58Writer interface {
	// CopyFrom copies all values from megasquirt_gp58Reader.
	CopyFrom(megasquirt_gp58Reader) *megasquirt_gp58
	// Setadv1 sets the physical value of the adv1 signal.
	Setadv1(float64) *megasquirt_gp58
	// Setadv2 sets the physical value of the adv2 signal.
	Setadv2(float64) *megasquirt_gp58
	// Setadv3 sets the physical value of the adv3 signal.
	Setadv3(float64) *megasquirt_gp58
	// Setadv4 sets the physical value of the adv4 signal.
	Setadv4(float64) *megasquirt_gp58
}

type megasquirt_gp58 struct {
	xxx_adv1 int16
	xxx_adv2 int16
	xxx_adv3 int16
	xxx_adv4 int16
}

func Newmegasquirt_gp58() *megasquirt_gp58 {
	m := &megasquirt_gp58{}
	m.Reset()
	return m
}

func (m *megasquirt_gp58) Reset() {
	m.xxx_adv1 = 0
	m.xxx_adv2 = 0
	m.xxx_adv3 = 0
	m.xxx_adv4 = 0
}

func (m *megasquirt_gp58) CopyFrom(o megasquirt_gp58Reader) *megasquirt_gp58 {
	m.Setadv1(o.adv1())
	m.Setadv2(o.adv2())
	m.Setadv3(o.adv3())
	m.Setadv4(o.adv4())
	return m
}

// Descriptor returns the megasquirt_gp58 descriptor.
func (m *megasquirt_gp58) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp58.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp58) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp58) adv1() float64 {
	return Messages().megasquirt_gp58.adv1.ToPhysical(float64(m.xxx_adv1))
}

func (m *megasquirt_gp58) Setadv1(v float64) *megasquirt_gp58 {
	m.xxx_adv1 = int16(Messages().megasquirt_gp58.adv1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp58) adv2() float64 {
	return Messages().megasquirt_gp58.adv2.ToPhysical(float64(m.xxx_adv2))
}

func (m *megasquirt_gp58) Setadv2(v float64) *megasquirt_gp58 {
	m.xxx_adv2 = int16(Messages().megasquirt_gp58.adv2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp58) adv3() float64 {
	return Messages().megasquirt_gp58.adv3.ToPhysical(float64(m.xxx_adv3))
}

func (m *megasquirt_gp58) Setadv3(v float64) *megasquirt_gp58 {
	m.xxx_adv3 = int16(Messages().megasquirt_gp58.adv3.FromPhysical(v))
	return m
}

func (m *megasquirt_gp58) adv4() float64 {
	return Messages().megasquirt_gp58.adv4.ToPhysical(float64(m.xxx_adv4))
}

func (m *megasquirt_gp58) Setadv4(v float64) *megasquirt_gp58 {
	m.xxx_adv4 = int16(Messages().megasquirt_gp58.adv4.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp58) Frame() can.Frame {
	md := Messages().megasquirt_gp58
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.adv1.MarshalSigned(&f.Data, int64(m.xxx_adv1))
	md.adv2.MarshalSigned(&f.Data, int64(m.xxx_adv2))
	md.adv3.MarshalSigned(&f.Data, int64(m.xxx_adv3))
	md.adv4.MarshalSigned(&f.Data, int64(m.xxx_adv4))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp58) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp58) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp58
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp58: expects ID 1578 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp58: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp58: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp58: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_adv1 = int16(md.adv1.UnmarshalSigned(f.Data))
	m.xxx_adv2 = int16(md.adv2.UnmarshalSigned(f.Data))
	m.xxx_adv3 = int16(md.adv3.UnmarshalSigned(f.Data))
	m.xxx_adv4 = int16(md.adv4.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp59Reader provides read access to a megasquirt_gp59 message.
type megasquirt_gp59Reader interface {
	// revlim_retard returns the physical value of the revlim_retard signal.
	revlim_retard() float64
	// als_timing returns the physical value of the als_timing signal.
	als_timing() float64
	// ext_advance returns the physical value of the ext_advance signal.
	ext_advance() float64
	// deadtime1 returns the physical value of the deadtime1 signal.
	deadtime1() float64
}

// megasquirt_gp59Writer provides write access to a megasquirt_gp59 message.
type megasquirt_gp59Writer interface {
	// CopyFrom copies all values from megasquirt_gp59Reader.
	CopyFrom(megasquirt_gp59Reader) *megasquirt_gp59
	// Setrevlim_retard sets the physical value of the revlim_retard signal.
	Setrevlim_retard(float64) *megasquirt_gp59
	// Setals_timing sets the physical value of the als_timing signal.
	Setals_timing(float64) *megasquirt_gp59
	// Setext_advance sets the physical value of the ext_advance signal.
	Setext_advance(float64) *megasquirt_gp59
	// Setdeadtime1 sets the physical value of the deadtime1 signal.
	Setdeadtime1(float64) *megasquirt_gp59
}

type megasquirt_gp59 struct {
	xxx_revlim_retard int16
	xxx_als_timing    int16
	xxx_ext_advance   int16
	xxx_deadtime1     int16
}

func Newmegasquirt_gp59() *megasquirt_gp59 {
	m := &megasquirt_gp59{}
	m.Reset()
	return m
}

func (m *megasquirt_gp59) Reset() {
	m.xxx_revlim_retard = 0
	m.xxx_als_timing = 0
	m.xxx_ext_advance = 0
	m.xxx_deadtime1 = 0
}

func (m *megasquirt_gp59) CopyFrom(o megasquirt_gp59Reader) *megasquirt_gp59 {
	m.Setrevlim_retard(o.revlim_retard())
	m.Setals_timing(o.als_timing())
	m.Setext_advance(o.ext_advance())
	m.Setdeadtime1(o.deadtime1())
	return m
}

// Descriptor returns the megasquirt_gp59 descriptor.
func (m *megasquirt_gp59) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp59.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp59) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp59) revlim_retard() float64 {
	return Messages().megasquirt_gp59.revlim_retard.ToPhysical(float64(m.xxx_revlim_retard))
}

func (m *megasquirt_gp59) Setrevlim_retard(v float64) *megasquirt_gp59 {
	m.xxx_revlim_retard = int16(Messages().megasquirt_gp59.revlim_retard.FromPhysical(v))
	return m
}

func (m *megasquirt_gp59) als_timing() float64 {
	return Messages().megasquirt_gp59.als_timing.ToPhysical(float64(m.xxx_als_timing))
}

func (m *megasquirt_gp59) Setals_timing(v float64) *megasquirt_gp59 {
	m.xxx_als_timing = int16(Messages().megasquirt_gp59.als_timing.FromPhysical(v))
	return m
}

func (m *megasquirt_gp59) ext_advance() float64 {
	return Messages().megasquirt_gp59.ext_advance.ToPhysical(float64(m.xxx_ext_advance))
}

func (m *megasquirt_gp59) Setext_advance(v float64) *megasquirt_gp59 {
	m.xxx_ext_advance = int16(Messages().megasquirt_gp59.ext_advance.FromPhysical(v))
	return m
}

func (m *megasquirt_gp59) deadtime1() float64 {
	return Messages().megasquirt_gp59.deadtime1.ToPhysical(float64(m.xxx_deadtime1))
}

func (m *megasquirt_gp59) Setdeadtime1(v float64) *megasquirt_gp59 {
	m.xxx_deadtime1 = int16(Messages().megasquirt_gp59.deadtime1.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp59) Frame() can.Frame {
	md := Messages().megasquirt_gp59
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.revlim_retard.MarshalSigned(&f.Data, int64(m.xxx_revlim_retard))
	md.als_timing.MarshalSigned(&f.Data, int64(m.xxx_als_timing))
	md.ext_advance.MarshalSigned(&f.Data, int64(m.xxx_ext_advance))
	md.deadtime1.MarshalSigned(&f.Data, int64(m.xxx_deadtime1))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp59) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp59) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp59
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp59: expects ID 1579 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp59: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp59: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp59: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_revlim_retard = int16(md.revlim_retard.UnmarshalSigned(f.Data))
	m.xxx_als_timing = int16(md.als_timing.UnmarshalSigned(f.Data))
	m.xxx_ext_advance = int16(md.ext_advance.UnmarshalSigned(f.Data))
	m.xxx_deadtime1 = int16(md.deadtime1.UnmarshalSigned(f.Data))
	return nil
}

// megasquirt_gp60Reader provides read access to a megasquirt_gp60 message.
type megasquirt_gp60Reader interface {
	// launch_timing returns the physical value of the launch_timing signal.
	launch_timing() float64
	// step3_timing returns the physical value of the step3_timing signal.
	step3_timing() float64
	// vsslaunch_retard returns the physical value of the vsslaunch_retard signal.
	vsslaunch_retard() float64
	// cel_status2 returns the value of the cel_status2 signal.
	cel_status2() uint16
}

// megasquirt_gp60Writer provides write access to a megasquirt_gp60 message.
type megasquirt_gp60Writer interface {
	// CopyFrom copies all values from megasquirt_gp60Reader.
	CopyFrom(megasquirt_gp60Reader) *megasquirt_gp60
	// Setlaunch_timing sets the physical value of the launch_timing signal.
	Setlaunch_timing(float64) *megasquirt_gp60
	// Setstep3_timing sets the physical value of the step3_timing signal.
	Setstep3_timing(float64) *megasquirt_gp60
	// Setvsslaunch_retard sets the physical value of the vsslaunch_retard signal.
	Setvsslaunch_retard(float64) *megasquirt_gp60
	// Setcel_status2 sets the value of the cel_status2 signal.
	Setcel_status2(uint16) *megasquirt_gp60
}

type megasquirt_gp60 struct {
	xxx_launch_timing    int16
	xxx_step3_timing     int16
	xxx_vsslaunch_retard int16
	xxx_cel_status2      uint16
}

func Newmegasquirt_gp60() *megasquirt_gp60 {
	m := &megasquirt_gp60{}
	m.Reset()
	return m
}

func (m *megasquirt_gp60) Reset() {
	m.xxx_launch_timing = 0
	m.xxx_step3_timing = 0
	m.xxx_vsslaunch_retard = 0
	m.xxx_cel_status2 = 0
}

func (m *megasquirt_gp60) CopyFrom(o megasquirt_gp60Reader) *megasquirt_gp60 {
	m.Setlaunch_timing(o.launch_timing())
	m.Setstep3_timing(o.step3_timing())
	m.Setvsslaunch_retard(o.vsslaunch_retard())
	m.xxx_cel_status2 = o.cel_status2()
	return m
}

// Descriptor returns the megasquirt_gp60 descriptor.
func (m *megasquirt_gp60) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp60.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp60) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp60) launch_timing() float64 {
	return Messages().megasquirt_gp60.launch_timing.ToPhysical(float64(m.xxx_launch_timing))
}

func (m *megasquirt_gp60) Setlaunch_timing(v float64) *megasquirt_gp60 {
	m.xxx_launch_timing = int16(Messages().megasquirt_gp60.launch_timing.FromPhysical(v))
	return m
}

func (m *megasquirt_gp60) step3_timing() float64 {
	return Messages().megasquirt_gp60.step3_timing.ToPhysical(float64(m.xxx_step3_timing))
}

func (m *megasquirt_gp60) Setstep3_timing(v float64) *megasquirt_gp60 {
	m.xxx_step3_timing = int16(Messages().megasquirt_gp60.step3_timing.FromPhysical(v))
	return m
}

func (m *megasquirt_gp60) vsslaunch_retard() float64 {
	return Messages().megasquirt_gp60.vsslaunch_retard.ToPhysical(float64(m.xxx_vsslaunch_retard))
}

func (m *megasquirt_gp60) Setvsslaunch_retard(v float64) *megasquirt_gp60 {
	m.xxx_vsslaunch_retard = int16(Messages().megasquirt_gp60.vsslaunch_retard.FromPhysical(v))
	return m
}

func (m *megasquirt_gp60) cel_status2() uint16 {
	return m.xxx_cel_status2
}

func (m *megasquirt_gp60) Setcel_status2(v uint16) *megasquirt_gp60 {
	m.xxx_cel_status2 = uint16(Messages().megasquirt_gp60.cel_status2.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp60) Frame() can.Frame {
	md := Messages().megasquirt_gp60
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.launch_timing.MarshalSigned(&f.Data, int64(m.xxx_launch_timing))
	md.step3_timing.MarshalSigned(&f.Data, int64(m.xxx_step3_timing))
	md.vsslaunch_retard.MarshalSigned(&f.Data, int64(m.xxx_vsslaunch_retard))
	md.cel_status2.MarshalUnsigned(&f.Data, uint64(m.xxx_cel_status2))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp60) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp60) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp60
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp60: expects ID 1580 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp60: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp60: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp60: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_launch_timing = int16(md.launch_timing.UnmarshalSigned(f.Data))
	m.xxx_step3_timing = int16(md.step3_timing.UnmarshalSigned(f.Data))
	m.xxx_vsslaunch_retard = int16(md.vsslaunch_retard.UnmarshalSigned(f.Data))
	m.xxx_cel_status2 = uint16(md.cel_status2.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp61Reader provides read access to a megasquirt_gp61 message.
type megasquirt_gp61Reader interface {
	// gps_latdeg returns the value of the gps_latdeg signal.
	gps_latdeg() int8
	// gps_latmin returns the value of the gps_latmin signal.
	gps_latmin() uint8
	// gps_latmmin returns the value of the gps_latmmin signal.
	gps_latmmin() uint16
	// gps_londeg returns the value of the gps_londeg signal.
	gps_londeg() uint8
	// gps_lonmin returns the value of the gps_lonmin signal.
	gps_lonmin() uint8
	// gps_lonmmin returns the value of the gps_lonmmin signal.
	gps_lonmmin() uint16
}

// megasquirt_gp61Writer provides write access to a megasquirt_gp61 message.
type megasquirt_gp61Writer interface {
	// CopyFrom copies all values from megasquirt_gp61Reader.
	CopyFrom(megasquirt_gp61Reader) *megasquirt_gp61
	// Setgps_latdeg sets the value of the gps_latdeg signal.
	Setgps_latdeg(int8) *megasquirt_gp61
	// Setgps_latmin sets the value of the gps_latmin signal.
	Setgps_latmin(uint8) *megasquirt_gp61
	// Setgps_latmmin sets the value of the gps_latmmin signal.
	Setgps_latmmin(uint16) *megasquirt_gp61
	// Setgps_londeg sets the value of the gps_londeg signal.
	Setgps_londeg(uint8) *megasquirt_gp61
	// Setgps_lonmin sets the value of the gps_lonmin signal.
	Setgps_lonmin(uint8) *megasquirt_gp61
	// Setgps_lonmmin sets the value of the gps_lonmmin signal.
	Setgps_lonmmin(uint16) *megasquirt_gp61
}

type megasquirt_gp61 struct {
	xxx_gps_latdeg  int8
	xxx_gps_latmin  uint8
	xxx_gps_latmmin uint16
	xxx_gps_londeg  uint8
	xxx_gps_lonmin  uint8
	xxx_gps_lonmmin uint16
}

func Newmegasquirt_gp61() *megasquirt_gp61 {
	m := &megasquirt_gp61{}
	m.Reset()
	return m
}

func (m *megasquirt_gp61) Reset() {
	m.xxx_gps_latdeg = 0
	m.xxx_gps_latmin = 0
	m.xxx_gps_latmmin = 0
	m.xxx_gps_londeg = 0
	m.xxx_gps_lonmin = 0
	m.xxx_gps_lonmmin = 0
}

func (m *megasquirt_gp61) CopyFrom(o megasquirt_gp61Reader) *megasquirt_gp61 {
	m.xxx_gps_latdeg = o.gps_latdeg()
	m.xxx_gps_latmin = o.gps_latmin()
	m.xxx_gps_latmmin = o.gps_latmmin()
	m.xxx_gps_londeg = o.gps_londeg()
	m.xxx_gps_lonmin = o.gps_lonmin()
	m.xxx_gps_lonmmin = o.gps_lonmmin()
	return m
}

// Descriptor returns the megasquirt_gp61 descriptor.
func (m *megasquirt_gp61) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp61.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp61) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp61) gps_latdeg() int8 {
	return m.xxx_gps_latdeg
}

func (m *megasquirt_gp61) Setgps_latdeg(v int8) *megasquirt_gp61 {
	m.xxx_gps_latdeg = int8(Messages().megasquirt_gp61.gps_latdeg.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp61) gps_latmin() uint8 {
	return m.xxx_gps_latmin
}

func (m *megasquirt_gp61) Setgps_latmin(v uint8) *megasquirt_gp61 {
	m.xxx_gps_latmin = uint8(Messages().megasquirt_gp61.gps_latmin.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp61) gps_latmmin() uint16 {
	return m.xxx_gps_latmmin
}

func (m *megasquirt_gp61) Setgps_latmmin(v uint16) *megasquirt_gp61 {
	m.xxx_gps_latmmin = uint16(Messages().megasquirt_gp61.gps_latmmin.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp61) gps_londeg() uint8 {
	return m.xxx_gps_londeg
}

func (m *megasquirt_gp61) Setgps_londeg(v uint8) *megasquirt_gp61 {
	m.xxx_gps_londeg = uint8(Messages().megasquirt_gp61.gps_londeg.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp61) gps_lonmin() uint8 {
	return m.xxx_gps_lonmin
}

func (m *megasquirt_gp61) Setgps_lonmin(v uint8) *megasquirt_gp61 {
	m.xxx_gps_lonmin = uint8(Messages().megasquirt_gp61.gps_lonmin.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp61) gps_lonmmin() uint16 {
	return m.xxx_gps_lonmmin
}

func (m *megasquirt_gp61) Setgps_lonmmin(v uint16) *megasquirt_gp61 {
	m.xxx_gps_lonmmin = uint16(Messages().megasquirt_gp61.gps_lonmmin.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp61) Frame() can.Frame {
	md := Messages().megasquirt_gp61
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.gps_latdeg.MarshalSigned(&f.Data, int64(m.xxx_gps_latdeg))
	md.gps_latmin.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_latmin))
	md.gps_latmmin.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_latmmin))
	md.gps_londeg.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_londeg))
	md.gps_lonmin.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_lonmin))
	md.gps_lonmmin.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_lonmmin))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp61) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp61) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp61
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp61: expects ID 1581 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp61: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp61: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp61: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_gps_latdeg = int8(md.gps_latdeg.UnmarshalSigned(f.Data))
	m.xxx_gps_latmin = uint8(md.gps_latmin.UnmarshalUnsigned(f.Data))
	m.xxx_gps_latmmin = uint16(md.gps_latmmin.UnmarshalUnsigned(f.Data))
	m.xxx_gps_londeg = uint8(md.gps_londeg.UnmarshalUnsigned(f.Data))
	m.xxx_gps_lonmin = uint8(md.gps_lonmin.UnmarshalUnsigned(f.Data))
	m.xxx_gps_lonmmin = uint16(md.gps_lonmmin.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp62Reader provides read access to a megasquirt_gp62 message.
type megasquirt_gp62Reader interface {
	// gps_west returns the value of the gps_west signal.
	gps_west() bool
	// gps_altk returns the value of the gps_altk signal.
	gps_altk() int8
	// gps_altm returns the value of the gps_altm signal.
	gps_altm() uint16
	// gps_speed returns the physical value of the gps_speed signal.
	gps_speed() float64
	// gps_course returns the physical value of the gps_course signal.
	gps_course() float64
}

// megasquirt_gp62Writer provides write access to a megasquirt_gp62 message.
type megasquirt_gp62Writer interface {
	// CopyFrom copies all values from megasquirt_gp62Reader.
	CopyFrom(megasquirt_gp62Reader) *megasquirt_gp62
	// Setgps_west sets the value of the gps_west signal.
	Setgps_west(bool) *megasquirt_gp62
	// Setgps_altk sets the value of the gps_altk signal.
	Setgps_altk(int8) *megasquirt_gp62
	// Setgps_altm sets the value of the gps_altm signal.
	Setgps_altm(uint16) *megasquirt_gp62
	// Setgps_speed sets the physical value of the gps_speed signal.
	Setgps_speed(float64) *megasquirt_gp62
	// Setgps_course sets the physical value of the gps_course signal.
	Setgps_course(float64) *megasquirt_gp62
}

type megasquirt_gp62 struct {
	xxx_gps_west   bool
	xxx_gps_altk   int8
	xxx_gps_altm   uint16
	xxx_gps_speed  uint16
	xxx_gps_course uint16
}

func Newmegasquirt_gp62() *megasquirt_gp62 {
	m := &megasquirt_gp62{}
	m.Reset()
	return m
}

func (m *megasquirt_gp62) Reset() {
	m.xxx_gps_west = false
	m.xxx_gps_altk = 0
	m.xxx_gps_altm = 0
	m.xxx_gps_speed = 0
	m.xxx_gps_course = 0
}

func (m *megasquirt_gp62) CopyFrom(o megasquirt_gp62Reader) *megasquirt_gp62 {
	m.xxx_gps_west = o.gps_west()
	m.xxx_gps_altk = o.gps_altk()
	m.xxx_gps_altm = o.gps_altm()
	m.Setgps_speed(o.gps_speed())
	m.Setgps_course(o.gps_course())
	return m
}

// Descriptor returns the megasquirt_gp62 descriptor.
func (m *megasquirt_gp62) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp62.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp62) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp62) gps_west() bool {
	return m.xxx_gps_west
}

func (m *megasquirt_gp62) Setgps_west(v bool) *megasquirt_gp62 {
	m.xxx_gps_west = v
	return m
}

func (m *megasquirt_gp62) gps_altk() int8 {
	return m.xxx_gps_altk
}

func (m *megasquirt_gp62) Setgps_altk(v int8) *megasquirt_gp62 {
	m.xxx_gps_altk = int8(Messages().megasquirt_gp62.gps_altk.SaturatedCastSigned(int64(v)))
	return m
}

func (m *megasquirt_gp62) gps_altm() uint16 {
	return m.xxx_gps_altm
}

func (m *megasquirt_gp62) Setgps_altm(v uint16) *megasquirt_gp62 {
	m.xxx_gps_altm = uint16(Messages().megasquirt_gp62.gps_altm.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp62) gps_speed() float64 {
	return Messages().megasquirt_gp62.gps_speed.ToPhysical(float64(m.xxx_gps_speed))
}

func (m *megasquirt_gp62) Setgps_speed(v float64) *megasquirt_gp62 {
	m.xxx_gps_speed = uint16(Messages().megasquirt_gp62.gps_speed.FromPhysical(v))
	return m
}

func (m *megasquirt_gp62) gps_course() float64 {
	return Messages().megasquirt_gp62.gps_course.ToPhysical(float64(m.xxx_gps_course))
}

func (m *megasquirt_gp62) Setgps_course(v float64) *megasquirt_gp62 {
	m.xxx_gps_course = uint16(Messages().megasquirt_gp62.gps_course.FromPhysical(v))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp62) Frame() can.Frame {
	md := Messages().megasquirt_gp62
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.gps_west.MarshalBool(&f.Data, bool(m.xxx_gps_west))
	md.gps_altk.MarshalSigned(&f.Data, int64(m.xxx_gps_altk))
	md.gps_altm.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_altm))
	md.gps_speed.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_speed))
	md.gps_course.MarshalUnsigned(&f.Data, uint64(m.xxx_gps_course))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp62) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp62) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp62
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp62: expects ID 1582 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp62: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp62: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp62: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_gps_west = bool(md.gps_west.UnmarshalBool(f.Data))
	m.xxx_gps_altk = int8(md.gps_altk.UnmarshalSigned(f.Data))
	m.xxx_gps_altm = uint16(md.gps_altm.UnmarshalUnsigned(f.Data))
	m.xxx_gps_speed = uint16(md.gps_speed.UnmarshalUnsigned(f.Data))
	m.xxx_gps_course = uint16(md.gps_course.UnmarshalUnsigned(f.Data))
	return nil
}

// megasquirt_gp63Reader provides read access to a megasquirt_gp63 message.
type megasquirt_gp63Reader interface {
	// generic_pid_duty1 returns the physical value of the generic_pid_duty1 signal.
	generic_pid_duty1() float64
	// generic_pid_duty2 returns the physical value of the generic_pid_duty2 signal.
	generic_pid_duty2() float64
	// spare63_1 returns the value of the spare63_1 signal.
	spare63_1() uint16
	// spare63_2 returns the value of the spare63_2 signal.
	spare63_2() uint16
	// spare63_3 returns the value of the spare63_3 signal.
	spare63_3() uint16
}

// megasquirt_gp63Writer provides write access to a megasquirt_gp63 message.
type megasquirt_gp63Writer interface {
	// CopyFrom copies all values from megasquirt_gp63Reader.
	CopyFrom(megasquirt_gp63Reader) *megasquirt_gp63
	// Setgeneric_pid_duty1 sets the physical value of the generic_pid_duty1 signal.
	Setgeneric_pid_duty1(float64) *megasquirt_gp63
	// Setgeneric_pid_duty2 sets the physical value of the generic_pid_duty2 signal.
	Setgeneric_pid_duty2(float64) *megasquirt_gp63
	// Setspare63_1 sets the value of the spare63_1 signal.
	Setspare63_1(uint16) *megasquirt_gp63
	// Setspare63_2 sets the value of the spare63_2 signal.
	Setspare63_2(uint16) *megasquirt_gp63
	// Setspare63_3 sets the value of the spare63_3 signal.
	Setspare63_3(uint16) *megasquirt_gp63
}

type megasquirt_gp63 struct {
	xxx_generic_pid_duty1 uint8
	xxx_generic_pid_duty2 uint8
	xxx_spare63_1         uint16
	xxx_spare63_2         uint16
	xxx_spare63_3         uint16
}

func Newmegasquirt_gp63() *megasquirt_gp63 {
	m := &megasquirt_gp63{}
	m.Reset()
	return m
}

func (m *megasquirt_gp63) Reset() {
	m.xxx_generic_pid_duty1 = 0
	m.xxx_generic_pid_duty2 = 0
	m.xxx_spare63_1 = 0
	m.xxx_spare63_2 = 0
	m.xxx_spare63_3 = 0
}

func (m *megasquirt_gp63) CopyFrom(o megasquirt_gp63Reader) *megasquirt_gp63 {
	m.Setgeneric_pid_duty1(o.generic_pid_duty1())
	m.Setgeneric_pid_duty2(o.generic_pid_duty2())
	m.xxx_spare63_1 = o.spare63_1()
	m.xxx_spare63_2 = o.spare63_2()
	m.xxx_spare63_3 = o.spare63_3()
	return m
}

// Descriptor returns the megasquirt_gp63 descriptor.
func (m *megasquirt_gp63) Descriptor() *descriptor.Message {
	return Messages().megasquirt_gp63.Message
}

// String returns a compact string representation of the message.
func (m *megasquirt_gp63) String() string {
	return cantext.MessageString(m)
}

func (m *megasquirt_gp63) generic_pid_duty1() float64 {
	return Messages().megasquirt_gp63.generic_pid_duty1.ToPhysical(float64(m.xxx_generic_pid_duty1))
}

func (m *megasquirt_gp63) Setgeneric_pid_duty1(v float64) *megasquirt_gp63 {
	m.xxx_generic_pid_duty1 = uint8(Messages().megasquirt_gp63.generic_pid_duty1.FromPhysical(v))
	return m
}

func (m *megasquirt_gp63) generic_pid_duty2() float64 {
	return Messages().megasquirt_gp63.generic_pid_duty2.ToPhysical(float64(m.xxx_generic_pid_duty2))
}

func (m *megasquirt_gp63) Setgeneric_pid_duty2(v float64) *megasquirt_gp63 {
	m.xxx_generic_pid_duty2 = uint8(Messages().megasquirt_gp63.generic_pid_duty2.FromPhysical(v))
	return m
}

func (m *megasquirt_gp63) spare63_1() uint16 {
	return m.xxx_spare63_1
}

func (m *megasquirt_gp63) Setspare63_1(v uint16) *megasquirt_gp63 {
	m.xxx_spare63_1 = uint16(Messages().megasquirt_gp63.spare63_1.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp63) spare63_2() uint16 {
	return m.xxx_spare63_2
}

func (m *megasquirt_gp63) Setspare63_2(v uint16) *megasquirt_gp63 {
	m.xxx_spare63_2 = uint16(Messages().megasquirt_gp63.spare63_2.SaturatedCastUnsigned(uint64(v)))
	return m
}

func (m *megasquirt_gp63) spare63_3() uint16 {
	return m.xxx_spare63_3
}

func (m *megasquirt_gp63) Setspare63_3(v uint16) *megasquirt_gp63 {
	m.xxx_spare63_3 = uint16(Messages().megasquirt_gp63.spare63_3.SaturatedCastUnsigned(uint64(v)))
	return m
}

// Frame returns a CAN frame representing the message.
func (m *megasquirt_gp63) Frame() can.Frame {
	md := Messages().megasquirt_gp63
	f := can.Frame{ID: md.ID, IsExtended: md.IsExtended, Length: md.Length}
	md.generic_pid_duty1.MarshalUnsigned(&f.Data, uint64(m.xxx_generic_pid_duty1))
	md.generic_pid_duty2.MarshalUnsigned(&f.Data, uint64(m.xxx_generic_pid_duty2))
	md.spare63_1.MarshalUnsigned(&f.Data, uint64(m.xxx_spare63_1))
	md.spare63_2.MarshalUnsigned(&f.Data, uint64(m.xxx_spare63_2))
	md.spare63_3.MarshalUnsigned(&f.Data, uint64(m.xxx_spare63_3))
	return f
}

// MarshalFrame encodes the message as a CAN frame.
func (m *megasquirt_gp63) MarshalFrame() (can.Frame, error) {
	return m.Frame(), nil
}

// UnmarshalFrame decodes the message from a CAN frame.
func (m *megasquirt_gp63) UnmarshalFrame(f can.Frame) error {
	md := Messages().megasquirt_gp63
	switch {
	case f.ID != md.ID:
		return fmt.Errorf(
			"unmarshal megasquirt_gp63: expects ID 1583 (got %s with ID %d)", f.String(), f.ID,
		)
	case f.Length != md.Length:
		return fmt.Errorf(
			"unmarshal megasquirt_gp63: expects length 8 (got %s with length %d)", f.String(), f.Length,
		)
	case f.IsRemote:
		return fmt.Errorf(
			"unmarshal megasquirt_gp63: expects non-remote frame (got remote frame %s)", f.String(),
		)
	case f.IsExtended != md.IsExtended:
		return fmt.Errorf(
			"unmarshal megasquirt_gp63: expects standard ID (got %s with extended ID)", f.String(),
		)
	}
	m.xxx_generic_pid_duty1 = uint8(md.generic_pid_duty1.UnmarshalUnsigned(f.Data))
	m.xxx_generic_pid_duty2 = uint8(md.generic_pid_duty2.UnmarshalUnsigned(f.Data))
	m.xxx_spare63_1 = uint16(md.spare63_1.UnmarshalUnsigned(f.Data))
	m.xxx_spare63_2 = uint16(md.spare63_2.UnmarshalUnsigned(f.Data))
	m.xxx_spare63_3 = uint16(md.spare63_3.UnmarshalUnsigned(f.Data))
	return nil
}

// Nodes returns the megasquirt node descriptors.
func Nodes() *NodesDescriptor {
	return nd
}

// NodesDescriptor contains all megasquirt node descriptors.
type NodesDescriptor struct {
}

// Messages returns the megasquirt message descriptors.
func Messages() *MessagesDescriptor {
	return md
}

// MessagesDescriptor contains all megasquirt message descriptors.
type MessagesDescriptor struct {
	megasquirt_gp0  *megasquirt_gp0Descriptor
	megasquirt_gp1  *megasquirt_gp1Descriptor
	megasquirt_gp2  *megasquirt_gp2Descriptor
	megasquirt_gp3  *megasquirt_gp3Descriptor
	megasquirt_gp4  *megasquirt_gp4Descriptor
	megasquirt_gp5  *megasquirt_gp5Descriptor
	megasquirt_gp6  *megasquirt_gp6Descriptor
	megasquirt_gp7  *megasquirt_gp7Descriptor
	megasquirt_gp8  *megasquirt_gp8Descriptor
	megasquirt_gp9  *megasquirt_gp9Descriptor
	megasquirt_gp10 *megasquirt_gp10Descriptor
	megasquirt_gp11 *megasquirt_gp11Descriptor
	megasquirt_gp12 *megasquirt_gp12Descriptor
	megasquirt_gp13 *megasquirt_gp13Descriptor
	megasquirt_gp14 *megasquirt_gp14Descriptor
	megasquirt_gp15 *megasquirt_gp15Descriptor
	megasquirt_gp16 *megasquirt_gp16Descriptor
	megasquirt_gp17 *megasquirt_gp17Descriptor
	megasquirt_gp18 *megasquirt_gp18Descriptor
	megasquirt_gp19 *megasquirt_gp19Descriptor
	megasquirt_gp20 *megasquirt_gp20Descriptor
	megasquirt_gp21 *megasquirt_gp21Descriptor
	megasquirt_gp22 *megasquirt_gp22Descriptor
	megasquirt_gp23 *megasquirt_gp23Descriptor
	megasquirt_gp24 *megasquirt_gp24Descriptor
	megasquirt_gp25 *megasquirt_gp25Descriptor
	megasquirt_gp26 *megasquirt_gp26Descriptor
	megasquirt_gp27 *megasquirt_gp27Descriptor
	megasquirt_gp28 *megasquirt_gp28Descriptor
	megasquirt_gp29 *megasquirt_gp29Descriptor
	megasquirt_gp30 *megasquirt_gp30Descriptor
	megasquirt_gp31 *megasquirt_gp31Descriptor
	megasquirt_gp32 *megasquirt_gp32Descriptor
	megasquirt_gp33 *megasquirt_gp33Descriptor
	megasquirt_gp34 *megasquirt_gp34Descriptor
	megasquirt_gp35 *megasquirt_gp35Descriptor
	megasquirt_gp36 *megasquirt_gp36Descriptor
	megasquirt_gp37 *megasquirt_gp37Descriptor
	megasquirt_gp38 *megasquirt_gp38Descriptor
	megasquirt_gp39 *megasquirt_gp39Descriptor
	megasquirt_gp40 *megasquirt_gp40Descriptor
	megasquirt_gp41 *megasquirt_gp41Descriptor
	megasquirt_gp42 *megasquirt_gp42Descriptor
	megasquirt_gp43 *megasquirt_gp43Descriptor
	megasquirt_gp44 *megasquirt_gp44Descriptor
	megasquirt_gp45 *megasquirt_gp45Descriptor
	megasquirt_gp46 *megasquirt_gp46Descriptor
	megasquirt_gp47 *megasquirt_gp47Descriptor
	megasquirt_gp48 *megasquirt_gp48Descriptor
	megasquirt_gp49 *megasquirt_gp49Descriptor
	megasquirt_gp50 *megasquirt_gp50Descriptor
	megasquirt_gp51 *megasquirt_gp51Descriptor
	megasquirt_gp52 *megasquirt_gp52Descriptor
	megasquirt_gp53 *megasquirt_gp53Descriptor
	megasquirt_gp54 *megasquirt_gp54Descriptor
	megasquirt_gp55 *megasquirt_gp55Descriptor
	megasquirt_gp56 *megasquirt_gp56Descriptor
	megasquirt_gp57 *megasquirt_gp57Descriptor
	megasquirt_gp58 *megasquirt_gp58Descriptor
	megasquirt_gp59 *megasquirt_gp59Descriptor
	megasquirt_gp60 *megasquirt_gp60Descriptor
	megasquirt_gp61 *megasquirt_gp61Descriptor
	megasquirt_gp62 *megasquirt_gp62Descriptor
	megasquirt_gp63 *megasquirt_gp63Descriptor
}

// UnmarshalFrame unmarshals the provided megasquirt CAN frame.
func (md *MessagesDescriptor) UnmarshalFrame(f can.Frame) (generated.Message, error) {
	switch f.ID {
	case md.megasquirt_gp0.ID:
		var msg megasquirt_gp0
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp1.ID:
		var msg megasquirt_gp1
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp2.ID:
		var msg megasquirt_gp2
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp3.ID:
		var msg megasquirt_gp3
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp4.ID:
		var msg megasquirt_gp4
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp5.ID:
		var msg megasquirt_gp5
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp6.ID:
		var msg megasquirt_gp6
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp7.ID:
		var msg megasquirt_gp7
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp8.ID:
		var msg megasquirt_gp8
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp9.ID:
		var msg megasquirt_gp9
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp10.ID:
		var msg megasquirt_gp10
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp11.ID:
		var msg megasquirt_gp11
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp12.ID:
		var msg megasquirt_gp12
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp13.ID:
		var msg megasquirt_gp13
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp14.ID:
		var msg megasquirt_gp14
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp15.ID:
		var msg megasquirt_gp15
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp16.ID:
		var msg megasquirt_gp16
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp17.ID:
		var msg megasquirt_gp17
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp18.ID:
		var msg megasquirt_gp18
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp19.ID:
		var msg megasquirt_gp19
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp20.ID:
		var msg megasquirt_gp20
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp21.ID:
		var msg megasquirt_gp21
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp22.ID:
		var msg megasquirt_gp22
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp23.ID:
		var msg megasquirt_gp23
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp24.ID:
		var msg megasquirt_gp24
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp25.ID:
		var msg megasquirt_gp25
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp26.ID:
		var msg megasquirt_gp26
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp27.ID:
		var msg megasquirt_gp27
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp28.ID:
		var msg megasquirt_gp28
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp29.ID:
		var msg megasquirt_gp29
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp30.ID:
		var msg megasquirt_gp30
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp31.ID:
		var msg megasquirt_gp31
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp32.ID:
		var msg megasquirt_gp32
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp33.ID:
		var msg megasquirt_gp33
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp34.ID:
		var msg megasquirt_gp34
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp35.ID:
		var msg megasquirt_gp35
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp36.ID:
		var msg megasquirt_gp36
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp37.ID:
		var msg megasquirt_gp37
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp38.ID:
		var msg megasquirt_gp38
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp39.ID:
		var msg megasquirt_gp39
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp40.ID:
		var msg megasquirt_gp40
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp41.ID:
		var msg megasquirt_gp41
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp42.ID:
		var msg megasquirt_gp42
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp43.ID:
		var msg megasquirt_gp43
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp44.ID:
		var msg megasquirt_gp44
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp45.ID:
		var msg megasquirt_gp45
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp46.ID:
		var msg megasquirt_gp46
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp47.ID:
		var msg megasquirt_gp47
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp48.ID:
		var msg megasquirt_gp48
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp49.ID:
		var msg megasquirt_gp49
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp50.ID:
		var msg megasquirt_gp50
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp51.ID:
		var msg megasquirt_gp51
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp52.ID:
		var msg megasquirt_gp52
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp53.ID:
		var msg megasquirt_gp53
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp54.ID:
		var msg megasquirt_gp54
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp55.ID:
		var msg megasquirt_gp55
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp56.ID:
		var msg megasquirt_gp56
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp57.ID:
		var msg megasquirt_gp57
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp58.ID:
		var msg megasquirt_gp58
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp59.ID:
		var msg megasquirt_gp59
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp60.ID:
		var msg megasquirt_gp60
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp61.ID:
		var msg megasquirt_gp61
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp62.ID:
		var msg megasquirt_gp62
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	case md.megasquirt_gp63.ID:
		var msg megasquirt_gp63
		if err := msg.UnmarshalFrame(f); err != nil {
			return nil, fmt.Errorf("unmarshal megasquirt frame: %w", err)
		}
		return &msg, nil
	default:
		return nil, fmt.Errorf("unmarshal megasquirt frame: ID not in database: %d", f.ID)
	}
}

type megasquirt_gp0Descriptor struct {
	*descriptor.Message
	seconds *descriptor.Signal
	pw1     *descriptor.Signal
	pw2     *descriptor.Signal
	rpm     *descriptor.Signal
}

type megasquirt_gp1Descriptor struct {
	*descriptor.Message
	adv_deg  *descriptor.Signal
	squirt   *descriptor.Signal
	engine   *descriptor.Signal
	afrtgt1  *descriptor.Signal
	afrtgt2  *descriptor.Signal
	wbo2_en1 *descriptor.Signal
	wbo2_en2 *descriptor.Signal
}

type megasquirt_gp2Descriptor struct {
	*descriptor.Message
	baro *descriptor.Signal
	map1 *descriptor.Signal
	mat  *descriptor.Signal
	clt  *descriptor.Signal
}

type megasquirt_gp3Descriptor struct {
	*descriptor.Message
	tps      *descriptor.Signal
	batt     *descriptor.Signal
	afr1_old *descriptor.Signal
	afr2_old *descriptor.Signal
}

type megasquirt_gp4Descriptor struct {
	*descriptor.Message
	knock   *descriptor.Signal
	egocor1 *descriptor.Signal
	egocor2 *descriptor.Signal
	aircor  *descriptor.Signal
}

type megasquirt_gp5Descriptor struct {
	*descriptor.Message
	warmcor    *descriptor.Signal
	tpsaccel   *descriptor.Signal
	tpsfuelcut *descriptor.Signal
	barocor    *descriptor.Signal
}

type megasquirt_gp6Descriptor struct {
	*descriptor.Message
	totalcor *descriptor.Signal
	ve1      *descriptor.Signal
	ve2      *descriptor.Signal
	iacstep  *descriptor.Signal
}

type megasquirt_gp7Descriptor struct {
	*descriptor.Message
	cold_adv_deg *descriptor.Signal
	TPSdot       *descriptor.Signal
	MAPdot       *descriptor.Signal
	RPMdot       *descriptor.Signal
}

type megasquirt_gp8Descriptor struct {
	*descriptor.Message
	MAFload  *descriptor.Signal
	fuelload *descriptor.Signal
	fuelcor  *descriptor.Signal
	MAF      *descriptor.Signal
}

type megasquirt_gp9Descriptor struct {
	*descriptor.Message
	egoV1     *descriptor.Signal
	egoV2     *descriptor.Signal
	dwell     *descriptor.Signal
	dwell_trl *descriptor.Signal
}

type megasquirt_gp10Descriptor struct {
	*descriptor.Message
	status1 *descriptor.Signal
	status2 *descriptor.Signal
	status3 *descriptor.Signal
	status4 *descriptor.Signal
	status5 *descriptor.Signal
	status6 *descriptor.Signal
	status7 *descriptor.Signal
}

type megasquirt_gp11Descriptor struct {
	*descriptor.Message
	fuelload2 *descriptor.Signal
	ignload   *descriptor.Signal
	ignload2  *descriptor.Signal
	airtemp   *descriptor.Signal
}

type megasquirt_gp12Descriptor struct {
	*descriptor.Message
	wallfuel1 *descriptor.Signal
	wallfuel2 *descriptor.Signal
}

type megasquirt_gp13Descriptor struct {
	*descriptor.Message
	sensors1 *descriptor.Signal
	sensors2 *descriptor.Signal
	sensors3 *descriptor.Signal
	sensors4 *descriptor.Signal
}

type megasquirt_gp14Descriptor struct {
	*descriptor.Message
	sensors5 *descriptor.Signal
	sensors6 *descriptor.Signal
	sensors7 *descriptor.Signal
	sensors8 *descriptor.Signal
}

type megasquirt_gp15Descriptor struct {
	*descriptor.Message
	sensors9  *descriptor.Signal
	sensors10 *descriptor.Signal
	sensors11 *descriptor.Signal
	sensors12 *descriptor.Signal
}

type megasquirt_gp16Descriptor struct {
	*descriptor.Message
	sensors13 *descriptor.Signal
	sensors14 *descriptor.Signal
	sensors15 *descriptor.Signal
	sensors16 *descriptor.Signal
}

type megasquirt_gp17Descriptor struct {
	*descriptor.Message
	boost_targ_1 *descriptor.Signal
	boost_targ_2 *descriptor.Signal
	boostduty    *descriptor.Signal
	boostduty2   *descriptor.Signal
	maf_volts    *descriptor.Signal
}

type megasquirt_gp18Descriptor struct {
	*descriptor.Message
	pwseq1 *descriptor.Signal
	pwseq2 *descriptor.Signal
	pwseq3 *descriptor.Signal
	pwseq4 *descriptor.Signal
}

type megasquirt_gp19Descriptor struct {
	*descriptor.Message
	pwseq5 *descriptor.Signal
	pwseq6 *descriptor.Signal
	pwseq7 *descriptor.Signal
	pwseq8 *descriptor.Signal
}

type megasquirt_gp20Descriptor struct {
	*descriptor.Message
	pwseq9  *descriptor.Signal
	pwseq10 *descriptor.Signal
	pwseq11 *descriptor.Signal
	pwseq12 *descriptor.Signal
}

type megasquirt_gp21Descriptor struct {
	*descriptor.Message
	pwseq13 *descriptor.Signal
	pwseq14 *descriptor.Signal
	pwseq15 *descriptor.Signal
	pwseq16 *descriptor.Signal
}

type megasquirt_gp22Descriptor struct {
	*descriptor.Message
	egt1 *descriptor.Signal
	egt2 *descriptor.Signal
	egt3 *descriptor.Signal
	egt4 *descriptor.Signal
}

type megasquirt_gp23Descriptor struct {
	*descriptor.Message
	egt5 *descriptor.Signal
	egt6 *descriptor.Signal
	egt7 *descriptor.Signal
	egt8 *descriptor.Signal
}

type megasquirt_gp24Descriptor struct {
	*descriptor.Message
	egt9  *descriptor.Signal
	egt10 *descriptor.Signal
	egt11 *descriptor.Signal
	egt12 *descriptor.Signal
}

type megasquirt_gp25Descriptor struct {
	*descriptor.Message
	egt13 *descriptor.Signal
	egt14 *descriptor.Signal
	egt15 *descriptor.Signal
	egt16 *descriptor.Signal
}

type megasquirt_gp26Descriptor struct {
	*descriptor.Message
	nitrous1_duty     *descriptor.Signal
	nitrous2_duty     *descriptor.Signal
	nitrous_timer_out *descriptor.Signal
	n2o_addfuel       *descriptor.Signal
	n2o_retard        *descriptor.Signal
}

type megasquirt_gp27Descriptor struct {
	*descriptor.Message
	canpwmin1 *descriptor.Signal
	canpwmin2 *descriptor.Signal
	canpwmin3 *descriptor.Signal
	canpwmin4 *descriptor.Signal
}

type megasquirt_gp28Descriptor struct {
	*descriptor.Message
	cl_idle_targ_rpm *descriptor.Signal
	tpsadc           *descriptor.Signal
	eaeload          *descriptor.Signal
	afrload          *descriptor.Signal
}

type megasquirt_gp29Descriptor struct {
	*descriptor.Message
	EAEfcor1 *descriptor.Signal
	EAEfcor2 *descriptor.Signal
	VSS1dot  *descriptor.Signal
	VSS2dot  *descriptor.Signal
}

type megasquirt_gp30Descriptor struct {
	*descriptor.Message
	accelx       *descriptor.Signal
	accely       *descriptor.Signal
	accelz       *descriptor.Signal
	stream_level *descriptor.Signal
	water_duty   *descriptor.Signal
}

type megasquirt_gp31Descriptor struct {
	*descriptor.Message
	AFR1 *descriptor.Signal
	AFR2 *descriptor.Signal
	AFR3 *descriptor.Signal
	AFR4 *descriptor.Signal
	AFR5 *descriptor.Signal
	AFR6 *descriptor.Signal
	AFR7 *descriptor.Signal
	AFR8 *descriptor.Signal
}

type megasquirt_gp32Descriptor struct {
	*descriptor.Message
	AFR9  *descriptor.Signal
	AFR10 *descriptor.Signal
	AFR11 *descriptor.Signal
	AFR12 *descriptor.Signal
	AFR13 *descriptor.Signal
	AFR14 *descriptor.Signal
	AFR15 *descriptor.Signal
	AFR16 *descriptor.Signal
}

type megasquirt_gp33Descriptor struct {
	*descriptor.Message
	duty_pwm1 *descriptor.Signal
	duty_pwm2 *descriptor.Signal
	duty_pwm3 *descriptor.Signal
	duty_pwm4 *descriptor.Signal
	duty_pwm5 *descriptor.Signal
	duty_pwm6 *descriptor.Signal
	gear      *descriptor.Signal
	status8   *descriptor.Signal
}

type megasquirt_gp34Descriptor struct {
	*descriptor.Message
	EGOv1 *descriptor.Signal
	EGOv2 *descriptor.Signal
	EGOv3 *descriptor.Signal
	EGOv4 *descriptor.Signal
}

type megasquirt_gp35Descriptor struct {
	*descriptor.Message
	EGOv5 *descriptor.Signal
	EGOv6 *descriptor.Signal
	EGOv7 *descriptor.Signal
	EGOv8 *descriptor.Signal
}

type megasquirt_gp36Descriptor struct {
	*descriptor.Message
	EGOv9  *descriptor.Signal
	EGOv10 *descriptor.Signal
	EGOv11 *descriptor.Signal
	EGOv12 *descriptor.Signal
}

type megasquirt_gp37Descriptor struct {
	*descriptor.Message
	EGOv13 *descriptor.Signal
	EGOv14 *descriptor.Signal
	EGOv15 *descriptor.Signal
	EGOv16 *descriptor.Signal
}

type megasquirt_gp38Descriptor struct {
	*descriptor.Message
	EGOcor1 *descriptor.Signal
	EGOcor2 *descriptor.Signal
	EGOcor3 *descriptor.Signal
	EGOcor4 *descriptor.Signal
}

type megasquirt_gp39Descriptor struct {
	*descriptor.Message
	EGOcor5 *descriptor.Signal
	EGOcor6 *descriptor.Signal
	EGOcor7 *descriptor.Signal
	EGOcor8 *descriptor.Signal
}

type megasquirt_gp40Descriptor struct {
	*descriptor.Message
	EGOcor9  *descriptor.Signal
	EGOcor10 *descriptor.Signal
	EGOcor11 *descriptor.Signal
	EGOcor12 *descriptor.Signal
}

type megasquirt_gp41Descriptor struct {
	*descriptor.Message
	EGOcor13 *descriptor.Signal
	EGOcor14 *descriptor.Signal
	EGOcor15 *descriptor.Signal
	EGOcor16 *descriptor.Signal
}

type megasquirt_gp42Descriptor struct {
	*descriptor.Message
	VSS1 *descriptor.Signal
	VSS2 *descriptor.Signal
	VSS3 *descriptor.Signal
	VSS4 *descriptor.Signal
}

type megasquirt_gp43Descriptor struct {
	*descriptor.Message
	synccnt    *descriptor.Signal
	syncreason *descriptor.Signal
	sd_filenum *descriptor.Signal
	sd_error   *descriptor.Signal
	sd_phase   *descriptor.Signal
	sd_status  *descriptor.Signal
	timing_err *descriptor.Signal
}

type megasquirt_gp44Descriptor struct {
	*descriptor.Message
	vvt_ang1 *descriptor.Signal
	vvt_ang2 *descriptor.Signal
	vvt_ang3 *descriptor.Signal
	vvt_ang4 *descriptor.Signal
}

type megasquirt_gp45Descriptor struct {
	*descriptor.Message
	vvt_target1 *descriptor.Signal
	vvt_target2 *descriptor.Signal
	vvt_target3 *descriptor.Signal
	vvt_target4 *descriptor.Signal
}

type megasquirt_gp46Descriptor struct {
	*descriptor.Message
	vvt_duty1      *descriptor.Signal
	vvt_duty2      *descriptor.Signal
	vvt_duty3      *descriptor.Signal
	vvt_duty4      *descriptor.Signal
	inj_timing_pri *descriptor.Signal
	inj_timing_sec *descriptor.Signal
}

type megasquirt_gp47Descriptor struct {
	*descriptor.Message
	fuel_pct  *descriptor.Signal
	tps_accel *descriptor.Signal
	SS1       *descriptor.Signal
	SS2       *descriptor.Signal
}

type megasquirt_gp48Descriptor struct {
	*descriptor.Message
	knock_cyl1 *descriptor.Signal
	knock_cyl2 *descriptor.Signal
	knock_cyl3 *descriptor.Signal
	knock_cyl4 *descriptor.Signal
	knock_cyl5 *descriptor.Signal
	knock_cyl6 *descriptor.Signal
	knock_cyl7 *descriptor.Signal
	knock_cyl8 *descriptor.Signal
}

type megasquirt_gp49Descriptor struct {
	*descriptor.Message
	knock_cyl9  *descriptor.Signal
	knock_cyl10 *descriptor.Signal
	knock_cyl11 *descriptor.Signal
	knock_cyl12 *descriptor.Signal
	knock_cyl13 *descriptor.Signal
	knock_cyl14 *descriptor.Signal
	knock_cyl15 *descriptor.Signal
	knock_cyl16 *descriptor.Signal
}

type megasquirt_gp50Descriptor struct {
	*descriptor.Message
	map_accel     *descriptor.Signal
	total_accel   *descriptor.Signal
	launch_timer  *descriptor.Signal
	launch_retard *descriptor.Signal
}

type megasquirt_gp51Descriptor struct {
	*descriptor.Message
	porta         *descriptor.Signal
	portb         *descriptor.Signal
	porteh        *descriptor.Signal
	portk         *descriptor.Signal
	portmj        *descriptor.Signal
	portp         *descriptor.Signal
	portt         *descriptor.Signal
	cel_errorcode *descriptor.Signal
}

type megasquirt_gp52Descriptor struct {
	*descriptor.Message
	canin1   *descriptor.Signal
	canin2   *descriptor.Signal
	canout   *descriptor.Signal
	knk_rtd  *descriptor.Signal
	fuelflow *descriptor.Signal
	fuelcons *descriptor.Signal
}

type megasquirt_gp53Descriptor struct {
	*descriptor.Message
	fuel_press1 *descriptor.Signal
	fuel_press2 *descriptor.Signal
	fuel_temp1  *descriptor.Signal
	fuel_temp2  *descriptor.Signal
}

type megasquirt_gp54Descriptor struct {
	*descriptor.Message
	batt_cur   *descriptor.Signal
	cel_status *descriptor.Signal
	fp_duty    *descriptor.Signal
	alt_duty   *descriptor.Signal
	load_duty  *descriptor.Signal
	alt_targv  *descriptor.Signal
}

type megasquirt_gp55Descriptor struct {
	*descriptor.Message
	looptime      *descriptor.Signal
	fueltemp_cor  *descriptor.Signal
	fuelpress_cor *descriptor.Signal
	ltt_cor       *descriptor.Signal
	sp1           *descriptor.Signal
}

type megasquirt_gp56Descriptor struct {
	*descriptor.Message
	tc_retard   *descriptor.Signal
	cel_retard  *descriptor.Signal
	fc_retard   *descriptor.Signal
	als_addfuel *descriptor.Signal
}

type megasquirt_gp57Descriptor struct {
	*descriptor.Message
	base_advance     *descriptor.Signal
	idle_cor_advance *descriptor.Signal
	mat_retard       *descriptor.Signal
	flex_advance     *descriptor.Signal
}

type megasquirt_gp58Descriptor struct {
	*descriptor.Message
	adv1 *descriptor.Signal
	adv2 *descriptor.Signal
	adv3 *descriptor.Signal
	adv4 *descriptor.Signal
}

type megasquirt_gp59Descriptor struct {
	*descriptor.Message
	revlim_retard *descriptor.Signal
	als_timing    *descriptor.Signal
	ext_advance   *descriptor.Signal
	deadtime1     *descriptor.Signal
}

type megasquirt_gp60Descriptor struct {
	*descriptor.Message
	launch_timing    *descriptor.Signal
	step3_timing     *descriptor.Signal
	vsslaunch_retard *descriptor.Signal
	cel_status2      *descriptor.Signal
}

type megasquirt_gp61Descriptor struct {
	*descriptor.Message
	gps_latdeg  *descriptor.Signal
	gps_latmin  *descriptor.Signal
	gps_latmmin *descriptor.Signal
	gps_londeg  *descriptor.Signal
	gps_lonmin  *descriptor.Signal
	gps_lonmmin *descriptor.Signal
}

type megasquirt_gp62Descriptor struct {
	*descriptor.Message
	gps_west   *descriptor.Signal
	gps_altk   *descriptor.Signal
	gps_altm   *descriptor.Signal
	gps_speed  *descriptor.Signal
	gps_course *descriptor.Signal
}

type megasquirt_gp63Descriptor struct {
	*descriptor.Message
	generic_pid_duty1 *descriptor.Signal
	generic_pid_duty2 *descriptor.Signal
	spare63_1         *descriptor.Signal
	spare63_2         *descriptor.Signal
	spare63_3         *descriptor.Signal
}

// Database returns the megasquirt database descriptor.
func (md *MessagesDescriptor) Database() *descriptor.Database {
	return d
}

var nd = &NodesDescriptor{}

var md = &MessagesDescriptor{
	megasquirt_gp0: &megasquirt_gp0Descriptor{
		Message: d.Messages[0],
		seconds: d.Messages[0].Signals[0],
		pw1:     d.Messages[0].Signals[1],
		pw2:     d.Messages[0].Signals[2],
		rpm:     d.Messages[0].Signals[3],
	},
	megasquirt_gp1: &megasquirt_gp1Descriptor{
		Message:  d.Messages[1],
		adv_deg:  d.Messages[1].Signals[0],
		squirt:   d.Messages[1].Signals[1],
		engine:   d.Messages[1].Signals[2],
		afrtgt1:  d.Messages[1].Signals[3],
		afrtgt2:  d.Messages[1].Signals[4],
		wbo2_en1: d.Messages[1].Signals[5],
		wbo2_en2: d.Messages[1].Signals[6],
	},
	megasquirt_gp2: &megasquirt_gp2Descriptor{
		Message: d.Messages[2],
		baro:    d.Messages[2].Signals[0],
		map1:    d.Messages[2].Signals[1],
		mat:     d.Messages[2].Signals[2],
		clt:     d.Messages[2].Signals[3],
	},
	megasquirt_gp3: &megasquirt_gp3Descriptor{
		Message:  d.Messages[3],
		tps:      d.Messages[3].Signals[0],
		batt:     d.Messages[3].Signals[1],
		afr1_old: d.Messages[3].Signals[2],
		afr2_old: d.Messages[3].Signals[3],
	},
	megasquirt_gp4: &megasquirt_gp4Descriptor{
		Message: d.Messages[4],
		knock:   d.Messages[4].Signals[0],
		egocor1: d.Messages[4].Signals[1],
		egocor2: d.Messages[4].Signals[2],
		aircor:  d.Messages[4].Signals[3],
	},
	megasquirt_gp5: &megasquirt_gp5Descriptor{
		Message:    d.Messages[5],
		warmcor:    d.Messages[5].Signals[0],
		tpsaccel:   d.Messages[5].Signals[1],
		tpsfuelcut: d.Messages[5].Signals[2],
		barocor:    d.Messages[5].Signals[3],
	},
	megasquirt_gp6: &megasquirt_gp6Descriptor{
		Message:  d.Messages[6],
		totalcor: d.Messages[6].Signals[0],
		ve1:      d.Messages[6].Signals[1],
		ve2:      d.Messages[6].Signals[2],
		iacstep:  d.Messages[6].Signals[3],
	},
	megasquirt_gp7: &megasquirt_gp7Descriptor{
		Message:      d.Messages[7],
		cold_adv_deg: d.Messages[7].Signals[0],
		TPSdot:       d.Messages[7].Signals[1],
		MAPdot:       d.Messages[7].Signals[2],
		RPMdot:       d.Messages[7].Signals[3],
	},
	megasquirt_gp8: &megasquirt_gp8Descriptor{
		Message:  d.Messages[8],
		MAFload:  d.Messages[8].Signals[0],
		fuelload: d.Messages[8].Signals[1],
		fuelcor:  d.Messages[8].Signals[2],
		MAF:      d.Messages[8].Signals[3],
	},
	megasquirt_gp9: &megasquirt_gp9Descriptor{
		Message:   d.Messages[9],
		egoV1:     d.Messages[9].Signals[0],
		egoV2:     d.Messages[9].Signals[1],
		dwell:     d.Messages[9].Signals[2],
		dwell_trl: d.Messages[9].Signals[3],
	},
	megasquirt_gp10: &megasquirt_gp10Descriptor{
		Message: d.Messages[10],
		status1: d.Messages[10].Signals[0],
		status2: d.Messages[10].Signals[1],
		status3: d.Messages[10].Signals[2],
		status4: d.Messages[10].Signals[3],
		status5: d.Messages[10].Signals[4],
		status6: d.Messages[10].Signals[5],
		status7: d.Messages[10].Signals[6],
	},
	megasquirt_gp11: &megasquirt_gp11Descriptor{
		Message:   d.Messages[11],
		fuelload2: d.Messages[11].Signals[0],
		ignload:   d.Messages[11].Signals[1],
		ignload2:  d.Messages[11].Signals[2],
		airtemp:   d.Messages[11].Signals[3],
	},
	megasquirt_gp12: &megasquirt_gp12Descriptor{
		Message:   d.Messages[12],
		wallfuel1: d.Messages[12].Signals[0],
		wallfuel2: d.Messages[12].Signals[1],
	},
	megasquirt_gp13: &megasquirt_gp13Descriptor{
		Message:  d.Messages[13],
		sensors1: d.Messages[13].Signals[0],
		sensors2: d.Messages[13].Signals[1],
		sensors3: d.Messages[13].Signals[2],
		sensors4: d.Messages[13].Signals[3],
	},
	megasquirt_gp14: &megasquirt_gp14Descriptor{
		Message:  d.Messages[14],
		sensors5: d.Messages[14].Signals[0],
		sensors6: d.Messages[14].Signals[1],
		sensors7: d.Messages[14].Signals[2],
		sensors8: d.Messages[14].Signals[3],
	},
	megasquirt_gp15: &megasquirt_gp15Descriptor{
		Message:   d.Messages[15],
		sensors9:  d.Messages[15].Signals[0],
		sensors10: d.Messages[15].Signals[1],
		sensors11: d.Messages[15].Signals[2],
		sensors12: d.Messages[15].Signals[3],
	},
	megasquirt_gp16: &megasquirt_gp16Descriptor{
		Message:   d.Messages[16],
		sensors13: d.Messages[16].Signals[0],
		sensors14: d.Messages[16].Signals[1],
		sensors15: d.Messages[16].Signals[2],
		sensors16: d.Messages[16].Signals[3],
	},
	megasquirt_gp17: &megasquirt_gp17Descriptor{
		Message:      d.Messages[17],
		boost_targ_1: d.Messages[17].Signals[0],
		boost_targ_2: d.Messages[17].Signals[1],
		boostduty:    d.Messages[17].Signals[2],
		boostduty2:   d.Messages[17].Signals[3],
		maf_volts:    d.Messages[17].Signals[4],
	},
	megasquirt_gp18: &megasquirt_gp18Descriptor{
		Message: d.Messages[18],
		pwseq1:  d.Messages[18].Signals[0],
		pwseq2:  d.Messages[18].Signals[1],
		pwseq3:  d.Messages[18].Signals[2],
		pwseq4:  d.Messages[18].Signals[3],
	},
	megasquirt_gp19: &megasquirt_gp19Descriptor{
		Message: d.Messages[19],
		pwseq5:  d.Messages[19].Signals[0],
		pwseq6:  d.Messages[19].Signals[1],
		pwseq7:  d.Messages[19].Signals[2],
		pwseq8:  d.Messages[19].Signals[3],
	},
	megasquirt_gp20: &megasquirt_gp20Descriptor{
		Message: d.Messages[20],
		pwseq9:  d.Messages[20].Signals[0],
		pwseq10: d.Messages[20].Signals[1],
		pwseq11: d.Messages[20].Signals[2],
		pwseq12: d.Messages[20].Signals[3],
	},
	megasquirt_gp21: &megasquirt_gp21Descriptor{
		Message: d.Messages[21],
		pwseq13: d.Messages[21].Signals[0],
		pwseq14: d.Messages[21].Signals[1],
		pwseq15: d.Messages[21].Signals[2],
		pwseq16: d.Messages[21].Signals[3],
	},
	megasquirt_gp22: &megasquirt_gp22Descriptor{
		Message: d.Messages[22],
		egt1:    d.Messages[22].Signals[0],
		egt2:    d.Messages[22].Signals[1],
		egt3:    d.Messages[22].Signals[2],
		egt4:    d.Messages[22].Signals[3],
	},
	megasquirt_gp23: &megasquirt_gp23Descriptor{
		Message: d.Messages[23],
		egt5:    d.Messages[23].Signals[0],
		egt6:    d.Messages[23].Signals[1],
		egt7:    d.Messages[23].Signals[2],
		egt8:    d.Messages[23].Signals[3],
	},
	megasquirt_gp24: &megasquirt_gp24Descriptor{
		Message: d.Messages[24],
		egt9:    d.Messages[24].Signals[0],
		egt10:   d.Messages[24].Signals[1],
		egt11:   d.Messages[24].Signals[2],
		egt12:   d.Messages[24].Signals[3],
	},
	megasquirt_gp25: &megasquirt_gp25Descriptor{
		Message: d.Messages[25],
		egt13:   d.Messages[25].Signals[0],
		egt14:   d.Messages[25].Signals[1],
		egt15:   d.Messages[25].Signals[2],
		egt16:   d.Messages[25].Signals[3],
	},
	megasquirt_gp26: &megasquirt_gp26Descriptor{
		Message:           d.Messages[26],
		nitrous1_duty:     d.Messages[26].Signals[0],
		nitrous2_duty:     d.Messages[26].Signals[1],
		nitrous_timer_out: d.Messages[26].Signals[2],
		n2o_addfuel:       d.Messages[26].Signals[3],
		n2o_retard:        d.Messages[26].Signals[4],
	},
	megasquirt_gp27: &megasquirt_gp27Descriptor{
		Message:   d.Messages[27],
		canpwmin1: d.Messages[27].Signals[0],
		canpwmin2: d.Messages[27].Signals[1],
		canpwmin3: d.Messages[27].Signals[2],
		canpwmin4: d.Messages[27].Signals[3],
	},
	megasquirt_gp28: &megasquirt_gp28Descriptor{
		Message:          d.Messages[28],
		cl_idle_targ_rpm: d.Messages[28].Signals[0],
		tpsadc:           d.Messages[28].Signals[1],
		eaeload:          d.Messages[28].Signals[2],
		afrload:          d.Messages[28].Signals[3],
	},
	megasquirt_gp29: &megasquirt_gp29Descriptor{
		Message:  d.Messages[29],
		EAEfcor1: d.Messages[29].Signals[0],
		EAEfcor2: d.Messages[29].Signals[1],
		VSS1dot:  d.Messages[29].Signals[2],
		VSS2dot:  d.Messages[29].Signals[3],
	},
	megasquirt_gp30: &megasquirt_gp30Descriptor{
		Message:      d.Messages[30],
		accelx:       d.Messages[30].Signals[0],
		accely:       d.Messages[30].Signals[1],
		accelz:       d.Messages[30].Signals[2],
		stream_level: d.Messages[30].Signals[3],
		water_duty:   d.Messages[30].Signals[4],
	},
	megasquirt_gp31: &megasquirt_gp31Descriptor{
		Message: d.Messages[31],
		AFR1:    d.Messages[31].Signals[0],
		AFR2:    d.Messages[31].Signals[1],
		AFR3:    d.Messages[31].Signals[2],
		AFR4:    d.Messages[31].Signals[3],
		AFR5:    d.Messages[31].Signals[4],
		AFR6:    d.Messages[31].Signals[5],
		AFR7:    d.Messages[31].Signals[6],
		AFR8:    d.Messages[31].Signals[7],
	},
	megasquirt_gp32: &megasquirt_gp32Descriptor{
		Message: d.Messages[32],
		AFR9:    d.Messages[32].Signals[0],
		AFR10:   d.Messages[32].Signals[1],
		AFR11:   d.Messages[32].Signals[2],
		AFR12:   d.Messages[32].Signals[3],
		AFR13:   d.Messages[32].Signals[4],
		AFR14:   d.Messages[32].Signals[5],
		AFR15:   d.Messages[32].Signals[6],
		AFR16:   d.Messages[32].Signals[7],
	},
	megasquirt_gp33: &megasquirt_gp33Descriptor{
		Message:   d.Messages[33],
		duty_pwm1: d.Messages[33].Signals[0],
		duty_pwm2: d.Messages[33].Signals[1],
		duty_pwm3: d.Messages[33].Signals[2],
		duty_pwm4: d.Messages[33].Signals[3],
		duty_pwm5: d.Messages[33].Signals[4],
		duty_pwm6: d.Messages[33].Signals[5],
		gear:      d.Messages[33].Signals[6],
		status8:   d.Messages[33].Signals[7],
	},
	megasquirt_gp34: &megasquirt_gp34Descriptor{
		Message: d.Messages[34],
		EGOv1:   d.Messages[34].Signals[0],
		EGOv2:   d.Messages[34].Signals[1],
		EGOv3:   d.Messages[34].Signals[2],
		EGOv4:   d.Messages[34].Signals[3],
	},
	megasquirt_gp35: &megasquirt_gp35Descriptor{
		Message: d.Messages[35],
		EGOv5:   d.Messages[35].Signals[0],
		EGOv6:   d.Messages[35].Signals[1],
		EGOv7:   d.Messages[35].Signals[2],
		EGOv8:   d.Messages[35].Signals[3],
	},
	megasquirt_gp36: &megasquirt_gp36Descriptor{
		Message: d.Messages[36],
		EGOv9:   d.Messages[36].Signals[0],
		EGOv10:  d.Messages[36].Signals[1],
		EGOv11:  d.Messages[36].Signals[2],
		EGOv12:  d.Messages[36].Signals[3],
	},
	megasquirt_gp37: &megasquirt_gp37Descriptor{
		Message: d.Messages[37],
		EGOv13:  d.Messages[37].Signals[0],
		EGOv14:  d.Messages[37].Signals[1],
		EGOv15:  d.Messages[37].Signals[2],
		EGOv16:  d.Messages[37].Signals[3],
	},
	megasquirt_gp38: &megasquirt_gp38Descriptor{
		Message: d.Messages[38],
		EGOcor1: d.Messages[38].Signals[0],
		EGOcor2: d.Messages[38].Signals[1],
		EGOcor3: d.Messages[38].Signals[2],
		EGOcor4: d.Messages[38].Signals[3],
	},
	megasquirt_gp39: &megasquirt_gp39Descriptor{
		Message: d.Messages[39],
		EGOcor5: d.Messages[39].Signals[0],
		EGOcor6: d.Messages[39].Signals[1],
		EGOcor7: d.Messages[39].Signals[2],
		EGOcor8: d.Messages[39].Signals[3],
	},
	megasquirt_gp40: &megasquirt_gp40Descriptor{
		Message:  d.Messages[40],
		EGOcor9:  d.Messages[40].Signals[0],
		EGOcor10: d.Messages[40].Signals[1],
		EGOcor11: d.Messages[40].Signals[2],
		EGOcor12: d.Messages[40].Signals[3],
	},
	megasquirt_gp41: &megasquirt_gp41Descriptor{
		Message:  d.Messages[41],
		EGOcor13: d.Messages[41].Signals[0],
		EGOcor14: d.Messages[41].Signals[1],
		EGOcor15: d.Messages[41].Signals[2],
		EGOcor16: d.Messages[41].Signals[3],
	},
	megasquirt_gp42: &megasquirt_gp42Descriptor{
		Message: d.Messages[42],
		VSS1:    d.Messages[42].Signals[0],
		VSS2:    d.Messages[42].Signals[1],
		VSS3:    d.Messages[42].Signals[2],
		VSS4:    d.Messages[42].Signals[3],
	},
	megasquirt_gp43: &megasquirt_gp43Descriptor{
		Message:    d.Messages[43],
		synccnt:    d.Messages[43].Signals[0],
		syncreason: d.Messages[43].Signals[1],
		sd_filenum: d.Messages[43].Signals[2],
		sd_error:   d.Messages[43].Signals[3],
		sd_phase:   d.Messages[43].Signals[4],
		sd_status:  d.Messages[43].Signals[5],
		timing_err: d.Messages[43].Signals[6],
	},
	megasquirt_gp44: &megasquirt_gp44Descriptor{
		Message:  d.Messages[44],
		vvt_ang1: d.Messages[44].Signals[0],
		vvt_ang2: d.Messages[44].Signals[1],
		vvt_ang3: d.Messages[44].Signals[2],
		vvt_ang4: d.Messages[44].Signals[3],
	},
	megasquirt_gp45: &megasquirt_gp45Descriptor{
		Message:     d.Messages[45],
		vvt_target1: d.Messages[45].Signals[0],
		vvt_target2: d.Messages[45].Signals[1],
		vvt_target3: d.Messages[45].Signals[2],
		vvt_target4: d.Messages[45].Signals[3],
	},
	megasquirt_gp46: &megasquirt_gp46Descriptor{
		Message:        d.Messages[46],
		vvt_duty1:      d.Messages[46].Signals[0],
		vvt_duty2:      d.Messages[46].Signals[1],
		vvt_duty3:      d.Messages[46].Signals[2],
		vvt_duty4:      d.Messages[46].Signals[3],
		inj_timing_pri: d.Messages[46].Signals[4],
		inj_timing_sec: d.Messages[46].Signals[5],
	},
	megasquirt_gp47: &megasquirt_gp47Descriptor{
		Message:   d.Messages[47],
		fuel_pct:  d.Messages[47].Signals[0],
		tps_accel: d.Messages[47].Signals[1],
		SS1:       d.Messages[47].Signals[2],
		SS2:       d.Messages[47].Signals[3],
	},
	megasquirt_gp48: &megasquirt_gp48Descriptor{
		Message:    d.Messages[48],
		knock_cyl1: d.Messages[48].Signals[0],
		knock_cyl2: d.Messages[48].Signals[1],
		knock_cyl3: d.Messages[48].Signals[2],
		knock_cyl4: d.Messages[48].Signals[3],
		knock_cyl5: d.Messages[48].Signals[4],
		knock_cyl6: d.Messages[48].Signals[5],
		knock_cyl7: d.Messages[48].Signals[6],
		knock_cyl8: d.Messages[48].Signals[7],
	},
	megasquirt_gp49: &megasquirt_gp49Descriptor{
		Message:     d.Messages[49],
		knock_cyl9:  d.Messages[49].Signals[0],
		knock_cyl10: d.Messages[49].Signals[1],
		knock_cyl11: d.Messages[49].Signals[2],
		knock_cyl12: d.Messages[49].Signals[3],
		knock_cyl13: d.Messages[49].Signals[4],
		knock_cyl14: d.Messages[49].Signals[5],
		knock_cyl15: d.Messages[49].Signals[6],
		knock_cyl16: d.Messages[49].Signals[7],
	},
	megasquirt_gp50: &megasquirt_gp50Descriptor{
		Message:       d.Messages[50],
		map_accel:     d.Messages[50].Signals[0],
		total_accel:   d.Messages[50].Signals[1],
		launch_timer:  d.Messages[50].Signals[2],
		launch_retard: d.Messages[50].Signals[3],
	},
	megasquirt_gp51: &megasquirt_gp51Descriptor{
		Message:       d.Messages[51],
		porta:         d.Messages[51].Signals[0],
		portb:         d.Messages[51].Signals[1],
		porteh:        d.Messages[51].Signals[2],
		portk:         d.Messages[51].Signals[3],
		portmj:        d.Messages[51].Signals[4],
		portp:         d.Messages[51].Signals[5],
		portt:         d.Messages[51].Signals[6],
		cel_errorcode: d.Messages[51].Signals[7],
	},
	megasquirt_gp52: &megasquirt_gp52Descriptor{
		Message:  d.Messages[52],
		canin1:   d.Messages[52].Signals[0],
		canin2:   d.Messages[52].Signals[1],
		canout:   d.Messages[52].Signals[2],
		knk_rtd:  d.Messages[52].Signals[3],
		fuelflow: d.Messages[52].Signals[4],
		fuelcons: d.Messages[52].Signals[5],
	},
	megasquirt_gp53: &megasquirt_gp53Descriptor{
		Message:     d.Messages[53],
		fuel_press1: d.Messages[53].Signals[0],
		fuel_press2: d.Messages[53].Signals[1],
		fuel_temp1:  d.Messages[53].Signals[2],
		fuel_temp2:  d.Messages[53].Signals[3],
	},
	megasquirt_gp54: &megasquirt_gp54Descriptor{
		Message:    d.Messages[54],
		batt_cur:   d.Messages[54].Signals[0],
		cel_status: d.Messages[54].Signals[1],
		fp_duty:    d.Messages[54].Signals[2],
		alt_duty:   d.Messages[54].Signals[3],
		load_duty:  d.Messages[54].Signals[4],
		alt_targv:  d.Messages[54].Signals[5],
	},
	megasquirt_gp55: &megasquirt_gp55Descriptor{
		Message:       d.Messages[55],
		looptime:      d.Messages[55].Signals[0],
		fueltemp_cor:  d.Messages[55].Signals[1],
		fuelpress_cor: d.Messages[55].Signals[2],
		ltt_cor:       d.Messages[55].Signals[3],
		sp1:           d.Messages[55].Signals[4],
	},
	megasquirt_gp56: &megasquirt_gp56Descriptor{
		Message:     d.Messages[56],
		tc_retard:   d.Messages[56].Signals[0],
		cel_retard:  d.Messages[56].Signals[1],
		fc_retard:   d.Messages[56].Signals[2],
		als_addfuel: d.Messages[56].Signals[3],
	},
	megasquirt_gp57: &megasquirt_gp57Descriptor{
		Message:          d.Messages[57],
		base_advance:     d.Messages[57].Signals[0],
		idle_cor_advance: d.Messages[57].Signals[1],
		mat_retard:       d.Messages[57].Signals[2],
		flex_advance:     d.Messages[57].Signals[3],
	},
	megasquirt_gp58: &megasquirt_gp58Descriptor{
		Message: d.Messages[58],
		adv1:    d.Messages[58].Signals[0],
		adv2:    d.Messages[58].Signals[1],
		adv3:    d.Messages[58].Signals[2],
		adv4:    d.Messages[58].Signals[3],
	},
	megasquirt_gp59: &megasquirt_gp59Descriptor{
		Message:       d.Messages[59],
		revlim_retard: d.Messages[59].Signals[0],
		als_timing:    d.Messages[59].Signals[1],
		ext_advance:   d.Messages[59].Signals[2],
		deadtime1:     d.Messages[59].Signals[3],
	},
	megasquirt_gp60: &megasquirt_gp60Descriptor{
		Message:          d.Messages[60],
		launch_timing:    d.Messages[60].Signals[0],
		step3_timing:     d.Messages[60].Signals[1],
		vsslaunch_retard: d.Messages[60].Signals[2],
		cel_status2:      d.Messages[60].Signals[3],
	},
	megasquirt_gp61: &megasquirt_gp61Descriptor{
		Message:     d.Messages[61],
		gps_latdeg:  d.Messages[61].Signals[0],
		gps_latmin:  d.Messages[61].Signals[1],
		gps_latmmin: d.Messages[61].Signals[2],
		gps_londeg:  d.Messages[61].Signals[3],
		gps_lonmin:  d.Messages[61].Signals[4],
		gps_lonmmin: d.Messages[61].Signals[5],
	},
	megasquirt_gp62: &megasquirt_gp62Descriptor{
		Message:    d.Messages[62],
		gps_west:   d.Messages[62].Signals[0],
		gps_altk:   d.Messages[62].Signals[1],
		gps_altm:   d.Messages[62].Signals[2],
		gps_speed:  d.Messages[62].Signals[3],
		gps_course: d.Messages[62].Signals[4],
	},
	megasquirt_gp63: &megasquirt_gp63Descriptor{
		Message:           d.Messages[63],
		generic_pid_duty1: d.Messages[63].Signals[0],
		generic_pid_duty2: d.Messages[63].Signals[1],
		spare63_1:         d.Messages[63].Signals[2],
		spare63_2:         d.Messages[63].Signals[3],
		spare63_3:         d.Messages[63].Signals[4],
	},
}

var d = (*descriptor.Database)(&descriptor.Database{
	SourceFile: (string)("dbc/megasquirt.dbc"),
	Version:    (string)(""),
	Messages: ([]*descriptor.Message)([]*descriptor.Message{
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp0"),
			ID:          (uint32)(1520),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("seconds"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("s"),
					Description:       (string)("Seconds ECU has been on"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pw1"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Main pulsewidth bank 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pw2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Main pulsewidth bank 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("rpm"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("RPM"),
					Description:       (string)("Engine RPM"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp1"),
			ID:          (uint32)(1521),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("adv_deg"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg BTDC"),
					Description:       (string)("Final ignition spark advance"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("squirt"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Bitfield of batch fire injector events"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("engine"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Bitfield of engine status"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("afrtgt1"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("Bank 1 AFR target"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("afrtgt2"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("Bank 2 AFR target"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("wbo2_en1"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("wbo2_en2"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp2"),
			ID:          (uint32)(1522),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("baro"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa"),
					Description:       (string)("Barometric pressure"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("map1"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa"),
					Description:       (string)("Manifold air pressure"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("mat"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg F"),
					Description:       (string)("Manifold air temperature"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("clt"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg F"),
					Description:       (string)("Coolant temperature"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp3"),
			ID:          (uint32)(1523),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("tps"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Throttle position"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("batt"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Battery voltage"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("afr1_old"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("afr2_old"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp4"),
			ID:          (uint32)(1524),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Indication of knock input"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egocor1"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO bank 1 correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egocor2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO bank 2 correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("aircor"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Air density correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp5"),
			ID:          (uint32)(1525),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("warmcor"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Warmup correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("tpsaccel"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("TPS-based acceleration"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("tpsfuelcut"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("TPS-based fuel cut"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("barocor"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Barometric fuel correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp6"),
			ID:          (uint32)(1526),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("totalcor"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Total fuel correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("ve1"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("VE table/bank 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("ve2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("VE table/bank 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("iacstep"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("step"),
					Description:       (string)("Stepper idle step number"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp7"),
			ID:          (uint32)(1527),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("cold_adv_deg"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Cold advance"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("TPSdot"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%/s"),
					Description:       (string)("Rate of change of TPS"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("MAPdot"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa/s"),
					Description:       (string)("Rate of change of MAP"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("RPMdot"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("RPM/s"),
					Description:       (string)("Rate of change of RPM"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp8"),
			ID:          (uint32)(1528),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("MAFload"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Synthetic 'load' from MAF"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuelload"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("'Load' used for fuel table lookup"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuelcor"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Adjustment to fuel from Flex"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("MAF"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("g/s"),
					Description:       (string)("Mass air flow"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp9"),
			ID:          (uint32)(1529),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egoV1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egoV2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("dwell"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Main ignition dwell"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("dwell_trl"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Trailing ignition dwell"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp10"),
			ID:          (uint32)(1530),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ECU status bitfield 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ECU status bitfield 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status3"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ECU status bitfield 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status4"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status5"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status6"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ECU status bitfield 6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status7"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ECU status bitfield 7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp11"),
			ID:          (uint32)(1531),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuelload2"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("'Load' used for modifier fuel table lookup"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("ignload"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("'Load' used for ignition table lookup"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("ignload2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("'Load' used for modifier ignition table lookup"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("airtemp"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("Estimated intake air temperature"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp12"),
			ID:          (uint32)(1532),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("wallfuel1"),
					Start:             (uint8)(7),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("us"),
					Description:       (string)("Calculated volume of fuel on intake walls from EAE - channel 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("wallfuel2"),
					Start:             (uint8)(39),
					Length:            (uint8)(32),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("us"),
					Description:       (string)("Calculated volume of fuel on intake walls from EAE - channel 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp13"),
			ID:          (uint32)(1533),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 1 (gpioadc0 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 2 (gpioadc1 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 3 (gpioadc2 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 4 (gpioadc3 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp14"),
			ID:          (uint32)(1534),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors5"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 5 (gpioadc4 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors6"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 6 (gpioadc5 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors7"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 7 (gpioadc6 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors8"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 8 (gpioadc7 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp15"),
			ID:          (uint32)(1535),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors9"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 19 (adc6 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors10"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 10 (adc7 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors11"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors12"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp16"),
			ID:          (uint32)(1536),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors13"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors14"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors15"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sensors16"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.01),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Generic sensor input 16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp17"),
			ID:          (uint32)(1537),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("boost_targ_1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa"),
					Description:       (string)("Target boost - channel 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("boost_targ_2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa"),
					Description:       (string)("Target boost - channel 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("boostduty"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Duty cycle on boost solenoid 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("boostduty2"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Duty cycle on boost solenoid 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("maf_volts"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("MAF voltage"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp18"),
			ID:          (uint32)(1538),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp19"),
			ID:          (uint32)(1539),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq5"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq6"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq7"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq8"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp20"),
			ID:          (uint32)(1540),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq9"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #9"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq10"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #10"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq11"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq12"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp21"),
			ID:          (uint32)(1541),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq13"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq14"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq15"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("pwseq16"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Sequential pulsewidth for cyl #16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp22"),
			ID:          (uint32)(1542),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp23"),
			ID:          (uint32)(1543),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt5"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt6"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt7"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt8"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp24"),
			ID:          (uint32)(1544),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt9"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #9"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt10"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #10"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt11"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt12"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp25"),
			ID:          (uint32)(1545),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt13"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt14"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt15"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("egt16"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("degF"),
					Description:       (string)("EGT cyl #16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp26"),
			ID:          (uint32)(1546),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("nitrous1_duty"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Duty cycle to nitrous solenoid 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("nitrous2_duty"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Duty cycle to nitrous solenoid 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("nitrous_timer_out"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("s"),
					Description:       (string)("Timer used internally for nitrous system"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("n2o_addfuel"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("Fuel pulsewidth added due to nitrous system"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("n2o_retard"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Timing retard due to nitrous system"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp27"),
			ID:          (uint32)(1547),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canpwmin1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("PWM period 1 from remote board"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canpwmin2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("PWM period 2 from remote board"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canpwmin3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("PWM period 3 from remote board"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canpwmin4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("PWM period 4 from remote board"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp28"),
			ID:          (uint32)(1548),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("cl_idle_targ_rpm"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("RPM"),
					Description:       (string)("Closed-loop idle target RPM"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("tpsadc"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ADC count from TPS"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("eaeload"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("'Load' used for EAE calc"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("afrload"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("'Load' used for AFR table lookups"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp29"),
			ID:          (uint32)(1549),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EAEfcor1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Fuel correction from EAE - channel 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EAEfcor2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Fuel correction from EAE - channel 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VSS1dot"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s^2"),
					Description:       (string)("Rate of change of VSS1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VSS2dot"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s^2"),
					Description:       (string)("Rate of change of VSS2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp30"),
			ID:          (uint32)(1550),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("accelx"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s^2"),
					Description:       (string)("External accelerometer X"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("accely"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s^2"),
					Description:       (string)("External accelerometer Y"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("accelz"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s^2"),
					Description:       (string)("External accelerometer Z"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("stream_level"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Volume level on audio input"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("water_duty"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Duty cycle to water injection solenoid"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp31"),
			ID:          (uint32)(1551),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("AFR cyl #2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR3"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR4"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR5"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR6"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR7"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR8"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp32"),
			ID:          (uint32)(1552),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR9"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #9"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR10"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #10"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR11"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR12"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR13"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR14"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR15"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("AFR16"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("AFR"),
					Description:       (string)("AFR cyl #16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp33"),
			ID:          (uint32)(1553),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("duty_pwm1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic PWM duty 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("duty_pwm2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic PWM duty 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("duty_pwm3"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic PWM duty 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("duty_pwm4"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic PWM duty 4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("duty_pwm5"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic PWM duty 5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("duty_pwm6"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic PWM duty 6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gear"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Current gear selected"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("status8"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("ECU status bitfield 8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp34"),
			ID:          (uint32)(1554),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp35"),
			ID:          (uint32)(1555),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv5"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv6"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv7"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv8"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp36"),
			ID:          (uint32)(1556),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv9"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #9"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv10"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #10"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv11"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv12"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp37"),
			ID:          (uint32)(1557),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv13"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv14"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv15"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOv16"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.00489),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Voltage from O2 cyl #16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp38"),
			ID:          (uint32)(1558),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp39"),
			ID:          (uint32)(1559),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor5"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor6"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor7"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor8"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp40"),
			ID:          (uint32)(1560),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor9"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#9"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor10"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#10"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor11"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor12"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp41"),
			ID:          (uint32)(1561),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor13"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor14"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor15"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("EGOcor16"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("EGO correction cyl#16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp42"),
			ID:          (uint32)(1562),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VSS1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s"),
					Description:       (string)("Vehicle speed 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VSS2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s"),
					Description:       (string)("Vehicle speed 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VSS3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s"),
					Description:       (string)("Vehicle speed 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("VSS4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s"),
					Description:       (string)("Vehicle speed 4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp43"),
			ID:          (uint32)(1563),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("synccnt"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Sync-loss counter"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("syncreason"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Sync-loss reason code"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sd_filenum"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("SDcard file number"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sd_error"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("SDcard error number"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sd_phase"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("SDcard internal code"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sd_status"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("SDcard status bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("timing_err"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Calculated error in ignition timing"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp44"),
			ID:          (uint32)(1564),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_ang1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT actual angle 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_ang2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT actual angle 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_ang3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT actual angle 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_ang4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT actual angle 4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp45"),
			ID:          (uint32)(1565),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_target1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT target angle 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_target2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT target angle 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_target3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT target angle 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_target4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("VVT target angle 4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp46"),
			ID:          (uint32)(1566),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_duty1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("VVT solenoid duty cycle 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_duty2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("VVT solenoid duty cycle 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_duty3"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("VVT solenoid duty cycle 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vvt_duty4"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("VVT solenoid duty cycle 4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("inj_timing_pri"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg BTDC"),
					Description:       (string)("Injection Timing Angle (Primary)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("inj_timing_sec"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg BTDC"),
					Description:       (string)("Injection Timing Angle (Secondary)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp47"),
			ID:          (uint32)(1567),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuel_pct"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Ethanol content of fuel from Flex sensor"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("tps_accel"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("TPSdot based accel"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("SS1"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("RPM"),
					Description:       (string)("Shaft Speed 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("SS2"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(10),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("RPM"),
					Description:       (string)("Shaft Speed 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp48"),
			ID:          (uint32)(1568),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl3"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl4"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl5"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #5"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl6"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #6"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl7"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #7"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl8"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #8"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp49"),
			ID:          (uint32)(1569),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl9"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #9"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl10"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #10"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl11"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #11"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl12"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #12"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl13"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #13"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl14"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #14"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl15"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #15"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knock_cyl16"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.4),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Knock % cyl #16"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp50"),
			ID:          (uint32)(1570),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("map_accel"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("MAPdot based accel"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("total_accel"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Total accel"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("launch_timer"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("s"),
					Description:       (string)("Timer for timed-launch retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("launch_retard"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Launch retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp51"),
			ID:          (uint32)(1571),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("porta"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portA bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("portb"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portB bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("porteh"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portE/portH bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("portk"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portK bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("portmj"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portM/portJ bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("portp"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portP bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("portt"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CPU portT bitfield"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("cel_errorcode"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CEL error code"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp52"),
			ID:          (uint32)(1572),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canin1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CAN input 1 bitfield (CAN port 1 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canin2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CAN input 2 bitfield (CAN port 2 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("canout"),
					Start:             (uint8)(23),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CAN output 1 bitfield (CAN port 3 on MS2)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("knk_rtd"),
					Start:             (uint8)(31),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Knock retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuelflow"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("cc/min"),
					Description:       (string)("Average fuel flow"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuelcons"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("l/km"),
					Description:       (string)("Average fuel consumption"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp53"),
			ID:          (uint32)(1573),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuel_press1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa"),
					Description:       (string)("Fuel pressure 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuel_press2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("kPa"),
					Description:       (string)("Fuel pressure 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuel_temp1"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg F"),
					Description:       (string)("Fuel temperature 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuel_temp2"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg F"),
					Description:       (string)("Fuel temperature 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp54"),
			ID:          (uint32)(1574),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("batt_cur"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Battery current (alternator system)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("cel_status"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CEL status bitfield 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fp_duty"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Fuel pump output duty"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("alt_duty"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Alternator field output duty"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("load_duty"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Alternator measured load-sense duty"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("alt_targv"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("V"),
					Description:       (string)("Alternator target voltage"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp55"),
			ID:          (uint32)(1575),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("looptime"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("us"),
					Description:       (string)("Main code loop execution time"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fueltemp_cor"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Fuel temperature correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fuelpress_cor"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Fuel pressure correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("ltt_cor"),
					Start:             (uint8)(55),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Long term trim correction"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("sp1"),
					Start:             (uint8)(63),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("Unused"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp56"),
			ID:          (uint32)(1576),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("tc_retard"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("def"),
					Description:       (string)("Traction control retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("cel_retard"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("CEL retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("fc_retard"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Fuel-cut (overrun) retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("als_addfuel"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("ms"),
					Description:       (string)("ALS added fuel"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp57"),
			ID:          (uint32)(1577),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("base_advance"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Base timing from tables"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("idle_cor_advance"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Idle correction advance"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("mat_retard"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("MAT retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("flex_advance"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Flex advance"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp58"),
			ID:          (uint32)(1578),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("adv1"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Timing lookup from table 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("adv2"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Timing lookup from table 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("adv3"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Timing lookup from table 3"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("adv4"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Timing lookup from table 4"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp59"),
			ID:          (uint32)(1579),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("revlim_retard"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Rev-limiter 'soft' retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("als_timing"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("ALS timing change"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("ext_advance"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("External advance (e.g. trans)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("deadtime1"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.001),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Injector deadtime in use (#1)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp60"),
			ID:          (uint32)(1580),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("launch_timing"),
					Start:             (uint8)(7),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Launch control timing"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("step3_timing"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("3-step timing"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("vsslaunch_retard"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("Wheel-speed based launch retard"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("cel_status2"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("CEL status bitfield 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp61"),
			ID:          (uint32)(1581),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_latdeg"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("GPS latitude degrees"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_latmin"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("min"),
					Description:       (string)("GPS latitude minutes"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_latmmin"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("mmin"),
					Description:       (string)("GPS latitude milli-minutes"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_londeg"),
					Start:             (uint8)(39),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("GPS longitude degrees"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_lonmin"),
					Start:             (uint8)(47),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("min"),
					Description:       (string)("GPS longitude minutes"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_lonmmin"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("mmin"),
					Description:       (string)("GPS longitude milli-minutes"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp62"),
			ID:          (uint32)(1582),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_west"),
					Start:             (uint8)(0),
					Length:            (uint8)(1),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)("GPS longitude west indicator bit (gps_outstatus)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_altk"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(true),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("km"),
					Description:       (string)("GPS altitude (km)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_altm"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m"),
					Description:       (string)("GPS altitude (m)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_speed"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("m/s"),
					Description:       (string)("GPS speed"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("gps_course"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("deg"),
					Description:       (string)("GPS course (heading)"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
		(*descriptor.Message)(&descriptor.Message{
			Name:        (string)("megasquirt_gp63"),
			ID:          (uint32)(1583),
			IsExtended:  (bool)(false),
			Length:      (uint8)(8),
			SendType:    (descriptor.SendType)(0),
			Description: (string)(""),
			Signals: ([]*descriptor.Signal)([]*descriptor.Signal{
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("generic_pid_duty1"),
					Start:             (uint8)(7),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic closed-loop duty 1"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("generic_pid_duty2"),
					Start:             (uint8)(15),
					Length:            (uint8)(8),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(0.392),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)("%"),
					Description:       (string)("Generic closed-loop duty 2"),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("spare63_1"),
					Start:             (uint8)(23),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("spare63_2"),
					Start:             (uint8)(39),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
				(*descriptor.Signal)(&descriptor.Signal{
					Name:              (string)("spare63_3"),
					Start:             (uint8)(55),
					Length:            (uint8)(16),
					IsBigEndian:       (bool)(true),
					IsSigned:          (bool)(false),
					IsMultiplexer:     (bool)(false),
					IsMultiplexed:     (bool)(false),
					MultiplexerValue:  (uint)(0),
					Offset:            (float64)(0),
					Scale:             (float64)(1),
					Min:               (float64)(0),
					Max:               (float64)(0),
					Unit:              (string)(""),
					Description:       (string)(""),
					ValueDescriptions: ([]*descriptor.ValueDescription)(nil),
					ReceiverNodes: ([]string)([]string{
						(string)("Vector__XXX"),
					}),
					DefaultValue: (int)(0),
				}),
			}),
			SenderNode: (string)("Vector__XXX"),
			CycleTime:  (time.Duration)(0),
			DelayTime:  (time.Duration)(0),
		}),
	}),
	Nodes: ([]*descriptor.Node)(nil),
})
