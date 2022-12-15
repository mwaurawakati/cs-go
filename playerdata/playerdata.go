package playerdata

import (
	//"fmt"
	//"time"

	"github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	//constants "github.com/markus-wa/demoinfocs-golang/v2/internal/constants"
	st "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/sendtables"
	"time"
)

func AimPunchAngle(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}

	val, _ := p.Entity.PropertyValue("localdata.m_Local.m_aimPunchAngle")
	return val.VectorVal

}

func AimPunchAngleLevel(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}

	val, _ := p.Entity.PropertyValue("localdata.m_Local.m_aimPunchAngleVel")
	return val.VectorVal

}

func BDucked(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "localdata.m_Local.m_bDucked")
}

func BDucking(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "localdata.m_Local.m_bDucking")
}

func WearingSuit(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "localdata.m_Local.m_bWearingSuit")
}

func AreaBits000(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_Local.m_chAreaBits.000")
}

func AreaBits001(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_Local.m_chAreaBits.001")
}

func VecVelocity0(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := (p.Entity.PropertyValue("localdata.m_vecVelocity[0]"))
	return float64(val.FloatVal)

}

func VecVelocity1(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_vecVelocity[1]")
	return float64(val.FloatVal)
}

func VecVelocity2(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_vecVelocity[2]")
	return float64(val.FloatVal)
}

func LaggedMovementValue(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_flLaggedMovementValue")
	return float64(val.FloatVal)
}

func VecBaseVelocity(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}

	val, _ := (p.Entity.PropertyValue("localdata.m_vecBaseVelocity"))
	return val.VectorVal
}

func TickBase(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_nTickBase")
}

func VecViewOffset(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_vecViewOffset[2]")
	return float64(val.FloatVal)
}

func NextThinkTick(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_nNextThinkTick")
}

func FOVRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_Local.m_flFOVRate")
	return float64(val.FloatVal)
}

func FallVelocity(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_Local.m_flFallVelocity")
	return float64(val.FloatVal)
}

func LastDuckTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_Local.m_flLastDuckTime")
	return float64(val.FloatVal)
}

func ViewPunchAngle(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}

	val, _ := p.Entity.PropertyValue("localdata.m_Local.m_viewPunchAngle")
	return val.VectorVal
}

func VecLadderNormal(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}

	val, _ := p.Entity.PropertyValue("m_vecLadderNormal")
	return val.VectorVal
}

func DeathTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("localdata.m_flDeathTime")
	return float64(val.FloatVal)
}

func LastWeapon(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_hLastWeapon")
}

func ArmorValue(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_ArmorValue")
}

func VecMaxs(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_Collision.m_vecMaxs")
	return float64(val.FloatVal)
}

func MaxSpeed(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flMaxspeed")
	return float64(val.FloatVal)
}

func VecMins(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_Collision.m_vecMins")
	return float64(val.FloatVal)
}

func LastHitGroup(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_LastHitGroup")
}

func EyeAngle0(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_angEyeAngles[0]")
	return float64(val.FloatVal)
}

func EyeAngle1(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_angEyeAngles[1]")
	return float64(val.FloatVal)
}

func NextAttack(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("bcc_localdata.m_flNextAttack")
	return float64(val.FloatVal)
}

func DuckOverride(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("cslocaldata.m_bDuckOverride")
	return float64(val.FloatVal)
}

func Stamina(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("cslocaldata.m_flStamina")
	return float64(val.FloatVal)
}

func VelocityModifier(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("cslocaldata.m_flVelocityModifier")
	return float64(val.FloatVal)
}

func ShotsFired(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "cslocaldata.m_iShotsFired")
}

func PhysicsFlags(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_afPhysicsFlags")
}

func VecOrigin(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("cslocaldata.m_vecOrigin")
	return float64(val.FloatVal)
}

func VecOrigin2(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("cslocaldata.m_vecOrigin[2]")
	return float64(val.FloatVal)
}

func DuckAmount(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flDuckAmount")
	return float64(val.FloatVal)
}

func DuckSpeed(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flDuckSpeed")
	return float64(val.FloatVal)
}

func FOVTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flFOVTime")
	return float64(val.FloatVal)
}

func FlashDuration(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flFlashDuration")
	return float64(val.FloatVal)
}

func FlashMaxAlpha(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flFlashMaxAlpha")
	return float64(val.FloatVal)
}

func GroundAccelLinearFracLastTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flGroundAccelLinearFracLastTime")
	return float64(val.FloatVal)
}

func LastMadeNoiseTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flLastMadeNoiseTime")
	return float64(val.FloatVal)
}

func LowerBodyYawTarget(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flLowerBodyYawTarget")
	return float64(val.FloatVal)
}

func TimeOfLastInjury(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_flTimeOfLastInjury")
	return float64(val.FloatVal)
}

func ColorCorrectionCtrl(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_hColorCorrectionCtrl")
}

func PhysicsCollisionState(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_vphysicsCollisionState")
}

func GroundEntity(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_hGroundEntity")
}

func PlayerPing(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_hPlayerPing")
}

func ForceBone(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_nForceBone")
}

func LastConcurrentKilled(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_nLastConcurrentKilled")
}

func HasMovedSinceSpawn(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bHasMovedSinceSpawn")
}

func IsHoldingLookAtWeapon(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bIsHoldingLookAtWeapon")
}

func IsLookingAtWeapon(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bIsLookingAtWeapon")
}

func ResumeZoom(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bResumeZoom")
}

func WaitForNoAttack(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bWaitForNoAttack")
}

func MolotovDamageTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_fMolotovDamageTime")
	return float64(val.FloatVal)
}

func MolotovUseTime(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_fMolotovUseTime")
	return float64(val.FloatVal)
}

func Strafing(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bStrafing")
}

func Spotted(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bSpotted")
}

func SpottedByMask000(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}
	return getBool(p.Entity, "m_bSpottedByMask.000")
}

func NLDVecOrigin2(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("csnonlocaldata.m_vecOrigin[2]")
	return float64(val.FloatVal)
}

func NLDVecOrigin(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("csnonlocaldata.m_vecOrigin")
	return float64(val.FloatVal)
}

func FOV(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_iFOV")
}

func DefaultFOV(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_iDefaultFOV")
}

func WaterLevel(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_nWaterLevel")
}

func LifeState(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_lifeState")
}

func LadderSurfaceProps(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_ladderSurfaceProps")
}

func FOVStart(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_iFOVStart")
}

func MoveState(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_iMoveState")
}

func PlayerState(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_iPlayerState")
}

func RelativeDirectionOfLastInjury(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_nRelativeDirectionOfLastInjury")
}

func TotalHitsOnServer(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_totalHitsOnServer")
}

func EFNoInterpParity(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_ubEFNoInterpParity")
}

func SimulationTime(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_flSimulationTime")
}

func AnimOverlay0PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.000.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay1PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.001.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay2PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.002.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay3PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.003.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay4PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.004.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay5PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.005.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay6PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.006.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay7PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.007.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay8PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.008.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay9PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.009.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay10PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.010.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay11PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.011.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay12PlaybackRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.012.m_flPlaybackRate")
	return float64(val.FloatVal)
}

func AnimOverlay0Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.000.m_nSequence")
}

func AnimOverlay1Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.001.m_nSequence")
}

func AnimOverlay2Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.001.m_nSequence")
}

func AnimOverlay3Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.003.m_nSequence")
}

func AnimOverlay4Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.004.m_nSequence")
}

func AnimOverlay5Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.005.m_nSequence")
}

func AnimOverlay6Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.006.m_nSequence")
}

func AnimOverlay7Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.007.m_nSequence")
}

func AnimOverlay8Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.008.m_nSequence")
}

func AnimOverlay9Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.009.m_nSequence")
}

func AnimOverlay10Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.010.m_nSequence")
}

func AnimOverlay11Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.011.m_nSequence")
}

func AnimOverlay12Sequence(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.012.m_nSequence")
}

func AnimOverlay1Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.001.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay0Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.000.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay2Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.002.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay3Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.003.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay4Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.004.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay5Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.005.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay6Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.006.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay7Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.007.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay8Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.008.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay9Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.009.m_flWeight")
	return float64(val.FloatVal)
}
func AnimOverlay10Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.010.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay11Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.011.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay12Weight(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.012.m_flWeight")
	return float64(val.FloatVal)
}

func AnimOverlay0Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.000.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay1Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.001.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay2Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.002.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay3Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.003.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay4Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.004.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay5Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.006.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay6Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.006.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay7Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.007.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay8Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.008.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay9Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.009.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay10Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.010.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay11Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.011.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay12Cycle(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.012.m_flCycle")
	return float64(val.FloatVal)
}

