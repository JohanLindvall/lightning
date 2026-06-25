#!/usr/bin/env python3
"""Render the `go test -bench ./...` output from the main module into a markdown
summary, one table per top-level benchmark (EscapeString, EscapeStringInto, …).

Usage: pkg_results_md.py <src pkg_results.txt> <dst .md> [cpu-name]. pkg_bench.sh
passes an architecture-specific destination (pkg_results_<goarch>.md) and, as an
optional third argument, a CPU brand string detected via github.com/klauspost/cpuid
(Go's own `cpu:` line is blank — "unknown" — on arm64)."""
import re
import sys

src, dst = sys.argv[1], sys.argv[2]
cpu_override = sys.argv[3].strip() if len(sys.argv) > 3 else ""
with open(src) as f:
    lines = f.readlines()

meta = [l[1:].strip() for l in lines if l.startswith("#") and l[1:].strip()]
meta = [m for m in meta if m.lower().startswith(("generated", "go "))]

# Surface the processor model (Go prints a `cpu:` line) and the core count (the
# GOMAXPROCS suffix on benchmark names, e.g. `...-16`) at the top of the summary.
cpu, cores = "", ""
for l in lines:
    if not cpu and l.startswith("cpu:"):
        cpu = l.split(":", 1)[1].strip()
    if not cores:
        m = re.match(r"Benchmark\S+?-(\d+)\s", l)
        if m:
            cores = m.group(1)
# Prefer the cpuid-detected brand string: Go's `cpu:` line is "unknown" on arm64,
# where cpuid still reports a real model (e.g. "Apple M2").
if cpu_override and (not cpu or cpu.lower() == "unknown"):
    cpu = cpu_override
if cpu or cores:
    label = "cpu: " + (cpu or "unknown")
    if cores:
        label += f" ({cores} cores)"
    meta.append(label)

# `go test ./...` prints a `pkg: <import path>` header before each package's rows.
pkg = re.compile(r"^pkg:\s+(\S+)")
row = re.compile(
    r"^Benchmark(\S+?)-\d+\s+\d+\s+([\d.]+)\s*ns/op"
    r"(?:\s+([\d.]+)\s*MB/s)?"
    r"(?:\s+(\d+)\s*B/op)?"
    r"(?:\s+(\d+)\s*allocs/op)?"
)

# One group per top-level benchmark name (the part before the first '/'); each
# sub-benchmark becomes a row keyed by the remaining "case" path. Groups are kept
# in first-seen order so the markdown mirrors the run order.
groups = {}
order = []
curpkg = None
for l in lines:
    m = pkg.match(l.strip())
    if m:
        curpkg = m.group(1)
        continue
    m = row.match(l.strip())
    if not m:
        continue
    full = m.group(1)
    group, _, case = full.partition("/")
    if group not in groups:
        groups[group] = {"pkg": curpkg, "rows": []}
        order.append(group)
    groups[group]["rows"].append({
        "case": case or "—",
        "ns": float(m.group(2)),
        "mbs": m.group(3),
        "bytes": m.group(4),
        "allocs": m.group(5),
    })

out = ["# lightning main-module benchmarks", ""]
out += [f"- {m}" for m in meta] + [""]
out.append("The Benchmark* functions in the lightning module itself "
           "(`pkg/json`, `pkg/unstable`, …), as opposed to the competitor-comparison "
           "suite in `bench/` (see `results_<arch>.md`). One table per benchmark; "
           "lower ns/op is better; throughput (MB/s) and allocations are reported by "
           "`-benchmem`.")
out.append("")

for group in order:
    g = groups[group]
    out.append(f"## {group}")
    out.append("")
    if g["pkg"]:
        out.append(f"`{g['pkg']}`")
        out.append("")
    out.append("| Case | ns/op | Throughput | B/op | allocs/op |")
    out.append("|---|--:|--:|--:|--:|")
    for r in g["rows"]:
        mbs = f"{r['mbs']} MB/s" if r["mbs"] else "—"
        b = r["bytes"] if r["bytes"] is not None else "—"
        a = r["allocs"] if r["allocs"] is not None else "—"
        out.append(f"| {r['case']} | {r['ns']:.1f} | {mbs} | {b} | {a} |")
    out.append("")

with open(dst, "w") as f:
    f.write("\n".join(out).rstrip() + "\n")
print(f"wrote {dst}")
