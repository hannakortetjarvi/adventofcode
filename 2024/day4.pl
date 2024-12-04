#!/usr/bin/perl -l
use warnings;
no warnings 'uninitialized';

# Read lines into matrix
my $INFILE;
open($INFILE, "inputs/day4.txt") or die "day4.txt doesn't open: $!";

my @crossword;
my $length = 0;
while (my $line = <$INFILE>) {
    chomp $line;
    my @letters = split //, $line;
    push @crossword, \@letters;
    $length++;
}
close $INFILE;

# ------ PART 1 ------

my $result = 0;
for my $i (0 .. $#crossword) {
    for my $j (0 .. $#{$crossword[$i]}) {
        if ($j + 2 < $length) {
            my $right = $crossword[$i][$j] . $crossword[$i][$j + 1] . $crossword[$i][$j + 2] . $crossword[$i][$j + 3];
            if ($right eq 'XMAS' or $right eq 'SAMX') {$result++;}
        }

        if ($i + 2 < $length) {
            my $down = $crossword[$i][$j] . $crossword[$i + 1][$j] . $crossword[$i + 2][$j] . $crossword[$i + 3][$j];
            if ($down eq 'XMAS' or $down eq 'SAMX') {$result++;}
        }

        if ($i + 2 < $length and $j + 2 < $length) {
            my $south_east = $crossword[$i][$j] . $crossword[$i + 1][$j + 1] . $crossword[$i + 2][$j + 2] . $crossword[$i + 3][$j + 3];
            if ($south_east eq 'XMAS' or $south_east eq 'SAMX') {$result++;}
        }

        if ($i + 2 < $length and $j - 3 >= 0) {
            my $south_west = $crossword[$i][$j] . $crossword[$i + 1][$j - 1] . $crossword[$i + 2][$j - 2] . $crossword[$i + 3][$j - 3];
            if ($south_west eq 'XMAS' or $south_west eq 'SAMX') {$result++;}
        }
    }
}

print "Answer to part 1: $result";

# ------ PART 2 ------

$result = 0;
for my $i (1 .. $#crossword - 1) {
    for my $j (1 .. $#{$crossword[$i]} - 1) {
        if ($crossword[$i][$j] ne 'A') {next;}

        my $ne = $crossword[$i - 1][$j + 1];
        my $nw = $crossword[$i - 1][$j - 1];
        my $se = $crossword[$i + 1][$j + 1];
        my $sw = $crossword[$i + 1][$j - 1];

        if ((($ne . $sw) eq 'SM' or ($ne . $sw) eq 'MS') and (($nw . $se) eq 'SM' or ($nw . $se) eq 'MS')) {$result++;}
    }
}

print "Answer to part 2: $result";