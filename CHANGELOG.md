# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [21.10] (unreleased)

### Added
- Add mqtt, example, handler for create, modify and get target [4](https://github.com/greenbone/eulabeia/pull/4)
- Add scan, sensor aggregate; extend sensor to register and deregister itself [5](https://github.com/greenbone/eulabeia/pull/5)
- Add scanner logic, scheduler and openvas module [10](https://github.com/greenbone/eulabeia/pull/10)
- Smoke Tests, tests to verify if the happy path is functioning [14](https://github.com/greenbone/eulabeia/pull/14)
- Unittests for openvas module [24](https://github.com/greenbone/eulabeia/pull/24)
- Adapt target credentials to allow different types [27](https://github.com/greenbone/eulabeia/pull/27)
- Adapt sensor so it is able to start a scan now [30](https://github.com/greenbone/eulabeia/pull/30)
- Add redis and vt loading [32](https://github.com/greenbone/eulabeia/pull/32)
- Possibility to preprocess messages [31](https://github.com/greenbone/eulabeia/pull/31)
- Possibility to send one start.scan event containing all the data [31](https://github.com/greenbone/eulabeia/pull/31)
- Add result model [36](https://github.com/greenbone/eulabeia/pull/36)
- Add model for vts [37](https://github.com/greenbone/eulabeia/pull/37)
- Tests for sensor, QueueList and handler(sensor) [39](https://github.com/greenbone/eulabeia/pull/39)
- Handling interrupted scans [39](https://github.com/greenbone/eulabeia/pull/39)
- Possibility to set VT preferences and select VT Groups [43](https://github.com/greenbone/eulabeia/pull/43)
- Feedservice to handle Get VT [54](https://github.com/greenbone/eulabeia/pull/54)
- Reolving VT Groups into OIDs [59](https://github.com/greenbone/eulabeia/pull/59)

### Changed
- Split cmds and info messages into own module [8](https://github.com/greenbone/eulabeia/pull/8)
- Normalized topic structure to `group/aggregate/event/destination`; setting topic based on return message rather than configuration [8](https://github.com/greenbone/eulabeia/pull/8)
- Simplified block until sigterm handling [11](https://github.com/greenbone/eulabeia/pull/11)
- When sensor is closing stop all scans simultaniously [52](https://github.com/greenbone/eulabeia/pull/52)
### Fixed
- Fix handling for finished scans [46](https://github.com/greenbone/eulabeia/pull/46)
### Removed
