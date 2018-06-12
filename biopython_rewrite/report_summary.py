#!/usr/bin/env python3
# Create tiny HTML summary from TSV stats
import sys
if len(sys.argv)<2:
    print('Usage: report_summary.py stats.tsv'); sys.exit(1)
print('<html><body><table>')
with open(sys.argv[1]) as fh:
    for line in fh:
        cols = line.strip().split('\t')
        print('<tr>' + ''.join(f'<td>{c}</td>' for c in cols) + '</tr>')
print('</table></body></html>')
