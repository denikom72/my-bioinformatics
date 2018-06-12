#!/usr/bin/env python3
# Parse BLAST tabular output and report best hit per query
import sys
from collections import defaultdict
if len(sys.argv)<2:
    print('Usage: blast_parser.py blast.tsv'); sys.exit(1)
best = {}
with open(sys.argv[1]) as fh:
    for line in fh:
        parts = line.strip().split('\t')
        q,s,pid,length,mm,gap,qstart,qend,sstart,send,evalue,bits = parts[:12]
        bits = float(bits)
        if q not in best or bits > best[q]['bits']:
            best[q] = {'subject':s,'pid':pid,'bits':bits,'evalue':evalue}
for q in sorted(best):
    b = best[q]
    print('\t'.join([q,b['subject'],b['pid'],str(b['bits']),b['evalue']]))
