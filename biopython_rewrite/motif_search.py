#!/usr/bin/env python3
# Regex motif search in FASTA
import sys, re
from Bio import SeqIO
if len(sys.argv)<3:
    print('Usage: motif_search.py input.fasta motif_regex'); sys.exit(1)
infile = sys.argv[1]; motif = re.compile(sys.argv[2])
for rec in SeqIO.parse(infile,'fasta'):
    if motif.search(str(rec.seq)):
        print(rec.id)
