package playerdata

import (
	"github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"testing"
)

var p = new(common.Player)

func TestAimPunchAngle(t *testing.T) {
	apa := AimPunchAngle(p)
	if apa != (r3.Vector{X: 0, Y: 0, Z: 0}) {
		t.Fatalf("Test FAIL")
	}
}

func TestAimPunchAngleLevel(t *testing.T) {
	apa := AimPunchAngleLevel(&common.Player{})
	if apa != (r3.Vector{}) {
		t.Fatalf("Test FAIL")
	}
}

func TestBDucked(t *testing.T) {
	apa := BDucked(&common.Player{})
	if apa != false {
		t.Fatalf("Test FAIL")
	}
}

func TestBDucking(t *testing.T) {
	apa := BDucking(&common.Player{})
	if apa != false {
		t.Fatalf("Test FAIL")
	}
}

func TestWearingSuit(t *testing.T) {
	apa := WearingSuit(&common.Player{})
	if apa != false {
		t.Fatalf("Test FAIL")
	}
}

func TestAreaBits000(t *testing.T) {
	apa := AreaBits000(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestAreaBits001(t *testing.T) {
	apa := AreaBits001(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestVecVelocity0(t *testing.T) {
	apa := VecVelocity0(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestVecVelocity1(t *testing.T) {
	apa := VecVelocity1(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestVecVelocity2(t *testing.T) {
	apa := VecVelocity2(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestVecBaseVelocity(t *testing.T) {
	apa := VecBaseVelocity(&common.Player{})
	if apa != (r3.Vector{}) {
		t.Fatalf("Test FAIL")
	}
}

func TestTickBase(t *testing.T) {
	apa := TickBase(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestVecViewOffset(t *testing.T) {
	apa := VecViewOffset(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestNextThinkTick(t *testing.T) {
	apa := NextThinkTick(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestFOVRate(t *testing.T) {
	apa := FOVRate(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestFallVelocity(t *testing.T) {
	apa := FallVelocity(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestLastDuckTime(t *testing.T) {
	apa := LastDuckTime(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}

func TestViewPunchAngle(t *testing.T) {
	apa := ViewPunchAngle(&common.Player{})
	if apa != 0 {
		t.Fatalf("Test FAIL")
	}
}
