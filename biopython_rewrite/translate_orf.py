#!/usr/bin/env python3
# Translate ORFs from FASTA (simple)
from Bio import SeqIO
from Bio.Seq import Seq
import sys
if len(sys.argv)<2:
    print('Usage: translate_orf.py input.fasta'); sys.exit(1)
for rec in SeqIO.parse(sys.argv[1],'fasta'):
    seq = rec.seq
    for frame in range(3):
        pep = seq[frame:].translate(to_stop=False)
        print(f">{rec.id}_frame{frame+1}\n{pep}")
