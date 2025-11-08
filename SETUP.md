# BB-Lock Setup Guide

## Quick Setup (5 Minutes)

### Step 1: Clone This Repository

```bash
git clone https://github.com/gefthi/bb-lock.git
cd bb-lock
```

### Step 2: Copy WoWSims Core Files

You need to copy some files from the WoWSims project:

```bash
# Clone WoWSims (temporary, for copying files)
cd /tmp
git clone https://github.com/wowsims/wotlk.git
cd wotlk

# Copy core engine files to your bb-lock directory
cp sim/core/sim.go /path/to/bb-lock/core/
cp sim/core/spell.go /path/to/bb-lock/core/
cp sim/core/character.go /path/to/bb-lock/core/
cp sim/core/unit.go /path/to/bb-lock/core/
cp sim/core/aura.go /path/to/bb-lock/core/
cp sim/core/dot.go /path/to/bb-lock/core/
cp sim/core/stats.go /path/to/bb-lock/core/
cp sim/core/pending_action.go /path/to/bb-lock/core/
cp sim/core/environment.go /path/to/bb-lock/core/
cp sim/core/cast.go /path/to/bb-lock/core/

# Copy proto files (vendored .pb.go files)
cp sim/core/proto/*.pb.go /path/to/bb-lock/proto/
```

### Step 3: Update Import Paths

Update imports in the copied files:

```bash
cd /path/to/bb-lock

# Find and replace in all Go files
find . -name "*.go" -type f -exec sed -i 's|github.com/wowsims/wotlk/sim/core/proto|github.com/gefthi/bb-lock/proto|g' {} \;
find . -name "*.go" -type f -exec sed -i 's|github.com/wowsims/wotlk/sim/core|github.com/gefthi/bb-lock/core|g' {} \;
```

### Step 4: Install Dependencies

```bash
go mod download
```

### Step 5: Test Build

```bash
go build
./bb-lock
```

You should see:
```
BB-Lock Warlock Simulator - Phase 1
=====================================

Initializing simulation...
✅ Simulator framework ready
⏳ Next: Implement Warlock spells
```

## Next Steps

Once setup is complete, refer to:
- **docs/CONTEXT.md** - Development context
- **docs/PROJECT-PLAN.md** - Full project roadmap

## Troubleshooting

### "cannot find package"

Make sure you've copied all core files and updated import paths.

### "undefined: proto.XXX"

Make sure you've copied all `.pb.go` files from `sim/core/proto/`.

### Build errors in core files

Some core files may reference other core files you haven't copied yet. Copy the dependency files as well.

## Files Checklist

### Core Files (Essential):
- [ ] sim.go
- [ ] spell.go
- [ ] character.go
- [ ] unit.go
- [ ] aura.go
- [ ] dot.go
- [ ] stats.go
- [ ] pending_action.go
- [ ] environment.go
- [ ] cast.go

### Proto Files:
- [ ] api.pb.go
- [ ] apl.pb.go
- [ ] common.pb.go

### Import Updates:
- [ ] Updated all imports in core/ files
- [ ] Updated all imports in proto/ files
- [ ] Tested build with `go build`

## Support

Open an issue at: https://github.com/gefthi/bb-lock/issues