func AnimOverlay0Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.000.m_nOrder")
}

func AnimOverlay1Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.001.m_nOrder")
}

func AnimOverlay2Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.002.m_nOrder")
}

func AnimOverlay3Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.003.m_nOrder")
}

func AnimOverlay4Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.004.m_nOrder")
}

func AnimOverlay5Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.005.m_nOrder")
}

func AnimOverlay6Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.006.m_nOrder")
}

func AnimOverlay7Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.007.m_nOrder")
}

func AnimOverlay8Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.008.m_nOrder")
}

func AnimOverlay9Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.009.m_nOrder")
}

func AnimOverlay10Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.010.m_nOrder")
}

func AnimOverlay11Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.011.m_nOrder")
}

func AnimOverlay12Order(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "m_AnimOverlay.012.m_nOrder")
}

func AnimOverlay0WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.000.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay1WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.001.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay2WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.002.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay3WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.003.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay4WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.004.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay5WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.005.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay6WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.006.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay7WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.007.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay8WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.008.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay9WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.009.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay10WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.010.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay11WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.011.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func AnimOverlay12WeightDeltaRate(p *common.Player) float64 {
	if p.Entity == nil {
		return float64(0)
	}

	val, _ := p.Entity.PropertyValue("m_AnimOverlay.012.m_flWeightDeltaRate")
	return float64(val.FloatVal)
}

func getInt(entity st.Entity, propName string) int {
	if entity == nil {
		return 0
	}

	val, _ := entity.PropertyValue(propName)
	return val.IntVal
}

func getFloat(entity st.Entity, propName string) float32 {
	if entity == nil {
		return 0
	}

	val, _ := entity.PropertyValue(propName)
	return val.FloatVal
}

func getString(entity st.Entity, propName string) string {
	if entity == nil {
		return ""
	}

	val, _ := entity.PropertyValue(propName)
	return val.StringVal
}

func getBool(entity st.Entity, propName string) bool {
	if entity == nil {
		return false
	}

	val, _ := entity.PropertyValue(propName)
	return val.BoolVal()
}

func getVector(entity st.Entity, propName string) r3.Vector {
	if entity == nil {
		return r3.Vector{}
	}

	val, _ := entity.PropertyValue(propName)
	return val.VectorVal
}

func EntIndex(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_Local.m_audio.entIndex")
}

func LocalBits(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}
	return getInt(p.Entity, "localdata.m_Local.m_audio.localBits")
}

func LocalSound1(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[1]")
}

func LocalSound2(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[2]")
}

func LocalSound3(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[3]")
}

func LocalSound4(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[4]")
}

func LocalSound5(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[5]")
}

func LocalSound6(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[6]")
}

func LocalSound7(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[7]")
}

func LocalSound0(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_audio.localSound[0]")
}

func SoundscapeIndex(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_audio.soundscapeIndex")
}

func AllowAutoMovement(p *common.Player) bool {
	return getBool(p.Entity, "localdata.m_Local.m_bAllowAutoMovement")
}

func InDuckJump(p *common.Player) bool {
	return getBool(p.Entity, "localdata.m_Local.m_bInDuckJump")
}

func DrawViewmodel(p *common.Player) bool {
	return getBool(p.Entity, "localdata.m_Local.m_bDrawViewmodel")
}

func Poisoned(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_bPoisoned")

	return val.BoolVal()
}

func StepSize(p *common.Player) float32 {
	return getFloat(p.Entity, "localdata.m_Local.m_flStepSize")
}

func HideHUD(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_iHideHUD")
}

func DuckJumpTimeMsecs(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_nDuckJumpTimeMsecs")
}

func DuckTimeMsecs(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_nDuckTimeMsecs")
}

func JumpTimeMsecs(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_nJumpTimeMsecs")
}

func Skybox3dArea(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_skybox3d.area")
}

func HDRColorScale(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_skybox3d.fog.HDRColorScale")
}

func Blend(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_skybox3d.fog.blend")
}

func ColorPrimary(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_skybox3d.fog.colorPrimary")
}

func ColorSecondary(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_skybox3d.fog.colorSecondary")
}

