import subprocess
import os

def test_gc_content():
    infile = os.path.join('legacy_data','example.fasta')
    result = subprocess.run(['python3','biopython_rewrite/gc_content.py', infile], capture_output=True, text=True)
    assert 'seq1' in result.stdout
    assert 'seq2' in result.stdout

def test_motif_search():
    infile = os.path.join('legacy_data','example.fasta')
    result = subprocess.run(['python3','biopython_rewrite/motif_search.py', infile, 'ATG'], capture_output=True, text=True)
    assert 'seq1' in result.stdout or 'seq2' in result.stdout
