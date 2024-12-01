#!/usr/bin/perl
use warnings;

# Read while into two arrays
my $INFILE;
open($INFILE, "inputs/day1.txt") or die "day1.txt doesn't open: $!";

my (@pair1, @pair2);
while (my $line = <$INFILE>) {
    chomp $line;
    my ($value1, $value2) = split / +/, $line;
    push @pair1, $value1;
    push @pair2, $value2;
}
close $INFILE;

# Sort rows
@pair1 = sort {$a <=> $b} @pair1;
@pair2 = sort {$a <=> $b} @pair2;

# Count distance between each sorted pair
my $distance = 0;
for my $i (0 .. $#pair1) {
    $distance = $distance + (abs($pair1[$i] - $pair2[$i]));
}

print "Answer to part 1: $distance\n";

# Count the similarity
my $similarity = 0;
for my $i (0 .. $#pair2) {
    my $count = scalar grep { $_ == $pair1[$i] } @pair2;
    $similarity += $count * $pair1[$i];
}

print "Answer to part 2: $similarity\n";