#!/usr/bin/env perl
# Parse BLAST tabular output and report best hit per query
use strict;
use warnings;
my $file = shift or die "Usage: $0 blast.tsv\n";
open my $fh, '<', $file or die $!;
my %best;
while(<$fh>){
    chomp;
    my ($q,$s,$pid,$len,$mismatch,$gap,$qstart,$qend,$sstart,$send,$evalue,$bits) = split /\t/;
    if(!exists $best{$q} || $bits > $best{$q}{bits}){
        $best{$q} = {subject=>$s, pid=>$pid, bits=>$bits, evalue=>$evalue};
    }
}
foreach my $q (sort keys %best){
    print join("\t", $q, $best{$q}{subject}, $best{$q}{pid}, $best{$q}{bits}, $best{$q}{evalue}), "\n";
}
