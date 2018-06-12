#!/usr/bin/env perl
# Translate ORFs from a FASTA file (simple approach)
use strict;
use warnings;
use Bio::SeqIO;
use Bio::Tools::CodonTable;
my $in = shift or die "Usage: $0 input.fasta\n";
my $seqio = Bio::SeqIO->new(-file=>$in, -format=>'fasta');
my $table = Bio::Tools::CodonTable->new(-id=>1);
while(my $r = $seqio->next_seq){
    my $seq = $r->seq;
    for(my $i=0;$i<3;$i++){
        my $pep = $table->translate($seq, $i, undef, 1);
        print ">", $r->display_id, "_frame", $i+1, "\n", $pep, "\n";
    }
}
