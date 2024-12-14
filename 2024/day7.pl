#!/usr/bin/perl -l
use warnings;

my $INFILE;
open($INFILE, "inputs/day7.txt") or die "day7.txt doesn't open: $!";

my $calibration = 0;
while (my $line = <$INFILE>) {
    chomp $line;

    my ($total, $calc) = split ':', $line;
    $calc =~ s/^\s+//;
    my @nums = split / +/, $calc;

    my @results;
    foreach my $num (@nums) {
        if (scalar(@results) == 0) {
            push @results, $num;
            next;
        }

        my @new_results;
        for my $i (0 .. $#results) {
            my $sum = $results[$i] + $num;
            my $mult = $results[$i] * $num;
            push @new_results, $sum;
            push @new_results, $mult;
        }
        @results = @new_results;
    }

    foreach my $result (@results) {
        if ($result == $total) {
            $calibration += $result;
            last;
        }
    }
}

print "Answer to part 1: $calibration\n";

# ----------------------------

open($INFILE, "inputs/day7.txt") or die "day7.txt doesn't open: $!";
$calibration = 0;
while (my $line = <$INFILE>) {
    chomp $line;

    my ($total, $calc) = split ':', $line;
    $calc =~ s/^\s+//;
    my @nums = split / +/, $calc;

    my @results;
    foreach my $num (@nums) {
        if (scalar(@results) == 0) {
            push @results, $num;
            next;
        }

        my @new_results;
        for my $i (0 .. $#results) {
            my $sum = $results[$i] + $num;
            my $mult = $results[$i] * $num;
            my $conc = $results[$i] . $num;
            push @new_results, $sum;
            push @new_results, $mult;
            push @new_results, $conc;
        }
        @results = @new_results;
    }

    foreach my $result (@results) {
        if ($result == $total) {
            $calibration += $result;
            last;
        }
    }
}

print "Answer to part 2: $calibration";