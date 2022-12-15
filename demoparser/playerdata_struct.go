package demoparser

import (
	h "github.com/David-Durst/head-position-model/pkg/headpositionmodel"
	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	common "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	play "github.com/mwaurawakati/csgo/playerdata"
	"math"
	"time"
)

// PlayerData contains player data with ActiveWeapon information that the player wields
type PlayerData struct {
	SteamID64 uint64
	//Name                          string
	FlashDuration float32
	Velocity0     float64
	Velocity1     float64
	Velocity2     float64
	//ActiveWeapon                  string
	LastAlivePosition0            float64
	LastAlivePosition1            float64
	LastAlivePosition2            float64
	ViewDirectionX                float32
	ViewDirectionY                float32
	IsDucking                     bool
	PlayerHealth                  int
	Firing                        bool
	AmmoInMagazine                int //
	Ping                          int
	IsScoped                      bool
	AmmoReserve                   int
	ZoomLevel                     common.ZoomLevel
	AmmoType                      int
	RecoilIndex                   float32
	BurstMode                     bool
	ReloadVisuallyComplete        bool
	SilencerOn                    bool
	AccuracyPenalty               float32
	LastShotTime                  float32
	DoneSwitchingSilencer         bool
	PostponeFireReadyTime         float32
	IronSightMode                 int
	WeaponMode                    bool
	AmmoLeft                      [32]int
	FlashTick                     int
	IsBot                         bool
	IsControllingBot              bool
	IsConnected                   bool
	IsDefusing                    bool
	IsPlanting                    bool
	IsReloading                   bool
	IsUnknown                     bool
	IsAlive                       bool
	IsBlinded                     bool
	IsAirborne                    bool
	FlashDurationTime             time.Duration
	FlashDurationTimeFull         time.Duration
	IsWalking                     bool
	IsDuckingInProgress           bool
	IsUnDuckingInProgress         bool
	IsStanding                    bool
	HasHelmet                     bool
	Armor                         int
	PositionX                     float64
	PositionY                     float64
	PositionZ                     float64
	PositionEyesX                 float64
	PositionEyesY                 float64
	PositionEyesZ                 float64
	HeadPositionX                 float64
	HeadPositionY                 float64
	HeadPositionZ                 float64
	PlayerFlags                   common.PlayerFlags
	OnGround                      bool
	DuckingKeyPressed             bool
	HasHeavyArmor                 bool
	HasMovedSinceSpawn            bool
	IsGrabbingHostage             bool
	IsHoldingLookAtWeapon         bool
	IsLookingAtWeapon             bool
	IsRescuing                    bool
	ResumeZone                    bool
	Strafing                      bool
	WaitForNoAttack               bool
	MolotovDamageTime             float64
	FlashMaxAlpha                 float64
	LowerYawBodyTarget            float64
	ThirdpersonRecoil             float32
	BlockingUseActionInProgress   int
	MoveState                     int
	PlayerState                   int
	ViewPunchAngleX               float64
	ViewPunchAngleY               float64
	ViewPunchAngleZ               float64
	AimPunchAngleX                float64
	AimPunchAngleY                float64
	AimPunchAngleZ                float64
	AimPunchAngleVelX             float64
	AimPunchAngleVelY             float64
	AimPunchAngleVelZ             float64
	InDuckJump                    bool
	FOVRate                       float64
	FallVelocity                  float64
	StepSize                      float32
	WearingSuit                   bool
	Poisoned                      bool
	DuckOverride                  float64
	Stamina                       float64
	VelocityModifier              float64
	Direction                     int
	ShotsFired                    int
	TimeOfLastInjury              float64
	RelativeDirectionOfLastInjury int
	VecLadderNormalX              float64
	VecLadderNormalY              float64
	VecLadderNormalZ              float64
	PhysicsCollisionState         int
	WaterLevel                    int
	LifeState                     int
	LadderSurfaceProps            int
	FOVStart                      int
	FOV                           int
	DefaultFOV                    int
	MaxSpeed                      float64
	FOVTime                       float64
	DuckSpeed                     float64
	DuckAmount                    float64
	PhysicsFlags                  int
	VecBaseVelocityX              float64
	VecBaseVelocityY              float64
	VecBaseVelocityZ              float64
	LaggedMovementValue           float64
	M3ReloadState                 bool
	NOVAReloadState               bool
	SawedoffReloadState           bool
	XM1014ReloadState             bool
	MoveType                      int
	MoveCollide                   int
	UseLookAtAngle                float32
	SimulationTime                int
	ShadowCastDistance            float32
	MaxFallVelocity               float32
	LastMadeNoiseTime             float64
	NextPrimaryAttack             float32
	NextSecondaryAttack           float32
	TimeWeaponIdle                float32
	WeaponState                   int
	BurstShotsRemaining           int
	ViewTarget                    int
	ForceBone                     int
	Yaw                           float32
	Pitch                         float32
	Position                      r3.Vector
	Footstep                      bool
	FallDamage                    float32
	WeaponFireOnEmpty             bool
	Tick                          int
	Frame                         int
}

