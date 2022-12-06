package day6;
import java.util.*;
import java.io.*;

public class DaySix {

    public static int calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return -1;
        }

        List<Character> chars = new ArrayList(4);
        int index = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                for (int i = 0; i < s.length(); i++) {
                    System.out.println(chars);
                    if (chars.size() != 4) {
                        chars.add(s.charAt(i));
                        continue;
                    }
                    chars.remove(0);
                    chars.add(s.charAt(i));
                    System.out.println(chars);
                    if (checkIfAllUnique(chars)) {
                        index = i + 1;
                        break;
                    }
                }
            }
        } finally {
            fi.close(); 
        }
        return index;
    }

    public static int calculateSecond(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return -1;
        }

        List<Character> chars = new ArrayList(14);
        int index = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                for (int i = 0; i < s.length(); i++) {
                    System.out.println(chars);
                    if (chars.size() != 14) {
                        chars.add(s.charAt(i));
                        continue;
                    }
                    chars.remove(0);
                    chars.add(s.charAt(i));
                    System.out.println(chars);
                    if (checkIfAllUnique(chars)) {
                        index = i + 1;
                        break;
                    }
                }
            }
        } finally {
            fi.close(); 
        }
        return index;
    }

    public static boolean checkIfAllUnique(List<Character> list) {
        for (int i = 0; i < list.size(); i++) {
            for (int j = 0; j < list.size(); j++) {
                if (i == j) continue;
                if (list.get(i) == list.get(j)) return false;
            }
        }
        return true;
    }

    public static void main(String args[]) {
        String fileInput = "input.txt";

        int firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        int secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
        
    }
}