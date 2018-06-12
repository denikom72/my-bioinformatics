#!/usr/bin/env perl
# Basic SAM flag/statistics parser
use strict;
use warnings;
my $sam = shift or die "Usage: $0 file.sam\n";
open my $fh, '<', $sam or die $!;
my ($mapped,$total)=(0,0);
while(<$fh>){
    next if /^@/;
    $total++;
    my @f = split /\t/;
    $mapped++ if $f[2] ne '*';
}
print "Mapped: $mapped / $total\n";
