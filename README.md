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
        - Note: You may need to override the row names if the [feature](https://github.com/vawser/Smithbox/issues/527) I requested hasn't been implemented yet.
          ![alt text](https://private-user-images.githubusercontent.com/33138104/623001185-6b95d2b7-c684-4c7e-8faf-1375dffbacdf.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3ODQyNTY3OTAsIm5iZiI6MTc4NDI1NjQ5MCwicGF0aCI6Ii8zMzEzODEwNC82MjMwMDExODUtNmI5NWQyYjctYzY4NC00YzdlLThmYWYtMTM3NWRmZmJhY2RmLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNjA3MTclMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjYwNzE3VDAyNDgxMFomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTQ3NzRhOGNmMWZiMjE0ZDgzNjI5ZmJjNDI2NGRmMDIzMjkxZWZlNzE1ZjRkZDUyMjRkYzY0ZjkyNjRhNDRiMGUmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0JnJlc3BvbnNlLWNvbnRlbnQtdHlwZT1pbWFnZSUyRnBuZyJ9.rXDXKrgkI0w8i4c8BPgmNrcrDTzWF1CEIkqggk__UsU)
    - PlayerLevelUpSoulsParam
    - PlayerStatusParam
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
