package demoparser

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/golang/geo/r3"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	msg "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
	"github.com/pkg/errors"
)

// Defines amount of frames to collect around attacks
const (
	//samplesPerSecond    = 32
	secondsBeforeAttack = 6
	secondsAfterAttack  = 2
	secondsPerAttack    = secondsBeforeAttack + secondsAfterAttack
)

// MissingTicksAndFrames parses through the demo file checking missing ticks and frames
func MissingTicksAndFrames(source, name string) ([]int, []int, error) {
	var frames []int
	var ticklist []int
	defer RecoverFromPanic()
	f, err := os.Open(source + name)
	defer f.Close()
	checkError(err)
	cfg := dem.DefaultParserConfig
	cfg.IgnoreErrBombsiteIndexNotFound = true
	p := dem.NewParserWithConfig(f, cfg)
	h, _ := p.ParseHeader()
	prevframe := -1
	tr := p.TickRate()
	fr := h.FrameRate()
	dv := math.Round(tr / (math.Round(fr)))
	gs := p.GameState()
	prevtick := -dv
	if int(dv) == 1 || int(dv) == 2 || int(dv) == 4 || int(dv) == 8 || int(dv) == 16 || int(dv) == 32 {

		for ok := true; ok; ok, err = p.ParseNextFrame() {
			if err != nil {
				log.Println("This demos :", name, "Expereinces EOF panic and therefore missing ticks and frames can not be handled")
				return nil, nil, err
			}

			if prevframe+1 != p.CurrentFrame() {
				frames = append(frames, -1)
				frames = append(frames, p.CurrentFrame())
				prevframe++
			} else {
				prevframe = p.CurrentFrame()
				frames = append(frames, p.CurrentFrame())
			}

			igt := gs.IngameTick()
			if float64(igt) < 0 {
				ticklist = append(ticklist, -1)
			} else if prevtick+dv != float64(igt) {
				prevtick = float64(igt)
				ticklist = append(ticklist, -1)
				ticklist = append(ticklist, igt)
			} else {
				prevtick = float64(igt)
				ticklist = append(ticklist, igt)
			}
		}
	} else {
		e := errors.New("Invalid TickRate")
		return nil, nil, errors.Wrap(e, FloatToString(dv))
	}
	return frames, ticklist, nil
}
func FrameRateCheck(source string, name string) {
	defer RecoverFromPanic()
	f, err := os.Open(source + name)
	checkError(err)
	cfg := dem.DefaultParserConfig
	cfg.IgnoreErrBombsiteIndexNotFound = true
	p := dem.NewParserWithConfig(f, cfg)
	h, _ := p.ParseHeader()
	FrameRate := h.FrameRate()
	//tr := p.TickRate()
	//dv := math.Round(tr / (math.Round(fr)))
	f.Close()
	frames, ticklist, er := MissingTicksAndFrames(source, name)
	if er != nil {
		fmt.Println(er)
	}

	if FrameRate <= 32 {
		parseDemo(source, name, 32, frames, ticklist)
	} else if FrameRate > 32 && FrameRate <= 64 {
		parseDemo(source, name, 32, frames, ticklist)
		parseDemo(source, name, 64, frames, ticklist)
	} else {
		parseDemo(source, name, 32, frames, ticklist)
		parseDemo(source, name, 64, frames, ticklist)
		parseDemo(source, name, 128, frames, ticklist)
	}
}

