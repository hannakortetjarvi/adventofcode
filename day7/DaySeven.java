package day7;
import java.util.*;
import java.io.*;

public class DaySeven {

    public static Integer calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open " + e.getMessage());
            return -1;
        }

        List<String> currentLoc = new ArrayList(100);
        List<AbstractMap.SimpleEntry<String, Integer>> sizes = new ArrayList(100);
        Integer sum = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                String[] parts = s.split(" ");
                switch(parts[0]){
                    case "$":
                        if (parts[1].equals("ls")) continue;
                        if (parts[1].equals("cd") && !parts[2].equals("..")) {
                            currentLoc.add(parts[2]); 
                            AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(parts[2], 0);
                            sizes.add(pair);
                        }
                        if (parts[1].equals("cd") && parts[2].equals("..")) {
                            for (int i = 0; i < sizes.size(); i++) {
                                if (sizes.get(i).getKey() == currentLoc.get(currentLoc.size() - 2)) {
                                    for (int j = 0; j < sizes.size(); j++) {
                                        if (sizes.get(j).getKey() == currentLoc.get(currentLoc.size() - 1)) {
                                            String folder = sizes.get(i).getKey();
                                            Integer newSize = sizes.get(i).getValue() + sizes.get(j).getValue();
                                            AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(folder, newSize);
                                            sizes.set(i, pair);
                                            currentLoc.remove(currentLoc.size() - 1);
                                            break;
                                        }
                                    }
                                    break;
                                }
                            }
                        }
                    case "dir":
                        continue;
                    default:
                        for (int i = 0; i < sizes.size(); i++) {
                            if (sizes.get(i).getKey() == currentLoc.get(currentLoc.size() - 1)) {
                                String folder = sizes.get(i).getKey();
                                Integer newSize = sizes.get(i).getValue() + Integer.parseInt(parts[0]);
                                AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(folder, newSize);
                                sizes.set(i, pair);
                                break;
                            }
                        }
                }
            }
        } finally {
            while (currentLoc.size() > 1) {
                for (int i = 0; i < sizes.size(); i++) {
                    if (sizes.get(i).getKey() == currentLoc.get(currentLoc.size() - 2)) {
                        for (int j = 0; j < sizes.size(); j++) {
                            if (sizes.get(j).getKey() == currentLoc.get(currentLoc.size() - 1)) {
                                String folder = sizes.get(i).getKey();
                                Integer newSize = sizes.get(i).getValue() + sizes.get(j).getValue();
                                AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(folder, newSize);
                                sizes.set(i, pair);
                                currentLoc.remove(currentLoc.size() - 1);
                                break;
                            }
                        }
                        break;
                    }
                }
            }
            
            for (int i = 0; i < sizes.size(); i++) {
                if (sizes.get(i).getValue() <= 100000) sum += sizes.get(i).getValue();
            }
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

        Integer largest = 70000000;
        List<String> currentLoc = new ArrayList(100);
        List<AbstractMap.SimpleEntry<String, Integer>> sizes = new ArrayList(100);
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                String[] parts = s.split(" ");
                switch(parts[0]){
                    case "$":
                        if (parts[1].equals("ls")) continue;
                        if (parts[1].equals("cd") && !parts[2].equals("..")) {
                            currentLoc.add(parts[2]); 
                            AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(parts[2], 0);
                            sizes.add(pair);
                        }
                        if (parts[1].equals("cd") && parts[2].equals("..")) {
                            for (int i = 0; i < sizes.size(); i++) {
                                if (sizes.get(i).getKey() == currentLoc.get(currentLoc.size() - 2)) {
                                    for (int j = 0; j < sizes.size(); j++) {
                                        if (sizes.get(j).getKey() == currentLoc.get(currentLoc.size() - 1)) {
                                            String folder = sizes.get(i).getKey();
                                            Integer newSize = sizes.get(i).getValue() + sizes.get(j).getValue();
                                            AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(folder, newSize);
                                            sizes.set(i, pair);
                                            currentLoc.remove(currentLoc.size() - 1);
                                            break;
                                        }
                                    }
                                    break;
                                }
                            }
                        }
                    case "dir":
                        continue;
                    default:
                        for (int i = 0; i < sizes.size(); i++) {
                            if (sizes.get(i).getKey() == currentLoc.get(currentLoc.size() - 1)) {
                                String folder = sizes.get(i).getKey();
                                Integer newSize = sizes.get(i).getValue() + Integer.parseInt(parts[0]);
                                AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(folder, newSize);
                                sizes.set(i, pair);
                                break;
                            }
                        }
                }
            }
        } finally {
            while (currentLoc.size() > 1) {
                for (int i = 0; i < sizes.size(); i++) {
                    if (sizes.get(i).getKey() == currentLoc.get(currentLoc.size() - 2)) {
                        for (int j = 0; j < sizes.size(); j++) {
                            if (sizes.get(j).getKey() == currentLoc.get(currentLoc.size() - 1)) {
                                String folder = sizes.get(i).getKey();
                                Integer newSize = sizes.get(i).getValue() + sizes.get(j).getValue();
                                AbstractMap.SimpleEntry<String, Integer> pair = new AbstractMap.SimpleEntry<String, Integer>(folder, newSize);
                                sizes.set(i, pair);
                                currentLoc.remove(currentLoc.size() - 1);
                                break;
                            }
                        }
                        break;
                    }
                }
            }
            
            largest = sizes.get(0).getValue();
            Integer neededToDelete = 0;
            Integer space = 70000000 - largest;
            neededToDelete = 30000000 - space;
            for (int i = 0; i < sizes.size(); i++) {

                if (sizes.get(i).getValue() < largest && sizes.get(i).getValue() > neededToDelete) largest = sizes.get(i).getValue();
            }
            fi.close(); 
        }
        return largest;
    }

    public static void main(String args[]) {
        String fileInput = "input.txt";

        Integer firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        Integer secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
        
    }
}