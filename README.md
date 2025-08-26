# cwgo

Character Works Library for Go(golang).

## Discovery helpers (read-only)

These helpers provide convenient, read-only discovery for auto-populating UI (e.g., Stream Deck Property Inspector). They wrap existing JSON actions and return user-facing, stable names.

```go
client := cw.NewCharacterWorks("localhost", 50)

// Layers
names, err := client.ListLayerNames(nil) // nil => default channel
if err != nil { /* handle */ }

// Fields for a given layer
fields, err := client.ListFieldNames("Scoreboard", nil)
if err != nil { /* handle */ }
// fields is map[string]string (name => optional metadata; currently empty string)

// Motions (names and IDs aligned with Play/Stop APIs)
motions, motionIDs, err := client.ListMotionNamesAndIDs(nil)
if err != nil { /* handle */ }
```

### Behavior
- channel nil: default channel
- channel set: queries specific channel (where supported by the underlying action)
- Returns stable, user-facing names; motionIDs align with `PlayMotions`/`StopMotions`.

## Underlying JSON actions (HTTP)

All requests are POSTed to the configured host (default `http://<host>:5201`) with an `application/json` body.

- list_layers
  - Request: `{ "action": "list_layers", "parent"?: string, "channel"?: string }`
  - Response: `{ "result": "ok", "layers": [ { "name": string, "path": string, "id": string, "type": string, "children"?: [...] } ] }`

- list_motions
  - Request: `{ "action": "list_motions" }`
  - Response: `{ "result": "ok", "motions": string[] }`

- list_motions_with_ids
  - Request: `{ "action": "list_motions_with_ids" }`
  - Response: `{ "result": "ok", "motions": [ { "name": string, "id": string } ] }`

- play_motions / stop_motions
  - Request: `{ "action": "play_motions"|"stop_motions", "motions": string[], "motion_ids": string[], "channel"?: string }`
  - Response: `{ "result": "ok", "warnings"?: [ { "code": string, "description": string } ] }`

- set_text
  - Request: `{ "action": "set_text", "layer": string, "layer_id": string[], "value": string, "channel"?: string }`
  - Response: `{ "result": "ok", "errors"?: [ { "code": string, "description": string } ] }`

## Examples
See `examples/basic` and `examples/grid` for end-to-end usage.