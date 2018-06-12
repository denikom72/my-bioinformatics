Note: These scripts are provided solely as reference implementations for my projects. All sensitive or proprietary data have been removed for legal and confidentiality reasons."

MIGRATION NOTES
---------------
migration of ~12 BioPerl utilities into Biopython equivalents.

Approach used by teams in industry/research:
- Inventory scripts and categorize by function (I/O, parsing, analysis, reporting)
- Prioritize critical scripts with tests and sample data
- Rewrite logic in Python, using Biopython for sequence I/O and standard libraries for parsing
- Validate outputs using regression tests against legacy outputs
- Keep legacy scripts archived and provide wrappers during cutover


- GC content, motif search, k-mer counting, BLAST best-hit parsing, basic SAM stats, VCF-like parsing
- A small Java module set for PDB/structure-related tasks
