# Original Perl script by I. Petrov (RU) 2012
#!/usr/bin/env perl
# legacy script: compute GC content per sequence in a FASTA file
use strict;
use warnings;
use Bio::SeqIO;
my $in = shift or die "Usage: $0 sequences.fasta\n";
my $seqio = Bio::SeqIO->new(-file => $in, -format => 'fasta');
while (my $seq = $seqio->next_seq) {
    my $id = $seq->display_id;
    my $seqstr = $seq->seq;
    my $gc = (($seqstr =~ tr/GCgc//) / length($seqstr)) * 100;
    printf "%s\t%.2f\n", $id, $gc;
}
