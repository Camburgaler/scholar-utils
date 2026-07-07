# Scholar

A Dark Souls II: SotFS build optimizer

## Setup

### Get param data

1. Install Dark Souls II: Scholar of the First Sin
2. Install [Smithbox](https://github.com/vawser/Smithbox)
3. Open Smithbox and create a new DS2:SotFS project _with param headers_
4. Go to the param editor and export as CSV the following files into the `inputs` folder:
    - ArmorParam
    - ArmorReinforceParam
    - CustomAttrSpecParam
    - ItemParam
    - RingParam
    - WeaponParam
    - WeaponReinforceParam
    - WeaponStatsAffectParam

### Get EMEVD data

Idk what goes here yet, I haven't accessed EMEVD data before. But you'll need DarkScript3, and you'll need to do this to get SpEffect data.

## Usage

1. Install [Go](https://go.dev/doc/install).
2. Run the following command:

```bash
go run .
```