func DirPrimary(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_skybox3d.fog.dirPrimary")
}

func FogEnable(p *common.Player) bool {
	return getBool(p.Entity, "localdata.m_Local.m_skybox3d.fog.enable")
}

func FogEnd(p *common.Player) float32 {
	return getFloat(p.Entity, "localdata.m_Local.m_skybox3d.fog.end")
}

func FogMaxDensity(p *common.Player) float32 {
	return getFloat(p.Entity, "localdata.m_Local.m_skybox3d.fog.maxdensity")
}

func FogStart(p *common.Player) float32 {
	return getFloat(p.Entity, "localdata.m_Local.m_skybox3d.fog.start")
}

func Skybox3DOrigin(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_skybox3d.origin")
}

func Skybox3DScale(p *common.Player) int {
	return getInt(p.Entity, "localdata.m_Local.m_skybox3d.scale")
}

func AimPunchAngleVel(p *common.Player) r3.Vector {
	return getVector(p.Entity, "localdata.m_Local.m_aimPunchAngleVel")
}

func ScaleType(p *common.Player) int {
	return getInt(p.Entity, "m_ScaleType")
}

func ClientSideAnimation(p *common.Player) bool {
	return getBool(p.Entity, "m_bClientSideAnimation")
}

func ClientSideFrameReset(p *common.Player) bool {
	return getBool(p.Entity, "m_bClientSideFrameReset")
}

func ClientSideRagdoll(p *common.Player) bool {
	return getBool(p.Entity, "m_bClientSideRagdoll")
}

func SuppressAnimSounds(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bSuppressAnimSounds")

	return val.BoolVal()
}
func Frozen(p *common.Player) int {
	return getInt(p.Entity, "m_flFrozen")
}

func ModelScale(p *common.Player) float32 {
	return getFloat(p.Entity, "m_flModelScale")
}

func Body(p *common.Player) int {
	return getInt(p.Entity, "m_nBody")
}

func HighlightColorB(p *common.Player) int {
	return getInt(p.Entity, "m_nHighlightColorB")
}

func HighlightColorG(p *common.Player) int {
	return getInt(p.Entity, "m_nHighlightColorG")
}

func HighlightColorR(p *common.Player) int {
	return getInt(p.Entity, "m_nHighlightColorR")
}

func NextPrimaryAttack(p *common.Player) float32 {

	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalActiveWeaponData.m_flNextPrimaryAttack")

	return val.FloatVal
}

func NextSecondaryAttack(p *common.Player) float32 {

	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalActiveWeaponData.m_flNextSecondaryAttack")

	return val.FloatVal
}

func TimeWeaponIdle(p *common.Player) float32 {

	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalActiveWeaponData.m_flTimeWeaponIdle")

	return val.FloatVal
}
func WeaponNextThinkTick(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalActiveWeaponData.m_nNextThinkTick")

	return val.IntVal
}
func PrimaryAmmoType(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalWeaponData.m_iPrimaryAmmoType")

	return val.IntVal
}
func SecondaryAmmoType(p *common.Player) int {

	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LLocalWeaponData.m_iSecondaryAmmoType")

	return val.IntVal
}
func WeaponModule(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalWeaponData.m_iWeaponModule")

	return val.IntVal
}
func WeaponOrigin(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalWeaponData.m_iWeaponOrigin")

	return val.IntVal
}
func ViewModelIndex(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("LocalWeaponData.m_nViewModelIndex")

	return val.IntVal
}

func CollisionGroup(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_CollisionGroup")

	return val.IntVal
}

