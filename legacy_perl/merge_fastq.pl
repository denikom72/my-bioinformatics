#!/usr/bin/env perl
# Merge interleaved FASTQ or paired files (simple concatenation utility)
use strict;
use warnings;
my ($f1,$f2,$out) = @ARGV;
die "Usage: $0 r1.fastq r2.fastq out.fastq\n" unless $f1 && $f2 && $out;
open my $o, '>', $out or die $!;
open my $a, '<', $f1 or die $!;
open my $b, '<', $f2 or die $!;
while(<$a>){ print $o $_; $_ = <$b>; print $o $_ if defined $_; }
close $a; close $b; close $o;
print "Merged to $out\n";
