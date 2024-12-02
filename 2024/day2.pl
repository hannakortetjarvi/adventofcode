#!/usr/bin/perl -l
use warnings;

# PART 1
my $INFILE;
open($INFILE, "inputs/day2.txt") or die "day2.txt doesn't open: $!";

my $safe_reports = 0;
while (my $line = <$INFILE>) {
    chomp $line;
    my @report = split / +/, $line;

    $safe_reports += check_safe(@report);
}
close $INFILE;

print "Answer to part 1: $safe_reports";

# PART 2
open($INFILE, "inputs/day2.txt") or die "day2.txt doesn't open: $!";

$safe_reports = 0;
while (my $line = <$INFILE>) {
    chomp $line;
    my @report = split / +/, $line;

    my $index = -1;
    foreach my $level (@report) {
        $index++;

        my $is_safe = check_safe(@report);
        if ($is_safe eq 1) {
            $safe_reports++;
            last;
        }

        my @temp_report = @report;
        splice(@temp_report, $index, 1);

        $is_safe = check_safe(@temp_report);
        if ($is_safe eq 1) {
            $safe_reports++;
            last;
        }
    }
}
close $INFILE;

print "Answer to part 2: $safe_reports";

sub check_safe {
    my @report = @_;
    my $safe = True;
    my $change;
    my $previous;

    foreach my $level (@report) {
        if (defined $previous) {
            if (defined $change) {
                my $result = $level - $previous;

                if ($change < 0 and (($result > -1 or $result < -3))) {
                    $safe = False;
                    last;
                }
                elsif ($change > 0 and (($result > 3 or $result < 1))) {
                    $safe = False;
                    last;
                }
            }
            else {
                $change = $level - $previous;
                if ($change > 3 or $change < -3 or $change eq 0) {
                    $safe = False;
                    last;
                }
            }
        }
        $previous = $level;
    }
    if ($safe eq True) {
        return 1;
    }
    return 0;
}