func AngRotation(p *common.Player) r3.Vector {
	if p.Entity == nil {
		return r3.Vector{}
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_angRotation")

	return val.VectorVal
}

func AlternateSorting(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bAlternateSorting")

	return val.BoolVal()
}

func AnimatedEveryTick(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bAnimatedEveryTick")

	return val.BoolVal()
}

func EligibleForScreenHighlight(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bEligibleForScreenHighlight")

	return val.BoolVal()
}

func IsAutoaimTarget(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bIsAutoaimTarget")

	return val.BoolVal()
}

func SimulatedEveryTick(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bSimulatedEveryTick")

	return val.BoolVal()
}

func CellX(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_cellX")

	return val.IntVal
}

func CellY(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_cellY")

	return val.IntVal
}

func CellZ(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_cellZ")

	return val.IntVal
}

func Cellbits(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_cellbits")

	return val.IntVal
}

func ClrRender(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_clrRender")

	return val.IntVal
}

func Effects(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_fEffects")

	return val.IntVal
}

func FadeMinDist(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_fadeMinDist")

	return val.FloatVal
}

func FadeMaxDist(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_fadeMaxDist")

	return val.FloatVal
}

func Elasticity(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flElasticity")

	return val.FloatVal
}

func FadeScale(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flFadeScale")

	return val.FloatVal
}

func MaxFallVelocity(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flMaxFallVelocity")

	return val.FloatVal
}

func ShadowCastDistance(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flShadowCastDistance")

	return val.FloatVal
}

func UseLookAtAngle(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flUseLookAtAngle")

	return val.FloatVal
}

func EffectEntity(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_hEffectEntity")

	return val.IntVal
}

func OwnerEntity(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_hOwnerEntity")

	return val.IntVal
}

func PendingTeamNum(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iPendingTeamNum")

	return val.IntVal
}

func ParentAttachment(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iParentAttachment")

	return val.IntVal
}

func TeamNum(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iTeamNum")

	return val.IntVal
}

func TextureFrameIndex(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iTextureFrameIndex")

	return val.IntVal
}

func MaxCPULevel(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nMaxCPULevel")

	return val.IntVal
}

func MaxGPULevel(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nMaxGPULevel")

	return val.IntVal
}

func MinCPULevel(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nMinCPULevel")

	return val.IntVal
}

func MinGPULevel(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nMinGPULevel")

	return val.IntVal
}

func ModelIndex(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nModelIndex")

	return val.IntVal
}

func RenderFX(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nRenderFX")

	return val.IntVal
}

func RenderMode(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nRenderMode")

	return val.IntVal
}

func ZoomLevel(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return -1
	}
	return int(p.ActiveWeapon().ZoomLevel())
}

func AmmoReserve(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return -1
	}
	return p.ActiveWeapon().AmmoReserve()
}

func NumEmptyAttacks(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_iNumEmptyAttacks")

	return val.IntVal
}

func Clip1(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_iClip1")

	return val.IntVal
}

func BurstShotsRemaining(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_iBurstShotsRemaining")

	return val.IntVal
}

func RecoilIndex(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_flRecoilIndex")

	return val.FloatVal
}

func PlaybackRate(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_flPlaybackRate")

	return val.FloatVal
}

func DoneSwitchingSilencer(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_flDoneSwitchingSilencer")

	return val.BoolVal()
}

func LastShotTime(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_fLastShotTime")

	return val.FloatVal
}

func AccuracyPenalty(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_fAccuracyPenalty")

	return val.FloatVal
}

func BurstMode(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_bBurstMode")

	return val.BoolVal()
}

func WeaponMode(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_weaponMode")

	return val.BoolVal()
}

func WeaponState(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_iState")

	return val.IntVal
}

func MuzzleFlashParity(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_nMuzzleFlashParity")

	return val.IntVal
}

func ReloadState(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_reloadState")

	return val.BoolVal()
}

func M3ReloadState(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_reloadState")

	return val.BoolVal()
}

func NOVAReloadState(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_reloadState")

	return val.BoolVal()
}

func SawedoffReloadState(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_reloadState")

	return val.BoolVal()
}

func XM1014ReloadState(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_reloadState")

	return val.BoolVal()
}

func ReloadVisuallyComplete(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_bReloadVisuallyComplete")

	return val.BoolVal()
}

func SilencerOn(p *common.Player) bool {
	if p.ActiveWeapon() == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_bSilencerOn")

	return val.BoolVal()
}

func IronSightMode(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_iIronSightMode")

	return val.IntVal
}

func PrimaryReserveAmmoCount(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_iPrimaryReserveAmmoCount")

	return val.IntVal
}

func ItemDefinitionIndex(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_AttributeManager.m_Item.m_iItemDefinitionIndex")

	return val.IntVal
}

func WeaponServerCycle(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("serveranimdata.m_flCycle")

	return val.FloatVal
}

func Inaccuracy(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 1
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_fInaccuracy")

	return float32(val.IntVal)
}

func Spread(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_fSpread")

	return val.FloatVal
}

