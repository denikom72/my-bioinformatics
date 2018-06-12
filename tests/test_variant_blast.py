import subprocess
import os

def test_variant_parser():
    infile = os.path.join('legacy_data','variants.tsv')
    result = subprocess.run(['python3','biopython_rewrite/variant_parser.py', infile], capture_output=True, text=True)
    assert 'SNP' in result.stdout or 'INDEL' in result.stdout

def test_blast_parser():
    infile = os.path.join('legacy_data','blast_hits.tsv')
    result = subprocess.run(['python3','biopython_rewrite/blast_parser.py', infile], capture_output=True, text=True)
    assert 'subjA' in result.stdout or 'subjB' in result.stdout
