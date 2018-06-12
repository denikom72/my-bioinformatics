#!/usr/bin/env perl
# Create a tiny HTML summary report from tabular stats
use strict; use warnings;
my $in = shift or die "Usage: $0 stats.tsv\n";
open my $fh,'<',$in or die $!;
print "<html><body><table>\n";
while(<$fh>){
    chomp; my @c=split/\t/;
    print "<tr>" . join('', map {"<td>$_</td>"} @c) . "</tr>\n";
}
print "</table></body></html>\n";