func WeaponVecOrigin(p *common.Player) r3.Vector {
	if p.ActiveWeapon() == nil {
		return r3.Vector{}
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_vecOrigin")

	return val.VectorVal
}

func WeaponOwner(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_hOwner")

	return val.IntVal
}

func WeaponWorldModel(p *common.Player) int {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_hWeaponWorldModel")

	return val.IntVal
}

func CanControlObservedBot(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bCanControlObservedBot")

	return val.BoolVal()
}

func CanMoveDuringFreezePeriod(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bCanMoveDuringFreezePeriod")

	return val.BoolVal()
}

func GunGameImmunity(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bGunGameImmunity")

	return val.BoolVal()
}

func HasControlledBotThisRound(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bHasControlledBotThisRound")

	return val.BoolVal()
}

func HasHeavyArmor(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bHasHeavyArmor")

	return val.BoolVal()
}

func HasNightVision(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bHasNightVision")

	return val.BoolVal()
}

func HideTargetID(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bHideTargetID")

	return val.BoolVal()
}

func Hud_MiniScoreHidden(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bHud_MiniScoreHidden")

	return val.BoolVal()
}

func Hud_RadarHidden(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bHud_RadarHidden")

	return val.BoolVal()
}

func InHostageRescueZone(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bInHostageRescueZone")

	return val.BoolVal()
}

func InNoDefuseArea(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bInNoDefuseArea")

	return val.BoolVal()
}

func IsGrabbingHostage(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bIsGrabbingHostage")

	return val.BoolVal()
}

func IsPlayerGhost(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bIsPlayerGhost")

	return val.BoolVal()
}

func IsRescuing(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bIsRescuing")

	return val.BoolVal()
}

func IsRespawningForDMBonus(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bIsRespawningForDMBonus")

	return val.BoolVal()
}

func IsSpawnRappelling(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bIsSpawnRappelling")

	return val.BoolVal()
}

func KilledByTaser(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bKilledByTaser")

	return val.BoolVal()
}

func MadeFinalGunGameProgressiveKill(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bMadeFinalGunGameProgressiveKill")

	return val.BoolVal()
}

func NightVisionOn(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bNightVisionOn")

	return val.BoolVal()
}

func CycleLatch(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_cycleLatch")

	return val.IntVal
}

func ImmuneToGunGameDamageTime(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_fImmuneToGunGameDamageTime")

	return val.IntVal
}

func AutoMoveTargetTime(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flAutoMoveTargetTime")

	return val.FloatVal
}

func DetectedByEnemySensorTime(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flDetectedByEnemySensorTime")

	return val.FloatVal
}

func GuardianTooFarDistFrac(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flGuardianTooFarDistFrac")

	return val.FloatVal
}

func HealthShotBoostExpirationTime(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flHealthShotBoostExpirationTime")

	return val.FloatVal
}

func LastExoJumpTime(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flLastExoJumpTime")

	return val.FloatVal
}

func ProgressBarStartTime(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flProgressBarStartTime")

	return val.FloatVal
}

func ThirdpersonRecoil(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_flThirdpersonRecoil")

	return val.FloatVal
}

func CarriedHostage(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_hCarriedHostage")

	return val.IntVal
}

func MoveType(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("movetype")

	return val.IntVal
}

func MoveCollide(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("movecollide")

	return val.IntVal
}

func CarriedHostageProp(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_hCarriedHostageProp")

	return val.IntVal
}

func Ragdoll(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_hRagdoll")

	return val.IntVal
}

func SurvivalAssassinationTarget(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_hSurvivalAssassinationTarget")

	return val.IntVal
}

func Account(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iAccount")

	return val.IntVal
}

func AddonBits(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iAddonBits")

	return val.IntVal
}

func BlockingUseActionInProgress(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iBlockingUseActionInProgress")

	return val.IntVal
}

func Class(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iClass")

	return val.IntVal
}

func ControlledBotEntIndex(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iControlledBotEntIndex")

	return val.IntVal
}

func GunGameProgressiveWeaponIndex(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iGunGameProgressiveWeaponIndex")

	return val.IntVal
}

func NumGunGameKillsWithCurrentWeapon(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iNumGunGameKillsWithCurrentWeapon")

	return val.IntVal
}

func NumGunGameTRKillPoints(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iNumGunGameTRKillPoints")

	return val.IntVal
}

func NumRoundKills(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iNumRoundKills")

	return val.IntVal
}

func NumRoundKillsHeadshots(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iNumRoundKillsHeadshots")

	return val.IntVal
}

func PrimaryAddon(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iPrimaryAddon")

	return val.IntVal
}

func ProgressBarDuration(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iProgressBarDuration")

	return val.IntVal
}

func SecondaryAddon(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iSecondaryAddon")

	return val.IntVal
}

func StartAccount(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iStartAccount")

	return val.IntVal
}

func ThrowGrenadeCounter(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_iThrowGrenadeCounter")

	return val.IntVal
}

func IsCurrentGunGameLeader(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_isCurrentGunGameLeader")

	return val.BoolVal()
}

func IsCurrentGunGameTeamLeader(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_isCurrentGunGameTeamLeader")

	return val.BoolVal()
}
func DeathCamMusic(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nDeathCamMusic")

	return val.IntVal
}

func HeavyAssaultSuitCooldownRemaining(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nHeavyAssaultSuitCooldownRemaining")

	return val.IntVal
}

func IsAutoMounting(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nIsAutoMounting")

	return val.IntVal
}

func LastKillerIndex(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nLastKillerIndex")

	return val.IntVal
}

func SurvivalTeam(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_nSurvivalTeam")

	return val.IntVal
}

func ArmsModel(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_szArmsModel")

	return val.IntVal
}

func TotalRoundDamageDealt(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_unTotalRoundDamageDealt")

	return val.IntVal
}

func ConstraintPastRadius(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_bConstraintPastRadius")

	return val.BoolVal()
}

func ForceTeam(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_fForceTeam")

	return val.FloatVal
}

func OnTarget(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_fOnTarget")

	return val.IntVal
}

func ConstraintRadius(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_flConstraintRadius")

	return val.IntVal
}

func ConstraintSpeedFactor(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_flConstraintSpeedFactor")

	return val.IntVal
}

func ConstraintWidth(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_flConstraintWidth")

	return val.IntVal
}

func Friction(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_flFriction")

	return val.IntVal
}

/*
func LaggedMovementValue(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_flLaggedMovementValue")

	return val.FloatVal
}*/

func NextDecalTime(p *common.Player) float32 {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("localdata.m_flNextDecalTime")

	return val.FloatVal
}

func RetakesMVPBoostExtraUtility(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_RetakesMVPBoostExtraUtility")

	return val.IntVal
}

func RetakesHasDefuseKit(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_bRetakesHasDefuseKit")

	return val.BoolVal()
}

func RetakesMVPLastRound(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_bRetakesMVPLastRound")

	return val.BoolVal()
}

func Direction(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_iDirection")

	return val.IntVal
}

func RetakesMVPBoostItem(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_iRetakesMVPBoostItem")

	return val.IntVal
}

func RetakesOffering(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_iRetakesOffering")

	return val.IntVal
}

func RetakesOfferingCard(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_iRetakesOfferingCard")

	return val.IntVal
}

func NumFastDucks(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_nNumFastDucks")

	return val.IntVal
}

func QuestProgressReason(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("cslocaldata.m_nQuestProgressReason")

	return val.IntVal
}

func ViewTarget(p *common.Player) int {
	if p.Entity == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_viewtarget")

	return val.IntVal
}

func ActiveCameraMan(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bActiveCameraMan")

	return val.BoolVal()
}

func CameraManOverview(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bCameraManOverview")

	return val.BoolVal()
}

func CameraManScoreBoard(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bCameraManScoreBoard")

	return val.BoolVal()
}

func CameraManXRay(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bCameraManXRay")

	return val.BoolVal()
}

func ShouldDrawPlayerWhileUsingViewEntity(p *common.Player) bool {
	if p.Entity == nil {
		return false
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.Entity.PropertyValue("m_bShouldDrawPlayerWhileUsingViewEntity")

	return val.BoolVal()
}

func PostponeFireReadyTime(p *common.Player) float32 {
	if p.ActiveWeapon() == nil {
		return 0
	}

	// if the property doesn't exist we return 0 by default
	val, _ := p.ActiveWeapon().Entity.PropertyValue("m_flPostponeFireReadyTime")

	return val.FloatVal
}

func FlashDurationTimeFull(p *common.Player) time.Duration {
	return time.Duration(float32(time.Second) * p.FlashDuration)
}
