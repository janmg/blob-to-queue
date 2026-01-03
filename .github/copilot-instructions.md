# Copilot instructions for blob-to-queue

Purpose: help an AI coding agent quickly understand, change and extend this project.

- Big picture:
  - Reads Azure blob PT1H.json flow logs, flattens entries and streams events to outputs.
  - Core components: `main` ([blob-to-queue.go](blob-to-queue.go)), configuration (`common.ConfigHandler()` in [common/config.go](common/config.go)), input workers ([input/blob.go](input/blob.go)), event format ([format/flatevent.go](format/flatevent.go)) and output workers (files under [output/](output/)).
  - Data flow: list blobs -> compare to registry/timestamp -> HTTP range read (partial or full) -> parse into `format.Flatevent` -> push into channel queue -> output worker consumes and sends to destination.

- Key patterns & conventions (project-specific):
  - Configuration is read with `spf13/viper` from `blob-to-queue.yaml` and can be live-reloaded via `fsnotify` (see `common.ConfigHandler`).
  - Resume strategies: `resumepolicy` = `timestamp` or `registry`. `startpolicy` = `start_over` or `start_fresh`.
  - Input concurrency: single `input.Blobworker(queue)` goroutine produces into a buffered channel `queue` sized by `config.Qsize`. Backpressure handled by `Qwatermark` checks in `doLoop` and a signal goroutine inside `read()`.
  - Outputs: `main` iterates `config.Output` and starts worker goroutines via a switch. Currently only some outputs are activated (e.g. Elasticsearch). To add an output: implement `output.<Name>Worker(queue chan format.Flatevent)` and add the case in the switch in [blob-to-queue.go](blob-to-queue.go). Fan-out to multiple consumers is not implemented — the current pattern assumes one worker per output case that reads from the single channel.
  - File registry/timestamp: registry stored in JSON (example path default `./registry.dat`), timestamp stored in `timestamp.json` by default. See `input/blob.go` for read/write helpers.

- Build / run / debug:
  - Quick run: `go run blob-to-queue.go` (see [README.md](README.md)).
  - The program prints a simple start banner and uses `fmt.Println` throughout — check stdout for progress and warnings.
  - Config changes: edit `blob-to-queue.yaml`; if `fsnotify` is enabled in the config, `ConfigHandler` will reload on save.

- When editing config or adding fields:
  - Add `mapstructure` tags to new fields in `common.Config` to ensure viper unmarshals them.
  - If you add complex nested config, update `ConfigHandler()` defaults accordingly.

- Example edits an AI agent commonly needs to do (with pointers):
  - Add a new output backend: implement worker in `output/`, add a `case` in the switch in [blob-to-queue.go](blob-to-queue.go), and add config options in `common.Config` + defaults in `ConfigHandler()` (see how `elasticsearch` is wired).
  - Change resume logic: modify `input.Blobworker` and `doLoop` in [input/blob.go](input/blob.go) — registry vs timestamp behavior is centralized there.

- Files to inspect for change examples:
  - [blob-to-queue.go](blob-to-queue.go) — program entry, worker startup, output switch
  - [common/config.go](common/config.go) — viper config, defaults, watch handler
  - [input/blob.go](input/blob.go) — blob listing, partial/full reads, registry/timestamp handling
  - [format/flatevent.go](format/flatevent.go) — event struct and flattening logic
  - [output/stdout.go](output/stdout.go) — example output worker pattern

- Testing & assumptions discovered in code:
  - There are no automated tests in the repository; verify changes by running locally against a dev Azure storage account or with mocked inputs.
  - Error handling uses `common.Error(err)` helper across the codebase — follow that pattern for consistency.

If anything here is unclear or you want me to expand a specific example (e.g. add a new output worker skeleton, or provide a minimal test harness), tell me which part to expand.
