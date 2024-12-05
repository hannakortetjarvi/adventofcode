#!/usr/bin/perl -l
use warnings;

# Read lines into matrix
my $INFILE;
open($INFILE, "inputs/day5.txt") or die "day5.txt doesn't open: $!";

my @pairs;
while (my $line = <$INFILE>) {
    chomp $line;

    if ($line eq '') {
        last;
    }
    push @pairs, $line;
}

# ------ PART 1 ------

my $result = 0;
my @incorrect;
while (my $line = <$INFILE>) {
    chomp $line;
    my $correct = 1;
    my @numbers = split ',', $line;
    for my $i (0 .. $#numbers) {
        foreach my $pair (@pairs) {
            my @values = split /\|/, $pair;
            if ($values[0] eq $numbers[$i]) {
                if ( $line =~ /\Q$values[1]\E.*\Q$values[0]\E/ ) {
                    $correct = 0;
                    push @incorrect, $line;
                    last;
                }
            } 
        }
        if ($correct == 0) {
            last;
        }
    }
    if ($correct == 1) {
        $result += $numbers[(scalar(@numbers) - 1) / 2];
    }
}
close $INFILE;

print "Answer to part 1: $result";

# ------ PART 2 ------

$result = 0;
for my $i (0 .. $#incorrect) {
    my @numbers = split ',', $incorrect[$i];
    for my $x (0 .. $#numbers) {
        for my $y (0 .. $#numbers) {
            if ($x == $y) {
                next; # Skip same indexes
            }
            my $pattern = "$numbers[$y]|$numbers[$x]";
            if ( grep {$_ eq $pattern} @pairs ) {
                my $temp = $numbers[$x];
                $numbers[$x] = $numbers[$y];
                $numbers[$y] = $temp;
            }
        }
    }
    $result += $numbers[(scalar(@numbers) - 1) / 2];
}

print "Answer to part 2: $result";