// ExtractPlayerData extracts player data
func ExtractPlayerData(frame int, player *common.Player, fireFrames map[FireFrameKey]bool, footstepFrames map[FootstepKey]bool, currentTime time.Duration, ingameTick int, weaponFOEFrame map[WeaponFOEKey]bool, fallDamageFrame map[FallDamageKey]float32) PlayerData {

	fixedPitch := float32(math.Mod(float64(player.ViewDirectionY())+90, 180))
	ViewDirectionX := player.ViewDirectionX()
	ViewDirectionY := player.ViewDirectionY()
	pos := player.LastAlivePosition
	id := 0
	if player == nil {
		id = 0
	} else {
		id = player.UserID
	}

	viewangle := r2.Point{X: float64(ViewDirectionX), Y: float64(ViewDirectionY)}
	headposition := h.ModelHeadPosition(pos, viewangle, play.DuckAmount(player))

	activeweapon := player.ActiveWeapon()
	var ammoleft int = -2
	var weapon int = -2
	//var a string
	var reserve int
	var zoomlevel common.ZoomLevel
	//var falldamage float32

	// ActiveWeapon might return nil
	if activeweapon == nil {
		ammoleft = -2
		weapon = -2
		//a = ""
		reserve = -2
		zoomlevel = 0

	} else {
		ammoleft = activeweapon.AmmoInMagazine()
		weapon = int(activeweapon.Type)
		//a = activeweapon.String()
		reserve = activeweapon.AmmoReserve()
		zoomlevel = activeweapon.ZoomLevel()
	}

	return PlayerData{
		player.SteamID64,
		//player.Name,
		player.FlashDuration,
		player.Velocity().X,
		player.Velocity().Y,
		player.Velocity().Z,
		//a,
		player.LastAlivePosition.X,
		player.LastAlivePosition.Y,
		player.LastAlivePosition.Z,
		player.ViewDirectionX(),
		player.ViewDirectionY(),
		player.IsDucking(),
		player.Health(),
		fireFrames[FireFrameKey{id, frame}],
		ammoleft,
		player.Ping(),
		player.IsScoped(),
		reserve,
		zoomlevel,
		weapon,
		play.RecoilIndex(player),
		play.BurstMode(player),
		play.ReloadVisuallyComplete(player),
		play.SilencerOn(player),
		play.AccuracyPenalty(player),
		TimeDifference(currentTime, play.LastShotTime(player)),
		play.DoneSwitchingSilencer(player),
		TimeDifference(currentTime, play.PostponeFireReadyTime(player)),
		play.IronSightMode(player),
		play.WeaponMode(player),
		player.AmmoLeft,
		ingameTick - player.FlashTick,
		player.IsBot,
		player.IsControllingBot(),
		player.IsConnected,
		player.IsDefusing,
		player.IsPlanting,
		player.IsReloading,
		player.IsUnknown,
		player.IsAlive(),
		player.IsBlinded(),
		player.IsAirborne(),
		player.FlashDurationTime(),
		play.FlashDurationTimeFull(player),
		player.IsWalking(),
		player.IsDuckingInProgress(),
		player.IsUnDuckingInProgress(),
		player.IsStanding(),
		player.HasHelmet(),
		player.Armor(),
		player.Position().X,
		player.Position().Y,
		player.Position().Z,
		player.PositionEyes().X,
		player.PositionEyes().Y,
		player.PositionEyes().Z,
		headposition.X,
		headposition.Y,
		headposition.Z,
		player.Flags(),
		player.Flags().OnGround(),
		player.Flags().DuckingKeyPressed(),
		play.HasHeavyArmor(player),
		play.HasMovedSinceSpawn(player),
		play.IsGrabbingHostage(player),
		play.IsHoldingLookAtWeapon(player),
		play.IsLookingAtWeapon(player),
		play.IsRescuing(player),
		play.ResumeZoom(player),
		play.Strafing(player),
		play.WaitForNoAttack(player),
		float64(TimeDifference(currentTime, float32(play.MolotovDamageTime(player)))),
		play.FlashMaxAlpha(player),
		play.LowerBodyYawTarget(player),
		play.ThirdpersonRecoil(player),
		play.BlockingUseActionInProgress(player),
		play.MoveState(player),
		play.PlayerState(player),
		play.ViewPunchAngle(player).X,
		play.ViewPunchAngle(player).Y,
		play.ViewPunchAngle(player).Z,
		play.AimPunchAngle(player).X,
		play.AimPunchAngle(player).Y,
		play.AimPunchAngle(player).Z,
		play.AimPunchAngleVel(player).X,
		play.AimPunchAngleVel(player).Y,
		play.AimPunchAngleVel(player).Z,
		play.InDuckJump(player),
		play.FOVRate(player),
		play.FallVelocity(player),
		play.StepSize(player),
		play.WearingSuit(player),
		play.Poisoned(player),
		play.DuckOverride(player),
		play.Stamina(player),
		play.VelocityModifier(player),
		play.Direction(player),
		play.ShotsFired(player),
		float64(TimeDifference(currentTime, float32(play.TimeOfLastInjury(player)))),
		play.RelativeDirectionOfLastInjury(player),
		play.VecLadderNormal(player).X,
		play.VecLadderNormal(player).Y,
		play.VecLadderNormal(player).Z,
		play.PhysicsCollisionState(player),
		play.WaterLevel(player),
		play.LifeState(player),
		play.LadderSurfaceProps(player),
		play.FOVStart(player),
		play.FOV(player),
		play.DefaultFOV(player),
		play.MaxSpeed(player),
		float64(TimeDifference(currentTime, float32(play.FOVTime(player)))),
		play.DuckSpeed(player),
		play.DuckAmount(player),
		play.PhysicsFlags(player),
		play.VecBaseVelocity(player).X,
		play.VecBaseVelocity(player).Y,
		play.VecBaseVelocity(player).Z,
		play.LaggedMovementValue(player),
		play.M3ReloadState(player),
		play.NOVAReloadState(player),
		play.SawedoffReloadState(player),
		play.XM1014ReloadState(player),
		play.MoveType(player),
		play.MoveCollide(player),
		play.UseLookAtAngle(player),
		play.SimulationTime(player),
		play.ShadowCastDistance(player),
		play.MaxFallVelocity(player),
		float64(TimeDifference(currentTime, float32(play.LastMadeNoiseTime(player)))),
		TimeDifference(currentTime, play.NextPrimaryAttack(player)),
		TimeDifference(currentTime, play.NextSecondaryAttack(player)),
		TimeDifference(currentTime, play.TimeWeaponIdle(player)),
		play.WeaponState(player),
		play.BurstShotsRemaining(player),
		play.ViewTarget(player),
		play.ForceBone(player),
		player.ViewDirectionX(),
		fixedPitch,
		player.LastAlivePosition,
		footstepFrames[FootstepKey{id, frame}],
		fallDamageFrame[FallDamageKey{int32(id), frame}],
		weaponFOEFrame[WeaponFOEKey{id, frame}],
		ingameTick,
		frame,
	}
}

// Time difference calculates time differences
func TimeDifference(currenttime time.Duration, othertime float32) float32 {
	return float32(currenttime) - (float32(time.Second) * othertime)
}
