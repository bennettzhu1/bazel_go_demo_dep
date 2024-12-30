doing some testing with bazel and dependencies. question i'm trying to answer is if indirect deps get translated by 
go_deps.from_file 


# can a dep be present in the child but not in the parent go.mod?
```
➜  multi-module-demo go mod why github.com/google/uuid

# github.com/google/uuid
github.com/bennettzhu1/multi-module-demo
github.com/bennettzhu1/multi-module-demo/lib
github.com/google/uuid
```

uuid required by lib required by main, even though in main go.mod it's an indirect dep. conclusion is if parent go.mod consumes child go.mod, parent go.mod will have all dependencies inside child go.mod although it may be listed as indirect

# different versions of deps inside child vs parent go.mod
➜  multi-module-demo go mod graph | grep pkg/errors
github.com/bennettzhu1/multi-module-demo github.com/pkg/errors@v0.9.1
github.com/bennettzhu1/multi-module-demo/lib@v0.0.0-00010101000000-000000000000 github.com/pkg/errors@v0.9.0

➜  lib go mod graph | grep pkg/errors
github.com/bennettzhu1/multi-module-demo/lib github.com/pkg/errors@v0.9.0

Shows that lib can depend on a different version than main. also that if parent depends on child, the parent go mod graph will be a superset of the child go mod graph. this assumes the parent consumes all children go.mods

in cire-obelisk, all are consumed except protos/contrib/mk-include


