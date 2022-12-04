package day4;
import java.util.*;
import java.io.*;

public class DayFour {

    public static Integer calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return 0;
        }

        int pairs = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                String[] vals = s.split(",");
                pairs += compare(vals[0], vals[1]);
            }
        } finally {
            fi.close(); 
        }

        return pairs;
    }

    public static int compare(String first, String second) {
        String[] firstNums = first.split("-");
        String[] secondNums = second.split("-");
        if ((Integer.parseInt(firstNums[0]) <= Integer.parseInt(secondNums[0]) && Integer.parseInt(secondNums[0]) <= Integer.parseInt(firstNums[1]))
            && (Integer.parseInt(firstNums[0]) <= Integer.parseInt(secondNums[1]) && Integer.parseInt(secondNums[1]) <= Integer.parseInt(firstNums[1])))
            return 1;
        if ((Integer.parseInt(secondNums[0]) <= Integer.parseInt(firstNums[0]) && Integer.parseInt(firstNums[0]) <= Integer.parseInt(secondNums[1]))
            && (Integer.parseInt(secondNums[0]) <= Integer.parseInt(firstNums[1]) && Integer.parseInt(firstNums[1]) <= Integer.parseInt(secondNums[1])))
            return 1;
        return 0;
    }

    public static Integer calculateSecond(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return 0;
        }

        int pairs = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                String[] vals = s.split(",");
                pairs += compareNumbers(vals[0], vals[1]);
            }
        } finally {
            fi.close(); 
        }

        return pairs;
    }

    public static int compareNumbers(String first, String second) {
        String[] firstNums = first.split("-");
        String[] secondNums = second.split("-");
        if ((Integer.parseInt(firstNums[0]) <= Integer.parseInt(secondNums[0]) && Integer.parseInt(secondNums[0]) <= Integer.parseInt(firstNums[1])))
            return 1;
        if ((Integer.parseInt(firstNums[0]) <= Integer.parseInt(secondNums[1]) && Integer.parseInt(secondNums[1]) <= Integer.parseInt(firstNums[1])))
            return 1;
        if ((Integer.parseInt(secondNums[0]) <= Integer.parseInt(firstNums[0]) && Integer.parseInt(firstNums[0]) <= Integer.parseInt(secondNums[1])))
            return 1;
        if ((Integer.parseInt(secondNums[0]) <= Integer.parseInt(firstNums[1]) && Integer.parseInt(firstNums[1]) <= Integer.parseInt(secondNums[1])))
            return 1;
        return 0;
    }


    public static void main(String args[]) {
        String fileInput = "input.txt";

        int firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        int secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
        
    }
}