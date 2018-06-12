#!/usr/bin/env python3
# Simple merge of two FASTQ files by interleaving lines
import sys
if len(sys.argv)!=4:
    print('Usage: merge_fastq.py r1.fastq r2.fastq out.fastq'); sys.exit(1)
r1, r2, out = sys.argv[1:4]
with open(r1) as a, open(r2) as b, open(out,'w') as o:
    while True:
        a_block = [a.readline() for _ in range(4)]
        b_block = [b.readline() for _ in range(4)]
        if not a_block[0] or not b_block[0]:
            break
        o.writelines(a_block); o.writelines(b_block)
print('Merged to', out)
