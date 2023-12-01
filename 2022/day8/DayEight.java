package day8;
import java.util.*;
import java.io.*;

public class DayEight {

    public static Integer calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open " + e.getMessage());
            return -1;
        }

        List<List<Integer>> trees = new ArrayList(100);
        Integer sum = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                List<Integer> row = new ArrayList(s.length());
                for (int i = 0; i < s.length(); i++) {
                    row.add(Character.getNumericValue(s.charAt(i)));
                }
                trees.add(row);
            }
        } finally {
            List<AbstractMap.SimpleEntry<Integer, Integer>> calculated = new ArrayList(1000);
            Integer columnSize = trees.get(0).size();
            for (int i = 0; i < trees.size(); i++) {
                Integer before = -1;
                Integer after = -1;
                for (int j = 0; j < columnSize; j++) {
                    Integer left = trees.get(i).get(j);
                    Integer right = trees.get(i).get(columnSize - 1 - j);
                    AbstractMap.SimpleEntry<Integer, Integer> coord1 = new AbstractMap.SimpleEntry<Integer,Integer>(i, j);
                    AbstractMap.SimpleEntry<Integer, Integer> coord2 = new AbstractMap.SimpleEntry<Integer,Integer>(i, columnSize - 1 - j);
                    if (i == 0 || j == 0 || i == trees.size() - 1 || j == columnSize - 1) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        before = left;
                        after = right;
                        continue;
                    }
                    if (left > before) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        before = left;
                    }
                    if (right > after) {
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        after = right;
                    }
                }
            }

            for (int i = 0; i < columnSize; i++) {
                Integer before = -1;
                Integer after = -1;
                for (int j = 0; j < trees.size(); j++) {
                    Integer up = trees.get(j).get(i);
                    Integer down = trees.get(trees.size() - 1 - j).get(i);
                    AbstractMap.SimpleEntry<Integer, Integer> coord1 = new AbstractMap.SimpleEntry<Integer,Integer>(j, i);
                    AbstractMap.SimpleEntry<Integer, Integer> coord2 = new AbstractMap.SimpleEntry<Integer,Integer>(trees.size() - 1 - j, i);
                    if (i == 0 || j == 0 || i == columnSize - 1 || j == trees.size() - 1) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        before = up;
                        after = down;
                        continue;
                    }
                    if (up > before) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        before = up;
                    }
                    if (down > after) {
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        after = down;
                    }
                }
            }

            sum = calculated.size();
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
            System.err.println("File doesn't open " + e.getMessage());
            return -1;
        }

        List<List<Integer>> trees = new ArrayList(100);
        Integer highest = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                List<Integer> row = new ArrayList(s.length());
                for (int i = 0; i < s.length(); i++) {
                    row.add(Character.getNumericValue(s.charAt(i)));
                }
                trees.add(row);
            }
        } finally {
            List<AbstractMap.SimpleEntry<Integer, Integer>> calculated = new ArrayList(1000);
            Integer columnSize = trees.get(0).size();
            for (int i = 0; i < trees.size(); i++) {
                Integer before = -1;
                Integer after = -1;
                for (int j = 0; j < columnSize; j++) {
                    Integer left = trees.get(i).get(j);
                    Integer right = trees.get(i).get(columnSize - 1 - j);
                    AbstractMap.SimpleEntry<Integer, Integer> coord1 = new AbstractMap.SimpleEntry<Integer,Integer>(i, j);
                    AbstractMap.SimpleEntry<Integer, Integer> coord2 = new AbstractMap.SimpleEntry<Integer,Integer>(i, columnSize - 1 - j);
                    if (i == 0 || j == 0 || i == trees.size() - 1 || j == columnSize - 1) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        before = left;
                        after = right;
                        continue;
                    }
                    if (left > before) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        before = left;
                    }
                    if (right > after) {
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        after = right;
                    }
                }
            }

            for (int i = 0; i < columnSize; i++) {
                Integer before = -1;
                Integer after = -1;
                for (int j = 0; j < trees.size(); j++) {
                    Integer up = trees.get(j).get(i);
                    Integer down = trees.get(trees.size() - 1 - j).get(i);
                    AbstractMap.SimpleEntry<Integer, Integer> coord1 = new AbstractMap.SimpleEntry<Integer,Integer>(j, i);
                    AbstractMap.SimpleEntry<Integer, Integer> coord2 = new AbstractMap.SimpleEntry<Integer,Integer>(trees.size() - 1 - j, i);
                    if (i == 0 || j == 0 || i == columnSize - 1 || j == trees.size() - 1) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        before = up;
                        after = down;
                        continue;
                    }
                    if (up > before) {
                        if (!calculated.contains(coord1)) calculated.add(coord1);
                        before = up;
                    }
                    if (down > after) {
                        if (!calculated.contains(coord2)) calculated.add(coord2);
                        after = down;
                    }
                }
            }

            for (int i = 0; i < calculated.size(); i++) {
                Integer y = calculated.get(i).getKey();
                Integer x = calculated.get(i).getValue();

                Integer orig = trees.get(y).get(x);

                Integer left = 0;
                System.out.println(y + " " + x);
                for (int j = 1; j < trees.get(y).size(); j++) {
                    Integer check = x - j;
                    if (check <= -1) break;
                    Integer val = trees.get(y).get(check);
                    if (val >= orig) {
                        left++;
                        break;
                    }
                    else left++;
                }

                Integer right = 0;
                for (int j = 1; j < trees.get(y).size(); j++) {
                    Integer check = x + j;
                    if (check >= trees.get(y).size()) break;
                    Integer val = trees.get(y).get(check);
                    if (val >= orig) {
                        right++;
                        break;
                    }
                    else right++;
                }

                Integer up = 0;
                for (int j = 1; j < trees.size(); j++) {
                    Integer check = y - j;
                    if (check <= -1) break;
                    Integer val = trees.get(check).get(x);
                    if (val >= orig) {
                        up++;
                        break;
                    }
                    else up++;
                }

                Integer down = 0;
                for (int j = 1; j < trees.size(); j++) {
                    Integer check = y + j;
                    if (check >= trees.size()) break;
                    Integer val = trees.get(check).get(x);
                    if (val >= orig) {
                        down++;
                        break;
                    }
                    else down++;
                }

                Integer calc = left * right * up * down;
                if (calc > highest) highest = calc;
            }
            fi.close(); 
        }
        return highest;
    }

    public static void main(String args[]) {
        String fileInput = "input.txt";

        Integer firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        Integer secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
    }
}