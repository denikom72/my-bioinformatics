#!/usr/bin/env python3
# Toy fetch: from best-hits TSV produce a FASTA with placeholder seqs
import sys
if len(sys.argv)<2:
    print('Usage: blast_fetch_ids.py best_hits.tsv'); sys.exit(1)
with open(sys.argv[1]) as fh:
    for line in fh:
        q,s,*rest = line.strip().split('\t')
        print(f">{s}\nNNNNNNNNNN\n")
