rule all:
    input:
        "results/gc_content.txt",
        "results/count_kmers.txt",
        "results/motif_search.txt",
        "results/variant_parser.txt",
        "results/report_summary.html"

rule gc_content:
    input:
        fasta="legacy_data/example.fasta"
    output:
        "results/gc_content.txt"
    shell:
        "python3 biopython_rewrite/gc_content.py {input.fasta} > {output}"

rule count_kmers:
    input:
        fasta="legacy_data/example.fasta"
    output:
        "results/count_kmers.txt"
    shell:
        "python3 biopython_rewrite/count_kmers.py {input.fasta} 3 > {output}"

rule motif_search:
    input:
        fasta="legacy_data/example.fasta"
    output:
        "results/motif_search.txt"
    shell:
        "python3 biopython_rewrite/motif_search.py {input.fasta} 'ATG' > {output}"

rule variant_parser:
    input:
        variants="legacy_data/variants.tsv"
    output:
        "results/variant_parser.txt"
    shell:
        "python3 biopython_rewrite/variant_parser.py {input.variants} > {output}"

rule report_summary:
    input:
        stats="legacy_data/stats.tsv"
    output:
        "results/report_summary.html"
    shell:
        "python3 biopython_rewrite/report_summary.py {input.stats} > {output}"
