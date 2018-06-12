#!/usr/bin/env python3
# Count k-mers in FASTA sequences
import sys
from collections import Counter
from Bio import SeqIO
infile = sys.argv[1] if len(sys.argv)>1 else None
k = int(sys.argv[2]) if len(sys.argv)>2 else 5
if not infile:
    print('Usage: count_kmers.py input.fasta [k]'); sys.exit(1)
cnt = Counter()
for rec in SeqIO.parse(infile,'fasta'):
    s = str(rec.seq)
    for i in range(len(s)-k+1):
        cnt[s[i:i+k]] += 1
for kmer, c in cnt.items():
    print(f"{kmer}\t{c}")
