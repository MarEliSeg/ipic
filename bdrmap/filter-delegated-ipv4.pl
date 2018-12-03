#!/usr/bin/perl
#
# $Id$

use strict;
use warnings;

foreach my $delegated (@ARGV)
{
    if($delegated =~ /\.bz2$/) {
	open(STATS, "bzcat $delegated |") or die "could not open $delegated\n";
    } elsif($delegated =~ /\.gz$/) {
	open(STATS, "gzcat $delegated |") or die "could not open $delegated\n";
    } else {
	open(STATS, $delegated) or die "could not open $delegated\n";
    }
    while(<STATS>)
    {
	next if(/^#/);
	chomp;
	my @bits = split(/\|/, $_);
	next if($bits[2] ne "ipv4" || scalar(@bits) < 7);
	print "$_\n";
    }
    close STATS;
}
