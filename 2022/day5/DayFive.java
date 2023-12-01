package day5;
import java.util.*;
import java.io.*;

public class DayFive {

    public static List<Character> calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return null;
        }

        List<Character> onTop = new ArrayList(10);
        List<Character> firstLine = new ArrayList(100);
        List<Character> secondLine = new ArrayList(100);
        List<Character> thirdLine = new ArrayList(100);
        List<List<Character>> lines = new ArrayList(10);
        boolean startMove = false;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                if (s.length() == 0) {
                    for (int i = 0; i < lines.size(); i++) {
                        List<Character> line = lines.get(i);
                        Collections.reverse(line);
                        lines.set(i, line);
                    }
                    continue;
                }
                if (startMove) {
                    String[] rules = s.split(" ");
                    int count = Integer.parseInt(rules[1]);
                    int from = Integer.parseInt(rules[3]) - 1;
                    int to = Integer.parseInt(rules[5]) - 1;

                    for (int i = 0; i < count; i++) {
                        if (lines.get(from).size() == 0) break;
                        List<Character> fromList = lines.get(from);
                        List<Character> toList = lines.get(to);

                        int length = fromList.size();
                        Character x = fromList.get(length - 1);
                        fromList.remove(length - 1);
                        toList.add(x);

                        lines.set(from, fromList);
                        lines.set(to, toList);
                    }
                }
                else if (s.substring(1, 2).equals("1")) {System.out.println("Third stop"); startMove = true; continue;}
                else {
                    for (int i = 1; i < s.length();) {
                        if (lines.size() < (i - 1) / 4 + 1) {
                            List<Character> line = new ArrayList(100);
                            String part = s.substring(i, i+1);
                            i += 4;
                            System.out.println(part);
                            if (!part.equals(" ")) line.add(part.charAt(0));
                            lines.add(line);
                        }
                        else {
                            List<Character> line = lines.get((i - 1) / 4);
                            String part = s.substring(i, i+1);
                            if (!part.equals(" ")) line.add(part.charAt(0));
                            i += 4;
                        }
                    }
                }
            }
        } finally {
            for (int i = 0; i < lines.size(); i++) {
                List<Character> line = lines.get(i);
                if (line.size() == 0) onTop.add('0');
                else onTop.add(line.get(line.size() - 1));
            }
            fi.close(); 
        }

        return onTop;
    }

    public static List<Character> calculateSecond(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return null;
        }

        List<Character> onTop = new ArrayList(10);
        List<Character> firstLine = new ArrayList(100);
        List<Character> secondLine = new ArrayList(100);
        List<Character> thirdLine = new ArrayList(100);
        List<List<Character>> lines = new ArrayList(10);
        boolean startMove = false;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                if (s.length() == 0) {
                    for (int i = 0; i < lines.size(); i++) {
                        List<Character> line = lines.get(i);
                        Collections.reverse(line);
                        lines.set(i, line);
                    }
                    continue;
                }
                if (startMove) {
                    String[] rules = s.split(" ");
                    int count = Integer.parseInt(rules[1]);
                    int from = Integer.parseInt(rules[3]) - 1;
                    int to = Integer.parseInt(rules[5]) - 1;

                    List<Character> fromList = lines.get(from);
                    int start = fromList.size() - count;
                    if (count >= fromList.size()) start = 0;
                    List<Character> sub = fromList.subList(start, fromList.size());
                    List<Character> toList = lines.get(to);
                    toList.addAll(sub);
                    for (int i = 0; i < count; i++) {
                        fromList.remove(fromList.size() - 1);
                    }
                    lines.set(from, fromList);
                    lines.set(to, toList);
                }
                else if (s.substring(1, 2).equals("1")) {startMove = true; continue;}
                else {
                    for (int i = 1; i < s.length();) {
                        if (lines.size() < (i - 1) / 4 + 1) {
                            List<Character> line = new ArrayList(100);
                            String part = s.substring(i, i+1);
                            i += 4;
                            if (!part.equals(" ")) line.add(part.charAt(0));
                            lines.add(line);
                        }
                        else {
                            List<Character> line = lines.get((i - 1) / 4);
                            String part = s.substring(i, i+1);
                            if (!part.equals(" ")) line.add(part.charAt(0));
                            i += 4;
                        }
                    }
                }
            }
        } finally {
            for (int i = 0; i < lines.size(); i++) {
                List<Character> line = lines.get(i);
                if (line.size() == 0) onTop.add('0');
                else onTop.add(line.get(line.size() - 1));
            }
            fi.close(); 
        }

        return onTop;
    }

    public static void main(String args[]) {
        String fileInput = "input.txt";

        List<Character> firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        List<Character> secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
        
    }
}