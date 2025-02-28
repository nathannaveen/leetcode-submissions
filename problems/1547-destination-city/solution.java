class Solution {
    public String destCity(List<List<String>> paths) {
        String isEndDestination = "";
        for (int i = 0; i < paths.size(); i++) {
            isEndDestination = paths.get(i).get(1);
            boolean end = true;
            for (int j = 0; j < paths.size(); j++) {
                if (paths.get(j).get(0).equals(isEndDestination)){
                    end = false;
                    break;
                }
            }
            if (end){
                return isEndDestination;
            }
        }
        return isEndDestination;
    }
}