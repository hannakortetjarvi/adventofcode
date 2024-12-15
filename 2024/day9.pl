#!/usr/bin/perl -l
use warnings;

my $INFILE;
open($INFILE, "inputs/day9_example.txt") or die "day9.txt doesn't open: $!";
my $line = <$INFILE>;

# ------ PART 1 ------

my @array = split(//, $line);
my @filesystem;
my $blocks = 0;
my $index = 0;
my $format_index = 0;

for my $letter (@array) {
    if ($index == 0) {
        $blocks = $letter;
        $index++;
        push @filesystem, ($format_index) x $blocks;
    } elsif ($index == 1) {
        my $free_space = $letter;
        $index--;
        push @filesystem, ('.') x $free_space;
        $format_index++;
    }
}

my @filesystem_prt_two = @filesystem;

my $dot_index = 0;
my $rev_index = $#filesystem;
while ($dot_index < $rev_index) {
    while ($dot_index < @filesystem && $filesystem[$dot_index] ne '.') {
        $dot_index++;
    }

    while ($rev_index > $dot_index && $filesystem[$rev_index] eq '.') {
        $rev_index--;
    }

    if ($dot_index < $rev_index) {
        $filesystem[$dot_index] = $filesystem[$rev_index];
        $filesystem[$rev_index] = '.';
    }
}

my $checksum = 0;
for my $i (0 .. $#filesystem) {
    my $char = $filesystem[$i];
    last if $char eq '.';
    $checksum += $i * $char;
}

print "Answer to part 1: $checksum";