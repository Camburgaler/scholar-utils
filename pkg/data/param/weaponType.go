package param

import "github.com/Camburgaler/scholar-utils/pkg/data/param/id"

// ValidWeaponTypeIDs is a list of valid WeaponTypeParam IDs
var ValidWeaponTypeIDs = []id.Range{
	{Start: 10, End: 1640},
}

type (
	// WeaponType is a struct for storing data from WeaponTypeParam.csv
	//
	// WeaponTypeParam.csv contains information about weapon types
	WeaponType struct {
		// The unique ID for this weapon type
		ID id.ID

		// The name of this weapon type
		Name string

		allowMagic                      bool
		allowMiracle                    bool
		allowPyromancy                  bool
		allowDarkMagic                  bool
		rightStance1HMotionID           id.ID
		leftStance1HMotionID            id.ID
		rightStance2HMotionID           id.ID
		leftStance2HMotionID            id.ID
		rightStanceDual1HMotionID       id.ID
		leftStanceDual1HMotionID        id.ID
		rightStanceDual2HMotionID       id.ID
		leftStanceDual2HMotionID        id.ID
		changeWeaponAnim                id.ID
		dmyPolyID1HandStanceRightChr    id.ID
		dmyPolyID1HandStanceRightWeapon id.ID
		dmyPolyID1HandStanceLeftChr     id.ID
		dmyPolyID1HandStanceLeftWeapon  id.ID
		dmyPolyIDR2HandStanceChrMain    id.ID
		dmyPolyIDR2HandStanceWeaponMain id.ID
		dmyPolyIDR2HandStanceChrSub     id.ID
		dmyPolyIDR2HandStanceWeaponSub  id.ID
		dmyPolyIDL2HandStanceChrMain    id.ID
		dmyPolyIDL2HandStanceWeaponMain id.ID
		dmyPolyIDL2HandStanceChrSub     id.ID
		dmyPolyIDL2HandStanceWeaponSub  id.ID
		dmyPolyIDRightStoreChr          id.ID
		dmyPolyIDRightStoreWeapon       id.ID
		dmyPolyIDLeftStoreChr           id.ID
		dmyPolyIDLeftStoreWeapon        id.ID
		shootCategory                   id.ID

		// The type of this weapon
		WeaponType id.ID

		attackSeCategory id.ID

		// Determines whether the player can dual wield this weapon
		DualWieldingPermission id.ID

		weaponStabilizeBlengWeight          id.ID
		weaponSpineTwistBlendWeight         id.ID
		parryStabYawCorrectAngle            float64
		guardBreakStabYawCorrectAngle       float64
		backStabYawCorrectAngle             float64
		guardPassiveLevelType               id.ID
		guardReleaseSpeedCategory           id.ID
		parryType                           id.ID
		allowDarkMiracle                    bool
		leftDamageScale                     float64
		leftStaminaDamageScale              float64
		leftDurabilityDamageScale           float64
		rightDamageScale                    float64
		rightDtaminaDamageScale             float64
		rightDurabilityDamageScale          float64
		counterDamageScale                  float64
		fallStrikeFallDistRate              float64
		jumpAttackVerticalSpeedRate         float64
		jumpAttackSideDistMinRate           float64
		jumpAttackSideDistMaxRate           float64
		jumpAttackSideSpeedMinRate          float64
		jumpAttackSideSpeedMaxRate          float64
		jumpFallAttackFallDistRate          float64
		jumpFallAttackVerticalSpeedRate     float64
		jumpFallAttackSideDistMinRate       float64
		jumpFallAttackSideDistMaxRate       float64
		jumpFallAttackSideSpeedMinRate      float64
		jumpFallAttackSideSpeedMaxRate      float64
		fallStabType                        id.ID
		damageWeaponType                    id.ID
		menuShotRange                       int
		backStabAttackAnimID                id.ID
		gdBreakStabAttackAnimID             id.ID
		lookAtRotationLeftRate              float64
		lookAtRotationRightRate             float64
		lookAtRotationUpperRate             float64
		lookAtRotationLowerRate             float64
		lookAtVisibilitySideRate            float64
		lookAtVisibilityUpperRate           float64
		lookAtRotationSideSpeedRate         float64
		lookAtRotationUpperSpeedRate        float64
		staminaRecoveryScale                float64
		stabilityAddRateSingle1Hand         float64
		stabilityAddRateSingle2Handed       float64
		stabilityAddRateDual1Hand           float64
		stabilityAddRateDual2Handed         float64
		meleeAutoPivotRotationSpeedScale    float64
		meleeAutoPivotLimitAngleScale       float64
		manualPivotRotationSpeedScale       float64
		manualPivotLimitAngleScale          float64
		moveAccelerationScale               float64
		moveMaxSpeedScale                   float64
		moveDecelerationScale               float64
		moveRotationSpeedScale              float64
		canPredictShot                      bool
		shotSeCategory                      id.ID
		dummy                               []byte
		rangedAutoPivotRotationSpeedScale   float64
		rangedManualPivotRotationSpeedScale float64
		spellChantTimeScale                 float64
		toughnessPeriodScale                float64
		decreaseSpeedRateParallelWalk       float64
		decreaseSpeedRateWalk               float64
		decreaseSpeedRateLockOnWalk         float64
		afterAttackAccelerationScale        float64
		afterAttackMaxSpeedScale            float64
		afterAttackDecelerationScale        float64
		afterAttackRotationSpeedScale       float64
		weaponMenuCategory                  id.ID
		snipeCameraParamID                  id.ID
		canGuardAttack                      bool
		unk138                              []byte
		abyssRate                           float64
		dummy1                              []byte
	}
)
