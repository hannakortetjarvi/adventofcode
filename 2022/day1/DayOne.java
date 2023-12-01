package day1;
import java.util.*;
import java.io.*;

public class DayOne {

    /**
     * Read file and calculate top elf's calories from the contents of file
     * @param file File that will be opened
     * @return Top calories
     */
    public static Integer calculateTopFromFile(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return 0;
        }

        int mostCalories = 0;
        int currentCalories = 0;
        int lines = 0;
        try {
            while (fi.hasNextLine()) {
                lines++;
                String s = fi.nextLine();
                try {
                    int num = Integer.parseInt(s);
                    currentCalories += num;
                } catch (NumberFormatException e) {
                    if (currentCalories > mostCalories) mostCalories = currentCalories;
                    currentCalories = 0;
                }
            }
        } finally {
            if (currentCalories > mostCalories) mostCalories = currentCalories;
            fi.close(); 
        }

        return mostCalories;
    }
    
    /**
     * Read file and calculate three top elves' calories from the contents of file
     * @param file File that will be opened
     * @return Total value of calories
     */
    public static Integer calculateTopThreeFromFile(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return 0;
        }

        int mostCalories = 0;
        int[] topCalories = {0,0,0};
        int currentCalories = 0;
        int lines = 0;
        try {
            while (fi.hasNextLine()) {
                lines++;
                String s = fi.nextLine();
                try {
                    int num = Integer.parseInt(s);
                    currentCalories += num;
                } catch (NumberFormatException e) {
                    topCalories = isLarger(topCalories, currentCalories);
                    currentCalories = 0;
                }
            }
        } finally {
            if (currentCalories > mostCalories) mostCalories = currentCalories;
            fi.close(); 
        }

        int totalCalories = 0;
        for (int i = 0; i < topCalories.length; i++) {
            totalCalories += topCalories[i];
        }
        return totalCalories;
    }

    /**
     * Checks if the new value is larger than values in array
     * @param arr Array with top values
     * @param val Checked value
     * @return New array with top values
     */
    public static int[] isLarger(int[] arr, int val) {
        Arrays.sort(arr);
        if (arr[0] < val) arr[0] = val;
        return arr;
    }

    public static void main(String args[]) {
        String fileInput = "input.txt";

        int firstResult = calculateTopFromFile(fileInput);
        System.out.println("Top elf is carrying "+ firstResult + " calories");

        int secondResult = calculateTopThreeFromFile(fileInput);
        System.out.println("Three top elfs are carrying "+ secondResult + " calories");
    }
}