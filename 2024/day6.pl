#!/usr/bin/perl -l
use warnings;

# Read lines into matrix
my $INFILE;
open($INFILE, "inputs/day6.txt") or die "day6.txt doesn't open: $!";

my @obstacles;
my $length = 0;
my $start_x;
my $start_y;
while (my $line = <$INFILE>) {
    chomp $line;
    for my $i (0 .. length($line) - 1) {
        if (substr($line, $i, 1) eq '#') {
            my $pair = "($i|$length)";
            push @obstacles, $pair;
        }

        if (substr($line, $i, 1) eq '^') {
            $start_x = $i;
            $start_y = $length;
        }
    }
    $length++;
}

# ------ PART 1 ------

my %obstacle_hash = map { $_ => 1 } @obstacles;
my %visited;

$current_dir = 0;
$x = $start_x;
$y = $start_y;

$result = 0;
while ($x >= 0 && $x < $length && $y >= 0 && $y < $length) {
    my $pair = "($x|$y)";

    my ($next_x, $next_y) = ($x, $y);
    if ($current_dir == 0) {
        $next_y--;
    } elsif ($current_dir == 1) {
        $next_x++;
    } elsif ($current_dir == 2) {
        $next_y++;
    } else {
        $next_x--;
    }

    my $next_pair = "($next_x|$next_y)";
    if (exists $obstacle_hash{$next_pair}) {
        $current_dir = ($current_dir + 1) % 4;
        next;
    }

    if (!exists $visited{$pair}) {
        $visited{$pair} = 1; # Mark as visited
        $result++;
    }

    ($x, $y) = ($next_x, $next_y);
}

print "Answer to part 1: $result";

# ------ PART 2 ------

%obstacle_hash = map { $_ => 1 } @obstacles;

$result = 0;
foreach my $pair (%visited) {
    my %visited;
    my $current_dir = 0;
    $x = $start_x;
    $y = $start_y;

    $obstacle_hash{$pair} = 1;
    while ($x >= 0 && $x < $length && $y >= 0 && $y < $length) {
        my $loc = "($x|$y)|$current_dir";

        if ($visited{$loc}) {
            $result++;
            last;
        }

        my ($next_x, $next_y) = ($x, $y);
        if ($current_dir == 0) {
            $next_y--;
        } elsif ($current_dir == 1) {
            $next_x++;
        } elsif ($current_dir == 2) {
            $next_y++;
        } else {
            $next_x--;
        }

        my $next_pos = "($next_x|$next_y)";
        if (exists $obstacle_hash{$next_pos}) {
            $current_dir = ($current_dir + 1) % 4;
            next;
        }

        $visited{$loc} = 1; # Mark as visited

        ($x, $y) = ($next_x, $next_y);
    }

    delete $obstacle_hash{$pair};
}

print "Answer to part 2: $result";
