#!/usr/bin/env python3
# Parse simple TSV variants and add consequence tag
import sys
if len(sys.argv)<2:
    print('Usage: variant_parser.py variants.tsv'); sys.exit(1)
with open(sys.argv[1]) as fh:
    for line in fh:
        chr,pos,ref,alt = line.strip().split('\t')[:4]
        cons = 'SNP' if len(ref)==len(alt) else 'INDEL'
        print('\t'.join([chr,pos,ref,alt,cons]))
