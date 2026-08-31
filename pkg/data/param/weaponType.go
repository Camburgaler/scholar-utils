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

		allowMagic                      bool  //nolint:unused
		allowMiracle                    bool  //nolint:unused
		allowPyromancy                  bool  //nolint:unused
		allowDarkMagic                  bool  //nolint:unused
		rightStance1HMotionID           id.ID //nolint:unused
		leftStance1HMotionID            id.ID //nolint:unused
		rightStance2HMotionID           id.ID //nolint:unused
		leftStance2HMotionID            id.ID //nolint:unused
		rightStanceDual1HMotionID       id.ID //nolint:unused
		leftStanceDual1HMotionID        id.ID //nolint:unused
		rightStanceDual2HMotionID       id.ID //nolint:unused
		leftStanceDual2HMotionID        id.ID //nolint:unused
		changeWeaponAnim                id.ID //nolint:unused
		dmyPolyID1HandStanceRightChr    id.ID //nolint:unused
		dmyPolyID1HandStanceRightWeapon id.ID //nolint:unused
		dmyPolyID1HandStanceLeftChr     id.ID //nolint:unused
		dmyPolyID1HandStanceLeftWeapon  id.ID //nolint:unused
		dmyPolyIDR2HandStanceChrMain    id.ID //nolint:unused
		dmyPolyIDR2HandStanceWeaponMain id.ID //nolint:unused
		dmyPolyIDR2HandStanceChrSub     id.ID //nolint:unused
		dmyPolyIDR2HandStanceWeaponSub  id.ID //nolint:unused
		dmyPolyIDL2HandStanceChrMain    id.ID //nolint:unused
		dmyPolyIDL2HandStanceWeaponMain id.ID //nolint:unused
		dmyPolyIDL2HandStanceChrSub     id.ID //nolint:unused
		dmyPolyIDL2HandStanceWeaponSub  id.ID //nolint:unused
		dmyPolyIDRightStoreChr          id.ID //nolint:unused
		dmyPolyIDRightStoreWeapon       id.ID //nolint:unused
		dmyPolyIDLeftStoreChr           id.ID //nolint:unused
		dmyPolyIDLeftStoreWeapon        id.ID //nolint:unused
		shootCategory                   id.ID //nolint:unused

		// The type of this weapon
		WeaponType id.ID

		attackSeCategory id.ID //nolint:unused

		// Determines whether the player can dual wield this weapon
		DualWieldingPermission id.ID

		weaponStabilizeBlengWeight          id.ID   //nolint:unused
		weaponSpineTwistBlendWeight         id.ID   //nolint:unused
		parryStabYawCorrectAngle            float64 //nolint:unused
		guardBreakStabYawCorrectAngle       float64 //nolint:unused
		backStabYawCorrectAngle             float64 //nolint:unused
		guardPassiveLevelType               id.ID   //nolint:unused
		guardReleaseSpeedCategory           id.ID   //nolint:unused
		parryType                           id.ID   //nolint:unused
		allowDarkMiracle                    bool    //nolint:unused
		leftDamageScale                     float64 //nolint:unused
		leftStaminaDamageScale              float64 //nolint:unused
		leftDurabilityDamageScale           float64 //nolint:unused
		rightDamageScale                    float64 //nolint:unused
		rightDtaminaDamageScale             float64 //nolint:unused
		rightDurabilityDamageScale          float64 //nolint:unused
		counterDamageScale                  float64 //nolint:unused
		fallStrikeFallDistRate              float64 //nolint:unused
		jumpAttackVerticalSpeedRate         float64 //nolint:unused
		jumpAttackSideDistMinRate           float64 //nolint:unused
		jumpAttackSideDistMaxRate           float64 //nolint:unused
		jumpAttackSideSpeedMinRate          float64 //nolint:unused
		jumpAttackSideSpeedMaxRate          float64 //nolint:unused
		jumpFallAttackFallDistRate          float64 //nolint:unused
		jumpFallAttackVerticalSpeedRate     float64 //nolint:unused
		jumpFallAttackSideDistMinRate       float64 //nolint:unused
		jumpFallAttackSideDistMaxRate       float64 //nolint:unused
		jumpFallAttackSideSpeedMinRate      float64 //nolint:unused
		jumpFallAttackSideSpeedMaxRate      float64 //nolint:unused
		fallStabType                        id.ID   //nolint:unused
		damageWeaponType                    id.ID   //nolint:unused
		menuShotRange                       int     //nolint:unused
		backStabAttackAnimID                id.ID   //nolint:unused
		gdBreakStabAttackAnimID             id.ID   //nolint:unused
		lookAtRotationLeftRate              float64 //nolint:unused
		lookAtRotationRightRate             float64 //nolint:unused
		lookAtRotationUpperRate             float64 //nolint:unused
		lookAtRotationLowerRate             float64 //nolint:unused
		lookAtVisibilitySideRate            float64 //nolint:unused
		lookAtVisibilityUpperRate           float64 //nolint:unused
		lookAtRotationSideSpeedRate         float64 //nolint:unused
		lookAtRotationUpperSpeedRate        float64 //nolint:unused
		staminaRecoveryScale                float64 //nolint:unused
		stabilityAddRateSingle1Hand         float64 //nolint:unused
		stabilityAddRateSingle2Handed       float64 //nolint:unused
		stabilityAddRateDual1Hand           float64 //nolint:unused
		stabilityAddRateDual2Handed         float64 //nolint:unused
		meleeAutoPivotRotationSpeedScale    float64 //nolint:unused
		meleeAutoPivotLimitAngleScale       float64 //nolint:unused
		manualPivotRotationSpeedScale       float64 //nolint:unused
		manualPivotLimitAngleScale          float64 //nolint:unused
		moveAccelerationScale               float64 //nolint:unused
		moveMaxSpeedScale                   float64 //nolint:unused
		moveDecelerationScale               float64 //nolint:unused
		moveRotationSpeedScale              float64 //nolint:unused
		canPredictShot                      bool    //nolint:unused
		shotSeCategory                      id.ID   //nolint:unused
		dummy                               []byte  //nolint:unused
		rangedAutoPivotRotationSpeedScale   float64 //nolint:unused
		rangedManualPivotRotationSpeedScale float64 //nolint:unused
		spellChantTimeScale                 float64 //nolint:unused
		toughnessPeriodScale                float64 //nolint:unused
		decreaseSpeedRateParallelWalk       float64 //nolint:unused
		decreaseSpeedRateWalk               float64 //nolint:unused
		decreaseSpeedRateLockOnWalk         float64 //nolint:unused
		afterAttackAccelerationScale        float64 //nolint:unused
		afterAttackMaxSpeedScale            float64 //nolint:unused
		afterAttackDecelerationScale        float64 //nolint:unused
		afterAttackRotationSpeedScale       float64 //nolint:unused
		weaponMenuCategory                  id.ID   //nolint:unused
		snipeCameraParamID                  id.ID   //nolint:unused
		canGuardAttack                      bool    //nolint:unused
		unk138                              []byte  //nolint:unused
		abyssRate                           float64 //nolint:unused
		dummy1                              []byte  //nolint:unused
	}
)
