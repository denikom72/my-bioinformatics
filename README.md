Note: These scripts are provided solely as reference implementations for my projects. All sensitive or proprietary data have been removed for legal and confidentiality reasons."

Migration demo repo - BioPerl -> Biopython + BioJava modules
================================================================
long-term project history (2016-2019) in an international team:
- Initial Perl scripts to Python migration with Biopython 
- BioJava modules for structure parsing 
- Regression tests and example data 


Structure:
- legacy_perl/: original BioPerl scripts (archived)
- biopython_rewrite/: rewritten scripts using Biopython
- biojava_module/: Java utilities for structure parsing / GC analysis
- legacy_data/: example FASTA/TSV files for testing
- tests/: pytest scripts to validate rewrites
- docs/: migration notes, usage and history

###package manager###

#Perl
cpanm --installdeps .

#For the new rewritten Python Scripts
python3 -m pip install -r requirements.txt

#Java
mvn clean install


# Example how to run scripts with 4 cores
snakemake -j 4
