class Solution {
    public int romanToInt(String s) {

        String[] h = s.split("");
        int result = 0;

        for (int i = h.length - 1; i >= 0; i--) {
            boolean subtract = false;

            if (i != h.length - 1){
                switch (h[i]){
                    case "I":
                        if (h[i + 1].equals("V") || h[i+1].equals("X")){
                            subtract = true;
                            result -= 1;
                        }
                        break;
                    case "X":
                        if (h[i + 1].equals("L") || h[i+1].equals("C")){
                            subtract = true;
                            result -= 10;
                        }
                        break;
                    case "C":
                        if (h[i + 1].equals("D") || h[i+1].equals("M")){
                            subtract = true;
                            result -= 100;
                        }
                        break;
                }
            }

            if (!subtract){

                switch (h[i]){
                    case "I":
                        result += 1;
                        break;
                    case "V":
                        result += 5;
                        break;
                    case "X":
                        result += 10;
                        break;
                    case "L":
                        result += 50;
                        break;
                    case "C":
                        result += 100;
                        break;
                    case "D":
                        result += 500;
                        break;
                    case "M":
                        result += 1000;
                        break;
                }
            }
        }
        return result;
    }
}