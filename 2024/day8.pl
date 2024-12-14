#!/usr/bin/perl -l
use warnings;

my $INFILE;
open($INFILE, "inputs/day8.txt") or die "day8.txt doesn't open: $!";

my %satellites;
my $length = 0;
while (my $line = <$INFILE>) {
    chomp $line;

    for my $i (0 .. length($line) - 1) {
        if (substr($line, $i, 1) eq '.') {
            next;
        }

        my $char = substr($line, $i, 1);
        if (exists $satellites{$char}) {
            push @{$satellites{$char}}, "($i|$length)";
        } else {
            $satellites{$char} = ["($i|$length)"];
        }
    }
    $length++;
}

# ------ PART 1 ------

my %locations;
my $unique = 0;
for (keys %satellites){
	for $pair1 (@{$satellites{$_}}) {
        for $pair2 (@{$satellites{$_}}) {
            if ($pair1 eq $pair2) {
                next;
            }
            my @nums1 = $pair1 =~ /([0-9]+)/g; 
            my @nums2 = $pair2 =~ /([0-9]+)/g; 
            
            my $x;
            my $y;

            if ($nums1[0] == $nums2[0] ) {
                $x = $nums1[0];
            } else {
                $x = 2 * $nums1[0] - $nums2[0]
            }

            if ($nums1[1] == $nums2[1] ) {
                $y = $nums1[1];
            } else {
                $y = 2 * $nums1[1] - $nums2[1]
            }

            my $antinode = "($x|$y)";
            if ($x >= 0 && $x < $length && $y >= 0 && $y < $length) {
                if (!exists $locations{$antinode}) {
                    $locations{$antinode} = 1;
                    $unique++;
                }
            }
        }
    }
}

print "Answer to part 1: $unique";

# ------ PART 2 ------

my %locations_new;
$unique = 0;
for (keys %satellites){
	for $pair1 (@{$satellites{$_}}) {
        for $pair2 (@{$satellites{$_}}) {
            if ($pair1 eq $pair2) {
                next;
            }

            my @nums1 = $pair1 =~ /([0-9]+)/g; 
            my $x = $nums1[0];
            my $y = $nums1[1];
            
            
            while ($x >= 0 && $x < $length && $y >= 0 && $y < $length) {
                my @nums2 = $pair2 =~ /([0-9]+)/g; 
                my $antinode = "($x|$y)";

                if (!exists $locations_new{$antinode}) {
                    $locations_new{$antinode} = 1;
                    $unique++;
                }

                if ($nums1[0] == $nums2[0] ) {
                    $x = $nums1[0];
                } else {
                    $x += $nums1[0] - $nums2[0];
                }

                if ($nums1[1] == $nums2[1] ) {
                    $y = $nums1[1];
                } else {
                    $y += $nums1[1] - $nums2[1];
                }
            }
        }
    }
}

print "Answer to part 2: $unique";