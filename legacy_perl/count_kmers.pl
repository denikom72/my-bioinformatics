#!/usr/bin/env perl
# Count k-mers in FASTA sequences
use strict;
use warnings;
use Bio::SeqIO;
my ($in, $k) = @ARGV;
$k ||= 5;
my $seqio = Bio::SeqIO->new(-file=>$in, -format=>'fasta');
my %counts;
while(my $r = $seqio->next_seq){
    my $s = $r->seq;
    for(my $i=0;$i<=length($s)-$k;$i++){
        $counts{substr($s,$i,$k)}++;
    }
}
foreach my $kmer (sort keys %counts){
    print "$kmer\t$counts{$kmer}\n";
}
