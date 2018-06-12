#!/usr/bin/env perl
# Filter FASTA by minimum length
use strict;
use warnings;
use Bio::SeqIO;
my ($in, $min) = @ARGV;
$min ||= 300;
die "Usage: $0 input.fasta [min_length]\n" unless $in;
my $inseq = Bio::SeqIO->new(-file=>$in, -format=>'fasta');
my $out = $in . '.filtered.fasta';
my $outseq = Bio::SeqIO->new(-file=>">$out", -format=>'fasta');
while(my $s = $inseq->next_seq){
    $outseq->write_seq($s) if $s->length >= $min;
}
print "Wrote $out\n";
