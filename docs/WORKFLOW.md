Mini Workflow (Demonstration)
=============================
This is a toy workflow emulating a small NGS/sequence analysis pipeline:

1. Input FASTA sequences in legacy_data/
2. GC content analysis: biopython_rewrite/gc_content.py
3. Motif search: biopython_rewrite/motif_search.py
4. Variant parsing: biopython_rewrite/variant_parser.py
5. Generate HTML summary report: biopython_rewrite/report_summary.py

Example run:
$ python3 biopython_rewrite/gc_content.py legacy_data/example.fasta
$ python3 biopython_rewrite/motif_search.py legacy_data/example.fasta 'ATG'
$ python3 biopython_rewrite/variant_parser.py legacy_data/variants.tsv
$ python3 biopython_rewrite/report_summary.py legacy_data/stats.tsv

Unit tests are provided in tests/ to verify functionality.
