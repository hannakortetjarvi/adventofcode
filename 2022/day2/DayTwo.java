package day1;
import java.util.*;
import java.io.*;

public class DayTwo {

    public static Integer calculateFirst(String file) {
        Scanner fi;
        try {
            File readFile = new File(file);
            fi = new Scanner(readFile);
        } catch (FileNotFoundException e) {
            System.err.println("File doesn't open" + e.getMessage());
            return 0;
        }

        int totalPoints = 0;
        int lines = 0;
        try {
            while (fi.hasNextLine()) {
                lines++;
                String s = fi.nextLine();
                String[] vals = s.split(" ");
                totalPoints += calc(getOpponent(vals[0]), getOwn(vals[1]));
            }
        } finally {
            fi.close(); 
        }

        return totalPoints;
    }

    public static String getOpponent(String input) {
        switch(input) {
            case "A":
                return "Rock";
            case "B":
                return "Paper";
            case "C":
                return "Scissors";
        }
        return "";
    }

    public static String getOwn(String input) {
        switch(input) {
            case "X":
                return "Rock";
            case "Y":
                return "Paper";
            case "Z":
                return "Scissors";
        }
        return "";
    }

    public static int calc(String opp, String own) {
        int points = 0;
        if (opp.equals(own)) {points += 3;}
        else if ((opp == "Rock" && own == "Paper") || 
            (opp == "Paper" && own == "Scissors") ||
            (opp == "Scissors" && own == "Rock"))
            {points += 6;}
        if (own.equals("Rock")) {points += 1;}
        else if (own.equals("Paper")) {points += 2;}
        else {points += 3;}
        return points;
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

        int totalPoints = 0;
        try {
            while (fi.hasNextLine()) {
                String s = fi.nextLine();
                String[] vals = s.split(" ");
                String opponent = getOpponent(vals[0]);
                totalPoints += calc(opponent, getOwnFromResult(opponent, vals[1]));
            }
        } finally {
            fi.close(); 
        }

        return totalPoints;
    }

    public static String getOwnFromResult(String opp, String own) {
        switch(own) {
            // Lose
            case "X":
                if (opp.equals("Rock")) return "Scissors";
                else if (opp.equals("Paper")) return "Rock";
                else return "Paper";
            // Draw
            case "Y":
                return opp;
            // Win
            case "Z":
                if (opp.equals("Rock")) return "Paper";
                else if (opp.equals("Paper")) return "Scissors";
                else return "Rock";
        }
        return "";
    }


    public static void main(String args[]) {
        String fileInput = "input.txt";

        int firstResult = calculateFirst(fileInput);
        System.out.println("First result: "+ firstResult);
        
        int secondResult = calculateSecond(fileInput);
        System.out.println("Second result: "+ secondResult);
        
    }
}