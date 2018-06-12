#!/usr/bin/env python3
# Filter FASTA by minimum length using Biopython
import sys
from Bio import SeqIO
infile = sys.argv[1] if len(sys.argv)>1 else None
min_len = int(sys.argv[2]) if len(sys.argv)>2 else 300
if not infile:
    print('Usage: filter_by_length.py input.fasta [min_length]'); sys.exit(1)
out = infile + '.filtered.fasta'
with open(out,'w') as o:
    for rec in SeqIO.parse(infile,'fasta'):
        if len(rec.seq) >= min_len:
            SeqIO.write(rec,o,'fasta')
print('Wrote', out)
