package day3;
import java.util.*;
import java.io.*;

public class DayThree {

    public static Integer calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return 0;
        }

        int sum = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                String[] vals = {s.substring(0, s.length()/2),s.substring(s.length()/2)};
                sum += bothComp(vals[0], vals[1]);
            }
        } finally {
            fi.close(); 
        }

        return sum;
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

        int sum = 0;
        List<String> arr = new ArrayList(3);
        int conts = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                arr.add(s);
                conts += 1;

                if (conts != 3) continue;
                sum += common(arr.get(0), arr.get(1), arr.get(2));
                arr.clear();
                conts = 0;
            }
        } finally {
            fi.close(); 
        }

        return sum;
    }

    public static int common(String first, String second, String third) {
        for (int i = 0; i < first.length(); i++) {
            if (second.indexOf(first.charAt(i)) != -1 && third.indexOf(first.charAt(i)) != -1) {
                return calcPri(first.charAt(i));
            }
        }
        return 0;
    }

    public static int bothComp(String first, String second) {
        int total = 0;
        List<Character> found = new ArrayList<>(first.length());
        for (int i = 0; i < first.length(); i++) {
            if (second.indexOf(first.charAt(i)) != -1 && found.contains(first.charAt(i)) == false) {
                total += calcPri(first.charAt(i));
                found.add(first.charAt(i));
            }
        }
        return total;
    }

    public static int calcPri(char x) {
        int asc = (int) x;
        if (97 <= asc && asc <= 122) return asc - 96;
        else if (65 <= asc && asc <= 90) return asc - 65 + 27;
        else return 0;
    }

    public static void main(String args[]) {
        String fileInput = "input.txt";

        int firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        int secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
        
    }
}