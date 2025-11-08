# BB-Lock Implementation Notes

## Directory Structure Corrections

### Stats Directory
- **Original wowsim**: `stats.go` was located in the `stats/` directory at the root level
- **BB-Lock**: Created `stats/` directory within `core/` directory
- **Decision**: Need to verify if this relocation is necessary. If not needed, will revert to original structure.

### Proto Files
- **Original wowsim**: Had `*_pb.go` proto compiled files
- **BB-Lock**: Only `.proto` source files were copied to the `proto/` directory
- **Note**: The `*_pb.go` files need to be generated using protoc compiler

## Code Modifications

### Time Import
- **Issue**: In several files, the `time` package import was commented out due to being unused
- **Location**: Check imports in spell implementation files
- **Action Required**: Uncomment `time` import when duration calculations are actually used
- **Example Files**: 
  - `warlock/immolate.go`
  - `warlock/incinerate.go`
  - `warlock/chaos_bolt.go`

### Import Paths
- **Issue**: Some files contain imports that reference the original wowsim git repository
- **Current State**: Left as-is for now
- **Future Decision**: Decide whether to update these to point to `github.com/gefthi/bb-lock`
- **Consideration**: May want to keep original references if maintaining compatibility, or update if forking completely

## Project Structure

### Class Directories
- **Original wowsim**: Has separate directories for each class (e.g., `/warlock`, `/mage`, `/warrior`)
- **BB-Lock**: Following the same structure
- **Current Implementation**: Created `/warlock` directory with spell implementations

### Root Documentation
- **Location**: Root directory of the project
- **Files**: Various `.md` files containing project information
- **Content**: Project overview, setup instructions, development guidelines

## Warlock Spell Implementation

### Implemented Spells (WotLK Version)
1. **Immolate** (`warlock/immolate.go`)
   - Spell ID: 47811 (Rank 11)
   - Type: Fire DoT with initial direct damage
   - Cast Time: 2 seconds
   - Duration: 15 seconds (5 ticks, 3 seconds each)
   - Direct Damage: 505-621
   - DoT Damage: 785 total (157 per tick)

2. **Incinerate** (`warlock/incinerate.go`)
   - Spell ID: 47838 (Rank 4)
   - Type: Direct Fire damage
   - Cast Time: 2.5 seconds
   - Base Damage: 583-676
   - Special: 25% increased damage if target has Immolate

3. **Chaos Bolt** (`warlock/chaos_bolt.go`)
   - Spell ID: 59172 (Rank 4)
   - Type: Direct Fire/Chaos damage
   - Cast Time: 2.5 seconds
   - Cooldown: 12 seconds
   - Base Damage: 1429-1813
   - Special: Always crits, cannot be resisted

## Build Notes

### Current Build Issues
- Time package imports commented out in spell files to allow compilation
- Need to uncomment when actually using time-based calculations

### Import References
- Several files reference original wowsim repository
- Decision pending on whether to update to bb-lock repository paths

## Next Steps

1. **Verify stats/ directory location**
   - Check if `core/stats/` is correct or if it should be at root level
   - Revert if necessary

2. **Generate proto files**
   - Run protoc compiler to generate `*_pb.go` files from `.proto` sources
   - Add to build process

3. **Review and update imports**
   - Audit all files for wowsim repository references
   - Update to bb-lock references if appropriate

4. **Uncomment time imports**
   - When duration/timing calculations are implemented
   - Update spell implementations to use time package

5. **Test spell implementations**
   - Verify spell mechanics
   - Check damage calculations
   - Validate spell interactions (e.g., Incinerate + Immolate bonus)

## References
- Original Project: wowsims GitHub repository
- Current Project: github.com/gefthi/bb-lock
- Version: WotLK (Wrath of the Lich King)
