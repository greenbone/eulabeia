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
### Changed
- Split cmds and info messages into own module [8](https://github.com/greenbone/eulabeia/pull/8)
- Normalized topic structure to `group/aggregate/event/destination`; setting topic based on return message rather than configuration [8](https://github.com/greenbone/eulabeia/pull/8)
- Simplified block until sigterm handling [11](https://github.com/greenbone/eulabeia/pull/11)
### Fixed
### Removed
