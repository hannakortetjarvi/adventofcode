#!/usr/bin/perl -l
use warnings;

# Read lines into matrix
my $INFILE;
open($INFILE, "inputs/day6.txt") or die "day6.txt doesn't open: $!";

my @obstacles;
my $length = 0;
my $x;
my $y;
while (my $line = <$INFILE>) {
    chomp $line;
    for my $i (0 .. length($line) - 1) {
        if (substr($line, $i, 1) eq '#') {
            my $pair = $i . '|' . $length;
            push @obstacles, $pair;
        }

        if (substr($line, $i, 1) eq '^') {
            $x = $i;
            $y = $length;
        }
    }
    $length++;
}

# ------ PART 1 ------

my @directions = (0, 1, 2, 3); # up, right, down, left
my $current_dir = 0;
my @visited_locations;

my $result = 0;
while ($x >= 0 && $x <= $length && $y >= 0 && $y <= $length) {
    my $pair = $x . '|' . $y;
    push @visited_locations, $pair;

    my $next_x = $x;
    my $next_y = $y;
    if ($current_dir == 0) {
        $next_y--;
    }
    elsif ($current_dir == 1) {
        $next_x++;
    }
    elsif ($current_dir == 2) {
        $next_y++;
    }
    else {
        $next_x--;
    }

    my $pattern = "$next_x|$next_y";
    if ( grep {$_ eq $pattern} @obstacles ) {
        $current_dir = ($current_dir + 1) % 4;
        next;
    }

    if ( not(grep {$_ eq $pattern} @visited_locations) ) {
        push @visited_locations, $pair;
        $result++;
    }

    $x = $next_x;
    $y = $next_y;
}

print "Answer to part 1: $result";
#4662 TOO LOW

# ------ PART 2 ------
