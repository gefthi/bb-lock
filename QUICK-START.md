# 🚀 BB-Lock Phase 1 Bootstrap - Quick Start

## What You Have

You now have **3 downloadable files** ready to go:

1. **CONTEXT.md** - The master context document for efficient future sessions
2. **warlock-sim-project-plan.md** - Complete development roadmap
3. **bb-lock-phase1-bootstrap.zip** - Phase 1 starter files

---

## ⚡ Next Steps (5 Minutes Setup)

### Step 1: Download & Extract

Download the ZIP file and extract it to your repository:

```bash
cd /path/to/your/repo
unzip bb-lock-phase1-bootstrap.zip
mv bb-lock-phase1/* .
rm -rf bb-lock-phase1
```

Your structure will look like:

```
bb-lock/
├── main.go
├── go.mod
├── README.md
├── SETUP.md
├── .gitignore
├── docs/
│   ├── CONTEXT.md           ← IMPORTANT: Keep this updated
│   └── PROJECT-PLAN.md
├── core/                    ← Copy WoWSims files here
├── proto/                   ← Copy WoWSims proto files here
├── warlock/
│   └── warlock.go
└── config/
```

### Step 2: Copy WoWSims Files

Follow **SETUP.md** for detailed instructions, or quick version:

```bash
# Clone WoWSims temporarily
cd /tmp
git clone https://github.com/wowsims/wotlk.git

# Copy core files
cd wotlk/sim/core
cp sim.go spell.go character.go unit.go aura.go dot.go stats.go \
   pending_action.go environment.go cast.go \
   /path/to/bb-lock/core/

# Copy proto files
cp proto/*.pb.go /path/to/bb-lock/proto/

# Update import paths (see SETUP.md)
cd /path/to/bb-lock
find . -name "*.go" -exec sed -i 's|github.com/wowsims/wotlk/sim/core|github.com/gefthi/bb-lock/core|g' {} \;
```

### Step 3: Commit to GitHub

```bash
cd /path/to/bb-lock
git add .
git commit -m "Phase 1 bootstrap - project structure"
git push origin main
```

### Step 4: Test Build

```bash
go mod download
go build
./bb-lock
```

Expected output:
```
BB-Lock Warlock Simulator - Phase 1
=====================================

Initializing simulation...
✅ Simulator framework ready
⏳ Next: Implement Warlock spells
```

---

## 🎯 For Next Session with Claude

When you're ready to continue development, start a new conversation with:

```
Context: https://github.com/gefthi/bb-lock/blob/main/docs/CONTEXT.md

Task: [what you want to implement next]

Examples:
- "Implement Chaos Bolt spell"
- "Add Immolate with DOT functionality"
- "Create the simulation loop"
```

**Why this works:**
- ✅ Claude fetches CONTEXT.md (~2K tokens)
- ✅ Understands full project structure
- ✅ Knows all conventions & patterns
- ✅ Only fetches files needed for your task
- ✅ Generates diffs/patches (not full rewrites)
- ✅ **Saves 70% tokens vs explaining from scratch**

---

## 📊 Token Usage Summary

**This Session:**
- Initial discussion: ~15K tokens
- Project planning: ~10K tokens
- Bootstrap files: ~5K tokens
- **Total: ~30K tokens** (84% remaining for future sessions!)

**Future Sessions:**
- Context loading: ~2K tokens
- Specific task: ~3-5K tokens
- **Per feature: ~5-7K tokens**

With 190K limit per session, you can implement **20-30 features** before needing a new conversation!

---

## 📝 Important Files

### Keep These Updated:

**docs/CONTEXT.md** - Update when you:
- Complete a phase
- Add new spells
- Change project structure
- Make architectural decisions

**README.md** - Update when you:
- Complete features
- Change usage instructions
- Add new capabilities

---

## 🎨 What's Included

### Ready to Use:
- ✅ Go module setup (go.mod)
- ✅ Main entry point (main.go)
- ✅ Project structure
- ✅ Documentation (CONTEXT.md, PROJECT-PLAN.md, SETUP.md)
- ✅ Warlock package stub
- ✅ Directory structure
- ✅ .gitignore

### TODO (Next Steps):
- [ ] Copy WoWSims core files
- [ ] Copy WoWSims proto files
- [ ] Update import paths
- [ ] Implement first spell (Chaos Bolt)
- [ ] Implement rotation logic
- [ ] Test simulation

---

## 🔗 Quick Links

- **Your Repo:** https://github.com/gefthi/bb-lock
- **WoWSims Source:** https://github.com/wowsims/wotlk
- **Context Doc:** https://github.com/gefthi/bb-lock/blob/main/docs/CONTEXT.md

---

## 🎯 Recommended Next Actions

1. **Today:**
   - Extract ZIP
   - Copy WoWSims files (following SETUP.md)
   - Commit to GitHub
   - Test build

2. **Tomorrow (New Session):**
   ```
   "Context: https://github.com/gefthi/bb-lock/blob/main/docs/CONTEXT.md
   
   Task: Implement Chaos Bolt spell in warlock/chaos_bolt.go
   following the spell implementation pattern from WoWSims"
   ```

3. **This Week:**
   - Complete Phase 1 (3-5 basic spells)
   - Get first DPS output
   - Celebrate! 🎉

---

## 💡 Pro Tips

1. **Always reference CONTEXT.md** in new sessions
2. **Ask for diffs/patches** for updates (not full files)
3. **Batch related questions** to save tokens
4. **Update CONTEXT.md** as project evolves
5. **Use GitHub URLs** to reference files

---

## 🚨 If You Get Stuck

1. Check SETUP.md for detailed instructions
2. Ensure all WoWSims files are copied
3. Verify import paths are updated
4. Try `go mod tidy`
5. Open an issue on GitHub

---

**Ready to build an awesome Warlock simulator!** 🔮✨

Start by extracting the ZIP, following SETUP.md, and pushing to GitHub. Then come back with the CONTEXT.md URL for the next session!
