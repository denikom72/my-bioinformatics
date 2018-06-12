#!/usr/bin/env perl
# Parse simple VCF-like tab delimited variants and annotate with consequence (toy example)
use strict;
use warnings;
my $in = shift or die "Usage: $0 variants.tsv\n";
open my $fh, '<', $in or die $!;
while(<$fh>){
    chomp; my ($chr,$pos,$ref,$alt) = split /\t/;
    my $cons = (length($ref) == length($alt)) ? 'SNP' : 'INDEL';
    print join('\t', $chr,$pos,$ref,$alt,$cons), "\n";
}
