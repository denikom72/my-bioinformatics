#!/usr/bin/env python3
# Basic SAM stats
import sys
if len(sys.argv)<2:
    print('Usage: sam_stats.py file.sam'); sys.exit(1)
mapped=0; total=0
with open(sys.argv[1]) as fh:
    for line in fh:
        if line.startswith('@'): continue
        total+=1
        cols=line.split('\t')
        if cols[2] != '*': mapped+=1
print(f"Mapped: {mapped} / {total}")
