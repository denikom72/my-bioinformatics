#!/usr/bin/env perl
# Given BLAST best-hits, fetch sequence IDs into a FASTA (toy, requires entrez utils normally)
use strict; use warnings;
my $in = shift or die "Usage: $0 best_hits.tsv\n";
open my $fh,'<',$in or die $!;
while(<$fh>){ chomp; my ($qid,$sid,undef)=split/\t/; print ">${sid}\nNNNNNNNNNN\n"; }
