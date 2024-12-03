#!/usr/bin/perl -l
use warnings;

# PART 1
my $INFILE;
open($INFILE, "inputs/day3.txt") or die "day3.txt doesn't open: $!";

my $result = 0;
while (my $line = <$INFILE>) {
    chomp $line;
    my $collected_letters = "";
    foreach $letter (split //, $line) {
        $collected_letters .= $letter;
        if ( $collected_letters =~ m/mul\((\d+),(\d+)\)$/ ) {
            $result += ($1 * $2);
        }
    }
}
close $INFILE;

print "Answer to part 1: $result";

# PART 2
open($INFILE, "inputs/day3.txt") or die "day3.txt doesn't open: $!";

$result = 0;
my $enabled = True;
while (my $line = <$INFILE>) {
    chomp $line;
    my $collected_letters = "";
    foreach $letter (split //, $line) {
        $collected_letters .= $letter;

        if ( $collected_letters =~ m/do\(\)$/ ) { # Enable mult
            $enabled = True;
            next;
        }
        
        if ( $collected_letters =~ m/don\'t\(\)$/ ) { # Disable mult
            $enabled = False;
            next;
        }

        if ( $enabled eq True and $collected_letters =~ m/mul\((\d+),(\d+)\)$/ ) {
            $result += ($1 * $2);
        }
    }
}
close $INFILE;

print "Answer to part 2: $result";