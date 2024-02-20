package runner

import (
	"context"
	"crypto/rand"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/experimental/sysfs"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

func Run(name string, wasm []byte) {
	ctx := context.Background()

	rtCfg := wazero.NewRuntimeConfig()
	uc, err := os.UserCacheDir()
	if err == nil {
		cache, err := wazero.NewCompilationCacheWithDir(filepath.Join(uc, "com.github.wasilibs"))
		if err == nil {
			rtCfg = rtCfg.WithCompilationCache(cache)
		}
	}

	rt := wazero.NewRuntimeWithConfig(ctx, rtCfg)

	wasi_snapshot_preview1.MustInstantiate(ctx, rt)

	args := []string{name}
	args = append(args, os.Args[1:]...)

	root := sysfs.DirFS(".")

	cfg := wazero.NewModuleConfig().
		WithSysNanosleep().
		WithSysNanotime().
		WithSysWalltime().
		WithStderr(os.Stderr).
		WithStdout(os.Stdout).
		WithStdin(os.Stdin).
		WithRandSource(rand.Reader).
		WithArgs(args...).
		WithFSConfig(wazero.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(root, "/"))
	for _, env := range os.Environ() {
		k, v, _ := strings.Cut(env, "=")
		cfg = cfg.WithEnv(k, v)
	}

	_, err = rt.InstantiateWithConfig(ctx, wasm, cfg)
	if err != nil {
		if sErr, ok := err.(*sys.ExitError); ok {
			os.Exit(int(sErr.ExitCode()))
		}
		log.Fatal(err)
	}
}
