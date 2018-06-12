#!/usr/bin/env python3
# Split multi-FASTA into files
import sys, os
from Bio import SeqIO
if len(sys.argv)<2:
    print('Usage: split_fasta.py multi.fasta'); sys.exit(1)
for rec in SeqIO.parse(sys.argv[1],'fasta'):
    out = rec.id + '.fasta'
    SeqIO.write(rec, out, 'fasta')
