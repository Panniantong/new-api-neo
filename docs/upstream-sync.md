# Upstream release sync

This fork keeps upstream release history separate from FluxNode integration work.

## Branch roles

- `main`: points at an official `QuantumNous/new-api` release tag. It must not contain FluxNode-only changes.
- `feat-new-ui-components-clean`: integration branch. Merge `main` into this branch after every release and validate the combined tree before publishing it.
- `archive/*`: immutable recovery points created before a one-time branch repair or history rewrite.

For `v1.0.0-rc.20`, the old UI integration commits are retained as merge history, while the code tree follows the official release. Upstream commit `8b2b03d27` already incorporates the Base UI migration, theme presets, rankings dashboard, and table-toolbar work; later upstream commits replace the old mock ranking and pricing-detail components.

## Sync a new official release

Use a signed release tag, not an untagged upstream `main` commit.

```bash
git fetch upstream --tags --prune
git fetch pnt --prune

release=v1.0.0-rc.20
release_commit=$(git rev-parse "${release}^{commit}")

# Fail fast unless this is a normal fast-forward from the current fork main.
git merge-base --is-ancestor pnt/main "$release_commit"

# main remains an exact, clean copy of the official release.
git push pnt "${release_commit}:refs/heads/main"
```

If the ancestry check fails, stop. Create and verify an `archive/*` branch before deciding whether a one-time `--force-with-lease` repair is appropriate.

Then merge the release into the integration branch from an isolated worktree:

```bash
git worktree add <worktree-path> -b codex/sync-<release> pnt/feat-new-ui-components-clean
git -C <worktree-path> merge --no-ff pnt/main
```

Resolve only genuine FluxNode deltas. If an old customization has already been absorbed or redesigned upstream, keep the current upstream implementation and retain the old commit only as merge history.

## Required validation

Use the official Dockerfile so the default and classic frontends keep their
separate dependency trees. The `builder2` target contains both frontend
artifacts and the Go toolchain pinned by the release:

```bash
docker build --target builder2 -t new-api-upstream-sync-test .
docker run --rm \
  --workdir /build \
  --entrypoint /usr/local/go/bin/go \
  new-api-upstream-sync-test test ./...
docker build -t new-api-upstream-sync-image .
```

Do not install both frontend dependency trees into one shared `node_modules`
directory for validation. Their independently pinned transitive dependencies
can conflict even when both official Dockerfile stages build successfully.

Also verify:

```bash
git diff --check
git status --short
git merge-base --is-ancestor pnt/main HEAD
git merge-base --is-ancestor pnt/feat-new-ui-components-clean HEAD
```

Only commit and push after every required check passes. Production deployment remains a separate operation and should pin the matching official image tag and image digest.
