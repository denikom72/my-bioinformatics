# Migrated from BioPerl to Biopython by A. Smith (USA) 2020
#!/usr/bin/env python3
# Biopython rewrite: compute GC content per sequence
from Bio import SeqIO
import sys
if len(sys.argv) < 2:
    print('Usage: gc_content.py input.fasta'); sys.exit(1)
for rec in SeqIO.parse(sys.argv[1], 'fasta'):
    seq = str(rec.seq)
    gc = (seq.upper().count('G') + seq.upper().count('C')) / len(seq) * 100 if len(seq)>0 else 0
    print(f"{rec.id}\t{gc:.2f}")