func parseDemo(source string, name string, framerate float64, frames []int, ticklist []int) {
	samplesPerAttack := int(framerate * secondsPerAttack)
	//defer wg.Done()
	//defer RecoverFromPanic()
	fmt.Println("Parsing demo", name)

	modelData := []AttackData{}
	// Times when a player is attacked
	var attackTimes = []AttackTime{}
	// Times when a player is killed
	var killTimes = map[int]Kill{}
	// Marks if a player is firing at a given frame.
	var fireFrames = map[FireFrameKey]bool{}
	//Marks of a player is footstepping at a given frame.
	var footstepFrames = map[FootstepKey]bool{}
	//Marks a fall damage event
	var fallDamageFrame = map[FallDamageKey]float32{}
	// Marks if a player is firing empty weapon at a given frame.
	var weaponFOEFrame = map[WeaponFOEKey]bool{}
	// Marks frames surrounding Killed that should not be gathered
	var isKillEvent = map[int]bool{}
	//Game events
	var ged map[int32]*msg.CSVCMsg_GameEventListDescriptorT
	// Stores the PlayerData for each player for each marked framed
	var markedFrameData = map[int]map[int]PlayerData{}
	// Stores the KillData for each marked framed
	var markedKillData = map[int]Kill{}

	var killframe []int
	var attackframe []int

	//Parsing server information
	f, err := os.Open(source + name)
	checkError(err)
	res := ParseServerInfo(f)
	f.Close()
	cfg := dem.DefaultParserConfig
	cfg.IgnoreErrBombsiteIndexNotFound = true
	FrameRate := framerate

	var framesBeforeAttack int
	var framesAfterAttack int
	if math.Abs(FrameRate-32.0) < 1 {
		framesBeforeAttack = secondsBeforeAttack * 32
		framesAfterAttack = secondsAfterAttack * 32
	} else if math.Abs(FrameRate-64.0) < 4 {
		framesBeforeAttack = secondsBeforeAttack * 64
		framesAfterAttack = secondsAfterAttack * 64
	} else if math.Abs(FrameRate-128) < 4 {
		framesBeforeAttack = secondsBeforeAttack * 128
		framesAfterAttack = secondsAfterAttack * 128
	} else {
		framesBeforeAttack = int(secondsBeforeAttack * FrameRate)
		framesAfterAttack = int(secondsAfterAttack * FrameRate)
	}

	framesPerAttack := framesBeforeAttack + framesAfterAttack
	framesPerSample := int(framesPerAttack / samplesPerAttack)
	var start int = 0
	var end int = 0
	var frame int = 0
	var attackFrame int

	//PlayerHurt Event
	attackCount := 0
	validAttacks := 0
	f, err = os.Open(source + name)
	p := dem.NewParserWithConfig(f, cfg)

	p.RegisterEventHandler(func(e events.PlayerHurt) {
		attackCount++

		attackFrame = p.CurrentFrame()
		attackframe = append(attackframe, attackFrame)

		s := p.GameState()
		var w string
		if e.Weapon == nil {
			w = ""
		} else {
			w = e.Weapon.String()
		}
		new := AttackTime{e.Attacker, e.Player, p.CurrentFrame(), e.HitGroup, s.GamePhase(), s.IsWarmupPeriod(), s.IsMatchStarted(), e.ArmorDamage, e.HealthDamageTaken, w}
		start = attackFrame - framesBeforeAttack
		end = attackFrame + framesAfterAttack
		attid := 0
		vicid := 0
		atteam := 0
		victeam := 0
		if e.Attacker == nil {
			attid = 0
			atteam = 0
		} else {
			attid = e.Attacker.UserID
			atteam = int(e.Attacker.Team)
		}
		if e.Player == nil {
			vicid = 0
			victeam = 0
		} else {
			vicid = e.Player.UserID
			victeam = int(e.Player.Team)
		}
		wip := 0
		if e.Weapon == nil {
			wip = 0
		} else {
			wip = int(e.Weapon.Type)
		}

		wep := Contains(wip)

		if !wep {

		} else if vicid == attid {

		} else if atteam == victeam {

		} else {
			attackTimes = append(attackTimes, new)
			validAttacks++
			//prevind := 0
			for frame = start - framesPerSample; frame < end; frame++ {
				fr := frame
				gs := p.GameState()
				tic := gs.IngameTick()
				if frames != nil {
					if Contains1(frames, frame) {
						fr = frame
						ind := indexOf(frame, frames)
						tic = ticklist[ind]

					} else {
						fr = -1
						tic = gs.IngameTick()

					}

				} else {
					fr = frame
					tic = gs.IngameTick()
				}

				var players = map[int]PlayerData{}

				currentTime := p.CurrentTime()
				for _, player := range gs.Participants().Playing() {
					players[player.UserID] = ExtractPlayerData(fr, player, fireFrames, footstepFrames, currentTime, tic, weaponFOEFrame, fallDamageFrame)
				}
				markedFrameData[frame] = players
			}
		}
	})
	p.ParseToEnd()
	f.Close()
	f, err = os.Open(source + name)
	p = dem.NewParserWithConfig(f, cfg)
	//Kill Event
	KillCount := 0
	var killFrame int
	start = 0
	end = 0
	p.RegisterEventHandler(func(e events.Kill) {
		KillCount++
		killFrame = p.CurrentFrame()
		killframe = append(killframe, killFrame)
		start = killFrame - framesBeforeAttack
		end = killFrame + framesAfterAttack
		for frame = start; frame < end; frame++ {
			isKillEvent[frame] = true
		}
		isKillEvent[start-framesPerSample] = true // For first sample delta angles

		new := Kill{e.Weapon,
			e.Victim,
			e.Killer,
			e.Assister,
			e.PenetratedObjects,
			e.IsHeadshot,
			e.AssistedFlash,
			e.AttackerBlind,
			e.NoScope,
			e.ThroughSmoke,
			e.Distance,
		}
		killTimes[killFrame] = new
	})
	//Weapon Fire Event
	// Track frames where a player fires a weapon
	p.RegisterEventHandler(func(e events.WeaponFire) {
		frame := p.CurrentFrame()
		// Include previous frames so that shot is not lost after sampling
		for i := 0; i < framesPerSample; i++ {
			sid := 0
			if e.Shooter == nil {
				sid = 0
			} else {
				sid = e.Shooter.UserID
			}

			fireFrames[FireFrameKey{sid, frame - i}] = true
		}
	})
	p.RegisterEventHandler(func(e events.Footstep) {
		frame := p.CurrentFrame()
		for i := 0; i < framesPerSample; i++ {
			pid := 0
			if e.Player == nil {
				pid = 0
			} else {
				pid = e.Player.UserID
			}

			footstepFrames[FootstepKey{pid, frame - i}] = true
		}
	})

	p.RegisterNetMessageHandler(func(d *msg.CSVCMsg_GameEventList) {
		ged = handleGameEventList(d)
	})
	p.RegisterNetMessageHandler(func(ge *msg.CSVCMsg_GameEvent) {
		desc := ged[ge.GetEventid()]
		frame := p.CurrentFrame()
		n, d := debugGameEvent(desc, ge)
		if n == "player_falldamage" {
			for i := 0; i < framesPerSample; i++ {
				dam := d["damage"].(float32)
				id := d["userid"].(int32)
				fallDamageFrame[FallDamageKey{id, frame - i}] = dam
			}
		}
		if n == "weapon_fire_on_empty" {
			for i := 0; i < framesPerSample; i++ {
				weaponFOEFrame[WeaponFOEKey{d["userid"], frame}] = true
			}
		}

	})

	p.ParseToEnd()

	fmt.Printf("Valid kills for %s is: %d\n", name, KillCount)
	fmt.Printf("Total attacks for %s is: %d\n", name, attackCount)
	fmt.Printf("Valid attacks for %s is: %d\n", name, validAttacks)
	fmt.Println(len(attackTimes))
	f.Close()
	//Extending kill data to hurt data
	for i, val := range attackframe {

		if intInSlice(val, killframe) {
			markedKillData[i] = killTimes[val]

		} else {
			markedKillData[i] = Kill{}
		}

	}
	i := 0
	for _, attack := range attackTimes {
		//weapon := markedFrameData[attack.attackFrame][attack.attacker].weapon
		attackData := AttackData{}
		var prevAttackerYaw float32
		var prevAttackerPitch float32
		var prevVictimYaw float32
		var prevVictimPitch float32
		var victimspotted_TargetPlayer bool
		var attackerspotted_TargetPlayer bool
		prevFrame := (attack.attackFrame - framesBeforeAttack) - framesPerSample
		if attack.attacker == nil {
			prevAttackerYaw = 0
			prevAttackerPitch = 0

		} else {
			prevAttackerYaw = markedFrameData[prevFrame][attack.attacker.UserID].Yaw
			prevAttackerPitch = markedFrameData[prevFrame][attack.attacker.UserID].Pitch
		}
		if attack.victim == nil {
			prevVictimYaw = 0
			prevVictimPitch = 0

		} else {
			prevVictimYaw = markedFrameData[prevFrame][attack.victim.UserID].Yaw
			prevVictimPitch = markedFrameData[prevFrame][attack.victim.UserID].Pitch
		}

		if attack.victim == nil {
			attackerspotted_TargetPlayer = false
			victimspotted_TargetPlayer = false
		} else if attack.attacker == nil {
			attackerspotted_TargetPlayer = false
			victimspotted_TargetPlayer = false
		} else {
			victimspotted_TargetPlayer = attack.victim.HasSpotted(attack.attacker)
			attackerspotted_TargetPlayer = attack.attacker.HasSpotted(attack.victim)
		}
		var attackerToVictim r3.Vector
		var attackerToVictimYaw float32
		var attackerToVictimPitch float32
		var sample int
		var attackerYaw float32
		var attackerPitch float32
		var victimToAttacker r3.Vector
		var victimToAttackerYaw float32
		var victimToAttackerPitch float32
		var victimYaw float32
		var victimPitch float32
		killdata := markedKillData[i]
		i++
		for sample = 0; sample < samplesPerAttack; sample++ {
			frame = framesPerSample*sample + (attack.attackFrame - framesBeforeAttack)
			victim := PlayerData{}
			attacker := PlayerData{}
			aid := 0
			vid := 0
			if attack.attacker == nil {
				attacker = PlayerData{}

			} else {
				aid = attack.attacker.UserID
				attacker = markedFrameData[frame][aid]

			}

			if attack.victim == nil {
				victim = PlayerData{}
			} else {
				vid = attack.victim.UserID
				victim = markedFrameData[frame][vid]
			}

			victimYaw = victim.Yaw
			victimPitch = victim.Pitch
			attackerYaw = attacker.Yaw
			attackerPitch = attacker.Pitch
			prevAttackerYaw = attackerYaw
			prevAttackerPitch = attackerPitch
			prevVictimYaw = victimYaw
			prevVictimPitch = victimPitch
			attackerToVictim = victim.Position.Sub(attacker.Position)
			victimToAttacker = attacker.Position.Sub(victim.Position)
			dXa := attackerToVictim.X
			dYa := attackerToVictim.Y
			dZa := attackerToVictim.Z
			dXv := victimToAttacker.X
			dYv := victimToAttacker.Y
			dZv := victimToAttacker.Z
			attackerToVictimYaw = 180 / math.Pi * float32(math.Atan2(dYa, dXa))
			attackerToVictimPitch = 180 / math.Pi * float32(math.Atan2(math.Sqrt(dXa*dXa+dYa*dYa), dZa))

			victimToAttackerYaw = 180 / math.Pi * float32(math.Atan2(dYv, dXv))
			victimToAttackerPitch = 180 / math.Pi * float32(math.Atan2(math.Sqrt(dXv*dXv+dYv*dYv), dZv))

			attackerYaw64 := float64(math.Pi / 180 * attackerYaw)
			attackerPitch64 := float64(math.Pi / 180 * attackerPitch)

			attackData.MapCrc = res.GetMapCrc()
			attackData.MapName = res.GetMapName()
			attackData.Frame = append(attackData.Frame, victim.Frame)

			//killdata
			if killdata.Assister == nil {
				attackData.Assisted = false
			} else {
				attackData.Assisted = true
			}
			attackData.KillPenetratedObjects = killdata.PenetratedObjects
			attackData.KillIsHeadshot = killdata.IsHeadshot
			attackData.KillAssistedFlash = killdata.AssistedFlash
			attackData.KillAttackerBlind = killdata.AttackerBlind
			attackData.KillNoScope = killdata.NoScope
			attackData.KillThroughSmoke = killdata.ThroughSmoke
			attackData.KillDistance = killdata.Distance
			if killdata.PenetratedObjects > 0 {
				attackData.IsWallBang = true
			} else {
				attackData.IsWallBang = false
			}

			attackData.AttackHitGroup = attack.HitGroup
			attackData.AttackGamePhase = attack.GamePhase
			attackData.IsWarmUpPeriodDuringAttack = attack.IsWarmupPeriod
			attackData.IsMatchStarted = attack.IsMatchStarted
			attackData.HealthDamageTaken = attack.HealthDamageTaken
			attackData.tick = append(attackData.tick, victim.Tick)

			//AttackerDATA
			attackData.attackerSteamID64 = attacker.SteamID64
			attackData.attackerFlashDuration = append(attackData.attackerFlashDuration, attacker.FlashDuration)
			attackData.attackerVelocity0 = append(attackData.attackerVelocity0, attacker.Velocity0)
			attackData.attackerVelocity1 = append(attackData.attackerVelocity1, attacker.Velocity1)
			attackData.attackerVelocity2 = append(attackData.attackerVelocity2, attacker.Velocity2)
			attackData.attackerLastAlivePosition0 = append(attackData.attackerLastAlivePosition0, attacker.LastAlivePosition0)
			attackData.attackerLastAlivePosition1 = append(attackData.attackerLastAlivePosition1, attacker.LastAlivePosition1)
			attackData.attackerLastAlivePosition2 = append(attackData.attackerLastAlivePosition2, attacker.LastAlivePosition2)
			attackData.attackerViewDirectionX = append(attackData.attackerViewDirectionX, attacker.ViewDirectionX)
			attackData.attackerViewDirectionY = append(attackData.attackerViewDirectionY, attacker.ViewDirectionY)
			attackData.attackerIsDucking = append(attackData.attackerIsDucking, attacker.IsDucking)
			attackData.attackerPlayerHealth = append(attackData.attackerPlayerHealth, attacker.PlayerHealth)
			attackData.attackerFiring = append(attackData.attackerFiring, attacker.Firing)
			attackData.attackerAmmoInMagazine = append(attackData.attackerAmmoInMagazine, attacker.AmmoInMagazine)
			attackData.attackerPing = append(attackData.attackerPing, attacker.Ping)
			attackData.attackerIsScoped = append(attackData.attackerIsScoped, attacker.IsScoped)
			attackData.attackerAmmoReserve = append(attackData.attackerAmmoReserve, attacker.AmmoReserve)
			attackData.attackerZoomLevel = append(attackData.attackerZoomLevel, attacker.ZoomLevel)
			attackData.attackerAmmoType = append(attackData.attackerAmmoType, attacker.AmmoType)
			attackData.attackerRecoilIndex = append(attackData.attackerRecoilIndex, attacker.RecoilIndex)
			attackData.attackerBurstMode = append(attackData.attackerBurstMode, attacker.BurstMode)
			attackData.attackerReloadVisuallyComplete = append(attackData.attackerReloadVisuallyComplete, attacker.ReloadVisuallyComplete)
			attackData.attackerSilencerOn = append(attackData.attackerSilencerOn, attacker.SilencerOn)
			attackData.attackerAccuracyPenalty = append(attackData.attackerAccuracyPenalty, attacker.AccuracyPenalty)
			attackData.attackerLastShotTime = append(attackData.attackerLastShotTime, attacker.LastShotTime)
			attackData.attackerDoneSwitchingSilencer = append(attackData.attackerDoneSwitchingSilencer, attacker.DoneSwitchingSilencer)
			attackData.attackerPostponeFireReadyTime = append(attackData.attackerPostponeFireReadyTime, attacker.PostponeFireReadyTime)
			attackData.attackerIronSightMode = append(attackData.attackerIronSightMode, attacker.IronSightMode)
			attackData.attackerWeaponMode = append(attackData.attackerWeaponMode, attacker.WeaponMode)
			attackData.attackerAmmoLeft = append(attackData.attackerAmmoLeft, attacker.AmmoLeft)
			attackData.attackerFlashTick = append(attackData.attackerFlashTick, attacker.FlashTick)
			attackData.attackerIsBot = append(attackData.attackerIsBot, attacker.IsBot)
			attackData.attackerIsControllingBot = append(attackData.attackerIsControllingBot, attacker.IsControllingBot)
			attackData.attackerIsConnected = append(attackData.attackerIsConnected, attacker.IsConnected)
			attackData.attackerIsDefusing = append(attackData.attackerIsDefusing, attacker.IsDefusing)
			attackData.attackerIsPlanting = append(attackData.attackerIsPlanting, attacker.IsPlanting)
			attackData.attackerIsReloading = append(attackData.attackerIsReloading, attacker.IsReloading)
			attackData.attackerIsUnknown = append(attackData.attackerIsUnknown, attacker.IsUnknown)
			attackData.attackerIsAlive = append(attackData.attackerIsAlive, attacker.IsAlive)
			attackData.attackerIsBlinded = append(attackData.attackerIsBlinded, attacker.IsBlinded)
			attackData.attackerIsAirborne = append(attackData.attackerIsAirborne, attacker.IsAirborne)
			attackData.attackerFlashDurationTime = append(attackData.attackerFlashDurationTime, attacker.FlashDurationTime)
			attackData.attackerFlashDurationTimeFull = append(attackData.attackerFlashDurationTimeFull, attacker.FlashDurationTimeFull)
			attackData.attackerIsWalking = append(attackData.attackerIsWalking, attacker.IsWalking)
			attackData.attackerIsDuckingInProgress = append(attackData.attackerIsDuckingInProgress, attacker.IsDuckingInProgress)
			attackData.attackerIsUnDuckingInProgress = append(attackData.attackerIsUnDuckingInProgress, attacker.IsUnDuckingInProgress)
			attackData.attackerIsStanding = append(attackData.attackerIsStanding, attacker.IsStanding)
			attackData.attackerHasHelmet = append(attackData.attackerHasHelmet, attacker.HasHelmet)
			attackData.attackerArmor = append(attackData.attackerArmor, attacker.Armor)
			attackData.attackerPositionX = append(attackData.attackerPositionX, attacker.PositionX)
			attackData.attackerPositionY = append(attackData.attackerPositionY, attacker.PositionY)
			attackData.attackerPositionZ = append(attackData.attackerPositionZ, attacker.PositionZ)
			attackData.attackerPositionEyesX = append(attackData.attackerPositionEyesX, attacker.PositionEyesX)
			attackData.attackerPositionEyesY = append(attackData.attackerPositionEyesY, attacker.PositionEyesY)
			attackData.attackerPositionEyesZ = append(attackData.attackerPositionEyesZ, attacker.PositionEyesZ)
			attackData.attackerHeadPositionX = append(attackData.attackerHeadPositionX, attacker.HeadPositionX)
			attackData.attackerHeadPositionY = append(attackData.attackerHeadPositionY, attacker.HeadPositionY)
			attackData.attackerHeadPositionZ = append(attackData.attackerHeadPositionZ, attacker.HeadPositionZ)
			attackData.attackerPlayerFlags = append(attackData.attackerPlayerFlags, attacker.PlayerFlags)
			attackData.attackerOnGround = append(attackData.attackerOnGround, attacker.OnGround)
			attackData.attackerDuckingKeyPressed = append(attackData.attackerDuckingKeyPressed, attacker.DuckingKeyPressed)
			attackData.attackerHasHeavyArmor = append(attackData.attackerHasHeavyArmor, attacker.HasHeavyArmor)
			attackData.attackerHasMovedSinceSpawn = append(attackData.attackerHasMovedSinceSpawn, attacker.HasMovedSinceSpawn)
			attackData.attackerIsGrabbingHostage = append(attackData.attackerIsGrabbingHostage, attacker.IsGrabbingHostage)
			attackData.attackerIsHoldingLookAtWeapon = append(attackData.attackerIsHoldingLookAtWeapon, attacker.IsHoldingLookAtWeapon)
			attackData.attackerIsLookingAtWeapon = append(attackData.attackerIsLookingAtWeapon, attacker.IsLookingAtWeapon)
			attackData.attackerIsRescuing = append(attackData.attackerIsRescuing, attacker.IsRescuing)
			attackData.attackerResumeZone = append(attackData.attackerResumeZone, attacker.ResumeZone)
			attackData.attackerStrafing = append(attackData.attackerStrafing, attacker.Strafing)
			attackData.attackerWaitForNoAttack = append(attackData.attackerWaitForNoAttack, attacker.WaitForNoAttack)
			attackData.attackerMolotovDamageTime = append(attackData.attackerMolotovDamageTime, attacker.MolotovDamageTime)
			attackData.attackerFlashMaxAlpha = append(attackData.attackerFlashMaxAlpha, attacker.FlashMaxAlpha)
			attackData.attackerLowerYawBodyTarget = append(attackData.attackerLowerYawBodyTarget, attacker.LowerYawBodyTarget)
			attackData.attackerThirdpersonRecoil = append(attackData.attackerThirdpersonRecoil, attacker.ThirdpersonRecoil)
			attackData.attackerBlockingUseActionInProgress = append(attackData.attackerBlockingUseActionInProgress, attacker.BlockingUseActionInProgress)
			attackData.attackerMoveState = append(attackData.attackerMoveState, attacker.MoveState)
			attackData.attackerPlayerState = append(attackData.attackerPlayerState, attacker.PlayerState)
			attackData.attackerViewPunchAngleX = append(attackData.attackerViewPunchAngleX, attacker.ViewPunchAngleX)
			attackData.attackerViewPunchAngleY = append(attackData.attackerViewPunchAngleY, attacker.ViewPunchAngleY)
			attackData.attackerViewPunchAngleZ = append(attackData.attackerViewPunchAngleZ, attacker.ViewPunchAngleZ)
			attackData.attackerAimPunchAngleX = append(attackData.attackerAimPunchAngleX, attacker.AimPunchAngleX)
			attackData.attackerAimPunchAngleY = append(attackData.attackerAimPunchAngleY, attacker.AimPunchAngleY)
			attackData.attackerAimPunchAngleZ = append(attackData.attackerAimPunchAngleZ, attacker.AimPunchAngleZ)
			attackData.attackerAimPunchAngleVelX = append(attackData.attackerAimPunchAngleVelX, attacker.AimPunchAngleVelX)
			attackData.attackerAimPunchAngleVelY = append(attackData.attackerAimPunchAngleVelY, attacker.AimPunchAngleVelY)
			attackData.attackerAimPunchAngleVelZ = append(attackData.attackerAimPunchAngleVelZ, attacker.AimPunchAngleVelZ)
			attackData.attackerInDuckJump = append(attackData.attackerInDuckJump, attacker.InDuckJump)
			attackData.attackerFOVRate = append(attackData.attackerFOVRate, attacker.FOVRate)
			attackData.attackerFallVelocity = append(attackData.attackerFallVelocity, attacker.FallVelocity)
			attackData.attackerStepSize = append(attackData.attackerStepSize, attacker.StepSize)
			attackData.attackerWearingSuit = append(attackData.attackerWearingSuit, attacker.WearingSuit)
			attackData.attackerPoisoned = append(attackData.attackerPoisoned, attacker.Poisoned)
			attackData.attackerDuckOverride = append(attackData.attackerDuckOverride, attacker.DuckOverride)
			attackData.attackerStamina = append(attackData.attackerStamina, attacker.Stamina)
			attackData.attackerVelocityModifier = append(attackData.attackerVelocityModifier, attacker.VelocityModifier)
			attackData.attackerDirection = append(attackData.attackerDirection, attacker.Direction)
			attackData.attackerShotsFired = append(attackData.attackerShotsFired, attacker.ShotsFired)
			attackData.attackerTimeOfLastInjury = append(attackData.attackerTimeOfLastInjury, attacker.TimeOfLastInjury)
			attackData.attackerRelativeDirectionOfLastInjury = append(attackData.attackerRelativeDirectionOfLastInjury, attacker.RelativeDirectionOfLastInjury)
			attackData.attackerVecLadderNormalX = append(attackData.attackerVecLadderNormalX, attacker.VecLadderNormalX)
			attackData.attackerVecLadderNormalY = append(attackData.attackerVecLadderNormalY, attacker.VecLadderNormalY)
			attackData.attackerVecLadderNormalZ = append(attackData.attackerVecLadderNormalZ, attacker.VecLadderNormalZ)
			attackData.attackerPhysicsCollisionState = append(attackData.attackerPhysicsCollisionState, attacker.PhysicsCollisionState)
			attackData.attackerWaterLevel = append(attackData.attackerWaterLevel, attacker.WaterLevel)
			attackData.attackerLifeState = append(attackData.attackerLifeState, attacker.LifeState)
			attackData.attackerLadderSurfaceProps = append(attackData.attackerLadderSurfaceProps, attacker.LadderSurfaceProps)
			attackData.attackerFOVStart = append(attackData.attackerFOVStart, attacker.FOVStart)
			attackData.attackerFOV = append(attackData.attackerFOV, attacker.FOV)
			attackData.attackerDefaultFOV = append(attackData.attackerDefaultFOV, attacker.DefaultFOV)
			attackData.attackerMaxSpeed = append(attackData.attackerMaxSpeed, attacker.MaxSpeed)
			attackData.attackerFOVTime = append(attackData.attackerFOVTime, attacker.FOVTime)
			attackData.attackerDuckSpeed = append(attackData.attackerDuckSpeed, attacker.DuckSpeed)
			attackData.attackerDuckAmount = append(attackData.attackerDuckAmount, attacker.DuckAmount)
			attackData.attackerPhysicsFlags = append(attackData.attackerPhysicsFlags, attacker.PhysicsFlags)
			attackData.attackerVecBaseVelocityX = append(attackData.attackerVecBaseVelocityX, attacker.VecBaseVelocityX)
			attackData.attackerVecBaseVelocityY = append(attackData.attackerVecBaseVelocityY, attacker.VecBaseVelocityY)
			attackData.attackerVecBaseVelocityZ = append(attackData.attackerVecBaseVelocityZ, attacker.VecBaseVelocityZ)
			attackData.attackerLaggedMovementValue = append(attackData.attackerLaggedMovementValue, attacker.LaggedMovementValue)
			attackData.attackerM3ReloadState = append(attackData.attackerM3ReloadState, attacker.M3ReloadState)
			attackData.attackerNOVAReloadState = append(attackData.attackerNOVAReloadState, attacker.NOVAReloadState)
			attackData.attackerSawedoffReloadState = append(attackData.attackerSawedoffReloadState, attacker.SawedoffReloadState)
			attackData.attackerXM1014ReloadState = append(attackData.attackerXM1014ReloadState, attacker.XM1014ReloadState)
			attackData.attackerMoveType = append(attackData.attackerMoveType, attacker.MoveType)
			attackData.attackerMoveCollide = append(attackData.attackerMoveCollide, attacker.MoveCollide)
			attackData.attackerUseLookAtAngle = append(attackData.attackerUseLookAtAngle, attacker.UseLookAtAngle)
			attackData.attackerSimulationTime = append(attackData.attackerSimulationTime, attacker.SimulationTime)
			attackData.attackerShadowCastDistance = append(attackData.attackerShadowCastDistance, attacker.ShadowCastDistance)
			attackData.attackerMaxFallVelocity = append(attackData.attackerMaxFallVelocity, attacker.MaxFallVelocity)
			attackData.attackerLastMadeNoiseTime = append(attackData.attackerLastMadeNoiseTime, attacker.LastMadeNoiseTime)
			attackData.attackerNextPrimaryAttack = append(attackData.attackerNextPrimaryAttack, attacker.NextPrimaryAttack)
			attackData.attackerNextSecondaryAttack = append(attackData.attackerNextSecondaryAttack, attacker.NextSecondaryAttack)
			attackData.attackerTimeWeaponIdle = append(attackData.attackerTimeWeaponIdle, attacker.TimeWeaponIdle)
			attackData.attackerWeaponState = append(attackData.attackerWeaponState, attacker.WeaponState)
			attackData.attackerBurstShotsRemaining = append(attackData.attackerBurstShotsRemaining, attacker.BurstShotsRemaining)
			attackData.attackerViewTarget = append(attackData.attackerViewTarget, attacker.ViewTarget)
			attackData.attackerForceBone = append(attackData.attackerForceBone, attacker.ForceBone)
			attackData.attackerDeltaYaw = append(attackData.attackerDeltaYaw, normalizeAngle(attackerYaw-prevAttackerYaw))
			attackData.attackerDeltaPitch = append(attackData.attackerDeltaPitch, attackerPitch-prevAttackerPitch)
			attackData.crosshairToVictimYaw = append(attackData.crosshairToVictimYaw, normalizeAngle(attackerToVictimYaw-attackerYaw))
			attackData.crosshairToVictimPitch = append(attackData.crosshairToVictimPitch, attackerToVictimPitch-attackerPitch)
			attackData.AttackerSpottedTargetPlayer = append(attackData.AttackerSpottedTargetPlayer, attackerspotted_TargetPlayer)
			attackData.attackerViewVectorX = append(attackData.attackerViewVectorX, r3.Vector{
				math.Cos(attackerYaw64) * math.Sin(attackerPitch64),
				math.Sin(attackerYaw64) * math.Sin(attackerPitch64),
				math.Cos(attackerPitch64)}.X)
			attackData.attackerViewVectorY = append(attackData.attackerViewVectorY, r3.Vector{
				math.Cos(attackerYaw64) * math.Sin(attackerPitch64),
				math.Sin(attackerYaw64) * math.Sin(attackerPitch64),
				math.Cos(attackerPitch64)}.Y)
			attackData.attackerViewVectorZ = append(attackData.attackerViewVectorZ, r3.Vector{
				math.Cos(attackerYaw64) * math.Sin(attackerPitch64),
				math.Sin(attackerYaw64) * math.Sin(attackerPitch64),
				math.Cos(attackerPitch64)}.Z)
			attackData.AttackerToVictimX = append(attackData.AttackerToVictimX, attackerToVictim.X)
			attackData.AttackerToVictimY = append(attackData.AttackerToVictimY, attackerToVictim.Y)
			attackData.AttackerToVictimZ = append(attackData.AttackerToVictimZ, attackerToVictim.Z)
			attackData.AttackerSpottedByVictim = append(attackData.AttackerSpottedByVictim, victimspotted_TargetPlayer)
			attackData.AttackerFootstep = append(attackData.AttackerFootstep, attacker.Footstep)
			attackData.AttackerDistance = append(attackData.AttackerDistance, float32(victimToAttacker.Norm()))
			attackData.AttackerFixedPitch = append(attackData.AttackerFixedPitch, attacker.Pitch)
			attackData.AttackerFallDamage = append(attackData.AttackerFallDamage, attacker.FallDamage)
			attackData.AttackerWeaponFireOnEmpty = append(attackData.AttackerWeaponFireOnEmpty, attacker.WeaponFireOnEmpty)

			//victimDATA
			attackData.victimSteamID64 = victim.SteamID64
			attackData.victimFlashDuration = append(attackData.victimFlashDuration, victim.FlashDuration)
			attackData.victimVelocity0 = append(attackData.victimVelocity0, victim.Velocity0)
			attackData.victimVelocity1 = append(attackData.victimVelocity1, victim.Velocity1)
			attackData.victimVelocity2 = append(attackData.victimVelocity2, victim.Velocity2)
			attackData.victimLastAlivePosition0 = append(attackData.victimLastAlivePosition0, victim.LastAlivePosition0)
			attackData.victimLastAlivePosition1 = append(attackData.victimLastAlivePosition1, victim.LastAlivePosition1)
			attackData.victimLastAlivePosition2 = append(attackData.victimLastAlivePosition2, victim.LastAlivePosition2)
			attackData.victimViewDirectionX = append(attackData.victimViewDirectionX, victim.ViewDirectionX)
			attackData.victimViewDirectionY = append(attackData.victimViewDirectionY, victim.ViewDirectionY)
			attackData.victimIsDucking = append(attackData.victimIsDucking, victim.IsDucking)
			attackData.victimPlayerHealth = append(attackData.victimPlayerHealth, victim.PlayerHealth)
			attackData.victimFiring = append(attackData.victimFiring, victim.Firing)
			attackData.victimAmmoInMagazine = append(attackData.victimAmmoInMagazine, victim.AmmoInMagazine)
			attackData.victimPing = append(attackData.victimPing, victim.Ping)
			attackData.victimIsScoped = append(attackData.victimIsScoped, victim.IsScoped)
			attackData.victimAmmoReserve = append(attackData.victimAmmoReserve, victim.AmmoReserve)
			attackData.victimZoomLevel = append(attackData.victimZoomLevel, victim.ZoomLevel)
			attackData.victimAmmoType = append(attackData.victimAmmoType, victim.AmmoType)
			attackData.victimRecoilIndex = append(attackData.victimRecoilIndex, victim.RecoilIndex)
			attackData.victimBurstMode = append(attackData.victimBurstMode, victim.BurstMode)
			attackData.victimReloadVisuallyComplete = append(attackData.victimReloadVisuallyComplete, victim.ReloadVisuallyComplete)
			attackData.victimSilencerOn = append(attackData.victimSilencerOn, victim.SilencerOn)
			attackData.victimAccuracyPenalty = append(attackData.victimAccuracyPenalty, victim.AccuracyPenalty)
			attackData.victimLastShotTime = append(attackData.victimLastShotTime, victim.LastShotTime)
			attackData.victimDoneSwitchingSilencer = append(attackData.victimDoneSwitchingSilencer, victim.DoneSwitchingSilencer)
			attackData.victimPostponeFireReadyTime = append(attackData.victimPostponeFireReadyTime, victim.PostponeFireReadyTime)
			attackData.victimIronSightMode = append(attackData.victimIronSightMode, victim.IronSightMode)
			attackData.victimWeaponMode = append(attackData.victimWeaponMode, victim.WeaponMode)
			attackData.victimAmmoLeft = append(attackData.victimAmmoLeft, victim.AmmoLeft)
			attackData.victimFlashTick = append(attackData.victimFlashTick, victim.FlashTick)
			attackData.victimIsBot = append(attackData.victimIsBot, victim.IsBot)
			attackData.victimIsControllingBot = append(attackData.victimIsControllingBot, victim.IsControllingBot)
			attackData.victimIsConnected = append(attackData.victimIsConnected, victim.IsConnected)
			attackData.victimIsDefusing = append(attackData.victimIsDefusing, victim.IsDefusing)
			attackData.victimIsPlanting = append(attackData.victimIsPlanting, victim.IsPlanting)
			attackData.victimIsReloading = append(attackData.victimIsReloading, victim.IsReloading)
			attackData.victimIsUnknown = append(attackData.victimIsUnknown, victim.IsUnknown)
			attackData.victimIsAlive = append(attackData.victimIsAlive, victim.IsAlive)
			attackData.victimIsBlinded = append(attackData.victimIsBlinded, victim.IsBlinded)
			attackData.victimIsAirborne = append(attackData.victimIsAirborne, victim.IsAirborne)
			attackData.victimFlashDurationTime = append(attackData.victimFlashDurationTime, victim.FlashDurationTime)
			attackData.victimFlashDurationTimeFull = append(attackData.victimFlashDurationTimeFull, victim.FlashDurationTimeFull)
			attackData.victimIsWalking = append(attackData.victimIsWalking, victim.IsWalking)
			attackData.victimIsDuckingInProgress = append(attackData.victimIsDuckingInProgress, victim.IsDuckingInProgress)
			attackData.victimIsUnDuckingInProgress = append(attackData.victimIsUnDuckingInProgress, victim.IsUnDuckingInProgress)
			attackData.victimIsStanding = append(attackData.victimIsStanding, victim.IsStanding)
			attackData.victimHasHelmet = append(attackData.victimHasHelmet, victim.HasHelmet)
			attackData.victimArmor = append(attackData.victimArmor, victim.Armor)
			attackData.victimPositionX = append(attackData.victimPositionX, victim.PositionX)
			attackData.victimPositionY = append(attackData.victimPositionY, victim.PositionY)
			attackData.victimPositionZ = append(attackData.victimPositionZ, victim.PositionZ)
			attackData.victimPositionEyesX = append(attackData.victimPositionEyesX, victim.PositionEyesX)
			attackData.victimPositionEyesY = append(attackData.victimPositionEyesY, victim.PositionEyesY)
			attackData.victimPositionEyesZ = append(attackData.victimPositionEyesZ, victim.PositionEyesZ)
			attackData.victimHeadPositionX = append(attackData.victimHeadPositionX, victim.HeadPositionX)
			attackData.victimHeadPositionY = append(attackData.victimHeadPositionY, victim.HeadPositionY)
			attackData.victimHeadPositionZ = append(attackData.victimHeadPositionZ, victim.HeadPositionZ)
			attackData.victimPlayerFlags = append(attackData.victimPlayerFlags, victim.PlayerFlags)
			attackData.victimOnGround = append(attackData.victimOnGround, victim.OnGround)
			attackData.victimDuckingKeyPressed = append(attackData.victimDuckingKeyPressed, victim.DuckingKeyPressed)
			attackData.victimHasHeavyArmor = append(attackData.victimHasHeavyArmor, victim.HasHeavyArmor)
			attackData.victimHasMovedSinceSpawn = append(attackData.victimHasMovedSinceSpawn, victim.HasMovedSinceSpawn)
			attackData.victimIsGrabbingHostage = append(attackData.victimIsGrabbingHostage, victim.IsGrabbingHostage)
			attackData.victimIsHoldingLookAtWeapon = append(attackData.victimIsHoldingLookAtWeapon, victim.IsHoldingLookAtWeapon)
			attackData.victimIsLookingAtWeapon = append(attackData.victimIsLookingAtWeapon, victim.IsLookingAtWeapon)
			attackData.victimIsRescuing = append(attackData.victimIsRescuing, victim.IsRescuing)
			attackData.victimResumeZone = append(attackData.victimResumeZone, victim.ResumeZone)
			attackData.victimStrafing = append(attackData.victimStrafing, victim.Strafing)
			attackData.victimWaitForNoAttack = append(attackData.victimWaitForNoAttack, victim.WaitForNoAttack)
			attackData.victimMolotovDamageTime = append(attackData.victimMolotovDamageTime, victim.MolotovDamageTime)
			attackData.victimFlashMaxAlpha = append(attackData.victimFlashMaxAlpha, victim.FlashMaxAlpha)
			attackData.victimLowerYawBodyTarget = append(attackData.victimLowerYawBodyTarget, victim.LowerYawBodyTarget)
			attackData.victimThirdpersonRecoil = append(attackData.victimThirdpersonRecoil, victim.ThirdpersonRecoil)
			attackData.victimBlockingUseActionInProgress = append(attackData.victimBlockingUseActionInProgress, victim.BlockingUseActionInProgress)
			attackData.victimMoveState = append(attackData.victimMoveState, victim.MoveState)
			attackData.victimPlayerState = append(attackData.victimPlayerState, victim.PlayerState)
			attackData.victimViewPunchAngleX = append(attackData.victimViewPunchAngleX, victim.ViewPunchAngleX)
			attackData.victimViewPunchAngleY = append(attackData.victimViewPunchAngleY, victim.ViewPunchAngleY)
			attackData.victimViewPunchAngleZ = append(attackData.victimViewPunchAngleZ, victim.ViewPunchAngleZ)
			attackData.victimAimPunchAngleX = append(attackData.victimAimPunchAngleX, victim.AimPunchAngleX)
			attackData.victimAimPunchAngleY = append(attackData.victimAimPunchAngleY, victim.AimPunchAngleY)
			attackData.victimAimPunchAngleZ = append(attackData.victimAimPunchAngleZ, victim.AimPunchAngleZ)
			attackData.victimAimPunchAngleVelX = append(attackData.victimAimPunchAngleVelX, victim.AimPunchAngleVelX)
			attackData.victimAimPunchAngleVelY = append(attackData.victimAimPunchAngleVelY, victim.AimPunchAngleVelY)
			attackData.victimAimPunchAngleVelZ = append(attackData.victimAimPunchAngleVelZ, victim.AimPunchAngleVelZ)
			attackData.victimInDuckJump = append(attackData.victimInDuckJump, victim.InDuckJump)
			attackData.victimFOVRate = append(attackData.victimFOVRate, victim.FOVRate)
			attackData.victimFallVelocity = append(attackData.victimFallVelocity, victim.FallVelocity)
			attackData.victimStepSize = append(attackData.victimStepSize, victim.StepSize)
			attackData.victimWearingSuit = append(attackData.victimWearingSuit, victim.WearingSuit)
			attackData.victimPoisoned = append(attackData.victimPoisoned, victim.Poisoned)
			attackData.victimDuckOverride = append(attackData.victimDuckOverride, victim.DuckOverride)
			attackData.victimStamina = append(attackData.victimStamina, victim.Stamina)
			attackData.victimVelocityModifier = append(attackData.victimVelocityModifier, victim.VelocityModifier)
			attackData.victimDirection = append(attackData.victimDirection, victim.Direction)
			attackData.victimShotsFired = append(attackData.victimShotsFired, victim.ShotsFired)
			attackData.victimTimeOfLastInjury = append(attackData.victimTimeOfLastInjury, victim.TimeOfLastInjury)
			attackData.victimRelativeDirectionOfLastInjury = append(attackData.victimRelativeDirectionOfLastInjury, victim.RelativeDirectionOfLastInjury)
			attackData.victimVecLadderNormalX = append(attackData.victimVecLadderNormalX, victim.VecLadderNormalX)
			attackData.victimVecLadderNormalY = append(attackData.victimVecLadderNormalY, victim.VecLadderNormalY)
			attackData.victimVecLadderNormalZ = append(attackData.victimVecLadderNormalZ, victim.VecLadderNormalZ)
			attackData.victimPhysicsCollisionState = append(attackData.victimPhysicsCollisionState, victim.PhysicsCollisionState)
			attackData.victimWaterLevel = append(attackData.victimWaterLevel, victim.WaterLevel)
			attackData.victimLifeState = append(attackData.victimLifeState, victim.LifeState)
			attackData.victimLadderSurfaceProps = append(attackData.victimLadderSurfaceProps, victim.LadderSurfaceProps)
			attackData.victimFOVStart = append(attackData.victimFOVStart, victim.FOVStart)
			attackData.victimFOV = append(attackData.victimFOV, victim.FOV)
			attackData.victimDefaultFOV = append(attackData.victimDefaultFOV, victim.DefaultFOV)
			attackData.victimMaxSpeed = append(attackData.victimMaxSpeed, victim.MaxSpeed)
			attackData.victimFOVTime = append(attackData.victimFOVTime, victim.FOVTime)
			attackData.victimDuckSpeed = append(attackData.victimDuckSpeed, victim.DuckSpeed)
			attackData.victimDuckAmount = append(attackData.victimDuckAmount, victim.DuckAmount)
			attackData.victimPhysicsFlags = append(attackData.victimPhysicsFlags, victim.PhysicsFlags)
			attackData.victimVecBaseVelocityX = append(attackData.victimVecBaseVelocityX, victim.VecBaseVelocityX)
			attackData.victimVecBaseVelocityY = append(attackData.victimVecBaseVelocityY, victim.VecBaseVelocityY)
			attackData.victimVecBaseVelocityZ = append(attackData.victimVecBaseVelocityZ, victim.VecBaseVelocityZ)
			attackData.victimLaggedMovementValue = append(attackData.victimLaggedMovementValue, victim.LaggedMovementValue)
			attackData.victimM3ReloadState = append(attackData.victimM3ReloadState, victim.M3ReloadState)
			attackData.victimNOVAReloadState = append(attackData.victimNOVAReloadState, victim.NOVAReloadState)
			attackData.victimSawedoffReloadState = append(attackData.victimSawedoffReloadState, victim.SawedoffReloadState)
			attackData.victimXM1014ReloadState = append(attackData.victimXM1014ReloadState, victim.XM1014ReloadState)
			attackData.victimMoveType = append(attackData.victimMoveType, victim.MoveType)
			attackData.victimMoveCollide = append(attackData.victimMoveCollide, victim.MoveCollide)
			attackData.victimUseLookAtAngle = append(attackData.victimUseLookAtAngle, victim.UseLookAtAngle)
			attackData.victimSimulationTime = append(attackData.victimSimulationTime, victim.SimulationTime)
			attackData.victimShadowCastDistance = append(attackData.victimShadowCastDistance, victim.ShadowCastDistance)
			attackData.victimMaxFallVelocity = append(attackData.victimMaxFallVelocity, victim.MaxFallVelocity)
			attackData.victimLastMadeNoiseTime = append(attackData.victimLastMadeNoiseTime, victim.LastMadeNoiseTime)
			attackData.victimNextPrimaryAttack = append(attackData.victimNextPrimaryAttack, victim.NextPrimaryAttack)
			attackData.victimNextSecondaryAttack = append(attackData.victimNextSecondaryAttack, victim.NextSecondaryAttack)
			attackData.victimTimeWeaponIdle = append(attackData.victimTimeWeaponIdle, victim.TimeWeaponIdle)
			attackData.victimWeaponState = append(attackData.victimWeaponState, victim.WeaponState)
			attackData.victimBurstShotsRemaining = append(attackData.victimBurstShotsRemaining, victim.BurstShotsRemaining)
			attackData.victimViewTarget = append(attackData.victimViewTarget, victim.ViewTarget)
			attackData.victimForceBone = append(attackData.victimForceBone, victim.ForceBone)
			attackData.victimDeltaYaw = append(attackData.victimDeltaYaw, normalizeAngle(victimYaw-prevVictimYaw))
			attackData.victimDeltaPitch = append(attackData.victimDeltaPitch, victimPitch-prevVictimPitch)
			attackData.crosshairToAttackerYaw = append(attackData.crosshairToAttackerYaw, normalizeAngle(victimToAttackerYaw-victimYaw))
			attackData.crosshairToAttackerPitch = append(attackData.crosshairToAttackerPitch, victimToAttackerPitch-victimPitch)
			attackData.VictimSpottedTargetPlayer = append(attackData.VictimSpottedTargetPlayer, victimspotted_TargetPlayer)
			attackData.VictimToAttackerX = append(attackData.VictimToAttackerX, victimToAttacker.X)
			attackData.VictimToAttackerY = append(attackData.VictimToAttackerY, victimToAttacker.Y)
			attackData.VictimToAttackerZ = append(attackData.VictimToAttackerZ, victimToAttacker.Z)
			attackData.VictimSpottedByAttacker = append(attackData.VictimSpottedByAttacker, attackerspotted_TargetPlayer)
			attackData.VictimFootstep = append(attackData.VictimFootstep, victim.Footstep)
			attackData.VictimDistance = append(attackData.VictimDistance, float32(attackerToVictim.Norm()))
			attackData.VictimFixedPitch = append(attackData.VictimFixedPitch, victim.Pitch)
			attackData.VictimFallDamage = append(attackData.VictimFallDamage, victim.FallDamage)
			attackData.VictimWeaponFireOnEmpty = append(attackData.VictimWeaponFireOnEmpty, victim.WeaponFireOnEmpty)
		}
		modelData = append(modelData, attackData)

	}
	f.Close()

	fmt.Println("done Parsing demo", name)
	CsvExport(name, modelData, framerate, samplesPerAttack)

}

func Contains1(array []int, val int) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}

	return false
}

// Finds the index of a value in an array
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

/*
func framesandticks(frame int,tick int, frames []int,ticklist int[]) frame,tick{
if frames != nil {
					if Contains1(frames, frame) {
						if (frame > frames[len(frames)]) || (frame < 0) {
							fr = -1
							tic = -1
						} else {
							fr = frame
							ind := indexOf(frame, frames)
							tic = ticklist[ind]
							prevind = ind
						}
					} else {
						fr = -1
						tic = ticklist[prevind+1]
						prevind++
					}
				}else{
					fr = frame
					tic = gs.IngameTick()
				}*/
