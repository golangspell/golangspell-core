# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.7.0] - 2024-07-01

### Added
- addgateway command

## [0.6.5] - 2024-06-07

### Changed
- Adjusted AddStatementToFunction behavior and renamed to AddCallToFunction

## [0.6.4] - 2024-06-07

### Added
- New feature of the code file management - add attribute to struct

## [0.6.3] - 2024-06-06

### Changed
- Adjusted version

## [0.6.2] - 2024-06-06

### Changed
- Fixed import at addusecase command

## [0.6.1] - 2024-06-06

### Changed
- Changed deprecated MIME type MIMEApplicationJSONCharsetUTF8 to MIMEApplicationJSON

## [0.6.0] - 2024-06-06

### Added
- addusecase command
- Code file domains, to provide features for code parsing and editing through syntax tree analysis
- RestClient added to the init template

### Changed
- Updated Golang version
- Updated dependencies

## [0.5.17] - 2022-09-22

### Added
- Added the CORS headers HeaderAccessControlAllowOrigin and HeaderAccessControlAllowMethods

## [0.5.16] - 2022-09-16

### Changed
- Fixed version

## [0.5.15] - 2022-09-16

### Changed
- Updated the Dockerfile with the newer Alpine releases

## [0.5.14] - 2022-07-28

### Changed
- Adjusted template main.got - added import to golangspell/gateway/filesystem - Config Repository

## [0.5.13] - 2022-07-28

### Changed
- Adjusted template cmd/buildconfig.got - added Spell version

## [0.5.12] - 2022-07-27

### Changed
- Fixed bug 

## [0.5.11] - 2022-07-27

### Changed
- Fixed context bug 

## [0.5.10] - 2022-07-27

### Changed
- Adjusted template rederer to the new go mod directory structure

## [0.5.9] - 2022-07-26

### Changed
- Updated logic for handling new directory structure of go mod 

## [0.5.8] - 2021-07-26

### Changed
- Updated dependencies

## [0.5.7] - 2021-07-26

### Changed
- Updated dependencies

## [0.5.6] - 2021-08-05

### Changed
- Fixed templates imports

## [0.5.5] - 2021-08-05

### Changed
- Removed viper dependency

## [0.5.4] - 2021-08-05

### Changed
- Updated dependencies

## [0.5.3] - 2021-08-04

### Changed
- Updated Cobra and Viper versions

## [0.5.2] - 2020-12-27

### Added
- template/init/.github/workflows/test.yml for go test and add badge into README;
- template/init/.github/workflows/lint.yml to run golangci-lint

### Changed
- Remove strategy matrix since it use the same pattern as go template and creates an issue inside the template and keep test only using ubuntu-latest.

## [0.5.1] - 2020-07-23

### Changed
- Adjusted log encoder for timestamp

## [0.5.0] - 2020-05-14

### Changed
- Moved to organization's repository

## [0.4.10] - 2020-05-10

### Added
- CHANGELOG.md created

### Changed
- Improved Dockerfile reducing the maintainance effort

## [0.4.9] - 2020-04-17

### Changed
- Updated documentation

## [0.4.8] - 2020-04-15

### Changed
- Adjusted string template rendering

## [0.4.7] - 2020-04-14

### Changed
- Adjusted string templating

## [0.4.6] - 2020-04-14

### Changed
- Adjusted string templating feature

## [0.4.5] - 2020-04-08

### Changed
- Adjusted initspell command feedback text, added lint and adjusted code

## [0.4.4] - 2020-04-08

### Changed
- addspellcommand adjusted to add the new command specification to the Spell's buildconfig command

## [0.4.3] - 2020-04-08

### Added
- Included golangci-lint

### Changed
-  Fixed lint issues

## [0.4.2] - 2020-04-03

### Changed
-  Removed unused imports from template

## [0.4.1] - 2020-04-03

### Changed
-  Fixed directory and naming issues

## [0.4.0] - 2020-04-08

### Added
- Added Prometheus to init template
- Created feature addspellcommand

### Changed
-  Updated echo to the v4

## [0.3.3] - 2020-03-30

### Changed
-  Removed race modifier from test example for performance

## [0.3.2] - 2020-03-30

### Changed
-  Adjusted logger abstraction level

## [0.3.1] - 2020-03-27

### Changed
-  Fixed command initspell description

## [0.3.0] - 2020-03-26

### Added
-  Added Spell version command

## [0.2.2] - 2020-03-20

### Changed
-  Adjusted for Windows compatibility

## [0.2.1] - 2020-03-20

### Changed
-  inispell template adjusted

## [0.2.0] - 2020-03-20

### Added
-  Added initspell command

## [0.1.5] - 2020-02-25

### Changed
-  Fixed log template

## [0.1.4] - 2020-02-25

### Changed
-  Fixed log template

## [0.1.3] - 2020-02-24

### Changed
-  Improved documentation

## [0.1.2] - 2020-02-24

### Added
-  Added args check on init cmd

## [0.1.1] - 2020-02-24

### Changed
-  Improved documentation

## [0.1.0] - 2020-02-24

### Changed
-  Shared template renderer and improved multithreading isolation

## [0.0.6] - 2020-02-09

### Changed
-  Improved file generation logic

## [0.0.5] - 2020-02-08

### Added
-  init command implemented

## [0.0.4] - 2020-02-08

### Changed
-  Improved setup

## [0.0.3] - 2020-02-08

### Added
-  Added templates workflow

## [0.0.2] - 2020-01-29

### Changed
-  Improved Spell setup

## [0.0.1] - 2020-01-28

### Changed
-  Plugin management adjusts

## [0.0.0] - 2020-01-27

### Added
-  Spell Core first test tag
