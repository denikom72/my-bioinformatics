import subprocess
import os

def test_report_summary():
    infile = os.path.join('legacy_data','stats.tsv')
    result = subprocess.run(['python3','biopython_rewrite/report_summary.py', infile], capture_output=True, text=True)
    assert '<table>' in result.stdout

def test_count_kmers():
    infile = os.path.join('legacy_data','example.fasta')
    result = subprocess.run(['python3','biopython_rewrite/count_kmers.py', infile, '3'], capture_output=True, text=True)
    assert len(result.stdout.strip().split('\n')) > 0
