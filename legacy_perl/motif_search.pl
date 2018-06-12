#!/usr/bin/env perl
# Search for regex motif in FASTA sequences
use strict; use warnings; use Bio::SeqIO;
my ($in,$motif) = @ARGV;
die "Usage: $0 input.fasta motif_regex\n" unless $in && $motif;
my $seqio = Bio::SeqIO->new(-file=>$in,-format=>'fasta');
while(my $r=$seqio->next_seq){
    if($r->seq =~ /$motif/){
        print $r->display_id, "\n";
    }
}
