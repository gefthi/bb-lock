# BB-Lock - Warlock Simulator

A standalone Warlock DPS simulator for WotLK private servers, featuring custom "Mystic Enchants" and YAML-based configuration.

## Overview

BB-Lock is a subset of [WoWSims](https://github.com/wowsims/wotlk) focused exclusively on Warlock simulation with:
- ✨ **Custom Mystic Enchants** - Server-specific mechanics
- 📝 **YAML Configuration** - No web UI, pure config files
- 🎯 **APL System** - Flexible rotation programming
- ⚡ **Fast & Accurate** - Based on proven WoWSims engine

## Project Status

### Phase 1: Bootstrap 🚧 IN PROGRESS
- [x] Repository initialized
- [x] Project structure defined
- [ ] Core engine integration
- [ ] Basic spell implementation
- [ ] Hardcoded rotation

### Phase 2: YAML Configs ⏳ PLANNED
- [ ] Character configuration
- [ ] Spell database
- [ ] Config loaders

### Phase 3: APL Compiler ⏳ PLANNED
- [ ] Human-readable APL format
- [ ] YAML → Proto compiler
- [ ] Integration with sim

### Phase 4: Mystic Enchants ⏳ PLANNED
- [ ] Enchant system
- [ ] Example enchants
- [ ] Custom mechanics

## Quick Start (Phase 1)

```bash
# Clone the repository
git clone https://github.com/gefthi/bb-lock.git
cd bb-lock

# Install dependencies
go mod download

# Run (currently just prints framework status)
go run main.go
```

## Documentation

- **[CONTEXT.md](docs/CONTEXT.md)** - Comprehensive context for development
- **[PROJECT-PLAN.md](docs/PROJECT-PLAN.md)** - Full development roadmap

## Development

### Prerequisites

- Go 1.21 or higher
- Git

### Building

```bash
go build -o bb-lock main.go
./bb-lock
```

### Testing

```bash
go test ./...
```

## Architecture

```
bb-lock/
├── main.go           # Entry point
├── core/             # Simulation engine (from WoWSims)
├── warlock/          # Warlock spells & mechanics
├── apl/              # APL compiler
├── mystic/           # Custom enchants
├── loader/           # Config loaders
└── config/           # User configurations
```

## Future Usage (After Phase 2)

```bash
# Run with custom configs
./bb-lock \
  --character config/my_warlock.yml \
  --rotation config/destro_st.yml \
  --iterations 10000

# Output
DPS: 8,432.5 ± 142.3
Chaos Bolt: 42% of damage
Immolate: 26% of damage
...
```

## Contributing

This is a personal project for a specific private server, but suggestions and improvements are welcome!

## License

MIT License - Based on [WoWSims](https://github.com/wowsims/wotlk)

## Attribution

Core simulation engine adapted from WoWSims by the WotLK Classic community.

## Links

- **Repository:** https://github.com/gefthi/bb-lock
- **WoWSims:** https://github.com/wowsims/wotlk
- **Issues:** https://github.com/gefthi/bb-lock/issues
