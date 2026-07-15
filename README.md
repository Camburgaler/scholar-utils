# Scholar

A Dark Souls II: SotFS build optimizer

## Setup

### Get param data

1. Install Dark Souls II: Scholar of the First Sin
2. Install [Smithbox](https://github.com/vawser/Smithbox)
3. Open Smithbox and create a new DS2:SotFS project _with param headers_
4. Go to the param editor and export as CSV the following params:
    - ArmorParam
    - ArmorReinforceParam
    - CustomAttrSpecParam
    - ItemParam
    - LevelUpStatusCalcParam
    - PlayerLevelUpSoulsParam
    - RingParam
    - VowParam
    - WeaponParam
    - WeaponReinforceParam
    - WeaponStatsAffectParam
5. Move the resulting files into a folder called `inputs` at the root of this project

### Get EMEVD data

1. Install [WitchyBND](https://github.com/ividyon/WitchyBND)
2. Using WitchyBND, unpack the `enc.regulation.bnd.dcx` file in your `Dark Souls II Scholar of the First Sin/Game` directory
3. Install [DarkScript3](https://github.com/AinTunez/DarkScript3)
4. Using DarkScript3, open the following files and convert them to JS:
    - `Dark Souls II Scholar of the First Sin/Game/enc_regulation-bnd-dcx/SpEffectArmor.emevd` -> `SpEffectArmor.emevd.js`
    - `Dark Souls II Scholar of the First Sin/Game/enc_regulation-bnd-dcx/SpEffectRing.emevd` -> `SpEffectRing.emevd.js`
    - `Dark Souls II Scholar of the First Sin/Game/enc_regulation-bnd-dcx/SpEffectWeapon.emevd` -> `SpEffectWeapon.emevd.js`
5. Move the resulting files into a folder called `inputs` at the root of this project

## Usage

1. Install [Go](https://go.dev/doc/install).
2. Run the following command:

```bash
go run .
```
