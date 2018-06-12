#!/usr/bin/env perl
# Split multi-FASTA into individual files
use strict; use warnings; use Bio::SeqIO;
my $in = shift or die "Usage: $0 multi.fasta\n";
my $seqio = Bio::SeqIO->new(-file=>$in,-format=>'fasta');
while(my $r=$seqio->next_seq){
    my $id = $r->display_id;
    my $out = $id . '.fasta';
    my $o = Bio::SeqIO->new(-file=>">$out", -format=>'fasta');
    $o->write_seq($r);
}
