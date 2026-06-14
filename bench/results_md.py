#!/usr/bin/env python3
"""Convert the `go test -bench` output in results.txt into the markdown summary.

Usage: results_md.py <src results.txt> <dst .md>. run_bench.sh passes an
architecture-specific destination (results_<goarch>.md)."""
import re
import sys

src, dst = sys.argv[1], sys.argv[2]
with open(src) as f:
    lines = f.readlines()

meta = [l[1:].strip() for l in lines if l.startswith("#") and l[1:].strip()]
# Drop the descriptive title line; keep the "generated ..." / "go version ..." facts.
meta = [m for m in meta if m.lower().startswith(("generated", "go "))]

# Parse sections: header "bench/N/  (input.json: NNNN bytes)" then Benchmark rows.
sections = []
cur = None
hdr = re.compile(r"^(bench/\S+?)/?\s+\(input\.json:\s*(\d+)\s*bytes\)")
row = re.compile(
    r"^Benchmark(\S+?)-\d+\s+\d+\s+([\d.]+)\s*ns/op"
    r"(?:\s+([\d.]+)\s*MB/s)?"
    r"(?:\s+(\d+)\s*B/op)?"
    r"(?:\s+(\d+)\s*allocs/op)?"
)
for l in lines:
    m = hdr.match(l.strip())
    if m:
        cur = {"name": m.group(1), "bytes": m.group(2), "rows": []}
        sections.append(cur)
        continue
    m = row.match(l.strip())
    if m and cur is not None:
        cur["rows"].append({
            "decoder": m.group(1),
            "ns": float(m.group(2)),
            "mbs": m.group(3),
            "bytes": m.group(4),
            "allocs": m.group(5),
        })

out = ["# JSON Deserialization Benchmarks", ""]
out += [f"- {m}" for m in meta] + [""]
out.append("Lower ns/op is better; throughput (MB/s) and allocations are reported "
           "by `-benchmem`. **Speedup** is relative to the `encoding/json` (Stdlib) baseline.")
out.append("")

for s in sections:
    out.append(f"## {s['name']} — {s['bytes']} byte input")
    out.append("")
    base = next((r["ns"] for r in s["rows"] if r["decoder"].lower() == "stdlib"), None)
    out.append("| Decoder | ns/op | Throughput | B/op | allocs/op | Speedup |")
    out.append("|---|--:|--:|--:|--:|--:|")
    for r in sorted(s["rows"], key=lambda r: r["ns"]):
        mbs = f"{r['mbs']} MB/s" if r["mbs"] else "—"
        b = r["bytes"] if r["bytes"] is not None else "—"
        a = r["allocs"] if r["allocs"] is not None else "—"
        speed = f"{base / r['ns']:.1f}×" if base else "—"
        out.append(f"| {r['decoder']} | {r['ns']:.0f} | {mbs} | {b} | {a} | {speed} |")
    out.append("")

with open(dst, "w") as f:
    f.write("\n".join(out).rstrip() + "\n")
print(f"wrote {dst}")
