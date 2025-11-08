# Protocol Buffer Files (Vendored)

This directory contains vendored `.pb.go` files from WoWSims.

## Strategy: Option 2 (Vendored Protobufs)

We're using **vendored** protobuf files - meaning we copy the already-generated `.pb.go` files, NOT the source `.proto` files.

### Benefits:
- ✅ No `protoc` compiler needed
- ✅ No build step required
- ✅ Just `go run main.go` works
- ✅ Simpler setup for contributors

## Files to Copy

From `https://github.com/wowsims/wotlk/tree/master/sim/core/proto`:

- [ ] `api.pb.go` - API definitions
- [ ] `apl.pb.go` - APL system definitions (CRITICAL for rotations)
- [ ] `common.pb.go` - Common types
- [ ] `warlock.pb.go` - Warlock-specific (if exists)

## Import Path Updates

After copying, update the package declaration in each file:

```go
// OLD (in WoWSims):
package proto

// NEW (in bb-lock):
package proto
```

Update imports in files that reference proto:

```go
// OLD:
import "github.com/wowsims/wotlk/sim/core/proto"

// NEW:
import "github.com/gefthi/bb-lock/proto"
```

## Instructions

```bash
# 1. Clone WoWSims (if not already)
git clone https://github.com/wowsims/wotlk.git

# 2. Copy proto files
cp wotlk/sim/core/proto/*.pb.go /path/to/bb-lock/proto/

# 3. Update imports in other files (core/, warlock/, etc.)
# Find and replace:
#   github.com/wowsims/wotlk/sim/core/proto
# With:
#   github.com/gefthi/bb-lock/proto
```

## Dependencies

Only runtime dependency needed:

```go
// go.mod
require (
    google.golang.org/protobuf v1.31.0
)
```

No `protoc` or `protoc-gen-go` required!
